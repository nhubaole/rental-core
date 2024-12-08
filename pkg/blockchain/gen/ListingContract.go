// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

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

// ListingContractMetaData contains all meta data concerning the ListingContract contract.
var ListingContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"owner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRent\",\"type\":\"bool\"}],\"name\":\"RoomCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"RoomDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"owner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRent\",\"type\":\"bool\"}],\"name\":\"RoomUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_owner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isRent\",\"type\":\"bool\"}],\"name\":\"createRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getRoom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rooms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"owner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRent\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isRent\",\"type\":\"bool\"}],\"name\":\"updateRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506109cc8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80630f9f583b146100595780631bae0ac814610075578063682b7100146100ad5780636d8a74cb146100c95780636e04e16e14610100575b5f5ffd5b610073600480360381019061006e91906105ee565b61011c565b005b61008f600480360381019061008a9190610665565b610212565b6040516100a4999897969594939291906106ae565b60405180910390f35b6100c760048036038101906100c29190610665565b610273565b005b6100e360048036038101906100de9190610665565b610333565b6040516100f7989796959493929190610739565b60405180910390f35b61011a600480360381019061011591906107b5565b6103f4565b005b5f5f5f8781526020019081526020015f20905060011515816008015f9054906101000a900460ff16151514610186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017d90610898565b60405180910390fd5b84816004018190555083816005018190555082816002018190555081816003015f6101000a81548160ff0219169083151502179055504281600701819055507f534c7676f2d5efde7c224ea38906c043ed4c7d44482023b66a5403806b5ce65f86826001015487878787604051610202969594939291906108b6565b60405180910390a1505050505050565b5f602052805f5260405f205f91509050805f015490806001015490806002015490806003015f9054906101000a900460ff1690806004015490806005015490806006015490806007015490806008015f9054906101000a900460ff16905089565b5f5f5f8381526020019081526020015f20905060011515816008015f9054906101000a900460ff161515146102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d490610898565b60405180910390fd5b5f816008015f6101000a81548160ff0219169083151502179055507f640a818bdd6ca14bdd2cf9cdac867db41b7a8832e5c24995d200910d60e7e4dc826040516103279190610915565b60405180910390a15050565b5f5f5f5f5f5f5f5f5f5f5f8b81526020019081526020015f20905060011515816008015f9054906101000a900460ff161515146103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c90610898565b60405180910390fd5b805f01548160010154826004015483600501548460020154856003015f9054906101000a900460ff16866006015487600701549850985098509850985098509850985050919395975091939597565b5f5f5f8881526020019081526020015f2090505f1515816008015f9054906101000a900460ff1615151461045d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045490610978565b60405180910390fd5b6040518061012001604052808881526020018781526020018481526020018315158152602001868152602001858152602001428152602001428152602001600115158152505f5f8981526020019081526020015f205f820151815f015560208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155610100820151816008015f6101000a81548160ff0219169083151502179055509050507f8ca55b2f4ddd3294467a47bcc1601b06616452e4fdb1c22169423a437392d032878787878787604051610571969594939291906108b6565b60405180910390a150505050505050565b5f5ffd5b5f819050919050565b61059881610586565b81146105a2575f5ffd5b50565b5f813590506105b38161058f565b92915050565b5f8115159050919050565b6105cd816105b9565b81146105d7575f5ffd5b50565b5f813590506105e8816105c4565b92915050565b5f5f5f5f5f60a0868803121561060757610606610582565b5b5f610614888289016105a5565b9550506020610625888289016105a5565b9450506040610636888289016105a5565b9350506060610647888289016105a5565b9250506080610658888289016105da565b9150509295509295909350565b5f6020828403121561067a57610679610582565b5b5f610687848285016105a5565b91505092915050565b61069981610586565b82525050565b6106a8816105b9565b82525050565b5f610120820190506106c25f83018c610690565b6106cf602083018b610690565b6106dc604083018a610690565b6106e9606083018961069f565b6106f66080830188610690565b61070360a0830187610690565b61071060c0830186610690565b61071d60e0830185610690565b61072b61010083018461069f565b9a9950505050505050505050565b5f6101008201905061074d5f83018b610690565b61075a602083018a610690565b6107676040830189610690565b6107746060830188610690565b6107816080830187610690565b61078e60a083018661069f565b61079b60c0830185610690565b6107a860e0830184610690565b9998505050505050505050565b5f5f5f5f5f5f60c087890312156107cf576107ce610582565b5b5f6107dc89828a016105a5565b96505060206107ed89828a016105a5565b95505060406107fe89828a016105a5565b945050606061080f89828a016105a5565b935050608061082089828a016105a5565b92505060a061083189828a016105da565b9150509295509295509295565b5f82825260208201905092915050565b7f526f6f6d20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f61088260138361083e565b915061088d8261084e565b602082019050919050565b5f6020820190508181035f8301526108af81610876565b9050919050565b5f60c0820190506108c95f830189610690565b6108d66020830188610690565b6108e36040830187610690565b6108f06060830186610690565b6108fd6080830185610690565b61090a60a083018461069f565b979650505050505050565b5f6020820190506109285f830184610690565b92915050565b7f526f6f6d20616c726561647920657869737473000000000000000000000000005f82015250565b5f61096260138361083e565b915061096d8261092e565b602082019050919050565b5f6020820190508181035f83015261098f81610956565b905091905056fea264697066735822122009f0f77cbe90f6ffe75f0358b2e2f83b781b493d6e99e3ccb762d94af3e8616164736f6c634300081c0033",
}

// ListingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ListingContractMetaData.ABI instead.
var ListingContractABI = ListingContractMetaData.ABI

// ListingContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ListingContractMetaData.Bin instead.
var ListingContractBin = ListingContractMetaData.Bin

// DeployListingContract deploys a new Ethereum contract, binding an instance of ListingContract to it.
func DeployListingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ListingContract, error) {
	parsed, err := ListingContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ListingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ListingContract{ListingContractCaller: ListingContractCaller{contract: contract}, ListingContractTransactor: ListingContractTransactor{contract: contract}, ListingContractFilterer: ListingContractFilterer{contract: contract}}, nil
}

// ListingContract is an auto generated Go binding around an Ethereum contract.
type ListingContract struct {
	ListingContractCaller     // Read-only binding to the contract
	ListingContractTransactor // Write-only binding to the contract
	ListingContractFilterer   // Log filterer for contract events
}

// ListingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListingContractSession struct {
	Contract     *ListingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListingContractCallerSession struct {
	Contract *ListingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ListingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListingContractTransactorSession struct {
	Contract     *ListingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ListingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListingContractRaw struct {
	Contract *ListingContract // Generic contract binding to access the raw methods on
}

// ListingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListingContractCallerRaw struct {
	Contract *ListingContractCaller // Generic read-only contract binding to access the raw methods on
}

// ListingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListingContractTransactorRaw struct {
	Contract *ListingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListingContract creates a new instance of ListingContract, bound to a specific deployed contract.
func NewListingContract(address common.Address, backend bind.ContractBackend) (*ListingContract, error) {
	contract, err := bindListingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ListingContract{ListingContractCaller: ListingContractCaller{contract: contract}, ListingContractTransactor: ListingContractTransactor{contract: contract}, ListingContractFilterer: ListingContractFilterer{contract: contract}}, nil
}

// NewListingContractCaller creates a new read-only instance of ListingContract, bound to a specific deployed contract.
func NewListingContractCaller(address common.Address, caller bind.ContractCaller) (*ListingContractCaller, error) {
	contract, err := bindListingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListingContractCaller{contract: contract}, nil
}

// NewListingContractTransactor creates a new write-only instance of ListingContract, bound to a specific deployed contract.
func NewListingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ListingContractTransactor, error) {
	contract, err := bindListingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListingContractTransactor{contract: contract}, nil
}

// NewListingContractFilterer creates a new log filterer instance of ListingContract, bound to a specific deployed contract.
func NewListingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ListingContractFilterer, error) {
	contract, err := bindListingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListingContractFilterer{contract: contract}, nil
}

