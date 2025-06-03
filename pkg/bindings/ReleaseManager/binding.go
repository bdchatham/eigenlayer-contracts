// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReleaseManager

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

// ReleaseManagerMetaData contains all meta data concerning the ReleaseManager contract.
var ReleaseManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allPublishedReleases\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deprecateRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecatedReleases\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeprecatedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublishedReleases\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManager.PublishedRelease[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseAtBlock\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManager.PublishedRelease\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publishedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReleaseCheckpointCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDeprecated\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSDeregistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistered\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleaseDeprecated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"deploymentDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AVSAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AVSNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyDeprecated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDigest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060015f55611cb6806100205f395ff3fe608060405234801561000f575f5ffd5b506004361061011c575f3560e01c806384ac33ec116100a9578063d64b79dc1161006e578063d64b79dc146102a3578063e07955cf146102b6578063f2753b1b1461031d578063f2fde38b14610330578063f330132c14610343575f5ffd5b806384ac33ec146102375780638da5cb5b1461024a578063b8861afe1461025b578063bf2d8e071461026e578063c4d66de814610290575f5ffd5b80634657e26a116100ef5780634657e26a146101ac57806355bac3ca146101d7578063715018a6146101f85780637754737d146102005780637e149b0914610224575f5ffd5b80630de4d5c01461012057806317d46c96146101575780632e924e8c146101775780634420e48614610197575b5f5ffd5b61014261012e36600461174b565b606a6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61016a61016536600461177d565b610363565b60405161014e9190611796565b61018a6101853660046117d8565b6103cc565b60405161014e9190611886565b6101aa6101a536600461177d565b610630565b005b6066546101bf906001600160a01b031681565b6040516001600160a01b03909116815260200161014e565b6101ea6101e536600461177d565b6106dc565b60405190815260200161014e565b6101aa6106f9565b61021361020e3660046117d8565b61070c565b60405161014e959493929190611898565b6101ea6102323660046117d8565b610864565b6101aa61024536600461177d565b61088f565b6034546001600160a01b03166101bf565b6101aa61026936600461191c565b610937565b61014261027c36600461177d565b60676020525f908152604090205460ff1681565b6101aa61029e36600461177d565b610bc2565b6101aa6102b13660046117d8565b610cee565b6101426102c43660046117d8565b6040516bffffffffffffffffffffffff19606084901b166020820152603481018290525f90819060540160408051808303601f1901815291815281516020928301205f908152606a90925290205460ff16949350505050565b61018a61032b36600461177d565b610e59565b6101aa61033e36600461177d565b611080565b61035661035136600461177d565b6110f9565b60405161014e91906119ad565b6001600160a01b0381165f908152606b60209081526040918290208054835181840281018401909452808452606093928301828280156103c057602002820191905f5260205f20905b8154815260200190600101908083116103ac575b50505050509050919050565b6103fb6040518060a001604052805f815260200160608152602001606081526020015f81526020015f81525090565b6001600160a01b0383165f90815260696020526040812061041c90846112a8565b6001600160e01b031690505f60685f866001600160a01b03166001600160a01b031681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156105d7578382905f5260205f2090600502016040518060a00160405290815f82015481526020016001820180546104a490611a10565b80601f01602080910402602001604051908101604052809291908181526020018280546104d090611a10565b801561051b5780601f106104f25761010080835404028352916020019161051b565b820191905f5260205f20905b8154815290600101906020018083116104fe57829003601f168201915b5050505050815260200160028201805461053490611a10565b80601f016020809104026020016040519081016040528092919081815260200182805461056090611a10565b80156105ab5780601f10610582576101008083540402835291602001916105ab565b820191905f5260205f20905b81548152906001019060200180831161058e57829003601f168201915b50505050508152602001600382015481526020016004820154815250508152602001906001019061046a565b50505050905080515f14806105ed575080518210155b1561060b5760405163050cc7ff60e31b815260040160405180910390fd5b80828151811061061d5761061d611a48565b6020026020010151925050505b92915050565b8061063b81336112ff565b610657576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0382165f9081526067602052604090205460ff16156106905760405163886f069560e01b815260040160405180910390fd5b6001600160a01b0382165f81815260676020526040808220805460ff19166001179055517f2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e9190a25050565b6001600160a01b0381165f9081526069602052604081205461062a565b6107016113b5565b61070a5f61140f565b565b6068602052815f5260405f208181548110610725575f80fd5b905f5260205f2090600502015f9150915050805f01549080600101805461074b90611a10565b80601f016020809104026020016040519081016040528092919081815260200182805461077790611a10565b80156107c25780601f10610799576101008083540402835291602001916107c2565b820191905f5260205f20905b8154815290600101906020018083116107a557829003601f168201915b5050505050908060020180546107d790611a10565b80601f016020809104026020016040519081016040528092919081815260200182805461080390611a10565b801561084e5780601f106108255761010080835404028352916020019161084e565b820191905f5260205f20905b81548152906001019060200180831161083157829003601f168201915b5050505050908060030154908060040154905085565b606b602052815f5260405f20818154811061087d575f80fd5b905f5260205f20015f91509150505481565b8061089a81336112ff565b6108b6576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0382165f9081526067602052604090205460ff166108ee57604051635ae33df360e11b815260040160405180910390fd5b6001600160a01b0382165f81815260676020526040808220805460ff19169055517ff7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab79190a25050565b61093f611460565b6001600160a01b0387165f90815260676020526040902054879060ff1661097957604051635ae33df360e11b815260040160405180910390fd5b8761098481336112ff565b6109a0576040516282b42960e81b815260040160405180910390fd5b825f036109c057604051631da7447960e21b815260040160405180910390fd5b876109de57604051633e78eb9360e01b815260040160405180910390fd5b5f8690036109ff57604051633e78eb9360e01b815260040160405180910390fd5b5f6040518060a001604052808a815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284375f9201829052509385525050506020808301889052426040938401526001600160a01b038e168252606881529181208054600181810183559183529183902084516005909302019182559183015192935083929091820190610ad79082611abc565b5060408201516002820190610aec9082611abc565b50606082015160038201556080909101516004909101556001600160a01b038a165f90815260686020526040812054610b2790600190611b8b565b6001600160a01b038c165f908152606960205260409020909150610b4c9043836114b7565b5050898787604051610b5f929190611b9e565b60405180910390208c6001600160a01b03167f5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd28c8c8a604051610ba493929190611bad565b60405180910390a450505050610bb960015f55565b50505050505050565b600154610100900460ff1615808015610bdf57506001805460ff16105b80610bf85750303b158015610bf857506001805460ff16145b610c605760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6001805460ff1916811790558015610c82576001805461ff0019166101001790555b610c8a6114d1565b606680546001600160a01b0319166001600160a01b0384161790558015610cea576001805461ff00191681556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b610cf6611460565b6001600160a01b0382165f90815260676020526040902054829060ff16610d3057604051635ae33df360e11b815260040160405180910390fd5b82610d3b81336112ff565b610d57576040516282b42960e81b815260040160405180910390fd5b82610d7557604051633e78eb9360e01b815260040160405180910390fd5b6040516bffffffffffffffffffffffff19606086901b166020820152603481018490525f9060540160408051601f1981840301815291815281516020928301205f818152606a90935291205490915060ff1615610de55760405163764c143b60e01b815260040160405180910390fd5b5f818152606a60209081526040808320805460ff191660019081179091556001600160a01b038916808552606b8452828520805492830181558552928420018790555186927f4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b491a3505050610cea60015f55565b610e886040518060a001604052805f815260200160608152602001606081526020015f81526020015f81525090565b6001600160a01b0382165f90815260686020908152604080832080548251818502810185019093528083529192909190849084015b8282101561102a578382905f5260205f2090600502016040518060a00160405290815f8201548152602001600182018054610ef790611a10565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2390611a10565b8015610f6e5780601f10610f4557610100808354040283529160200191610f6e565b820191905f5260205f20905b815481529060010190602001808311610f5157829003601f168201915b50505050508152602001600282018054610f8790611a10565b80601f0160208091040260200160405190810160405280929190818152602001828054610fb390611a10565b8015610ffe5780601f10610fd557610100808354040283529160200191610ffe565b820191905f5260205f20905b815481529060010190602001808311610fe157829003601f168201915b505050505081526020016003820154815260200160048201548152505081526020019060010190610ebd565b50505050905080515f036110515760405163050cc7ff60e31b815260040160405180910390fd5b80600182516110609190611b8b565b8151811061107057611070611a48565b6020026020010151915050919050565b6110886113b5565b6001600160a01b0381166110ed5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610c57565b6110f68161140f565b50565b6001600160a01b0381165f908152606860209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561129d578382905f5260205f2090600502016040518060a00160405290815f820154815260200160018201805461116a90611a10565b80601f016020809104026020016040519081016040528092919081815260200182805461119690611a10565b80156111e15780601f106111b8576101008083540402835291602001916111e1565b820191905f5260205f20905b8154815290600101906020018083116111c457829003601f168201915b505050505081526020016002820180546111fa90611a10565b80601f016020809104026020016040519081016040528092919081815260200182805461122690611a10565b80156112715780601f1061124857610100808354040283529160200191611271565b820191905f5260205f20905b81548152906001019060200180831161125457829003601f168201915b505050505081526020016003820154815260200160048201548152505081526020019060010190611130565b505050509050919050565b81545f90816112b985858385611500565b905080156112f4576112dd856112d0600184611b8b565b5f91825260209091200190565b5464010000000090046001600160e01b03166112f6565b5f5b95945050505050565b5f816001600160a01b0316836001600160a01b0316036113215750600161062a565b6066546001600160a01b0316156113ad57606654604051639100674560e01b81526001600160a01b038581166004830152848116602483015290911690639100674590604401602060405180830381865afa158015611382573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113a69190611be4565b905061062a565b505f92915050565b6034546001600160a01b0316331461070a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c57565b603480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60025f54036114b15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610c57565b60025f55565b5f806114c485858561155b565b915091505b935093915050565b600154610100900460ff166114f85760405162461bcd60e51b8152600401610c5790611c03565b61070a6116fa565b5f5b81831015611553575f611515848461172a565b5f8781526020902090915063ffffffff86169082015463ffffffff16111561153f5780925061154d565b61154a816001611c4e565b93505b50611502565b509392505050565b82545f90819080156116a1575f611577876112d0600185611b8b565b60408051808201909152905463ffffffff8082168084526401000000009092046001600160e01b0316602084015291925090871610156115f95760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b65797300000000006044820152606401610c57565b805163ffffffff808816911603611642578461161a886112d0600186611b8b565b80546001600160e01b03929092166401000000000263ffffffff909216919091179055611691565b6040805180820190915263ffffffff80881682526001600160e01b0380881660208085019182528b54600181018d555f8d81529190912094519151909216640100000000029216919091179101555b6020015192508391506114c99050565b50506040805180820190915263ffffffff80851682526001600160e01b0380851660208085019182528854600181018a555f8a8152918220955192519093166401000000000291909316179201919091559050816114c9565b600154610100900460ff166117215760405162461bcd60e51b8152600401610c5790611c03565b61070a3361140f565b5f6117386002848418611c61565b61174490848416611c4e565b9392505050565b5f6020828403121561175b575f5ffd5b5035919050565b80356001600160a01b0381168114611778575f5ffd5b919050565b5f6020828403121561178d575f5ffd5b61174482611762565b602080825282518282018190525f918401906040840190835b818110156117cd5783518352602093840193909201916001016117af565b509095945050505050565b5f5f604083850312156117e9575f5ffd5b6117f283611762565b946020939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b805182525f602082015160a0602085015261184c60a0850182611800565b9050604083015184820360408601526118658282611800565b91505060608301516060850152608083015160808501528091505092915050565b602081525f611744602083018461182e565b85815260a060208201525f6118b060a0830187611800565b82810360408401526118c28187611800565b60608401959095525050608001529392505050565b5f5f83601f8401126118e7575f5ffd5b50813567ffffffffffffffff8111156118fe575f5ffd5b602083019150836020828501011115611915575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a031215611932575f5ffd5b61193b88611762565b965060208801359550604088013567ffffffffffffffff81111561195d575f5ffd5b6119698a828b016118d7565b909650945050606088013567ffffffffffffffff811115611988575f5ffd5b6119948a828b016118d7565b989b979a50959894979596608090950135949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611a0457603f198786030184526119ef85835161182e565b945060209384019391909101906001016119d3565b50929695505050505050565b600181811c90821680611a2457607f821691505b602082108103611a4257634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b601f821115611ab757805f5260205f20601f840160051c81016020851015611a955750805b601f840160051c820191505b81811015611ab4575f8155600101611aa1565b50505b505050565b815167ffffffffffffffff811115611ad657611ad6611a5c565b611aea81611ae48454611a10565b84611a70565b6020601f821160018114611b1c575f8315611b055750848201515b5f19600385901b1c1916600184901b178455611ab4565b5f84815260208120601f198516915b82811015611b4b5787850151825560209485019460019092019101611b2b565b5084821015611b6857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561062a5761062a611b77565b818382375f9101908152919050565b60408152826040820152828460608301375f606084830101525f6060601f19601f8601168301019050826020830152949350505050565b5f60208284031215611bf4575f5ffd5b81518015158114611744575f5ffd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b8082018082111561062a5761062a611b77565b5f82611c7b57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220d25ae3da7779a50c57cf4b245c2b631e787f9628f8f47f1011ea07d949898fc364736f6c634300081b0033",
}

// ReleaseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReleaseManagerMetaData.ABI instead.
var ReleaseManagerABI = ReleaseManagerMetaData.ABI

// ReleaseManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReleaseManagerMetaData.Bin instead.
var ReleaseManagerBin = ReleaseManagerMetaData.Bin

// DeployReleaseManager deploys a new Ethereum contract, binding an instance of ReleaseManager to it.
func DeployReleaseManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReleaseManager, error) {
	parsed, err := ReleaseManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReleaseManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReleaseManager{ReleaseManagerCaller: ReleaseManagerCaller{contract: contract}, ReleaseManagerTransactor: ReleaseManagerTransactor{contract: contract}, ReleaseManagerFilterer: ReleaseManagerFilterer{contract: contract}}, nil
}

// ReleaseManager is an auto generated Go binding around an Ethereum contract.
type ReleaseManager struct {
	ReleaseManagerCaller     // Read-only binding to the contract
	ReleaseManagerTransactor // Write-only binding to the contract
	ReleaseManagerFilterer   // Log filterer for contract events
}

// ReleaseManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseManagerSession struct {
	Contract     *ReleaseManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReleaseManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseManagerCallerSession struct {
	Contract *ReleaseManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReleaseManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseManagerTransactorSession struct {
	Contract     *ReleaseManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReleaseManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseManagerRaw struct {
	Contract *ReleaseManager // Generic contract binding to access the raw methods on
}

// ReleaseManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseManagerCallerRaw struct {
	Contract *ReleaseManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseManagerTransactorRaw struct {
	Contract *ReleaseManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseManager creates a new instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManager(address common.Address, backend bind.ContractBackend) (*ReleaseManager, error) {
	contract, err := bindReleaseManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseManager{ReleaseManagerCaller: ReleaseManagerCaller{contract: contract}, ReleaseManagerTransactor: ReleaseManagerTransactor{contract: contract}, ReleaseManagerFilterer: ReleaseManagerFilterer{contract: contract}}, nil
}

// NewReleaseManagerCaller creates a new read-only instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerCaller(address common.Address, caller bind.ContractCaller) (*ReleaseManagerCaller, error) {
	contract, err := bindReleaseManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerCaller{contract: contract}, nil
}

// NewReleaseManagerTransactor creates a new write-only instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseManagerTransactor, error) {
	contract, err := bindReleaseManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerTransactor{contract: contract}, nil
}

// NewReleaseManagerFilterer creates a new log filterer instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseManagerFilterer, error) {
	contract, err := bindReleaseManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerFilterer{contract: contract}, nil
}

