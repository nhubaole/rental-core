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

// ContractManagementMetaData contains all meta data concerning the ContractManagement contract.
var ContractManagementMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"landlord\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tenant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beginDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"MContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"MContractDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"MContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"approveReturnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_landlord\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tenant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_roomId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beginDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_paymentMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electricityMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_electricityCost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_waterMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_waterCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_internetCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_parkingFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_responsibilityA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_responsibilityB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_generalResponsibility\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signatureA\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_signedTimeA\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_signatureB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_signedTimeB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_contractTemplateId\",\"type\":\"uint256\"}],\"name\":\"createMContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createReturnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumContractManagement.PreRentalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.RentalProcessStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.PostRentalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"inputMeterReading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"landlord\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tenant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roomId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beginDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"paymentMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electricityMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"electricityCost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"waterMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"waterCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"internetCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parkingFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"responsibilityA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"responsibilityB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"generalResponsibility\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatureA\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signedTimeA\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signatureB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signedTimeB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractTemplateId\",\"type\":\"uint256\"},{\"internalType\":\"enumContractManagement.PreRentalStatus\",\"name\":\"preRentalStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.RentalProcessStatus\",\"name\":\"rentalProcessStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.PostRentalStatus\",\"name\":\"postRentalStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"payBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"payDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_signatureB\",\"type\":\"string\"}],\"name\":\"signContractByTenant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611fa3908161001a8239f35b5f80fdfe6105006040526004361015610012575f80fd5b5f3560e01c80631c8f078914611c5f57806322f3bd3b146118de57806328cd3f18146118095780636f039de7146115fe578063712edc611461153b5780637874336f14611475578063a08bae231461131f578063a5244e6414610fa8578063d6d433db1461015d5763f097519014610088575f80fd5b3461015957602036600319011261015957600435805f525f60205260405f20906100bd600160ff601b85015416151514611f01565b601882019060ff825460081c16600481101561014557600203610108575f516020611f4e5f395f51905f529260209261030061ff0019825416179055601a42910155604051908152a1005b60405162461bcd60e51b81526020600482015260156024820152744e6f20756e706169642062696c6c20746f2070617960581b6044820152606490fd5b634e487b7160e01b5f52602160045260245ffd5b5f80fd5b3461015957610300366003190112610159576024356001600160401b0381116101595761018e903690600401611dd4565b610124356001600160401b038111610159576101ae903690600401611dd4565b610144356001600160401b038111610159576101ce903690600401611dd4565b610184356001600160401b038111610159576101ee903690600401611dd4565b610204356001600160401b0381116101595761020e903690600401611dd4565b610224356001600160401b0381116101595761022e903690600401611dd4565b610244356001600160401b0381116101595761024e903690600401611dd4565b91610264356001600160401b0381116101595761026f903690600401611dd4565b936102a4356001600160401b03811161015957610290903690600401611dd4565b956004355f525f60205260ff601b60405f20015416610f6357604051976103c089018981106001600160401b038211176109735760405260043589528960208a015260443560408a015260643560608a015260843560808a015260a43560a08a015260c43560c08a015260e43560e08a0152610104356101008a0152610120890152610140880152610164356101608801526101808701526101a4356101a08701526101c4356101c08701526101e4356101e0870152610200860152610220850152610240840152610260830152610284356102808301526102a08201526102c4356102c08201526102e4356102e08201525f6103008201525f6103208201525f610340820152426103608201524261038082015260016103a08201526004355f525f60205260405f208151815560208201519283516001600160401b038111610973576103e16001840154611e29565b601f8111610f1f575b506020601f8211600114610eb55781929394955f92610eaa575b50508160011b915f199060031b1c19161760018301555b60408301516002830155606083015160038301556080830151600483015560a0830151600583015560c0830151600683015560e0830151600783015561010083015160088301556101208301519283516001600160401b038111610973576104866009850154611e29565b601f8111610e66575b506020601f8211600114610dfc5781929394955f92610df1575b50508160011b915f199060031b1c19161760098401555b6101408101519283516001600160401b038111610973576104e4600a830154611e29565b601f8111610dad575b506020601f8211600114610d435781929394955f92610d38575b50508160011b915f199060031b1c191617600a8201555b610160820151600b8201556101808201519283516001600160401b0381116109735761054d600c840154611e29565b601f8111610cf4575b506020601f8211600114610c8a5781929394955f92610c7f575b50508160011b915f199060031b1c191617600c8301555b6101a0830151600d8301556101c0830151600e8301556101e0830151600f830155601082016102008401518051906001600160401b0382116109735781906105cf8454611e29565b601f8111610c2f575b50602090601f8311600114610bcc575f92610bc1575b50508160011b915f199060031b1c19161790555b601182016102208401518051906001600160401b0382116109735781906106298454611e29565b601f8111610b71575b50602090601f8311600114610b0e575f92610b03575b50508160011b915f199060031b1c19161790555b601282016102408401518051906001600160401b0382116109735781906106838454611e29565b601f8111610ab3575b50602090601f8311600114610a50575f92610a45575b50508160011b915f199060031b1c19161790555b601382016102608401518051906001600160401b0382116109735781906106dd8454611e29565b601f81116109f5575b50602090601f8311600114610992575f92610987575b50508160011b915f199060031b1c19161790555b6102808301516014830155601582016102a08401518051906001600160401b0382116109735781906107428454611e29565b601f8111610923575b50602090601f83116001146108c0575f926108b5575b50508160011b915f199060031b1c19161790555b6102c083015160168301556102e08301516017830155601882019261030081015161079f81611d9c565b6107a881611d9c565b60ff8019865416911617845561032081015193600485101561014557601b610876946103a0927fada3db4246fb7379c45e0f4bad05d62f46c69cbc3b3c774d8e9aa0f94c98af949761ff0082549160081b169061ff00191617815561034085015161081281611d9c565b61081b81611d9c565b62ff000082549160101b169062ff000019161790556103608401516019820155610380840151601a82015501910151151560ff8019835416911617905560405191829160043583526101206020840152610120830190611d78565b60443560408301526064356060830152608435608083015260a43560a083015260c43560c083015260e43560e0830152610104356101008301520390a1005b015190508680610761565b5f8581528281209350601f198516905b81811061090b57509084600195949392106108f3575b505050811b019055610775565b01515f1960f88460031b161c191690558680806108e6565b929360206001819287860151815501950193016108d0565b909150835f5260205f20601f840160051c81019160208510610969575b90601f859493920160051c01905b81811061095b575061074b565b5f815584935060010161094e565b9091508190610940565b634e487b7160e01b5f52604160045260245ffd5b0151905086806106fc565b5f8581528281209350601f198516905b8181106109dd57509084600195949392106109c5575b505050811b019055610710565b01515f1960f88460031b161c191690558680806109b8565b929360206001819287860151815501950193016109a2565b909150835f5260205f20601f840160051c81019160208510610a3b575b90601f859493920160051c01905b818110610a2d57506106e6565b5f8155849350600101610a20565b9091508190610a12565b0151905086806106a2565b5f8581528281209350601f198516905b818110610a9b5750908460019594939210610a83575b505050811b0190556106b6565b01515f1960f88460031b161c19169055868080610a76565b92936020600181928786015181550195019301610a60565b909150835f5260205f20601f840160051c81019160208510610af9575b90601f859493920160051c01905b818110610aeb575061068c565b5f8155849350600101610ade565b9091508190610ad0565b015190508680610648565b5f8581528281209350601f198516905b818110610b595750908460019594939210610b41575b505050811b01905561065c565b01515f1960f88460031b161c19169055868080610b34565b92936020600181928786015181550195019301610b1e565b909150835f5260205f20601f840160051c81019160208510610bb7575b90601f859493920160051c01905b818110610ba95750610632565b5f8155849350600101610b9c565b9091508190610b8e565b0151905086806105ee565b5f8581528281209350601f198516905b818110610c175750908460019594939210610bff575b505050811b019055610602565b01515f1960f88460031b161c19169055868080610bf2565b92936020600181928786015181550195019301610bdc565b909150835f5260205f20601f840160051c81019160208510610c75575b90601f859493920160051c01905b818110610c6757506105d8565b5f8155849350600101610c5a565b9091508190610c4c565b015190508580610570565b600c84015f52805f20905f5b601f1984168110610cdc575060019394959683601f19811610610cc4575b505050811b01600c830155610587565b01515f1960f88460031b161c19169055858080610cb4565b9091602060018192858b015181550193019101610c96565b600c84015f5260205f20601f830160051c810160208410610d31575b601f830160051c82018110610d26575050610556565b5f8155600101610d10565b5080610d10565b015190508580610507565b600a83015f52805f20905f5b601f1984168110610d95575060019394959683601f19811610610d7d575b505050811b01600a82015561051e565b01515f1960f88460031b161c19169055858080610d6d565b9091602060018192858b015181550193019101610d4f565b600a83015f5260205f20601f830160051c810160208410610dea575b601f830160051c82018110610ddf5750506104ed565b5f8155600101610dc9565b5080610dc9565b0151905085806104a9565b600985015f52805f20905f5b601f1984168110610e4e575060019394959683601f19811610610e36575b505050811b0160098401556104c0565b01515f1960f88460031b161c19169055858080610e26565b9091602060018192858b015181550193019101610e08565b600985015f5260205f20601f830160051c810160208410610ea3575b601f830160051c82018110610e9857505061048f565b5f8155600101610e82565b5080610e82565b015190508580610404565b600184015f52805f20905f5b601f1984168110610f07575060019394959683601f19811610610eef575b505050811b01600183015561041b565b01515f1960f88460031b161c19169055858080610edf565b9091602060018192858b015181550193019101610ec1565b600184015f5260205f20601f830160051c810160208410610f5c575b601f830160051c82018110610f515750506103ea565b5f8155600101610f3b565b5080610f3b565b60405162461bcd60e51b815260206004820152601860248201527f4d436f6e747261637420616c72656164792065786973747300000000000000006044820152606490fd5b34610159576020366003190112610159576004355f525f60205260405f206102a05261127e61125c61124661123061121a6102a05154610fed60016102a05101611e61565b9060026102a051015460e05260036102a05101546101005260046102a0510154610240526111e760056102a051015460066102a05101546111cc60076102a0510154926111b960086102a05101549161104b60096102a05101611e61565b61105a600a6102a05101611e61565b93600b6102a051015497611073600c6102a05101611e61565b97600d6102a05101546101a05261116a600e6102a05101549d600f6102a05101549d6110a460106102a05101611e61565b6101c0526110b760116102a05101611e61565b6101e0526110ca60126102a05101611e61565b610200526110dd60136102a05101611e61565b6102805260146102a0510154610260526110fc60156102a05101611e61565b6102205260166102a051015460c05260176102a05101546101805260186102a05101546101605260196102a051015461014052601a6102a05101546101205260ff601b6102a05101541660a052604051608052608051526103c0602060805101526103c06080510190611d78565b9460e0516040608051015261010051606060805101526102405160808051015260a0608051015260c0608051015260e06080510152610100608051015260805182036101206080510152611d78565b9060805182036101406080510152611d78565b91610160608051015260805182036101806080510152611d78565b916101a0516101a060805101526101c060805101526101e06080510152608051810361020060805101526101c051611d78565b608051810361022060805101526101e051611d78565b6080518103610240608051015261020051611d78565b6080518103610260608051015261028051611d78565b61026051610280608051015260805181036102a0608051015261022051611d78565b60c0516102c06080510152610180516102e060805101526112a460ff6101605116611d9c565b60ff610160511661030060805101526112cc6103206080510160ff6101605160081c16611da6565b6112de60ff6101605160101c16611d9c565b60ff6101605160101c16610340608051015261014051610360608051015261012051610380608051015260a05115156103a060805101526080519003608051f35b3461015957602036600319011261015957600435805f525f60205260405f2090611354600160ff601b85015416151514611f01565b60188201908154600260ff821661136a81611d9c565b0361141c5760081c60ff16600481101561014557156113b1575f516020611f4e5f395f51905f529260209261010061ff0019825416179055601a42910155604051908152a1005b60405162461bcd60e51b815260206004820152603860248201527f4d657465722072656164696e672063616e206f6e6c79206265207265636f726460448201527f65642061667465722074686520666972737420706861736500000000000000006064820152608490fd5b60405162461bcd60e51b815260206004820152602b60248201527f4465706f736974206d7573742062652070616964206265666f7265207265636f60448201526a393234b7339036b2ba32b960a91b6064820152608490fd5b3461015957602036600319011261015957600435805f525f60205260405f206114a9600160ff601b84015416151514611f01565b6018810191825491600160ff8460101c166114c381611d9c565b036114f6575f516020611f4e5f395f51905f52936202000060209462ff00001916179055601a42910155604051908152a1005b60405162461bcd60e51b815260206004820152601960248201527f4e6f2070656e64696e672072657475726e2072657175657374000000000000006044820152606490fd5b3461015957602036600319011261015957600435805f525f60205260405f2061156f600160ff601b84015416151514611f01565b601881019182549160ff8360101c1661158781611d9c565b6115b9575f516020611f4e5f395f51905f52936201000060209462ff00001916179055601a42910155604051908152a1005b60405162461bcd60e51b815260206004820152601d60248201527f52657475726e207265717565737420616c7265616479206578697374730000006044820152606490fd5b34610159576040366003190112610159576004356024356001600160401b03811161015957611631903690600401611dd4565b815f525f60205260405f2090611652600160ff601b85015416151514611f01565b601882019060ff82541661166581611d9c565b6117ba57805160158401916001600160401b038211610973576116888354611e29565b601f8111611775575b50602090601f83116001146116fe57918060209694925f516020611f4e5f395f51905f529896945f926116f3575b50508160011b915f199060031b1c19161790555b426016830155600160ff19825416179055601a42910155604051908152a1005b0151905088806116bf565b90601f19831691845f52815f20925f5b81811061175d57509260019285925f516020611f4e5f395f51905f529a989660209a989610611745575b505050811b0190556116d3565b01515f1960f88460031b161c19169055888080611738565b9293602060018192878601518155019501930161170e565b835f5260205f20601f840160051c810191602085106117b0575b601f0160051c01905b8181106117a55750611691565b5f8155600101611798565b909150819061178f565b60405162461bcd60e51b815260206004820152602160248201527f436f6e7472616374206973206e6f7420726561647920666f72207369676e696e6044820152606760f81b6064820152608490fd5b3461015957602036600319011261015957600435805f525f60205260405f2061183d600160ff601b84015416151514611f01565b6018810191825491600160ff841661185481611d9c565b03611883575f516020611f4e5f395f51905f5293600260209460ff1916179055601a42910155604051908152a1005b60405162461bcd60e51b815260206004820152602d60248201527f436f6e7472616374206d757374206265207369676e6564206265666f7265207060448201526c185e5a5b99c819195c1bdcda5d609a1b6064820152608490fd5b34610159576020366003190112610159576004355f525f60205260405f2061046052611918600160ff601b61046051015416151514611f01565b6104605180546104e05260028101546103c05260038101546103e05260048101546104405260058101546006820154600783015460088401546104c052600b840154600d850154600e860154600f87015460148801546104205260168801546103a05260178801546103605260188801546104805260198801546104a052601a88015461038052611bc397611b9e97611b8697611b6e97611b5697949594611b229490939192611b04929091611aef91906119d590600101611e61565b946119e560096104605101611e61565b906119f5600a6104605101611e61565b94611a91611a08600c6104605101611e61565b98611a1860106104605101611e61565b6102c052611a2b60116104605101611e61565b6102e052611a3e60126104605101611e61565b61032052611a5160136104605101611e61565b61040052611a6460156104605101611e61565b61034052604051610300526104e05161030051526103a060206103005101526103a0610300510190611d78565b936103c05160406103005101526103e051606061030051015261044051608061030051015260a061030051015260c061030051015260e06103005101526104c051610100610300510152610300518203610120610300510152611d78565b90610300518203610140610300510152611d78565b91610160610300510152610300518203610180610300510152611d78565b926101a06103005101526101c06103005101526101e06103005101526103005181036102006103005101526102c051611d78565b6103005181036102206103005101526102e051611d78565b61030051810361024061030051015261032051611d78565b61030051810361026061030051015261040051611d78565b610420516102806103005101526103005181036102a061030051015261034051611d78565b6103a0516102c0610300510152610360516102e0610300510152611bec60ff6104805116611d9c565b60ff610480511661030080510152611c14610320610300510160ff6104805160081c16611da6565b611c2660ff6104805160101c16611d9c565b60ff6104805160101c166103406103005101526104a0516103606103005101526103805161038061030051015261030051900361030051f35b3461015957602036600319011261015957600435805f525f60205260405f2090611c94600160ff601b85015416151514611f01565b601882019060ff825460081c166004811015610145578015908115611d6d575b5015611ce8575f516020611f4e5f395f51905f529260209261020061ff0019825416179055601a42910155604051908152a1005b60405162461bcd60e51b815260206004820152605160248201527f42696c6c206372656174696f6e2072657175697265732066697273742070686160448201527f7365206f72207265636f72646564206d657465722072656164696e6720696e2060648201527073756273657175656e742070686173657360781b608482015260a490fd5b600191501484611cb4565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b6003111561014557565b9060048210156101455752565b90601f801991011681019081106001600160401b0382111761097357604052565b81601f82011215610159578035906001600160401b0382116109735760405192611e08601f8401601f191660200185611db3565b8284526020838301011161015957815f926020809301838601378301015290565b90600182811c92168015611e57575b6020831014611e4357565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611e38565b9060405191825f825492611e7484611e29565b8084529360018116908115611edf5750600114611e9b575b50611e9992500383611db3565b565b90505f9291925260205f20905f915b818310611ec3575050906020611e99928201015f611e8c565b6020919350806001915483858901015201910190918492611eaa565b905060209250611e9994915060ff191682840152151560051b8201015f611e8c565b15611f0857565b60405162461bcd60e51b815260206004820152601860248201527f4d436f6e747261637420646f6573206e6f7420657869737400000000000000006044820152606490fdfe2752a706b121abafb4dd58bef7d110189317241e7578a771ac0cfa40ce58b58fa2646970667358221220ffdeeb70eec8b4fcb8fa2b30bd8eac49f6fcee994cbf3b882185aa4f77967cd564736f6c634300081c0033",
}

// ContractManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractManagementMetaData.ABI instead.
var ContractManagementABI = ContractManagementMetaData.ABI

// ContractManagementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractManagementMetaData.Bin instead.
var ContractManagementBin = ContractManagementMetaData.Bin

// DeployContractManagement deploys a new Ethereum contract, binding an instance of ContractManagement to it.
func DeployContractManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractManagement, error) {
	parsed, err := ContractManagementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractManagement{ContractManagementCaller: ContractManagementCaller{contract: contract}, ContractManagementTransactor: ContractManagementTransactor{contract: contract}, ContractManagementFilterer: ContractManagementFilterer{contract: contract}}, nil
}

// ContractManagement is an auto generated Go binding around an Ethereum contract.
type ContractManagement struct {
	ContractManagementCaller     // Read-only binding to the contract
	ContractManagementTransactor // Write-only binding to the contract
	ContractManagementFilterer   // Log filterer for contract events
}

// ContractManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractManagementSession struct {
	Contract     *ContractManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractManagementCallerSession struct {
	Contract *ContractManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractManagementTransactorSession struct {
	Contract     *ContractManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractManagementRaw struct {
	Contract *ContractManagement // Generic contract binding to access the raw methods on
}

// ContractManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractManagementCallerRaw struct {
	Contract *ContractManagementCaller // Generic read-only contract binding to access the raw methods on
}

// ContractManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractManagementTransactorRaw struct {
	Contract *ContractManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractManagement creates a new instance of ContractManagement, bound to a specific deployed contract.
func NewContractManagement(address common.Address, backend bind.ContractBackend) (*ContractManagement, error) {
	contract, err := bindContractManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractManagement{ContractManagementCaller: ContractManagementCaller{contract: contract}, ContractManagementTransactor: ContractManagementTransactor{contract: contract}, ContractManagementFilterer: ContractManagementFilterer{contract: contract}}, nil
}

// NewContractManagementCaller creates a new read-only instance of ContractManagement, bound to a specific deployed contract.
func NewContractManagementCaller(address common.Address, caller bind.ContractCaller) (*ContractManagementCaller, error) {
	contract, err := bindContractManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractManagementCaller{contract: contract}, nil
}

// NewContractManagementTransactor creates a new write-only instance of ContractManagement, bound to a specific deployed contract.
func NewContractManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractManagementTransactor, error) {
	contract, err := bindContractManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractManagementTransactor{contract: contract}, nil
}

