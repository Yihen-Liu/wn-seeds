// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seeds

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SeedsABI is the input ABI used to generate the binding from.
const SeedsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seed\",\"type\":\"address\"}],\"name\":\"addSeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteSeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seedSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"updateCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seed\",\"type\":\"address\"}],\"name\":\"updateSeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"updateSeedSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SeedsFuncSigs maps the 4-byte function signature to its string representation.
var SeedsFuncSigs = map[string]string{
	"10cb98a5": "addSeed(address)",
	"5cfc1a51": "capacity()",
	"06661abd": "count()",
	"735d406f": "deleteSeed(uint256)",
	"4a610148": "seedSetter()",
	"f0503e80": "seeds(uint256)",
	"e4e6a86a": "updateCapacity(uint256)",
	"217ab447": "updateSeed(uint256,address)",
	"6d69dfbe": "updateSeedSetter(address)",
}

// SeedsBin is the compiled bytecode used for deploying new contracts.
var SeedsBin = "0x608060405234801561001057600080fd5b5060405161065838038061065883398101604081905261002f9161005d565b600380546001600160a01b0319166001600160a01b039290921691909117905560008055601060015561008b565b60006020828403121561006e578081fd5b81516001600160a01b0381168114610084578182fd5b9392505050565b6105be8061009a6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635cfc1a51116100665780635cfc1a51146101075780636d69dfbe14610110578063735d406f14610123578063e4e6a86a14610136578063f0503e801461014957610093565b806306661abd1461009857806310cb98a5146100b4578063217ab447146100c95780634a610148146100dc575b600080fd5b6100a160005481565b6040519081526020015b60405180910390f35b6100c76100c2366004610500565b610172565b005b6100c76100d7366004610539565b61025f565b6003546100ef906001600160a01b031681565b6040516001600160a01b0390911681526020016100ab565b6100a160015481565b6100c761011e366004610500565b610357565b6100c7610131366004610521565b6103d3565b6100c7610144366004610521565b61048c565b6100ef610157366004610521565b6002602052600090815260409020546001600160a01b031681565b6003546001600160a01b031633146101c75760405162461bcd60e51b815260206004820152601360248201527229b2b2b21020b2321d102327a92124a22222a760691b60448201526064015b60405180910390fd5b600154600054106102255760405162461bcd60e51b815260206004820152602260248201527f4361706163697479206861766520616368697665207761726e696e67206c696e604482015261329760f11b60648201526084016101be565b60008054815260026020526040812080546001600160a01b0319166001600160a01b03841617905554610259906001610564565b60005550565b6003546001600160a01b031633146102b25760405162461bcd60e51b815260206004820152601660248201527529b2b2b2102ab83230ba329d102327a92124a22222a760511b60448201526064016101be565b60005482106103295760405162461bcd60e51b815260206004820152603760248201527f496e646578206d75737420626520736d616c6c6572207468616e207265616c6c60448201527f7920636170616369747928636f756e742076616c75652900000000000000000060648201526084016101be565b60009182526002602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6003546001600160a01b031633146103b15760405162461bcd60e51b815260206004820152601760248201527f4f776e6572205570646174653a20464f5242494444454e00000000000000000060448201526064016101be565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031633146104265760405162461bcd60e51b815260206004820152601660248201527529b2b2b2102232b632ba329d102327a92124a22222a760511b60448201526064016101be565b600054811061046e5760405162461bcd60e51b815260206004820152601460248201527314995c5d5a5c994e881a5b99195e0f18dbdd5b9d60621b60448201526064016101be565b600090815260026020526040902080546001600160a01b0319169055565b6003546001600160a01b031633146104df5760405162461bcd60e51b815260206004820152601660248201527529b2b2b2102ab83230ba329d102327a92124a22222a760511b60448201526064016101be565b600155565b80356001600160a01b03811681146104fb57600080fd5b919050565b600060208284031215610511578081fd5b61051a826104e4565b9392505050565b600060208284031215610532578081fd5b5035919050565b6000806040838503121561054b578081fd5b8235915061055b602084016104e4565b90509250929050565b6000821982111561058357634e487b7160e01b81526011600452602481fd5b50019056fea2646970667358221220f8041667e51ac5d888e3524e375391cd69a65b2257b348ae1fadb617f2c6ce3864736f6c63430008020033"