// bindReleaseManager binds a generic wrapper to an already deployed contract.
func bindReleaseManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReleaseManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManager *ReleaseManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManager.Contract.ReleaseManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManager *ReleaseManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManager.Contract.ReleaseManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManager *ReleaseManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManager.Contract.ReleaseManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManager *ReleaseManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManager *ReleaseManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManager *ReleaseManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManager.Contract.contract.Transact(opts, method, params...)
}

// AllPublishedReleases is a free data retrieval call binding the contract method 0x7754737d.
//
// Solidity: function allPublishedReleases(address , uint256 ) view returns(bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline, uint256 publishedAt)
func (_ReleaseManager *ReleaseManagerCaller) AllPublishedReleases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "allPublishedReleases", arg0, arg1)

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
func (_ReleaseManager *ReleaseManagerSession) AllPublishedReleases(arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	return _ReleaseManager.Contract.AllPublishedReleases(&_ReleaseManager.CallOpts, arg0, arg1)
}

// AllPublishedReleases is a free data retrieval call binding the contract method 0x7754737d.
//
// Solidity: function allPublishedReleases(address , uint256 ) view returns(bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline, uint256 publishedAt)
func (_ReleaseManager *ReleaseManagerCallerSession) AllPublishedReleases(arg0 common.Address, arg1 *big.Int) (struct {
	Digest             [32]byte
	RegistryUrl        string
	Version            string
	DeploymentDeadline *big.Int
	PublishedAt        *big.Int
}, error) {
	return _ReleaseManager.Contract.AllPublishedReleases(&_ReleaseManager.CallOpts, arg0, arg1)
}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManager *ReleaseManagerCaller) DeprecatedReleases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "deprecatedReleases", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManager *ReleaseManagerSession) DeprecatedReleases(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ReleaseManager.Contract.DeprecatedReleases(&_ReleaseManager.CallOpts, arg0, arg1)
}

