// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_contarct

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
)

// EthDemoMetaData contains all meta data concerning the EthDemo contract.
var EthDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200092f3803806200092f833981810160405281019062000037919062000185565b80600090805190602001906200004f92919062000057565b50506200035a565b82805462000065906200026b565b90600052602060002090601f016020900481019282620000895760008555620000d5565b82601f10620000a457805160ff1916838001178555620000d5565b82800160010185558215620000d5579182015b82811115620000d4578251825591602001919060010190620000b7565b5b509050620000e49190620000e8565b5090565b5b8082111562000103576000816000905550600101620000e9565b5090565b60006200011e6200011884620001ff565b620001d6565b9050828152602081018484840111156200013d576200013c6200033a565b5b6200014a84828562000235565b509392505050565b600082601f8301126200016a576200016962000335565b5b81516200017c84826020860162000107565b91505092915050565b6000602082840312156200019e576200019d62000344565b5b600082015167ffffffffffffffff811115620001bf57620001be6200033f565b5b620001cd8482850162000152565b91505092915050565b6000620001e2620001f5565b9050620001f08282620002a1565b919050565b6000604051905090565b600067ffffffffffffffff8211156200021d576200021c62000306565b5b620002288262000349565b9050602081019050919050565b60005b838110156200025557808201518184015260208101905062000238565b8381111562000265576000848401525b50505050565b600060028204905060018216806200028457607f821691505b602082108114156200029b576200029a620002d7565b5b50919050565b620002ac8262000349565b810181811067ffffffffffffffff82111715620002ce57620002cd62000306565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6105c5806200036a6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063209652551461003b57806393a0935214610059575b600080fd5b610043610075565b6040516100509190610387565b60405180910390f35b610073600480360381019061006e9190610285565b610107565b005b606060008054610084906104a9565b80601f01602080910402602001604051908101604052809291908181526020018280546100b0906104a9565b80156100fd5780601f106100d2576101008083540402835291602001916100fd565b820191906000526020600020905b8154815290600101906020018083116100e057829003601f168201915b5050505050905090565b806000908051906020019061011d929190610172565b503373ffffffffffffffffffffffffffffffffffffffff167fe826f71647b8486f2bae59832124c70792fba044036720a54ec8dacdd5df4fcb6000836040516101679291906103a9565b60405180910390a250565b82805461017e906104a9565b90600052602060002090601f0160209004810192826101a057600085556101e7565b82601f106101b957805160ff19168380011785556101e7565b828001600101855582156101e7579182015b828111156101e65782518255916020019190600101906101cb565b5b5090506101f491906101f8565b5090565b5b808211156102115760008160009055506001016101f9565b5090565b600061022861022384610405565b6103e0565b9050828152602081018484840111156102445761024361056f565b5b61024f848285610467565b509392505050565b600082601f83011261026c5761026b61056a565b5b813561027c848260208601610215565b91505092915050565b60006020828403121561029b5761029a610579565b5b600082013567ffffffffffffffff8111156102b9576102b8610574565b5b6102c584828501610257565b91505092915050565b60006102d98261044b565b6102e38185610456565b93506102f3818560208601610476565b6102fc8161057e565b840191505092915050565b60008154610314816104a9565b61031e8186610456565b94506001821660008114610339576001811461034b5761037e565b60ff198316865260208601935061037e565b61035485610436565b60005b8381101561037657815481890152600182019150602081019050610357565b808801955050505b50505092915050565b600060208201905081810360008301526103a181846102ce565b905092915050565b600060408201905081810360008301526103c38185610307565b905081810360208301526103d781846102ce565b90509392505050565b60006103ea6103fb565b90506103f682826104db565b919050565b6000604051905090565b600067ffffffffffffffff8211156104205761041f61053b565b5b6104298261057e565b9050602081019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600082825260208201905092915050565b82818337600083830152505050565b60005b83811015610494578082015181840152602081019050610479565b838111156104a3576000848401525b50505050565b600060028204905060018216806104c157607f821691505b602082108114156104d5576104d461050c565b5b50919050565b6104e48261057e565b810181811067ffffffffffffffff821117156105035761050261053b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f830116905091905056fea2646970667358221220d8699ad881a2dca470539d50ba2080293aef35c9cbfc9f5c60948ecbce14e40664736f6c63430008060033",
}

// EthDemoABI is the input ABI used to generate the binding from.
// Deprecated: Use EthDemoMetaData.ABI instead.
var EthDemoABI = EthDemoMetaData.ABI

// EthDemoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthDemoMetaData.Bin instead.
var EthDemoBin = EthDemoMetaData.Bin

// DeployEthDemo deploys a new Ethereum contract, binding an instance of EthDemo to it.
func DeployEthDemo(auth *bind.TransactOpts, backend bind.ContractBackend, value string) (common.Address, *types.Transaction, *EthDemo, error) {
	parsed, err := EthDemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthDemoBin), backend, value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthDemo{EthDemoCaller: EthDemoCaller{contract: contract}, EthDemoTransactor: EthDemoTransactor{contract: contract}, EthDemoFilterer: EthDemoFilterer{contract: contract}}, nil
}

// EthDemo is an auto generated Go binding around an Ethereum contract.
type EthDemo struct {
	EthDemoCaller     // Read-only binding to the contract
	EthDemoTransactor // Write-only binding to the contract
	EthDemoFilterer   // Log filterer for contract events
}

// EthDemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthDemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthDemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthDemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthDemoSession struct {
	Contract     *EthDemo          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthDemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthDemoCallerSession struct {
	Contract *EthDemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EthDemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthDemoTransactorSession struct {
	Contract     *EthDemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthDemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthDemoRaw struct {
	Contract *EthDemo // Generic contract binding to access the raw methods on
}

// EthDemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthDemoCallerRaw struct {
	Contract *EthDemoCaller // Generic read-only contract binding to access the raw methods on
}

// EthDemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthDemoTransactorRaw struct {
	Contract *EthDemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthDemo creates a new instance of EthDemo, bound to a specific deployed contract.
func NewEthDemo(address common.Address, backend bind.ContractBackend) (*EthDemo, error) {
	contract, err := bindEthDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthDemo{EthDemoCaller: EthDemoCaller{contract: contract}, EthDemoTransactor: EthDemoTransactor{contract: contract}, EthDemoFilterer: EthDemoFilterer{contract: contract}}, nil
}

// NewEthDemoCaller creates a new read-only instance of EthDemo, bound to a specific deployed contract.
func NewEthDemoCaller(address common.Address, caller bind.ContractCaller) (*EthDemoCaller, error) {
	contract, err := bindEthDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthDemoCaller{contract: contract}, nil
}

// NewEthDemoTransactor creates a new write-only instance of EthDemo, bound to a specific deployed contract.
func NewEthDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*EthDemoTransactor, error) {
	contract, err := bindEthDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthDemoTransactor{contract: contract}, nil
}

// NewEthDemoFilterer creates a new log filterer instance of EthDemo, bound to a specific deployed contract.
func NewEthDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*EthDemoFilterer, error) {
	contract, err := bindEthDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthDemoFilterer{contract: contract}, nil
}

// bindEthDemo binds a generic wrapper to an already deployed contract.
func bindEthDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthDemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDemo *EthDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDemo.Contract.EthDemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDemo *EthDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDemo.Contract.EthDemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDemo *EthDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDemo.Contract.EthDemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDemo *EthDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDemo *EthDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDemo *EthDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDemo.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(string)
func (_EthDemo *EthDemoCaller) GetValue(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthDemo.contract.Call(opts, &out, "getValue")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(string)
func (_EthDemo *EthDemoSession) GetValue() (string, error) {
	return _EthDemo.Contract.GetValue(&_EthDemo.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(string)
func (_EthDemo *EthDemoCallerSession) GetValue() (string, error) {
	return _EthDemo.Contract.GetValue(&_EthDemo.CallOpts)
}

// SetValue is a paid mutator transaction binding the contract method 0x93a09352.
//
// Solidity: function setValue(string value) returns()
func (_EthDemo *EthDemoTransactor) SetValue(opts *bind.TransactOpts, value string) (*types.Transaction, error) {
	return _EthDemo.contract.Transact(opts, "setValue", value)
}

// SetValue is a paid mutator transaction binding the contract method 0x93a09352.
//
// Solidity: function setValue(string value) returns()
func (_EthDemo *EthDemoSession) SetValue(value string) (*types.Transaction, error) {
	return _EthDemo.Contract.SetValue(&_EthDemo.TransactOpts, value)
}

// SetValue is a paid mutator transaction binding the contract method 0x93a09352.
//
// Solidity: function setValue(string value) returns()
func (_EthDemo *EthDemoTransactorSession) SetValue(value string) (*types.Transaction, error) {
	return _EthDemo.Contract.SetValue(&_EthDemo.TransactOpts, value)
}

// EthDemoValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the EthDemo contract.
type EthDemoValueChangedIterator struct {
	Event *EthDemoValueChanged // Event containing the contract specifics and raw log

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
func (it *EthDemoValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDemoValueChanged)
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
		it.Event = new(EthDemoValueChanged)
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
func (it *EthDemoValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDemoValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDemoValueChanged represents a ValueChanged event raised by the EthDemo contract.
type EthDemoValueChanged struct {
	Owner    common.Address
	OldValue string
	NewValue string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0xe826f71647b8486f2bae59832124c70792fba044036720a54ec8dacdd5df4fcb.
//
// Solidity: event ValueChanged(address indexed owner, string oldValue, string newValue)
func (_EthDemo *EthDemoFilterer) FilterValueChanged(opts *bind.FilterOpts, owner []common.Address) (*EthDemoValueChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EthDemo.contract.FilterLogs(opts, "ValueChanged", ownerRule)
	if err != nil {
		return nil, err
	}
	return &EthDemoValueChangedIterator{contract: _EthDemo.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0xe826f71647b8486f2bae59832124c70792fba044036720a54ec8dacdd5df4fcb.
//
// Solidity: event ValueChanged(address indexed owner, string oldValue, string newValue)
func (_EthDemo *EthDemoFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *EthDemoValueChanged, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EthDemo.contract.WatchLogs(opts, "ValueChanged", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDemoValueChanged)
				if err := _EthDemo.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0xe826f71647b8486f2bae59832124c70792fba044036720a54ec8dacdd5df4fcb.
//
// Solidity: event ValueChanged(address indexed owner, string oldValue, string newValue)
func (_EthDemo *EthDemoFilterer) ParseValueChanged(log types.Log) (*EthDemoValueChanged, error) {
	event := new(EthDemoValueChanged)
	if err := _EthDemo.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
