package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v6/x/ugov"

	"github.com/umee-network/umee/v6/x/metoken"
)

// Builder constructs Keeper by preparing all related dependencies (notably the store).
type Builder struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	bankKeeper     metoken.BankKeeper
	leverageKeeper metoken.LeverageKeeper
	oracleKeeper   metoken.OracleKeeper
	ugov           ugov.EmergencyGroupBuilder
	rewardsAuction sdk.AccAddress
}

// NewBuilder returns Builder object.
func NewBuilder(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper metoken.BankKeeper,
	leverageKeeper metoken.LeverageKeeper,
	oracleKeeper metoken.OracleKeeper,
	ugov ugov.EmergencyGroupBuilder,
	rewardsAuction sdk.AccAddress,
) Builder {
	return Builder{
		cdc:            cdc,
		storeKey:       storeKey,
		bankKeeper:     bankKeeper,
		leverageKeeper: leverageKeeper,
		oracleKeeper:   oracleKeeper,
		ugov:           ugov,
		rewardsAuction: rewardsAuction,
	}
}

type Keeper struct {
	cdc            codec.BinaryCodec
	store          sdk.KVStore
	bankKeeper     metoken.BankKeeper
	leverageKeeper metoken.LeverageKeeper
	oracleKeeper   metoken.OracleKeeper
	ugov           ugov.EmergencyGroupBuilder
	rewardsAuction sdk.AccAddress

	// TODO: ctx should be removed when we migrate leverageKeeper and oracleKeeper
	ctx *sdk.Context
}

// Keeper creates a new Keeper object
func (b Builder) Keeper(ctx *sdk.Context) Keeper {
	return Keeper{
		cdc:            b.cdc,
		store:          ctx.KVStore(b.storeKey),
		bankKeeper:     b.bankKeeper,
		leverageKeeper: b.leverageKeeper,
		oracleKeeper:   b.oracleKeeper,
		ugov:           b.ugov,
		rewardsAuction: b.rewardsAuction,
		ctx:            ctx,
	}
}

// Logger returns module Logger
func (k Keeper) Logger() log.Logger {
	return k.ctx.Logger().With("module", "x/"+metoken.ModuleName)
}
