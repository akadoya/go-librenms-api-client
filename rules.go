package librenms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAlertRules - Returns list of Alert Rules
func (c *Client) GetAlertRules(authToken *string) ([]AlertRule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rules", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	rules := []AlertRule{}

	err = json.Unmarshal(body, &rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

// GetAlertRule - Returns specific Alert Rule
func (c *Client) GetAlertRule(ruleID string, authToken *string) ([]AlertRule, error) {
	var intermediate struct {
		Status string        `json:"status"`
		Rules  []AlertRule   `json:"rules"`
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rules/%s", c.BaseURL, ruleID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	// read rules list from response body
	err = json.Unmarshal(body, &intermediate)
	if err != nil {
		return nil, err
	}

	return intermediate.Rules, nil
}

// CreateAlertRule - Create new AlertRule
func (c *Client) CreateAlertRule(rule AlertRule, authToken *string) (*AlertRule, error) {
	rb, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rules", c.BaseURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newAlertRule := AlertRule{}
	err = json.Unmarshal(body, &newAlertRule)
	if err != nil {
		return nil, err
	}

	return &newAlertRule, nil
}
