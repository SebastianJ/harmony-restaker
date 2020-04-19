package config

// PersistentFlags represents the persistent flags
type PersistentFlags struct {
	Network       string
	Mode          string
	Node          string
	Nodes         []string
	From          string
	To            string
	Passphrase    string
	Timeout       int
	Interval      int
	Verbose       bool
	VerboseGoSDK  bool
	GasCost       string
	GasPrice      string
	GasLimit      int64
	RetryAttempts int
	RetryTimeout  int
}
