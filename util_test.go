package labdeploy_test

import (
	"labdeploy"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRegion(t *testing.T) {

	expectedValues := "eu-central-1"
	os.Setenv("AWS_REGION",expectedValues)

	computedValue := labdeploy.GetRegion()
	
	assert.Equal(t,expectedValues, computedValue)

}