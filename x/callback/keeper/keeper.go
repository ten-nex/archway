package keeper

import (
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/query"
    "github.com/archway-network/archway/x/callback/types"
)

// Keeper struct
type Keeper struct {
    cdc      codec.Codec
    storeKey sdk.StoreKey
}

// NewKeeper creates a new keeper instance.
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) Keeper {
    return Keeper{
        cdc:      cdc,
        storeKey: storeKey,
    }
}

// GetCallbacksByHeightPaginated retrieves callbacks by height with pagination
func (k Keeper) GetCallbacksByHeightPaginated(ctx sdk.Context, height int64, pagination *query.PageRequest) ([]types.Callback, *query.PageResponse, error) {
    store := ctx.KVStore(k.storeKey)
    callbackStore := prefix.NewStore(store, types.GetCallbackPrefix(height))

    var callbacks []types.Callback
    pageRes, err := query.Paginate(callbackStore, pagination, func(key []byte, value []byte) error {
        var callback types.Callback
        if err := k.cdc.Unmarshal(value, &callback); err != nil {
            return err
        }
        callbacks = append(callbacks, callback)
        return nil
    })
    if err != nil {
        return nil, nil, err
    }

    return callbacks, pageRes, nil
}