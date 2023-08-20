package domain

type (
	ThreadsConfig interface {
		IsEnableMultiThreads() bool // IsEnableMultiThreads return if application is allowed multi thread
		GetMaximumThreads() int     // GetMaximumThreads return the maximum number of the application how many useable threads
	}

	FileConfig interface {
		GetDefaultDirectory() (string, error) // GetDefaultDirectory return the destination folder
	}

	LogConfig interface {
		GetDefaultLogDirectory() string // GetDefaultLogDirectory return the path of default log directory
		GetBaseLogFileName() string     // GetBaseLogFileName return the name of log file
	}

	AppConfig struct {
		EnableThreads       bool   `json:"enable_threads"`     // EnableThreads is a option to use threads
		MaximumThreads      int    `json:"maximum_threads"`    // MaximumThreads is a maximum number to be used
		DefaultDirectory    string `json:"default_target_dir"` // DefaultDirectory is a path to move file
		DefaultLogDirectory string `json:"default_log_dir"`    // DefaultLogDirectory is a path to save log
		LogFileName         string `json:"log_file_name"`      // LogFileName is a name of log file
	}
)