// DeprecatedReleases is a free data retrieval call binding the contract method 0x7e149b09.
//
// Solidity: function deprecatedReleases(address , uint256 ) view returns(bytes32)
func (_ReleaseManager *ReleaseManagerCallerSession) DeprecatedReleases(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ReleaseManager.Contract.DeprecatedReleases(&_ReleaseManager.CallOpts, arg0, arg1)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManager *ReleaseManagerCaller) GetDeprecatedReleases(opts *bind.CallOpts, avs common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getDeprecatedReleases", avs)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManager *ReleaseManagerSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _ReleaseManager.Contract.GetDeprecatedReleases(&_ReleaseManager.CallOpts, avs)
}

// GetDeprecatedReleases is a free data retrieval call binding the contract method 0x17d46c96.
//
// Solidity: function getDeprecatedReleases(address avs) view returns(bytes32[])
func (_ReleaseManager *ReleaseManagerCallerSession) GetDeprecatedReleases(avs common.Address) ([][32]byte, error) {
	return _ReleaseManager.Contract.GetDeprecatedReleases(&_ReleaseManager.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerCaller) GetLatestRelease(opts *bind.CallOpts, avs common.Address) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getLatestRelease", avs)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetLatestRelease(&_ReleaseManager.CallOpts, avs)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xf2753b1b.
//
// Solidity: function getLatestRelease(address avs) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerCallerSession) GetLatestRelease(avs common.Address) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetLatestRelease(&_ReleaseManager.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManager *ReleaseManagerCaller) GetPublishedReleases(opts *bind.CallOpts, avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getPublishedReleases", avs)

	if err != nil {
		return *new([]IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new([]IReleaseManagerPublishedRelease)).(*[]IReleaseManagerPublishedRelease)

	return out0, err

}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManager *ReleaseManagerSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetPublishedReleases(&_ReleaseManager.CallOpts, avs)
}

