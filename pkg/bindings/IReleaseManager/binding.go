// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IReleaseManager

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

// IReleaseManagerPublishedRelease is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerPublishedRelease struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}

// IReleaseManagerMetaData contains all meta data concerning the IReleaseManager contract.
var IReleaseManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deprecateRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeprecatedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublishedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManager.PublishedRelease[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseAtBlock\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseCheckpointCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSDeregistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AVSAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AVSNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyDeprecated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDigest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// IReleaseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IReleaseManagerMetaData.ABI instead.
var IReleaseManagerABI = IReleaseManagerMetaData.ABI

// IReleaseManager is an auto generated Go binding around an Ethereum contract.
type IReleaseManager struct {
	IReleaseManagerCaller     // Read-only binding to the contract
	IReleaseManagerTransactor // Write-only binding to the contract
	IReleaseManagerFilterer   // Log filterer for contract events
}

// IReleaseManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReleaseManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReleaseManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReleaseManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReleaseManagerSession struct {
	Contract     *IReleaseManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReleaseManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReleaseManagerCallerSession struct {
	Contract *IReleaseManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IReleaseManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReleaseManagerTransactorSession struct {
	Contract     *IReleaseManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IReleaseManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReleaseManagerRaw struct {
	Contract *IReleaseManager // Generic contract binding to access the raw methods on
}

// IReleaseManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReleaseManagerCallerRaw struct {
	Contract *IReleaseManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IReleaseManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReleaseManagerTransactorRaw struct {
	Contract *IReleaseManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReleaseManager creates a new instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManager(address common.Address, backend bind.ContractBackend) (*IReleaseManager, error) {
	contract, err := bindIReleaseManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReleaseManager{IReleaseManagerCaller: IReleaseManagerCaller{contract: contract}, IReleaseManagerTransactor: IReleaseManagerTransactor{contract: contract}, IReleaseManagerFilterer: IReleaseManagerFilterer{contract: contract}}, nil
}

// NewIReleaseManagerCaller creates a new read-only instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerCaller(address common.Address, caller bind.ContractCaller) (*IReleaseManagerCaller, error) {
	contract, err := bindIReleaseManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerCaller{contract: contract}, nil
}

// NewIReleaseManagerTransactor creates a new write-only instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IReleaseManagerTransactor, error) {
	contract, err := bindIReleaseManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerTransactor{contract: contract}, nil
}

// NewIReleaseManagerFilterer creates a new log filterer instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IReleaseManagerFilterer, error) {
	contract, err := bindIReleaseManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerFilterer{contract: contract}, nil
}

// bindIReleaseManager binds a generic wrapper to an already deployed contract.
func bindIReleaseManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReleaseManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReleaseManager *IReleaseManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReleaseManager.Contract.IReleaseManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReleaseManager *IReleaseManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReleaseManager.Contract.IReleaseManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReleaseManager *IReleaseManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReleaseManager.Contract.IReleaseManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReleaseManager *IReleaseManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReleaseManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReleaseManager *IReleaseManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReleaseManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReleaseManager *IReleaseManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReleaseManager.Contract.contract.Transact(opts, method, params...)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_IReleaseManager *IReleaseManagerCaller) GetDeprecatedReleases(opts *bind.CallOpts, avs common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getDeprecatedReleases", avs)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_IReleaseManager *IReleaseManagerSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _IReleaseManager.Contract.GetDeprecatedReleases(&_IReleaseManager.CallOpts, avs)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_IReleaseManager *IReleaseManagerCallerSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _IReleaseManager.Contract.GetDeprecatedReleases(&_IReleaseManager.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerCaller) GetLatestRelease(opts *bind.CallOpts, avs common.Address) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getLatestRelease", avs)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetLatestRelease(&_IReleaseManager.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerCallerSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetLatestRelease(&_IReleaseManager.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_IReleaseManager *IReleaseManagerCaller) GetPublishedReleases(opts *bind.CallOpts, avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getPublishedReleases", avs)

	if err != nil {
		return *new([]IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new([]IReleaseManagerPublishedRelease)).(*[]IReleaseManagerPublishedRelease)

	return out0, err

}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_IReleaseManager *IReleaseManagerSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetPublishedReleases(&_IReleaseManager.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_IReleaseManager *IReleaseManagerCallerSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetPublishedReleases(&_IReleaseManager.CallOpts, avs)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerCaller) GetReleaseAtBlock(opts *bind.CallOpts, avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getReleaseAtBlock", avs, blockNumber)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetReleaseAtBlock(&_IReleaseManager.CallOpts, avs, blockNumber)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_IReleaseManager *IReleaseManagerCallerSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _IReleaseManager.Contract.GetReleaseAtBlock(&_IReleaseManager.CallOpts, avs, blockNumber)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_IReleaseManager *IReleaseManagerCaller) GetReleaseCheckpointCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getReleaseCheckpointCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_IReleaseManager *IReleaseManagerSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _IReleaseManager.Contract.GetReleaseCheckpointCount(&_IReleaseManager.CallOpts, avs)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_IReleaseManager *IReleaseManagerCallerSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _IReleaseManager.Contract.GetReleaseCheckpointCount(&_IReleaseManager.CallOpts, avs)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_IReleaseManager *IReleaseManagerCaller) IsReleaseDeprecated(opts *bind.CallOpts, avs common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "isReleaseDeprecated", avs, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_IReleaseManager *IReleaseManagerSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _IReleaseManager.Contract.IsReleaseDeprecated(&_IReleaseManager.CallOpts, avs, digest)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_IReleaseManager *IReleaseManagerCallerSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _IReleaseManager.Contract.IsReleaseDeprecated(&_IReleaseManager.CallOpts, avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_IReleaseManager *IReleaseManagerTransactor) DeprecateRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "deprecateRelease", avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_IReleaseManager *IReleaseManagerSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _IReleaseManager.Contract.DeprecateRelease(&_IReleaseManager.TransactOpts, avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_IReleaseManager *IReleaseManagerTransactorSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _IReleaseManager.Contract.DeprecateRelease(&_IReleaseManager.TransactOpts, avs, digest)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_IReleaseManager *IReleaseManagerTransactor) Deregister(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "deregister", avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_IReleaseManager *IReleaseManagerSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.Contract.Deregister(&_IReleaseManager.TransactOpts, avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_IReleaseManager *IReleaseManagerTransactorSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.Contract.Deregister(&_IReleaseManager.TransactOpts, avs)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_IReleaseManager *IReleaseManagerTransactor) PublishRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "publishRelease", avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_IReleaseManager *IReleaseManagerSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishRelease(&_IReleaseManager.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_IReleaseManager *IReleaseManagerTransactorSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishRelease(&_IReleaseManager.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_IReleaseManager *IReleaseManagerTransactor) Register(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "register", avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_IReleaseManager *IReleaseManagerSession) Register(avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.Contract.Register(&_IReleaseManager.TransactOpts, avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_IReleaseManager *IReleaseManagerTransactorSession) Register(avs common.Address) (*types.Transaction, error) {
	return _IReleaseManager.Contract.Register(&_IReleaseManager.TransactOpts, avs)
}

// IReleaseManagerAVSDeregisteredIterator is returned from FilterAVSDeregistered and is used to iterate over the raw logs and unpacked data for AVSDeregistered events raised by the IReleaseManager contract.
type IReleaseManagerAVSDeregisteredIterator struct {
	Event *IReleaseManagerAVSDeregistered // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerAVSDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerAVSDeregistered)
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
		it.Event = new(IReleaseManagerAVSDeregistered)
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
func (it *IReleaseManagerAVSDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerAVSDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerAVSDeregistered represents a AVSDeregistered event raised by the IReleaseManager contract.
type IReleaseManagerAVSDeregistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSDeregistered is a free log retrieval operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) FilterAVSDeregistered(opts *bind.FilterOpts, avs []common.Address) (*IReleaseManagerAVSDeregisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerAVSDeregisteredIterator{contract: _IReleaseManager.contract, event: "AVSDeregistered", logs: logs, sub: sub}, nil
}

// WatchAVSDeregistered is a free log subscription operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) WatchAVSDeregistered(opts *bind.WatchOpts, sink chan<- *IReleaseManagerAVSDeregistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerAVSDeregistered)
				if err := _IReleaseManager.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
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

// ParseAVSDeregistered is a log parse operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) ParseAVSDeregistered(log types.Log) (*IReleaseManagerAVSDeregistered, error) {
	event := new(IReleaseManagerAVSDeregistered)
	if err := _IReleaseManager.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReleaseManagerAVSRegisteredIterator is returned from FilterAVSRegistered and is used to iterate over the raw logs and unpacked data for AVSRegistered events raised by the IReleaseManager contract.
type IReleaseManagerAVSRegisteredIterator struct {
	Event *IReleaseManagerAVSRegistered // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerAVSRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerAVSRegistered)
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
		it.Event = new(IReleaseManagerAVSRegistered)
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
func (it *IReleaseManagerAVSRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerAVSRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerAVSRegistered represents a AVSRegistered event raised by the IReleaseManager contract.
type IReleaseManagerAVSRegistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistered is a free log retrieval operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) FilterAVSRegistered(opts *bind.FilterOpts, avs []common.Address) (*IReleaseManagerAVSRegisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerAVSRegisteredIterator{contract: _IReleaseManager.contract, event: "AVSRegistered", logs: logs, sub: sub}, nil
}

// WatchAVSRegistered is a free log subscription operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) WatchAVSRegistered(opts *bind.WatchOpts, sink chan<- *IReleaseManagerAVSRegistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerAVSRegistered)
				if err := _IReleaseManager.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
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

// ParseAVSRegistered is a log parse operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_IReleaseManager *IReleaseManagerFilterer) ParseAVSRegistered(log types.Log) (*IReleaseManagerAVSRegistered, error) {
	event := new(IReleaseManagerAVSRegistered)
	if err := _IReleaseManager.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReleaseManagerReleaseDeprecatedIterator is returned from FilterReleaseDeprecated and is used to iterate over the raw logs and unpacked data for ReleaseDeprecated events raised by the IReleaseManager contract.
type IReleaseManagerReleaseDeprecatedIterator struct {
	Event *IReleaseManagerReleaseDeprecated // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerReleaseDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerReleaseDeprecated)
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
		it.Event = new(IReleaseManagerReleaseDeprecated)
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
func (it *IReleaseManagerReleaseDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerReleaseDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerReleaseDeprecated represents a ReleaseDeprecated event raised by the IReleaseManager contract.
type IReleaseManagerReleaseDeprecated struct {
	Avs    common.Address
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleaseDeprecated is a free log retrieval operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_IReleaseManager *IReleaseManagerFilterer) FilterReleaseDeprecated(opts *bind.FilterOpts, avs []common.Address, digest [][32]byte) (*IReleaseManagerReleaseDeprecatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerReleaseDeprecatedIterator{contract: _IReleaseManager.contract, event: "ReleaseDeprecated", logs: logs, sub: sub}, nil
}

// WatchReleaseDeprecated is a free log subscription operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_IReleaseManager *IReleaseManagerFilterer) WatchReleaseDeprecated(opts *bind.WatchOpts, sink chan<- *IReleaseManagerReleaseDeprecated, avs []common.Address, digest [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerReleaseDeprecated)
				if err := _IReleaseManager.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
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

// ParseReleaseDeprecated is a log parse operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_IReleaseManager *IReleaseManagerFilterer) ParseReleaseDeprecated(log types.Log) (*IReleaseManagerReleaseDeprecated, error) {
	event := new(IReleaseManagerReleaseDeprecated)
	if err := _IReleaseManager.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReleaseManagerReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the IReleaseManager contract.
type IReleaseManagerReleasePublishedIterator struct {
	Event *IReleaseManagerReleasePublished // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerReleasePublished)
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
		it.Event = new(IReleaseManagerReleasePublished)
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
func (it *IReleaseManagerReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerReleasePublished represents a ReleasePublished event raised by the IReleaseManager contract.
type IReleaseManagerReleasePublished struct {
	Avs                common.Address
	Version            common.Hash
	Digest             [32]byte
	RegistryUrl        string
	DeploymentDeadline *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReleasePublished is a free log retrieval operation binding the contract event 0x5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd2.
//
// Solidity: event ReleasePublished(address indexed avs, string indexed version, bytes32 indexed digest, string registryUrl, uint256 deploymentDeadline)
func (_IReleaseManager *IReleaseManagerFilterer) FilterReleasePublished(opts *bind.FilterOpts, avs []common.Address, version []string, digest [][32]byte) (*IReleaseManagerReleasePublishedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerReleasePublishedIterator{contract: _IReleaseManager.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd2.
//
// Solidity: event ReleasePublished(address indexed avs, string indexed version, bytes32 indexed digest, string registryUrl, uint256 deploymentDeadline)
func (_IReleaseManager *IReleaseManagerFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *IReleaseManagerReleasePublished, avs []common.Address, version []string, digest [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerReleasePublished)
				if err := _IReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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

// ParseReleasePublished is a log parse operation binding the contract event 0x5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd2.
//
// Solidity: event ReleasePublished(address indexed avs, string indexed version, bytes32 indexed digest, string registryUrl, uint256 deploymentDeadline)
func (_IReleaseManager *IReleaseManagerFilterer) ParseReleasePublished(log types.Log) (*IReleaseManagerReleasePublished, error) {
	event := new(IReleaseManagerReleasePublished)
	if err := _IReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
