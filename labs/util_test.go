package labs_test
import (
	labdeploy "labdeploy/labs"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

const integration = false
func TestGetRegion(t *testing.T) {
	if integration {

		expectedValues := "eu-central-1"
		os.Setenv("AWS_REGION",expectedValues)
		
		computedValue := labdeploy.GetRegion()
		
		assert.Equal(t,expectedValues, computedValue)
	}

}