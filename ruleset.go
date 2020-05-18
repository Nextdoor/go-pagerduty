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

type EventRule struct {
	APIObject
	Position int `json:"position,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	CatchAll bool `json:"catch_all,omitempty"`
	Conditions `json:"conditions,omitempty"`
	TimeFrame `json:"time_frame,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Conditions struct {
	Operator string `json:"operator,omitempty"`
	Subconditions []Subcondition `json:"subconditions,omitempty"`
}

type Subcondition struct {
	Operator string `json:"operator,omitempty"`
	Parameters []Parameter `json:"parameter,omitempty"`
}

type Parameter struct {
	Path string `json:"path,omitempty"`
	Value `json:"value,omitempty"`
	Options interface{} `json:"options,omitempty"`
}

type TimeFrame struct {
	ActiveBetween `json:"active_between,omitempty"`
	ScheduledWeekly `json:"scheduled_weekly,omitempty"`
}

type ActiveBetween struct {
	StartTime int `json:"start_time,omitempty"`
	EndTime int `json:"end_time,omitempty"`
}

type ScheduledWeekly struct {
	StartTime int `json:"start_time,omitempty"`
	Duration int `json:"duration,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Weekdays []int `json:"weekdays,omitempty"`
}

type Action struct {
	Annotate Value `json:"annotate,omitempty"`
	EventAction Value `json:"event_action,omitempty"`
	Extractions []Extraction `json:"extractions,omitempty"`
	Priority Value `json:"priority,omitempty"`
	Route Value `json:"route,omitempty"`
	Severity Value `json:"severity,omitempty"`
	Suppress `json:"suppress,omitempty"`
}

type Extraction struct {
	Target string `json:"target,omitempty"`
	Source string `json:"source,omitempty"`
	Regex string `json:"regex,omitempty"`
}

type Value struct {
	Value string `json:"value,omitempty"`
}

type Suppress struct {
	Value bool `json:"value,omitempty"`
	ThresholdValue int `json:"threshold_value,omitempty"`
	ThresholdTimeUnit string `json:"threshold_time_unit,omitempty"`
	ThresholdTimeAmount int `json:"threshold_time_amount,omitempty"`
}

// CreateEventRule creates a new event rule.
func (c *Client) CreateEventRule(rulesetId string, r EventRule) (*EventRule, error) {
	data := make(map[string]EventRule)
	data["rule"] = r
	resp, err := c.post("/rulesets/" + rulesetId + "/rules", data, nil)
	if err != nil {
		return nil, err
	}
	return getEventRuleFromResponse(c, resp)
}

// GetEventRule shows information about the given event rule.
func (c *Client) GetEventRule(rulesetId, id string) (*EventRule, error) {
	resp, err := c.get("/rulesets/" + rulesetId + "rules" + id)
	if err != nil {
		return nil, err
	}
	return getEventRuleFromResponse(c, resp)
}

// UpdateEventRule updates an existing event rule.
func (c *Client) UpdateEventRule(rulesetId string, r EventRule) (*EventRule, error) {
	v := make(map[string]EventRule)
	v["rule"] = r
	resp, err := c.put("/rulesets/" + rulesetId + "rules" + r.ID, v, nil)
	if err != nil {
		return nil, err
	}
	return getEventRuleFromResponse(c, resp)
}

// DeleteEventRule removes the event rule.
func (c *Client) DeleteEventRule(rulesetId, id string) error {
	_, err := c.delete("/rulesets/" + rulesetId + "rules" + id)
	return err
}

func getEventRuleFromResponse(c *Client, resp *http.Response) (*EventRule, error) {
	var target map[string]EventRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	rootNode := "rule"
	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}
	return &t, nil
}
