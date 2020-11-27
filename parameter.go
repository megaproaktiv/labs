package labdeploy

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Parameters Cloudformation parameters
type Parameters struct {
	parameter []Parameter `yaml:"parameter"`
}

// Parameter Cloudformation
type Parameter struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// ReadParameter read cloudformation parameter
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
