// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// HashTimeLockABI is the input ABI used to generate the binding from.
const HashTimeLockABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"swapRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"secret\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"secret\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"returnToSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expireHeight\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"NewSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"secret\",\"type\":\"bytes\"}],\"name\":\"SuccessSwap\",\"type\":\"event\"}]"

// HashTimeLockBin is the compiled bytecode used for deploying new contracts.
var HashTimeLockBin = "0x608060405234801561001057600080fd5b50610d6a806100206000396000f3fe60806040526004361061003f5760003560e01c80633e029427146100445780635fc7e133146101735780637f1ce9eb14610245578063a80de0e814610280575b600080fd5b34801561005057600080fd5b5061007d6004803603602081101561006757600080fd5b81019080803590602001909291905050506102d8565b60405180878152602001868152602001806020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825286818151815260200191508051906020019080838360005b83811015610133578082015181840152602081019050610118565b50505050905090810190601f1680156101605780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561017f57600080fd5b506102436004803603604081101561019657600080fd5b8101908080359060200190929190803590602001906401000000008111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460018302840111640100000000831117156101f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506103ec565b005b34801561025157600080fd5b5061027e6004803603602081101561026857600080fd5b810190808035906020019092919050505061064b565b005b6102d66004803603606081101561029657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506107da565b005b6000602052806000526040600020600091509050806000015490806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103905780601f1061036557610100808354040283529160200191610390565b820191906000526020600020905b81548152906001019060200180831161037357829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905086565b60008060008481526020019081526020016000206000015411610477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f73776170206973206f766572000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008083815260200190815260200160002060010154431115610502576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f74696d656f7574206973206e6f74206f7665722079657400000000000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc600080858152602001908152602001600020600001549081150290604051600060405180830381858888f1935050505015801561055d573d6000803e3d6000fd5b50806000808481526020019081526020016000206002019080519060200190610587929190610c10565b506000806000848152602001908152602001600020600001819055507fa2c430bbccac8c84ea0c055b0b3d0c9d56bb4a1aa8699f183d0f814ec8417a3f82826040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561060c5780820151818401526020810190506105f1565b50505050905090810190601f1680156106395780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b600080600083815260200190815260200160002060000154116106d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f73776170206973206f766572000000000000000000000000000000000000000081525060200191505060405180910390fd5b600080828152602001908152602001600020600101544311610760576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f74696d656f7574206973206e6f74206f7665722079657400000000000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc600080848152602001908152602001600020600001549081150290604051600060405180830381858888f193505050501580156107bb573d6000803e3d6000fd5b5060008060008381526020019081526020016000206000018190555050565b600030338585604051602001808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200194505050505060405160208183030381529060405280519060200120905060003411610920576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f76616c7565206c657373206f7220657175616c7320300000000000000000000081525060200191505060405180910390fd5b60008211610996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f65787069726520686569676874206c657373206f7220657175616c732030000081525060200191505060405180910390fd5b60008060008381526020019081526020016000206001015414610a21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f737761702065786973740000000000000000000000000000000000000000000081525060200191505060405180910390fd5b60606040518060c001604052803481526020018481526020018281526020018581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681525060008084815260200190815260200160002060008201518160000155602082015181600101556040820151816002019080519060200190610abe929190610c90565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507fe27fd08f602ea73841cdb9720ab29ad0d8806779aad9a248110b56e1e215033a8284873388604051808681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019550505050505060405180910390a15050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c5157805160ff1916838001178555610c7f565b82800160010185558215610c7f579182015b82811115610c7e578251825591602001919060010190610c63565b5b509050610c8c9190610d10565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cd157805160ff1916838001178555610cff565b82800160010185558215610cff579182015b82811115610cfe578251825591602001919060010190610ce3565b5b509050610d0c9190610d10565b5090565b610d3291905b80821115610d2e576000816000905550600101610d16565b5090565b9056fea265627a7a723158206e6887e8e776f9a9a947815e6dab656a2f9db99438a770c55955478f61e7640964736f6c634300050b0032"

