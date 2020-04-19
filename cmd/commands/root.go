package commands

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/SebastianJ/harmony-restaker/config"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// VersionWrap - version displayed in case of errors
	VersionWrap = fmt.Sprintf("%s/%s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	// RootCmd - main entry point for Cobra commands
	RootCmd = &cobra.Command{
		Use:          "restake",
		Short:        "Harmony restaker",
		SilenceUsage: true,
		Long:         "Harmony restaker - automatically restake earned validator or delegator rewards",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	config.Args = config.PersistentFlags{}
	RootCmd.PersistentFlags().StringVar(&config.Args.Network, "network", "pangaea", "--network <name>")
	RootCmd.PersistentFlags().StringVar(&config.Args.Mode, "mode", "api", "--mode <mode>")
	RootCmd.PersistentFlags().StringVar(&config.Args.Node, "node", "", "--node <node>")
	RootCmd.PersistentFlags().StringSliceVar(&config.Args.Nodes, "nodes", []string{}, "--nodes node1,node2")

	RootCmd.PersistentFlags().StringVar(&config.Args.From, "from", "", "--from <from>")
	RootCmd.PersistentFlags().StringVar(&config.Args.To, "to", "", "--to <to>")
	RootCmd.PersistentFlags().StringVar(&config.Args.Passphrase, "passphrase", "", "--passphrase <passphrase>")

	RootCmd.PersistentFlags().IntVar(&config.Args.Interval, "interval", 1, "--interval <interval>")

	RootCmd.PersistentFlags().IntVar(&config.Args.Timeout, "timeout", 60, "--timeout <timeout>")
	RootCmd.PersistentFlags().StringVar(&config.Args.GasPrice, "gas.price", "1", "--gas.price <price>")
	RootCmd.PersistentFlags().Int64Var(&config.Args.GasLimit, "gas.limit", -1, "--gas.limit <limit>")   // Let go-lib decide the gas limit
	RootCmd.PersistentFlags().StringVar(&config.Args.GasCost, "gas.cost", "0.001", "--gas.cost <cost>") // Expected gas cost for sending a tx

	RootCmd.PersistentFlags().IntVar(&config.Args.RetryAttempts, "retry.attempts", 6, "--retry.attempts <attempts>")
	RootCmd.PersistentFlags().IntVar(&config.Args.RetryTimeout, "retry.timeout", 10, "--retry.attempts <attempts>")

	RootCmd.PersistentFlags().BoolVar(&config.Args.Verbose, "verbose", false, "--verbose")
	RootCmd.PersistentFlags().BoolVar(&config.Args.VerboseGoSDK, "verbose-go-sdk", false, "--verbose-go-sdk")

	RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(os.Stderr, "SebastianJ (C) 2020. %v, version %s/%s-%s\n", path.Base(os.Args[0]), runtime.Version(), runtime.GOOS, runtime.GOARCH)
			os.Exit(0)
			return nil
		},
	})
}

// ParseArgs - parse arguments using Cobra
func ParseArgs() {
	RootCmd.SilenceErrors = true
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(errors.Wrapf(err, "commit: %s, error", VersionWrap).Error())
		os.Exit(1)
	}
}
