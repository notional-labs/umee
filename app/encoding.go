package app

import (
	sdkparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/umee-network/umee/v4/app/params"
)

// MakeEncodingConfig returns the application's encoding configuration with all
// types and interfaces registered.
func MakeEncodingConfig() sdkparams.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
