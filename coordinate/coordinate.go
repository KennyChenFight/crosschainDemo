// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coordinate

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

// CoordinateABI is the input ABI used to generate the binding from.
const CoordinateABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TransactionSet\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"addTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CoordinateBin is the compiled bytecode used for deploying new contracts.
var CoordinateBin = "0x608060405234801561001057600080fd5b506104b1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063022236f41461003b5780631083fa231461016f575b600080fd5b6100f46004803603602081101561005157600080fd5b810190808035906020019064010000000081111561006e57600080fd5b82018360208201111561008057600080fd5b803590602001918460018302840111640100000000831117156100a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061023d565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610134578082015181840152602081019050610119565b50505050905090810190601f1680156101615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61023b6004803603604081101561018557600080fd5b81019080803590602001906401000000008111156101a257600080fd5b8201836020820111156101b457600080fd5b803590602001918460018302840111640100000000831117156101d657600080fd5b9091929391929390803590602001906401000000008111156101f757600080fd5b82018360208201111561020957600080fd5b8035906020019184600183028401116401000000008311171561022b57600080fd5b9091929391929390505050610303565b005b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102fb5780601f106102d0576101008083540402835291602001916102fb565b820191906000526020600020905b8154815290600101906020018083116102de57829003601f168201915b505050505081565b818160008686604051808383808284378083019250505092505050908152602001604051809103902091906103399291906103d7565b507f487097b8f2eecdf832266b502ff8357117b5ba21d87cc91e1824875ebe631210848484846040518080602001806020018381038352878782818152602001925080828437600081840152601f19601f8201169050808301925050508381038252858582818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060405180910390a150505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061041857803560ff1916838001178555610446565b82800160010185558215610446579182015b8281111561044557823582559160200191906001019061042a565b5b5090506104539190610457565b5090565b61047991905b8082111561047557600081600090555060010161045d565b5090565b9056fea265627a7a72315820e284c39b62fc433373330f56ecbccf32db110035c07de1dcf89d3f598c5834e064736f6c634300050d0032"