// GetPublishedReleases is a free data retrieval call binding the contract method 0xf330132c.
//
// Solidity: function getPublishedReleases(address avs) view returns((bytes32,string,string,uint256,uint256)[])
func (_ReleaseManager *ReleaseManagerCallerSession) GetPublishedReleases(avs common.Address) ([]IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetPublishedReleases(&_ReleaseManager.CallOpts, avs)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerCaller) GetReleaseAtBlock(opts *bind.CallOpts, avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getReleaseAtBlock", avs, blockNumber)

	if err != nil {
		return *new(IReleaseManagerPublishedRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerPublishedRelease)).(*IReleaseManagerPublishedRelease)

	return out0, err

}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetReleaseAtBlock(&_ReleaseManager.CallOpts, avs, blockNumber)
}

// GetReleaseAtBlock is a free data retrieval call binding the contract method 0x2e924e8c.
//
// Solidity: function getReleaseAtBlock(address avs, uint256 blockNumber) view returns((bytes32,string,string,uint256,uint256))
func (_ReleaseManager *ReleaseManagerCallerSession) GetReleaseAtBlock(avs common.Address, blockNumber *big.Int) (IReleaseManagerPublishedRelease, error) {
	return _ReleaseManager.Contract.GetReleaseAtBlock(&_ReleaseManager.CallOpts, avs, blockNumber)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManager *ReleaseManagerCaller) GetReleaseCheckpointCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getReleaseCheckpointCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManager *ReleaseManagerSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _ReleaseManager.Contract.GetReleaseCheckpointCount(&_ReleaseManager.CallOpts, avs)
}