// bindListingContract binds a generic wrapper to an already deployed contract.
func bindListingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ListingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListingContract *ListingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListingContract.Contract.ListingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListingContract *ListingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListingContract.Contract.ListingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListingContract *ListingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListingContract.Contract.ListingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListingContract *ListingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListingContract *ListingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListingContract *ListingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListingContract.Contract.contract.Transact(opts, method, params...)
}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 _id) view returns(uint256, uint256, uint256, uint256, uint256, bool, uint256, uint256)
func (_ListingContract *ListingContractCaller) GetRoom(opts *bind.CallOpts, _id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ListingContract.contract.Call(opts, &out, "getRoom", _id)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 _id) view returns(uint256, uint256, uint256, uint256, uint256, bool, uint256, uint256)
func (_ListingContract *ListingContractSession) GetRoom(_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	return _ListingContract.Contract.GetRoom(&_ListingContract.CallOpts, _id)
}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 _id) view returns(uint256, uint256, uint256, uint256, uint256, bool, uint256, uint256)
func (_ListingContract *ListingContractCallerSession) GetRoom(_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	return _ListingContract.Contract.GetRoom(&_ListingContract.CallOpts, _id)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, uint256 owner, uint256 status, bool isRent, uint256 totalPrice, uint256 deposit, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ListingContract *ListingContractCaller) Rooms(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id         *big.Int
	Owner      *big.Int
	Status     *big.Int
	IsRent     bool
	TotalPrice *big.Int
	Deposit    *big.Int
	CreatedAt  *big.Int
	UpdatedAt  *big.Int
	Exists     bool
}, error) {
	var out []interface{}
	err := _ListingContract.contract.Call(opts, &out, "rooms", arg0)

	outstruct := new(struct {
		Id         *big.Int
		Owner      *big.Int
		Status     *big.Int
		IsRent     bool
		TotalPrice *big.Int
		Deposit    *big.Int
		CreatedAt  *big.Int
		UpdatedAt  *big.Int
		Exists     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsRent = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.TotalPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deposit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, uint256 owner, uint256 status, bool isRent, uint256 totalPrice, uint256 deposit, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ListingContract *ListingContractSession) Rooms(arg0 *big.Int) (struct {
	Id         *big.Int
	Owner      *big.Int
	Status     *big.Int
	IsRent     bool
	TotalPrice *big.Int
	Deposit    *big.Int
	CreatedAt  *big.Int
	UpdatedAt  *big.Int
	Exists     bool
}, error) {
	return _ListingContract.Contract.Rooms(&_ListingContract.CallOpts, arg0)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, uint256 owner, uint256 status, bool isRent, uint256 totalPrice, uint256 deposit, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ListingContract *ListingContractCallerSession) Rooms(arg0 *big.Int) (struct {
	Id         *big.Int
	Owner      *big.Int
	Status     *big.Int
	IsRent     bool
	TotalPrice *big.Int
	Deposit    *big.Int
	CreatedAt  *big.Int
	UpdatedAt  *big.Int
	Exists     bool
}, error) {
	return _ListingContract.Contract.Rooms(&_ListingContract.CallOpts, arg0)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x6e04e16e.
//
// Solidity: function createRoom(uint256 _id, uint256 _owner, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractTransactor) CreateRoom(opts *bind.TransactOpts, _id *big.Int, _owner *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.contract.Transact(opts, "createRoom", _id, _owner, _totalPrice, _deposit, _status, _isRent)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x6e04e16e.
//
// Solidity: function createRoom(uint256 _id, uint256 _owner, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractSession) CreateRoom(_id *big.Int, _owner *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.Contract.CreateRoom(&_ListingContract.TransactOpts, _id, _owner, _totalPrice, _deposit, _status, _isRent)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x6e04e16e.
//
// Solidity: function createRoom(uint256 _id, uint256 _owner, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractTransactorSession) CreateRoom(_id *big.Int, _owner *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.Contract.CreateRoom(&_ListingContract.TransactOpts, _id, _owner, _totalPrice, _deposit, _status, _isRent)
}

// DeleteRoom is a paid mutator transaction binding the contract method 0x682b7100.
//
// Solidity: function deleteRoom(uint256 _id) returns()
func (_ListingContract *ListingContractTransactor) DeleteRoom(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ListingContract.contract.Transact(opts, "deleteRoom", _id)
}

// DeleteRoom is a paid mutator transaction binding the contract method 0x682b7100.
//
// Solidity: function deleteRoom(uint256 _id) returns()
func (_ListingContract *ListingContractSession) DeleteRoom(_id *big.Int) (*types.Transaction, error) {
	return _ListingContract.Contract.DeleteRoom(&_ListingContract.TransactOpts, _id)
}

// DeleteRoom is a paid mutator transaction binding the contract method 0x682b7100.
//
// Solidity: function deleteRoom(uint256 _id) returns()
func (_ListingContract *ListingContractTransactorSession) DeleteRoom(_id *big.Int) (*types.Transaction, error) {
	return _ListingContract.Contract.DeleteRoom(&_ListingContract.TransactOpts, _id)
}

// UpdateRoom is a paid mutator transaction binding the contract method 0x0f9f583b.
//
// Solidity: function updateRoom(uint256 _id, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractTransactor) UpdateRoom(opts *bind.TransactOpts, _id *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.contract.Transact(opts, "updateRoom", _id, _totalPrice, _deposit, _status, _isRent)
}

// UpdateRoom is a paid mutator transaction binding the contract method 0x0f9f583b.
//
// Solidity: function updateRoom(uint256 _id, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractSession) UpdateRoom(_id *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.Contract.UpdateRoom(&_ListingContract.TransactOpts, _id, _totalPrice, _deposit, _status, _isRent)
}

