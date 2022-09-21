package labs
import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/megaproaktiv/cfdl"
)


// **** Stacks ***

// **** User ***

// CreateUserIfnotExist awsstudent go
func CreateUserIfnotExist(client IAMInterface) {
	log.Println("Create User")
	password := GenerateMyPassword()
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

