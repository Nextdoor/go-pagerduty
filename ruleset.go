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

// GetRuleset shows detailed information about a schedule, including entries for each layer and sub-schedule.
//func (c *Client) GetRuleset(id string, o GetRulesetOptions) (*Ruleset, error) {
//	v, err := query.Values(o)
//	if err != nil {
//		return nil, fmt.Errorf("Could not parse values for query: %v", err)
//	}
//	resp, err := c.get("/rulesets/" + id + "?" + v.Encode())
//	if err != nil {
//		return nil, err
//	}
//	return getRulesetFromResponse(c, resp)
//}

// UpdateRuleset updates an existing ruleset.
func (c *Client) UpdateRuleset(id string, r Ruleset) (*Ruleset, error) {
	v := make(map[string]Ruleset)
	v["ruleset"] = r
	resp, err := c.put("/rulesets/"+id, v, nil)
	if err != nil {
		return nil, err
	}
	return getRulesetFromResponse(c, resp)
}

// DeleteRuleset removes an override.
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
