package main

import (
	"fmt"
	"log"
	"os"
	"time"

	cmd "github.com/SebastianJ/harmony-restaker/cmd/commands"
	"github.com/SebastianJ/harmony-restaker/config"
	"github.com/SebastianJ/harmony-restaker/logger"
	"github.com/SebastianJ/harmony-restaker/staking"
)

func main() {
	// Force usage of Go's own DNS implementation
	os.Setenv("GODEBUG", "netdns=go")

	if err := setup(); err != nil {
		log.Fatalln(err)
	}

	run()
}

func setup() error {
	cmd.ParseArgs()

	if err := config.Configure(); err != nil {
		return err
	}

	return nil
}

func run() {
	logger.Title()

	for {
		if err := staking.Restake(); err != nil {
			logger.ErrorLog(fmt.Sprintf("Failed to perform restaking for account %s - error: %s\n", config.Configuration.Account.Address, err.Error()))
		}

		logger.InfoLog(fmt.Sprintf("Waiting %d minute(s) until next restaking attempt\n", config.Configuration.Interval))
		time.Sleep(time.Duration(config.Configuration.Interval) * time.Minute)
	}
}
