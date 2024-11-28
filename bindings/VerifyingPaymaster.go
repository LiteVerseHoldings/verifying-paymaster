// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifyingPaymaster

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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// VerifyingPaymasterPaymasterData is an auto generated low-level Go binding around an user-defined struct.
type VerifyingPaymasterPaymasterData struct {
	ValidUntil         *big.Int
	ValidAfter         *big.Int
	SponsorUUID        *big.Int
	AllowAnyBundler    bool
	PrecheckBalance    bool
	PrepaymentRequired bool
	Token              common.Address
	Receiver           common.Address
	ExchangeRate       *big.Int
	PostOpGas          *big.Int
}

// VerifyingPaymasterMetaData contains all meta data concerning the VerifyingPaymaster contract.
var VerifyingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"entryPoint\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"},{\"name\":\"initialVerifyingSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStake\",\"inputs\":[{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entryPoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHash\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"paymasterData\",\"type\":\"tuple\",\"internalType\":\"structVerifyingPaymaster.PaymasterData\",\"components\":[{\"name\":\"validUntil\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"validAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"sponsorUUID\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"allowAnyBundler\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"precheckBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"prepaymentRequired\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exchangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"postOpGas\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"name\":\"paymasterValidationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"postOpGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isBundlerAllowed\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerWithdrawERC20\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"parsePaymasterData\",\"inputs\":[{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"paymasterData\",\"type\":\"tuple\",\"internalType\":\"structVerifyingPaymaster.PaymasterData\",\"components\":[{\"name\":\"validUntil\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"validAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"sponsorUUID\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"allowAnyBundler\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"precheckBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"prepaymentRequired\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exchangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"postOpGas\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingVerifyingSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postOp\",\"inputs\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumIPaymaster.PostOpMode\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rotateVerifyingSigner\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPendingVerifyingSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBundlerAllowlist\",\"inputs\":[{\"name\":\"bundlers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePaymasterUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyingSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTo\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BundlerAllowlistUpdated\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingVerifyingSignerSet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationSponsored\",\"inputs\":[{\"name\":\"userOperationHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sponsorUUID\",\"type\":\"uint128\",\"indexed\":true,\"internalType\":\"uint128\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationSponsoredWithERC20\",\"inputs\":[{\"name\":\"userOperationHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sponsorUUID\",\"type\":\"uint128\",\"indexed\":true,\"internalType\":\"uint128\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyingSignerRotated\",\"inputs\":[{\"name\":\"oldSigner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BundlerNotAllowed\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DespositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPendingSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PostOpGasLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RenouceOwnershipDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderTokenBalanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxTokenCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// VerifyingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyingPaymasterMetaData.ABI instead.
var VerifyingPaymasterABI = VerifyingPaymasterMetaData.ABI

// VerifyingPaymaster is an auto generated Go binding around an Ethereum contract.
type VerifyingPaymaster struct {
	VerifyingPaymasterCaller     // Read-only binding to the contract
	VerifyingPaymasterTransactor // Write-only binding to the contract
	VerifyingPaymasterFilterer   // Log filterer for contract events
}

// VerifyingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyingPaymasterSession struct {
	Contract     *VerifyingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifyingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyingPaymasterCallerSession struct {
	Contract *VerifyingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifyingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyingPaymasterTransactorSession struct {
	Contract     *VerifyingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifyingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyingPaymasterRaw struct {
	Contract *VerifyingPaymaster // Generic contract binding to access the raw methods on
}

// VerifyingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCallerRaw struct {
	Contract *VerifyingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactorRaw struct {
	Contract *VerifyingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyingPaymaster creates a new instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymaster(address common.Address, backend bind.ContractBackend) (*VerifyingPaymaster, error) {
	contract, err := bindVerifyingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymaster{VerifyingPaymasterCaller: VerifyingPaymasterCaller{contract: contract}, VerifyingPaymasterTransactor: VerifyingPaymasterTransactor{contract: contract}, VerifyingPaymasterFilterer: VerifyingPaymasterFilterer{contract: contract}}, nil
}

// NewVerifyingPaymasterCaller creates a new read-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*VerifyingPaymasterCaller, error) {
	contract, err := bindVerifyingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterCaller{contract: contract}, nil
}

// NewVerifyingPaymasterTransactor creates a new write-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyingPaymasterTransactor, error) {
	contract, err := bindVerifyingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterTransactor{contract: contract}, nil
}

// NewVerifyingPaymasterFilterer creates a new log filterer instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyingPaymasterFilterer, error) {
	contract, err := bindVerifyingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterFilterer{contract: contract}, nil
}

// bindVerifyingPaymaster binds a generic wrapper to an already deployed contract.
func bindVerifyingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyingPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x38dcc981.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, (uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, uint256 paymasterValidationGasLimit, uint256 postOpGasLimit) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetHash(opts *bind.CallOpts, userOp PackedUserOperation, paymasterData VerifyingPaymasterPaymasterData, paymasterValidationGasLimit *big.Int, postOpGasLimit *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getHash", userOp, paymasterData, paymasterValidationGasLimit, postOpGasLimit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x38dcc981.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, (uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, uint256 paymasterValidationGasLimit, uint256 postOpGasLimit) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetHash(userOp PackedUserOperation, paymasterData VerifyingPaymasterPaymasterData, paymasterValidationGasLimit *big.Int, postOpGasLimit *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, paymasterData, paymasterValidationGasLimit, postOpGasLimit)
}

// GetHash is a free data retrieval call binding the contract method 0x38dcc981.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, (uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, uint256 paymasterValidationGasLimit, uint256 postOpGasLimit) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetHash(userOp PackedUserOperation, paymasterData VerifyingPaymasterPaymasterData, paymasterValidationGasLimit *big.Int, postOpGasLimit *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, paymasterData, paymasterValidationGasLimit, postOpGasLimit)
}