// DeploySeeds deploys a new Ethereum contract, binding an instance of Seeds to it.
func DeploySeeds(auth *bind.TransactOpts, backend bind.ContractBackend, setter common.Address) (common.Address, *types.Transaction, *Seeds, error) {
	parsed, err := abi.JSON(strings.NewReader(SeedsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SeedsBin), backend, setter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seeds{SeedsCaller: SeedsCaller{contract: contract}, SeedsTransactor: SeedsTransactor{contract: contract}, SeedsFilterer: SeedsFilterer{contract: contract}}, nil
}

// Seeds is an auto generated Go binding around an Ethereum contract.
type Seeds struct {
	SeedsCaller     // Read-only binding to the contract
	SeedsTransactor // Write-only binding to the contract
	SeedsFilterer   // Log filterer for contract events
}

// SeedsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeedsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeedsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeedsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeedsSession struct {
	Contract     *Seeds            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeedsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeedsCallerSession struct {
	Contract *SeedsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SeedsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeedsTransactorSession struct {
	Contract     *SeedsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeedsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeedsRaw struct {
	Contract *Seeds // Generic contract binding to access the raw methods on
}

// SeedsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeedsCallerRaw struct {
	Contract *SeedsCaller // Generic read-only contract binding to access the raw methods on
}

// SeedsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeedsTransactorRaw struct {
	Contract *SeedsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeeds creates a new instance of Seeds, bound to a specific deployed contract.
func NewSeeds(address common.Address, backend bind.ContractBackend) (*Seeds, error) {
	contract, err := bindSeeds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seeds{SeedsCaller: SeedsCaller{contract: contract}, SeedsTransactor: SeedsTransactor{contract: contract}, SeedsFilterer: SeedsFilterer{contract: contract}}, nil
}

// NewSeedsCaller creates a new read-only instance of Seeds, bound to a specific deployed contract.
func NewSeedsCaller(address common.Address, caller bind.ContractCaller) (*SeedsCaller, error) {
	contract, err := bindSeeds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeedsCaller{contract: contract}, nil
}

// NewSeedsTransactor creates a new write-only instance of Seeds, bound to a specific deployed contract.
func NewSeedsTransactor(address common.Address, transactor bind.ContractTransactor) (*SeedsTransactor, error) {
	contract, err := bindSeeds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeedsTransactor{contract: contract}, nil
}

// NewSeedsFilterer creates a new log filterer instance of Seeds, bound to a specific deployed contract.
func NewSeedsFilterer(address common.Address, filterer bind.ContractFilterer) (*SeedsFilterer, error) {
	contract, err := bindSeeds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeedsFilterer{contract: contract}, nil
}

// bindSeeds binds a generic wrapper to an already deployed contract.
func bindSeeds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeedsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seeds *SeedsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seeds.Contract.SeedsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seeds *SeedsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seeds.Contract.SeedsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seeds *SeedsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seeds.Contract.SeedsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seeds *SeedsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seeds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seeds *SeedsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seeds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seeds *SeedsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seeds.Contract.contract.Transact(opts, method, params...)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Seeds *SeedsCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Seeds.contract.Call(opts, &out, "capacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Seeds *SeedsSession) Capacity() (*big.Int, error) {
	return _Seeds.Contract.Capacity(&_Seeds.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Seeds *SeedsCallerSession) Capacity() (*big.Int, error) {
	return _Seeds.Contract.Capacity(&_Seeds.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Seeds *SeedsCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Seeds.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Seeds *SeedsSession) Count() (*big.Int, error) {
	return _Seeds.Contract.Count(&_Seeds.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Seeds *SeedsCallerSession) Count() (*big.Int, error) {
	return _Seeds.Contract.Count(&_Seeds.CallOpts)
}

// SeedSetter is a free data retrieval call binding the contract method 0x4a610148.
//
// Solidity: function seedSetter() view returns(address)
func (_Seeds *SeedsCaller) SeedSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Seeds.contract.Call(opts, &out, "seedSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SeedSetter is a free data retrieval call binding the contract method 0x4a610148.
//
// Solidity: function seedSetter() view returns(address)
func (_Seeds *SeedsSession) SeedSetter() (common.Address, error) {
	return _Seeds.Contract.SeedSetter(&_Seeds.CallOpts)
}

// SeedSetter is a free data retrieval call binding the contract method 0x4a610148.
//
// Solidity: function seedSetter() view returns(address)
func (_Seeds *SeedsCallerSession) SeedSetter() (common.Address, error) {
	return _Seeds.Contract.SeedSetter(&_Seeds.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(address)
func (_Seeds *SeedsCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Seeds.contract.Call(opts, &out, "seeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(address)
func (_Seeds *SeedsSession) Seeds(arg0 *big.Int) (common.Address, error) {
	return _Seeds.Contract.Seeds(&_Seeds.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(address)
func (_Seeds *SeedsCallerSession) Seeds(arg0 *big.Int) (common.Address, error) {
	return _Seeds.Contract.Seeds(&_Seeds.CallOpts, arg0)
}

// AddSeed is a paid mutator transaction binding the contract method 0x10cb98a5.
//
// Solidity: function addSeed(address seed) returns()
func (_Seeds *SeedsTransactor) AddSeed(opts *bind.TransactOpts, seed common.Address) (*types.Transaction, error) {
	return _Seeds.contract.Transact(opts, "addSeed", seed)
}

// AddSeed is a paid mutator transaction binding the contract method 0x10cb98a5.
//
// Solidity: function addSeed(address seed) returns()
func (_Seeds *SeedsSession) AddSeed(seed common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.AddSeed(&_Seeds.TransactOpts, seed)
}

// AddSeed is a paid mutator transaction binding the contract method 0x10cb98a5.
//
// Solidity: function addSeed(address seed) returns()
func (_Seeds *SeedsTransactorSession) AddSeed(seed common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.AddSeed(&_Seeds.TransactOpts, seed)
}

// DeleteSeed is a paid mutator transaction binding the contract method 0x735d406f.
//
// Solidity: function deleteSeed(uint256 index) returns()
func (_Seeds *SeedsTransactor) DeleteSeed(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Seeds.contract.Transact(opts, "deleteSeed", index)
}

// DeleteSeed is a paid mutator transaction binding the contract method 0x735d406f.
//
// Solidity: function deleteSeed(uint256 index) returns()
func (_Seeds *SeedsSession) DeleteSeed(index *big.Int) (*types.Transaction, error) {
	return _Seeds.Contract.DeleteSeed(&_Seeds.TransactOpts, index)
}

// DeleteSeed is a paid mutator transaction binding the contract method 0x735d406f.
//
// Solidity: function deleteSeed(uint256 index) returns()
func (_Seeds *SeedsTransactorSession) DeleteSeed(index *big.Int) (*types.Transaction, error) {
	return _Seeds.Contract.DeleteSeed(&_Seeds.TransactOpts, index)
}

// UpdateCapacity is a paid mutator transaction binding the contract method 0xe4e6a86a.
//
// Solidity: function updateCapacity(uint256 cap) returns()
func (_Seeds *SeedsTransactor) UpdateCapacity(opts *bind.TransactOpts, cap *big.Int) (*types.Transaction, error) {
	return _Seeds.contract.Transact(opts, "updateCapacity", cap)
}

// UpdateCapacity is a paid mutator transaction binding the contract method 0xe4e6a86a.
//
// Solidity: function updateCapacity(uint256 cap) returns()
func (_Seeds *SeedsSession) UpdateCapacity(cap *big.Int) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateCapacity(&_Seeds.TransactOpts, cap)
}

// UpdateCapacity is a paid mutator transaction binding the contract method 0xe4e6a86a.
//
// Solidity: function updateCapacity(uint256 cap) returns()
func (_Seeds *SeedsTransactorSession) UpdateCapacity(cap *big.Int) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateCapacity(&_Seeds.TransactOpts, cap)
}

// UpdateSeed is a paid mutator transaction binding the contract method 0x217ab447.
//
// Solidity: function updateSeed(uint256 index, address seed) returns()
func (_Seeds *SeedsTransactor) UpdateSeed(opts *bind.TransactOpts, index *big.Int, seed common.Address) (*types.Transaction, error) {
	return _Seeds.contract.Transact(opts, "updateSeed", index, seed)
}

// UpdateSeed is a paid mutator transaction binding the contract method 0x217ab447.
//
// Solidity: function updateSeed(uint256 index, address seed) returns()
func (_Seeds *SeedsSession) UpdateSeed(index *big.Int, seed common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateSeed(&_Seeds.TransactOpts, index, seed)
}

// UpdateSeed is a paid mutator transaction binding the contract method 0x217ab447.
//
// Solidity: function updateSeed(uint256 index, address seed) returns()
func (_Seeds *SeedsTransactorSession) UpdateSeed(index *big.Int, seed common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateSeed(&_Seeds.TransactOpts, index, seed)
}

// UpdateSeedSetter is a paid mutator transaction binding the contract method 0x6d69dfbe.
//
// Solidity: function updateSeedSetter(address setter) returns()
func (_Seeds *SeedsTransactor) UpdateSeedSetter(opts *bind.TransactOpts, setter common.Address) (*types.Transaction, error) {
	return _Seeds.contract.Transact(opts, "updateSeedSetter", setter)
}

// UpdateSeedSetter is a paid mutator transaction binding the contract method 0x6d69dfbe.
//
// Solidity: function updateSeedSetter(address setter) returns()
func (_Seeds *SeedsSession) UpdateSeedSetter(setter common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateSeedSetter(&_Seeds.TransactOpts, setter)
}

// UpdateSeedSetter is a paid mutator transaction binding the contract method 0x6d69dfbe.
//
// Solidity: function updateSeedSetter(address setter) returns()
func (_Seeds *SeedsTransactorSession) UpdateSeedSetter(setter common.Address) (*types.Transaction, error) {
	return _Seeds.Contract.UpdateSeedSetter(&_Seeds.TransactOpts, setter)
}