// NewContractManagementFilterer creates a new log filterer instance of ContractManagement, bound to a specific deployed contract.
func NewContractManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractManagementFilterer, error) {
	contract, err := bindContractManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractManagementFilterer{contract: contract}, nil
}

// bindContractManagement binds a generic wrapper to an already deployed contract.
func bindContractManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractManagement *ContractManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractManagement.Contract.ContractManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractManagement *ContractManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractManagement.Contract.ContractManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractManagement *ContractManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractManagement.Contract.ContractManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractManagement *ContractManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractManagement *ContractManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractManagement *ContractManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractManagement.Contract.contract.Transact(opts, method, params...)
}

// GetMContract is a free data retrieval call binding the contract method 0x22f3bd3b.
//
// Solidity: function getMContract(uint256 _id) view returns(uint256, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256, string, string, uint256, string, uint256, uint256, uint256, string, string, string, string, uint256, string, uint256, uint256, uint8, uint8, uint8, uint256, uint256)
func (_ContractManagement *ContractManagementCaller) GetMContract(opts *bind.CallOpts, _id *big.Int) (*big.Int, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, string, string, *big.Int, string, *big.Int, *big.Int, *big.Int, string, string, string, string, *big.Int, string, *big.Int, *big.Int, uint8, uint8, uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ContractManagement.contract.Call(opts, &out, "getMContract", _id)

	if err != nil {
		return *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(string), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(string), *new(string), *new(string), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), *new(uint8), *new(uint8), *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(string)).(*string)
	out10 := *abi.ConvertType(out[10], new(string)).(*string)
	out11 := *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	out12 := *abi.ConvertType(out[12], new(string)).(*string)
	out13 := *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)
	out14 := *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)
	out15 := *abi.ConvertType(out[15], new(*big.Int)).(**big.Int)
	out16 := *abi.ConvertType(out[16], new(string)).(*string)
	out17 := *abi.ConvertType(out[17], new(string)).(*string)
	out18 := *abi.ConvertType(out[18], new(string)).(*string)
	out19 := *abi.ConvertType(out[19], new(string)).(*string)
	out20 := *abi.ConvertType(out[20], new(*big.Int)).(**big.Int)
	out21 := *abi.ConvertType(out[21], new(string)).(*string)
	out22 := *abi.ConvertType(out[22], new(*big.Int)).(**big.Int)
	out23 := *abi.ConvertType(out[23], new(*big.Int)).(**big.Int)
	out24 := *abi.ConvertType(out[24], new(uint8)).(*uint8)
	out25 := *abi.ConvertType(out[25], new(uint8)).(*uint8)
	out26 := *abi.ConvertType(out[26], new(uint8)).(*uint8)
	out27 := *abi.ConvertType(out[27], new(*big.Int)).(**big.Int)
	out28 := *abi.ConvertType(out[28], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, out14, out15, out16, out17, out18, out19, out20, out21, out22, out23, out24, out25, out26, out27, out28, err

}