// IsBundlerAllowed is a free data retrieval call binding the contract method 0x4031c20e.
//
// Solidity: function isBundlerAllowed(address bundler) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) IsBundlerAllowed(opts *bind.CallOpts, bundler common.Address) (bool, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "isBundlerAllowed", bundler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBundlerAllowed is a free data retrieval call binding the contract method 0x4031c20e.
//
// Solidity: function isBundlerAllowed(address bundler) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterSession) IsBundlerAllowed(bundler common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.IsBundlerAllowed(&_VerifyingPaymaster.CallOpts, bundler)
}

// IsBundlerAllowed is a free data retrieval call binding the contract method 0x4031c20e.
//
// Solidity: function isBundlerAllowed(address bundler) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) IsBundlerAllowed(bundler common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.IsBundlerAllowed(&_VerifyingPaymaster.CallOpts, bundler)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// ParsePaymasterData is a free data retrieval call binding the contract method 0x81e4d81d.
//
// Solidity: function parsePaymasterData(bytes paymasterAndData) pure returns((uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) ParsePaymasterData(opts *bind.CallOpts, paymasterAndData []byte) (struct {
	PaymasterData VerifyingPaymasterPaymasterData
	Signature     []byte
}, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "parsePaymasterData", paymasterAndData)

	outstruct := new(struct {
		PaymasterData VerifyingPaymasterPaymasterData
		Signature     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymasterData = *abi.ConvertType(out[0], new(VerifyingPaymasterPaymasterData)).(*VerifyingPaymasterPaymasterData)
	outstruct.Signature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ParsePaymasterData is a free data retrieval call binding the contract method 0x81e4d81d.
//
// Solidity: function parsePaymasterData(bytes paymasterAndData) pure returns((uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ParsePaymasterData(paymasterAndData []byte) (struct {
	PaymasterData VerifyingPaymasterPaymasterData
	Signature     []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// ParsePaymasterData is a free data retrieval call binding the contract method 0x81e4d81d.
//
// Solidity: function parsePaymasterData(bytes paymasterAndData) pure returns((uint48,uint48,uint128,bool,bool,bool,address,address,uint256,uint48) paymasterData, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) ParsePaymasterData(paymasterAndData []byte) (struct {
	PaymasterData VerifyingPaymasterPaymasterData
	Signature     []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) PendingOwner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.PendingOwner(&_VerifyingPaymaster.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) PendingOwner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.PendingOwner(&_VerifyingPaymaster.CallOpts)
}

