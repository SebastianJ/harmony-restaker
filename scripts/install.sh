#!/usr/bin/env bash

echo "Installing Restaker"
curl -LOs http://tools.harmony.one.s3.amazonaws.com/release/linux-x86_64/restaker && chmod u+x restaker
echo "Harmony Restaker is now ready to use!"
echo "Invoke it using ./restaker - see ./restaker --help for all available options!"
