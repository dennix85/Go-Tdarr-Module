package tdarr

import "encoding/json"

// DeleteBackupRequest is the request body for DeleteBackup.
type DeleteBackupRequest struct {
	Name string `json:"name"`
}

// CreateBackup - For creating a backup of the Tdarr database.
func (c *Client) CreateBackup() (bool, error) {
	resp, err := c.post(c.baseURL+"/api/v2/create-backup", nil)
	if err != nil {
		return false, err
	}
	return c.decodeBool(resp)
}

// DeleteBackup - For deleting a backup of the Tdarr database.
func (c *Client) DeleteBackup(req DeleteBackupRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL+"/api/v2/delete-backup", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// GetBackupStatus - For getting the status of a Tdarr backup in progress.
func (c *Client) GetBackupStatus() ([]BackupStatus, error) {
	resp, err := c.post(c.baseURL+"/api/v2/get-backup-status", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out []BackupStatus
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetBackups - For getting a list of backups of the Tdarr database.
func (c *Client) GetBackups() ([]Backup, error) {
	resp, err := c.post(c.baseURL+"/api/v2/get-backups", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out []Backup
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResetBackupStatus - For resetting the backup status.
func (c *Client) ResetBackupStatus() (string, error) {
	resp, err := c.post(c.baseURL+"/api/v2/reset-backup-status", nil)
	if err != nil {
		return "", err
	}
	return c.decodeString(resp)
}
