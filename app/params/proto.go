package params

import (
	"cosmossdk.io/simapp/params"
)

// MakeEncodingConfig creates an EncodingConfig for Amino-based tests.
func MakeEncodingConfig() params.EncodingConfig {
	return params.MakeTestEncodingConfig()
}
