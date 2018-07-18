package teamv2

// DeleteTeamMemberRequest is a request for deleting a team member.
type DeleteTeamMemberRequest struct {
	*Identifier
	ApiKey string
}

// DeleteTeamMemberResponse is a response of deleting a team member.
type DeleteTeamMemberResponse struct {
	ResponseMeta
	ActionResult
}

// GetApiKey returns api key.
func (r *DeleteTeamMemberRequest) GetApiKey() string {
	return r.ApiKey
}
