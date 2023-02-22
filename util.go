package wanode_near

import (
	"github.com/Marreck/wacore"
	"time"
)

var _ wacore.BlockchainUtils = Util{} // Make sure Util implements wacore.BlockchainUtils

const (
	AverageBlockTime = time.Second * 1 // Average time between blocks
	GenesisBlock     = 0               // First block height
	BlockchainID     = 15              // Unique id for this blockchain
)

// Util is an empty struct that is used to implement interfaces from wacore.
// External packages can call functions like the address formatting methods that are specific for this blockchain.
type Util struct{}

func (u Util) BaseSymbol() wacore.Symbol {
	return NEAR
}

var Utils Util

// Name returns the blockchain Name.
// This implements part of the wacore.BlockchainConfig interface.
func (Util) Name() string {
	return Name
}

// ID returns the blockchain id.
// This implements part of the wacore.BlockchainConfig interface.
func (Util) ID() wacore.BlockchainID {
	return BlockchainID
}

// AvgBlockTime returns the average time between two consecutive blocks.
// This implements part of the wacore.BlockchainConfig interface.
func (Util) AvgBlockTime() time.Duration {
	return AverageBlockTime
}

// GenesisBlock returns the height of the first block in this blockchain
// This implements part of the wacore.BlockchainConfig interface.
func (Util) GenesisBlock() uint64 {
	return GenesisBlock
}
