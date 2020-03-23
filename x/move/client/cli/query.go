package cli

import (
	"github.com/co-move/move/x/move/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	moveQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the Move smart contract module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	moveQueryCmd.AddCommand(flags.GetCommands(
		//GetCmdResolveName(storeKey, cdc),
	)...)

	return moveQueryCmd
}

//// GetCmdResolveName queries information about a name
//func GetCmdResolveName(queryRoute string, cdc *codec.Codec) *cobra.Command {
//	return &cobra.Command{
//		Use:   "resolve [name]",
//		Short: "resolve name",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx := context.NewCLIContext().WithCodec(cdc)
//			name := args[0]
//
//			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)
//			if err != nil {
//				fmt.Printf("could not resolve name - %s \n", name)
//				return nil
//			}
//
//			var out types.QueryResResolve
//			cdc.MustUnmarshalJSON(res, &out)
//			return cliCtx.PrintOutput(out)
//		},
//	}
//}
