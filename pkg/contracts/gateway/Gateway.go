// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EntryPointMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type EntryPointMemoryUserOp struct {
	Sender               common.Address
	Nonce                *big.Int
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	Paymaster            common.Address
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []UserOperation
	Aggregator common.Address
	Signature  []byte
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"targetSuccess\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"targetResult\",\"type\":\"bytes\"}],\"name\":\"ExecutionResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResult\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPoint.AggregatorStakeInfo\",\"name\":\"aggregatorInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResultWithAggregation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"GatewayInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SIG_VALIDATION_FAILED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"_validateSenderAndPaymaster\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"simulateHandleOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simulateValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052604051620000129062000183565b604051809103905ff0801580156200002c573d5f803e3d5ffd5b506001600160a01b031660805234801562000045575f80fd5b506001600255620000563362000066565b62000060620000c1565b62000191565b600380546001600160a01b038381166201000081810262010000600160b01b031985161790945560405193909204169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600354610100900460ff16156200012e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60035460ff908116101562000181576003805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6101f980620042af83390190565b6080516140fe620001b15f395f81816113e60152612f5901526140fe5ff3fe608060405260043610610128575f3560e01c80630396cb601461013c5780630bd28e3b1461014f57806313af40351461016e5780631b2e01b81461018d5780631d732756146101d65780631fad948c146101f5578063205c287814610214578063319775c41461023357806335567e1a146102525780634b1d7cf5146102715780635287ce121461029057806370a08231146103a9578063715018a6146103c85780638da5cb5b146103dc5780638f41ec5a146103fd578063957122ab146104115780639b249f6914610430578063a61935311461044f578063b760faf91461046e578063bb9fe6bf14610481578063c23a5cea14610495578063d6383f94146104b4578063ee219423146104d3578063f2fde38b146104f2578063fc7e286d14610511575f80fd5b3661013857610136336105c3565b005b5f80fd5b61013661014a36600461324e565b610629565b34801561015a575f80fd5b5061013661016936600461328c565b6108a4565b348015610179575f80fd5b506101366101883660046132c4565b6108da565b348015610198575f80fd5b506101c36101a73660046132df565b600160209081525f928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b3480156101e1575f80fd5b506101c36101f03660046134cd565b6108ee565b348015610200575f80fd5b5061013661020f3660046135ca565b610a4f565b34801561021f575f80fd5b5061013661022e36600461361c565b610bb1565b34801561023e575f80fd5b5061013661024d3660046132c4565b610d1c565b34801561025d575f80fd5b506101c361026c3660046132df565b610e62565b34801561027c575f80fd5b5061013661028b3660046135ca565b610ea6565b34801561029b575f80fd5b506103516102aa3660046132c4565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160a01b03165f9081526020818152604091829020825160a08101845281546001600160701b038082168352600160701b820460ff16151594830194909452600160781b90049092169282019290925260019091015463ffffffff81166060830152600160201b900465ffffffffffff16608082015290565b6040805182516001600160701b03908116825260208085015115159083015283830151169181019190915260608083015163ffffffff169082015260809182015165ffffffffffff169181019190915260a0016101cd565b3480156103b4575f80fd5b506101c36103c33660046132c4565b61128e565b3480156103d3575f80fd5b506101366112b1565b3480156103e7575f80fd5b506103f06112c4565b6040516101cd9190613646565b348015610408575f80fd5b506101c3600181565b34801561041c575f80fd5b5061013661042b36600461365a565b6112d9565b34801561043b575f80fd5b5061013661044a3660046136d8565b6113cd565b34801561045a575f80fd5b506101c361046936600461372d565b61147a565b61013661047c3660046132c4565b6105c3565b34801561048c575f80fd5b506101366114bb565b3480156104a0575f80fd5b506101366104af3660046132c4565b6115e0565b3480156104bf575f80fd5b506101366104ce36600461375e565b6117ff565b3480156104de575f80fd5b506101366104ed36600461372d565b6118ee565b3480156104fd575f80fd5b5061013661050c3660046132c4565b611aa6565b34801561051c575f80fd5b5061057d61052b3660046132c4565b5f60208190529081526040902080546001909101546001600160701b0380831692600160701b810460ff1692600160781b9091049091169063ffffffff811690600160201b900465ffffffffffff1685565b604080516001600160701b0396871681529415156020860152929094169183019190915263ffffffff16606082015265ffffffffffff909116608082015260a0016101cd565b6105cd8134611b1c565b6001600160a01b0381165f8181526020818152604091829020805492516001600160701b03909316835292917f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c491015b60405180910390a25050565b335f90815260208190526040902063ffffffff821661068c5760405162461bcd60e51b815260206004820152601a6024820152796d757374207370656369667920756e7374616b652064656c617960301b60448201526064015b60405180910390fd5b600181015463ffffffff90811690831610156106e95760405162461bcd60e51b815260206004820152601c60248201527b63616e6e6f7420646563726561736520756e7374616b652074696d6560201b6044820152606401610683565b80545f90610708903490600160781b90046001600160701b03166137ce565b90505f811161074e5760405162461bcd60e51b81526020600482015260126024820152711b9bc81cdd185ad9481cdc1958da599a595960721b6044820152606401610683565b6001600160701b038111156107965760405162461bcd60e51b815260206004820152600e60248201526d7374616b65206f766572666c6f7760901b6044820152606401610683565b6040805160a08101825283546001600160701b0390811682526001602080840182815286841685870190815263ffffffff808b16606088019081525f608089018181523380835296829052908a902098518954955194518916600160781b02600160781b600160e81b0319951515600160701b026001600160781b03199097169190991617949094179290921695909517865551949092018054925165ffffffffffff16600160201b026001600160501b0319909316949093169390931717905590517fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c0190610897908490879091825263ffffffff16602082015260400190565b60405180910390a2505050565b335f9081526001602090815260408083206001600160c01b038516845290915281208054916108d2836137e1565b919050555050565b6108e2611bb7565b6108eb81611aa6565b50565b5f805a905033301461093c5760405162461bcd60e51b81526020600482015260176024820152764141393220696e7465726e616c2063616c6c206f6e6c7960481b6044820152606401610683565b8451604081015160608201518101611388015a10156109645763deaddead60e01b5f5260205ffd5b87515f90156109f2575f61097d845f01515f8c86611c16565b9050806109f0575f610990610800611c2c565b8051909150156109ea57845f01516001600160a01b03168a602001517f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a2018760200151846040516109e1929190613846565b60405180910390a35b60019250505b505b5f88608001515a8603019050610a415f838b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250611c57915050565b9a9950505050505050505050565b610a57611f41565b815f816001600160401b03811115610a7157610a71613312565b604051908082528060200260200182016040528015610aaa57816020015b610a97613199565b815260200190600190039081610a8f5790505b5090505f5b82811015610b1f575f828281518110610aca57610aca61385e565b602002602001015190505f80610b04848a8a87818110610aec57610aec61385e565b9050602002810190610afe9190613872565b85611f98565b91509150610b148483835f61216e565b505050600101610aaf565b506040515f905f805160206140a9833981519152908290a15f5b83811015610b9457610b8881888884818110610b5757610b5761385e565b9050602002810190610b699190613872565b858481518110610b7b57610b7b61385e565b6020026020010151612302565b90910190600101610b39565b50610b9f8482612422565b505050610bac6001600255565b505050565b335f90815260208190526040902080546001600160701b0316821115610c155760405162461bcd60e51b8152602060048201526019602482015278576974686472617720616d6f756e7420746f6f206c6172676560381b6044820152606401610683565b8054610c2b9083906001600160701b0316613891565b81546001600160701b0319166001600160701b039190911617815560405133907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb90610c7a90869086906138a4565b60405180910390a25f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114610ccb576040519150601f19603f3d011682016040523d82523d5f602084013e610cd0565b606091505b5050905080610d165760405162461bcd60e51b81526020600482015260126024820152716661696c656420746f20776974686472617760701b6044820152606401610683565b50505050565b600354610100900460ff1615808015610d3c5750600354600160ff909116105b80610d565750303b158015610d56575060035460ff166001145b610db95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610683565b6003805460ff191660011790558015610ddc576003805461ff0019166101001790555b610de582611aa6565b6040516001600160a01b038316907ff10bfb1f400dfca17da71024877aa92dc6afd7f01bf79f55bfd2ed169a3a841f905f90a28015610e5e576003805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6001600160a01b0382165f9081526001602090815260408083206001600160c01b038516845290915290819020549082901b6001600160401b031916175b92915050565b610eae611f41565b815f805b8281101561100d5736868683818110610ecd57610ecd61385e565b9050602002810190610edf91906138bd565b9050365f610eed83806138d1565b90925090505f610f0360408501602086016132c4565b90505f196001600160a01b03821601610f585760405162461bcd60e51b815260206004820152601760248201527620a09c9b1034b73b30b634b21030b3b3b932b3b0ba37b960491b6044820152606401610683565b6001600160a01b03811615610fea576001600160a01b03811663e3563a4f8484610f856040890189613916565b6040518563ffffffff1660e01b8152600401610fa49493929190613ab7565b5f6040518083038186803b158015610fba575f80fd5b505afa925050508015610fcb575060015b610fea578060405163086a9f7560e41b81526004016106839190613646565b610ff482876137ce565b9550505050508080611005906137e1565b915050610eb2565b505f816001600160401b0381111561102757611027613312565b60405190808252806020026020018201604052801561106057816020015b61104d613199565b8152602001906001900390816110455790505b506040519091505f805160206140a9833981519152905f90a15f805b8481101561115b57368888838181106110975761109761385e565b90506020028101906110a991906138bd565b9050365f6110b783806138d1565b90925090505f6110cd60408501602086016132c4565b9050815f5b81811015611142575f8989815181106110ed576110ed61385e565b602002602001015190505f8061110f8b898987818110610aec57610aec61385e565b9150915061111f8483838961216e565b8a611129816137e1565b9b5050505050808061113a906137e1565b9150506110d2565b5050505050508080611153906137e1565b91505061107c565b505f8091505f5b8581101561125d573689898381811061117d5761117d61385e565b905060200281019061118f91906138bd565b90506111a160408201602083016132c4565b6001600160a01b03165f8051602061408983398151915260405160405180910390a2365f6111cf83806138d1565b9092509050805f5b8181101561124557611219888585848181106111f5576111f561385e565b90506020028101906112079190613872565b8b8b81518110610b7b57610b7b61385e565b61122390886137ce565b96508761122f816137e1565b985050808061123d906137e1565b9150506111d7565b50505050508080611255906137e1565b915050611162565b506040515f905f80516020614089833981519152908290a261127f8682612422565b5050505050610bac6001600255565b6001600160a01b03165f908152602081905260409020546001600160701b031690565b6112b9611bb7565b6112c25f612512565b565b6003546201000090046001600160a01b031690565b831580156112ef57506001600160a01b0383163b155b156113385760405162461bcd60e51b815260206004820152601960248201527810504c8c081858d8dbdd5b9d081b9bdd0819195c1b1bde5959603a1b6044820152606401610683565b601481106113ac575f61134e6014828486613b33565b61135791613b5a565b60601c9050803b5f036113aa5760405162461bcd60e51b815260206004820152601b60248201527a10504ccc081c185e5b585cdd195c881b9bdd0819195c1b1bde5959602a1b6044820152606401610683565b505b60405162461bcd60e51b8152602060048201525f6024820152604401610683565b604051632b870d1b60e11b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063570e1a369061141d9086908690600401613b8a565b6020604051808303815f875af1158015611439573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061145d9190613b9d565b905080604051633653dc0360e11b81526004016106839190613646565b5f6114848261256d565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b335f9081526020819052604081206001810154909163ffffffff90911690036115135760405162461bcd60e51b815260206004820152600a6024820152691b9bdd081cdd185ad95960b21b6044820152606401610683565b8054600160701b900460ff1661155f5760405162461bcd60e51b8152602060048201526011602482015270616c726561647920756e7374616b696e6760781b6044820152606401610683565b60018101545f906115769063ffffffff1642613bb8565b60018301805465ffffffffffff60201b1916600160201b65ffffffffffff841690810291909117909155835460ff60701b1916845560405190815290915033907ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a9060200161061d565b335f9081526020819052604090208054600160781b90046001600160701b0316806116445760405162461bcd60e51b81526020600482015260146024820152734e6f207374616b6520746f20776974686472617760601b6044820152606401610683565b6001820154600160201b900465ffffffffffff166116a45760405162461bcd60e51b815260206004820152601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000006044820152606401610683565b600182015442600160201b90910465ffffffffffff1611156117065760405162461bcd60e51b815260206004820152601b60248201527a5374616b65207769746864726177616c206973206e6f742064756560281b6044820152606401610683565b6001820180546001600160501b03191690558154600160781b600160e81b031916825560405133907fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda39061175d90869085906138a4565b60405180910390a25f836001600160a01b0316826040515f6040518083038185875af1925050503d805f81146117ae576040519150601f19603f3d011682016040523d82523d5f602084013e6117b3565b606091505b5050905080610d165760405162461bcd60e51b81526020600482015260186024820152776661696c656420746f207769746864726177207374616b6560401b6044820152606401610683565b611807613199565b61181085612585565b5f8061181d5f8885611f98565b915091505f61182c8383612653565b9050611836435f52565b5f6118425f8a87612302565b905061184c435f52565b5f60606001600160a01b038a16156118bd57896001600160a01b03168989604051611878929190613bde565b5f604051808303815f865af19150503d805f81146118b1576040519150601f19603f3d011682016040523d82523d5f602084013e6118b6565b606091505b5090925090505b866080015183856020015186604001518585604051630116f59360e71b815260040161068396959493929190613bed565b6118f6613199565b6118ff82612585565b5f8061190c5f8585611f98565b915091505f611921845f015160a0015161270a565b8451519091505f906119329061270a565b905061193c613217565b365f61194b60408a018a613916565b90925090505f6014821015611960575f61197a565b61196d60145f8486613b33565b61197691613b5a565b60601c5b90506119858161270a565b93505050505f6119958686612653565b90505f815f015190505f60016001600160a01b0316826001600160a01b03161490505f6040518060c001604052808b6080015181526020018b6040015181526020018315158152602001856020015165ffffffffffff168152602001856040015165ffffffffffff168152602001611a0e8c6060015190565b905290506001600160a01b03831615801590611a3457506001600160a01b038316600114155b15611a85575f6040518060400160405280856001600160a01b03168152602001611a5d8661270a565b81525090508187878a84604051633ebb2d3960e21b8152600401610683959493929190613c9c565b8086868960405163e0cff05f60e01b81526004016106839493929190613d07565b611aae611bb7565b6001600160a01b038116611b135760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610683565b6108eb81612512565b6001600160a01b0382165f9081526020819052604081208054909190611b4c9084906001600160701b03166137ce565b90506001600160701b03811115611b985760405162461bcd60e51b815260206004820152601060248201526f6465706f736974206f766572666c6f7760801b6044820152606401610683565b81546001600160701b0319166001600160701b03919091161790555050565b33611bc06112c4565b6001600160a01b0316146112c25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610683565b5f805f845160208601878987f195945050505050565b60603d82811115611c3a5750815b604051602082018101604052818152815f602083013e9392505050565b5f805a85519091505f9081611c6b82612751565b60a08301519091506001600160a01b038116611c8a5782519350611e28565b8093505f88511115611e2857868202955060028a6002811115611caf57611caf613d4b565b14611d1c57606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611ce9908e908d908c90600401613d5f565b5f604051808303815f88803b158015611d00575f80fd5b5087f1158015611d12573d5f803e3d5ffd5b5050505050611e28565b606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611d51908e908d908c90600401613d5f565b5f604051808303815f88803b158015611d68575f80fd5b5087f193505050508015611d7a575060015b611e2857611d86613da3565b806308c379a003611ddf5750611d9a613dbc565b80611da55750611de1565b8b81604051602001611db79190613e44565b60408051601f1981840301815290829052631101335b60e11b82526106839291600401613846565b505b8a604051631101335b60e11b81526004016106839181526040602082018190526012908201527110504d4c081c1bdcdd13dc081c995d995c9d60721b606082015260800190565b5a85038701965081870295508589604001511015611e91578a604051631101335b60e11b815260040161068391815260406020808301829052908201527f414135312070726566756e642062656c6f772061637475616c476173436f7374606082015260800190565b6040890151869003611ea38582611b1c565b5f808c6002811115611eb757611eb7613d4b565b1490508460a001516001600160a01b0316855f01516001600160a01b03168c602001517f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f8860200151858d8f604051611f29949392919093845291151560208401526040830152606082015260800190565b60405180910390a45050505050505095945050505050565b6002805403611f925760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610683565b60028055565b5f805f5a8451909150611fab8682612780565b611fb48661147a565b6020860152604081015160608201516080830151171760e087013517610100870135176001600160781b038111156120295760405162461bcd60e51b815260206004820152601860248201527741413934206761732076616c756573206f766572666c6f7760401b6044820152606401610683565b5f8061203484612876565b90506120428a8a8a846128c1565b855160208701519199509193506120599190612ad6565b6120ac5789604051631101335b60e11b8152600401610683918152604060208201819052601a90820152794141323520696e76616c6964206163636f756e74206e6f6e636560301b606082015260800190565b6120b4435f52565b60a08401516060906001600160a01b0316156120dc576120d78b8b8b8587612b22565b975090505b5f5a87039050808b60a001351015612140578b604051631101335b60e11b8152600401610683918152604060208201819052601e908201527f41413430206f76657220766572696669636174696f6e4761734c696d69740000606082015260800190565b60408a018390528160608b015260c08b01355a8803018a608001818152505050505050505050935093915050565b5f8061217985612d3e565b91509150816001600160a01b0316836001600160a01b0316146121df5785604051631101335b60e11b81526004016106839181526040602082018190526014908201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b80156122315785604051631101335b60e11b815260040161068391815260406020820181905260179082015276414132322065787069726564206f72206e6f742064756560481b606082015260800190565b5f61223b85612d3e565b925090506001600160a01b038116156122975786604051631101335b60e11b81526004016106839181526040602082018190526014908201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b81156122f95786604051631101335b60e11b81526004016106839181526040602082018190526021908201527f41413332207061796d61737465722065787069726564206f72206e6f742064756060820152606560f81b608082015260a00190565b50505050505050565b5f805a90505f612313846060015190565b905030631d7327566123286060880188613916565b87856040518563ffffffff1660e01b81526004016123499493929190613e81565b6020604051808303815f875af1925050508015612383575060408051601f3d908101601f1916820190925261238091810190613f33565b60015b612416575f60205f803e505f51632152215360e01b81016123e25786604051631101335b60e11b8152600401610683918152604060208201819052600f908201526e41413935206f7574206f662067617360881b606082015260800190565b5f85608001515a6123f39086613891565b6123fd91906137ce565b905061240d886002888685611c57565b94505050612419565b92505b50509392505050565b6001600160a01b0382166124735760405162461bcd60e51b81526020600482015260186024820152774141393020696e76616c69642062656e656669636961727960401b6044820152606401610683565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146124bc576040519150601f19603f3d011682016040523d82523d5f602084013e6124c1565b606091505b5050905080610bac5760405162461bcd60e51b815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606401610683565b600380546001600160a01b038381166201000081810262010000600160b01b031985161790945560405193909204169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f61257782612d8d565b805190602001209050919050565b3063957122ab6125986040840184613916565b6125a560208601866132c4565b6125b3610120870187613916565b6040518663ffffffff1660e01b81526004016125d3959493929190613f4a565b5f6040518083038186803b1580156125e9575f80fd5b505afa9250505080156125fa575060015b6108eb57612606613da3565b806308c379a003612649575061261a613dbc565b80612625575061264b565b805115610e5e575f81604051631101335b60e11b8152600401610683929190613846565b505b3d5f803e3d5ffd5b61265b61322f565b5f61266584612e5d565b90505f61267184612e5d565b82519091506001600160a01b038116612688575080515b602080840151604080860151928501519085015191929165ffffffffffff80831690851610156126b6578193505b8065ffffffffffff168365ffffffffffff1611156126d2578092505b5050604080516060810182526001600160a01b03909416845265ffffffffffff92831660208501529116908201529250505092915050565b612712613217565b6001600160a01b039091165f9081526020818152604090912080546001600160701b03600160781b9091041683526001015463ffffffff169082015290565b60c081015160e08201515f919080820361276c575092915050565b61277882488301612eb8565b949350505050565b61278d60208301836132c4565b6001600160a01b0316815260208083013590820152608080830135604083015260a0830135606083015260c0808401359183019190915260e0808401359183019190915261010083013590820152365f6127eb610120850185613916565b9092509050801561286a5760148110156128475760405162461bcd60e51b815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606401610683565b61285460145f8385613b33565b61285d91613b5a565b60601c60a0840152610d16565b5f60a084015250505050565b60a08101515f9081906001600160a01b0316612893576001612896565b60035b60ff1690505f8360800151828560600151028560400151010190508360c00151810292505050919050565b5f805f5a85518051919250906128e489886128df60408c018c613916565b612ecf565b60a08201516128f1435f52565b5f6001600160a01b038216612921575f61290a8461128e565b905088811161291b5780890361291d565b5f5b9150505b606084015160208a0151604051633a871cdd60e01b81526001600160a01b03861692633a871cdd92909161295b918f918790600401613f7f565b6020604051808303815f8887f193505050508015612996575060408051601f3d908101601f1916820190925261299391810190613f33565b60015b612a20576129a2613da3565b806308c379a0036129d357506129b6613dbc565b806129c157506129d5565b8b81604051602001611db79190613fa3565b505b8a604051631101335b60e11b8152600401610683918152604060208201819052601690820152754141323320726576657274656420286f72204f4f472960501b606082015260800190565b95506001600160a01b038216612ac3576001600160a01b0383165f90815260208190526040902080546001600160701b0316808a1115612aa6578c604051631101335b60e11b81526004016106839181526040602082018190526017908201527610504c8c48191a591b89dd081c185e481c1c99599d5b99604a1b606082015260800190565b81546001600160701b031916908a90036001600160701b03161790555b5a85039650505050505094509492505050565b6001600160a01b0382165f90815260016020908152604080832084821c80855292528220805484916001600160401b038316919085612b14836137e1565b909155501495945050505050565b825160608181015190915f91848111612b7d5760405162461bcd60e51b815260206004820152601f60248201527f4141343120746f6f206c6974746c6520766572696669636174696f6e476173006044820152606401610683565b60a08201516001600160a01b0381165f90815260208190526040902080548784039291906001600160701b031689811015612c04578c604051631101335b60e11b8152600401610683918152604060208201819052601e908201527f41413331207061796d6173746572206465706f73697420746f6f206c6f770000606082015260800190565b898103825f015f6101000a8154816001600160701b0302191690836001600160701b03160217905550826001600160a01b031663f465c77e858e8e602001518e6040518563ffffffff1660e01b8152600401612c6293929190613f7f565b5f604051808303815f8887f193505050508015612ca057506040513d5f823e601f3d908101601f19168201604052612c9d9190810190613fd9565b60015b612d2a57612cac613da3565b806308c379a003612cdd5750612cc0613dbc565b80612ccb5750612cdf565b8d81604051602001611db7919061405f565b505b8c604051631101335b60e11b8152600401610683918152604060208201819052601690820152754141333320726576657274656420286f72204f4f472960501b606082015260800190565b909e909d509b505050505050505050505050565b5f80825f03612d5157505f928392509050565b5f612d5b84612e5d565b9050806040015165ffffffffffff16421180612d825750806020015165ffffffffffff1642105b905194909350915050565b6060813560208301355f612dac612da76040870187613916565b613187565b90505f612dbf612da76060880188613916565b9050608086013560a087013560c088013560e08901356101008a01355f612ded612da76101208e018e613916565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b612e6561322f565b8160a081901c65ffffffffffff81165f03612e83575065ffffffffffff5b604080516060810182526001600160a01b03909316835260d09490941c602083015265ffffffffffff16928101929092525090565b5f818310612ec65781612ec8565b825b9392505050565b8015610d16578251516001600160a01b0381163b15612f3a5784604051631101335b60e11b8152600401610683918152604060208201819052601f908201527f414131302073656e64657220616c726561647920636f6e737472756374656400606082015260800190565b835160600151604051632b870d1b60e11b81525f916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163570e1a369190612f919088908890600401613b8a565b6020604051808303815f8887f1158015612fad573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190612fd29190613b9d565b90506001600160a01b0381166130325785604051631101335b60e11b8152600401610683918152604060208201819052601b908201527a4141313320696e6974436f6465206661696c6564206f72204f4f4760281b606082015260800190565b816001600160a01b0316816001600160a01b03161461309c5785604051631101335b60e11b815260040161068391815260406020808301829052908201527f4141313420696e6974436f6465206d7573742072657475726e2073656e646572606082015260800190565b806001600160a01b03163b5f036130fe5785604051631101335b60e11b815260040161068391815260406020808301829052908201527f4141313520696e6974436f6465206d757374206372656174652073656e646572606082015260800190565b5f61310c6014828688613b33565b61311591613b5a565b60601c9050826001600160a01b031686602001517fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d83895f015160a001516040516131769291906001600160a01b0392831681529116602082015260400190565b60405180910390a350505050505050565b5f604051828085833790209392505050565b6040518060a001604052806131f66040518061010001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f6001600160a01b031681526020015f81526020015f81525090565b81526020015f80191681526020015f81526020015f81526020015f81525090565b60405180604001604052805f81526020015f81525090565b604080516060810182525f808252602082018190529181019190915290565b5f6020828403121561325e575f80fd5b813563ffffffff81168114612ec8575f80fd5b80356001600160c01b0381168114613287575f80fd5b919050565b5f6020828403121561329c575f80fd5b612ec882613271565b6001600160a01b03811681146108eb575f80fd5b8035613287816132a5565b5f602082840312156132d4575f80fd5b8135612ec8816132a5565b5f80604083850312156132f0575f80fd5b82356132fb816132a5565b915061330960208401613271565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b60a081016001600160401b038111828210171561334557613345613312565b60405250565b61010081016001600160401b038111828210171561334557613345613312565b601f8201601f191681016001600160401b038111828210171561339057613390613312565b6040525050565b5f6001600160401b038211156133af576133af613312565b50601f01601f191660200190565b5f8183036101808112156133cf575f80fd5b6040516133db81613326565b809250610100808312156133ed575f80fd5b60405192506133fb8361334b565b613404856132b9565b83526020850135602084015260408501356040840152606085013560608401526080850135608084015261343a60a086016132b9565b60a084015260c085013560c084015260e085013560e084015282825280850135602083015250610120840135604082015261014084013560608201526101608401356080820152505092915050565b5f8083601f840112613499575f80fd5b5081356001600160401b038111156134af575f80fd5b6020830191508360208285010111156134c6575f80fd5b9250929050565b5f805f806101c085870312156134e1575f80fd5b84356001600160401b03808211156134f7575f80fd5b818701915087601f83011261350a575f80fd5b813561351581613397565b604051613522828261336b565b8281528a6020848701011115613536575f80fd5b826020860160208301375f6020848301015280985050505061355b88602089016133bd565b94506101a0870135915080821115613571575f80fd5b5061357e87828801613489565b95989497509550505050565b5f8083601f84011261359a575f80fd5b5081356001600160401b038111156135b0575f80fd5b6020830191508360208260051b85010111156134c6575f80fd5b5f805f604084860312156135dc575f80fd5b83356001600160401b038111156135f1575f80fd5b6135fd8682870161358a565b9094509250506020840135613611816132a5565b809150509250925092565b5f806040838503121561362d575f80fd5b8235613638816132a5565b946020939093013593505050565b6001600160a01b0391909116815260200190565b5f805f805f6060868803121561366e575f80fd5b85356001600160401b0380821115613684575f80fd5b61369089838a01613489565b9097509550602088013591506136a5826132a5565b909350604087013590808211156136ba575f80fd5b506136c788828901613489565b969995985093965092949392505050565b5f80602083850312156136e9575f80fd5b82356001600160401b038111156136fe575f80fd5b61370a85828601613489565b90969095509350505050565b5f6101608284031215613727575f80fd5b50919050565b5f6020828403121561373d575f80fd5b81356001600160401b03811115613752575f80fd5b61277884828501613716565b5f805f8060608587031215613771575f80fd5b84356001600160401b0380821115613787575f80fd5b61379388838901613716565b9550602087013591506137a5826132a5565b90935060408601359080821115613571575f80fd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610ea057610ea06137ba565b5f600182016137f2576137f26137ba565b5060010190565b5f5b838110156138135781810151838201526020016137fb565b50505f910152565b5f81518084526138328160208601602086016137f9565b601f01601f19169290920160200192915050565b828152604060208201525f612778604083018461381b565b634e487b7160e01b5f52603260045260245ffd5b5f823561015e19833603018112613887575f80fd5b9190910192915050565b81810381811115610ea057610ea06137ba565b6001600160a01b03929092168252602082015260400190565b5f8235605e19833603018112613887575f80fd5b5f808335601e198436030181126138e6575f80fd5b8301803591506001600160401b038211156138ff575f80fd5b6020019150600581901b36038213156134c6575f80fd5b5f808335601e1984360301811261392b575f80fd5b8301803591506001600160401b03821115613944575f80fd5b6020019150368190038213156134c6575f80fd5b5f808335601e1984360301811261396d575f80fd5b83016020810192503590506001600160401b0381111561398b575f80fd5b8036038213156134c6575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f6101606139df846139d2856132b9565b6001600160a01b03169052565b602083013560208501526139f66040840184613958565b826040870152613a098387018284613999565b92505050613a1a6060840184613958565b8583036060870152613a2d838284613999565b925050506080830135608085015260a083013560a085015260c083013560c085015260e083013560e0850152610100808401358186015250610120613a7481850185613958565b86840383880152613a86848284613999565b9350505050610140613a9a81850185613958565b86840383880152613aac848284613999565b979650505050505050565b604080825281018490525f6060600586901b830181019083018783805b89811015613b1c57868503605f190184528235368c900361015e19018112613afa578283fd5b613b06868d83016139c1565b9550506020938401939290920191600101613ad4565b505050508281036020840152613aac818587613999565b5f8085851115613b41575f80fd5b83861115613b4d575f80fd5b5050820193919092039150565b6001600160601b03198135818116916014851015613b825780818660140360031b1b83161692505b505092915050565b602081525f612778602083018486613999565b5f60208284031215613bad575f80fd5b8151612ec8816132a5565b65ffffffffffff818116838216019080821115613bd757613bd76137ba565b5092915050565b818382375f9101908152919050565b8681528560208201525f65ffffffffffff8087166040840152808616606084015250831515608083015260c060a0830152613c2b60c083018461381b565b98975050505050505050565b80518252602081015160208301526040810151151560408301525f606082015165ffffffffffff8082166060860152806080850151166080860152505060a082015160c060a085015261277860c085018261381b565b80518252602090810151910152565b5f610140808352613caf81840189613c37565b915050613cbf6020830187613c8d565b613ccc6060830186613c8d565b613cd960a0830185613c8d565b82516001600160a01b031660e08301526020830151613cfc610100840182613c8d565b509695505050505050565b60e081525f613d1960e0830187613c37565b9050613d286020830186613c8d565b613d356060830185613c8d565b613d4260a0830184613c8d565b95945050505050565b634e487b7160e01b5f52602160045260245ffd5b5f60038510613d7c57634e487b7160e01b5f52602160045260245ffd5b84825260606020830152613d93606083018561381b565b9050826040830152949350505050565b5f60033d1115613db95760045f803e505f5160e01c5b90565b5f60443d1015613dc95790565b6040516003193d81016004833e81513d6001600160401b038083116024840183101715613df857505050505090565b8285019150815181811115613e105750505050505090565b843d8701016020828501011115613e2a5750505050505090565b613e396020828601018761336b565b509095945050505050565b75020a09a98103837b9ba27b8103932bb32b93a32b21d160551b81525f8251613e748160168501602087016137f9565b9190910160160192915050565b5f6101c0808352613e958184018789613999565b9050845160018060a01b03808251166020860152602082015160408601526040820151606086015260608201516080860152608082015160a08601528060a08301511660c08601525060c081015160e085015260e08101516101008501525060208501516101208401526040850151610140840152606085015161016084015260808501516101808401528281036101a0840152613aac818561381b565b5f60208284031215613f43575f80fd5b5051919050565b606081525f613f5d606083018789613999565b6001600160a01b03861660208401528281036040840152613c2b818587613999565b606081525f613f9160608301866139c1565b60208301949094525060400152919050565b6e020a09919903932bb32b93a32b21d1608d1b81525f8251613fcc81600f8501602087016137f9565b91909101600f0192915050565b5f8060408385031215613fea575f80fd5b82516001600160401b03811115613fff575f80fd5b8301601f8101851361400f575f80fd5b805161401a81613397565b604051614027828261336b565b82815287602084860101111561403b575f80fd5b61404c8360208301602087016137f9565b6020969096015195979596505050505050565b6e020a09999903932bb32b93a32b21d1608d1b81525f8251613fcc81600f8501602087016137f956fe575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4dbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972a2646970667358221220b3db0396cac47f85ad001a4a1ee92e05d471d9b51b39b922a7bbe78d3305f63464736f6c63430008140033608060405234801561000f575f80fd5b506101dc8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063570e1a361461002d575b5f80fd5b61004061003b3660046100e4565b61005c565b6040516001600160a01b03909116815260200160405180910390f35b5f8061006b601482858761014f565b61007491610176565b60601c90505f610087846014818861014f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92018290525084519495509360209350849250905082850182875af190505f519350806100db575f93505b50505092915050565b5f80602083850312156100f5575f80fd5b82356001600160401b038082111561010b575f80fd5b818501915085601f83011261011e575f80fd5b81358181111561012c575f80fd5b86602082850101111561013d575f80fd5b60209290920196919550909350505050565b5f808585111561015d575f80fd5b83861115610169575f80fd5b5050820193919092039150565b6001600160601b0319813581811691601485101561019e5780818660140360031b1b83161692505b50509291505056fea26469706673582212201b2bf7acb234a7b63fa6f2487163a7793ce00d08ac218dc9001ef3ae71eb96c564736f6c63430008140033",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMetaData.Bin instead.
var GatewayBin = GatewayMetaData.Bin

