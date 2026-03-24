package tdarr

import (
	"encoding/json"
)

// GetProcessInfo - For getting process information from server and nodes with child process relationships
func (c *Client) GetProcessInfo() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-process-info", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}
