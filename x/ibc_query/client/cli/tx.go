package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)


func NewMsgCrossChainQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ibc-query [src-port] [src-channel] [query-path] [query-height] [timeout-timeheight] [timeout-timestamp]",
		Short: "submit an ibc query",
		Long:  strings.TrimSpace(`submit an ibc query to queried chain.
		Timeout height and the timeout stamp are the local timeoutHeight and the local timeoutstamp, respectively, 
		which are used to validate the query result`),
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			creator := clientCtx.GetFromAddress().String()

			srcPort := args[0]
			srcChannel := args[1]
			path := args[2]
			queryHeight, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			timeoutHeight, err := clienttypes.ParseHeight( args[4])
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
	

			msg := types.NewMsgSubmitCrossChainQuery("", path, timeoutHeight, timeoutTimestamp, queryHeight, creator, srcPort, srcChannel)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}