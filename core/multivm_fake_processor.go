// +build !sputnikvm

package core

import (
	"math/big"

	"github.com/Victorium-org/go-victorium/core/state"
	"github.com/Victorium-org/go-victorium/core/types"
	evm "github.com/Victorium-org/go-victorium/core/vm"
)

const SputnikVMExists = false
var UseSputnikVM = false

func ApplyMultiVmTransaction(config *ChainConfig, bc *BlockChain, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, totalUsedGas *big.Int) (*types.Receipt, evm.Logs, *big.Int, error) {
	panic("not implemented")
}
