// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hospital

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

// HospitalABI is the input ABI used to generate the binding from.
const HospitalABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"InfoSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"infos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HospitalBin is the compiled bytecode used for deploying new contracts.
var HospitalBin = "0x608060405234801561001057600080fd5b50610652806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637b973bd014610046578063a26493e91461017a578063a923fc401461024d575b600080fd5b6100ff6004803603602081101561005c57600080fd5b810190808035906020019064010000000081111561007957600080fd5b82018360208201111561008b57600080fd5b803590602001918460018302840111640100000000831117156100ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061031b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013f578082015181840152602081019050610124565b50505050905090810190601f16801561016c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102336004803603602081101561019057600080fd5b81019080803590602001906401000000008111156101ad57600080fd5b8201836020820111156101bf57600080fd5b803590602001918460018302840111640100000000831117156101e157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506103e1565b604051808215151515815260200191505060405180910390f35b6103196004803603604081101561026357600080fd5b810190808035906020019064010000000081111561028057600080fd5b82018360208201111561029257600080fd5b803590602001918460018302840111640100000000831117156102b457600080fd5b9091929391929390803590602001906401000000008111156102d557600080fd5b8201836020820111156102e757600080fd5b8035906020019184600183028401116401000000008311171561030957600080fd5b90919293919293905050506104a4565b005b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d95780601f106103ae576101008083540402835291602001916103d9565b820191906000526020600020905b8154815290600101906020018083116103bc57829003601f168201915b505050505081565b600060405160200180807f636c696e69630000000000000000000000000000000000000000000000000000815250600601905060405160208183030381529060405280519060200120826040516020018082805190602001908083835b60208310610461578051825260208201915060208101905060208303925061043e565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120149050919050565b818160008686604051808383808284378083019250505092505050908152602001604051809103902091906104da929190610578565b507fbb4663e939ed98b60f591a4f61e38a306d703e68557951430c51a3cc0eb4f3fa848484846040518080602001806020018381038352878782818152602001925080828437600081840152601f19601f8201169050808301925050508381038252858582818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060405180910390a150505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105b957803560ff19168380011785556105e7565b828001600101855582156105e7579182015b828111156105e65782358255916020019190600101906105cb565b5b5090506105f491906105f8565b5090565b61061a91905b808211156106165760008160009055506001016105fe565b5090565b9056fea265627a7a723158202737f566e405af09485ac8b5c5a81f02fcd245a173fc514ba83c859d4f8856d364736f6c634300050d0032"

