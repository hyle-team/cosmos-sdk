package v046_28

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// MigrateStore performs in-place store migrations from v0.43/v0.44/v0.45 to v0.46.
// The migration includes:
//
// - Setting the MinDelegationAmount param in the paramstore
func MigrateStore(ctx sdk.Context, _ storetypes.StoreKey, _ codec.BinaryCodec, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)

	return nil
}

func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	if paramstore.HasKeyTable() {
		paramstore.Set(ctx, types.KeyMinCommissionRate, types.DefaultMinCommissionRate)
		return
	}
	paramstore.WithKeyTable(types.ParamKeyTable())
	paramstore.Set(ctx, types.KeyMinCommissionRate, types.DefaultMinCommissionRate)
}
