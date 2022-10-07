package newrelic

// AlertCondition describes what triggers an alert for a specific policy.
type AlertCondition struct {
	Enabled     bool                 `json:"enabled,omitempty"`
	Entities    []string             `json:"entities,omitempty"`
	ID          int                  `json:"id,omitempty"`
	Metric      string               `json:"metric,omitempty"`
	Name        string               `json:"name,omitempty"`
	RunbookURL  string               `json:"runbook_url,omitempty"`
	Terms       []AlertConditionTerm `json:"terms,omitempty"`
	Type        string               `json:"type,omitempty"`
	UserDefined AlertUserDefined     `json:"user_defined,omitempty"`
}

// AlertConditionTerm defines thresholds that trigger an AlertCondition.
type AlertConditionTerm struct {
	Duration     string `json:"duration,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Priority     string `json:"priority,omitempty"`
	Threshold    string `json:"threshold,omitempty"`
	TimeFunction string `json:"time_function,omitempty"`
}

// AlertUserDefined describes user-defined behavior for an AlertCondition.
type AlertUserDefined struct {
	Metric        string `json:"metric,omitempty"`
	ValueFunction string `json:"value_function,omitempty"`
}

// AlertConditionOptions define filters for GetAlertConditions.
type AlertConditionOptions struct {
	policyID int
	Page     int
}

func (o *AlertConditionOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]any{
		"policy_id": o.policyID,
		"page":      o.Page,
	})
}

// GetAlertConditions will return any AlertCondition defined for a given
// policy, optionally filtered by AlertConditionOptions.
func (c *Client) GetAlertConditions(policy int, options *AlertConditionOptions) ([]AlertCondition, error) {
	resp := &struct {
		Conditions []AlertCondition `json:"conditions,omitempty"`
	}{}
	options.policyID = policy
	err := c.doGet("alerts_conditions.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Conditions, nil
}
