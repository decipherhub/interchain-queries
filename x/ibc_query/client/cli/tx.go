package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	captypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

func NewMsgCrossChainQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [query-path] [query-height] [chain-id] [timeout-timeheight] [timeout-timestamp]",
		Short: "submit an ibc query",
		Long: strings.TrimSpace(`submit an ibc query to queried chain.
		Timeout height and the timeout stamp are the local timeoutHeight and the local timeoutstamp, respectively, 
		which are used to validate the query result`),
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			path := args[0]
			queryHeight, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			chainId := args[2]

			timeoutHeight, err := clienttypes.ParseHeight(args[3])
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitCrossChainQuery(path, timeoutHeight, timeoutTimestamp, queryHeight, chainId, creator)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewMsgPruneCrossChainQueryResultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prune [query-id] [capability-key]",
		Short: "prune ibc query result",
		Long:  strings.TrimSpace(`prune ibc query result`),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			queryId := args[0]
			capKeyIndex, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			capKey := captypes.NewCapability(capKeyIndex)
			msg := types.NewMsgSubmitPruneCrossChainQueryResult(queryId, capKey, creator)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
