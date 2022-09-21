package labs
import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/megaproaktiv/cfdl"

)


// **** Stacks ***

// **** User ***

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
		if err != nil {
			log.Println(err)
		}
		_, err = client.DeleteUser(context.TODO(), &iam.DeleteUserInput{
			UserName: aws.String(labUser),
		})
		if err != nil {
			panic(err)
		}
	}
}