// GetMContract is a free data retrieval call binding the contract method 0x22f3bd3b.
//
// Solidity: function getMContract(uint256 _id) view returns(uint256, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256, string, string, uint256, string, uint256, uint256, uint256, string, string, string, string, uint256, string, uint256, uint256, uint8, uint8, uint8, uint256, uint256)
func (_ContractManagement *ContractManagementSession) GetMContract(_id *big.Int) (*big.Int, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, string, string, *big.Int, string, *big.Int, *big.Int, *big.Int, string, string, string, string, *big.Int, string, *big.Int, *big.Int, uint8, uint8, uint8, *big.Int, *big.Int, error) {
	return _ContractManagement.Contract.GetMContract(&_ContractManagement.CallOpts, _id)
}

// GetMContract is a free data retrieval call binding the contract method 0x22f3bd3b.
//
// Solidity: function getMContract(uint256 _id) view returns(uint256, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256, string, string, uint256, string, uint256, uint256, uint256, string, string, string, string, uint256, string, uint256, uint256, uint8, uint8, uint8, uint256, uint256)
func (_ContractManagement *ContractManagementCallerSession) GetMContract(_id *big.Int) (*big.Int, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, string, string, *big.Int, string, *big.Int, *big.Int, *big.Int, string, string, string, string, *big.Int, string, *big.Int, *big.Int, uint8, uint8, uint8, *big.Int, *big.Int, error) {
	return _ContractManagement.Contract.GetMContract(&_ContractManagement.CallOpts, _id)
}

