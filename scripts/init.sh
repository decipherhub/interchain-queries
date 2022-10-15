set -x
rm -rf ~/.interchain-queries

MONIKER=interchain-queries-test
CHAINID=interchain-queries-1
KEYRING=test

MNEMONIC_FILE1=mnemonic1.txt
MNEMONIC_FILE2=mnemonic2.txt
MNEMONIC_FILE3=mnemonic3.txt
MNEMONIC_FILE4=mnemonic4.txt
MNEMONIC_FILE5=mnemonic5.txt
MNEMONIC_FILE6=mnemonic6.txt

interchain-queriesd init $MONIKER --chain-id $CHAINID

interchain-queriesd config keyring-backend $KEYRING
interchain-queriesd config broadcast-mode block

interchain-queriesd keys add genkey --recover < $MNEMONIC_FILE1
interchain-queriesd keys add relayer --recover < $MNEMONIC_FILE2
interchain-queriesd keys add boosik --recover < $MNEMONIC_FILE3
interchain-queriesd keys add sangyun --recover < $MNEMONIC_FILE4
interchain-queriesd keys add soojin --recover < $MNEMONIC_FILE5
interchain-queriesd keys add wujinger --recover < $MNEMONIC_FILE6

interchain-queriesd add-genesis-account $(interchain-queriesd keys show genkey -a) 10000000000uiq
interchain-queriesd add-genesis-account $(interchain-queriesd keys show relayer -a) 10000000000uiq
interchain-queriesd add-genesis-account $(interchain-queriesd keys show boosik -a) 10000000000uiq
interchain-queriesd add-genesis-account $(interchain-queriesd keys show sangyun -a) 10000000000uiq
interchain-queriesd add-genesis-account $(interchain-queriesd keys show soojin -a) 10000000000uiq
interchain-queriesd add-genesis-account $(interchain-queriesd keys show wujinger -a) 10000000000uiq

interchain-queriesd gentx genkey 1000000000uiq --chain-id $CHAINID


interchain-queriesd collect-gentxs

# update staking genesis
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uiq"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["staking"]["params"]["max_entries"]="10"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json

# update mint denom
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uiq"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json

# update crisis variable to uiq
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uiq"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json

# udpate gov genesis
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uiq"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json
cat $HOME/.interchain-queries/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uiq"' > $HOME/.interchain-queries/config/tmp_genesis.json && mv $HOME/.interchain-queries/config/tmp_genesis.json $HOME/.interchain-queries/config/genesis.json
