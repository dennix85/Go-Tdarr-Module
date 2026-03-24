package tdarr

import (
	"encoding/json"
)

// CopyCommunityToLocalRequest is the request body for this endpoint.
type CopyCommunityToLocalRequest struct {
	PluginID string `json:"pluginID"`
	ForceOverwrite bool `json:"forceOverwrite"`
}

// CreatePluginRequest is the request body for this endpoint.
type CreatePluginRequest struct {
	Details map[string]interface{} `json:"details"`
	ConditionalsString string `json:"conditionalsString"`
	ConditionalNotes string `json:"conditionalNotes"`
	Action map[string]interface{} `json:"action"`
}

// DeletePluginRequest is the request body for this endpoint.
type DeletePluginRequest struct {
	PluginSource string `json:"pluginSource"`
	PluginID string `json:"pluginID"`
}

// ReadPluginTextRequest is the request body for this endpoint.
type ReadPluginTextRequest struct {
	PluginSource string `json:"pluginSource"`
	PluginID string `json:"pluginID"`
}

// SavePluginTextRequest is the request body for this endpoint.
type SavePluginTextRequest struct {
	PluginSource string `json:"pluginSource"`
	PluginID string `json:"pluginID"`
	Text string `json:"text"`
}

// SearchFlowPluginsRequest is the request body for this endpoint.
type SearchFlowPluginsRequest struct {
	String string `json:"string"`
	PluginType string `json:"pluginType"`
}

// SearchFlowTemplatesRequest is the request body for this endpoint.
type SearchFlowTemplatesRequest struct {
	String string `json:"string"`
	PluginType string `json:"pluginType"`
}

// SearchPluginsRequest is the request body for this endpoint.
type SearchPluginsRequest struct {
	String string `json:"string"`
	PluginType string `json:"pluginType"`
}

// UpdatePluginsRequest is the request body for this endpoint.
type UpdatePluginsRequest struct {
	Force bool `json:"force"`
}

// CopyCommunityToLocal - For copying a community plugin to local plugins
func (c *Client) CopyCommunityToLocal(req CopyCommunityToLocalRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/copy-community-to-local", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// CreatePlugin - For creating a basic classic plugin using the classic plugin creator
func (c *Client) CreatePlugin(req CreatePluginRequest) (bool, error) {
	resp, err := c.post(c.baseURL + "/api/v2/create-plugin", map[string]interface{}{"data": req})
	if err != nil {
		return false, err
	}
	return c.decodeBool(resp)
}

// DeletePlugin - For deleting a plugin
func (c *Client) DeletePlugin(req DeletePluginRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-plugin", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// ReadPluginText - For the classic plugin editor to read a plugin file
func (c *Client) ReadPluginText(req ReadPluginTextRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/read-plugin-text", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SavePluginText - For the classic plugin editor to save plugin text
func (c *Client) SavePluginText(req SavePluginTextRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/save-plugin-text", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SearchFlowPlugins - For searching flow plugins
func (c *Client) SearchFlowPlugins(req SearchFlowPluginsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/search-flow-plugins", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SearchFlowTemplates - For searching flow templates
func (c *Client) SearchFlowTemplates(req SearchFlowTemplatesRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/search-flow-templates", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SearchPlugins - For searching classic plugins
func (c *Client) SearchPlugins(req SearchPluginsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/search-plugins", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// UpdatePlugins - For requesting the server to update community plugins
func (c *Client) UpdatePlugins(req UpdatePluginsRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-plugins", map[string]interface{}{"data": req})
	if err != nil {
		return "", err
	}
	return c.decodeString(resp)
}
