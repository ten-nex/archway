package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/archway-network/archway/x/callback/types"
)

// Keeper defines the callback module Keeper.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      types.Codec
}

// NewKeeper creates a new callback Keeper instance.
func NewKeeper(cdc types.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// GetCallbacksByHeightPaginated fetches callbacks by height with pagination.
func (k Keeper) GetCallbacksByHeightPaginated(ctx sdk.Context, height int64, pageReq *query.PageRequest) ([]types.Callback, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	callbackStore := prefix.NewStore(store, types.CallbackKey(height))

	var callbacks []types.Callback
	pageRes, err := query.Paginate(callbackStore, pageReq, func(key []byte, value []byte) error {
		var callback types.Callback
		if err := k.cdc.Unmarshal(value, &callback); err != nil {
			return err
		}
		
		callbacks = append(callbacks, callback)
		return nil
	})
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "failed to unmarshal callback: %s", err)
	}

	return callbacks, pageRes, nil
}