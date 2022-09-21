package labs

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
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

// IAMInterface used iam actions
type IAMInterface interface {
	AttachUserPolicy(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error)
	DetachUserPolicy(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error)

	CreateAccessKey(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error)
	DeleteAccessKey(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error)
	ListAccessKeys(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)

	CreateLoginProfile(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error)
	DeleteLoginProfile(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error)



	ListAttachedUserPolicies(ctx context.Context, params *iam.ListAttachedUserPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error)

	ListUsers(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
	CreateUser(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error)
	DeleteUser(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error)

	ListPolicies(ctx context.Context,params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) //
	CreatePolicy(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error)
	DeletePolicy(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error)
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

func GenerateMyPassword()(string){
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