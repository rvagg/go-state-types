package ethaccount

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(value *abi.EmptyValue) *abi.EmptyValue)},
	builtin.MustGenerateFRCMethodNum("AuthenticateMessage"): {"AuthenticateMessage", *new(func(*AuthenticateMessageParams) *abi.EmptyValue)},  // AuthenticateMessage
	builtin.MustGenerateFRCMethodNum("Receive"):             {"UniversalReceiverHook", *new(func(*abi.CborBytesTransparent) *abi.EmptyValue)}, // UniversalReceiverHook
}
