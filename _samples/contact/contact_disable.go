package main

import(
	ogcli "github.com/jayhyuk/opsgenie-go-sdk/client"
	"github.com/jayhyuk/opsgenie-go-sdk/_samples/constants"
	contacts "github.com/jayhyuk/opsgenie-go-sdk/contact"
	"fmt"
)

func main(){
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	contactCli, cliErr := cli.Contact()
	if cliErr != nil {
		panic(cliErr)
	}

	contactReq := contacts.CreateContactRequest{ Method: "sms", To: "1-9999999999", Username: "fazilet@test.com"}

	contactResp, contactErr := contactCli.Create(contactReq)
	if contactErr != nil {
		panic(contactErr)
	}

	contactDisableReq := contacts.DisableContactRequest{ Id: contactResp.Id, Username: "fazilet@test.com"}
	contactDisableResp, contactDisableErr := contactCli.Disable(contactDisableReq)
	if contactDisableErr != nil {
		panic(contactDisableErr)
	}

	fmt.Printf("status: %s\n", contactDisableResp.Status)
}