// DeployHashTimeLock deploys a new Ethereum contract, binding an instance of HashTimeLock to it.
func DeployHashTimeLock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashTimeLock, error) {
	parsed, err := abi.JSON(strings.NewReader(HashTimeLockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashTimeLockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashTimeLock{HashTimeLockCaller: HashTimeLockCaller{contract: contract}, HashTimeLockTransactor: HashTimeLockTransactor{contract: contract}, HashTimeLockFilterer: HashTimeLockFilterer{contract: contract}}, nil
}

// HashTimeLock is an auto generated Go binding around an Ethereum contract.
type HashTimeLock struct {
	HashTimeLockCaller     // Read-only binding to the contract
	HashTimeLockTransactor // Write-only binding to the contract
	HashTimeLockFilterer   // Log filterer for contract events
}

// HashTimeLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashTimeLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTimeLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashTimeLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTimeLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashTimeLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTimeLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashTimeLockSession struct {
	Contract     *HashTimeLock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashTimeLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashTimeLockCallerSession struct {
	Contract *HashTimeLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HashTimeLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashTimeLockTransactorSession struct {
	Contract     *HashTimeLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HashTimeLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashTimeLockRaw struct {
	Contract *HashTimeLock // Generic contract binding to access the raw methods on
}

// HashTimeLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashTimeLockCallerRaw struct {
	Contract *HashTimeLockCaller // Generic read-only contract binding to access the raw methods on
}

// HashTimeLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashTimeLockTransactorRaw struct {
	Contract *HashTimeLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashTimeLock creates a new instance of HashTimeLock, bound to a specific deployed contract.
func NewHashTimeLock(address common.Address, backend bind.ContractBackend) (*HashTimeLock, error) {
	contract, err := bindHashTimeLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashTimeLock{HashTimeLockCaller: HashTimeLockCaller{contract: contract}, HashTimeLockTransactor: HashTimeLockTransactor{contract: contract}, HashTimeLockFilterer: HashTimeLockFilterer{contract: contract}}, nil
}

// NewHashTimeLockCaller creates a new read-only instance of HashTimeLock, bound to a specific deployed contract.
func NewHashTimeLockCaller(address common.Address, caller bind.ContractCaller) (*HashTimeLockCaller, error) {
	contract, err := bindHashTimeLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashTimeLockCaller{contract: contract}, nil
}

// NewHashTimeLockTransactor creates a new write-only instance of HashTimeLock, bound to a specific deployed contract.
func NewHashTimeLockTransactor(address common.Address, transactor bind.ContractTransactor) (*HashTimeLockTransactor, error) {
	contract, err := bindHashTimeLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashTimeLockTransactor{contract: contract}, nil
}

// NewHashTimeLockFilterer creates a new log filterer instance of HashTimeLock, bound to a specific deployed contract.
func NewHashTimeLockFilterer(address common.Address, filterer bind.ContractFilterer) (*HashTimeLockFilterer, error) {
	contract, err := bindHashTimeLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashTimeLockFilterer{contract: contract}, nil
}

// bindHashTimeLock binds a generic wrapper to an already deployed contract.
func bindHashTimeLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashTimeLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashTimeLock *HashTimeLockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashTimeLock.Contract.HashTimeLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashTimeLock *HashTimeLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashTimeLock.Contract.HashTimeLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashTimeLock *HashTimeLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashTimeLock.Contract.HashTimeLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashTimeLock *HashTimeLockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashTimeLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashTimeLock *HashTimeLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashTimeLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashTimeLock *HashTimeLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashTimeLock.Contract.contract.Transact(opts, method, params...)
}

// SwapRequests is a free data retrieval call binding the contract method 0x3e029427.
//
// Solidity: function swapRequests(bytes32 ) constant returns(uint256 amount, uint256 expireHeight, bytes secret, bytes32 secretHash, address sender, address recipient)
func (_HashTimeLock *HashTimeLockCaller) SwapRequests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount       *big.Int
	ExpireHeight *big.Int
	Secret       []byte
	SecretHash   [32]byte
	Sender       common.Address
	Recipient    common.Address
}, error) {
	ret := new(struct {
		Amount       *big.Int
		ExpireHeight *big.Int
		Secret       []byte
		SecretHash   [32]byte
		Sender       common.Address
		Recipient    common.Address
	})
	out := ret
	err := _HashTimeLock.contract.Call(opts, out, "swapRequests", arg0)
	return *ret, err
}

