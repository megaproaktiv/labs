// go get -u github.com/aws/aws-sdk-go/...

package labdeploy

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const policyName = "iam_policy.json"

// CreateLabPolicy policy for awsstudent
func CreateLabPolicy(client IAMInterface,lab int) {
	policyName := policyIAMName(lab)
	policyFileName := policyFileName(lab)
	
	log.Println("Create Lab Policy ", policyName," from: ",policyFileName)

	policyFile, err := ioutil.ReadFile(policyFileName)
	// fmt.Printf("File contents: %s", policyFile)

    // Lab Policy
    policyResponse, err := client.CreatePolicy(context.TODO(),&iam.CreatePolicyInput{
        PolicyName: aws.String(policyName),
        Description: aws.String(policyName),
        PolicyDocument: aws.String(string(policyFile)),
	})
	if err != nil {
		log.Println("CreateLabPolicy error: ",err);
	}
	labPolicyArn := policyResponse.Policy.Arn
	
	// User policy
	userPolicyName := policyIAMAdditionalName(lab);
    policyResponse, err = client.CreatePolicy(context.TODO(),&iam.CreatePolicyInput{
        PolicyName: aws.String(userPolicyName),
        Description: aws.String(userPolicyName),
        PolicyDocument: aws.String(getUserPolicy()),
	})
	if err != nil {
		log.Println("CreateLabPolicy error: ",err);
	}
	userPolicyArn := policyResponse.Policy.Arn

    _, err = client.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
        PolicyArn: labPolicyArn,
        UserName: aws.String("awsstudent"),
    })
    _, err = client.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
        PolicyArn: userPolicyArn,
        UserName: aws.String("awsstudent"),
    })

}

// DeleteLabPolicy policy for awsstudent
func DeleteLabPolicy(client IAMInterface,lab int) {
	
	log.Println("Delete Lab Policy")
	
	listResponse, err := client.ListPolicies(context.TODO(),&iam.ListPoliciesInput{})
	if err != nil {
		log.Println("Delete Lab Policy error: ",err);
	}
	policies := listResponse.Policies
	
	for _, policy := range policies {
		policyName := policy.PolicyName
		if *policyName == policyIAMName(lab) || *policyName == policyIAMAdditionalName(lab) {
			_, err :=client.DetachUserPolicy(context.TODO(), &iam.DetachUserPolicyInput{
				UserName: aws.String(labUser),
				PolicyArn: policy.Arn,
			})
			if err != nil {
				log.Println("Detach Lab Policy error: ",err);
			}
			// ************************
			_,err = client.DeletePolicy(context.TODO(), &iam.DeletePolicyInput{
				PolicyArn: policy.Arn,
				
			})
			if err != nil {
				log.Println("Delete Lab Policy error: ",err);
			}

		}
	}


}

func policyIAMName(lab int) string{
	return "lab"+strconv.Itoa(lab)+"-policy"
}
func policyIAMAdditionalName(lab int) string{
	return "lab-"+strconv.Itoa(lab)+"-user-policy"
}

func policyFileName(lab int) string{
	return "./lab"+strconv.Itoa(lab)+"/"+policyName
}


// Additional policy attached to the user
func getUserPolicy()(string){
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

	part2 :=`"}]}`

	up := part1+GetAccount()+":user/"+labUser+part2
	return up;

}