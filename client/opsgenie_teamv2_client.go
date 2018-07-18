package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/teamv2"
)

// OpsGenieTeamV2Client is the data type to make TeamAPI requests.
type OpsGenieTeamV2Client struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieTeamV2Client.
func (cli *OpsGenieTeamV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// DeleteTeamMember method deletes a team member.
func (cli *OpsGenieTeamV2Client) DeleteTeamMember(req teamv2.DeleteTeamMemberRequest) (*teamv2.DeleteTeamMemberResponse, error) {
	var response teamv2.DeleteTeamMemberResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
