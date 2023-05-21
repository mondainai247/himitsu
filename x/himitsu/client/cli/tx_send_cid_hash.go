package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"himitsu/x/himitsu/types"
)

var _ = strconv.Itoa(0)

func CmdSendCidHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-cid-hash [tomodachi] [cid-hash]",
		Short: "Broadcast message send-cid-hash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTomodachi := args[0]
			argCidHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendCidHash(
				clientCtx.GetFromAddress().String(),
				argTomodachi,
				argCidHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
