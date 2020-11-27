package labdeploy

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// DeployInterface interface for deployment
type DeployInterface interface {
	CreateStack(ctx context.Context, params *cfn.CreateStackInput, optFns ...func(*cfn.Options)) (*cfn.CreateStackOutput, error)
	DescribeStackEvents(ctx context.Context, params *cfn.DescribeStackEventsInput, optFns ...func(*cfn.Options)) (*cfn.DescribeStackEventsOutput, error)
	DeleteStack(ctx context.Context, params *cfn.DeleteStackInput, optFns ...func(*cfn.Options)) (*cfn.DeleteStackOutput, error)
	UpdateStack(ctx context.Context, params *cfn.UpdateStackInput, optFns ...func(*cfn.Options)) (*cfn.UpdateStackOutput, error)
}



// **** Stacks ***

// CreateStack creates a lab stack
func CreateStack(client DeployInterface, name string, template []byte) {
	log.Println("Create Stack")

	templateBody := string(template)

	params := &cfn.CreateStackInput{
		StackName:    &name,
		TemplateBody: &templateBody,
		Parameters: []*types.Parameter{
			{ParameterKey: aws.String("KeyName"), ParameterValue: aws.String(labKey)},
			{ParameterKey: aws.String("AWSSecretAccessKey"), ParameterValue: aws.String("nixda")},
			{ParameterKey: aws.String("AWSAccessKey"), ParameterValue: aws.String("nixda")},
		},
		Capabilities: []types.Capability{"CAPABILITY_NAMED_IAM"},
	}

	_, err := client.CreateStack(context.TODO(), params)
	if err != nil {
		log.Println("Error with stack creation: ",err);
	}
}





// ClientEC2 an ec2client
func ClientEC2() *ec2.Client {
	client := ec2.NewFromConfig(*getCfg())
	return client
}



func getCfg() *aws.Config {
	cfg, err := config.LoadDefaultConfig(config.WithRegion("eu-central-1"))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	return &cfg
}
