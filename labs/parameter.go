package labs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

// Template Parameters Cloudformation parameters
type Template struct {
	AWSTemplateFormatVersion string                 `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string                 `yaml:"Description,omitempty"`
	Metadata                 map[string]interface{} `yaml:"Metadata,omitempty"`
	Parameters               Parameters             `yaml:"Parameters,omitempty"`
	Mappings                 map[string]interface{} `yaml:"Mappings,omitempty"`
	Conditions               map[string]interface{} `yaml:"Conditions,omitempty"`

}

// Parameters map parms
type Parameters map[string]Parameter

// Parameter map parms
type Parameter struct {
	Type                  string      `yaml:"Type"`
	Description           string      `yaml:"Description,omitempty"`
	Default               interface{} `yaml:"Default,omitempty"`
	AllowedPattern        string      `yaml:"AllowedPattern,omitempty"`
	AllowedValues         []string    `yaml:"AllowedValues,omitempty"`
	ConstraintDescription string      `yaml:"ConstraintDescription,omitempty"`
	MaxLength             int         `yaml:"MaxLength,omitempty"`
	MinLength             int         `yaml:"MinLength,omitempty"`
	MaxValue              float64     `yaml:"MaxValue,omitempty"`
	MinValue              float64     `yaml:"MinValue,omitempty"`
	NoEcho                bool        `yaml:"NoEcho,omitempty"`
}
// ReadParameter read cloudformation parameter - NO USED YET
func (p *Parameters) ReadParameter() *Parameters {

	yamlFile, err := ioutil.ReadFile("parameters.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return p
}

// CountParameter - how mny params in the cfn template
func CountParameter(template []byte )(int){
	var stack Template

	err := yaml.Unmarshal([]byte(template), &stack)
	if err != nil {
		log.Fatalln("Unmarshal error: ",err)
	}
	// lelog.Println("Count: ",stack);
	// for k, v := range stack.Parameters {
	// 	fmt.Println("k:", k, "v:", v)
	// }
	
	count := len(stack.Parameters)

	return count
}


// ReadParameter - return parameters in template
func ReadParameter(template []byte )(Parameters){
	var stack Template

	params := make( Parameters);

	err := yaml.Unmarshal([]byte(template), &stack)
	if err != nil {
		log.Fatalln("Unmarshal error: ",err)
	}
	// lelog.Println("Count: ",stack);
	for k, v := range stack.Parameters {
		//fmt.Println("k:", k, "v:", v)
		params[k] = v

		
	}
	
	return params
}

