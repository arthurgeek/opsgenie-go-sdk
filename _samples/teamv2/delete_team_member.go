package main

import (
	"fmt"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/teamv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	teamCli, _ := cli.TeamV2()

	request := teamv2.DeleteTeamMemberRequest{
		Identifier: &teamv2.Identifier{
			UserName: "user@company.com",
			Name:     "Acme",
		},
	}

	response, err := teamCli.DeleteTeamMember(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
