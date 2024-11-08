// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transaction

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

// FinancialTransactionTransaction is an auto generated low-level Go binding around an user-defined struct.
type FinancialTransactionTransaction struct {
	TransactionId   string
	SenderAccount   string
	ReceiverAccount string
	Amount          *big.Int
	Currency        uint8
	Timestamp       *big.Int
	BlockTimestamp  *big.Int
	Note            string
}

// FinancialTransactionMetaData contains all meta data concerning the FinancialTransaction contract.
var FinancialTransactionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"name\":\"TransactionNoteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"senderAccount\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiverAccount\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumFinancialTransaction.Currency\",\"name\":\"currency\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"name\":\"TransactionRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserUnauthorized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"authorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"unauthorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_transactionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_senderAccount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_receiverAccount\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"enumFinancialTransaction.Currency\",\"name\":\"_currency\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_note\",\"type\":\"string\"}],\"name\":\"recordTransaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_transactionHash\",\"type\":\"bytes32\"}],\"name\":\"getTransactionByHash\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"senderAccount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiverAccount\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumFinancialTransaction.Currency\",\"name\":\"currency\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"internalType\":\"structFinancialTransaction.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"}],\"name\":\"getAccountTransactions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumFinancialTransaction.Currency\",\"name\":\"_currency\",\"type\":\"uint8\"}],\"name\":\"getCurrencyName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_note\",\"type\":\"string\"}],\"name\":\"updateTransactionNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FinancialTransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use FinancialTransactionMetaData.ABI instead.
var FinancialTransactionABI = FinancialTransactionMetaData.ABI

// FinancialTransaction is an auto generated Go binding around an Ethereum contract.
type FinancialTransaction struct {
	FinancialTransactionCaller     // Read-only binding to the contract
	FinancialTransactionTransactor // Write-only binding to the contract
	FinancialTransactionFilterer   // Log filterer for contract events
}

// FinancialTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinancialTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinancialTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinancialTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinancialTransactionSession struct {
	Contract     *FinancialTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FinancialTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinancialTransactionCallerSession struct {
	Contract *FinancialTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FinancialTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinancialTransactionTransactorSession struct {
	Contract     *FinancialTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FinancialTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinancialTransactionRaw struct {
	Contract *FinancialTransaction // Generic contract binding to access the raw methods on
}

// FinancialTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinancialTransactionCallerRaw struct {
	Contract *FinancialTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// FinancialTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinancialTransactionTransactorRaw struct {
	Contract *FinancialTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinancialTransaction creates a new instance of FinancialTransaction, bound to a specific deployed contract.
func NewFinancialTransaction(address common.Address, backend bind.ContractBackend) (*FinancialTransaction, error) {
	contract, err := bindFinancialTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinancialTransaction{FinancialTransactionCaller: FinancialTransactionCaller{contract: contract}, FinancialTransactionTransactor: FinancialTransactionTransactor{contract: contract}, FinancialTransactionFilterer: FinancialTransactionFilterer{contract: contract}}, nil
}

// NewFinancialTransactionCaller creates a new read-only instance of FinancialTransaction, bound to a specific deployed contract.
func NewFinancialTransactionCaller(address common.Address, caller bind.ContractCaller) (*FinancialTransactionCaller, error) {
	contract, err := bindFinancialTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionCaller{contract: contract}, nil
}

// NewFinancialTransactionTransactor creates a new write-only instance of FinancialTransaction, bound to a specific deployed contract.
func NewFinancialTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*FinancialTransactionTransactor, error) {
	contract, err := bindFinancialTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionTransactor{contract: contract}, nil
}

// NewFinancialTransactionFilterer creates a new log filterer instance of FinancialTransaction, bound to a specific deployed contract.
func NewFinancialTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*FinancialTransactionFilterer, error) {
	contract, err := bindFinancialTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionFilterer{contract: contract}, nil
}

// bindFinancialTransaction binds a generic wrapper to an already deployed contract.
func bindFinancialTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinancialTransactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinancialTransaction *FinancialTransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinancialTransaction.Contract.FinancialTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinancialTransaction *FinancialTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.FinancialTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinancialTransaction *FinancialTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.FinancialTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinancialTransaction *FinancialTransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinancialTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinancialTransaction *FinancialTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinancialTransaction *FinancialTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.contract.Transact(opts, method, params...)
}