// MContracts is a free data retrieval call binding the contract method 0xa5244e64.
//
// Solidity: function mContracts(uint256 ) view returns(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate, string paymentMethod, string electricityMethod, uint256 electricityCost, string waterMethod, uint256 waterCost, uint256 internetCost, uint256 parkingFee, string responsibilityA, string responsibilityB, string generalResponsibility, string signatureA, uint256 signedTimeA, string signatureB, uint256 signedTimeB, uint256 contractTemplateId, uint8 preRentalStatus, uint8 rentalProcessStatus, uint8 postRentalStatus, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ContractManagement *ContractManagementCaller) MContracts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                    *big.Int
	Code                  string
	Landlord              *big.Int
	Tenant                *big.Int
	RoomId                *big.Int
	ActualPrice           *big.Int
	Deposit               *big.Int
	BeginDate             *big.Int
	EndDate               *big.Int
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       *big.Int
	WaterMethod           string
	WaterCost             *big.Int
	InternetCost          *big.Int
	ParkingFee            *big.Int
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           *big.Int
	SignatureB            string
	SignedTimeB           *big.Int
	ContractTemplateId    *big.Int
	PreRentalStatus       uint8
	RentalProcessStatus   uint8
	PostRentalStatus      uint8
	CreatedAt             *big.Int
	UpdatedAt             *big.Int
	Exists                bool
}, error) {
	var out []interface{}
	err := _ContractManagement.contract.Call(opts, &out, "mContracts", arg0)

	outstruct := new(struct {
		Id                    *big.Int
		Code                  string
		Landlord              *big.Int
		Tenant                *big.Int
		RoomId                *big.Int
		ActualPrice           *big.Int
		Deposit               *big.Int
		BeginDate             *big.Int
		EndDate               *big.Int
		PaymentMethod         string
		ElectricityMethod     string
		ElectricityCost       *big.Int
		WaterMethod           string
		WaterCost             *big.Int
		InternetCost          *big.Int
		ParkingFee            *big.Int
		ResponsibilityA       string
		ResponsibilityB       string
		GeneralResponsibility string
		SignatureA            string
		SignedTimeA           *big.Int
		SignatureB            string
		SignedTimeB           *big.Int
		ContractTemplateId    *big.Int
		PreRentalStatus       uint8
		RentalProcessStatus   uint8
		PostRentalStatus      uint8
		CreatedAt             *big.Int
		UpdatedAt             *big.Int
		Exists                bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Code = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Landlord = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tenant = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RoomId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ActualPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Deposit = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BeginDate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.PaymentMethod = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.ElectricityMethod = *abi.ConvertType(out[10], new(string)).(*string)
	outstruct.ElectricityCost = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.WaterMethod = *abi.ConvertType(out[12], new(string)).(*string)
	outstruct.WaterCost = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)
	outstruct.InternetCost = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)
	outstruct.ParkingFee = *abi.ConvertType(out[15], new(*big.Int)).(**big.Int)
	outstruct.ResponsibilityA = *abi.ConvertType(out[16], new(string)).(*string)
	outstruct.ResponsibilityB = *abi.ConvertType(out[17], new(string)).(*string)
	outstruct.GeneralResponsibility = *abi.ConvertType(out[18], new(string)).(*string)
	outstruct.SignatureA = *abi.ConvertType(out[19], new(string)).(*string)
	outstruct.SignedTimeA = *abi.ConvertType(out[20], new(*big.Int)).(**big.Int)
	outstruct.SignatureB = *abi.ConvertType(out[21], new(string)).(*string)
	outstruct.SignedTimeB = *abi.ConvertType(out[22], new(*big.Int)).(**big.Int)
	outstruct.ContractTemplateId = *abi.ConvertType(out[23], new(*big.Int)).(**big.Int)
	outstruct.PreRentalStatus = *abi.ConvertType(out[24], new(uint8)).(*uint8)
	outstruct.RentalProcessStatus = *abi.ConvertType(out[25], new(uint8)).(*uint8)
	outstruct.PostRentalStatus = *abi.ConvertType(out[26], new(uint8)).(*uint8)
	outstruct.CreatedAt = *abi.ConvertType(out[27], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[28], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[29], new(bool)).(*bool)

	return *outstruct, err

}

// MContracts is a free data retrieval call binding the contract method 0xa5244e64.
//
// Solidity: function mContracts(uint256 ) view returns(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate, string paymentMethod, string electricityMethod, uint256 electricityCost, string waterMethod, uint256 waterCost, uint256 internetCost, uint256 parkingFee, string responsibilityA, string responsibilityB, string generalResponsibility, string signatureA, uint256 signedTimeA, string signatureB, uint256 signedTimeB, uint256 contractTemplateId, uint8 preRentalStatus, uint8 rentalProcessStatus, uint8 postRentalStatus, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ContractManagement *ContractManagementSession) MContracts(arg0 *big.Int) (struct {
	Id                    *big.Int
	Code                  string
	Landlord              *big.Int
	Tenant                *big.Int
	RoomId                *big.Int
	ActualPrice           *big.Int
	Deposit               *big.Int
	BeginDate             *big.Int
	EndDate               *big.Int
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       *big.Int
	WaterMethod           string
	WaterCost             *big.Int
	InternetCost          *big.Int
	ParkingFee            *big.Int
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           *big.Int
	SignatureB            string
	SignedTimeB           *big.Int
	ContractTemplateId    *big.Int
	PreRentalStatus       uint8
	RentalProcessStatus   uint8
	PostRentalStatus      uint8
	CreatedAt             *big.Int
	UpdatedAt             *big.Int
	Exists                bool
}, error) {
	return _ContractManagement.Contract.MContracts(&_ContractManagement.CallOpts, arg0)
}