// SwapRequests is a free data retrieval call binding the contract method 0x3e029427.
//
// Solidity: function swapRequests(bytes32 ) constant returns(uint256 amount, uint256 expireHeight, bytes secret, bytes32 secretHash, address sender, address recipient)
func (_HashTimeLock *HashTimeLockSession) SwapRequests(arg0 [32]byte) (struct {
	Amount       *big.Int
	ExpireHeight *big.Int
	Secret       []byte
	SecretHash   [32]byte
	Sender       common.Address
	Recipient    common.Address
}, error) {
	return _HashTimeLock.Contract.SwapRequests(&_HashTimeLock.CallOpts, arg0)
}

// SwapRequests is a free data retrieval call binding the contract method 0x3e029427.
//
// Solidity: function swapRequests(bytes32 ) constant returns(uint256 amount, uint256 expireHeight, bytes secret, bytes32 secretHash, address sender, address recipient)
func (_HashTimeLock *HashTimeLockCallerSession) SwapRequests(arg0 [32]byte) (struct {
	Amount       *big.Int
	ExpireHeight *big.Int
	Secret       []byte
	SecretHash   [32]byte
	Sender       common.Address
	Recipient    common.Address
}, error) {
	return _HashTimeLock.Contract.SwapRequests(&_HashTimeLock.CallOpts, arg0)
}

// Lock is a paid mutator transaction binding the contract method 0xa80de0e8.
//
// Solidity: function lock(address recipient, bytes32 secretHash, uint256 expireHeight) returns()
func (_HashTimeLock *HashTimeLockTransactor) Lock(opts *bind.TransactOpts, recipient common.Address, secretHash [32]byte, expireHeight *big.Int) (*types.Transaction, error) {
	return _HashTimeLock.contract.Transact(opts, "lock", recipient, secretHash, expireHeight)
}

// Lock is a paid mutator transaction binding the contract method 0xa80de0e8.
//
// Solidity: function lock(address recipient, bytes32 secretHash, uint256 expireHeight) returns()
func (_HashTimeLock *HashTimeLockSession) Lock(recipient common.Address, secretHash [32]byte, expireHeight *big.Int) (*types.Transaction, error) {
	return _HashTimeLock.Contract.Lock(&_HashTimeLock.TransactOpts, recipient, secretHash, expireHeight)
}

// Lock is a paid mutator transaction binding the contract method 0xa80de0e8.
//
// Solidity: function lock(address recipient, bytes32 secretHash, uint256 expireHeight) returns()
func (_HashTimeLock *HashTimeLockTransactorSession) Lock(recipient common.Address, secretHash [32]byte, expireHeight *big.Int) (*types.Transaction, error) {
	return _HashTimeLock.Contract.Lock(&_HashTimeLock.TransactOpts, recipient, secretHash, expireHeight)
}

// ReturnToSender is a paid mutator transaction binding the contract method 0x7f1ce9eb.
//
// Solidity: function returnToSender(bytes32 requestHash) returns()
func (_HashTimeLock *HashTimeLockTransactor) ReturnToSender(opts *bind.TransactOpts, requestHash [32]byte) (*types.Transaction, error) {
	return _HashTimeLock.contract.Transact(opts, "returnToSender", requestHash)
}

