package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper( storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}
