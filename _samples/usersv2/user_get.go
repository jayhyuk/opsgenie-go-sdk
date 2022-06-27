package main

import (
	"fmt"

	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	"github.com/jayhyuk/opsgenie-go-sdk/userv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	userCli, _ := cli.UserV2()

	request := userv2.GetUserRequest{
		Identifier: &userv2.Identifier{
			ID:     "example-user-id",
		},
	}

	response, err := userCli.Get(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.User)
	}
}
