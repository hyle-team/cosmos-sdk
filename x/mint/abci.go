package mint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	// fetch stored minter & params
	params := k.GetParams(ctx)

	// skip if all tokens already minted
	if uint64(ctx.BlockHeight()) >= params.EndBlock {
		return
	}
	monthReward := sdk.NewDecFromInt(params.MonthReward.Amount)
	mintedAmount := monthReward.QuoInt(sdk.NewInt(int64(params.BlocksPerMonth)))

	// mint coins, update supply
	mintedCoin := sdk.NewCoin(params.MintDenom, mintedAmount.TruncateInt())
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.SendFromAccumulator(ctx, mintedCoins)
	if err != nil {
		k.Logger(ctx).Error("failed to send tokens from accumulator")
		return
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
