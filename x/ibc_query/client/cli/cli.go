package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for ibc_query
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "ibc_query",
		Short:                      "IBC query command",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdQueryCrossChainQueryResult(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for ibc_query
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "ibc_query",
		Short:                      "Query cross chain query result",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewMsgCrossChainQueryCmd(),
	)

	return txCmd
}
