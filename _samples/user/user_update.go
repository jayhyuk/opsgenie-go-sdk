package main

import (
	"fmt"

	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	"github.com/jayhyuk/opsgenie-go-sdk/user"
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	userCli, cliErr := cli.User()

	if cliErr != nil {
		panic(cliErr)
	}

	req := user.UpdateUserRequest{Id: "", Fullname: "", Role: "", Locale: "", Timezone: "", SkypeUsername: ""}
	response, userErr := userCli.Update(req)

	if userErr != nil {
		panic(userErr)
	}

	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)
}
