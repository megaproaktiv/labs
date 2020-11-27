package labdeploy_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labdeploy"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"
)

func TestSomethingThatUsesEc2Interface(t *testing.T) {

	expectedValues := "ami-06032c95ea1ffa069"

	// make and configure a mocked Ec2Interface
	mockedEc2Interface := &labdeploy.Ec2InterfaceMock{
		
		DescribeImagesFunc: func(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
			var out ec2.DescribeImagesOutput
			data, err := ioutil.ReadFile("test/ami.json")
			if err != nil {
				fmt.Println("File reading error", err)
			}
			json.Unmarshal(data, &out);		
			return &out, nil
		},

	}

	computedValue := labdeploy.FindAmi(mockedEc2Interface, "x86_64,Windows_Server-2012-R2_RTM-English-64Bit-Base")
	
	assert.Equal(t,expectedValues, computedValue)

}