package tdarr

import (
	"encoding/json"
	"fmt"
)

// AuthStatusRequest is the request body for this endpoint.
type AuthStatusRequest struct {
	SaU bool `json:"saU"`
}

// BulkDeleteFilesRequest is the request body for this endpoint.
type BulkDeleteFilesRequest struct {
	FileIds []string `json:"fileIds"`
}

// BulkUpdateFilesRequest is the request body for this endpoint.
type BulkUpdateFilesRequest struct {
	FileIds []string `json:"fileIds"`
	UpdatedObj map[string]interface{} `json:"updatedObj"`
}

// ClientClienttypeRequest is the request body for this endpoint.
type ClientClienttypeRequest struct {
	Start float64 `json:"start"`
	PageSize float64 `json:"pageSize"`
	Filters []map[string]interface{} `json:"filters"`
	Sorts []map[string]interface{} `json:"sorts"`
	Opts map[string]interface{} `json:"opts"`
}

// CreateSampleRequest is the request body for this endpoint.
type CreateSampleRequest struct {
	FilePath string `json:"filePath"`
}

// CruddbRequest is the request body for this endpoint.
type CruddbRequest struct {
	// Allowed: ['SettingsGlobalJSONDB', 'LibrarySettingsJSONDB', 'NodeJSONDB', 'StatisticsJSONDB', 'VariablesJSONDB', 'UsersJSONDB', 'ApiKeysJSONDB', 'FlowsJSONDB', 'StagedJSONDB', 'FileJSONDB', 'F2FOutputJSONDB', 'WorkerVerdictHistoryJSONDB', 'JobsJSONDB', 'AutomationsJSONDB']
	Collection string `json:"collection"`
	// Allowed: ['getById', 'getByIndex', 'getAll', 'insert', 'update', 'incdec', 'removeOne', 'removeByDB', 'removeAll']
	Mode string `json:"mode"`
	DocID string `json:"docID,omitempty"`
	Obj map[string]interface{} `json:"obj,omitempty"`
	Filters []map[string]interface{} `json:"filters,omitempty"`
}

// DeleteCacheFileRequest is the request body for this endpoint.
type DeleteCacheFileRequest struct {
	File string `json:"file"`
}

// DeleteFileRequest is the request body for this endpoint.
type DeleteFileRequest struct {
	File map[string]interface{} `json:"file"`
}

// DeleteJobReportRequest is the request body for this endpoint.
type DeleteJobReportRequest struct {
	// The ID of the job report to delete
	JobId string `json:"jobId"`
}

// DeleteJobReportsRequest is the request body for this endpoint.
type DeleteJobReportsRequest struct {
	// Optional filters to apply when selecting job reports to delete
	Filters []map[string]interface{} `json:"filters,omitempty"`
}

// DeleteUnhealthyFilesRequest is the request body for this endpoint.
type DeleteUnhealthyFilesRequest struct {
	Table string `json:"table"`
}

// FindDuplicatesRequest is the request body for this endpoint.
type FindDuplicatesRequest struct {
	Threshold float64 `json:"threshold"`
	Count float64 `json:"count"`
}

// ListFootprintidReportsRequest is the request body for this endpoint.
type ListFootprintidReportsRequest struct {
	FootprintId string `json:"footprintId"`
}

// ReadJobFileRequest is the request body for this endpoint.
type ReadJobFileRequest struct {
	FootprintId string `json:"footprintId"`
	JobId string `json:"jobId"`
	JobFileId string `json:"jobFileId"`
}

// RescanFileRequest is the request body for this endpoint.
type RescanFileRequest struct {
	_id string `json:"_id"`
	DB string `json:"DB"`
}

// RunHelpCommandRequest is the request body for this endpoint.
type RunHelpCommandRequest struct {
	Mode string `json:"mode"`
	Text string `json:"text"`
}

// ScanIndividualFileRequest is the request body for this endpoint.
type ScanIndividualFileRequest struct {
	ScanTypes map[string]interface{} `json:"scanTypes"`
	File map[string]interface{} `json:"file"`
}

// SearchDbRequest is the request body for this endpoint.
type SearchDbRequest struct {
	String string `json:"string"`
	GreaterThanGB float64 `json:"greaterThanGB"`
	LessThanGB float64 `json:"lessThanGB"`
}

// SearchJobReportsRequest is the request body for this endpoint.
type SearchJobReportsRequest struct {
	SearchTerms string `json:"searchTerms"`
}

// SetAllStatusRequest is the request body for this endpoint.
type SetAllStatusRequest struct {
	DbID string `json:"dbID"`
	Mode string `json:"mode"`
	Table string `json:"table"`
	ProcessStatus string `json:"processStatus"`
}

// TranscodeUserVerdictRequest is the request body for this endpoint.
type TranscodeUserVerdictRequest struct {
	Obj map[string]interface{} `json:"obj"`
	Verdict string `json:"verdict"`
}

// UpdaterCheckRequest is the request body for this endpoint.
type UpdaterCheckRequest struct {
	ResetUpdater bool `json:"resetUpdater"`
	DownloadUpdate bool `json:"downloadUpdate"`
	ApplyUpdate bool `json:"applyUpdate"`
}

// UseTokenRequest is the request body for this endpoint.
type UseTokenRequest struct {
	Token string `json:"token"`
	Redirect string `json:"redirect"`
}

