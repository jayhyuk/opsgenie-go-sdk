package main

import (
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
	"github.com/jayhyuk/opsgenie-go-sdk/alertsv2"
	"fmt"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, _ := cli.AlertV2()

	request := alertsv2.RemoveTagsRequest{
		Identifier: &alertsv2.Identifier{
			TinyID: "2",
		},
		User:   "test",
		Source: "Source",
		Note:   "Note",
		Tags:   []string{"tag1"},
	}
	response, err := alertCli.RemoveTags(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("RequestID: " + response.RequestID)
	}
}
