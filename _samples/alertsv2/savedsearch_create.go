package main

import (
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
	"github.com/jayhyuk/opsgenie-go-sdk/alertsv2"
	"fmt"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	"github.com/jayhyuk/opsgenie-go-sdk/alertsv2/savedsearches"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, _ := cli.AlertV2()

	request := savedsearches.CreateSavedSearchRequest{
		Name: "list-blue-team-alerts",
		Owner: alertsv2.User{
			Username: "user@opsgenie.com",
		},
		Teams: []alertsv2.Team{
			{Name: "blue_team"},
		},
		Description: "description",
		Query:       "status: Open",
	}

	response, err := alertCli.CreateSavedSearch(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response.SavedSearch

		fmt.Println("ID: " + savedSearch.ID)
		fmt.Println("Name: " + savedSearch.Name)
	}
}
