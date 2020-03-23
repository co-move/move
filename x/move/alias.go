package move

import (
	"github.com/co-move/move/x/move/internal/keeper"
	"github.com/co-move/move/x/move/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
)
