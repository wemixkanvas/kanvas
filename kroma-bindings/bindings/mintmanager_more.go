// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MintManagerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"minted\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"recipients\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1003,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"shareOf\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_uint256)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\",\"base\":\"t_address\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var MintManagerStorageLayout = new(solc.StorageLayout)

var MintManagerDeployedBin = "0x608060405234801561001057600080fd5b50600436106100df5760003560e01c80637eb118451161008c578063baee5ed411610066578063baee5ed4146101f0578063d1bc76a1146101f8578063e4fc6b6d1461020b578063f2fde38b1461021357600080fd5b80637eb11845146101bd5780638da5cb5b146101c757806398f1312e146101e557600080fd5b8063457c3977116100bd578063457c39771461016d5780634f02c42014610180578063715018a6146101b557600080fd5b80631249c58b146100e457806321e5e2c4146100ee5780632efd46d614610121575b600080fd5b6100ec610226565b005b61010e6100fc366004610b74565b60026020526000908152604090205481565b6040519081526020015b60405180910390f35b6101487f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610118565b6100ec61017b366004610b74565b610508565b6000546101a59074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610118565b6100ec6105b4565b61010e620186a081565b60005473ffffffffffffffffffffffffffffffffffffffff16610148565b61010e633b9aca0081565b6100ec6105c8565b610148610206366004610bb1565b610652565b6100ec610689565b6100ec610221366004610b74565b6109c7565b61022e610a7e565b60005474010000000000000000000000000000000000000000900460ff16156102de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4d696e744d616e616765723a20616c7265616479206d696e746564206f6e207460448201527f68697320636861696e000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036f9190610bca565b61037a90600a610d3e565b61038890633b9aca00610d4d565b90506000805b60015481101561041e576000600182815481106103ad576103ad610d8a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090620186a06103f08388610d4d565b6103fa9190610db9565b90506104068186610df4565b9450505050808061041690610e0c565b91505061038e565b506040517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906340c10f1990604401600060405180830381600087803b1580156104ad57600080fd5b505af11580156104c1573d6000803e3d6000fd5b5050600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905550505050565b610510610a7e565b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b90602401600060405180830381600087803b15801561059957600080fd5b505af11580156105ad573d6000803e3d6000fd5b5050505050565b6105bc610a7e565b6105c66000610aff565b565b6105d0610a7e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561063857600080fd5b505af115801561064c573d6000803e3d6000fd5b50505050565b6001818154811061066257600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b610691610a7e565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107229190610bca565b61072d90600a610d3e565b61073b90633b9aca00610d4d565b905060005b6001548110156108815760006001828154811061075f5761075f610d8a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090620186a06107a28387610d4d565b6107ac9190610db9565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190610e44565b50505050808061087990610e0c565b915050610740565b506040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561090f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109339190610e66565b905080156109c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d696e744d616e616765723a20746f6b656e732072656d61696e20616674657260448201527f20646973747269627574696f6e0000000000000000000000000000000000000060648201526084016102d5565b5050565b6109cf610a7e565b73ffffffffffffffffffffffffffffffffffffffff8116610a72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102d5565b610a7b81610aff565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d5565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610b8657600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610baa57600080fd5b9392505050565b600060208284031215610bc357600080fd5b5035919050565b600060208284031215610bdc57600080fd5b815160ff81168114610baa57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600181815b80851115610c7557817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c5b57610c5b610bed565b80851615610c6857918102915b93841c9390800290610c21565b509250929050565b600082610c8c57506001610d38565b81610c9957506000610d38565b8160018114610caf5760028114610cb957610cd5565b6001915050610d38565b60ff841115610cca57610cca610bed565b50506001821b610d38565b5060208310610133831016604e8410600b8410161715610cf8575081810a610d38565b610d028383610c1c565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610d3457610d34610bed565b0290505b92915050565b6000610baa60ff841683610c7d565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610d8557610d85610bed565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082610def577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008219821115610e0757610e07610bed565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e3d57610e3d610bed565b5060010190565b600060208284031215610e5657600080fd5b81518015158114610baa57600080fd5b600060208284031215610e7857600080fd5b505191905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(MintManagerStorageLayoutJSON), MintManagerStorageLayout); err != nil {
		panic(err)
	}

	layouts["MintManager"] = MintManagerStorageLayout
	deployedBytecodes["MintManager"] = MintManagerDeployedBin
	immutableReferences["MintManager"] = true
}
