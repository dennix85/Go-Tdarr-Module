package tdarr

// Typed response structs for endpoints with well-defined return shapes.

// TokenResponse is returned by login and register endpoints.
type TokenResponse struct {
	Token string `json:"token"`
}

// TempPasswordResponse is returned by admin register/reset-password.
type TempPasswordResponse struct {
	TemporaryPassword string `json:"temporaryPassword"`
}

// VerifyTokenResponse is returned by AuthVerifyToken.
type VerifyTokenResponse struct {
	Roles     []string `json:"roles"`
	LocalAuth bool     `json:"localAuth"`
}

// Backup represents a single database backup entry.
type Backup struct {
	Name     string                 `json:"name"`
	Size     string                 `json:"size"`
	StatSync map[string]interface{} `json:"statSync"`
	Date     int64                  `json:"date"`
}

// BackupStatus represents the status of a backup in progress.
type BackupStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ServerStatus is returned by the /status endpoint.
type ServerStatus struct {
	Status       string `json:"status"`
	IsProduction bool   `json:"isProduction"`
	OS           string `json:"os"`
	Version      string `json:"version"`
	BuildDate    string `json:"buildDate"`
	Uptime       int    `json:"uptime"`
}

// UpdaterCheckResponse is returned by UpdaterCheck.
type UpdaterCheckResponse struct {
	NewVersionAvailable bool   `json:"newVersionAvailable"`
	CurrentVersion      string `json:"currentVersion"`
	RequiredVersion     string `json:"requiredVersion"`
	Message             string `json:"message"`
}

// WorkerLimits holds per-worker-type limits and queue lengths.
type WorkerLimits struct {
	HealthCheckCPU int `json:"healthcheckcpu"`
	HealthCheckGPU int `json:"healthcheckgpu"`
	TranscodeCPU   int `json:"transcodecpu"`
	TranscodeGPU   int `json:"transcodegpu"`
}

// PollWorkerLimitsResponse is returned by PollWorkerLimits.
type PollWorkerLimitsResponse struct {
	WorkerLimits    WorkerLimits `json:"workerLimits"`
	QueueLengths    WorkerLimits `json:"queueLengths"`
	ProcessPriority string       `json:"processPriority"`
}

// ResStats holds server resource statistics.
type ResStats struct {
	Process struct {
		Uptime      int    `json:"uptime"`
		HeapUsedMB  string `json:"heapUsedMB"`
		HeapTotalMB string `json:"heapTotalMB"`
	} `json:"process"`
	OS struct {
		CPUPerc    string `json:"cpuPerc"`
		MemUsedGB  string `json:"memUsedGB"`
		MemTotalGB string `json:"memTotalGB"`
	} `json:"os"`
}

// AutomationResponse is returned by RunAutomation.
type AutomationResponse struct {
	Success bool `json:"success"`
}
