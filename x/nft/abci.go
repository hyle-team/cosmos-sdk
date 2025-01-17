package nft

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/x/nft/types"
	"time"
)

// update vesting state for each nft
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	for _, nft := range k.GetNFTs(ctx) {
		if ctx.BlockTime().Unix()-nft.LastVestingTime < nft.VestingPeriod {
			continue
		}

		if nft.VestingCounter >= nft.VestingPeriodsCount {
			continue
		}

		nft.AvailableToWithdraw = nft.AvailableToWithdraw.Add(sdk.NewCoin(nft.Denom, nft.RewardPerPeriod.Amount))
		nft.VestingCounter++
		nft.LastVestingTime = ctx.BlockTime().Unix()

		k.SetNFT(ctx, nft)
	}

}
