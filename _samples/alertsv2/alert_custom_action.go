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

	response, err := alertCli.ExecuteCustomAction(alertsv2.ExecuteCustomActionRequest{
		Identifier: &identifier,
		User:       "test",
		Source:     "Source",
		Note:       "Note",
		ActionName: "customActionNew",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("RequestID: " + response.RequestID)

	}
}