// GetReleaseCheckpointCount is a free data retrieval call binding the contract method 0x55bac3ca.
//
// Solidity: function getReleaseCheckpointCount(address avs) view returns(uint256)
func (_ReleaseManager *ReleaseManagerCallerSession) GetReleaseCheckpointCount(avs common.Address) (*big.Int, error) {
	return _ReleaseManager.Contract.GetReleaseCheckpointCount(&_ReleaseManager.CallOpts, avs)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManager *ReleaseManagerCaller) IsDeprecated(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "isDeprecated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManager *ReleaseManagerSession) IsDeprecated(arg0 [32]byte) (bool, error) {
	return _ReleaseManager.Contract.IsDeprecated(&_ReleaseManager.CallOpts, arg0)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x0de4d5c0.
//
// Solidity: function isDeprecated(bytes32 ) view returns(bool)
func (_ReleaseManager *ReleaseManagerCallerSession) IsDeprecated(arg0 [32]byte) (bool, error) {
	return _ReleaseManager.Contract.IsDeprecated(&_ReleaseManager.CallOpts, arg0)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManager *ReleaseManagerCaller) IsReleaseDeprecated(opts *bind.CallOpts, avs common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "isReleaseDeprecated", avs, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManager *ReleaseManagerSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _ReleaseManager.Contract.IsReleaseDeprecated(&_ReleaseManager.CallOpts, avs, digest)
}

// IsReleaseDeprecated is a free data retrieval call binding the contract method 0xe07955cf.
//
// Solidity: function isReleaseDeprecated(address avs, bytes32 digest) view returns(bool)
func (_ReleaseManager *ReleaseManagerCallerSession) IsReleaseDeprecated(avs common.Address, digest [32]byte) (bool, error) {
	return _ReleaseManager.Contract.IsReleaseDeprecated(&_ReleaseManager.CallOpts, avs, digest)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReleaseManager *ReleaseManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReleaseManager *ReleaseManagerSession) Owner() (common.Address, error) {
	return _ReleaseManager.Contract.Owner(&_ReleaseManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReleaseManager *ReleaseManagerCallerSession) Owner() (common.Address, error) {
	return _ReleaseManager.Contract.Owner(&_ReleaseManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerSession) PermissionController() (common.Address, error) {
	return _ReleaseManager.Contract.PermissionController(&_ReleaseManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerCallerSession) PermissionController() (common.Address, error) {
	return _ReleaseManager.Contract.PermissionController(&_ReleaseManager.CallOpts)
}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManager *ReleaseManagerCaller) RegisteredAVS(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "registeredAVS", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManager *ReleaseManagerSession) RegisteredAVS(arg0 common.Address) (bool, error) {
	return _ReleaseManager.Contract.RegisteredAVS(&_ReleaseManager.CallOpts, arg0)
}

// RegisteredAVS is a free data retrieval call binding the contract method 0xbf2d8e07.
//
// Solidity: function registeredAVS(address ) view returns(bool)
func (_ReleaseManager *ReleaseManagerCallerSession) RegisteredAVS(arg0 common.Address) (bool, error) {
	return _ReleaseManager.Contract.RegisteredAVS(&_ReleaseManager.CallOpts, arg0)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManager *ReleaseManagerTransactor) DeprecateRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "deprecateRelease", avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManager *ReleaseManagerSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManager.Contract.DeprecateRelease(&_ReleaseManager.TransactOpts, avs, digest)
}

// DeprecateRelease is a paid mutator transaction binding the contract method 0xd64b79dc.
//
// Solidity: function deprecateRelease(address avs, bytes32 digest) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) DeprecateRelease(avs common.Address, digest [32]byte) (*types.Transaction, error) {
	return _ReleaseManager.Contract.DeprecateRelease(&_ReleaseManager.TransactOpts, avs, digest)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManager *ReleaseManagerTransactor) Deregister(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "deregister", avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManager *ReleaseManagerSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Deregister(&_ReleaseManager.TransactOpts, avs)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address avs) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) Deregister(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Deregister(&_ReleaseManager.TransactOpts, avs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionController) returns()
func (_ReleaseManager *ReleaseManagerTransactor) Initialize(opts *bind.TransactOpts, _permissionController common.Address) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "initialize", _permissionController)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionController) returns()
func (_ReleaseManager *ReleaseManagerSession) Initialize(_permissionController common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Initialize(&_ReleaseManager.TransactOpts, _permissionController)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionController) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) Initialize(_permissionController common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Initialize(&_ReleaseManager.TransactOpts, _permissionController)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManager *ReleaseManagerTransactor) PublishRelease(opts *bind.TransactOpts, avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "publishRelease", avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManager *ReleaseManagerSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishRelease(&_ReleaseManager.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// PublishRelease is a paid mutator transaction binding the contract method 0xb8861afe.
//
// Solidity: function publishRelease(address avs, bytes32 digest, string registryUrl, string version, uint256 deploymentDeadline) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) PublishRelease(avs common.Address, digest [32]byte, registryUrl string, version string, deploymentDeadline *big.Int) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishRelease(&_ReleaseManager.TransactOpts, avs, digest, registryUrl, version, deploymentDeadline)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManager *ReleaseManagerTransactor) Register(opts *bind.TransactOpts, avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "register", avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManager *ReleaseManagerSession) Register(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Register(&_ReleaseManager.TransactOpts, avs)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address avs) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) Register(avs common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.Register(&_ReleaseManager.TransactOpts, avs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseManager *ReleaseManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseManager *ReleaseManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReleaseManager.Contract.RenounceOwnership(&_ReleaseManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReleaseManager.Contract.RenounceOwnership(&_ReleaseManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseManager *ReleaseManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseManager *ReleaseManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.TransferOwnership(&_ReleaseManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseManager.Contract.TransferOwnership(&_ReleaseManager.TransactOpts, newOwner)
}

// ReleaseManagerAVSDeregisteredIterator is returned from FilterAVSDeregistered and is used to iterate over the raw logs and unpacked data for AVSDeregistered events raised by the ReleaseManager contract.
type ReleaseManagerAVSDeregisteredIterator struct {
	Event *ReleaseManagerAVSDeregistered // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerAVSDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerAVSDeregistered)
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
		it.Event = new(ReleaseManagerAVSDeregistered)
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
func (it *ReleaseManagerAVSDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerAVSDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerAVSDeregistered represents a AVSDeregistered event raised by the ReleaseManager contract.
type ReleaseManagerAVSDeregistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSDeregistered is a free log retrieval operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_ReleaseManager *ReleaseManagerFilterer) FilterAVSDeregistered(opts *bind.FilterOpts, avs []common.Address) (*ReleaseManagerAVSDeregisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerAVSDeregisteredIterator{contract: _ReleaseManager.contract, event: "AVSDeregistered", logs: logs, sub: sub}, nil
}

// WatchAVSDeregistered is a free log subscription operation binding the contract event 0xf7cd17cf5978e63a941e1b110c3afd213843bff041513266571af35a6cec8ab7.
//
// Solidity: event AVSDeregistered(address indexed avs)
func (_ReleaseManager *ReleaseManagerFilterer) WatchAVSDeregistered(opts *bind.WatchOpts, sink chan<- *ReleaseManagerAVSDeregistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "AVSDeregistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerAVSDeregistered)
				if err := _ReleaseManager.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
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
func (_ReleaseManager *ReleaseManagerFilterer) ParseAVSDeregistered(log types.Log) (*ReleaseManagerAVSDeregistered, error) {
	event := new(ReleaseManagerAVSDeregistered)
	if err := _ReleaseManager.contract.UnpackLog(event, "AVSDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerAVSRegisteredIterator is returned from FilterAVSRegistered and is used to iterate over the raw logs and unpacked data for AVSRegistered events raised by the ReleaseManager contract.
type ReleaseManagerAVSRegisteredIterator struct {
	Event *ReleaseManagerAVSRegistered // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerAVSRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerAVSRegistered)
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
		it.Event = new(ReleaseManagerAVSRegistered)
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
func (it *ReleaseManagerAVSRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerAVSRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerAVSRegistered represents a AVSRegistered event raised by the ReleaseManager contract.
type ReleaseManagerAVSRegistered struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistered is a free log retrieval operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_ReleaseManager *ReleaseManagerFilterer) FilterAVSRegistered(opts *bind.FilterOpts, avs []common.Address) (*ReleaseManagerAVSRegisteredIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerAVSRegisteredIterator{contract: _ReleaseManager.contract, event: "AVSRegistered", logs: logs, sub: sub}, nil
}

// WatchAVSRegistered is a free log subscription operation binding the contract event 0x2c7ccee1b83a57ffa52bfd71692c05a6b8b9dc9b1e73a6d25c78bab22a98b06e.
//
// Solidity: event AVSRegistered(address indexed avs)
func (_ReleaseManager *ReleaseManagerFilterer) WatchAVSRegistered(opts *bind.WatchOpts, sink chan<- *ReleaseManagerAVSRegistered, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "AVSRegistered", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerAVSRegistered)
				if err := _ReleaseManager.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
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
func (_ReleaseManager *ReleaseManagerFilterer) ParseAVSRegistered(log types.Log) (*ReleaseManagerAVSRegistered, error) {
	event := new(ReleaseManagerAVSRegistered)
	if err := _ReleaseManager.contract.UnpackLog(event, "AVSRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReleaseManager contract.
type ReleaseManagerInitializedIterator struct {
	Event *ReleaseManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerInitialized)
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
		it.Event = new(ReleaseManagerInitialized)
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
func (it *ReleaseManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerInitialized represents a Initialized event raised by the ReleaseManager contract.
type ReleaseManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReleaseManager *ReleaseManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReleaseManagerInitializedIterator, error) {

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerInitializedIterator{contract: _ReleaseManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReleaseManager *ReleaseManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReleaseManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerInitialized)
				if err := _ReleaseManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReleaseManager *ReleaseManagerFilterer) ParseInitialized(log types.Log) (*ReleaseManagerInitialized, error) {
	event := new(ReleaseManagerInitialized)
	if err := _ReleaseManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReleaseManager contract.
type ReleaseManagerOwnershipTransferredIterator struct {
	Event *ReleaseManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerOwnershipTransferred)
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
		it.Event = new(ReleaseManagerOwnershipTransferred)
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
func (it *ReleaseManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ReleaseManager contract.
type ReleaseManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReleaseManager *ReleaseManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReleaseManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerOwnershipTransferredIterator{contract: _ReleaseManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReleaseManager *ReleaseManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReleaseManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerOwnershipTransferred)
				if err := _ReleaseManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReleaseManager *ReleaseManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ReleaseManagerOwnershipTransferred, error) {
	event := new(ReleaseManagerOwnershipTransferred)
	if err := _ReleaseManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerReleaseDeprecatedIterator is returned from FilterReleaseDeprecated and is used to iterate over the raw logs and unpacked data for ReleaseDeprecated events raised by the ReleaseManager contract.
type ReleaseManagerReleaseDeprecatedIterator struct {
	Event *ReleaseManagerReleaseDeprecated // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerReleaseDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerReleaseDeprecated)
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
		it.Event = new(ReleaseManagerReleaseDeprecated)
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
func (it *ReleaseManagerReleaseDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerReleaseDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerReleaseDeprecated represents a ReleaseDeprecated event raised by the ReleaseManager contract.
type ReleaseManagerReleaseDeprecated struct {
	Avs    common.Address
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleaseDeprecated is a free log retrieval operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_ReleaseManager *ReleaseManagerFilterer) FilterReleaseDeprecated(opts *bind.FilterOpts, avs []common.Address, digest [][32]byte) (*ReleaseManagerReleaseDeprecatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerReleaseDeprecatedIterator{contract: _ReleaseManager.contract, event: "ReleaseDeprecated", logs: logs, sub: sub}, nil
}

// WatchReleaseDeprecated is a free log subscription operation binding the contract event 0x4311226575e8128ce22dbc1d402e92129fbb3abea101f9c5e559863f5c6786b4.
//
// Solidity: event ReleaseDeprecated(address indexed avs, bytes32 indexed digest)
func (_ReleaseManager *ReleaseManagerFilterer) WatchReleaseDeprecated(opts *bind.WatchOpts, sink chan<- *ReleaseManagerReleaseDeprecated, avs []common.Address, digest [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "ReleaseDeprecated", avsRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerReleaseDeprecated)
				if err := _ReleaseManager.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
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
func (_ReleaseManager *ReleaseManagerFilterer) ParseReleaseDeprecated(log types.Log) (*ReleaseManagerReleaseDeprecated, error) {
	event := new(ReleaseManagerReleaseDeprecated)
	if err := _ReleaseManager.contract.UnpackLog(event, "ReleaseDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the ReleaseManager contract.
type ReleaseManagerReleasePublishedIterator struct {
	Event *ReleaseManagerReleasePublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerReleasePublished)
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
		it.Event = new(ReleaseManagerReleasePublished)
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
func (it *ReleaseManagerReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerReleasePublished represents a ReleasePublished event raised by the ReleaseManager contract.
type ReleaseManagerReleasePublished struct {
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
func (_ReleaseManager *ReleaseManagerFilterer) FilterReleasePublished(opts *bind.FilterOpts, avs []common.Address, version []string, digest [][32]byte) (*ReleaseManagerReleasePublishedIterator, error) {

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

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerReleasePublishedIterator{contract: _ReleaseManager.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x5e59c6e9cba1f5ee2d1fcd429bf1060da4bb5b738b23b49d5419c860c218fcd2.
//
// Solidity: event ReleasePublished(address indexed avs, string indexed version, bytes32 indexed digest, string registryUrl, uint256 deploymentDeadline)
func (_ReleaseManager *ReleaseManagerFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerReleasePublished, avs []common.Address, version []string, digest [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "ReleasePublished", avsRule, versionRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerReleasePublished)
				if err := _ReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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
func (_ReleaseManager *ReleaseManagerFilterer) ParseReleasePublished(log types.Log) (*ReleaseManagerReleasePublished, error) {
	event := new(ReleaseManagerReleasePublished)
	if err := _ReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
