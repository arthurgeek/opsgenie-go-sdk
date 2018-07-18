package teamv2

import (
	"errors"
	"net/url"
)

// Identifier defines the set of attributes for identification.
type Identifier struct {
	Name     string `json:"-"`
	UserName string `json:"-"`
}

// GenerateUrl generates API url using specified attributes of identifier.
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseURL := "/v2/teams/"

	params := url.Values{}

	if request.Name != "" {
		url := baseURL + request.Name

		params.Set("teamIdentifierType", "name")

		if request.UserName != "" {
			return url + "/members/" + request.UserName, params, nil
		}
	}

	return "", nil, errors.New("Name and UserName should be provided")
}