// DeployCoordinate deploys a new Ethereum contract, binding an instance of Coordinate to it.
func DeployCoordinate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Coordinate, error) {
	parsed, err := abi.JSON(strings.NewReader(CoordinateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoordinateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coordinate{CoordinateCaller: CoordinateCaller{contract: contract}, CoordinateTransactor: CoordinateTransactor{contract: contract}, CoordinateFilterer: CoordinateFilterer{contract: contract}}, nil
}

// Coordinate is an auto generated Go binding around an Ethereum contract.
type Coordinate struct {
	CoordinateCaller     // Read-only binding to the contract
	CoordinateTransactor // Write-only binding to the contract
	CoordinateFilterer   // Log filterer for contract events
}

// CoordinateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoordinateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoordinateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoordinateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoordinateSession struct {
	Contract     *Coordinate       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoordinateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoordinateCallerSession struct {
	Contract *CoordinateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoordinateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoordinateTransactorSession struct {
	Contract     *CoordinateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoordinateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoordinateRaw struct {
	Contract *Coordinate // Generic contract binding to access the raw methods on
}

// CoordinateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoordinateCallerRaw struct {
	Contract *CoordinateCaller // Generic read-only contract binding to access the raw methods on
}

// CoordinateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoordinateTransactorRaw struct {
	Contract *CoordinateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoordinate creates a new instance of Coordinate, bound to a specific deployed contract.
func NewCoordinate(address common.Address, backend bind.ContractBackend) (*Coordinate, error) {
	contract, err := bindCoordinate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coordinate{CoordinateCaller: CoordinateCaller{contract: contract}, CoordinateTransactor: CoordinateTransactor{contract: contract}, CoordinateFilterer: CoordinateFilterer{contract: contract}}, nil
}

// NewCoordinateCaller creates a new read-only instance of Coordinate, bound to a specific deployed contract.
func NewCoordinateCaller(address common.Address, caller bind.ContractCaller) (*CoordinateCaller, error) {
	contract, err := bindCoordinate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinateCaller{contract: contract}, nil
}

// NewCoordinateTransactor creates a new write-only instance of Coordinate, bound to a specific deployed contract.
func NewCoordinateTransactor(address common.Address, transactor bind.ContractTransactor) (*CoordinateTransactor, error) {
	contract, err := bindCoordinate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinateTransactor{contract: contract}, nil
}

// NewCoordinateFilterer creates a new log filterer instance of Coordinate, bound to a specific deployed contract.
func NewCoordinateFilterer(address common.Address, filterer bind.ContractFilterer) (*CoordinateFilterer, error) {
	contract, err := bindCoordinate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoordinateFilterer{contract: contract}, nil
}

// bindCoordinate binds a generic wrapper to an already deployed contract.
func bindCoordinate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoordinateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinate *CoordinateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coordinate.Contract.CoordinateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinate *CoordinateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinate.Contract.CoordinateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinate *CoordinateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinate.Contract.CoordinateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinate *CoordinateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coordinate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinate *CoordinateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinate *CoordinateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinate.Contract.contract.Transact(opts, method, params...)
}

// Transactions is a free data retrieval call binding the contract method 0x022236f4.
//
// Solidity: function transactions(string ) constant returns(string)
func (_Coordinate *CoordinateCaller) Transactions(opts *bind.CallOpts, arg0 string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coordinate.contract.Call(opts, out, "transactions", arg0)
	return *ret0, err
}

// Transactions is a free data retrieval call binding the contract method 0x022236f4.
//
// Solidity: function transactions(string ) constant returns(string)
func (_Coordinate *CoordinateSession) Transactions(arg0 string) (string, error) {
	return _Coordinate.Contract.Transactions(&_Coordinate.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x022236f4.
//
// Solidity: function transactions(string ) constant returns(string)
func (_Coordinate *CoordinateCallerSession) Transactions(arg0 string) (string, error) {
	return _Coordinate.Contract.Transactions(&_Coordinate.CallOpts, arg0)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x1083fa23.
//
// Solidity: function addTransaction(string key, string value) returns()
func (_Coordinate *CoordinateTransactor) AddTransaction(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Coordinate.contract.Transact(opts, "addTransaction", key, value)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x1083fa23.
//
// Solidity: function addTransaction(string key, string value) returns()
func (_Coordinate *CoordinateSession) AddTransaction(key string, value string) (*types.Transaction, error) {
	return _Coordinate.Contract.AddTransaction(&_Coordinate.TransactOpts, key, value)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x1083fa23.
//
// Solidity: function addTransaction(string key, string value) returns()
func (_Coordinate *CoordinateTransactorSession) AddTransaction(key string, value string) (*types.Transaction, error) {
	return _Coordinate.Contract.AddTransaction(&_Coordinate.TransactOpts, key, value)
}

// CoordinateTransactionSetIterator is returned from FilterTransactionSet and is used to iterate over the raw logs and unpacked data for TransactionSet events raised by the Coordinate contract.
type CoordinateTransactionSetIterator struct {
	Event *CoordinateTransactionSet // Event containing the contract specifics and raw log

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
func (it *CoordinateTransactionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinateTransactionSet)
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
		it.Event = new(CoordinateTransactionSet)
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
func (it *CoordinateTransactionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinateTransactionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinateTransactionSet represents a TransactionSet event raised by the Coordinate contract.
type CoordinateTransactionSet struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransactionSet is a free log retrieval operation binding the contract event 0x487097b8f2eecdf832266b502ff8357117b5ba21d87cc91e1824875ebe631210.
//
// Solidity: event TransactionSet(string key, string value)
func (_Coordinate *CoordinateFilterer) FilterTransactionSet(opts *bind.FilterOpts) (*CoordinateTransactionSetIterator, error) {

	logs, sub, err := _Coordinate.contract.FilterLogs(opts, "TransactionSet")
	if err != nil {
		return nil, err
	}
	return &CoordinateTransactionSetIterator{contract: _Coordinate.contract, event: "TransactionSet", logs: logs, sub: sub}, nil
}

// WatchTransactionSet is a free log subscription operation binding the contract event 0x487097b8f2eecdf832266b502ff8357117b5ba21d87cc91e1824875ebe631210.
//
// Solidity: event TransactionSet(string key, string value)
func (_Coordinate *CoordinateFilterer) WatchTransactionSet(opts *bind.WatchOpts, sink chan<- *CoordinateTransactionSet) (event.Subscription, error) {

	logs, sub, err := _Coordinate.contract.WatchLogs(opts, "TransactionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinateTransactionSet)
				if err := _Coordinate.contract.UnpackLog(event, "TransactionSet", log); err != nil {
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

// ParseTransactionSet is a log parse operation binding the contract event 0x487097b8f2eecdf832266b502ff8357117b5ba21d87cc91e1824875ebe631210.
//
// Solidity: event TransactionSet(string key, string value)
func (_Coordinate *CoordinateFilterer) ParseTransactionSet(log types.Log) (*CoordinateTransactionSet, error) {
	event := new(CoordinateTransactionSet)
	if err := _Coordinate.contract.UnpackLog(event, "TransactionSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