// PendingVerifyingSigner is a free data retrieval call binding the contract method 0xff1ff13a.
//
// Solidity: function pendingVerifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) PendingVerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "pendingVerifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVerifyingSigner is a free data retrieval call binding the contract method 0xff1ff13a.
//
// Solidity: function pendingVerifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) PendingVerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.PendingVerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// PendingVerifyingSigner is a free data retrieval call binding the contract method 0xff1ff13a.
//
// Solidity: function pendingVerifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) PendingVerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.PendingVerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_VerifyingPaymaster *VerifyingPaymasterCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) RenounceOwnership() error {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) RenounceOwnership() error {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) VerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "verifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AcceptOwnership(&_VerifyingPaymaster.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AcceptOwnership(&_VerifyingPaymaster.TransactOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// OwnerWithdrawERC20 is a paid mutator transaction binding the contract method 0x93563a95.
//
// Solidity: function ownerWithdrawERC20(address asset, address to, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) OwnerWithdrawERC20(opts *bind.TransactOpts, asset common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "ownerWithdrawERC20", asset, to, amount)
}

// OwnerWithdrawERC20 is a paid mutator transaction binding the contract method 0x93563a95.
//
// Solidity: function ownerWithdrawERC20(address asset, address to, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) OwnerWithdrawERC20(asset common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.OwnerWithdrawERC20(&_VerifyingPaymaster.TransactOpts, asset, to, amount)
}

// OwnerWithdrawERC20 is a paid mutator transaction binding the contract method 0x93563a95.
//
// Solidity: function ownerWithdrawERC20(address asset, address to, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) OwnerWithdrawERC20(asset common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.OwnerWithdrawERC20(&_VerifyingPaymaster.TransactOpts, asset, to, amount)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RotateVerifyingSigner is a paid mutator transaction binding the contract method 0x1f338ed8.
//
// Solidity: function rotateVerifyingSigner() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) RotateVerifyingSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "rotateVerifyingSigner")
}

// RotateVerifyingSigner is a paid mutator transaction binding the contract method 0x1f338ed8.
//
// Solidity: function rotateVerifyingSigner() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) RotateVerifyingSigner() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RotateVerifyingSigner(&_VerifyingPaymaster.TransactOpts)
}

// RotateVerifyingSigner is a paid mutator transaction binding the contract method 0x1f338ed8.
//
// Solidity: function rotateVerifyingSigner() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) RotateVerifyingSigner() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RotateVerifyingSigner(&_VerifyingPaymaster.TransactOpts)
}

// SetPendingVerifyingSigner is a paid mutator transaction binding the contract method 0x6fc45f4a.
//
// Solidity: function setPendingVerifyingSigner(address signer) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) SetPendingVerifyingSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "setPendingVerifyingSigner", signer)
}

// SetPendingVerifyingSigner is a paid mutator transaction binding the contract method 0x6fc45f4a.
//
// Solidity: function setPendingVerifyingSigner(address signer) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) SetPendingVerifyingSigner(signer common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetPendingVerifyingSigner(&_VerifyingPaymaster.TransactOpts, signer)
}

// SetPendingVerifyingSigner is a paid mutator transaction binding the contract method 0x6fc45f4a.
//
// Solidity: function setPendingVerifyingSigner(address signer) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) SetPendingVerifyingSigner(signer common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetPendingVerifyingSigner(&_VerifyingPaymaster.TransactOpts, signer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// UpdateBundlerAllowlist is a paid mutator transaction binding the contract method 0x7dd345cb.
//
// Solidity: function updateBundlerAllowlist(address[] bundlers, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) UpdateBundlerAllowlist(opts *bind.TransactOpts, bundlers []common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "updateBundlerAllowlist", bundlers, allowed)
}

