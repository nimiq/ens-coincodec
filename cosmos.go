package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	hrp = "cosmos"
)

func init() {
	toBytesMap[slip44.ATOM] = CosmosDecodeToBytes
	toStringMap[slip44.ATOM] = CosmosEncodeToString
}

// Converts the input string to a byte array
func CosmosDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrp, input)
	return bytes, err
}

// Converts the input byte array to a string representation of the Cosmos address.
func CosmosEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrp, bytes)
	return output, err
}
