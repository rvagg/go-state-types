// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package init

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufState = []byte{130}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Bytecode (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Bytecode); err != nil {
		return xerrors.Errorf("failed to write cid field t.Bytecode: %w", err)
	}

	// t.ContractState (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.ContractState); err != nil {
		return xerrors.Errorf("failed to write cid field t.ContractState: %w", err)
	}

	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bytecode (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Bytecode: %w", err)
		}

		t.Bytecode = c

	}
	// t.ContractState (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ContractState: %w", err)
		}

		t.ContractState = c

	}
	return nil
}

var lengthBufConstructorParams = []byte{130}

func (t *ConstructorParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufConstructorParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Bytecode ([]uint8) (slice)
	if len(t.Bytecode) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Bytecode was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Bytecode))); err != nil {
		return err
	}

	if _, err := w.Write(t.Bytecode[:]); err != nil {
		return err
	}

	// t.InputData ([]uint8) (slice)
	if len(t.InputData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.InputData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.InputData))); err != nil {
		return err
	}

	if _, err := w.Write(t.InputData[:]); err != nil {
		return err
	}
	return nil
}

func (t *ConstructorParams) UnmarshalCBOR(r io.Reader) error {
	*t = ConstructorParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bytecode ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Bytecode: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Bytecode = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Bytecode[:]); err != nil {
		return err
	}
	// t.InputData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.InputData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.InputData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.InputData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufInvokeParams = []byte{129}

func (t *InvokeParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufInvokeParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.InputData ([]uint8) (slice)
	if len(t.InputData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.InputData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.InputData))); err != nil {
		return err
	}

	if _, err := w.Write(t.InputData[:]); err != nil {
		return err
	}
	return nil
}

func (t *InvokeParams) UnmarshalCBOR(r io.Reader) error {
	*t = InvokeParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.InputData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.InputData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.InputData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.InputData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufInvokeReturn = []byte{129}

func (t *InvokeReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufInvokeReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.OutputData ([]uint8) (slice)
	if len(t.OutputData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.OutputData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.OutputData))); err != nil {
		return err
	}

	if _, err := w.Write(t.OutputData[:]); err != nil {
		return err
	}
	return nil
}

func (t *InvokeReturn) UnmarshalCBOR(r io.Reader) error {
	*t = InvokeReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.OutputData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.OutputData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.OutputData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.OutputData[:]); err != nil {
		return err
	}
	return nil
}