// UpdateBundlerAllowlist is a paid mutator transaction binding the contract method 0x7dd345cb.
//
// Solidity: function updateBundlerAllowlist(address[] bundlers, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) UpdateBundlerAllowlist(bundlers []common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UpdateBundlerAllowlist(&_VerifyingPaymaster.TransactOpts, bundlers, allowed)
}

// UpdateBundlerAllowlist is a paid mutator transaction binding the contract method 0x7dd345cb.
//
// Solidity: function updateBundlerAllowlist(address[] bundlers, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) UpdateBundlerAllowlist(bundlers []common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UpdateBundlerAllowlist(&_VerifyingPaymaster.TransactOpts, bundlers, allowed)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Receive() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Receive(&_VerifyingPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Receive(&_VerifyingPaymaster.TransactOpts)
}

// VerifyingPaymasterBundlerAllowlistUpdatedIterator is returned from FilterBundlerAllowlistUpdated and is used to iterate over the raw logs and unpacked data for BundlerAllowlistUpdated events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterBundlerAllowlistUpdatedIterator struct {
	Event *VerifyingPaymasterBundlerAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterBundlerAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterBundlerAllowlistUpdated)
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
		it.Event = new(VerifyingPaymasterBundlerAllowlistUpdated)
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
func (it *VerifyingPaymasterBundlerAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterBundlerAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterBundlerAllowlistUpdated represents a BundlerAllowlistUpdated event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterBundlerAllowlistUpdated struct {
	Bundler common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBundlerAllowlistUpdated is a free log retrieval operation binding the contract event 0x8ff8c5211f68ef53b4bdd15ab2ea6d87be8a3dbf58865bd8325c984057e4fcb4.
//
// Solidity: event BundlerAllowlistUpdated(address bundler, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterBundlerAllowlistUpdated(opts *bind.FilterOpts) (*VerifyingPaymasterBundlerAllowlistUpdatedIterator, error) {

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "BundlerAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterBundlerAllowlistUpdatedIterator{contract: _VerifyingPaymaster.contract, event: "BundlerAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchBundlerAllowlistUpdated is a free log subscription operation binding the contract event 0x8ff8c5211f68ef53b4bdd15ab2ea6d87be8a3dbf58865bd8325c984057e4fcb4.
//
// Solidity: event BundlerAllowlistUpdated(address bundler, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchBundlerAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterBundlerAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "BundlerAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterBundlerAllowlistUpdated)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "BundlerAllowlistUpdated", log); err != nil {
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

// ParseBundlerAllowlistUpdated is a log parse operation binding the contract event 0x8ff8c5211f68ef53b4bdd15ab2ea6d87be8a3dbf58865bd8325c984057e4fcb4.
//
// Solidity: event BundlerAllowlistUpdated(address bundler, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseBundlerAllowlistUpdated(log types.Log) (*VerifyingPaymasterBundlerAllowlistUpdated, error) {
	event := new(VerifyingPaymasterBundlerAllowlistUpdated)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "BundlerAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferStartedIterator struct {
	Event *VerifyingPaymasterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterOwnershipTransferStarted)
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
		it.Event = new(VerifyingPaymasterOwnershipTransferStarted)
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
func (it *VerifyingPaymasterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyingPaymasterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterOwnershipTransferStartedIterator{contract: _VerifyingPaymaster.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterOwnershipTransferStarted)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseOwnershipTransferStarted(log types.Log) (*VerifyingPaymasterOwnershipTransferStarted, error) {
	event := new(VerifyingPaymasterOwnershipTransferStarted)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferredIterator struct {
	Event *VerifyingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
		it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterOwnershipTransferredIterator{contract: _VerifyingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterOwnershipTransferred)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*VerifyingPaymasterOwnershipTransferred, error) {
	event := new(VerifyingPaymasterOwnershipTransferred)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterPendingVerifyingSignerSetIterator is returned from FilterPendingVerifyingSignerSet and is used to iterate over the raw logs and unpacked data for PendingVerifyingSignerSet events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterPendingVerifyingSignerSetIterator struct {
	Event *VerifyingPaymasterPendingVerifyingSignerSet // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterPendingVerifyingSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterPendingVerifyingSignerSet)
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
		it.Event = new(VerifyingPaymasterPendingVerifyingSignerSet)
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
func (it *VerifyingPaymasterPendingVerifyingSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterPendingVerifyingSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterPendingVerifyingSignerSet represents a PendingVerifyingSignerSet event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterPendingVerifyingSignerSet struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPendingVerifyingSignerSet is a free log retrieval operation binding the contract event 0xe6846361b77f46dd08e26fd167e64b6166b0fe77d822b2178bd06cf629c5dd93.
//
// Solidity: event PendingVerifyingSignerSet(address signer)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterPendingVerifyingSignerSet(opts *bind.FilterOpts) (*VerifyingPaymasterPendingVerifyingSignerSetIterator, error) {

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "PendingVerifyingSignerSet")
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterPendingVerifyingSignerSetIterator{contract: _VerifyingPaymaster.contract, event: "PendingVerifyingSignerSet", logs: logs, sub: sub}, nil
}

// WatchPendingVerifyingSignerSet is a free log subscription operation binding the contract event 0xe6846361b77f46dd08e26fd167e64b6166b0fe77d822b2178bd06cf629c5dd93.
//
// Solidity: event PendingVerifyingSignerSet(address signer)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchPendingVerifyingSignerSet(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterPendingVerifyingSignerSet) (event.Subscription, error) {

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "PendingVerifyingSignerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterPendingVerifyingSignerSet)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "PendingVerifyingSignerSet", log); err != nil {
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

// ParsePendingVerifyingSignerSet is a log parse operation binding the contract event 0xe6846361b77f46dd08e26fd167e64b6166b0fe77d822b2178bd06cf629c5dd93.
//
// Solidity: event PendingVerifyingSignerSet(address signer)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParsePendingVerifyingSignerSet(log types.Log) (*VerifyingPaymasterPendingVerifyingSignerSet, error) {
	event := new(VerifyingPaymasterPendingVerifyingSignerSet)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "PendingVerifyingSignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterUserOperationSponsoredIterator is returned from FilterUserOperationSponsored and is used to iterate over the raw logs and unpacked data for UserOperationSponsored events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterUserOperationSponsoredIterator struct {
	Event *VerifyingPaymasterUserOperationSponsored // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterUserOperationSponsoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterUserOperationSponsored)
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
		it.Event = new(VerifyingPaymasterUserOperationSponsored)
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
func (it *VerifyingPaymasterUserOperationSponsoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterUserOperationSponsoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterUserOperationSponsored represents a UserOperationSponsored event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterUserOperationSponsored struct {
	UserOperationHash [32]byte
	SponsorUUID       *big.Int
	Token             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserOperationSponsored is a free log retrieval operation binding the contract event 0xfc6d279aa22eb9d6ba081f1bd2443d1e1c8eb3bcf089906d78064616798fbad0.
//
// Solidity: event UserOperationSponsored(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address token)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterUserOperationSponsored(opts *bind.FilterOpts, userOperationHash [][32]byte, sponsorUUID []*big.Int) (*VerifyingPaymasterUserOperationSponsoredIterator, error) {

	var userOperationHashRule []interface{}
	for _, userOperationHashItem := range userOperationHash {
		userOperationHashRule = append(userOperationHashRule, userOperationHashItem)
	}
	var sponsorUUIDRule []interface{}
	for _, sponsorUUIDItem := range sponsorUUID {
		sponsorUUIDRule = append(sponsorUUIDRule, sponsorUUIDItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "UserOperationSponsored", userOperationHashRule, sponsorUUIDRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterUserOperationSponsoredIterator{contract: _VerifyingPaymaster.contract, event: "UserOperationSponsored", logs: logs, sub: sub}, nil
}

// WatchUserOperationSponsored is a free log subscription operation binding the contract event 0xfc6d279aa22eb9d6ba081f1bd2443d1e1c8eb3bcf089906d78064616798fbad0.
//
// Solidity: event UserOperationSponsored(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address token)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchUserOperationSponsored(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterUserOperationSponsored, userOperationHash [][32]byte, sponsorUUID []*big.Int) (event.Subscription, error) {

	var userOperationHashRule []interface{}
	for _, userOperationHashItem := range userOperationHash {
		userOperationHashRule = append(userOperationHashRule, userOperationHashItem)
	}
	var sponsorUUIDRule []interface{}
	for _, sponsorUUIDItem := range sponsorUUID {
		sponsorUUIDRule = append(sponsorUUIDRule, sponsorUUIDItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "UserOperationSponsored", userOperationHashRule, sponsorUUIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterUserOperationSponsored)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "UserOperationSponsored", log); err != nil {
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

// ParseUserOperationSponsored is a log parse operation binding the contract event 0xfc6d279aa22eb9d6ba081f1bd2443d1e1c8eb3bcf089906d78064616798fbad0.
//
// Solidity: event UserOperationSponsored(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address token)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseUserOperationSponsored(log types.Log) (*VerifyingPaymasterUserOperationSponsored, error) {
	event := new(VerifyingPaymasterUserOperationSponsored)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "UserOperationSponsored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterUserOperationSponsoredWithERC20Iterator is returned from FilterUserOperationSponsoredWithERC20 and is used to iterate over the raw logs and unpacked data for UserOperationSponsoredWithERC20 events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterUserOperationSponsoredWithERC20Iterator struct {
	Event *VerifyingPaymasterUserOperationSponsoredWithERC20 // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterUserOperationSponsoredWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterUserOperationSponsoredWithERC20)
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
		it.Event = new(VerifyingPaymasterUserOperationSponsoredWithERC20)
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
func (it *VerifyingPaymasterUserOperationSponsoredWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterUserOperationSponsoredWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterUserOperationSponsoredWithERC20 represents a UserOperationSponsoredWithERC20 event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterUserOperationSponsoredWithERC20 struct {
	UserOperationHash [32]byte
	SponsorUUID       *big.Int
	Token             common.Address
	Receiver          common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserOperationSponsoredWithERC20 is a free log retrieval operation binding the contract event 0xe2459e4c2a3c2f358a092e76af1f3b2337c91875a1d49fe8afd6126342444f6f.
//
// Solidity: event UserOperationSponsoredWithERC20(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address indexed token, address receiver, uint256 amount)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterUserOperationSponsoredWithERC20(opts *bind.FilterOpts, userOperationHash [][32]byte, sponsorUUID []*big.Int, token []common.Address) (*VerifyingPaymasterUserOperationSponsoredWithERC20Iterator, error) {

	var userOperationHashRule []interface{}
	for _, userOperationHashItem := range userOperationHash {
		userOperationHashRule = append(userOperationHashRule, userOperationHashItem)
	}
	var sponsorUUIDRule []interface{}
	for _, sponsorUUIDItem := range sponsorUUID {
		sponsorUUIDRule = append(sponsorUUIDRule, sponsorUUIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "UserOperationSponsoredWithERC20", userOperationHashRule, sponsorUUIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterUserOperationSponsoredWithERC20Iterator{contract: _VerifyingPaymaster.contract, event: "UserOperationSponsoredWithERC20", logs: logs, sub: sub}, nil
}

// WatchUserOperationSponsoredWithERC20 is a free log subscription operation binding the contract event 0xe2459e4c2a3c2f358a092e76af1f3b2337c91875a1d49fe8afd6126342444f6f.
//
// Solidity: event UserOperationSponsoredWithERC20(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address indexed token, address receiver, uint256 amount)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchUserOperationSponsoredWithERC20(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterUserOperationSponsoredWithERC20, userOperationHash [][32]byte, sponsorUUID []*big.Int, token []common.Address) (event.Subscription, error) {

	var userOperationHashRule []interface{}
	for _, userOperationHashItem := range userOperationHash {
		userOperationHashRule = append(userOperationHashRule, userOperationHashItem)
	}
	var sponsorUUIDRule []interface{}
	for _, sponsorUUIDItem := range sponsorUUID {
		sponsorUUIDRule = append(sponsorUUIDRule, sponsorUUIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "UserOperationSponsoredWithERC20", userOperationHashRule, sponsorUUIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterUserOperationSponsoredWithERC20)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "UserOperationSponsoredWithERC20", log); err != nil {
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

// ParseUserOperationSponsoredWithERC20 is a log parse operation binding the contract event 0xe2459e4c2a3c2f358a092e76af1f3b2337c91875a1d49fe8afd6126342444f6f.
//
// Solidity: event UserOperationSponsoredWithERC20(bytes32 indexed userOperationHash, uint128 indexed sponsorUUID, address indexed token, address receiver, uint256 amount)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseUserOperationSponsoredWithERC20(log types.Log) (*VerifyingPaymasterUserOperationSponsoredWithERC20, error) {
	event := new(VerifyingPaymasterUserOperationSponsoredWithERC20)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "UserOperationSponsoredWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterVerifyingSignerRotatedIterator is returned from FilterVerifyingSignerRotated and is used to iterate over the raw logs and unpacked data for VerifyingSignerRotated events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterVerifyingSignerRotatedIterator struct {
	Event *VerifyingPaymasterVerifyingSignerRotated // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterVerifyingSignerRotatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterVerifyingSignerRotated)
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
		it.Event = new(VerifyingPaymasterVerifyingSignerRotated)
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
func (it *VerifyingPaymasterVerifyingSignerRotatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterVerifyingSignerRotatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterVerifyingSignerRotated represents a VerifyingSignerRotated event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterVerifyingSignerRotated struct {
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerifyingSignerRotated is a free log retrieval operation binding the contract event 0x25beca4c8409103108ee57a3f82258e57ef286914874cd10638d85c37d1427f7.
//
// Solidity: event VerifyingSignerRotated(address oldSigner, address newSigner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterVerifyingSignerRotated(opts *bind.FilterOpts) (*VerifyingPaymasterVerifyingSignerRotatedIterator, error) {

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "VerifyingSignerRotated")
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterVerifyingSignerRotatedIterator{contract: _VerifyingPaymaster.contract, event: "VerifyingSignerRotated", logs: logs, sub: sub}, nil
}

// WatchVerifyingSignerRotated is a free log subscription operation binding the contract event 0x25beca4c8409103108ee57a3f82258e57ef286914874cd10638d85c37d1427f7.
//
// Solidity: event VerifyingSignerRotated(address oldSigner, address newSigner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchVerifyingSignerRotated(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterVerifyingSignerRotated) (event.Subscription, error) {

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "VerifyingSignerRotated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterVerifyingSignerRotated)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "VerifyingSignerRotated", log); err != nil {
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

// ParseVerifyingSignerRotated is a log parse operation binding the contract event 0x25beca4c8409103108ee57a3f82258e57ef286914874cd10638d85c37d1427f7.
//
// Solidity: event VerifyingSignerRotated(address oldSigner, address newSigner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseVerifyingSignerRotated(log types.Log) (*VerifyingPaymasterVerifyingSignerRotated, error) {
	event := new(VerifyingPaymasterVerifyingSignerRotated)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "VerifyingSignerRotated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
