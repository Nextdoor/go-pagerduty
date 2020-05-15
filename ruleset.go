package pagerduty

import (
	"fmt"
	"net/http"
)

type Ruleset struct {
	APIObject
	Name string `json:"name,omitempty"`
	RoutingKeys []string `json:"routing_keys,omitempty"`
	Team *APIObject `json:"team,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Creator *APIObject `json:"creator,omitempty"`
	Updater *APIObject `json:"updater,omitempty"`
}

// CreateRuleset creates a new ruleset.
func (c *Client) CreateRuleset(r Ruleset) (*Ruleset, error) {
	data := make(map[string]Ruleset)
	data["ruleset"] = r
	resp, err := c.post("/rulesets", data, nil)
	if err != nil {
		return nil, err
	}
	return getRulesetFromResponse(c, resp)
}

// GetRuleset shows information about the given ruleset.
func (c *Client) GetRuleset(id string) (*Ruleset, error) {
	resp, err := c.get("/rulesets/" + id)
	if err != nil {
		return nil, err
	}
	return getRulesetFromResponse(c, resp)
}

// UpdateRuleset updates an existing ruleset.
func (c *Client) UpdateRuleset(id string, r Ruleset) (*Ruleset, error) {
	v := make(map[string]Ruleset)
	v["ruleset"] = r
	resp, err := c.put("/rulesets/" + id, v, nil)
	if err != nil {
		return nil, err
	}
	return getRulesetFromResponse(c, resp)
}

// DeleteRuleset removes the ruleset.
func (c *Client) DeleteRuleset(id string) error {
	_, err := c.delete("/rulesets/" + id)
	return err
}

func getRulesetFromResponse(c *Client, resp *http.Response) (*Ruleset, error) {
	var target map[string]Ruleset
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	rootNode := "ruleset"
	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}
	return &t, nil
}
