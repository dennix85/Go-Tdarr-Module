package tdarr

// Worker process direction for AlterWorkerLimit.
const (
	ProcessIncrease = "increase"
	ProcessDecrease = "decrease"
)

// WorkerType identifies the kind of worker on a node.
const (
	WorkerTypeHealthCheckCPU = "healthcheckcpu"
	WorkerTypeHealthCheckGPU = "healthcheckgpu"
	WorkerTypeTranscodeCPU   = "transcodecpu"
	WorkerTypeTranscodeGPU   = "transcodegpu"
)

// ScanMode controls how a library scan is run.
const (
	ScanModeFresh         = "scanFresh"
	ScanModeFindNew       = "scanFindNew"
	ScanModeFolderWatcher = "scanFolderWatcher"
)

// CrudCollection names the available database collections.
const (
	CollectionSettingsGlobal       = "SettingsGlobalJSONDB"
	CollectionLibrarySettings      = "LibrarySettingsJSONDB"
	CollectionNode                 = "NodeJSONDB"
	CollectionStatistics           = "StatisticsJSONDB"
	CollectionVariables            = "VariablesJSONDB"
	CollectionUsers                = "UsersJSONDB"
	CollectionApiKeys              = "ApiKeysJSONDB"
	CollectionFlows                = "FlowsJSONDB"
	CollectionStaged               = "StagedJSONDB"
	CollectionFile                 = "FileJSONDB"
	CollectionF2FOutput            = "F2FOutputJSONDB"
	CollectionWorkerVerdictHistory = "WorkerVerdictHistoryJSONDB"
	CollectionJobs                 = "JobsJSONDB"
	CollectionAutomations          = "AutomationsJSONDB"
)

// CrudMode names the available database operations.
const (
	CrudModeGetByID    = "getById"
	CrudModeGetByIndex = "getByIndex"
	CrudModeGetAll     = "getAll"
	CrudModeInsert     = "insert"
	CrudModeUpdate     = "update"
	CrudModeIncDec     = "incdec"
	CrudModeRemoveOne  = "removeOne"
	CrudModeRemoveByDB = "removeByDB"
	CrudModeRemoveAll  = "removeAll"
)
