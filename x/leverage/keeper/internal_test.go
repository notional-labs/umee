package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/umee-network/umee/v6/tests/accs"
	"github.com/umee-network/umee/v6/x/leverage/types"
	ugovmocks "github.com/umee-network/umee/v6/x/ugov/mocks"
)

// TestKeeper is a keeper with some normally
// unexported methods exposed for testing.
type TestKeeper struct {
	*Keeper
}

// NewTestKeeper returns a new leverage keeper, and
// an additional TestKeeper that exposes normally
// unexported methods for testing.
func NewTestKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	akStoreKey storetypes.StoreKey,
	bk types.BankKeeper,
	ak authkeeper.AccountKeeper,
	ok types.OracleKeeper,
	enableLiquidatorQuery bool,
) (Keeper, TestKeeper) {
	k := NewKeeper(
		cdc,
		storeKey,
		akStoreKey,
		bk,
		ak,
		ok,
		ugovmocks.NewSimpleEmergencyGroupBuilder(),
		enableLiquidatorQuery,
		accs.GenerateAddr("auction.Rewards"),
	)
	return k, TestKeeper{&k}
}

func (tk *TestKeeper) SetBadDebtAddress(ctx sdk.Context, addr sdk.AccAddress, denom string, hasDebt bool) error {
	return tk.setBadDebtAddress(ctx, addr, denom, hasDebt)
}

func (tk *TestKeeper) SetBorrow(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) error {
	return tk.setBorrow(ctx, addr, amount)
}

func (tk *TestKeeper) SetCollateral(ctx sdk.Context, addr sdk.AccAddress, collateral sdk.Coin) error {
	return tk.setCollateral(ctx, addr, collateral)
}

func (tk *TestKeeper) SetInterestScalar(ctx sdk.Context, denom string, scalar sdk.Dec) error {
	return tk.setInterestScalar(ctx, denom, scalar)
}

func (tk *TestKeeper) SetReserveAmount(ctx sdk.Context, coin sdk.Coin) error {
	return tk.setReserves(ctx, coin)
}
