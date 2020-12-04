package labdeploy

import (
	"context"
	"fmt"
	"os"
	"log"
	"strings"
	"time"
	"sort"
	"github.com/alexeyco/simpletable"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	tm "github.com/buger/goterm"
)

// StatusCreateComplete CloudFormation Status
const StatusCreateComplete="CREATE_COMPLETE"
// StatusCreateInProgress CloudFormation Status
const StatusCreateInProgress = "CREATE_IN_PROGRESS"
// StatusDeleteComplete CloudFormation Status
const StatusDeleteComplete = "DELETE_COMPLETE"

const (
	// ColorDefault default color
	ColorDefault = "\x1b[39m"
	// ColorRed red for screen
	ColorRed   = "\x1b[91m"
	// ColorGreen green for screen
	ColorGreen = "\x1b[32m"
	// ColorBlue blue for screen
	ColorBlue  = "\x1b[94m"
	// ColorGray for screen
	ColorGray  = "\x1b[90m"
)

// CloudFormationResource holder for status
type CloudFormationResource struct {
	LogicalResourceID string
	PhysicalResourceID string
	Status string
	Type string
	Timestamp time.Time
}


// ShowStatus status of stack
func ShowStatus(client DeployInterface, name string, template []byte){
	log.Println("Show Status")
    data := map[string]CloudFormationResource{}
	var lastTimeStamp time.Time

		// Draw
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "ID"},
			{Align: simpletable.AlignLeft, Text: "State"},
			{Align: simpletable.AlignLeft, Text: "Type"},
			{Align: simpletable.AlignLeft, Text: "PhysicalResourceID"},
		},
		
	}
	table.SetStyle(simpletable.StyleCompactLite)
	
	first := true
	for !IsStackCompleted(data) || !first {
		tm.Clear()
		tm.MoveCursor(1,1)
		data, lastTimeStamp = PopulateData(client, name, data);
		if IsStackCompleted(data){
			break
		}
		var statustext string
		// Sort
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for i, k := range keys {
			v := data[k]
			id := data[k].LogicalResourceID
			if( v.Status == StatusCreateComplete){
				statustext = green(StatusCreateComplete)
			}else if v.Status == StatusDeleteComplete {
				statustext = red(StatusDeleteComplete)
			} else{		
				statustext = gray(v.Status)
			}

			shortPhysical := fmt.Sprintf("%.80s", v.PhysicalResourceID)
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: id},
				{Align: simpletable.AlignLeft, Text: statustext},
				{Align: simpletable.AlignLeft, Text: v.Type},
				{Align: simpletable.AlignLeft, Text: shortPhysical},
			}
			if  !first {
				if i < len(table.Body.Cells) {

					table.Body.Cells[i]=r
				}else {
					table.Body.Cells = append(table.Body.Cells, r)
				}
			}else{
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
		first = false
		tm.Println(table.String())
		tm.Println(lastTimeStamp)
		tm.Flush()
		time.Sleep(1 * time.Second) 
	}
	
	
}


// PopulateData update status from describe call
func PopulateData(client DeployInterface, name string,data map[string]CloudFormationResource)( map[string]CloudFormationResource, time.Time){
	params := &cfn.DescribeStackEventsInput{
		StackName: &name,
	}
	
	output, error := client.DescribeStackEvents(context.TODO(), params)
	if( error != nil){
		msg  := error.Error()
		if strings.Contains(msg, "does not exist"){
			fmt.Println("Stack <",name,"> does not exist");
			os.Exit(0);
		}

		panic(error)
	}
	var lastTimeStamp time.Time
	// Update Status and Timestamp if newer
	for i := 0; i < len(output.StackEvents); i++ {
		
		
		event := output.StackEvents[i];		
		if i == 0 {
			lastTimeStamp = *event.Timestamp
		}
		item := data[*event.LogicalResourceId]

		if( event.Timestamp.After(item.Timestamp) ){
			item.LogicalResourceID = *event.LogicalResourceId;
			item.Status = string(event.ResourceStatus);
			item.Timestamp = *event.Timestamp;
			if event.Timestamp.After(lastTimeStamp){
				lastTimeStamp = *event.Timestamp
			}
			item.PhysicalResourceID = *event.PhysicalResourceId
			item.Type = *event.ResourceType
			data[*event.LogicalResourceId] = item;
			
		}
		
	}
	return data, lastTimeStamp;

}

// IsStackCompleted check for everything "completed"
func IsStackCompleted(data map[string]CloudFormationResource) bool {
	if len(data) == 0 {
		return false;
	}
	end := "_COMPLETE"
	for _, value := range data {
		if(!strings.HasSuffix(value.Status, end)){
			return false
		}
	}
	return true;
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
