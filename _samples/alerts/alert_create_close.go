package main

import (
	"fmt"

	alerts "github.com/jayhyuk/opsgenie-go-sdk/alerts"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	samples "github.com/jayhyuk/opsgenie-go-sdk/_samples"
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
)

func main() {

	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertID)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// close the alert
	cloreq := alerts.CloseAlertRequest{ID: response.AlertID, Notify: constants.NotifyArr}
	cloresponse, alertErr := alertCli.Close(cloreq)
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("status: %s\n", cloresponse.Status)
	fmt.Printf("code: %d\n", cloresponse.Code)
}