// DeployHospital deploys a new Ethereum contract, binding an instance of Hospital to it.
func DeployHospital(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hospital, error) {
	parsed, err := abi.JSON(strings.NewReader(HospitalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HospitalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hospital{HospitalCaller: HospitalCaller{contract: contract}, HospitalTransactor: HospitalTransactor{contract: contract}, HospitalFilterer: HospitalFilterer{contract: contract}}, nil
}

// Hospital is an auto generated Go binding around an Ethereum contract.
type Hospital struct {
	HospitalCaller     // Read-only binding to the contract
	HospitalTransactor // Write-only binding to the contract
	HospitalFilterer   // Log filterer for contract events
}

// HospitalCaller is an auto generated read-only Go binding around an Ethereum contract.
type HospitalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HospitalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HospitalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HospitalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HospitalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HospitalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HospitalSession struct {
	Contract     *Hospital         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HospitalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HospitalCallerSession struct {
	Contract *HospitalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HospitalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HospitalTransactorSession struct {
	Contract     *HospitalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HospitalRaw is an auto generated low-level Go binding around an Ethereum contract.
type HospitalRaw struct {
	Contract *Hospital // Generic contract binding to access the raw methods on
}

// HospitalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HospitalCallerRaw struct {
	Contract *HospitalCaller // Generic read-only contract binding to access the raw methods on
}

// HospitalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HospitalTransactorRaw struct {
	Contract *HospitalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHospital creates a new instance of Hospital, bound to a specific deployed contract.
func NewHospital(address common.Address, backend bind.ContractBackend) (*Hospital, error) {
	contract, err := bindHospital(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hospital{HospitalCaller: HospitalCaller{contract: contract}, HospitalTransactor: HospitalTransactor{contract: contract}, HospitalFilterer: HospitalFilterer{contract: contract}}, nil
}

// NewHospitalCaller creates a new read-only instance of Hospital, bound to a specific deployed contract.
func NewHospitalCaller(address common.Address, caller bind.ContractCaller) (*HospitalCaller, error) {
	contract, err := bindHospital(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HospitalCaller{contract: contract}, nil
}

// NewHospitalTransactor creates a new write-only instance of Hospital, bound to a specific deployed contract.
func NewHospitalTransactor(address common.Address, transactor bind.ContractTransactor) (*HospitalTransactor, error) {
	contract, err := bindHospital(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HospitalTransactor{contract: contract}, nil
}

// NewHospitalFilterer creates a new log filterer instance of Hospital, bound to a specific deployed contract.
func NewHospitalFilterer(address common.Address, filterer bind.ContractFilterer) (*HospitalFilterer, error) {
	contract, err := bindHospital(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HospitalFilterer{contract: contract}, nil
}

// bindHospital binds a generic wrapper to an already deployed contract.
func bindHospital(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HospitalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hospital *HospitalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hospital.Contract.HospitalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hospital *HospitalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hospital.Contract.HospitalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hospital *HospitalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hospital.Contract.HospitalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hospital *HospitalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hospital.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hospital *HospitalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hospital.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hospital *HospitalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hospital.Contract.contract.Transact(opts, method, params...)
}

// Infos is a free data retrieval call binding the contract method 0x7b973bd0.
//
// Solidity: function infos(string ) constant returns(string)
func (_Hospital *HospitalCaller) Infos(opts *bind.CallOpts, arg0 string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hospital.contract.Call(opts, out, "infos", arg0)
	return *ret0, err
}

// Infos is a free data retrieval call binding the contract method 0x7b973bd0.
//
// Solidity: function infos(string ) constant returns(string)
func (_Hospital *HospitalSession) Infos(arg0 string) (string, error) {
	return _Hospital.Contract.Infos(&_Hospital.CallOpts, arg0)
}

// Infos is a free data retrieval call binding the contract method 0x7b973bd0.
//
// Solidity: function infos(string ) constant returns(string)
func (_Hospital *HospitalCallerSession) Infos(arg0 string) (string, error) {
	return _Hospital.Contract.Infos(&_Hospital.CallOpts, arg0)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xa26493e9.
//
// Solidity: function verifyMessage(string message) constant returns(bool)
func (_Hospital *HospitalCaller) VerifyMessage(opts *bind.CallOpts, message string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Hospital.contract.Call(opts, out, "verifyMessage", message)
	return *ret0, err
}

// VerifyMessage is a free data retrieval call binding the contract method 0xa26493e9.
//
// Solidity: function verifyMessage(string message) constant returns(bool)
func (_Hospital *HospitalSession) VerifyMessage(message string) (bool, error) {
	return _Hospital.Contract.VerifyMessage(&_Hospital.CallOpts, message)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xa26493e9.
//
// Solidity: function verifyMessage(string message) constant returns(bool)
func (_Hospital *HospitalCallerSession) VerifyMessage(message string) (bool, error) {
	return _Hospital.Contract.VerifyMessage(&_Hospital.CallOpts, message)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string key, string value) returns()
func (_Hospital *HospitalTransactor) SetInfo(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Hospital.contract.Transact(opts, "setInfo", key, value)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string key, string value) returns()
func (_Hospital *HospitalSession) SetInfo(key string, value string) (*types.Transaction, error) {
	return _Hospital.Contract.SetInfo(&_Hospital.TransactOpts, key, value)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string key, string value) returns()
func (_Hospital *HospitalTransactorSession) SetInfo(key string, value string) (*types.Transaction, error) {
	return _Hospital.Contract.SetInfo(&_Hospital.TransactOpts, key, value)
}

// HospitalInfoSetIterator is returned from FilterInfoSet and is used to iterate over the raw logs and unpacked data for InfoSet events raised by the Hospital contract.
type HospitalInfoSetIterator struct {
	Event *HospitalInfoSet // Event containing the contract specifics and raw log

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
func (it *HospitalInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HospitalInfoSet)
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
		it.Event = new(HospitalInfoSet)
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
func (it *HospitalInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HospitalInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HospitalInfoSet represents a InfoSet event raised by the Hospital contract.
type HospitalInfoSet struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInfoSet is a free log retrieval operation binding the contract event 0xbb4663e939ed98b60f591a4f61e38a306d703e68557951430c51a3cc0eb4f3fa.
//
// Solidity: event InfoSet(string key, string value)
func (_Hospital *HospitalFilterer) FilterInfoSet(opts *bind.FilterOpts) (*HospitalInfoSetIterator, error) {

	logs, sub, err := _Hospital.contract.FilterLogs(opts, "InfoSet")
	if err != nil {
		return nil, err
	}
	return &HospitalInfoSetIterator{contract: _Hospital.contract, event: "InfoSet", logs: logs, sub: sub}, nil
}

// WatchInfoSet is a free log subscription operation binding the contract event 0xbb4663e939ed98b60f591a4f61e38a306d703e68557951430c51a3cc0eb4f3fa.
//
// Solidity: event InfoSet(string key, string value)
func (_Hospital *HospitalFilterer) WatchInfoSet(opts *bind.WatchOpts, sink chan<- *HospitalInfoSet) (event.Subscription, error) {

	logs, sub, err := _Hospital.contract.WatchLogs(opts, "InfoSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HospitalInfoSet)
				if err := _Hospital.contract.UnpackLog(event, "InfoSet", log); err != nil {
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

// ParseInfoSet is a log parse operation binding the contract event 0xbb4663e939ed98b60f591a4f61e38a306d703e68557951430c51a3cc0eb4f3fa.
//
// Solidity: event InfoSet(string key, string value)
func (_Hospital *HospitalFilterer) ParseInfoSet(log types.Log) (*HospitalInfoSet, error) {
	event := new(HospitalInfoSet)
	if err := _Hospital.contract.UnpackLog(event, "InfoSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
