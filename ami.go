package labdeploy

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// FindAmi find ami id
func FindAmi(client Ec2Interface, amiParameter string) (string){
	architecture := extractArchitecture(amiParameter)
	name := extractName(amiParameter)
	params := &ec2.DescribeImagesInput{
		ExecutableUsers: []string{ "amazon",},
		Filters: []types.Filter{
			{ 
			Name: aws.String("architecture"), 	
			Values: []string{architecture,},},		
		 },
	};
	resp, err := client.DescribeImages(context.TODO(), params);
		 if err != nil {
			 panic("Could not find image "+amiParameter)
		 }

	bestImage := types.Image{
		CreationDate: aws.String("2018-11-11T07:00:45.000Z"),
		ImageId: aws.String("unknown"),
	};
	for _, key := range resp.Images{
		fullName := *key.Name
		leftName := strings.Split(fullName,".")[0]
		// todo creation date
		architectures := key.Architecture;
		if strings.Contains(string(architectures), architecture){
			
			if strings.Contains(string(leftName),name) {
				//  "2020-11-11T07:00:45.000Z",
				layout := "2006-01-02T15:04:05.000Z"
				str := key.CreationDate
				t1, _ := time.Parse(layout, *str)
				t2, _ := time.Parse(layout, *bestImage.CreationDate)
				if t1.After(t2) {
					bestImage = key
					
				}
			}
		}
	}
	return *bestImage.ImageId
}

func extractArchitecture(amiParameter string)(string){
	return strings.Split(amiParameter,",")[0]
}

func extractName(amiParameter string)(string){
	return strings.Split(amiParameter,",")[1]
}