package config

import (
	"github.com/gookit/color"
	sdkAccounts "github.com/harmony-one/go-lib/accounts"
	sdkCommonTypes "github.com/harmony-one/go-lib/network/types/common"
	sdkNetworkTypes "github.com/harmony-one/go-lib/network/types/network"
)

// Config - general config
type Config struct {
	Network  Network
	Account  sdkAccounts.Account
	To       string
	Interval int // In minutes
	Verbose  bool
	Styling  Styling
}

// Network - represents the network settings group
type Network struct {
	Name    string
	Mode    string
	Node    string
	Nodes   []string
	Shards  int
	Gas     sdkNetworkTypes.Gas
	API     sdkNetworkTypes.Network
	Retry   sdkCommonTypes.Retry
	Timeout int
}

// Styling - represents settings for styling the log output
type Styling struct {
	Header      *color.Style
	Info        *color.Style
	Default     *color.Style
	Account     *color.Style
	Funding     *color.Style
	Balance     *color.Style
	Transaction *color.Style
	Staking     *color.Style
	Teardown    *color.Style
	Success     *color.Style
	Warning     *color.Style
	Error       *color.Style
	Padding     string
}
