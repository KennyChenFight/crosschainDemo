// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clinic

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClinicABI is the input ABI used to generate the binding from.
const ClinicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"InfoSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"infos\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ClinicBin is the compiled bytecode used for deploying new contracts.
var ClinicBin = "0x608060405234801561001057600080fd5b5061015c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063369e398c1461003b578063cf3998dc1461007d575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100b5565b6040518082815260200191505060405180910390f35b6100b36004803603604081101561009357600080fd5b8101908080359060200190929190803590602001909291905050506100cd565b005b60006020528060005260406000206000915090505481565b80600080848152602001908152602001600020819055507f6297a27c0f9211725834ab3dc8d3f3960145a253d5fdbfebb042db832a5082238282604051808381526020018281526020019250505060405180910390a1505056fea265627a7a72315820f1946e96bc6e11a97d79a3241c946b7070b01d78d52c694c2d97b8c678d2cb1564736f6c634300050d0032"

// DeployClinic deploys a new Ethereum contract, binding an instance of Clinic to it.
func DeployClinic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clinic, error) {
	parsed, err := abi.JSON(strings.NewReader(ClinicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClinicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clinic{ClinicCaller: ClinicCaller{contract: contract}, ClinicTransactor: ClinicTransactor{contract: contract}, ClinicFilterer: ClinicFilterer{contract: contract}}, nil
}

// Clinic is an auto generated Go binding around an Ethereum contract.
type Clinic struct {
	ClinicCaller     // Read-only binding to the contract
	ClinicTransactor // Write-only binding to the contract
	ClinicFilterer   // Log filterer for contract events
}

// ClinicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClinicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClinicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClinicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClinicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClinicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClinicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClinicSession struct {
	Contract     *Clinic           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClinicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClinicCallerSession struct {
	Contract *ClinicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClinicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClinicTransactorSession struct {
	Contract     *ClinicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClinicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClinicRaw struct {
	Contract *Clinic // Generic contract binding to access the raw methods on
}

// ClinicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClinicCallerRaw struct {
	Contract *ClinicCaller // Generic read-only contract binding to access the raw methods on
}

// ClinicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClinicTransactorRaw struct {
	Contract *ClinicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClinic creates a new instance of Clinic, bound to a specific deployed contract.
func NewClinic(address common.Address, backend bind.ContractBackend) (*Clinic, error) {
	contract, err := bindClinic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clinic{ClinicCaller: ClinicCaller{contract: contract}, ClinicTransactor: ClinicTransactor{contract: contract}, ClinicFilterer: ClinicFilterer{contract: contract}}, nil
}

// NewClinicCaller creates a new read-only instance of Clinic, bound to a specific deployed contract.
func NewClinicCaller(address common.Address, caller bind.ContractCaller) (*ClinicCaller, error) {
	contract, err := bindClinic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClinicCaller{contract: contract}, nil
}

// NewClinicTransactor creates a new write-only instance of Clinic, bound to a specific deployed contract.
func NewClinicTransactor(address common.Address, transactor bind.ContractTransactor) (*ClinicTransactor, error) {
	contract, err := bindClinic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClinicTransactor{contract: contract}, nil
}

// NewClinicFilterer creates a new log filterer instance of Clinic, bound to a specific deployed contract.
func NewClinicFilterer(address common.Address, filterer bind.ContractFilterer) (*ClinicFilterer, error) {
	contract, err := bindClinic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClinicFilterer{contract: contract}, nil
}

// bindClinic binds a generic wrapper to an already deployed contract.
func bindClinic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClinicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clinic *ClinicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clinic.Contract.ClinicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clinic *ClinicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clinic.Contract.ClinicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clinic *ClinicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clinic.Contract.ClinicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clinic *ClinicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clinic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clinic *ClinicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clinic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clinic *ClinicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clinic.Contract.contract.Transact(opts, method, params...)
}

// Infos is a free data retrieval call binding the contract method 0x369e398c.
//
// Solidity: function infos(bytes32 ) constant returns(bytes32)
func (_Clinic *ClinicCaller) Infos(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Clinic.contract.Call(opts, out, "infos", arg0)
	return *ret0, err
}

// Infos is a free data retrieval call binding the contract method 0x369e398c.
//
// Solidity: function infos(bytes32 ) constant returns(bytes32)
func (_Clinic *ClinicSession) Infos(arg0 [32]byte) ([32]byte, error) {
	return _Clinic.Contract.Infos(&_Clinic.CallOpts, arg0)
}

// Infos is a free data retrieval call binding the contract method 0x369e398c.
//
// Solidity: function infos(bytes32 ) constant returns(bytes32)
func (_Clinic *ClinicCallerSession) Infos(arg0 [32]byte) ([32]byte, error) {
	return _Clinic.Contract.Infos(&_Clinic.CallOpts, arg0)
}

// SetInfo is a paid mutator transaction binding the contract method 0xcf3998dc.
//
// Solidity: function setInfo(bytes32 key, bytes32 value) returns()
func (_Clinic *ClinicTransactor) SetInfo(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Clinic.contract.Transact(opts, "setInfo", key, value)
}

// SetInfo is a paid mutator transaction binding the contract method 0xcf3998dc.
//
// Solidity: function setInfo(bytes32 key, bytes32 value) returns()
func (_Clinic *ClinicSession) SetInfo(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Clinic.Contract.SetInfo(&_Clinic.TransactOpts, key, value)
}

// SetInfo is a paid mutator transaction binding the contract method 0xcf3998dc.
//
// Solidity: function setInfo(bytes32 key, bytes32 value) returns()
func (_Clinic *ClinicTransactorSession) SetInfo(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Clinic.Contract.SetInfo(&_Clinic.TransactOpts, key, value)
}

// ClinicInfoSetIterator is returned from FilterInfoSet and is used to iterate over the raw logs and unpacked data for InfoSet events raised by the Clinic contract.
type ClinicInfoSetIterator struct {
	Event *ClinicInfoSet // Event containing the contract specifics and raw log

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
func (it *ClinicInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClinicInfoSet)
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
		it.Event = new(ClinicInfoSet)
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
func (it *ClinicInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClinicInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClinicInfoSet represents a InfoSet event raised by the Clinic contract.
type ClinicInfoSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInfoSet is a free log retrieval operation binding the contract event 0x6297a27c0f9211725834ab3dc8d3f3960145a253d5fdbfebb042db832a508223.
//
// Solidity: event InfoSet(bytes32 key, bytes32 value)
func (_Clinic *ClinicFilterer) FilterInfoSet(opts *bind.FilterOpts) (*ClinicInfoSetIterator, error) {

	logs, sub, err := _Clinic.contract.FilterLogs(opts, "InfoSet")
	if err != nil {
		return nil, err
	}
	return &ClinicInfoSetIterator{contract: _Clinic.contract, event: "InfoSet", logs: logs, sub: sub}, nil
}

// WatchInfoSet is a free log subscription operation binding the contract event 0x6297a27c0f9211725834ab3dc8d3f3960145a253d5fdbfebb042db832a508223.
//
// Solidity: event InfoSet(bytes32 key, bytes32 value)
func (_Clinic *ClinicFilterer) WatchInfoSet(opts *bind.WatchOpts, sink chan<- *ClinicInfoSet) (event.Subscription, error) {

	logs, sub, err := _Clinic.contract.WatchLogs(opts, "InfoSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClinicInfoSet)
				if err := _Clinic.contract.UnpackLog(event, "InfoSet", log); err != nil {
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

// ParseInfoSet is a log parse operation binding the contract event 0x6297a27c0f9211725834ab3dc8d3f3960145a253d5fdbfebb042db832a508223.
//
// Solidity: event InfoSet(bytes32 key, bytes32 value)
func (_Clinic *ClinicFilterer) ParseInfoSet(log types.Log) (*ClinicInfoSet, error) {
	event := new(ClinicInfoSet)
	if err := _Clinic.contract.UnpackLog(event, "InfoSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
