package tdarr

import (
	"encoding/json"
)

// AddAudioCodecExcludeRequest is the request body for this endpoint.
type AddAudioCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
}

// AddPluginIncludeRequest is the request body for this endpoint.
type AddPluginIncludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
	Source string `json:"source"`
	Index int `json:"index"`
}

// AddVideoCodecExcludeRequest is the request body for this endpoint.
type AddVideoCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
}

// GetFilescannerStatusRequest is the request body for this endpoint.
type GetFilescannerStatusRequest struct {
	DbID string `json:"dbID"`
}

// GetSubdirectoriesRequest is the request body for this endpoint.
type GetSubdirectoriesRequest struct {
	FolderPath string `json:"folderPath"`
}

// KillFileScannerRequest is the request body for this endpoint.
type KillFileScannerRequest struct {
	DbID string `json:"dbID"`
}

// RemoveAudioCodecExcludeRequest is the request body for this endpoint.
type RemoveAudioCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
}

// RemoveLibraryFilesRequest is the request body for this endpoint.
type RemoveLibraryFilesRequest struct {
	// The ID of the library to remove files from.
	DB string `json:"DB"`
}

// RemovePluginIncludeRequest is the request body for this endpoint.
type RemovePluginIncludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
}

// RemoveVideoCodecExcludeRequest is the request body for this endpoint.
type RemoveVideoCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
}

// ScanFilesRequest is the request body for this endpoint.
type ScanFilesRequest struct {
	ScanConfig map[string]interface{} `json:"scanConfig"`
}

// ToggleFolderWatchRequest is the request body for this endpoint.
type ToggleFolderWatchRequest struct {
	// Whether request has been triggered automatically (true) or by user (false).
	Auto bool `json:"auto"`
	// The path to the folder to watch.
	Folder string `json:"folder"`
	// The library ID.
	DbID string `json:"dbID"`
	// If the folder watcher should be enabled or disabled.
	Status bool `json:"status"`
}

// ToggleScheduleRequest is the request body for this endpoint.
type ToggleScheduleRequest struct {
	DbID string `json:"dbID"`
	Start int `json:"start"`
	End int `json:"end"`
	Type string `json:"type"`
}

// UpdateAudioCodecExcludeRequest is the request body for this endpoint.
type UpdateAudioCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
	Status bool `json:"status"`
}

// UpdatePluginIncludeRequest is the request body for this endpoint.
type UpdatePluginIncludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
	Status bool `json:"status"`
}

// UpdateScheduleBlockRequest is the request body for this endpoint.
type UpdateScheduleBlockRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
	Status bool `json:"status"`
}

// UpdateVideoCodecExcludeRequest is the request body for this endpoint.
type UpdateVideoCodecExcludeRequest struct {
	DbID string `json:"dbID"`
	Ele string `json:"ele"`
	Status bool `json:"status"`
}

// VerifyFolderExistsRequest is the request body for this endpoint.
type VerifyFolderExistsRequest struct {
	FolderPath string `json:"folderPath"`
}

// VerifyPluginRequest is the request body for this endpoint.
type VerifyPluginRequest struct {
	PluginID string `json:"pluginID"`
	Community bool `json:"community"`
}

// AddAudioCodecExclude - For adding an audio codec to be excluded/included in basic audio transcoding settings
func (c *Client) AddAudioCodecExclude(req AddAudioCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/add-audio-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// AddPluginInclude - For adding a plugin to a classic plugin stack
func (c *Client) AddPluginInclude(req AddPluginIncludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/add-plugin-include", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// AddVideoCodecExclude - For adding an video codec to be excluded/included in basic video transcoding settings
func (c *Client) AddVideoCodecExclude(req AddVideoCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/add-video-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// GetFilescannerStatus - For getting the status of a file scanner in progress
func (c *Client) GetFilescannerStatus(req GetFilescannerStatusRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-filescanner-status", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// GetSubdirectories - For getting subdirectories of a folder
func (c *Client) GetSubdirectories(req GetSubdirectoriesRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/get-subdirectories", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// KillFileScanner - For killing a file scanner in progress
func (c *Client) KillFileScanner(req KillFileScannerRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/kill-file-scanner", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RemoveAudioCodecExclude - For removing an audio codec to be excluded/included in basic audio transcoding settings
func (c *Client) RemoveAudioCodecExclude(req RemoveAudioCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/remove-audio-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RemoveLibraryFiles - For removing all files from a Tdarr library DB, files on disk aren't removed
func (c *Client) RemoveLibraryFiles(req RemoveLibraryFilesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/remove-library-files", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RemovePluginInclude - For removing a plugin from a classic plugin stack
func (c *Client) RemovePluginInclude(req RemovePluginIncludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/remove-plugin-include", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// RemoveVideoCodecExclude - For removing an video codec to be excluded/included in basic video transcoding settings
func (c *Client) RemoveVideoCodecExclude(req RemoveVideoCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/remove-video-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// ScanFiles - 
  For running a scanFresh, scanFindNew or scanFolderWatcher on a library
  scanFresh & scanFindNew require a single string directory path
  scanFolderWatcher requires an array of file paths to scanned
  
func (c *Client) ScanFiles(req ScanFilesRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/scan-files", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// ToggleFolderWatch - For enabling/disabling folder watching on a library
func (c *Client) ToggleFolderWatch(req ToggleFolderWatchRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/toggle-folder-watch", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// ToggleSchedule - For updating the schedule of a library
func (c *Client) ToggleSchedule(req ToggleScheduleRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/toggle-schedule", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdateAudioCodecExclude - For updating an audio codec to be excluded/included in basic audio transcoding settings
func (c *Client) UpdateAudioCodecExclude(req UpdateAudioCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-audio-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdatePluginInclude - For enabling/disabling a plugin in a classic plugin stack
func (c *Client) UpdatePluginInclude(req UpdatePluginIncludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-plugin-include", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdateScheduleBlock - For updating a block in a library schedule
func (c *Client) UpdateScheduleBlock(req UpdateScheduleBlockRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-schedule-block", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// UpdateVideoCodecExclude - For updating an video codec to be excluded/included in basic video transcoding settings
func (c *Client) UpdateVideoCodecExclude(req UpdateVideoCodecExcludeRequest) (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/update-video-codec-exclude", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// VerifyFolderExists - For verifying if a folder exists
func (c *Client) VerifyFolderExists(req VerifyFolderExistsRequest) (bool, error) {
	resp, err := c.post(c.baseURL + "/api/v2/verify-folder-exists", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeBool(resp)
}

// VerifyPlugin - For verifying if a classic plugin exists
func (c *Client) VerifyPlugin(req VerifyPluginRequest) (bool, error) {
	resp, err := c.post(c.baseURL + "/api/v2/verify-plugin", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decodeBool(resp)
}
