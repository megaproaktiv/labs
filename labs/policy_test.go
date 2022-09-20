// go get -u github.com/aws/aws-sdk-go/...

package labs_test

import (
	"testing"
	"labdeploy/labs" 
	"github.com/stretchr/testify/assert"
)

func Test_policyIAMName(t *testing.T) {
	assert.Equal(t, labs.PolicyIAMName(1), "lab1-policy")
}

func Test_PolicyIAMAdditionalName(t *testing.T) {
	assert.Equal(t, labs.PolicyIAMAdditionalName(1), "lab-1-user-policy")
}

