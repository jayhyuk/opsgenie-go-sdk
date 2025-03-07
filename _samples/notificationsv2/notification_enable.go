package main

import (
	"fmt"

	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	"github.com/jayhyuk/opsgenie-go-sdk/notificationv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.NotificationV2()

	response, err := notificationCli.Enable(notificationv2.EnableNotificationRequest{
		Identifier: &notificationv2.Identifier{
			Username: "user@company.com",
			RuleID: "example-notification-id",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Result: %s, Request ID: %s\n", response.Result, response.RequestID)
	}
}
