package labs_test
import (
	"io/ioutil"
	labdeploy "labdeploy/labs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountParamsNoParms(t *testing.T) {
	template, _ := ioutil.ReadFile("test/cfn_wo_parms.txt")
	count := labdeploy.CountParameter(template)
	assert.Equal(t,0, count)
}

func TestCountParams5Parms(t *testing.T) {
	template, _ := ioutil.ReadFile("test/cfn_w_parms.txt")
	// log.Println(string(template))
	count := labdeploy.CountParameter(template)
	assert.Equal(t,5, count)
}

func TestParms(t *testing.T) {
	template, _ := ioutil.ReadFile("test/cfn_w_parms.txt")
	// log.Println(string(template))
	parms := labdeploy.ReadParameter(template)
	_, ok := parms["AWSAccessKey"]
	assert.Equal(t, true , ok)
}