package staking

import (
	"fmt"
	"time"

	"github.com/SebastianJ/harmony-restaker/config"
	"github.com/SebastianJ/harmony-restaker/logger"
	sdkNetworkNonce "github.com/harmony-one/go-lib/network/rpc/nonces"
	"github.com/harmony-one/harmony/numeric"
	"github.com/pkg/errors"
)

// Restake - collect rewards, then delegate to the target validator
func Restake() error {
	nonce := sdkNetworkNonce.CurrentNonce(config.Configuration.Network.API.Shards[0].RPCClient, config.Configuration.Account.Address)
	gasPrice := config.Configuration.Network.Gas.Price

	logger.StakingLog(fmt.Sprintf("Sending collect rewards transaction for address %s - nonce: %d", config.Configuration.Account.Address, nonce))

	rewardsTx, err := CollectRewards(nonce, gasPrice)
	if err != nil {
		return errors.Wrapf(err, "failed to collect rewards for address %s", config.Configuration.Account.Address)
	}
	logger.StakingLog(fmt.Sprintf("Collect rewards tx hash: %s", rewardsTx.TransactionHash))

	if rewardsTx.Success {
		logger.SuccessLog(fmt.Sprintf("Successfully collected rewards for address: %s!", config.Configuration.Account.Address))
		nonce++
	} else {
		if rewardsTx.Error != nil {
			logger.ErrorLog(fmt.Sprintf("Failed to collect rewards for address: %s - error: %s", config.Configuration.Account.Address, rewardsTx.Error.Error()))
		} else {
			logger.ErrorLog(fmt.Sprintf("Failed to collect rewards for address: %s - maybe you don't have any rewards to collect?", config.Configuration.Account.Address))
		}
	}

	logger.BalanceLog(fmt.Sprintf("Waiting %d seconds for balance to get updated", config.Configuration.Network.Retry.Wait))
	time.Sleep(time.Second * time.Duration(config.Configuration.Network.Retry.Wait))

	logger.BalanceLog(fmt.Sprintf("Checking balance for address %s", config.Configuration.Account.Address))

	balance, err := config.Configuration.Network.API.GetShardBalance(config.Configuration.Account.Address, 0)
	if err != nil {
		return errors.Wrapf(err, "failed to check balance for address %s", config.Configuration.Account.Address)
	}

	logger.BalanceLog(fmt.Sprintf("Balance for address %s after collecting rewards is %f", config.Configuration.Account.Address, balance))
	delegationAmount := balance.Sub(config.Configuration.Network.Gas.Cost)

	if delegationAmount.LT(numeric.NewDec(1000)) {
		logger.WarningLog("You need at least 1000 ONE to delegate - stopping here and waiting for next round of collecting rewards!")
		return nil
	}

	logger.StakingLog(fmt.Sprintf("Subtracting %f as a gas cost for the delegation transaction, will try to restake/delegate a total of %f ONE", config.Configuration.Network.Gas.Cost, delegationAmount))

	logger.StakingLog(fmt.Sprintf("Sending delegation transaction of %f ONE to validator %s - nonce: %d", delegationAmount, config.Configuration.To, nonce))
	delegationTx, err := Delegate(delegationAmount, nonce, gasPrice)
	if err != nil {
		return errors.Wrapf(err, "failed to delegate %f ONE to validator %s", delegationAmount, config.Configuration.To)
	}
	logger.StakingLog(fmt.Sprintf("Delegation tx hash: %s", delegationTx.TransactionHash))

	if delegationTx.Success {
		logger.SuccessLog(fmt.Sprintf("Successfully delegated %f ONE from %s to %s!", delegationAmount, config.Configuration.Account.Address, config.Configuration.To))
		nonce++
	} else {
		if delegationTx.Error != nil {
			logger.ErrorLog(fmt.Sprintf("Failed to delegate %f ONE from %s to %s - error: %s", delegationAmount, config.Configuration.Account.Address, config.Configuration.To, delegationTx.Error.Error()))
		} else {
			logger.ErrorLog(fmt.Sprintf("Failed to delegate %f ONE from %s to %s!", delegationAmount, config.Configuration.Account.Address, config.Configuration.To))
		}
	}

	logger.BalanceLog(fmt.Sprintf("Waiting %d seconds for balance to get updated after delegating", config.Configuration.Network.Retry.Wait))
	time.Sleep(time.Second * time.Duration(config.Configuration.Network.Retry.Wait))

	balance, err = config.Configuration.Network.API.GetShardBalance(config.Configuration.Account.Address, 0)
	if err != nil {
		return errors.Wrapf(err, "failed to check balance for address %s", config.Configuration.Account.Address)
	}

	logger.BalanceLog(fmt.Sprintf("Balance for address %s after delegating is %f", config.Configuration.Account.Address, balance))

	return nil
}
