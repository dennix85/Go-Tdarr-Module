package tdarr

import (
	"encoding/json"
)

// StatsGetPiesRequest is the request body for this endpoint.
type StatsGetPiesRequest struct {
	LibraryId string `json:"libraryId,omitempty"`
}

// StatsGetResHistRequest is the request body for this endpoint.
type StatsGetResHistRequest struct {
	Timeframe string `json:"timeframe,omitempty"`
}

// StatsGetRunningWorkerHistRequest is the request body for this endpoint.
type StatsGetRunningWorkerHistRequest struct {
	Timeframe string `json:"timeframe,omitempty"`
}

// StatsGetSpaceSavedRequest is the request body for this endpoint.
type StatsGetSpaceSavedRequest struct {
	LibraryId string `json:"libraryId,omitempty"`
	Timeframe string `json:"timeframe,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	WorkerType string `json:"workerType,omitempty"`
	PluginId string `json:"pluginId,omitempty"`
}

// StatsGetStreamsRequest is the request body for this endpoint.
type StatsGetStreamsRequest struct {
	LibraryId string `json:"libraryId,omitempty"`
}

// StatsGetWorkerVerdictHistRequest is the request body for this endpoint.
type StatsGetWorkerVerdictHistRequest struct {
	Timeframe string `json:"timeframe,omitempty"`
	LibraryId string `json:"libraryId,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	WorkerType string `json:"workerType,omitempty"`
}

// StatsSpaceSavedAddRequest is the request body for this endpoint.
type StatsSpaceSavedAddRequest struct {
	LibraryId string `json:"libraryId,omitempty"`
	SpaceSaved float64 `json:"spaceSaved,omitempty"`
}

// StatsGetPies - For getting all or library pie stats
func (c *Client) StatsGetPies(req StatsGetPiesRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-pies", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsGetResHist - For getting server resource history
func (c *Client) StatsGetResHist(req StatsGetResHistRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-res-hist", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsGetRunningWorkerHist - For getting running worker history
func (c *Client) StatsGetRunningWorkerHist(req StatsGetRunningWorkerHistRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-running-worker-hist", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsGetSpaceSaved - For getting space saved history
func (c *Client) StatsGetSpaceSaved(req StatsGetSpaceSavedRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-space-saved", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsGetStreams - For getting stream stats info
func (c *Client) StatsGetStreams(req StatsGetStreamsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-streams", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsGetWorkerVerdictHist - For getting worker verdict history
func (c *Client) StatsGetWorkerVerdictHist(req StatsGetWorkerVerdictHistRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/get-worker-verdict-hist", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// StatsSpaceSavedAdd - For adding space saved record
func (c *Client) StatsSpaceSavedAdd(req StatsSpaceSavedAddRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/stats/space-saved-add", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}
