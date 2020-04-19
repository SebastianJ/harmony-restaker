# Harmony Restaker
An automatic restaking/redelegation tool for Harmony validators and delegators alike

## Prerequisites
The address you want to use for restaking/redelegation must exist in the keystore on your local machine.

The restaking tool doesn't require hmy - it's entirely standalone.

## Installation

```
rm -rf harmony-restaker && mkdir -p harmony-restaker && cd harmony-restaker
bash <(curl -s -S -L https://raw.githubusercontent.com/SebastianJ/harmony-restaker/master/scripts/install.sh)
```

## Usage

### Restake your earned validator rewards

```
./restaker --network pangaea --from YOUR_VALIDATOR_ADDRESS
```

### Restake all of your earner delegator rewards to a specific validator

```
./restaker --network pangaea --from YOUR_DELEGATOR_ADDRESS --to VALIDATOR_ADDRESS
```

### All options:

```
$ ./restaker --help
Harmony restaker - automatically restake earned validator or delegator rewards

Usage:
  restake [flags]
  restake [command]

Available Commands:
  help        Help about any command
  version     Show version

Flags:
      --from string          --from <from>
      --gas.cost string      --gas.cost <cost> (default "0.001")
      --gas.limit int        --gas.limit <limit> (default -1)
      --gas.price string     --gas.price <price> (default "1")
  -h, --help                 help for restake
      --interval int         --interval <interval> (default 1)
      --mode string          --mode <mode> (default "api")
      --network string       --network <name> (default "pangaea")
      --node string          --node <node>
      --nodes strings        --nodes node1,node2
      --passphrase string    --passphrase <passphrase>
      --retry.attempts int   --retry.attempts <attempts> (default 6)
      --retry.timeout int    --retry.attempts <attempts> (default 10)
      --timeout int          --timeout <timeout> (default 60)
      --to string            --to <to>
      --verbose              --verbose
      --verbose-go-sdk       --verbose-go-sdk

Use "restake [command] --help" for more information about a command.
```