// GetAccountTransactions is a free data retrieval call binding the contract method 0x9cae6559.
//
// Solidity: function getAccountTransactions(string _account) view returns(bytes32[])
func (_FinancialTransaction *FinancialTransactionCaller) GetAccountTransactions(opts *bind.CallOpts, _account string) ([][32]byte, error) {
	var out []interface{}
	err := _FinancialTransaction.contract.Call(opts, &out, "getAccountTransactions", _account)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAccountTransactions is a free data retrieval call binding the contract method 0x9cae6559.
//
// Solidity: function getAccountTransactions(string _account) view returns(bytes32[])
func (_FinancialTransaction *FinancialTransactionSession) GetAccountTransactions(_account string) ([][32]byte, error) {
	return _FinancialTransaction.Contract.GetAccountTransactions(&_FinancialTransaction.CallOpts, _account)
}

// GetAccountTransactions is a free data retrieval call binding the contract method 0x9cae6559.
//
// Solidity: function getAccountTransactions(string _account) view returns(bytes32[])
func (_FinancialTransaction *FinancialTransactionCallerSession) GetAccountTransactions(_account string) ([][32]byte, error) {
	return _FinancialTransaction.Contract.GetAccountTransactions(&_FinancialTransaction.CallOpts, _account)
}

// GetCurrencyName is a free data retrieval call binding the contract method 0x8c4f2f11.
//
// Solidity: function getCurrencyName(uint8 _currency) pure returns(string)
func (_FinancialTransaction *FinancialTransactionCaller) GetCurrencyName(opts *bind.CallOpts, _currency uint8) (string, error) {
	var out []interface{}
	err := _FinancialTransaction.contract.Call(opts, &out, "getCurrencyName", _currency)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCurrencyName is a free data retrieval call binding the contract method 0x8c4f2f11.
//
// Solidity: function getCurrencyName(uint8 _currency) pure returns(string)
func (_FinancialTransaction *FinancialTransactionSession) GetCurrencyName(_currency uint8) (string, error) {
	return _FinancialTransaction.Contract.GetCurrencyName(&_FinancialTransaction.CallOpts, _currency)
}

// GetCurrencyName is a free data retrieval call binding the contract method 0x8c4f2f11.
//
// Solidity: function getCurrencyName(uint8 _currency) pure returns(string)
func (_FinancialTransaction *FinancialTransactionCallerSession) GetCurrencyName(_currency uint8) (string, error) {
	return _FinancialTransaction.Contract.GetCurrencyName(&_FinancialTransaction.CallOpts, _currency)
}

// GetTransactionByHash is a free data retrieval call binding the contract method 0xee3d7ad3.
//
// Solidity: function getTransactionByHash(bytes32 _transactionHash) view returns((string,string,string,uint256,uint8,uint256,uint256,string))
func (_FinancialTransaction *FinancialTransactionCaller) GetTransactionByHash(opts *bind.CallOpts, _transactionHash [32]byte) (FinancialTransactionTransaction, error) {
	var out []interface{}
	err := _FinancialTransaction.contract.Call(opts, &out, "getTransactionByHash", _transactionHash)

	if err != nil {
		return *new(FinancialTransactionTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(FinancialTransactionTransaction)).(*FinancialTransactionTransaction)

	return out0, err

}

// GetTransactionByHash is a free data retrieval call binding the contract method 0xee3d7ad3.
//
// Solidity: function getTransactionByHash(bytes32 _transactionHash) view returns((string,string,string,uint256,uint8,uint256,uint256,string))
func (_FinancialTransaction *FinancialTransactionSession) GetTransactionByHash(_transactionHash [32]byte) (FinancialTransactionTransaction, error) {
	return _FinancialTransaction.Contract.GetTransactionByHash(&_FinancialTransaction.CallOpts, _transactionHash)
}

// GetTransactionByHash is a free data retrieval call binding the contract method 0xee3d7ad3.
//
// Solidity: function getTransactionByHash(bytes32 _transactionHash) view returns((string,string,string,uint256,uint8,uint256,uint256,string))
func (_FinancialTransaction *FinancialTransactionCallerSession) GetTransactionByHash(_transactionHash [32]byte) (FinancialTransactionTransaction, error) {
	return _FinancialTransaction.Contract.GetTransactionByHash(&_FinancialTransaction.CallOpts, _transactionHash)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_FinancialTransaction *FinancialTransactionCaller) IsAuthorized(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _FinancialTransaction.contract.Call(opts, &out, "isAuthorized", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_FinancialTransaction *FinancialTransactionSession) IsAuthorized(user common.Address) (bool, error) {
	return _FinancialTransaction.Contract.IsAuthorized(&_FinancialTransaction.CallOpts, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_FinancialTransaction *FinancialTransactionCallerSession) IsAuthorized(user common.Address) (bool, error) {
	return _FinancialTransaction.Contract.IsAuthorized(&_FinancialTransaction.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinancialTransaction *FinancialTransactionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinancialTransaction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinancialTransaction *FinancialTransactionSession) Owner() (common.Address, error) {
	return _FinancialTransaction.Contract.Owner(&_FinancialTransaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinancialTransaction *FinancialTransactionCallerSession) Owner() (common.Address, error) {
	return _FinancialTransaction.Contract.Owner(&_FinancialTransaction.CallOpts)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionTransactor) AuthorizeUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "authorizeUser", user)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionSession) AuthorizeUser(user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.AuthorizeUser(&_FinancialTransaction.TransactOpts, user)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionTransactorSession) AuthorizeUser(user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.AuthorizeUser(&_FinancialTransaction.TransactOpts, user)
}

// RecordTransaction is a paid mutator transaction binding the contract method 0x53a9e3db.
//
// Solidity: function recordTransaction(string _transactionId, string _senderAccount, string _receiverAccount, uint256 _amount, uint8 _currency, uint256 _timestamp, string _note) returns(bytes32)
func (_FinancialTransaction *FinancialTransactionTransactor) RecordTransaction(opts *bind.TransactOpts, _transactionId string, _senderAccount string, _receiverAccount string, _amount *big.Int, _currency uint8, _timestamp *big.Int, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "recordTransaction", _transactionId, _senderAccount, _receiverAccount, _amount, _currency, _timestamp, _note)
}

// RecordTransaction is a paid mutator transaction binding the contract method 0x53a9e3db.
//
// Solidity: function recordTransaction(string _transactionId, string _senderAccount, string _receiverAccount, uint256 _amount, uint8 _currency, uint256 _timestamp, string _note) returns(bytes32)
func (_FinancialTransaction *FinancialTransactionSession) RecordTransaction(_transactionId string, _senderAccount string, _receiverAccount string, _amount *big.Int, _currency uint8, _timestamp *big.Int, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.RecordTransaction(&_FinancialTransaction.TransactOpts, _transactionId, _senderAccount, _receiverAccount, _amount, _currency, _timestamp, _note)
}

// RecordTransaction is a paid mutator transaction binding the contract method 0x53a9e3db.
//
// Solidity: function recordTransaction(string _transactionId, string _senderAccount, string _receiverAccount, uint256 _amount, uint8 _currency, uint256 _timestamp, string _note) returns(bytes32)
func (_FinancialTransaction *FinancialTransactionTransactorSession) RecordTransaction(_transactionId string, _senderAccount string, _receiverAccount string, _amount *big.Int, _currency uint8, _timestamp *big.Int, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.RecordTransaction(&_FinancialTransaction.TransactOpts, _transactionId, _senderAccount, _receiverAccount, _amount, _currency, _timestamp, _note)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinancialTransaction *FinancialTransactionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinancialTransaction *FinancialTransactionSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinancialTransaction.Contract.RenounceOwnership(&_FinancialTransaction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinancialTransaction *FinancialTransactionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinancialTransaction.Contract.RenounceOwnership(&_FinancialTransaction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinancialTransaction *FinancialTransactionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinancialTransaction *FinancialTransactionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.TransferOwnership(&_FinancialTransaction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinancialTransaction *FinancialTransactionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.TransferOwnership(&_FinancialTransaction.TransactOpts, newOwner)
}

// UnauthorizeUser is a paid mutator transaction binding the contract method 0x478aa69e.
//
// Solidity: function unauthorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionTransactor) UnauthorizeUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "unauthorizeUser", user)
}

// UnauthorizeUser is a paid mutator transaction binding the contract method 0x478aa69e.
//
// Solidity: function unauthorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionSession) UnauthorizeUser(user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.UnauthorizeUser(&_FinancialTransaction.TransactOpts, user)
}

// UnauthorizeUser is a paid mutator transaction binding the contract method 0x478aa69e.
//
// Solidity: function unauthorizeUser(address user) returns()
func (_FinancialTransaction *FinancialTransactionTransactorSession) UnauthorizeUser(user common.Address) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.UnauthorizeUser(&_FinancialTransaction.TransactOpts, user)
}

// UpdateTransactionNote is a paid mutator transaction binding the contract method 0xd3e1d1fc.
//
// Solidity: function updateTransactionNote(bytes32 _transactionHash, string _note) returns()
func (_FinancialTransaction *FinancialTransactionTransactor) UpdateTransactionNote(opts *bind.TransactOpts, _transactionHash [32]byte, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.contract.Transact(opts, "updateTransactionNote", _transactionHash, _note)
}

// UpdateTransactionNote is a paid mutator transaction binding the contract method 0xd3e1d1fc.
//
// Solidity: function updateTransactionNote(bytes32 _transactionHash, string _note) returns()
func (_FinancialTransaction *FinancialTransactionSession) UpdateTransactionNote(_transactionHash [32]byte, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.UpdateTransactionNote(&_FinancialTransaction.TransactOpts, _transactionHash, _note)
}

// UpdateTransactionNote is a paid mutator transaction binding the contract method 0xd3e1d1fc.
//
// Solidity: function updateTransactionNote(bytes32 _transactionHash, string _note) returns()
func (_FinancialTransaction *FinancialTransactionTransactorSession) UpdateTransactionNote(_transactionHash [32]byte, _note string) (*types.Transaction, error) {
	return _FinancialTransaction.Contract.UpdateTransactionNote(&_FinancialTransaction.TransactOpts, _transactionHash, _note)
}

// FinancialTransactionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinancialTransaction contract.
type FinancialTransactionOwnershipTransferredIterator struct {
	Event *FinancialTransactionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinancialTransactionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialTransactionOwnershipTransferred)
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
		it.Event = new(FinancialTransactionOwnershipTransferred)
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
func (it *FinancialTransactionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialTransactionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialTransactionOwnershipTransferred represents a OwnershipTransferred event raised by the FinancialTransaction contract.
type FinancialTransactionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinancialTransaction *FinancialTransactionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinancialTransactionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinancialTransaction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionOwnershipTransferredIterator{contract: _FinancialTransaction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinancialTransaction *FinancialTransactionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinancialTransactionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinancialTransaction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialTransactionOwnershipTransferred)
				if err := _FinancialTransaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FinancialTransaction *FinancialTransactionFilterer) ParseOwnershipTransferred(log types.Log) (*FinancialTransactionOwnershipTransferred, error) {
	event := new(FinancialTransactionOwnershipTransferred)
	if err := _FinancialTransaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialTransactionTransactionNoteUpdatedIterator is returned from FilterTransactionNoteUpdated and is used to iterate over the raw logs and unpacked data for TransactionNoteUpdated events raised by the FinancialTransaction contract.
type FinancialTransactionTransactionNoteUpdatedIterator struct {
	Event *FinancialTransactionTransactionNoteUpdated // Event containing the contract specifics and raw log

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
func (it *FinancialTransactionTransactionNoteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialTransactionTransactionNoteUpdated)
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
		it.Event = new(FinancialTransactionTransactionNoteUpdated)
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
func (it *FinancialTransactionTransactionNoteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialTransactionTransactionNoteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialTransactionTransactionNoteUpdated represents a TransactionNoteUpdated event raised by the FinancialTransaction contract.
type FinancialTransactionTransactionNoteUpdated struct {
	Hash [32]byte
	Note string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransactionNoteUpdated is a free log retrieval operation binding the contract event 0xa009e6b329e30e84a7362ac8291febe4eae3b9333a71a690d24df6c2727a430c.
//
// Solidity: event TransactionNoteUpdated(bytes32 indexed hash, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) FilterTransactionNoteUpdated(opts *bind.FilterOpts, hash [][32]byte) (*FinancialTransactionTransactionNoteUpdatedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _FinancialTransaction.contract.FilterLogs(opts, "TransactionNoteUpdated", hashRule)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionTransactionNoteUpdatedIterator{contract: _FinancialTransaction.contract, event: "TransactionNoteUpdated", logs: logs, sub: sub}, nil
}

// WatchTransactionNoteUpdated is a free log subscription operation binding the contract event 0xa009e6b329e30e84a7362ac8291febe4eae3b9333a71a690d24df6c2727a430c.
//
// Solidity: event TransactionNoteUpdated(bytes32 indexed hash, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) WatchTransactionNoteUpdated(opts *bind.WatchOpts, sink chan<- *FinancialTransactionTransactionNoteUpdated, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _FinancialTransaction.contract.WatchLogs(opts, "TransactionNoteUpdated", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialTransactionTransactionNoteUpdated)
				if err := _FinancialTransaction.contract.UnpackLog(event, "TransactionNoteUpdated", log); err != nil {
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

// ParseTransactionNoteUpdated is a log parse operation binding the contract event 0xa009e6b329e30e84a7362ac8291febe4eae3b9333a71a690d24df6c2727a430c.
//
// Solidity: event TransactionNoteUpdated(bytes32 indexed hash, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) ParseTransactionNoteUpdated(log types.Log) (*FinancialTransactionTransactionNoteUpdated, error) {
	event := new(FinancialTransactionTransactionNoteUpdated)
	if err := _FinancialTransaction.contract.UnpackLog(event, "TransactionNoteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialTransactionTransactionRecordedIterator is returned from FilterTransactionRecorded and is used to iterate over the raw logs and unpacked data for TransactionRecorded events raised by the FinancialTransaction contract.
type FinancialTransactionTransactionRecordedIterator struct {
	Event *FinancialTransactionTransactionRecorded // Event containing the contract specifics and raw log

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
func (it *FinancialTransactionTransactionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialTransactionTransactionRecorded)
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
		it.Event = new(FinancialTransactionTransactionRecorded)
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
func (it *FinancialTransactionTransactionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialTransactionTransactionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialTransactionTransactionRecorded represents a TransactionRecorded event raised by the FinancialTransaction contract.
type FinancialTransactionTransactionRecorded struct {
	Hash            [32]byte
	TransactionId   string
	SenderAccount   string
	ReceiverAccount string
	Amount          *big.Int
	Currency        uint8
	Timestamp       *big.Int
	BlockTimestamp  *big.Int
	Note            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionRecorded is a free log retrieval operation binding the contract event 0xb607de420036a4d0aeca386fc2185b7a865baebac29240ab811c596059cf1507.
//
// Solidity: event TransactionRecorded(bytes32 indexed hash, string transactionId, string senderAccount, string receiverAccount, uint256 amount, uint8 currency, uint256 timestamp, uint256 blockTimestamp, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) FilterTransactionRecorded(opts *bind.FilterOpts, hash [][32]byte) (*FinancialTransactionTransactionRecordedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _FinancialTransaction.contract.FilterLogs(opts, "TransactionRecorded", hashRule)
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionTransactionRecordedIterator{contract: _FinancialTransaction.contract, event: "TransactionRecorded", logs: logs, sub: sub}, nil
}

// WatchTransactionRecorded is a free log subscription operation binding the contract event 0xb607de420036a4d0aeca386fc2185b7a865baebac29240ab811c596059cf1507.
//
// Solidity: event TransactionRecorded(bytes32 indexed hash, string transactionId, string senderAccount, string receiverAccount, uint256 amount, uint8 currency, uint256 timestamp, uint256 blockTimestamp, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) WatchTransactionRecorded(opts *bind.WatchOpts, sink chan<- *FinancialTransactionTransactionRecorded, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _FinancialTransaction.contract.WatchLogs(opts, "TransactionRecorded", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialTransactionTransactionRecorded)
				if err := _FinancialTransaction.contract.UnpackLog(event, "TransactionRecorded", log); err != nil {
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

// ParseTransactionRecorded is a log parse operation binding the contract event 0xb607de420036a4d0aeca386fc2185b7a865baebac29240ab811c596059cf1507.
//
// Solidity: event TransactionRecorded(bytes32 indexed hash, string transactionId, string senderAccount, string receiverAccount, uint256 amount, uint8 currency, uint256 timestamp, uint256 blockTimestamp, string note)
func (_FinancialTransaction *FinancialTransactionFilterer) ParseTransactionRecorded(log types.Log) (*FinancialTransactionTransactionRecorded, error) {
	event := new(FinancialTransactionTransactionRecorded)
	if err := _FinancialTransaction.contract.UnpackLog(event, "TransactionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialTransactionUserAuthorizedIterator is returned from FilterUserAuthorized and is used to iterate over the raw logs and unpacked data for UserAuthorized events raised by the FinancialTransaction contract.
type FinancialTransactionUserAuthorizedIterator struct {
	Event *FinancialTransactionUserAuthorized // Event containing the contract specifics and raw log

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
func (it *FinancialTransactionUserAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialTransactionUserAuthorized)
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
		it.Event = new(FinancialTransactionUserAuthorized)
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
func (it *FinancialTransactionUserAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialTransactionUserAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialTransactionUserAuthorized represents a UserAuthorized event raised by the FinancialTransaction contract.
type FinancialTransactionUserAuthorized struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserAuthorized is a free log retrieval operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) FilterUserAuthorized(opts *bind.FilterOpts) (*FinancialTransactionUserAuthorizedIterator, error) {

	logs, sub, err := _FinancialTransaction.contract.FilterLogs(opts, "UserAuthorized")
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionUserAuthorizedIterator{contract: _FinancialTransaction.contract, event: "UserAuthorized", logs: logs, sub: sub}, nil
}

// WatchUserAuthorized is a free log subscription operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) WatchUserAuthorized(opts *bind.WatchOpts, sink chan<- *FinancialTransactionUserAuthorized) (event.Subscription, error) {

	logs, sub, err := _FinancialTransaction.contract.WatchLogs(opts, "UserAuthorized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialTransactionUserAuthorized)
				if err := _FinancialTransaction.contract.UnpackLog(event, "UserAuthorized", log); err != nil {
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

// ParseUserAuthorized is a log parse operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) ParseUserAuthorized(log types.Log) (*FinancialTransactionUserAuthorized, error) {
	event := new(FinancialTransactionUserAuthorized)
	if err := _FinancialTransaction.contract.UnpackLog(event, "UserAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialTransactionUserUnauthorizedIterator is returned from FilterUserUnauthorized and is used to iterate over the raw logs and unpacked data for UserUnauthorized events raised by the FinancialTransaction contract.
type FinancialTransactionUserUnauthorizedIterator struct {
	Event *FinancialTransactionUserUnauthorized // Event containing the contract specifics and raw log

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
func (it *FinancialTransactionUserUnauthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialTransactionUserUnauthorized)
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
		it.Event = new(FinancialTransactionUserUnauthorized)
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
func (it *FinancialTransactionUserUnauthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialTransactionUserUnauthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialTransactionUserUnauthorized represents a UserUnauthorized event raised by the FinancialTransaction contract.
type FinancialTransactionUserUnauthorized struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserUnauthorized is a free log retrieval operation binding the contract event 0x39814f442b71862bed3e83dececad92d06dba81a8dbbd9459237ae953efd4a92.
//
// Solidity: event UserUnauthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) FilterUserUnauthorized(opts *bind.FilterOpts) (*FinancialTransactionUserUnauthorizedIterator, error) {

	logs, sub, err := _FinancialTransaction.contract.FilterLogs(opts, "UserUnauthorized")
	if err != nil {
		return nil, err
	}
	return &FinancialTransactionUserUnauthorizedIterator{contract: _FinancialTransaction.contract, event: "UserUnauthorized", logs: logs, sub: sub}, nil
}

// WatchUserUnauthorized is a free log subscription operation binding the contract event 0x39814f442b71862bed3e83dececad92d06dba81a8dbbd9459237ae953efd4a92.
//
// Solidity: event UserUnauthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) WatchUserUnauthorized(opts *bind.WatchOpts, sink chan<- *FinancialTransactionUserUnauthorized) (event.Subscription, error) {

	logs, sub, err := _FinancialTransaction.contract.WatchLogs(opts, "UserUnauthorized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialTransactionUserUnauthorized)
				if err := _FinancialTransaction.contract.UnpackLog(event, "UserUnauthorized", log); err != nil {
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

// ParseUserUnauthorized is a log parse operation binding the contract event 0x39814f442b71862bed3e83dececad92d06dba81a8dbbd9459237ae953efd4a92.
//
// Solidity: event UserUnauthorized(address user)
func (_FinancialTransaction *FinancialTransactionFilterer) ParseUserUnauthorized(log types.Log) (*FinancialTransactionUserUnauthorized, error) {
	event := new(FinancialTransactionUserUnauthorized)
	if err := _FinancialTransaction.contract.UnpackLog(event, "UserUnauthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
