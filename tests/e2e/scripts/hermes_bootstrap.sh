#!/bin/bash

set -ex

# initialize Hermes relayer configuration
mkdir -p /root/.hermes/
touch /root/.hermes/config.toml

# setup Hermes relayer configuration
tee /root/.hermes/config.toml <<EOF
[global]
log_level = 'info'

[mode]

[mode.clients]
enabled = true
refresh = true
misbehaviour = true

[mode.connections]
enabled = false

[mode.channels]
enabled = false

[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true

[rest]
enabled = true
host = '0.0.0.0'
port = 3031

[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001

[[chains]]
id = '$UMEE_E2E_UMEE_CHAIN_ID'
rpc_addr = 'http://$UMEE_E2E_UMEE_VAL_HOST:26657'
grpc_addr = 'http://$UMEE_E2E_UMEE_VAL_HOST:9090'
websocket_addr = 'ws://$UMEE_E2E_UMEE_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'umee'
key_name = 'val01-umee'
store_prefix = 'ibc'
max_gas = 6000000
gas_price = { price = 0.05, denom = 'uumee' }
gas_multiplier = 2
clock_drift = '1m' # to accommodate docker containers
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }

[[chains]]
id = '$UMEE_E2E_GAIA_CHAIN_ID'
rpc_addr = 'http://$UMEE_E2E_GAIA_VAL_HOST:26657'
grpc_addr = 'http://$UMEE_E2E_GAIA_VAL_HOST:9090'
websocket_addr = 'ws://$UMEE_E2E_GAIA_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'cosmos'
key_name = 'val01-gaia'
store_prefix = 'ibc'
max_gas = 6000000
gas_price = { price = 0.001, denom = 'stake' }
gas_multiplier = 2
clock_drift = '1m' # to accommodate docker containers
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
EOF

echo $UMEE_E2E_GAIA_VAL_MNEMONIC > gaia_val_mnemonic.txt
echo $UMEE_E2E_UMEE_VAL_MNEMONIC > umee_val_mnemonic.txt

# import gaia and umee keys
hermes keys add --chain ${UMEE_E2E_GAIA_CHAIN_ID} --key-name "val01-gaia" --mnemonic-file gaia_val_mnemonic.txt
hermes keys add --chain ${UMEE_E2E_UMEE_CHAIN_ID} --key-name "val01-umee" --mnemonic-file umee_val_mnemonic.txt

# start Hermes relayer
hermes start
