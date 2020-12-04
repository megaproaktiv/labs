package labdeploy

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/megaproaktiv/cfdl"
)

const labUser = "awsstudent"

// AccessKey short access key structure	
type AccessKey struct {
	AccessKeyID string
	SecretAccessKey string
}	

var keypair AccessKey

var (
    lowerCharSet   = "abcdedfghijklmnopqrstuvwxyz"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*§(){}"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)


//go:generate moq -out user_moq_test.go . IAMInterface


// IAMInterface used iam actions
type IAMInterface interface {
	AttachUserPolicy(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error)
	DetachUserPolicy(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error)

	CreateAccessKey(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error)
	DeleteAccessKey(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error)
	ListAccessKeys(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)

	CreateLoginProfile(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error)
	DeleteLoginProfile(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error)



	ListUsers(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
	CreateUser(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error)
	DeleteUser(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error)

	ListPolicies(ctx context.Context,params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) //
	CreatePolicy(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error)
	DeletePolicy(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error)
}

// **** Stacks ***

// **** User ***

// CreateUserIfnotExist awsstudent go
func CreateUserIfnotExist(client IAMInterface) {
	log.Println("Create User")
	password := generateMyPassword()
	if !studentExists(client) {
		_, err :=client.CreateUser(context.TODO(), &iam.CreateUserInput{
			UserName: aws.String(labUser),
		})
		if err != nil {
			log.Println("Error creating user: ",err)
		}
		_, err = client.CreateLoginProfile(context.TODO(), &iam.CreateLoginProfileInput{
			UserName: aws.String(labUser),
			Password: aws.String(password),
			PasswordResetRequired: false,
		})
	
		if err != nil {
			log.Println("Error creating login profile: ",err)
		}else{
			cfdl.Logger.Info("Created login Password: ",password)
		}

		f, err := os.Create("password.txt")
		if err != nil {
			cfdl.Logger.Error("File creation error: ",err)
		}
		err = os.Chmod(sshKeyFileName(), 0600)
		if err != nil {
			cfdl.Logger.Error("File change permission error: ",err)
		}
		_, err = f.WriteString(password)
	
		if err != nil {
			cfdl.Logger.Error("Credentials file write error: ",err)
			f.Close()
		}
		err = f.Close()
		if err != nil {
			cfdl.Logger.Error("Credentials file close error: ",err)			
		}	
	}
}

// CreateAccessKey create key
func CreateAccessKey(client IAMInterface) {
	log.Println("Create Access Key")
	params := &iam.CreateAccessKeyInput{
		UserName: aws.String(labUser),
	}
	response, err := client.CreateAccessKey(context.TODO(), params)
	if err != nil {
		panic(err)
	}
	var key,secret string
	key = fmt.Sprintf("%s%s","export AWS_ACCESS_KEY_ID=",*response.AccessKey.AccessKeyId)
	secret = fmt.Sprintf("%s%s","\nexport AWS_SECRET_ACCESS_KEY=",*response.AccessKey.SecretAccessKey)
	f, err := os.Create("credentials.txt")
	if err != nil {
		cfdl.Logger.Error("File creation error: ",err)
	}
	err = os.Chmod(sshKeyFileName(), 0600)
	if err != nil {
		cfdl.Logger.Error("File change permission error: ",err)
	}

	_, err = f.WriteString(key)
	_, err = f.WriteString(secret)
	if err != nil {
		cfdl.Logger.Error("Credentials file write error: ",err)
		f.Close()
	}
	err = f.Close()
	if err != nil {
		cfdl.Logger.Error("Credentials file close error: ",err)			
	}	

	

	keypair = AccessKey{
		AccessKeyID: *response.AccessKey.AccessKeyId,
		SecretAccessKey: *response.AccessKey.SecretAccessKey,
	}
	cfdl.Logger.Info("Key ",keypair.AccessKeyID)
	
}

// DeleteAccessKey deletes
func DeleteAccessKey(client IAMInterface) {
	log.Println("DeleteAccessKey Access Key")

	listKeyResponse, err :=client.ListAccessKeys(context.TODO(),&iam.ListAccessKeysInput{
        UserName: aws.String("awsstudent"),
	})
	if err != nil {
		cfdl.Logger.Error("Error list access key: ",err)
	}else {
		keys := listKeyResponse.AccessKeyMetadata
		
		for _, key := range keys {
			keyID := key.AccessKeyId
			log.Println("Deleting key:", *keyID)
			cfdl.Logger.Info("Deleting key:", *keyID)
			
			params := &iam.DeleteAccessKeyInput{
				UserName: aws.String(labUser),
				AccessKeyId: key.AccessKeyId,
			}
			_, err := client.DeleteAccessKey(context.TODO(), params)
			if err != nil {
				cfdl.Logger.Error("Error deleting access key: ",err)
			}
		}
	}
	
}


// DeleteUserIfExist cleanup
func DeleteUserIfExist(client IAMInterface) {
	log.Println("Delete user")
	if studentExists(client) {
		fmt.Println("Found ", labUser, "deleting...")
		_,err := client.DeleteLoginProfile(context.TODO(), &iam.DeleteLoginProfileInput{
			UserName: aws.String(labUser),
		});
		_, err = client.DeleteUser(context.TODO(), &iam.DeleteUserInput{
			UserName: aws.String(labUser),
		})
		if err != nil {
			panic(err)
		}
	}
}

func studentExists(client IAMInterface) bool {
	var studentExists bool
	studentExists = false
	listResponse, err := client.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		panic(err)
	}
	for _, user := range listResponse.Users {
		if *user.UserName == labUser {
			studentExists = true
		}
	}
	return studentExists
}



// ClientIAM IAM
func ClientIAM() *iam.Client {
	client := iam.NewFromConfig(*getCfg())
	return client
}

func generateMyPassword()(string){
	rand.Seed(time.Now().Unix())
    minSpecialChar := 1
    minNum := 1
    minUpperCase := 1
    passwordLength := 16
    password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
    return password

}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
    var password strings.Builder

    //Set special character
    for i := 0; i < minSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < minNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}