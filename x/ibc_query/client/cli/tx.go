package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/client/utils"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
)

func NewMsgCrossChainQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cross-chain-query [src-port] [src-channel] [query-path] [query-height]",
		Short: "Request ibc query on a given channel.",
		Long:  strings.TrimSpace(`Register a payee address on a given channel.`),
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			queryId, err := utils.GetQueryIdentifier()
			if err != nil {
				return err
			}

			srcPort := args[0]
			srcChannel := args[1]
			path := args[2]
			queryHeight, _ := strconv.ParseUint(args[2], 10, 64)

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitCrossChainQuery(queryId, path, timeoutHeight, timeoutTimestamp, queryHeight, creator, srcPort, srcChannel)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
