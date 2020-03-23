package cli

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/co-move/move/x/move/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	moveTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Move smart contract transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	moveTxCmd.AddCommand(flags.PostCommands(
		GetCmdRumMoveScript(cdc),
	)...)

	return moveTxCmd
}

// GetCmdRumMoveScript is the CLI command for sending a BuyName transaction
func GetCmdRumMoveScript(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "run-script [script] [args]",
		Short: "Run a Move transaction script on Cosmos",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			data, err1 := ioutil.ReadFile(args[0])
			if err1 != nil {
				return err1
			}

			fmt.Println("mv file:", hex.EncodeToString(data))

			msg := types.NewMsgRunMoveScript(hex.EncodeToString(data), args[1:], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
