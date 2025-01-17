package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/nft/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group nft queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryNFTByAddress(),
		CmdQueryAllNft(),
		CmdQueryAllOwners(),
		CmdQueryNFTsByOwner(),
	)
	return cmd
}
