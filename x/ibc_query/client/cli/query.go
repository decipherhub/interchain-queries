package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

// GetCmdQueryCrossChainQueryResult defines the command to query CrossChainQueryResult from store
func GetCmdQueryCrossChainQueryResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query-ibc_query-result [query id]",
		Short:   "query cross chain query result with query id",
		Long:    "query cross chain query result with query id",
		Example: fmt.Sprintf("%s query ibc_query query-ibc_query-result 3", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryCrossChainQueryResult{
				Id: args[0],
			}

			res, err := queryClient.CrossChainQueryResult(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
