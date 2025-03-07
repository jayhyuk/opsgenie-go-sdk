package main

import (
	"fmt"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	sch "github.com/jayhyuk/opsgenie-go-sdk/schedule"
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	schCli, cliErr := cli.Schedule()

	if cliErr != nil {
		panic(cliErr)
	}

	req := sch.ListSchedulesRequest{}
	response, schErr := schCli.List(req)

	if schErr != nil {
		panic(schErr)
	}
	for _, sch := range response.Schedules {
		fmt.Printf("Id: %s\n", sch.Id)
		fmt.Printf("Name: %s\n", sch.Name)
		fmt.Printf("Team: %s\n", sch.Team)
		fmt.Printf("Rules:\n")
		for _, rule := range sch.Rules {
			fmt.Printf("Id: %s\n", rule.Id)
			fmt.Printf("Name: %s\n", rule.Name)
			fmt.Printf("StartDate: %s\n", rule.StartDate)
			fmt.Printf("EndDate: %s\n", rule.EndDate)
			fmt.Printf("Rotation Type: %s\n", rule.RotationType)
			fmt.Printf("Rotation Length: %d\n", rule.RotationLength)
			fmt.Printf("Participants: %s\n", rule.Participants)
			fmt.Printf("Restrictions:\n")
			for _, restriction := range rule.Restrictions {
				fmt.Printf("Start Day: %s\n", restriction.StartDay)
				fmt.Printf("Start Time: %s\n", restriction.StartTime)
				fmt.Printf("End Day: %s\n", restriction.EndDay)
				fmt.Printf("End Time: %s\n", restriction.EndTime)
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}