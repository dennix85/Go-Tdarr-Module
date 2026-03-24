package tdarr

import (
	"encoding/json"
)

// AlterWorkerLimitRequest is the request body for this endpoint.
type AlterWorkerLimitRequest struct {
	NodeID string `json:"nodeID"`
	// Allowed: ['increase', 'decrease']
	Process string `json:"process"`
	// Allowed: ['healthcheckcpu', 'healthcheckgpu', 'transcodecpu', 'transcodegpu']
	WorkerType string `json:"workerType"`
}

// CancelWorkerItemRequest is the request body for this endpoint.
type CancelWorkerItemRequest struct {
	NodeID string `json:"nodeID"`
	WorkerID string `json:"workerID"`
	Cause string `json:"cause"`
}

// DisconnectNodeRequest is the request body for this endpoint.
type DisconnectNodeRequest struct {
	NodeID string `json:"nodeID"`
}

// FileDownloadRequest is the request body for this endpoint.
type FileDownloadRequest struct {
	FilePath string `json:"filePath,omitempty"`
}

// GetNewTaskRequest is the request body for this endpoint.
type GetNewTaskRequest struct {
	NodeID string `json:"nodeID"`
	WorkerID string `json:"workerID"`
	WorkerType string `json:"workerType"`
	Automation map[string]interface{} `json:"automation,omitempty"`
}

// GetNodeLogRequest is the request body for this endpoint.
type GetNodeLogRequest struct {
	NodeID string `json:"nodeID"`
}

// ItemProcEndRequest is the request body for this endpoint.
type ItemProcEndRequest struct {
	NodeID string `json:"nodeID"`
	Obj map[string]interface{} `json:"obj"`
}

// KillWorkerRequest is the request body for this endpoint.
type KillWorkerRequest struct {
	// The ID of the node where the worker is running.
	NodeID string `json:"nodeID"`
	// The ID of the worker to kill.
	WorkerID string `json:"workerID"`
}

// LogJobReportRequest is the request body for this endpoint.
type LogJobReportRequest struct {
	Date float64 `json:"date"`
	Job map[string]interface{} `json:"job"`
	Text string `json:"text"`
}

// NodesVersionCheckRequest is the request body for this endpoint.
type NodesVersionCheckRequest struct {
	NodeID string `json:"nodeID,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
	NodeVersion string `json:"nodeVersion"`
	NodeType string `json:"nodeType,omitempty"`
	InDocker bool `json:"inDocker,omitempty"`
	ProcessPid int `json:"processPid,omitempty"`
}

// PollWorkerLimitsRequest is the request body for this endpoint.
type PollWorkerLimitsRequest struct {
	NodeID string `json:"nodeID"`
}

// ReadPluginRequest is the request body for this endpoint.
type ReadPluginRequest struct {
	Plugin map[string]interface{} `json:"plugin"`
}

// RestartNodeRequest is the request body for this endpoint.
type RestartNodeRequest struct {
	// The ID of the node to restart.
	NodeID string `json:"nodeID"`
}

// UpdateNodeRequest is the request body for this endpoint.
type UpdateNodeRequest struct {
	// The ID of the node to update.
	NodeID string `json:"nodeID"`
	// The updates to apply to the node.
	NodeUpdates map[string]interface{} `json:"nodeUpdates"`
}

// UpdateNodeRelayRequest is the request body for this endpoint.
type UpdateNodeRelayRequest struct {
	NodeID string `json:"nodeID"`
	ResStats map[string]interface{} `json:"resStats"`
	Workers map[string]interface{} `json:"workers"`
}

// AlterWorkerLimit - For changing the number of running workers of a specific type on a specific node
func (c *Client) AlterWorkerLimit(req AlterWorkerLimitRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/alter-worker-limit", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// CancelWorkerItem - For cancelling a running worker item on a specific node
func (c *Client) CancelWorkerItem(req CancelWorkerItemRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/cancel-worker-item", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// DisconnectNode - For forcefully disconnecting a node
func (c *Client) DisconnectNode(req DisconnectNodeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/disconnect-node", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// DownloadPlugins - For nodes to download the latest plugins zip
func (c *Client) DownloadPlugins() (string, error) {
	resp, err := c.get(c.baseURL + "/api/v2/download-plugins")
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// FileDownload - For downloading a file
func (c *Client) FileDownload(req FileDownloadRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/file/download", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// FileUpload - For uploading a file
func (c *Client) FileUpload() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/file/upload", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// GetNewTask - For a node to request a new task
func (c *Client) GetNewTask(req GetNewTaskRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-new-task", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// GetNodeLog - For getting the log of a node
func (c *Client) GetNodeLog(req GetNodeLogRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-node-log", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// GetNodes - For getting connected nodes information
func (c *Client) GetNodes() (json.RawMessage, error) {
	resp, err := c.get(c.baseURL + "/api/v2/get-nodes")
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// ItemProcEnd - For when a node completes processing an item
func (c *Client) ItemProcEnd(req ItemProcEndRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/item-proc-end", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// KillWorker - For killing a worker on a node
func (c *Client) KillWorker(req KillWorkerRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/kill-worker", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// LogJobReport - For updating a job report
func (c *Client) LogJobReport(req LogJobReportRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/log-job-report", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// NodesVersionCheck - For nodes to verify the server version before establishing a full connection
func (c *Client) NodesVersionCheck(req NodesVersionCheckRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/nodes/version-check", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// PollWorkerLimits - For a node to get its worker limits and check if there's anything in the queue
func (c *Client) PollWorkerLimits(req PollWorkerLimitsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/poll-worker-limits", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// ReadPlugin - For a node to read a plugin
func (c *Client) ReadPlugin(req ReadPluginRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/read-plugin", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// RestartNode - For restarting a specific node
func (c *Client) RestartNode(req RestartNodeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/restart-node", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// SyncPlugins - For syncing plugins from server to all nodes
func (c *Client) SyncPlugins() (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/sync-plugins", nil)
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdateNode - For the UI to update a connected node
func (c *Client) UpdateNode(req UpdateNodeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-node", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdateNodeRelay - For nodes to update the server with their status
func (c *Client) UpdateNodeRelay(req UpdateNodeRelayRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-node-relay", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}
