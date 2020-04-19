package staking

import (
	"github.com/SebastianJ/harmony-restaker/config"
	sdkRewards "github.com/harmony-one/go-lib/staking/rewards"
	sdkTxs "github.com/harmony-one/go-lib/transactions"
	"github.com/harmony-one/harmony/numeric"
)

// CollectRewards - collects rewards
func CollectRewards(nonce uint64, gasPrice numeric.Dec) (sdkTxs.Transaction, error) {
	rawTx, err := sdkRewards.CollectRewards(
		config.Configuration.Account.Keystore,
		config.Configuration.Account.Account,
		config.Configuration.Network.API.Shards[0].RPCClient,
		config.Configuration.Network.API.ChainID,
		config.Configuration.Account.Address,
		config.Configuration.Network.Gas.Limit,
		gasPrice,
		nonce,
		config.Configuration.Account.Passphrase,
		config.Configuration.Network.API.NodeAddress(0),
		config.Configuration.Network.Timeout,
	)
	if err != nil {
		return sdkTxs.Transaction{}, err
	}

	tx := sdkTxs.ToTransaction(config.Configuration.Account.Address, 0, config.Configuration.Account.Address, 0, rawTx, err)

	return tx, nil
}