// ReturnToSender is a paid mutator transaction binding the contract method 0x7f1ce9eb.
//
// Solidity: function returnToSender(bytes32 requestHash) returns()
func (_HashTimeLock *HashTimeLockSession) ReturnToSender(requestHash [32]byte) (*types.Transaction, error) {
	return _HashTimeLock.Contract.ReturnToSender(&_HashTimeLock.TransactOpts, requestHash)
}

// ReturnToSender is a paid mutator transaction binding the contract method 0x7f1ce9eb.
//
// Solidity: function returnToSender(bytes32 requestHash) returns()
func (_HashTimeLock *HashTimeLockTransactorSession) ReturnToSender(requestHash [32]byte) (*types.Transaction, error) {
	return _HashTimeLock.Contract.ReturnToSender(&_HashTimeLock.TransactOpts, requestHash)
}

// Unlock is a paid mutator transaction binding the contract method 0x5fc7e133.
//
// Solidity: function unlock(bytes32 requestHash, bytes secret) returns()
func (_HashTimeLock *HashTimeLockTransactor) Unlock(opts *bind.TransactOpts, requestHash [32]byte, secret []byte) (*types.Transaction, error) {
	return _HashTimeLock.contract.Transact(opts, "unlock", requestHash, secret)
}

// Unlock is a paid mutator transaction binding the contract method 0x5fc7e133.
//
// Solidity: function unlock(bytes32 requestHash, bytes secret) returns()
func (_HashTimeLock *HashTimeLockSession) Unlock(requestHash [32]byte, secret []byte) (*types.Transaction, error) {
	return _HashTimeLock.Contract.Unlock(&_HashTimeLock.TransactOpts, requestHash, secret)
}

// Unlock is a paid mutator transaction binding the contract method 0x5fc7e133.
//
// Solidity: function unlock(bytes32 requestHash, bytes secret) returns()
func (_HashTimeLock *HashTimeLockTransactorSession) Unlock(requestHash [32]byte, secret []byte) (*types.Transaction, error) {
	return _HashTimeLock.Contract.Unlock(&_HashTimeLock.TransactOpts, requestHash, secret)
}

