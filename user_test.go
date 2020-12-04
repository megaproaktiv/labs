package labdeploy_test

import (
	"context"
	"labdeploy"
	"testing"
    "github.com/megaproaktiv/cfdl"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
)

func init() {
	cfdl.InitLogger()		
	defer cfdl.Logger.Sync()
}

func TestSomethingThatUsesIAMInterface(t *testing.T) {

        // make and configure a mocked IAMInterface
        mockedIAMInterface := &labdeploy.IAMInterfaceMock{
            AttachUserPolicyFunc: func(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error) {
	               panic("mock out the AttachUserPolicy method")
            },
            CreateAccessKeyFunc: func(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error) {
	               panic("mock out the CreateAccessKey method")
            },
            CreateLoginProfileFunc: func(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error) {
	               return &iam.CreateLoginProfileOutput{},nil
            },
            CreatePolicyFunc: func(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error) {
	               panic("mock out the CreatePolicy method")
            },
            CreateUserFunc: func(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error) {
	               return &iam.CreateUserOutput{},nil
            },
            DeleteAccessKeyFunc: func(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error) {
	               panic("mock out the DeleteAccessKey method")
            },
            DeleteLoginProfileFunc: func(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error) {
	               panic("mock out the DeleteLoginProfile method")
            },
            DeletePolicyFunc: func(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error) {
	               panic("mock out the DeletePolicy method")
            },
            DeleteUserFunc: func(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error) {
	               panic("mock out the DeleteUser method")
            },
            DetachUserPolicyFunc: func(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error) {
	               panic("mock out the DetachUserPolicy method")
            },
            ListAccessKeysFunc: func(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
	               panic("mock out the ListAccessKeys method")
            },
            ListPoliciesFunc: func(ctx context.Context, params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) {
	               panic("mock out the ListPolicies method")
            },
            ListUsersFunc: func(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	               return &iam.ListUsersOutput{},nil
            },
        }

        labdeploy.CreateUserIfnotExist(mockedIAMInterface)
        callsToCreate := len(mockedIAMInterface.CreateUserCalls())
        assert.Equal(t,1,callsToCreate)
    }
