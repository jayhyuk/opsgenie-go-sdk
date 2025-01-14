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

	identifier := alertsv2.Identifier{
		TinyID: "2",
	};
	ackRequest := alertsv2.AcknowledgeRequest{
		Identifier: &identifier,
		User:       "test",
		Source:     "Source",
		Note:       "Note",
	}

	response, err := alertCli.Acknowledge(ackRequest)

	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("RequestId: " + response.RequestID)
	}
}