// UpdateRoom is a paid mutator transaction binding the contract method 0x0f9f583b.
//
// Solidity: function updateRoom(uint256 _id, uint256 _totalPrice, uint256 _deposit, uint256 _status, bool _isRent) returns()
func (_ListingContract *ListingContractTransactorSession) UpdateRoom(_id *big.Int, _totalPrice *big.Int, _deposit *big.Int, _status *big.Int, _isRent bool) (*types.Transaction, error) {
	return _ListingContract.Contract.UpdateRoom(&_ListingContract.TransactOpts, _id, _totalPrice, _deposit, _status, _isRent)
}

// ListingContractRoomCreatedIterator is returned from FilterRoomCreated and is used to iterate over the raw logs and unpacked data for RoomCreated events raised by the ListingContract contract.
type ListingContractRoomCreatedIterator struct {
	Event *ListingContractRoomCreated // Event containing the contract specifics and raw log

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
func (it *ListingContractRoomCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingContractRoomCreated)
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
		it.Event = new(ListingContractRoomCreated)
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
func (it *ListingContractRoomCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingContractRoomCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingContractRoomCreated represents a RoomCreated event raised by the ListingContract contract.
type ListingContractRoomCreated struct {
	Id         *big.Int
	Owner      *big.Int
	TotalPrice *big.Int
	Deposit    *big.Int
	Status     *big.Int
	IsRent     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoomCreated is a free log retrieval operation binding the contract event 0x8ca55b2f4ddd3294467a47bcc1601b06616452e4fdb1c22169423a437392d032.
//
// Solidity: event RoomCreated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) FilterRoomCreated(opts *bind.FilterOpts) (*ListingContractRoomCreatedIterator, error) {

	logs, sub, err := _ListingContract.contract.FilterLogs(opts, "RoomCreated")
	if err != nil {
		return nil, err
	}
	return &ListingContractRoomCreatedIterator{contract: _ListingContract.contract, event: "RoomCreated", logs: logs, sub: sub}, nil
}

// WatchRoomCreated is a free log subscription operation binding the contract event 0x8ca55b2f4ddd3294467a47bcc1601b06616452e4fdb1c22169423a437392d032.
//
// Solidity: event RoomCreated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) WatchRoomCreated(opts *bind.WatchOpts, sink chan<- *ListingContractRoomCreated) (event.Subscription, error) {

	logs, sub, err := _ListingContract.contract.WatchLogs(opts, "RoomCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingContractRoomCreated)
				if err := _ListingContract.contract.UnpackLog(event, "RoomCreated", log); err != nil {
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

// ParseRoomCreated is a log parse operation binding the contract event 0x8ca55b2f4ddd3294467a47bcc1601b06616452e4fdb1c22169423a437392d032.
//
// Solidity: event RoomCreated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) ParseRoomCreated(log types.Log) (*ListingContractRoomCreated, error) {
	event := new(ListingContractRoomCreated)
	if err := _ListingContract.contract.UnpackLog(event, "RoomCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ListingContractRoomDeletedIterator is returned from FilterRoomDeleted and is used to iterate over the raw logs and unpacked data for RoomDeleted events raised by the ListingContract contract.
type ListingContractRoomDeletedIterator struct {
	Event *ListingContractRoomDeleted // Event containing the contract specifics and raw log

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
func (it *ListingContractRoomDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingContractRoomDeleted)
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
		it.Event = new(ListingContractRoomDeleted)
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
func (it *ListingContractRoomDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingContractRoomDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingContractRoomDeleted represents a RoomDeleted event raised by the ListingContract contract.
type ListingContractRoomDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRoomDeleted is a free log retrieval operation binding the contract event 0x640a818bdd6ca14bdd2cf9cdac867db41b7a8832e5c24995d200910d60e7e4dc.
//
// Solidity: event RoomDeleted(uint256 id)
func (_ListingContract *ListingContractFilterer) FilterRoomDeleted(opts *bind.FilterOpts) (*ListingContractRoomDeletedIterator, error) {

	logs, sub, err := _ListingContract.contract.FilterLogs(opts, "RoomDeleted")
	if err != nil {
		return nil, err
	}
	return &ListingContractRoomDeletedIterator{contract: _ListingContract.contract, event: "RoomDeleted", logs: logs, sub: sub}, nil
}

// WatchRoomDeleted is a free log subscription operation binding the contract event 0x640a818bdd6ca14bdd2cf9cdac867db41b7a8832e5c24995d200910d60e7e4dc.
//
// Solidity: event RoomDeleted(uint256 id)
func (_ListingContract *ListingContractFilterer) WatchRoomDeleted(opts *bind.WatchOpts, sink chan<- *ListingContractRoomDeleted) (event.Subscription, error) {

	logs, sub, err := _ListingContract.contract.WatchLogs(opts, "RoomDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingContractRoomDeleted)
				if err := _ListingContract.contract.UnpackLog(event, "RoomDeleted", log); err != nil {
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

// ParseRoomDeleted is a log parse operation binding the contract event 0x640a818bdd6ca14bdd2cf9cdac867db41b7a8832e5c24995d200910d60e7e4dc.
//
// Solidity: event RoomDeleted(uint256 id)
func (_ListingContract *ListingContractFilterer) ParseRoomDeleted(log types.Log) (*ListingContractRoomDeleted, error) {
	event := new(ListingContractRoomDeleted)
	if err := _ListingContract.contract.UnpackLog(event, "RoomDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ListingContractRoomUpdatedIterator is returned from FilterRoomUpdated and is used to iterate over the raw logs and unpacked data for RoomUpdated events raised by the ListingContract contract.
type ListingContractRoomUpdatedIterator struct {
	Event *ListingContractRoomUpdated // Event containing the contract specifics and raw log

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
func (it *ListingContractRoomUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingContractRoomUpdated)
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
		it.Event = new(ListingContractRoomUpdated)
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
func (it *ListingContractRoomUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingContractRoomUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingContractRoomUpdated represents a RoomUpdated event raised by the ListingContract contract.
type ListingContractRoomUpdated struct {
	Id         *big.Int
	Owner      *big.Int
	TotalPrice *big.Int
	Deposit    *big.Int
	Status     *big.Int
	IsRent     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoomUpdated is a free log retrieval operation binding the contract event 0x534c7676f2d5efde7c224ea38906c043ed4c7d44482023b66a5403806b5ce65f.
//
// Solidity: event RoomUpdated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) FilterRoomUpdated(opts *bind.FilterOpts) (*ListingContractRoomUpdatedIterator, error) {

	logs, sub, err := _ListingContract.contract.FilterLogs(opts, "RoomUpdated")
	if err != nil {
		return nil, err
	}
	return &ListingContractRoomUpdatedIterator{contract: _ListingContract.contract, event: "RoomUpdated", logs: logs, sub: sub}, nil
}

// WatchRoomUpdated is a free log subscription operation binding the contract event 0x534c7676f2d5efde7c224ea38906c043ed4c7d44482023b66a5403806b5ce65f.
//
// Solidity: event RoomUpdated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) WatchRoomUpdated(opts *bind.WatchOpts, sink chan<- *ListingContractRoomUpdated) (event.Subscription, error) {

	logs, sub, err := _ListingContract.contract.WatchLogs(opts, "RoomUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingContractRoomUpdated)
				if err := _ListingContract.contract.UnpackLog(event, "RoomUpdated", log); err != nil {
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

// ParseRoomUpdated is a log parse operation binding the contract event 0x534c7676f2d5efde7c224ea38906c043ed4c7d44482023b66a5403806b5ce65f.
//
// Solidity: event RoomUpdated(uint256 id, uint256 owner, uint256 totalPrice, uint256 deposit, uint256 status, bool isRent)
func (_ListingContract *ListingContractFilterer) ParseRoomUpdated(log types.Log) (*ListingContractRoomUpdated, error) {
	event := new(ListingContractRoomUpdated)
	if err := _ListingContract.contract.UnpackLog(event, "RoomUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
