package main

import (
	// go
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	cfd "github.com/megaproaktiv/cfdl"

	// own
	labs "labdeploy"

	"github.com/thatisuday/clapper"
)

func init() {
	cfd.InitLogger()		
	defer cfd.Logger.Sync()
}

func main() {

	cfd.Logger.Info("Starting")
	
	const cmdDestroyString = "destroy"
	const cmdDeployString = "deploy"
	const cmdStatusString = "status"
	
	const cmdHelpString = "help"

	const flagLab = "lab"

	// Look for commands
	registry := clapper.NewRegistry()
	cmdDeploy, _ := registry.Register(cmdDeployString)
	cmdDestroy, _ := registry.Register(cmdDestroyString)
	// cmdStatus, _ := registry.Register(cmdStatusString)
	registry.Register(cmdDestroyString)
	registry.Register(cmdStatusString)
	registry.Register(cmdHelpString)
	
	
	cmdDeploy.AddFlag(flagLab,"l",false,"")
	cmdDestroy.AddFlag(flagLab,"l",false,"")
	// cmdStatus.AddFlag(flagLab,"l",false,"")

	// parse command-line arguments
	command, err := registry.Parse(os.Args[1:])

	// check for command line error
	if err != nil {
		fmt.Printf("error => %#v\n", err)
		help()
		return
	}

	// no command
	if( len(command.Name) == 0 ) {
		help()
		os.Exit(1)
	}
	cmd := command.Name;
	

	clientCfd := cfd.Client(labs.GetRegion())
	clientEc2 := labs.ClientEC2()
	clientIAM := labs.ClientIAM()
	
	
	const nameTemplate = "cloudformation_template.txt"
	
	var template []byte
	var stackname = "lab"
	var labNumber string
	for flagName, flagValue := range command.Flags {
		if strings.Compare(flagName, flagLab) == 0 {
			// Check exist
			labNumber = (*&flagValue.Value)
			
			path := "./lab"+labNumber+"/"+nameTemplate

			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("Template file ",path, "does not exist")
				panic(err)
			}
			
			template, err = ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	
	log.Println("Lab number: ",labNumber);

	var lab int
	lab, _ =strconv.Atoi(labNumber)
	if cmd == cmdDeployString {
		cfd.Logger.Info("Deploy ",lab)
		
		// Create a new CloudFormation template	
		// Check template switch
		labs.CreateKeyIfNotExist(clientEc2)
		labs.CreateUserIfnotExist(clientIAM)
		labs.CreateAccessKey(clientIAM)
		labs.CreateLabPolicy(clientIAM,lab);
		labs.CreateStack(clientCfd,stackname, template)
		labs.ShowStatus(clientCfd,stackname,template);
	}
	
	if cmd == cmdDestroyString {
		cfd.Logger.Info("Destroy ",lab)
		cfd.DeleteStack(clientCfd,stackname)
		labs.DeleteKey(clientEc2)
		labs.DeleteLabPolicy(clientIAM,lab)
		labs.DeleteAccessKey(clientIAM)
		labs.DeleteUserIfExist(clientIAM)
		labs.ShowStatus(clientCfd,stackname,template);
		
	}
	
	if cmd == cmdStatusString {
		labs.ShowStatus(clientCfd,stackname,template);
		//cfd.ShowStatus(client,stackname,template,cfd.StatusCreateComplete);
	}


	if cmd == cmdHelpString {
		help();
		os.Exit(0);
	}

}




func help(){
	fmt.Println("CloudFormation deploy app. ")
	fmt.Println("CloudFormation Template is generated automatically.")
	fmt.Println("Please call with  cfd [deploy|destroy] -l x - act on labx/cloudformation_template.txt")
	fmt.Println("Please call with  cfd status - show stack name lab  ")
	fmt.Println("Please call with  cfd help  ")
	fmt.Println("where x is the number of the lab, so labx subdirectory contains the  cloudformation_template.txt and the iam_policy.json is saved.")
	
}