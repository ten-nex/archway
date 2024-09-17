package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/callback interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&Msg{}, "callback/Msg", nil)
	cdc.RegisterConcrete(&Callback{}, "callback/Callback", nil)
}

// RegisterInterfaces registers implementations for protobuf Any.
func RegisterInterfaces(registry codec.TypesRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
	)
}