// MContracts is a free data retrieval call binding the contract method 0xa5244e64.
//
// Solidity: function mContracts(uint256 ) view returns(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate, string paymentMethod, string electricityMethod, uint256 electricityCost, string waterMethod, uint256 waterCost, uint256 internetCost, uint256 parkingFee, string responsibilityA, string responsibilityB, string generalResponsibility, string signatureA, uint256 signedTimeA, string signatureB, uint256 signedTimeB, uint256 contractTemplateId, uint8 preRentalStatus, uint8 rentalProcessStatus, uint8 postRentalStatus, uint256 createdAt, uint256 updatedAt, bool exists)
func (_ContractManagement *ContractManagementCallerSession) MContracts(arg0 *big.Int) (struct {
	Id                    *big.Int
	Code                  string
	Landlord              *big.Int
	Tenant                *big.Int
	RoomId                *big.Int
	ActualPrice           *big.Int
	Deposit               *big.Int
	BeginDate             *big.Int
	EndDate               *big.Int
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       *big.Int
	WaterMethod           string
	WaterCost             *big.Int
	InternetCost          *big.Int
	ParkingFee            *big.Int
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           *big.Int
	SignatureB            string
	SignedTimeB           *big.Int
	ContractTemplateId    *big.Int
	PreRentalStatus       uint8
	RentalProcessStatus   uint8
	PostRentalStatus      uint8
	CreatedAt             *big.Int
	UpdatedAt             *big.Int
	Exists                bool
}, error) {
	return _ContractManagement.Contract.MContracts(&_ContractManagement.CallOpts, arg0)
}

// ApproveReturnRequest is a paid mutator transaction binding the contract method 0x7874336f.
//
// Solidity: function approveReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) ApproveReturnRequest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "approveReturnRequest", _id)
}

// ApproveReturnRequest is a paid mutator transaction binding the contract method 0x7874336f.
//
// Solidity: function approveReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) ApproveReturnRequest(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.ApproveReturnRequest(&_ContractManagement.TransactOpts, _id)
}

// ApproveReturnRequest is a paid mutator transaction binding the contract method 0x7874336f.
//
// Solidity: function approveReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) ApproveReturnRequest(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.ApproveReturnRequest(&_ContractManagement.TransactOpts, _id)
}

// CreateBill is a paid mutator transaction binding the contract method 0x1c8f0789.
//
// Solidity: function createBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) CreateBill(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "createBill", _id)
}

// CreateBill is a paid mutator transaction binding the contract method 0x1c8f0789.
//
// Solidity: function createBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) CreateBill(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateBill(&_ContractManagement.TransactOpts, _id)
}

// CreateBill is a paid mutator transaction binding the contract method 0x1c8f0789.
//
// Solidity: function createBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) CreateBill(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateBill(&_ContractManagement.TransactOpts, _id)
}

// CreateMContract is a paid mutator transaction binding the contract method 0xd6d433db.
//
// Solidity: function createMContract(uint256 _id, string _code, uint256 _landlord, uint256 _tenant, uint256 _roomId, uint256 _actualPrice, uint256 _deposit, uint256 _beginDate, uint256 _endDate, string _paymentMethod, string _electricityMethod, uint256 _electricityCost, string _waterMethod, uint256 _waterCost, uint256 _internetCost, uint256 _parkingFee, string _responsibilityA, string _responsibilityB, string _generalResponsibility, string _signatureA, uint256 _signedTimeA, string _signatureB, uint256 _signedTimeB, uint256 _contractTemplateId) returns()
func (_ContractManagement *ContractManagementTransactor) CreateMContract(opts *bind.TransactOpts, _id *big.Int, _code string, _landlord *big.Int, _tenant *big.Int, _roomId *big.Int, _actualPrice *big.Int, _deposit *big.Int, _beginDate *big.Int, _endDate *big.Int, _paymentMethod string, _electricityMethod string, _electricityCost *big.Int, _waterMethod string, _waterCost *big.Int, _internetCost *big.Int, _parkingFee *big.Int, _responsibilityA string, _responsibilityB string, _generalResponsibility string, _signatureA string, _signedTimeA *big.Int, _signatureB string, _signedTimeB *big.Int, _contractTemplateId *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "createMContract", _id, _code, _landlord, _tenant, _roomId, _actualPrice, _deposit, _beginDate, _endDate, _paymentMethod, _electricityMethod, _electricityCost, _waterMethod, _waterCost, _internetCost, _parkingFee, _responsibilityA, _responsibilityB, _generalResponsibility, _signatureA, _signedTimeA, _signatureB, _signedTimeB, _contractTemplateId)
}

// CreateMContract is a paid mutator transaction binding the contract method 0xd6d433db.
//
// Solidity: function createMContract(uint256 _id, string _code, uint256 _landlord, uint256 _tenant, uint256 _roomId, uint256 _actualPrice, uint256 _deposit, uint256 _beginDate, uint256 _endDate, string _paymentMethod, string _electricityMethod, uint256 _electricityCost, string _waterMethod, uint256 _waterCost, uint256 _internetCost, uint256 _parkingFee, string _responsibilityA, string _responsibilityB, string _generalResponsibility, string _signatureA, uint256 _signedTimeA, string _signatureB, uint256 _signedTimeB, uint256 _contractTemplateId) returns()
func (_ContractManagement *ContractManagementSession) CreateMContract(_id *big.Int, _code string, _landlord *big.Int, _tenant *big.Int, _roomId *big.Int, _actualPrice *big.Int, _deposit *big.Int, _beginDate *big.Int, _endDate *big.Int, _paymentMethod string, _electricityMethod string, _electricityCost *big.Int, _waterMethod string, _waterCost *big.Int, _internetCost *big.Int, _parkingFee *big.Int, _responsibilityA string, _responsibilityB string, _generalResponsibility string, _signatureA string, _signedTimeA *big.Int, _signatureB string, _signedTimeB *big.Int, _contractTemplateId *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateMContract(&_ContractManagement.TransactOpts, _id, _code, _landlord, _tenant, _roomId, _actualPrice, _deposit, _beginDate, _endDate, _paymentMethod, _electricityMethod, _electricityCost, _waterMethod, _waterCost, _internetCost, _parkingFee, _responsibilityA, _responsibilityB, _generalResponsibility, _signatureA, _signedTimeA, _signatureB, _signedTimeB, _contractTemplateId)
}

