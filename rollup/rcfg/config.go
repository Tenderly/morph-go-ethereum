package rcfg

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
)

var (
	// L2MessageQueueAddress is the address of the L2MessageQueue
	// predeploy
	// see contracts/src/L2/predeploys/L2MessageQueue.sol
	L2MessageQueueAddress = common.HexToAddress("0x5300000000000000000000000000000000000001")
	WithdrawTrieRootSlot  = common.BigToHash(big.NewInt(33))

	// SequencerAddress is the address of the Sequencer
	// predeploy
	// set contracts/contracts/l2/staking/Sequencer.sol
	SequencerAddress           = common.HexToAddress("0x5300000000000000000000000000000000000017")
	SequencerSetVerifyHashSlot = common.BigToHash(big.NewInt(101))

	// ScrollFeeVaultAddress is the address of the L2TxFeeVault
	// predeploy
	// see scroll-tech/scroll/contracts/src/L2/predeploys/L2TxFeeVault.sol
	ScrollFeeVaultAddress = common.HexToAddress("0x530000000000000000000000000000000000000a")

	// L1GasPriceOracleAddress is the address of the L1GasPriceOracle
	// predeploy
	// see scroll-tech/scroll/contracts/src/L2/predeploys/L1GasPriceOracle.sol
	L1GasPriceOracleAddress = common.HexToAddress("0x530000000000000000000000000000000000000F")
	Precision               = new(big.Int).SetUint64(1e9)
	L1BaseFeeSlot           = common.BigToHash(big.NewInt(1))
	OverheadSlot            = common.BigToHash(big.NewInt(2))
	ScalarSlot              = common.BigToHash(big.NewInt(3))

	// New fields added in the Curie hard fork
	L1BlobBaseFeeSlot = common.BigToHash(big.NewInt(6))
	CommitScalarSlot  = common.BigToHash(big.NewInt(7))
	BlobScalarSlot    = common.BigToHash(big.NewInt(8))
	IsCurieSlot       = common.BigToHash(big.NewInt(9))

	InitialCommitScalar = big.NewInt(230759955285)
	InitialBlobScalar   = big.NewInt(417565260)
)