// AuthStatus - For checking Tdarr Pro status
func (c *Client) AuthStatus(req AuthStatusRequest) (bool, error) {
	resp, err := c.post(c.baseURL + "/api/v2/auth-status", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeBool(resp)
}

// BulkDeleteFiles - For bulk deleting selected files from disk and removing them from the database
func (c *Client) BulkDeleteFiles(req BulkDeleteFilesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/bulk-delete-files", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// BulkUpdateFiles - For bulk updating specific files by their IDs
func (c *Client) BulkUpdateFiles(req BulkUpdateFilesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/bulk-update-files", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// ClientClienttype - For loading and updating data in various tables found around the Tdarr UI
func (c *Client) ClientClienttype(clientType string, req ClientClienttypeRequest) (json.RawMessage, error) {
	resp, err := c.post(fmt.Sprintf(c.baseURL+"/api/v2/client/%s", clientType), map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// CreateSample - For creating a 30 second sample of a file
func (c *Client) CreateSample(req CreateSampleRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/create-sample", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// Cruddb - 
For interacting with the database

insert:    requires collection, docID, obj (with keys/values to insert)
getById:   requires collection, docID
getByIndex:requires collection, docID
getAll:    requires collection
update:    requires collection, docID, obj (with keys/values to update)
removeOne: requires collection, docID
removeAll: requires collection, optional filters (with filter conditions)
  
func (c *Client) Cruddb(req CruddbRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/cruddb", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// Debug - For getting a page with various debug info
func (c *Client) Debug() (string, error) {
	resp, err := c.get(c.baseURL + "/api/v2/debug")
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// DebugVarsType - For getting various debug info
func (c *Client) DebugVarsType(type string) (json.RawMessage, error) {
	resp, err := c.get(fmt.Sprintf(c.baseURL+"/api/v2/debug-vars/%s", type))
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// DeleteCacheFile - For deleting a cache file
func (c *Client) DeleteCacheFile(req DeleteCacheFileRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-cache-file", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// DeleteFile - For deleting a file on disk of a file in Tdarr DB
func (c *Client) DeleteFile(req DeleteFileRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-file", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// DeleteJobReport - Delete a single job report from both database and disk
func (c *Client) DeleteJobReport(req DeleteJobReportRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-job-report", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// DeleteJobReports - Delete multiple job reports from both database and disk (with optional filters)
func (c *Client) DeleteJobReports(req DeleteJobReportsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-job-reports", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// DeleteUnhealthyFiles - For deleting files which have failed to transcode (table3) or unhealthy files (table6)
func (c *Client) DeleteUnhealthyFiles(req DeleteUnhealthyFilesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/delete-unhealthy-files", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// FindDuplicates - For starting the find duplicates process
func (c *Client) FindDuplicates(req FindDuplicatesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/find-duplicates", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// GetDbStatuses - For getting the statuses of the Tdarr database
func (c *Client) GetDbStatuses() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-db-statuses", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// GetResStats - For getting server resource information.
func (c *Client) GetResStats() (*ResStats, error) {
	resp, err := c.post(c.baseURL+"/api/v2/get-res-stats", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out ResStats
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetServerLog - For getting the server log
func (c *Client) GetServerLog() (string, error) {
	resp, err := c.get(c.baseURL + "/api/v2/get-server-log")
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// GetTimeNow - For getting the current time on the server
func (c *Client) GetTimeNow() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-time-now", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// IsServerAlive - Old endpoint for checking if the server is alive (user 'status' instead)
func (c *Client) IsServerAlive() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/is-server-alive", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// ListFootprintidReports - For listing all job reports for a specific footprintId
func (c *Client) ListFootprintidReports(req ListFootprintidReportsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/list-footprintId-reports", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// PerformanceStats - For various performance stat info
func (c *Client) PerformanceStats() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/performance-stats", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// ReadJobFile - For reading a job report
func (c *Client) ReadJobFile(req ReadJobFileRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/read-job-file", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// RescanFile - For rescanning a file
func (c *Client) RescanFile(req RescanFileRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/rescan-file", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RestartServer - For restarting Tdarr Server
func (c *Client) RestartServer() (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/restart-server", nil)
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RestartUi - For restarting the UI
func (c *Client) RestartUi() (string, error) {
	resp, err := c.get(c.baseURL + "/api/v2/restart-ui")
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RunHelpCommand - For running an ffmpeg or handbrake help command on the Help tab
func (c *Client) RunHelpCommand(req RunHelpCommandRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/run-help-command", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// ScanIndividualFile - For scanning an individual file with various tools
func (c *Client) ScanIndividualFile(req ScanIndividualFileRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/scan-individual-file", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SearchDb - Old endpoint for searching the file database (use 'client' endpoint instead)
func (c *Client) SearchDb(req SearchDbRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/search-db", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SearchJobReports - For searching job reports
func (c *Client) SearchJobReports(req SearchJobReportsRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/search-job-reports", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// SetAllStatus - For requeueing files for transcode or health check for a specific library
func (c *Client) SetAllStatus(req SetAllStatusRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/set-all-status", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// Status - For checking server status.
func (c *Client) Status() (*ServerStatus, error) {
	resp, err := c.get(c.baseURL + "/api/v2/status")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out ServerStatus
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// StopDedupe - For stopping the dedupe process
func (c *Client) StopDedupe() (string, error) {
	resp, err := c.get(c.baseURL + "/api/v2/stop-dedupe")
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// TranscodeUserVerdict - For taking action on a staged item
func (c *Client) TranscodeUserVerdict(req TranscodeUserVerdictRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/transcode-user-verdict", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdaterCheck - For checking if an update is available.
func (c *Client) UpdaterCheck(req UpdaterCheckRequest) (*UpdaterCheckResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/updater/check", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out UpdaterCheckResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdaterPackageIndex - For getting the package index
func (c *Client) UpdaterPackageIndex() (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/updater/package-index", nil)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// UpdaterRelaunch - For relaunching Tdarr Server when an update is ready
func (c *Client) UpdaterRelaunch() (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/updater/relaunch", nil)
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UseToken - For using a token
func (c *Client) UseToken(req UseTokenRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/use-token", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}