// HashTimeLockNewSwapIterator is returned from FilterNewSwap and is used to iterate over the raw logs and unpacked data for NewSwap events raised by the HashTimeLock contract.
type HashTimeLockNewSwapIterator struct {
	Event *HashTimeLockNewSwap // Event containing the contract specifics and raw log

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
func (it *HashTimeLockNewSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashTimeLockNewSwap)
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
		it.Event = new(HashTimeLockNewSwap)
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
func (it *HashTimeLockNewSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashTimeLockNewSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashTimeLockNewSwap represents a NewSwap event raised by the HashTimeLock contract.
type HashTimeLockNewSwap struct {
	RequestHash  [32]byte
	ExpireHeight *big.Int
	Recipient    common.Address
	Sender       common.Address
	SecretHash   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSwap is a free log retrieval operation binding the contract event 0xe27fd08f602ea73841cdb9720ab29ad0d8806779aad9a248110b56e1e215033a.
//
// Solidity: event NewSwap(bytes32 requestHash, uint256 expireHeight, address recipient, address sender, bytes32 secretHash)
func (_HashTimeLock *HashTimeLockFilterer) FilterNewSwap(opts *bind.FilterOpts) (*HashTimeLockNewSwapIterator, error) {

	logs, sub, err := _HashTimeLock.contract.FilterLogs(opts, "NewSwap")
	if err != nil {
		return nil, err
	}
	return &HashTimeLockNewSwapIterator{contract: _HashTimeLock.contract, event: "NewSwap", logs: logs, sub: sub}, nil
}

// WatchNewSwap is a free log subscription operation binding the contract event 0xe27fd08f602ea73841cdb9720ab29ad0d8806779aad9a248110b56e1e215033a.
//
// Solidity: event NewSwap(bytes32 requestHash, uint256 expireHeight, address recipient, address sender, bytes32 secretHash)
func (_HashTimeLock *HashTimeLockFilterer) WatchNewSwap(opts *bind.WatchOpts, sink chan<- *HashTimeLockNewSwap) (event.Subscription, error) {

	logs, sub, err := _HashTimeLock.contract.WatchLogs(opts, "NewSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashTimeLockNewSwap)
				if err := _HashTimeLock.contract.UnpackLog(event, "NewSwap", log); err != nil {
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

// ParseNewSwap is a log parse operation binding the contract event 0xe27fd08f602ea73841cdb9720ab29ad0d8806779aad9a248110b56e1e215033a.
//
// Solidity: event NewSwap(bytes32 requestHash, uint256 expireHeight, address recipient, address sender, bytes32 secretHash)
func (_HashTimeLock *HashTimeLockFilterer) ParseNewSwap(log types.Log) (*HashTimeLockNewSwap, error) {
	event := new(HashTimeLockNewSwap)
	if err := _HashTimeLock.contract.UnpackLog(event, "NewSwap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HashTimeLockSuccessSwapIterator is returned from FilterSuccessSwap and is used to iterate over the raw logs and unpacked data for SuccessSwap events raised by the HashTimeLock contract.
type HashTimeLockSuccessSwapIterator struct {
	Event *HashTimeLockSuccessSwap // Event containing the contract specifics and raw log

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
func (it *HashTimeLockSuccessSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashTimeLockSuccessSwap)
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
		it.Event = new(HashTimeLockSuccessSwap)
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
func (it *HashTimeLockSuccessSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashTimeLockSuccessSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashTimeLockSuccessSwap represents a SuccessSwap event raised by the HashTimeLock contract.
type HashTimeLockSuccessSwap struct {
	RequestHash [32]byte
	Secret      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSuccessSwap is a free log retrieval operation binding the contract event 0xa2c430bbccac8c84ea0c055b0b3d0c9d56bb4a1aa8699f183d0f814ec8417a3f.
//
// Solidity: event SuccessSwap(bytes32 requestHash, bytes secret)
func (_HashTimeLock *HashTimeLockFilterer) FilterSuccessSwap(opts *bind.FilterOpts) (*HashTimeLockSuccessSwapIterator, error) {

	logs, sub, err := _HashTimeLock.contract.FilterLogs(opts, "SuccessSwap")
	if err != nil {
		return nil, err
	}
	return &HashTimeLockSuccessSwapIterator{contract: _HashTimeLock.contract, event: "SuccessSwap", logs: logs, sub: sub}, nil
}

// WatchSuccessSwap is a free log subscription operation binding the contract event 0xa2c430bbccac8c84ea0c055b0b3d0c9d56bb4a1aa8699f183d0f814ec8417a3f.
//
// Solidity: event SuccessSwap(bytes32 requestHash, bytes secret)
func (_HashTimeLock *HashTimeLockFilterer) WatchSuccessSwap(opts *bind.WatchOpts, sink chan<- *HashTimeLockSuccessSwap) (event.Subscription, error) {

	logs, sub, err := _HashTimeLock.contract.WatchLogs(opts, "SuccessSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashTimeLockSuccessSwap)
				if err := _HashTimeLock.contract.UnpackLog(event, "SuccessSwap", log); err != nil {
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

// ParseSuccessSwap is a log parse operation binding the contract event 0xa2c430bbccac8c84ea0c055b0b3d0c9d56bb4a1aa8699f183d0f814ec8417a3f.
//
// Solidity: event SuccessSwap(bytes32 requestHash, bytes secret)
func (_HashTimeLock *HashTimeLockFilterer) ParseSuccessSwap(log types.Log) (*HashTimeLockSuccessSwap, error) {
	event := new(HashTimeLockSuccessSwap)
	if err := _HashTimeLock.contract.UnpackLog(event, "SuccessSwap", log); err != nil {
		return nil, err
	}
	return event, nil
}
