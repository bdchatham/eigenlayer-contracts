// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReleaseManagerStorage

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

// ReleaseManagerStorageMetaData contains all meta data concerning the ReleaseManagerStorage contract.
var ReleaseManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allPublishedReleases\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deprecateRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecatedReleases\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeprecatedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublishedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManager.PublishedRelease[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseAtBlock\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseCheckpointCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDeprecated\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSDeregistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AVSAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AVSNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyDeprecated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDigest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// ReleaseManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ReleaseManagerStorageMetaData.ABI instead.
var ReleaseManagerStorageABI = ReleaseManagerStorageMetaData.ABI

// ReleaseManagerStorage is an auto generated Go binding around an Ethereum contract.
type ReleaseManagerStorage struct {
	ReleaseManagerStorageCaller     // Read-only binding to the contract
	ReleaseManagerStorageTransactor // Write-only binding to the contract
	ReleaseManagerStorageFilterer   // Log filterer for contract events
}

// ReleaseManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseManagerStorageSession struct {
	Contract     *ReleaseManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReleaseManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseManagerStorageCallerSession struct {
	Contract *ReleaseManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ReleaseManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseManagerStorageTransactorSession struct {
	Contract     *ReleaseManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ReleaseManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseManagerStorageRaw struct {
	Contract *ReleaseManagerStorage // Generic contract binding to access the raw methods on
}

// ReleaseManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseManagerStorageCallerRaw struct {
	Contract *ReleaseManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseManagerStorageTransactorRaw struct {
	Contract *ReleaseManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseManagerStorage creates a new instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorage(address common.Address, backend bind.ContractBackend) (*ReleaseManagerStorage, error) {
	contract, err := bindReleaseManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorage{ReleaseManagerStorageCaller: ReleaseManagerStorageCaller{contract: contract}, ReleaseManagerStorageTransactor: ReleaseManagerStorageTransactor{contract: contract}, ReleaseManagerStorageFilterer: ReleaseManagerStorageFilterer{contract: contract}}, nil
}

// NewReleaseManagerStorageCaller creates a new read-only instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*ReleaseManagerStorageCaller, error) {
	contract, err := bindReleaseManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageCaller{contract: contract}, nil
}

// NewReleaseManagerStorageTransactor creates a new write-only instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseManagerStorageTransactor, error) {
	contract, err := bindReleaseManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageTransactor{contract: contract}, nil
}

// NewReleaseManagerStorageFilterer creates a new log filterer instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseManagerStorageFilterer, error) {
	contract, err := bindReleaseManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageFilterer{contract: contract}, nil
}

// bindReleaseManagerStorage binds a generic wrapper to an already deployed contract.
func bindReleaseManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReleaseManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// AllPublishedReleases is a free data retrieval call binding the contract method 0x7754737d.
//
// Solidity: function allPublishedReleases(address , uint256 ) view returns(bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline, uint256 publishedAt)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) AllPublishedReleases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "allPublishedReleases", arg0, arg1)

	outstruct := new(struct {
		Digest             [32]byte
		RegistryUrl        string
		Version            string
		DeploymentDeadline *big.Int
		PublishedAt        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Digest = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RegistryUrl = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.DeploymentDeadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PublishedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllPublishedReleases is a free data retrieval call binding the contract method 0x7754737d.
//
// Solidity: function allPublishedReleases(address , uint256 ) view returns(bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline, uint256 publishedAt)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) AllPublishedReleases(arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	return _ReleaseManagerStorage.Contract.AllPublishedReleases(&_ReleaseManagerStorage.CallOpts, arg0, arg1)
}

// AllPublishedReleases is a free data retrieval call binding the contract method 0x7754737d.
//
// Solidity: function allPublishedReleases(address , uint256 ) view returns(bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline, uint256 publishedAt)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) AllPublishedReleases(arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	return _ReleaseManagerStorage.Contract.AllPublishedReleases(&_ReleaseManagerStorage.CallOpts, arg0, arg1)
}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) DeprecatedReleases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "deprecatedReleases", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) DeprecatedReleases(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ReleaseManagerStorage.Contract.DeprecatedReleases(&_ReleaseManagerStorage.CallOpts, arg0, arg1)
}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) DeprecatedReleases(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ReleaseManagerStorage.Contract.DeprecatedReleases(&_ReleaseManagerStorage.CallOpts, arg0, arg1)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetDeprecatedReleases(opts *bind.CallOpts, avs common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getDeprecatedReleases", avs)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _ReleaseManagerStorage.Contract.GetDeprecatedReleases(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _ReleaseManagerStorage.Contract.GetDeprecatedReleases(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetLatestRelease(opts *bind.CallOpts, avs common.Address) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getLatestRelease", avs)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetLatestRelease(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetLatestRelease(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetPublishedReleases(opts *bind.CallOpts, avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getPublishedReleases", avs)

	if err != nil {
		return *new([]IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new([]IReleaseManagerPublishedRelease)).(*[]IReleaseManagerPublishedRelease)

	return out0, err

}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetPublishedReleases(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetPublishedReleases(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetReleaseAtBlock(opts *bind.CallOpts, avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getReleaseAtBlock", avs, blockNumber)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetReleaseAtBlock(&_ReleaseManagerStorage.CallOpts, avs, blockNumber)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManagerStorage.Contract.GetReleaseAtBlock(&_ReleaseManagerStorage.CallOpts, avs, blockNumber)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetReleaseCheckpointCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getReleaseCheckpointCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _ReleaseManagerStorage.Contract.GetReleaseCheckpointCount(&_ReleaseManagerStorage.CallOpts, avs)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _ReleaseManagerStorage.Contract.GetReleaseCheckpointCount(&_ReleaseManagerStorage.CallOpts, avs)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) IsDeprecated(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "isDeprecated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) IsDeprecated(arg0 [32]byte) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsDeprecated(&_ReleaseManagerStorage.CallOpts, arg0)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) IsDeprecated(arg0 [32]byte) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsDeprecated(&_ReleaseManagerStorage.CallOpts, arg0)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) IsReleaseDeprecated(opts *bind.CallOpts, avs common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "isReleaseDeprecated", avs, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsReleaseDeprecated(&_ReleaseManagerStorage.CallOpts, avs, digest)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsReleaseDeprecated(&_ReleaseManagerStorage.CallOpts, avs, digest)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) PermissionController() (common.Address, error) {
	return _ReleaseManagerStorage.Contract.PermissionController(&_ReleaseManagerStorage.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) PermissionController() (common.Address, error) {
	return _ReleaseManagerStorage.Contract.PermissionController(&_ReleaseManagerStorage.CallOpts)
}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) RegisteredAVS(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "registeredAVS", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) RegisteredAVS(arg0 common.Address) (bool, error) {
	return _ReleaseManagerStorage.Contract.RegisteredAVS(&_ReleaseManagerStorage.CallOpts, arg0)
}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) RegisteredAVS(arg0 common.Address) (bool, error) {
	return _ReleaseManagerStorage.Contract.RegisteredAVS(&_ReleaseManagerStorage.CallOpts, arg0)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) DeprecateRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "deprecateRelease", avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.DeprecateRelease(&_ReleaseManagerStorage.TransactOpts, avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.DeprecateRelease(&_ReleaseManagerStorage.TransactOpts, avs, digest)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) Deregister(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "deregister", avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.Deregister(&_ReleaseManagerStorage.TransactOpts, avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.Deregister(&_ReleaseManagerStorage.TransactOpts, avs)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) PublishRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "publishRelease", avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishRelease(&_ReleaseManagerStorage.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishRelease(&_ReleaseManagerStorage.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) Register(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "register", avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) Register(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.Register(&_ReleaseManagerStorage.TransactOpts, avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) Register(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.Register(&_ReleaseManagerStorage.TransactOpts, avs)
}

// ReleaseManagerStorageAVSDeregisteredIterator is returned from FilterAVSDeregistered and is used to iterate over the raw logs and unpacked data for AVSDeregistered events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageAVSDeregisteredIterator struct {
	Event *ReleaseManagerStorageAVSDeregistered // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageAVSDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageAVSDeregistered)
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
		it.Event = new(ReleaseManagerStorageAVSDeregistered)
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
func (it *ReleaseManagerStorageAVSDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageAVSDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageAVSDeregistered represents a AVSDeregistered event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageAVSDeregistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSDeregistered is a free log retrieval operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterAVSDeregistered(opts *bind.FilterOpts, avs []common.Address) (*ReleaseManagerStorageAVSDeregisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageAVSDeregisteredIterator{contract: _ReleaseManagerStorage.contract, event: "AVSDeregistered", logs: logs, sub: sub}, nil
}

// WatchAVSDeregistered is a free log subscription operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchAVSDeregistered(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageAVSDeregistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageAVSDeregistered)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseAVSDeregistered(log types.Log) (*ReleaseManagerStorageAVSDeregistered, error) {
	event := new(ReleaseManagerStorageAVSDeregistered)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerStorageAVSRegisteredIterator is returned from FilterAVSRegistered and is used to iterate over the raw logs and unpacked data for AVSRegistered events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageAVSRegisteredIterator struct {
	Event *ReleaseManagerStorageAVSRegistered // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageAVSRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageAVSRegistered)
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
		it.Event = new(ReleaseManagerStorageAVSRegistered)
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
func (it *ReleaseManagerStorageAVSRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageAVSRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageAVSRegistered represents a AVSRegistered event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageAVSRegistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistered is a free log retrieval operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterAVSRegistered(opts *bind.FilterOpts, avs []common.Address) (*ReleaseManagerStorageAVSRegisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageAVSRegisteredIterator{contract: _ReleaseManagerStorage.contract, event: "AVSRegistered", logs: logs, sub: sub}, nil
}

// WatchAVSRegistered is a free log subscription operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchAVSRegistered(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageAVSRegistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageAVSRegistered)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseAVSRegistered(log types.Log) (*ReleaseManagerStorageAVSRegistered, error) {
	event := new(ReleaseManagerStorageAVSRegistered)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerStorageReleaseDeprecatedIterator is returned from FilterReleaseDeprecated and is used to iterate over the raw logs and unpacked data for ReleaseDeprecated events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleaseDeprecatedIterator struct {
	Event *ReleaseManagerStorageReleaseDeprecated // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageReleaseDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageReleaseDeprecated)
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
		it.Event = new(ReleaseManagerStorageReleaseDeprecated)
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
func (it *ReleaseManagerStorageReleaseDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageReleaseDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageReleaseDeprecated represents a ReleaseDeprecated event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleaseDeprecated struct {
	Avs    common.Address
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleaseDeprecated is a free log retrieval operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterReleaseDeprecated(opts *bind.FilterOpts, avs []common.Address, digest [][32]byte) (*ReleaseManagerStorageReleaseDeprecatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageReleaseDeprecatedIterator{contract: _ReleaseManagerStorage.contract, event: "ReleaseDeprecated", logs: logs, sub: sub}, nil
}

// WatchReleaseDeprecated is a free log subscription operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchReleaseDeprecated(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageReleaseDeprecated, avs []common.Address, digest [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageReleaseDeprecated)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseReleaseDeprecated(log types.Log) (*ReleaseManagerStorageReleaseDeprecated, error) {
	event := new(ReleaseManagerStorageReleaseDeprecated)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerStorageReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleasePublishedIterator struct {
	Event *ReleaseManagerStorageReleasePublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageReleasePublished)
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
		it.Event = new(ReleaseManagerStorageReleasePublished)
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
func (it *ReleaseManagerStorageReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageReleasePublished represents a ReleasePublished event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleasePublished struct {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterReleasePublished(opts *bind.FilterOpts, avs []common.Address, version []string, digest [][32]byte) (*ReleaseManagerStorageReleasePublishedIterator, error) {

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

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageReleasePublishedIterator{contract: _ReleaseManagerStorage.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd2.
//
// Solidity: event ReleasePublished(address indexed avs, string indexed version, bytes32 indexed digest, string registryUrl, uint256 deploymentDeadline)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageReleasePublished, avs []common.Address, version []string, digest [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageReleasePublished)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseReleasePublished(log types.Log) (*ReleaseManagerStorageReleasePublished, error) {
	event := new(ReleaseManagerStorageReleasePublished)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