// CreateMContract is a paid mutator transaction binding the contract method 0xd6d433db.
//
// Solidity: function createMContract(uint256 _id, string _code, uint256 _landlord, uint256 _tenant, uint256 _roomId, uint256 _actualPrice, uint256 _deposit, uint256 _beginDate, uint256 _endDate, string _paymentMethod, string _electricityMethod, uint256 _electricityCost, string _waterMethod, uint256 _waterCost, uint256 _internetCost, uint256 _parkingFee, string _responsibilityA, string _responsibilityB, string _generalResponsibility, string _signatureA, uint256 _signedTimeA, string _signatureB, uint256 _signedTimeB, uint256 _contractTemplateId) returns()
func (_ContractManagement *ContractManagementTransactorSession) CreateMContract(_id *big.Int, _code string, _landlord *big.Int, _tenant *big.Int, _roomId *big.Int, _actualPrice *big.Int, _deposit *big.Int, _beginDate *big.Int, _endDate *big.Int, _paymentMethod string, _electricityMethod string, _electricityCost *big.Int, _waterMethod string, _waterCost *big.Int, _internetCost *big.Int, _parkingFee *big.Int, _responsibilityA string, _responsibilityB string, _generalResponsibility string, _signatureA string, _signedTimeA *big.Int, _signatureB string, _signedTimeB *big.Int, _contractTemplateId *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateMContract(&_ContractManagement.TransactOpts, _id, _code, _landlord, _tenant, _roomId, _actualPrice, _deposit, _beginDate, _endDate, _paymentMethod, _electricityMethod, _electricityCost, _waterMethod, _waterCost, _internetCost, _parkingFee, _responsibilityA, _responsibilityB, _generalResponsibility, _signatureA, _signedTimeA, _signatureB, _signedTimeB, _contractTemplateId)
}

// CreateReturnRequest is a paid mutator transaction binding the contract method 0x712edc61.
//
// Solidity: function createReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) CreateReturnRequest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "createReturnRequest", _id)
}

// CreateReturnRequest is a paid mutator transaction binding the contract method 0x712edc61.
//
// Solidity: function createReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) CreateReturnRequest(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateReturnRequest(&_ContractManagement.TransactOpts, _id)
}

// CreateReturnRequest is a paid mutator transaction binding the contract method 0x712edc61.
//
// Solidity: function createReturnRequest(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) CreateReturnRequest(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.CreateReturnRequest(&_ContractManagement.TransactOpts, _id)
}

// InputMeterReading is a paid mutator transaction binding the contract method 0xa08bae23.
//
// Solidity: function inputMeterReading(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) InputMeterReading(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "inputMeterReading", _id)
}

// InputMeterReading is a paid mutator transaction binding the contract method 0xa08bae23.
//
// Solidity: function inputMeterReading(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) InputMeterReading(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.InputMeterReading(&_ContractManagement.TransactOpts, _id)
}

// InputMeterReading is a paid mutator transaction binding the contract method 0xa08bae23.
//
// Solidity: function inputMeterReading(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) InputMeterReading(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.InputMeterReading(&_ContractManagement.TransactOpts, _id)
}

// PayBill is a paid mutator transaction binding the contract method 0xf0975190.
//
// Solidity: function payBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) PayBill(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "payBill", _id)
}

// PayBill is a paid mutator transaction binding the contract method 0xf0975190.
//
// Solidity: function payBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) PayBill(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.PayBill(&_ContractManagement.TransactOpts, _id)
}

// PayBill is a paid mutator transaction binding the contract method 0xf0975190.
//
// Solidity: function payBill(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) PayBill(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.PayBill(&_ContractManagement.TransactOpts, _id)
}

// PayDeposit is a paid mutator transaction binding the contract method 0x28cd3f18.
//
// Solidity: function payDeposit(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactor) PayDeposit(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "payDeposit", _id)
}

// PayDeposit is a paid mutator transaction binding the contract method 0x28cd3f18.
//
// Solidity: function payDeposit(uint256 _id) returns()
func (_ContractManagement *ContractManagementSession) PayDeposit(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.PayDeposit(&_ContractManagement.TransactOpts, _id)
}

// PayDeposit is a paid mutator transaction binding the contract method 0x28cd3f18.
//
// Solidity: function payDeposit(uint256 _id) returns()
func (_ContractManagement *ContractManagementTransactorSession) PayDeposit(_id *big.Int) (*types.Transaction, error) {
	return _ContractManagement.Contract.PayDeposit(&_ContractManagement.TransactOpts, _id)
}

// SignContractByTenant is a paid mutator transaction binding the contract method 0x6f039de7.
//
// Solidity: function signContractByTenant(uint256 _id, string _signatureB) returns()
func (_ContractManagement *ContractManagementTransactor) SignContractByTenant(opts *bind.TransactOpts, _id *big.Int, _signatureB string) (*types.Transaction, error) {
	return _ContractManagement.contract.Transact(opts, "signContractByTenant", _id, _signatureB)
}

// SignContractByTenant is a paid mutator transaction binding the contract method 0x6f039de7.
//
// Solidity: function signContractByTenant(uint256 _id, string _signatureB) returns()
func (_ContractManagement *ContractManagementSession) SignContractByTenant(_id *big.Int, _signatureB string) (*types.Transaction, error) {
	return _ContractManagement.Contract.SignContractByTenant(&_ContractManagement.TransactOpts, _id, _signatureB)
}

// SignContractByTenant is a paid mutator transaction binding the contract method 0x6f039de7.
//
// Solidity: function signContractByTenant(uint256 _id, string _signatureB) returns()
func (_ContractManagement *ContractManagementTransactorSession) SignContractByTenant(_id *big.Int, _signatureB string) (*types.Transaction, error) {
	return _ContractManagement.Contract.SignContractByTenant(&_ContractManagement.TransactOpts, _id, _signatureB)
}

