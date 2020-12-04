package labdeploy_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labdeploy"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"

	"testing"
	//"github.com/stretchr/testify/assert"
)


func TestEmptyKeypair(t *testing.T) {

    // make and configure a mocked Ec2Interface
    mockedEc2Interface := &labdeploy.Ec2InterfaceMock{
        CreateKeyPairFunc: func(ctx context.Context, params *ec2.CreateKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error) {
                return &ec2.CreateKeyPairOutput{
                        KeyFingerprint: aws.String(".kjn"),
                        KeyMaterial: aws.String("KeyLockOpen"),
                }, nil

        },
        DeleteKeyPairFunc: func(ctx context.Context, params *ec2.DeleteKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error) {
                panic("mock out the DeleteKeyPair method")
        },
        DescribeImagesFunc: func(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
                panic("mock out the DescribeImages method")
        },
        DescribeKeyPairsFunc: func(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
            response := ec2.DescribeKeyPairsOutput{
                    KeyPairs:  []types.KeyPairInfo{},

            };

            return &response, nil
        },

    }

    // No Keys there => call create
    labdeploy.CreateKeyIfNotExist(mockedEc2Interface)
    callsToCreate := len(mockedEc2Interface.CreateKeyPairCalls())
    assert.Equal(t,1,callsToCreate)

}
func TestOtherKeyPair(t *testing.T) {

        // make and configure a mocked Ec2Interface
        mockedEc2Interface := &labdeploy.Ec2InterfaceMock{
            CreateKeyPairFunc: func(ctx context.Context, params *ec2.CreateKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error) {
                    return &ec2.CreateKeyPairOutput{
                            KeyFingerprint: aws.String(".kjn"),
                            KeyMaterial: aws.String("KeyLockOpen"),
                    }, nil
    
            },
            DeleteKeyPairFunc: func(ctx context.Context, params *ec2.DeleteKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error) {
                    panic("mock out the DeleteKeyPair method")
            },
            DescribeImagesFunc: func(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
                    panic("mock out the DescribeImages method")
            },
            DescribeKeyPairsFunc: func(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
                
                var out ec2.DescribeKeyPairsOutput
                data, err := ioutil.ReadFile("testdata/ec2-keypairs.json")
                if err != nil {
                        fmt.Println("File reading error", err)
                }
                json.Unmarshal(data, &out);		
                
    
                return &out, nil
            },
    
        }
    
        // No Keys there => call create
        labdeploy.CreateKeyIfNotExist(mockedEc2Interface)
        callsToCreate := len(mockedEc2Interface.CreateKeyPairCalls())
        assert.Equal(t,1,callsToCreate)
    
    }