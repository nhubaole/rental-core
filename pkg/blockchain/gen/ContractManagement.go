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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"landlord\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tenant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beginDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"MContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"MContractDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"MContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_landlord\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tenant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_roomId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beginDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_paymentMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electricityMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_electricityCost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_waterMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_waterCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_internetCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_parkingFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_responsibilityA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_responsibilityB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_generalResponsibility\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signatureA\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_signedTimeA\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_signatureB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_signedTimeB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_contractTemplateId\",\"type\":\"uint256\"}],\"name\":\"createMContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumContractManagement.PreRentalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.RentalProcessStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.PostRentalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"landlord\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tenant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roomId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beginDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"paymentMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electricityMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"electricityCost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"waterMethod\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"waterCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"internetCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parkingFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"responsibilityA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"responsibilityB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"generalResponsibility\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatureA\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signedTimeA\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signatureB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signedTimeB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractTemplateId\",\"type\":\"uint256\"},{\"internalType\":\"enumContractManagement.PreRentalStatus\",\"name\":\"preRentalStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.RentalProcessStatus\",\"name\":\"rentalProcessStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumContractManagement.PostRentalStatus\",\"name\":\"postRentalStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611789908161001a8239f35b5f80fdfe610500806040526004361015610013575f80fd5b5f3560e01c90816322f3bd3b1461123957508063a5244e6414610e935763d6d433db1461003e575f80fd5b34610e8f57610300366003190112610e8f576024356001600160401b038111610e8f5761006f9036906004016116fe565b610124356001600160401b038111610e8f5761008f9036906004016116fe565b610144356001600160401b038111610e8f576100af9036906004016116fe565b610184356001600160401b038111610e8f576100cf9036906004016116fe565b610204356001600160401b038111610e8f576100ef9036906004016116fe565b610224356001600160401b038111610e8f5761010f9036906004016116fe565b610244356001600160401b038111610e8f5761012f9036906004016116fe565b91610264356001600160401b038111610e8f576101509036906004016116fe565b936102a4356001600160401b038111610e8f576101719036906004016116fe565b956004355f525f60205260ff601b60405f20015416610e4a57604051976103c089018981106001600160401b0382111761085a5760405260043589528960208a015260443560408a015260643560608a015260843560808a015260a43560a08a015260c43560c08a015260e43560e08a0152610104356101008a0152610120890152610140880152610164356101608801526101808701526101a4356101a08701526101c4356101c08701526101e4356101e0870152610200860152610220850152610240840152610260830152610284356102808301526102a08201526102c4356102c08201526102e4356102e08201525f6103008201525f6103208201525f610340820152426103608201524261038082015260016103a08201526004355f525f60205260405f20908051825560208101519283516001600160401b03811161085a576102c36001850154611605565b601f8111610e06575b506020601f8211600114610d9c5781929394955f92610d91575b50508160011b915f199060031b1c19161760018401555b60408201516002840155606082015160038401556080820151600484015560a0820151600584015560c0820151600684015560e0820151600784015561010082015160088401556101208201519283516001600160401b03811161085a576103686009830154611605565b601f8111610d4d575b506020601f8211600114610ce35781929394955f92610cd8575b50508160011b915f199060031b1c19161760098201555b6101408301519283516001600160401b03811161085a576103c6600a840154611605565b601f8111610c94575b506020601f8211600114610c2a5781929394955f92610c1f575b50508160011b915f199060031b1c191617600a8301555b610160810151600b8301556101808101519283516001600160401b03811161085a5761042f600c850154611605565b601f8111610bdb575b506020601f8211600114610b715781929394955f92610b66575b50508160011b915f199060031b1c191617600c8401555b6101a0820151600d8401556101c0820151600e8401556101e0820151600f840155601083016102008301518051906001600160401b03821161085a5781906104b18454611605565b601f8111610b16575b50602090601f8311600114610ab3575f92610aa8575b50508160011b915f199060031b1c19161790555b601183016102208301518051906001600160401b03821161085a57819061050b8454611605565b601f8111610a58575b50602090601f83116001146109f5575f926109ea575b50508160011b915f199060031b1c19161790555b601283016102408301518051906001600160401b03821161085a5781906105658454611605565b601f811161099a575b50602090601f8311600114610937575f9261092c575b50508160011b915f199060031b1c19161790555b601383016102608301518051906001600160401b03821161085a5781906105bf8454611605565b601f81116108dc575b50602090601f8311600114610879575f9261086e575b50508160011b915f199060031b1c19161790555b6102808201516014840155601583016102a08301518051906001600160401b03821161085a5781906106248454611605565b601f811161080a575b50602090601f83116001146107a7575f9261079c575b50508160011b915f199060031b1c19161790555b6102c082015160168401556102e082015160178401556103008201519161067d836115ee565b610686836115ee565b6018840154610320820151906004821015610788577fada3db4246fb7379c45e0f4bad05d62f46c69cbc3b3c774d8e9aa0f94c98af9495610749956103a09361ff00601b9460ff62ff00006103408a01516106e0816115ee565b6106e9816115ee565b60101b1694169062ffffff1916179160081b16171760188201556103608401516019820155610380840151601a82015501910151151560ff80198354169116179055604051918291600435835261012060208401526101208301906115ca565b60443560408301526064356060830152608435608083015260a43560a083015260c43560c083015260e43560e0830152610104356101008301520390a1005b634e487b7160e01b5f52602160045260245ffd5b015190505f80610643565b5f8581528281209350601f198516905b8181106107f257509084600195949392106107da575b505050811b019055610657565b01515f1960f88460031b161c191690555f80806107cd565b929360206001819287860151815501950193016107b7565b909150835f5260205f20601f840160051c81019160208510610850575b90601f859493920160051c01905b818110610842575061062d565b5f8155849350600101610835565b9091508190610827565b634e487b7160e01b5f52604160045260245ffd5b015190505f806105de565b5f8581528281209350601f198516905b8181106108c457509084600195949392106108ac575b505050811b0190556105f2565b01515f1960f88460031b161c191690555f808061089f565b92936020600181928786015181550195019301610889565b909150835f5260205f20601f840160051c81019160208510610922575b90601f859493920160051c01905b81811061091457506105c8565b5f8155849350600101610907565b90915081906108f9565b015190505f80610584565b5f8581528281209350601f198516905b818110610982575090846001959493921061096a575b505050811b019055610598565b01515f1960f88460031b161c191690555f808061095d565b92936020600181928786015181550195019301610947565b909150835f5260205f20601f840160051c810191602085106109e0575b90601f859493920160051c01905b8181106109d2575061056e565b5f81558493506001016109c5565b90915081906109b7565b015190505f8061052a565b5f8581528281209350601f198516905b818110610a405750908460019594939210610a28575b505050811b01905561053e565b01515f1960f88460031b161c191690555f8080610a1b565b92936020600181928786015181550195019301610a05565b909150835f5260205f20601f840160051c81019160208510610a9e575b90601f859493920160051c01905b818110610a905750610514565b5f8155849350600101610a83565b9091508190610a75565b015190505f806104d0565b5f8581528281209350601f198516905b818110610afe5750908460019594939210610ae6575b505050811b0190556104e4565b01515f1960f88460031b161c191690555f8080610ad9565b92936020600181928786015181550195019301610ac3565b909150835f5260205f20601f840160051c81019160208510610b5c575b90601f859493920160051c01905b818110610b4e57506104ba565b5f8155849350600101610b41565b9091508190610b33565b015190505f80610452565b600c85015f52805f20905f5b601f1984168110610bc3575060019394959683601f19811610610bab575b505050811b01600c840155610469565b01515f1960f88460031b161c191690555f8080610b9b565b9091602060018192858b015181550193019101610b7d565b600c85015f5260205f20601f830160051c810160208410610c18575b601f830160051c82018110610c0d575050610438565b5f8155600101610bf7565b5080610bf7565b015190505f806103e9565b600a84015f52805f20905f5b601f1984168110610c7c575060019394959683601f19811610610c64575b505050811b01600a830155610400565b01515f1960f88460031b161c191690555f8080610c54565b9091602060018192858b015181550193019101610c36565b600a84015f5260205f20601f830160051c810160208410610cd1575b601f830160051c82018110610cc65750506103cf565b5f8155600101610cb0565b5080610cb0565b015190505f8061038b565b600983015f52805f20905f5b601f1984168110610d35575060019394959683601f19811610610d1d575b505050811b0160098201556103a2565b01515f1960f88460031b161c191690555f8080610d0d565b9091602060018192858b015181550193019101610cef565b600983015f5260205f20601f830160051c810160208410610d8a575b601f830160051c82018110610d7f575050610371565b5f8155600101610d69565b5080610d69565b015190505f806102e6565b600185015f52805f20905f5b601f1984168110610dee575060019394959683601f19811610610dd6575b505050811b0160018401556102fd565b01515f1960f88460031b161c191690555f8080610dc6565b9091602060018192858b015181550193019101610da8565b600185015f5260205f20601f830160051c810160208410610e43575b601f830160051c82018110610e385750506102cc565b5f8155600101610e22565b5080610e22565b60405162461bcd60e51b815260206004820152601860248201527f4d436f6e747261637420616c72656164792065786973747300000000000000006044820152606490fd5b5f80fd5b34610e8f576020366003190112610e8f576004355f525f60205260405f206104e05261118e6111696111516111396111216104e05154610ed860016104e0510161165e565b9060026104e05101546103205260036104e05101546103405260046104e0510154610480526110e960056104e051015460066104e05101546110cb60076104e0510154926110b660086104e051015491610f3760096104e0510161165e565b610f46600a6104e0510161165e565b93600b6104e051015497610f5f600c6104e0510161165e565b97600d6104e05101546103e05261105c600e6104e05101549d600f6104e05101549d610f9060106104e0510161165e565b61040052610fa360116104e0510161165e565b61042052610fb660126104e0510161165e565b61044052610fc960136104e0510161165e565b6104c05260146104e05101546104a052610fe860156104e0510161165e565b6104605260166104e05101546103005260176104e05101546103c05260186104e05101546103a05260196104e051015461038052601a6104e05101546103605260ff601b6104e0510154166102e0526040516102c0526102c051526103c060206102c05101526103c06102c05101906115ca565b946103205160406102c05101526103405160606102c05101526104805160806102c051015260a06102c051015260c06102c051015260e06102c05101526101006102c05101526102c05182036101206102c05101526115ca565b906102c05182036101406102c05101526115ca565b916101606102c05101526102c05182036101806102c05101526115ca565b916103e0516101a06102c05101526101c06102c05101526101e06102c05101526102c05181036102006102c0510152610400516115ca565b6102c05181036102206102c0510152610420516115ca565b6102c05181036102406102c0510152610440516115ca565b6102c05181036102606102c05101526104c0516115ca565b6104a0516102806102c05101526102c05181036102a06102c0510152610460516115ca565b610300516102c0805101526103c0516102e06102c05101526111b560ff6103a051166115ee565b60ff6103a051166103006102c05101526111df6103206102c0510160ff6103a05160081c166115f8565b6111f160ff6103a05160101c166115ee565b60ff6103a05160101c166103406102c0510152610380516103606102c0510152610360516103806102c05101526102e05115156103a06102c05101526102c05190036102c051f35b34610e8f576020366003190112610e8f576004355f525f60205260405f2061022052600160ff601b61022051015416151503611588576102205180546102a05260028101546101805260038101546101a052600481015461020052600581015460068201546007830154600884015461028052600b840154600d850154600e860154600f87015460148801546101e052601688015461016052601788015461012052601888015461024052601988015461026052601a880154610140526114f3976114d1976114bb976114a697611491979495946114639490939192611448929091611435919061132c9060010161165e565b9461133c6009610220510161165e565b9061134c600a610220510161165e565b946113e161135f600c610220510161165e565b9861136f6010610220510161165e565b6080526113816011610220510161165e565b60a0526113936012610220510161165e565b60e0526113a56013610220510161165e565b6101c0526113b86015610220510161165e565b6101005260405160c0526102a05160c051526103a0602060c05101526103a060c05101906115ca565b9361018051604060c05101526101a051606060c051015261020051608060c051015260a060c051015260c08051015260e060c05101526102805161010060c051015260c051820361012060c05101526115ca565b9060c051820361014060c05101526115ca565b9161016060c051015260c051820361018060c05101526115ca565b926101a060c05101526101c060c05101526101e060c051015260c051810361020060c05101526080516115ca565b60c051810361022060c051015260a0516115ca565b60c051810361024060c051015260e0516115ca565b60c051810361026060c05101526101c0516115ca565b6101e05161028060c051015260c05181036102a060c0510152610100516115ca565b610160516102c060c0510152610120516102e060c051015261151a60ff61024051166115ee565b60ff610240511661030060c051015261154261032060c0510160ff6102405160081c166115f8565b61155460ff6102405160101c166115ee565b60ff6102405160101c1661034060c05101526102605161036060c05101526101405161038060c051015260c051900360c051f35b62461bcd60e51b815260206004820152601860248201527f4d436f6e747261637420646f6573206e6f7420657869737400000000000000006044820152606490fd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b6003111561078857565b9060048210156107885752565b90600182811c92168015611633575b602083101461161f57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611614565b90601f801991011681019081106001600160401b0382111761085a57604052565b9060405191825f82549261167184611605565b80845293600181169081156116dc5750600114611698575b506116969250038361163d565b565b90505f9291925260205f20905f915b8183106116c0575050906020611696928201015f611689565b60209193508060019154838589010152019101909184926116a7565b90506020925061169694915060ff191682840152151560051b8201015f611689565b81601f82011215610e8f578035906001600160401b03821161085a5760405192611732601f8401601f19166020018561163d565b82845260208383010111610e8f57815f92602080930183860137830101529056fea264697066735822122004ccd092247f4af1297165a07b064599c0f83b05ffdd6d1111f4da1824ad02cc64736f6c634300081c0033",
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