// ContractManagementMContractCreatedIterator is returned from FilterMContractCreated and is used to iterate over the raw logs and unpacked data for MContractCreated events raised by the ContractManagement contract.
type ContractManagementMContractCreatedIterator struct {
	Event *ContractManagementMContractCreated // Event containing the contract specifics and raw log

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
func (it *ContractManagementMContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractManagementMContractCreated)
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
		it.Event = new(ContractManagementMContractCreated)
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
func (it *ContractManagementMContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractManagementMContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractManagementMContractCreated represents a MContractCreated event raised by the ContractManagement contract.
type ContractManagementMContractCreated struct {
	Id          *big.Int
	Code        string
	Landlord    *big.Int
	Tenant      *big.Int
	RoomId      *big.Int
	ActualPrice *big.Int
	Deposit     *big.Int
	BeginDate   *big.Int
	EndDate     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMContractCreated is a free log retrieval operation binding the contract event 0xada3db4246fb7379c45e0f4bad05d62f46c69cbc3b3c774d8e9aa0f94c98af94.
//
// Solidity: event MContractCreated(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate)
func (_ContractManagement *ContractManagementFilterer) FilterMContractCreated(opts *bind.FilterOpts) (*ContractManagementMContractCreatedIterator, error) {

	logs, sub, err := _ContractManagement.contract.FilterLogs(opts, "MContractCreated")
	if err != nil {
		return nil, err
	}
	return &ContractManagementMContractCreatedIterator{contract: _ContractManagement.contract, event: "MContractCreated", logs: logs, sub: sub}, nil
}

// WatchMContractCreated is a free log subscription operation binding the contract event 0xada3db4246fb7379c45e0f4bad05d62f46c69cbc3b3c774d8e9aa0f94c98af94.
//
// Solidity: event MContractCreated(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate)
func (_ContractManagement *ContractManagementFilterer) WatchMContractCreated(opts *bind.WatchOpts, sink chan<- *ContractManagementMContractCreated) (event.Subscription, error) {

	logs, sub, err := _ContractManagement.contract.WatchLogs(opts, "MContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractManagementMContractCreated)
				if err := _ContractManagement.contract.UnpackLog(event, "MContractCreated", log); err != nil {
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

// ParseMContractCreated is a log parse operation binding the contract event 0xada3db4246fb7379c45e0f4bad05d62f46c69cbc3b3c774d8e9aa0f94c98af94.
//
// Solidity: event MContractCreated(uint256 id, string code, uint256 landlord, uint256 tenant, uint256 roomId, uint256 actualPrice, uint256 deposit, uint256 beginDate, uint256 endDate)
func (_ContractManagement *ContractManagementFilterer) ParseMContractCreated(log types.Log) (*ContractManagementMContractCreated, error) {
	event := new(ContractManagementMContractCreated)
	if err := _ContractManagement.contract.UnpackLog(event, "MContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractManagementMContractDeletedIterator is returned from FilterMContractDeleted and is used to iterate over the raw logs and unpacked data for MContractDeleted events raised by the ContractManagement contract.
type ContractManagementMContractDeletedIterator struct {
	Event *ContractManagementMContractDeleted // Event containing the contract specifics and raw log

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
func (it *ContractManagementMContractDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractManagementMContractDeleted)
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
		it.Event = new(ContractManagementMContractDeleted)
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
func (it *ContractManagementMContractDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractManagementMContractDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractManagementMContractDeleted represents a MContractDeleted event raised by the ContractManagement contract.
type ContractManagementMContractDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMContractDeleted is a free log retrieval operation binding the contract event 0x51fce3624363739b0ba973f0a380a50be7199078766ca77865e67bd8493b6575.
//
// Solidity: event MContractDeleted(uint256 id)
func (_ContractManagement *ContractManagementFilterer) FilterMContractDeleted(opts *bind.FilterOpts) (*ContractManagementMContractDeletedIterator, error) {

	logs, sub, err := _ContractManagement.contract.FilterLogs(opts, "MContractDeleted")
	if err != nil {
		return nil, err
	}
	return &ContractManagementMContractDeletedIterator{contract: _ContractManagement.contract, event: "MContractDeleted", logs: logs, sub: sub}, nil
}

// WatchMContractDeleted is a free log subscription operation binding the contract event 0x51fce3624363739b0ba973f0a380a50be7199078766ca77865e67bd8493b6575.
//
// Solidity: event MContractDeleted(uint256 id)
func (_ContractManagement *ContractManagementFilterer) WatchMContractDeleted(opts *bind.WatchOpts, sink chan<- *ContractManagementMContractDeleted) (event.Subscription, error) {

	logs, sub, err := _ContractManagement.contract.WatchLogs(opts, "MContractDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractManagementMContractDeleted)
				if err := _ContractManagement.contract.UnpackLog(event, "MContractDeleted", log); err != nil {
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

// ParseMContractDeleted is a log parse operation binding the contract event 0x51fce3624363739b0ba973f0a380a50be7199078766ca77865e67bd8493b6575.
//
// Solidity: event MContractDeleted(uint256 id)
func (_ContractManagement *ContractManagementFilterer) ParseMContractDeleted(log types.Log) (*ContractManagementMContractDeleted, error) {
	event := new(ContractManagementMContractDeleted)
	if err := _ContractManagement.contract.UnpackLog(event, "MContractDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractManagementMContractUpdatedIterator is returned from FilterMContractUpdated and is used to iterate over the raw logs and unpacked data for MContractUpdated events raised by the ContractManagement contract.
type ContractManagementMContractUpdatedIterator struct {
	Event *ContractManagementMContractUpdated // Event containing the contract specifics and raw log

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
func (it *ContractManagementMContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractManagementMContractUpdated)
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
		it.Event = new(ContractManagementMContractUpdated)
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
func (it *ContractManagementMContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractManagementMContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractManagementMContractUpdated represents a MContractUpdated event raised by the ContractManagement contract.
type ContractManagementMContractUpdated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMContractUpdated is a free log retrieval operation binding the contract event 0x2752a706b121abafb4dd58bef7d110189317241e7578a771ac0cfa40ce58b58f.
//
// Solidity: event MContractUpdated(uint256 id)
func (_ContractManagement *ContractManagementFilterer) FilterMContractUpdated(opts *bind.FilterOpts) (*ContractManagementMContractUpdatedIterator, error) {

	logs, sub, err := _ContractManagement.contract.FilterLogs(opts, "MContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractManagementMContractUpdatedIterator{contract: _ContractManagement.contract, event: "MContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMContractUpdated is a free log subscription operation binding the contract event 0x2752a706b121abafb4dd58bef7d110189317241e7578a771ac0cfa40ce58b58f.
//
// Solidity: event MContractUpdated(uint256 id)
func (_ContractManagement *ContractManagementFilterer) WatchMContractUpdated(opts *bind.WatchOpts, sink chan<- *ContractManagementMContractUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractManagement.contract.WatchLogs(opts, "MContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractManagementMContractUpdated)
				if err := _ContractManagement.contract.UnpackLog(event, "MContractUpdated", log); err != nil {
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

// ParseMContractUpdated is a log parse operation binding the contract event 0x2752a706b121abafb4dd58bef7d110189317241e7578a771ac0cfa40ce58b58f.
//
// Solidity: event MContractUpdated(uint256 id)
func (_ContractManagement *ContractManagementFilterer) ParseMContractUpdated(log types.Log) (*ContractManagementMContractUpdated, error) {
	event := new(ContractManagementMContractUpdated)
	if err := _ContractManagement.contract.UnpackLog(event, "MContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
