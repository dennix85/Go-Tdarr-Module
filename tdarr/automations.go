package tdarr

import (
	"encoding/json"
)

// RunAutomationRequest is the request body for this endpoint.
type RunAutomationRequest struct {
	ConfigId string `json:"configId"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	LibraryIds []string `json:"libraryIds,omitempty"`
	TargetNodeNames []string `json:"targetNodeNames,omitempty"`
	TargetNodeIds []string `json:"targetNodeIds,omitempty"`
	ExecuteImmediately bool `json:"executeImmediately,omitempty"`
	BypassWorkerLimits bool `json:"bypassWorkerLimits,omitempty"`
	BypassStagedFileLimit bool `json:"bypassStagedFileLimit,omitempty"`
}

// RunAutomation - Manually trigger an automation
func (c *Client) RunAutomation(req RunAutomationRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/run-automation", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}
