package evm

import "github.com/ipfs/go-cid"

type ConstructorParams struct {
	Creator  []byte
	Initcode []byte
}

type ResurrectParams = ConstructorParams

type GetStorageAtParams struct {
	StorageKey []byte
}

type DelegateCallParams struct {
	Code  cid.Cid
	Input []byte
}
