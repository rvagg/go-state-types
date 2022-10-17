package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

var Methods = map[uint64]interface{}{
	1: Constructor,   // Constructor
	2: PubkeyAddress, // PubkeyAddress
}

func Constructor(interface{}, *address.Address) *abi.EmptyValue {
	return nil
}

func PubkeyAddress(interface{}, *abi.EmptyValue) *address.Address {
	return nil
}