// DeployGateway deploys a new Ethereum contract, binding an instance of Gateway to it.
func DeployGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gateway, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewayCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewaySession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Gateway.Contract.SIGVALIDATIONFAILED(&_Gateway.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewayCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Gateway.Contract.SIGVALIDATIONFAILED(&_Gateway.CallOpts)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewayCaller) ValidateSenderAndPaymaster(opts *bind.CallOpts, initCode []byte, sender common.Address, paymasterAndData []byte) error {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "_validateSenderAndPaymaster", initCode, sender, paymasterAndData)

	if err != nil {
		return err
	}

	return err

}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewaySession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Gateway.Contract.ValidateSenderAndPaymaster(&_Gateway.CallOpts, initCode, sender, paymasterAndData)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewayCallerSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Gateway.Contract.ValidateSenderAndPaymaster(&_Gateway.CallOpts, initCode, sender, paymasterAndData)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewayCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewaySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gateway.Contract.BalanceOf(&_Gateway.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewayCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gateway.Contract.BalanceOf(&_Gateway.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewayCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewaySession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Gateway.Contract.Deposits(&_Gateway.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewayCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Gateway.Contract.Deposits(&_Gateway.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewayCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewaySession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _Gateway.Contract.GetDepositInfo(&_Gateway.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewayCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _Gateway.Contract.GetDepositInfo(&_Gateway.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewayCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewaySession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetNonce(&_Gateway.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewayCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetNonce(&_Gateway.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewayCaller) GetUserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewaySession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Gateway.Contract.GetUserOpHash(&_Gateway.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewayCallerSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Gateway.Contract.GetUserOpHash(&_Gateway.CallOpts, userOp)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewayCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewaySession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gateway.Contract.NonceSequenceNumber(&_Gateway.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewayCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gateway.Contract.NonceSequenceNumber(&_Gateway.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewaySession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCallerSession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x319775c4.
//
// Solidity: function _initialize(address _owner) returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "_initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x319775c4.
//
// Solidity: function _initialize(address _owner) returns()
func (_Gateway *GatewaySession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x319775c4.
//
// Solidity: function _initialize(address _owner) returns()
func (_Gateway *GatewayTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _owner)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewayTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewaySession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.Contract.AddStake(&_Gateway.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewayTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.Contract.AddStake(&_Gateway.TransactOpts, unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewayTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewaySession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.DepositTo(&_Gateway.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewayTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.DepositTo(&_Gateway.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewayTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewaySession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Gateway.Contract.GetSenderAddress(&_Gateway.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewayTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Gateway.Contract.GetSenderAddress(&_Gateway.TransactOpts, initCode)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewayTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewaySession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleAggregatedOps(&_Gateway.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewayTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleAggregatedOps(&_Gateway.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewayTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewaySession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleOps(&_Gateway.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewayTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleOps(&_Gateway.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewayTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewaySession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.IncrementNonce(&_Gateway.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewayTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.IncrementNonce(&_Gateway.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewayTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewaySession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.Contract.InnerHandleOp(&_Gateway.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewayTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.Contract.InnerHandleOp(&_Gateway.TransactOpts, callData, opInfo, context)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Gateway *GatewayTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Gateway *GatewaySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetOwner(&_Gateway.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Gateway *GatewayTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetOwner(&_Gateway.TransactOpts, _owner)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewayTransactor) SimulateHandleOp(opts *bind.TransactOpts, op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewaySession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateHandleOp(&_Gateway.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewayTransactorSession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateHandleOp(&_Gateway.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewayTransactor) SimulateValidation(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "simulateValidation", userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewaySession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateValidation(&_Gateway.TransactOpts, userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewayTransactorSession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateValidation(&_Gateway.TransactOpts, userOp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewayTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewaySession) UnlockStake() (*types.Transaction, error) {
	return _Gateway.Contract.UnlockStake(&_Gateway.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewayTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Gateway.Contract.UnlockStake(&_Gateway.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewayTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewaySession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawStake(&_Gateway.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewayTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawStake(&_Gateway.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewayTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewaySession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawTo(&_Gateway.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewayTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawTo(&_Gateway.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewaySession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// GatewayAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the Gateway contract.
type GatewayAccountDeployedIterator struct {
	Event *GatewayAccountDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayAccountDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayAccountDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayAccountDeployed represents a AccountDeployed event raised by the Gateway contract.
type GatewayAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GatewayAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayAccountDeployedIterator{contract: _Gateway.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *GatewayAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayAccountDeployed)
				if err := _Gateway.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) ParseAccountDeployed(log types.Log) (*GatewayAccountDeployed, error) {
	event := new(GatewayAccountDeployed)
	if err := _Gateway.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the Gateway contract.
type GatewayBeforeExecutionIterator struct {
	Event *GatewayBeforeExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayBeforeExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayBeforeExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayBeforeExecution represents a BeforeExecution event raised by the Gateway contract.
type GatewayBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*GatewayBeforeExecutionIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &GatewayBeforeExecutionIterator{contract: _Gateway.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *GatewayBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayBeforeExecution)
				if err := _Gateway.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeforeExecution is a log parse operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) ParseBeforeExecution(log types.Log) (*GatewayBeforeExecution, error) {
	event := new(GatewayBeforeExecution)
	if err := _Gateway.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Gateway contract.
type GatewayDepositedIterator struct {
	Event *GatewayDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDeposited represents a Deposited event raised by the Gateway contract.
type GatewayDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*GatewayDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayDepositedIterator{contract: _Gateway.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDeposited)
				if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) ParseDeposited(log types.Log) (*GatewayDeposited, error) {
	event := new(GatewayDeposited)
	if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayGatewayInitializedIterator is returned from FilterGatewayInitialized and is used to iterate over the raw logs and unpacked data for GatewayInitialized events raised by the Gateway contract.
type GatewayGatewayInitializedIterator struct {
	Event *GatewayGatewayInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayGatewayInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayGatewayInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayGatewayInitialized represents a GatewayInitialized event raised by the Gateway contract.
type GatewayGatewayInitialized struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGatewayInitialized is a free log retrieval operation binding the contract event 0xf10bfb1f400dfca17da71024877aa92dc6afd7f01bf79f55bfd2ed169a3a841f.
//
// Solidity: event GatewayInitialized(address indexed owner)
func (_Gateway *GatewayFilterer) FilterGatewayInitialized(opts *bind.FilterOpts, owner []common.Address) (*GatewayGatewayInitializedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "GatewayInitialized", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayGatewayInitializedIterator{contract: _Gateway.contract, event: "GatewayInitialized", logs: logs, sub: sub}, nil
}

// WatchGatewayInitialized is a free log subscription operation binding the contract event 0xf10bfb1f400dfca17da71024877aa92dc6afd7f01bf79f55bfd2ed169a3a841f.
//
// Solidity: event GatewayInitialized(address indexed owner)
func (_Gateway *GatewayFilterer) WatchGatewayInitialized(opts *bind.WatchOpts, sink chan<- *GatewayGatewayInitialized, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "GatewayInitialized", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayGatewayInitialized)
				if err := _Gateway.contract.UnpackLog(event, "GatewayInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGatewayInitialized is a log parse operation binding the contract event 0xf10bfb1f400dfca17da71024877aa92dc6afd7f01bf79f55bfd2ed169a3a841f.
//
// Solidity: event GatewayInitialized(address indexed owner)
func (_Gateway *GatewayFilterer) ParseGatewayInitialized(log types.Log) (*GatewayGatewayInitialized, error) {
	event := new(GatewayGatewayInitialized)
	if err := _Gateway.contract.UnpackLog(event, "GatewayInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gateway contract.
type GatewayInitializedIterator struct {
	Event *GatewayInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayInitialized represents a Initialized event raised by the Gateway contract.
type GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gateway *GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayInitializedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayInitializedIterator{contract: _Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gateway *GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayInitialized)
				if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gateway *GatewayFilterer) ParseInitialized(log types.Log) (*GatewayInitialized, error) {
	event := new(GatewayInitialized)
	if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gateway contract.
type GatewayOwnershipTransferredIterator struct {
	Event *GatewayOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the Gateway contract.
type GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOwnershipTransferredIterator{contract: _Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOwnershipTransferred)
				if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayOwnershipTransferred, error) {
	event := new(GatewayOwnershipTransferred)
	if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewaySignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the Gateway contract.
type GatewaySignatureAggregatorChangedIterator struct {
	Event *GatewaySignatureAggregatorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewaySignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewaySignatureAggregatorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewaySignatureAggregatorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewaySignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewaySignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewaySignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the Gateway contract.
type GatewaySignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*GatewaySignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &GatewaySignatureAggregatorChangedIterator{contract: _Gateway.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *GatewaySignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewaySignatureAggregatorChanged)
				if err := _Gateway.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) ParseSignatureAggregatorChanged(log types.Log) (*GatewaySignatureAggregatorChanged, error) {
	event := new(GatewaySignatureAggregatorChanged)
	if err := _Gateway.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the Gateway contract.
type GatewayStakeLockedIterator struct {
	Event *GatewayStakeLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeLocked represents a StakeLocked event raised by the Gateway contract.
type GatewayStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeLockedIterator{contract: _Gateway.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *GatewayStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeLocked)
				if err := _Gateway.contract.UnpackLog(event, "StakeLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) ParseStakeLocked(log types.Log) (*GatewayStakeLocked, error) {
	event := new(GatewayStakeLocked)
	if err := _Gateway.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the Gateway contract.
type GatewayStakeUnlockedIterator struct {
	Event *GatewayStakeUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeUnlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeUnlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeUnlocked represents a StakeUnlocked event raised by the Gateway contract.
type GatewayStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeUnlockedIterator{contract: _Gateway.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *GatewayStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeUnlocked)
				if err := _Gateway.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) ParseStakeUnlocked(log types.Log) (*GatewayStakeUnlocked, error) {
	event := new(GatewayStakeUnlocked)
	if err := _Gateway.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Gateway contract.
type GatewayStakeWithdrawnIterator struct {
	Event *GatewayStakeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeWithdrawn represents a StakeWithdrawn event raised by the Gateway contract.
type GatewayStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeWithdrawnIterator{contract: _Gateway.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeWithdrawn)
				if err := _Gateway.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) ParseStakeWithdrawn(log types.Log) (*GatewayStakeWithdrawn, error) {
	event := new(GatewayStakeWithdrawn)
	if err := _Gateway.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the Gateway contract.
type GatewayUserOperationEventIterator struct {
	Event *GatewayUserOperationEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUserOperationEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayUserOperationEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUserOperationEvent represents a UserOperationEvent event raised by the Gateway contract.
type GatewayUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*GatewayUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayUserOperationEventIterator{contract: _Gateway.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *GatewayUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUserOperationEvent)
				if err := _Gateway.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationEvent is a log parse operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) ParseUserOperationEvent(log types.Log) (*GatewayUserOperationEvent, error) {
	event := new(GatewayUserOperationEvent)
	if err := _Gateway.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the Gateway contract.
type GatewayUserOperationRevertReasonIterator struct {
	Event *GatewayUserOperationRevertReason // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUserOperationRevertReason)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayUserOperationRevertReason)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUserOperationRevertReason represents a UserOperationRevertReason event raised by the Gateway contract.
type GatewayUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GatewayUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayUserOperationRevertReasonIterator{contract: _Gateway.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *GatewayUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUserOperationRevertReason)
				if err := _Gateway.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) ParseUserOperationRevertReason(log types.Log) (*GatewayUserOperationRevertReason, error) {
	event := new(GatewayUserOperationRevertReason)
	if err := _Gateway.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Gateway contract.
type GatewayWithdrawnIterator struct {
	Event *GatewayWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawn represents a Withdrawn event raised by the Gateway contract.
type GatewayWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*GatewayWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawnIterator{contract: _Gateway.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawn)
				if err := _Gateway.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) ParseWithdrawn(log types.Log) (*GatewayWithdrawn, error) {
	event := new(GatewayWithdrawn)
	if err := _Gateway.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
