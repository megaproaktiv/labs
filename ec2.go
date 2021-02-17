package labdeploy

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

//go:generate moq -out ec2_moq_test.go . Ec2Interface


// Ec2Interface all actions which are used from ec2
type Ec2Interface interface {
	DescribeKeyPairs(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	CreateKeyPair(ctx context.Context, params *ec2.CreateKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error)
	DeleteKeyPair(ctx context.Context, params *ec2.DeleteKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error)
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)

}

const labKey = "labkey"

// **** Keys ***

// CreateKeyIfNotExist ssh key handling
func CreateKeyIfNotExist(client Ec2Interface) {
	log.Println("Create EC2 Key: ",labKey)

	key := "labkey"
	// Check existance
	params := &ec2.DescribeKeyPairsInput{
		KeyNames: []string{key,},
	}
	var createKey bool
	createKey = false
	response, err := client.DescribeKeyPairs(context.TODO(), params)
	if err != nil {
		createKey = true
	}

	if createKey == false {
		if len(response.KeyPairs) == 0 {
			createKey = true
			log.Println("Key not found in account, thus creating a new one")
		}else{
			createKey = true
			for _, keypair := range response.KeyPairs {
				if *keypair.KeyName == labKey {
					createKey = false
				}
			}
		}
		
	}
		
	if !createKey {
		log.Println("Key <",labKey,"> already found in account, thus not creating a new one")
	}

	if createKey {
		log.Println("Writing ssh key local: ", sshKeyFileName())
		CreateKeyPairOutput, err := client.CreateKeyPair(context.TODO(), &ec2.CreateKeyPairInput{
			KeyName: aws.String(labKey),
		})
		if err != nil {
			panic(err)
		}else {
			key := fmt.Sprintf("%s",*CreateKeyPairOutput.KeyMaterial)
			
			f, err := os.Create(sshKeyFileName())
			if err != nil {
				log.Println("File creation error: ",err)
			}
			err = os.Chmod(sshKeyFileName(), 0600)
			if err != nil {
				log.Println("File change permission error: ",err)
			}
			
			_, err = f.WriteString(key)
			if err != nil {
				log.Println("Pem file write error: ",err)
				f.Close()
			}
			err = f.Close()
			if err != nil {
				log.Println("Pem file close error: ",err)			
			}		
		}
				
	}

	response, err = client.DescribeKeyPairs(context.TODO(), &ec2.DescribeKeyPairsInput{})
	if err != nil {
		panic(err)
	}
	if len(response.KeyPairs) == 0 {	
		log.Println("Key found in account.")
	}
		
}

func sshKeyFileName() string {
	return labKey+".pem"
}

// DeleteKey ec2 ssh key
func DeleteKey(client Ec2Interface) {
	log.Println("Delete EC2 Key: ",labKey)
	// Check existance
	params := &ec2.DeleteKeyPairInput{
		KeyName: aws.String(labKey),
	}

	_, err := client.DeleteKeyPair(context.TODO(), params)
	if err != nil {
		log.Println("Delete key error: ",err)			

	}
	
}