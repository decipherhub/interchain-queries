package simapp

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/interchain-queries/app"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v5/testing"
	mock "github.com/cosmos/ibc-go/v5/testing/mock"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := tmdb.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	testApp := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, simapp.DefaultNodeHome, 5, encoding, simapp.EmptyAppOptions{}).(ibctesting.TestingApp)
	return testApp, app.NewDefaultGenesisState(encoding.Marshaler)
}

// NewCoordinator initializes Coordinator with 2 TestChains
func NewCrossChainQueryCoordinator(t *testing.T) (*ibctesting.Coordinator, *ibctesting.TestChain, *ibctesting.TestChain) {
	ibctesting.DefaultTestingAppInit = SetupTestingApp
	coordinator := ibctesting.NewCoordinator(t, 0)
	queryingChain := NewTestChain(t, coordinator, ibctesting.GetChainID(1))
	queriedChain := NewTestChain(t, coordinator, ibctesting.GetChainID(2))
	return coordinator, queryingChain, queriedChain
}

// NewTestChain initializes a new test chain with a default of 4 validators
// Use this function if the tests do not need custom control over the validator set
func NewTestChain(t *testing.T, coord *ibctesting.Coordinator, chainID string) *ibctesting.TestChain {
	// generate validators private/public key
	var (
		validatorsPerChain = 4
		validators         []*tmtypes.Validator
		signersByAddress   = make(map[string]tmtypes.PrivValidator, validatorsPerChain)
	)

	for i := 0; i < validatorsPerChain; i++ {
		privVal := mock.NewPV()
		pubKey, err := privVal.GetPubKey()
		require.NoError(t, err)
		validators = append(validators, tmtypes.NewValidator(pubKey, 1))
		signersByAddress[pubKey.Address().String()] = privVal
	}

	// construct validator set;
	// Note that the validators are sorted by voting power
	// or, if equal, by address lexical order
	valSet := tmtypes.NewValidatorSet(validators)

	return NewTestChainWithValSet(t, coord, chainID, valSet, signersByAddress)
}

func NewTestChainWithValSet(t *testing.T, coord *ibctesting.Coordinator, chainID string, valSet *tmtypes.ValidatorSet, signers map[string]tmtypes.PrivValidator) *ibctesting.TestChain {
	genAccs := []authtypes.GenesisAccount{}
	genBals := []banktypes.Balance{}
	senderAccs := []ibctesting.SenderAccount{}

	var MaxAccounts = 10

	// generate genesis accounts
	for i := 0; i < MaxAccounts; i++ {
		senderPrivKey := secp256k1.GenPrivKey()
		acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), uint64(i), 0)
		amount, ok := sdk.NewIntFromString("10000000000000000000")
		require.True(t, ok)

		// add sender account
		balance := banktypes.Balance{
			Address: acc.GetAddress().String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amount)),
		}

		genAccs = append(genAccs, acc)
		genBals = append(genBals, balance)

		senderAcc := ibctesting.SenderAccount{
			SenderAccount: acc,
			SenderPrivKey: senderPrivKey,
		}

		senderAccs = append(senderAccs, senderAcc)
	}

	// add mock module account balance
	genBals = append(genBals, banktypes.Balance{
		Address: authtypes.NewModuleAddress(mock.ModuleName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000000))},
	})

	app := ibctesting.SetupWithGenesisValSet(t, valSet, genAccs, chainID, sdk.DefaultPowerReduction, genBals...)

	// create current header and call begin block
	header := tmproto.Header{
		ChainID: chainID,
		Height:  1,
		Time:    coord.CurrentTime.UTC(),
	}

	txConfig := app.GetTxConfig()

	// create an account to send transactions from
	chain := &ibctesting.TestChain{
		T:              t,
		Coordinator:    coord,
		ChainID:        chainID,
		App:            app,
		CurrentHeader:  header,
		QueryServer:    app.GetIBCKeeper(),
		TxConfig:       txConfig,
		Codec:          app.AppCodec(),
		Vals:           valSet,
		NextVals:       valSet,
		Signers:        signers,
		SenderPrivKey:  senderAccs[0].SenderPrivKey,
		SenderAccount:  senderAccs[0].SenderAccount,
		SenderAccounts: senderAccs,
	}

	coord.CommitBlock(chain)

	return chain
}
