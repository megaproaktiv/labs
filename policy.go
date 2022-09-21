// go get -u github.com/aws/aws-sdk-go/...

package labs

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const policyName = "iam_policy.json"

// CreateLabPolicy policy for awsstudent
func CreateLabPolicy(client IAMInterface, lab int) {
	policyName := PolicyIAMName(lab)
	policyFileName := policyFileName(lab)

	log.Println("Create Lab Policy ", policyName, " from: ", policyFileName)

	policyFile, err := os.ReadFile(policyFileName)
	// fmt.Printf("File contents: %s", policyFile)

	// Lab Policy
	policyResponse, err := client.CreatePolicy(context.TODO(), &iam.CreatePolicyInput{
		PolicyName:     aws.String(policyName),
		Description:    aws.String(policyName),
		PolicyDocument: aws.String(string(policyFile)),
	})
	if err != nil {
		log.Println("CreateLabPolicy error: ", err)
	} else {
		log.Printf("Created policy: %v\n", policyName)
	}
	labPolicyArn := policyResponse.Policy.Arn

	// User policy
	userPolicyName := PolicyIAMAdditionalName(lab)
	policyResponse, err = client.CreatePolicy(context.TODO(), &iam.CreatePolicyInput{
		PolicyName:     aws.String(userPolicyName),
		Description:    aws.String(userPolicyName),
		PolicyDocument: aws.String(getUserPolicy()),
	})
	if err != nil {
		log.Println("CreateLabPolicy error: ", err)
	} else {
		log.Printf("Created policy: %v\n", userPolicyName)
	}
	userPolicyArn := policyResponse.Policy.Arn

	_, err = client.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
		PolicyArn: labPolicyArn,
		UserName:  aws.String("awsstudent"),
	})
	_, err = client.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
		PolicyArn: userPolicyArn,
		UserName:  aws.String("awsstudent"),
	})

}

// DeleteLabPolicy policy for awsstudent
func DeleteLabPolicy(client IAMInterface, lab int) {

	log.Println("Search Lab Policy to delete")

	params := &iam.ListPoliciesInput{}

	
	paginator := iam.NewListPoliciesPaginator(client, params)
	for paginator.HasMorePages() {
		listResponse, err := paginator.NextPage(context.TODO())
		if err != nil {
			// handle error
			log.Println("List Policy error: ", err)
		}
		policies := listResponse.Policies

		for _, policy := range policies {
			policyName := policy.PolicyName
			if (*policyName == PolicyIAMName(lab)) || (*policyName == PolicyIAMAdditionalName(lab)) {
				fmt.Printf("Policy to delete is %v\n", *policyName)
				_, err := client.DetachUserPolicy(context.TODO(), &iam.DetachUserPolicyInput{
					UserName:  aws.String(labUser),
					PolicyArn: policy.Arn,
				})
				if err != nil {
					log.Println("Detach Lab Policy error: ", err)
				} else {
					log.Printf("Detached policy:%v\n", policy.Arn)
				}
				// ************************
				_, err = client.DeletePolicy(context.TODO(), &iam.DeletePolicyInput{
					PolicyArn: policy.Arn,
				})
				if err != nil {
					log.Println("Delete Lab Policy error: ", err)
				} else {
					log.Printf("Deleted policy: %v\n", policy.Arn)
				}

			}
		}

	}

}

func PolicyIAMName(lab int) string {
	return "lab" + strconv.Itoa(lab) + "-policy"
}
func PolicyIAMAdditionalName(lab int) string {
	return "lab-" + strconv.Itoa(lab) + "-user-policy"
}

func policyFileName(lab int) string {
	return "./lab" + strconv.Itoa(lab) + "/" + policyName
}

// Additional policy attached to the user
func getUserPolicy() string {
	part1 := `{
"Version": "2012-10-17",
"Statement": [
	{
	"Effect": "Allow",
	"Action": "iam:GetAccountPasswordPolicy",
	"Resource": "*"
	},
	{
	"Effect": "Allow",
	"Action": "iam:ChangePassword",
	"Resource": "arn:aws:iam::`

	part2 := `"}]}`

	up := part1 + GetAccount() + ":user/" + labUser + part2
	return up

}
