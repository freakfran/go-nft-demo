// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package moodNft

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

// MoodNftMetaData contains all meta data concerning the MoodNft contract.
var MoodNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sadSvgImgUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"happySvgImgUri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MoodNft__CannotFlipMoodIfNotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"flipMood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHappySvgImgUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSadSvgImgUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenIdToMood\",\"outputs\":[{\"internalType\":\"enumMoodNft.Mood\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612b10380380612b1083398181016040528101906100319190610239565b6040518060400160405280600881526020017f4d6f6f64204e46540000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f4d4e000000000000000000000000000000000000000000000000000000000000815250815f90816100ab91906104bc565b5080600190816100bb91906104bc565b5050505f60068190555081600790816100d491906104bc565b5080600890816100e491906104bc565b50505061058b565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61014b82610105565b810181811067ffffffffffffffff8211171561016a57610169610115565b5b80604052505050565b5f61017c6100ec565b90506101888282610142565b919050565b5f67ffffffffffffffff8211156101a7576101a6610115565b5b6101b082610105565b9050602081019050919050565b8281835e5f83830152505050565b5f6101dd6101d88461018d565b610173565b9050828152602081018484840111156101f9576101f8610101565b5b6102048482856101bd565b509392505050565b5f82601f8301126102205761021f6100fd565b5b81516102308482602086016101cb565b91505092915050565b5f806040838503121561024f5761024e6100f5565b5b5f83015167ffffffffffffffff81111561026c5761026b6100f9565b5b6102788582860161020c565b925050602083015167ffffffffffffffff811115610299576102986100f9565b5b6102a58582860161020c565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102fd57607f821691505b6020821081036103105761030f6102b9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103727fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610337565b61037c8683610337565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103c06103bb6103b684610394565b61039d565b610394565b9050919050565b5f819050919050565b6103d9836103a6565b6103ed6103e5826103c7565b848454610343565b825550505050565b5f90565b6104016103f5565b61040c8184846103d0565b505050565b5b8181101561042f576104245f826103f9565b600181019050610412565b5050565b601f8211156104745761044581610316565b61044e84610328565b8101602085101561045d578190505b61047161046985610328565b830182610411565b50505b505050565b5f82821c905092915050565b5f6104945f1984600802610479565b1980831691505092915050565b5f6104ac8383610485565b9150826002028217905092915050565b6104c5826102af565b67ffffffffffffffff8111156104de576104dd610115565b5b6104e882546102e6565b6104f3828285610433565b5f60209050601f831160018114610524575f8415610512578287015190505b61051c85826104a1565b865550610583565b601f19841661053286610316565b5f5b8281101561055957848901518255600182019150602085019450602081019050610534565b868310156105765784890151610572601f891682610485565b8355505b6001600288020188555050505b505050505050565b612578806105985f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c806370a08231116100ab578063c1a147a01161006f578063c1a147a014610317578063c2229fea14610333578063c87b56dd1461033d578063e6159a721461036d578063e985e9c51461038b5761011f565b806370a082311461026157806395d89b4114610291578063a22cb465146102af578063a789237c146102cb578063b88d4fde146102fb5761011f565b806323a3dd4d116100f257806323a3dd4d146101bd57806323b872dd146101db57806342842e0e146101f75780636352211e146102135780636e02007d146102435761011f565b806301ffc9a71461012357806306fdde0314610153578063081812fc14610171578063095ea7b3146101a1575b5f80fd5b61013d60048036038101906101389190611a30565b6103bb565b60405161014a9190611a75565b60405180910390f35b61015b61049c565b6040516101689190611afe565b60405180910390f35b61018b60048036038101906101869190611b51565b61052b565b6040516101989190611bbb565b60405180910390f35b6101bb60048036038101906101b69190611bfe565b610546565b005b6101c561055c565b6040516101d29190611afe565b60405180910390f35b6101f560048036038101906101f09190611c3c565b6105ec565b005b610211600480360381019061020c9190611c3c565b6106eb565b005b61022d60048036038101906102289190611b51565b61070a565b60405161023a9190611bbb565b60405180910390f35b61024b61071b565b6040516102589190611c9b565b60405180910390f35b61027b60048036038101906102769190611cb4565b610724565b6040516102889190611c9b565b60405180910390f35b6102996107da565b6040516102a69190611afe565b60405180910390f35b6102c960048036038101906102c49190611d09565b61086a565b005b6102e560048036038101906102e09190611b51565b610880565b6040516102f29190611dba565b60405180910390f35b61031560048036038101906103109190611eff565b6108a6565b005b610331600480360381019061032c9190611b51565b6108c3565b005b61033b6109ce565b005b61035760048036038101906103529190611b51565b610a2d565b6040516103649190611afe565b60405180910390f35b610375610bf5565b6040516103829190611afe565b60405180910390f35b6103a560048036038101906103a09190611f7f565b610c85565b6040516103b29190611a75565b60405180910390f35b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061048557507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610495575061049482610d13565b5b9050919050565b60605f80546104aa90611fea565b80601f01602080910402602001604051908101604052809291908181526020018280546104d690611fea565b80156105215780601f106104f857610100808354040283529160200191610521565b820191905f5260205f20905b81548152906001019060200180831161050457829003601f168201915b5050505050905090565b5f61053582610d7c565b5061053f82610e02565b9050919050565b6105588282610553610e3b565b610e42565b5050565b60606008805461056b90611fea565b80601f016020809104026020016040519081016040528092919081815260200182805461059790611fea565b80156105e25780601f106105b9576101008083540402835291602001916105e2565b820191905f5260205f20905b8154815290600101906020018083116105c557829003601f168201915b5050505050905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361065c575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016106539190611bbb565b60405180910390fd5b5f61066f838361066a610e3b565b610e54565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106e5578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016106dc9392919061201a565b60405180910390fd5b50505050565b61070583838360405180602001604052805f8152506108a6565b505050565b5f61071482610d7c565b9050919050565b5f600654905090565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610795575f6040517f89c62b6400000000000000000000000000000000000000000000000000000000815260040161078c9190611bbb565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600180546107e990611fea565b80601f016020809104026020016040519081016040528092919081815260200182805461081590611fea565b80156108605780601f1061083757610100808354040283529160200191610860565b820191905f5260205f20905b81548152906001019060200180831161084357829003601f168201915b5050505050905090565b61087c610875610e3b565b838361105f565b5050565b5f60095f8381526020019081526020015f205f9054906101000a900460ff169050919050565b6108b18484846105ec565b6108bd848484846111c8565b50505050565b6108d66108cf8261070a565b338361137a565b61090c576040517fe02184b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600181111561091f5761091e611d47565b5b60095f8381526020019081526020015f205f9054906101000a900460ff16600181111561094f5761094e611d47565b5b0361099257600160095f8381526020019081526020015f205f6101000a81548160ff0219169083600181111561098857610987611d47565b5b02179055506109cb565b5f60095f8381526020019081526020015f205f6101000a81548160ff021916908360018111156109c5576109c4611d47565b5b02179055505b50565b6109da3360065461143a565b5f60095f60065481526020019081526020015f205f6101000a81548160ff02191690836001811115610a0f57610a0e611d47565b5b021790555060065f815480929190610a269061207c565b9190505550565b6060805f6001811115610a4357610a42611d47565b5b60095f8581526020019081526020015f205f9054906101000a900460ff166001811115610a7357610a72611d47565b5b03610b085760088054610a8590611fea565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab190611fea565b8015610afc5780601f10610ad357610100808354040283529160200191610afc565b820191905f5260205f20905b815481529060010190602001808311610adf57829003601f168201915b50505050509050610b94565b60078054610b1590611fea565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4190611fea565b8015610b8c5780601f10610b6357610100808354040283529160200191610b8c565b820191905f5260205f20905b815481529060010190602001808311610b6f57829003601f168201915b505050505090505b610b9c611457565b610bcd610ba761049c565b83604051602001610bb99291906122bd565b604051602081830303815290604052611494565b604051602001610bde92919061230c565b604051602081830303815290604052915050919050565b606060078054610c0490611fea565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3090611fea565b8015610c7b5780601f10610c5257610100808354040283529160200191610c7b565b820191905f5260205f20905b815481529060010190602001808311610c5e57829003601f168201915b5050505050905090565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f80610d8783611601565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610df957826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610df09190611c9b565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610e4f838383600161163a565b505050565b5f80610e5f84611601565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ea057610e9f8184866117f9565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f2b57610edf5f855f8061163a565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610faa57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036110cf57816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016110c69190611bbb565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516111bb9190611a75565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115611374578273ffffffffffffffffffffffffffffffffffffffff1663150b7a0261120b610e3b565b8685856040518563ffffffff1660e01b815260040161122d9493929190612381565b6020604051808303815f875af192505050801561126857506040513d601f19601f8201168201806040525081019061126591906123df565b60015b6112e9573d805f8114611296576040519150601f19603f3d011682016040523d82523d5f602084013e61129b565b606091505b505f8151036112e157836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016112d89190611bbb565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461137257836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016113699190611bbb565b60405180910390fd5b505b50505050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561143157508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806113f257506113f18484610c85565b5b8061143057508273ffffffffffffffffffffffffffffffffffffffff1661141883610e02565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b611453828260405180602001604052805f8152506118bc565b5050565b60606040518060400160405280601d81526020017f646174613a6170706c69636174696f6e2f6a736f6e3b6261736536342c000000815250905090565b60605f8251036114b45760405180602001604052805f81525090506115fc565b5f6040518060600160405280604081526020016125036040913990505f6003600285516114e1919061240a565b6114eb919061246a565b60046114f7919061249a565b67ffffffffffffffff8111156115105761150f611ddb565b5b6040519080825280601f01601f1916602001820160405280156115425781602001600182028036833780820191505090505b509050600182016020820185865187016020810180515f82525b828410156115b7576003840193508351603f8160121c168701518653600186019550603f81600c1c168701518653600186019550603f8160061c168701518653600186019550603f811687015186536001860195505061155c565b80825260038a5106600181146115d457600281146115e7576115ef565b603d6001870353603d60028703536115ef565b603d60018703535b5050505050505080925050505b919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061167257505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156117a4575f61168184610d7c565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156116eb57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156116fe57506116fc8184610c85565b155b1561174057826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016117379190611bbb565b60405180910390fd5b81156117a257838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b61180483838361137a565b6118b7575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361187857806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161186f9190611c9b565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016118ae9291906124db565b60405180910390fd5b505050565b6118c683836118d7565b6118d25f8484846111c8565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611947575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161193e9190611bbb565b60405180910390fd5b5f61195383835f610e54565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146119c5575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016119bc9190611bbb565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611a0f816119db565b8114611a19575f80fd5b50565b5f81359050611a2a81611a06565b92915050565b5f60208284031215611a4557611a446119d3565b5b5f611a5284828501611a1c565b91505092915050565b5f8115159050919050565b611a6f81611a5b565b82525050565b5f602082019050611a885f830184611a66565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611ad082611a8e565b611ada8185611a98565b9350611aea818560208601611aa8565b611af381611ab6565b840191505092915050565b5f6020820190508181035f830152611b168184611ac6565b905092915050565b5f819050919050565b611b3081611b1e565b8114611b3a575f80fd5b50565b5f81359050611b4b81611b27565b92915050565b5f60208284031215611b6657611b656119d3565b5b5f611b7384828501611b3d565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611ba582611b7c565b9050919050565b611bb581611b9b565b82525050565b5f602082019050611bce5f830184611bac565b92915050565b611bdd81611b9b565b8114611be7575f80fd5b50565b5f81359050611bf881611bd4565b92915050565b5f8060408385031215611c1457611c136119d3565b5b5f611c2185828601611bea565b9250506020611c3285828601611b3d565b9150509250929050565b5f805f60608486031215611c5357611c526119d3565b5b5f611c6086828701611bea565b9350506020611c7186828701611bea565b9250506040611c8286828701611b3d565b9150509250925092565b611c9581611b1e565b82525050565b5f602082019050611cae5f830184611c8c565b92915050565b5f60208284031215611cc957611cc86119d3565b5b5f611cd684828501611bea565b91505092915050565b611ce881611a5b565b8114611cf2575f80fd5b50565b5f81359050611d0381611cdf565b92915050565b5f8060408385031215611d1f57611d1e6119d3565b5b5f611d2c85828601611bea565b9250506020611d3d85828601611cf5565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110611d8557611d84611d47565b5b50565b5f819050611d9582611d74565b919050565b5f611da482611d88565b9050919050565b611db481611d9a565b82525050565b5f602082019050611dcd5f830184611dab565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611e1182611ab6565b810181811067ffffffffffffffff82111715611e3057611e2f611ddb565b5b80604052505050565b5f611e426119ca565b9050611e4e8282611e08565b919050565b5f67ffffffffffffffff821115611e6d57611e6c611ddb565b5b611e7682611ab6565b9050602081019050919050565b828183375f83830152505050565b5f611ea3611e9e84611e53565b611e39565b905082815260208101848484011115611ebf57611ebe611dd7565b5b611eca848285611e83565b509392505050565b5f82601f830112611ee657611ee5611dd3565b5b8135611ef6848260208601611e91565b91505092915050565b5f805f8060808587031215611f1757611f166119d3565b5b5f611f2487828801611bea565b9450506020611f3587828801611bea565b9350506040611f4687828801611b3d565b925050606085013567ffffffffffffffff811115611f6757611f666119d7565b5b611f7387828801611ed2565b91505092959194509250565b5f8060408385031215611f9557611f946119d3565b5b5f611fa285828601611bea565b9250506020611fb385828601611bea565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061200157607f821691505b60208210810361201457612013611fbd565b5b50919050565b5f60608201905061202d5f830186611bac565b61203a6020830185611c8c565b6120476040830184611bac565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61208682611b1e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120b8576120b761204f565b5b600182019050919050565b5f81905092915050565b7f7b226e616d65223a2200000000000000000000000000000000000000000000005f82015250565b5f6121016009836120c3565b915061210c826120cd565b600982019050919050565b5f61212182611a8e565b61212b81856120c3565b935061213b818560208601611aa8565b80840191505092915050565b7f222c20226465736372697074696f6e223a22416e204e465420746861742072655f8201527f666c6563747320746865206d6f6f64206f6620746865206f776e65722c20313060208201527f3025206f6e20436861696e21222c200000000000000000000000000000000000604082015250565b5f6121c7604f836120c3565b91506121d282612147565b604f82019050919050565b7f2261747472696275746573223a205b7b2274726169745f74797065223a20226d5f8201527f6f6f64696e657373222c202276616c7565223a203130307d5d2c2022696d616760208201527f65223a2200000000000000000000000000000000000000000000000000000000604082015250565b5f61225d6044836120c3565b9150612268826121dd565b604482019050919050565b7f227d0000000000000000000000000000000000000000000000000000000000005f82015250565b5f6122a76002836120c3565b91506122b282612273565b600282019050919050565b5f6122c7826120f5565b91506122d38285612117565b91506122de826121bb565b91506122e982612251565b91506122f58284612117565b91506123008261229b565b91508190509392505050565b5f6123178285612117565b91506123238284612117565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6123538261232f565b61235d8185612339565b935061236d818560208601611aa8565b61237681611ab6565b840191505092915050565b5f6080820190506123945f830187611bac565b6123a16020830186611bac565b6123ae6040830185611c8c565b81810360608301526123c08184612349565b905095945050505050565b5f815190506123d981611a06565b92915050565b5f602082840312156123f4576123f36119d3565b5b5f612401848285016123cb565b91505092915050565b5f61241482611b1e565b915061241f83611b1e565b92508282019050808211156124375761243661204f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61247482611b1e565b915061247f83611b1e565b92508261248f5761248e61243d565b5b828204905092915050565b5f6124a482611b1e565b91506124af83611b1e565b92508282026124bd81611b1e565b915082820484148315176124d4576124d361204f565b5b5092915050565b5f6040820190506124ee5f830185611bac565b6124fb6020830184611c8c565b939250505056fe4142434445464748494a4b4c4d4e4f505152535455565758595a6162636465666768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2fa2646970667358221220e6be63c1154811871f7908fc4eb84b15cd3016a381eaa9b313c28998e87b624e64736f6c634300081a0033",
}

// MoodNftABI is the input ABI used to generate the binding from.
// Deprecated: Use MoodNftMetaData.ABI instead.
var MoodNftABI = MoodNftMetaData.ABI

// MoodNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MoodNftMetaData.Bin instead.
var MoodNftBin = MoodNftMetaData.Bin

// DeployMoodNft deploys a new Ethereum contract, binding an instance of MoodNft to it.
func DeployMoodNft(auth *bind.TransactOpts, backend bind.ContractBackend, sadSvgImgUri string, happySvgImgUri string) (common.Address, *types.Transaction, *MoodNft, error) {
	parsed, err := MoodNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MoodNftBin), backend, sadSvgImgUri, happySvgImgUri)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MoodNft{MoodNftCaller: MoodNftCaller{contract: contract}, MoodNftTransactor: MoodNftTransactor{contract: contract}, MoodNftFilterer: MoodNftFilterer{contract: contract}}, nil
}

// MoodNft is an auto generated Go binding around an Ethereum contract.
type MoodNft struct {
	MoodNftCaller     // Read-only binding to the contract
	MoodNftTransactor // Write-only binding to the contract
	MoodNftFilterer   // Log filterer for contract events
}

// MoodNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoodNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoodNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoodNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoodNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoodNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoodNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoodNftSession struct {
	Contract     *MoodNft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoodNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoodNftCallerSession struct {
	Contract *MoodNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MoodNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoodNftTransactorSession struct {
	Contract     *MoodNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MoodNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoodNftRaw struct {
	Contract *MoodNft // Generic contract binding to access the raw methods on
}

// MoodNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoodNftCallerRaw struct {
	Contract *MoodNftCaller // Generic read-only contract binding to access the raw methods on
}

// MoodNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoodNftTransactorRaw struct {
	Contract *MoodNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoodNft creates a new instance of MoodNft, bound to a specific deployed contract.
func NewMoodNft(address common.Address, backend bind.ContractBackend) (*MoodNft, error) {
	contract, err := bindMoodNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MoodNft{MoodNftCaller: MoodNftCaller{contract: contract}, MoodNftTransactor: MoodNftTransactor{contract: contract}, MoodNftFilterer: MoodNftFilterer{contract: contract}}, nil
}

// NewMoodNftCaller creates a new read-only instance of MoodNft, bound to a specific deployed contract.
func NewMoodNftCaller(address common.Address, caller bind.ContractCaller) (*MoodNftCaller, error) {
	contract, err := bindMoodNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoodNftCaller{contract: contract}, nil
}

// NewMoodNftTransactor creates a new write-only instance of MoodNft, bound to a specific deployed contract.
func NewMoodNftTransactor(address common.Address, transactor bind.ContractTransactor) (*MoodNftTransactor, error) {
	contract, err := bindMoodNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoodNftTransactor{contract: contract}, nil
}

// NewMoodNftFilterer creates a new log filterer instance of MoodNft, bound to a specific deployed contract.
func NewMoodNftFilterer(address common.Address, filterer bind.ContractFilterer) (*MoodNftFilterer, error) {
	contract, err := bindMoodNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoodNftFilterer{contract: contract}, nil
}

// bindMoodNft binds a generic wrapper to an already deployed contract.
func bindMoodNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MoodNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoodNft *MoodNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoodNft.Contract.MoodNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoodNft *MoodNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoodNft.Contract.MoodNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoodNft *MoodNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoodNft.Contract.MoodNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoodNft *MoodNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoodNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoodNft *MoodNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoodNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoodNft *MoodNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoodNft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MoodNft *MoodNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MoodNft *MoodNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MoodNft.Contract.BalanceOf(&_MoodNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MoodNft *MoodNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MoodNft.Contract.BalanceOf(&_MoodNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MoodNft.Contract.GetApproved(&_MoodNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MoodNft.Contract.GetApproved(&_MoodNft.CallOpts, tokenId)
}

// GetHappySvgImgUri is a free data retrieval call binding the contract method 0x23a3dd4d.
//
// Solidity: function getHappySvgImgUri() view returns(string)
func (_MoodNft *MoodNftCaller) GetHappySvgImgUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "getHappySvgImgUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHappySvgImgUri is a free data retrieval call binding the contract method 0x23a3dd4d.
//
// Solidity: function getHappySvgImgUri() view returns(string)
func (_MoodNft *MoodNftSession) GetHappySvgImgUri() (string, error) {
	return _MoodNft.Contract.GetHappySvgImgUri(&_MoodNft.CallOpts)
}

// GetHappySvgImgUri is a free data retrieval call binding the contract method 0x23a3dd4d.
//
// Solidity: function getHappySvgImgUri() view returns(string)
func (_MoodNft *MoodNftCallerSession) GetHappySvgImgUri() (string, error) {
	return _MoodNft.Contract.GetHappySvgImgUri(&_MoodNft.CallOpts)
}

// GetSadSvgImgUri is a free data retrieval call binding the contract method 0xe6159a72.
//
// Solidity: function getSadSvgImgUri() view returns(string)
func (_MoodNft *MoodNftCaller) GetSadSvgImgUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "getSadSvgImgUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSadSvgImgUri is a free data retrieval call binding the contract method 0xe6159a72.
//
// Solidity: function getSadSvgImgUri() view returns(string)
func (_MoodNft *MoodNftSession) GetSadSvgImgUri() (string, error) {
	return _MoodNft.Contract.GetSadSvgImgUri(&_MoodNft.CallOpts)
}

// GetSadSvgImgUri is a free data retrieval call binding the contract method 0xe6159a72.
//
// Solidity: function getSadSvgImgUri() view returns(string)
func (_MoodNft *MoodNftCallerSession) GetSadSvgImgUri() (string, error) {
	return _MoodNft.Contract.GetSadSvgImgUri(&_MoodNft.CallOpts)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_MoodNft *MoodNftCaller) GetTokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "getTokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_MoodNft *MoodNftSession) GetTokenCounter() (*big.Int, error) {
	return _MoodNft.Contract.GetTokenCounter(&_MoodNft.CallOpts)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_MoodNft *MoodNftCallerSession) GetTokenCounter() (*big.Int, error) {
	return _MoodNft.Contract.GetTokenCounter(&_MoodNft.CallOpts)
}

// GetTokenIdToMood is a free data retrieval call binding the contract method 0xa789237c.
//
// Solidity: function getTokenIdToMood(uint256 tokenId) view returns(uint8)
func (_MoodNft *MoodNftCaller) GetTokenIdToMood(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "getTokenIdToMood", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIdToMood is a free data retrieval call binding the contract method 0xa789237c.
//
// Solidity: function getTokenIdToMood(uint256 tokenId) view returns(uint8)
func (_MoodNft *MoodNftSession) GetTokenIdToMood(tokenId *big.Int) (uint8, error) {
	return _MoodNft.Contract.GetTokenIdToMood(&_MoodNft.CallOpts, tokenId)
}

// GetTokenIdToMood is a free data retrieval call binding the contract method 0xa789237c.
//
// Solidity: function getTokenIdToMood(uint256 tokenId) view returns(uint8)
func (_MoodNft *MoodNftCallerSession) GetTokenIdToMood(tokenId *big.Int) (uint8, error) {
	return _MoodNft.Contract.GetTokenIdToMood(&_MoodNft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MoodNft *MoodNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MoodNft *MoodNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MoodNft.Contract.IsApprovedForAll(&_MoodNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MoodNft *MoodNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MoodNft.Contract.IsApprovedForAll(&_MoodNft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MoodNft *MoodNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MoodNft *MoodNftSession) Name() (string, error) {
	return _MoodNft.Contract.Name(&_MoodNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MoodNft *MoodNftCallerSession) Name() (string, error) {
	return _MoodNft.Contract.Name(&_MoodNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MoodNft.Contract.OwnerOf(&_MoodNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MoodNft *MoodNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MoodNft.Contract.OwnerOf(&_MoodNft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MoodNft *MoodNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MoodNft *MoodNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MoodNft.Contract.SupportsInterface(&_MoodNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MoodNft *MoodNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MoodNft.Contract.SupportsInterface(&_MoodNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MoodNft *MoodNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MoodNft *MoodNftSession) Symbol() (string, error) {
	return _MoodNft.Contract.Symbol(&_MoodNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MoodNft *MoodNftCallerSession) Symbol() (string, error) {
	return _MoodNft.Contract.Symbol(&_MoodNft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MoodNft *MoodNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MoodNft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MoodNft *MoodNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MoodNft.Contract.TokenURI(&_MoodNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MoodNft *MoodNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MoodNft.Contract.TokenURI(&_MoodNft.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.Approve(&_MoodNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.Approve(&_MoodNft.TransactOpts, to, tokenId)
}

// FlipMood is a paid mutator transaction binding the contract method 0xc1a147a0.
//
// Solidity: function flipMood(uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactor) FlipMood(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "flipMood", tokenId)
}

// FlipMood is a paid mutator transaction binding the contract method 0xc1a147a0.
//
// Solidity: function flipMood(uint256 tokenId) returns()
func (_MoodNft *MoodNftSession) FlipMood(tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.FlipMood(&_MoodNft.TransactOpts, tokenId)
}

// FlipMood is a paid mutator transaction binding the contract method 0xc1a147a0.
//
// Solidity: function flipMood(uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactorSession) FlipMood(tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.FlipMood(&_MoodNft.TransactOpts, tokenId)
}

// MintNft is a paid mutator transaction binding the contract method 0xc2229fea.
//
// Solidity: function mintNft() returns()
func (_MoodNft *MoodNftTransactor) MintNft(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "mintNft")
}

// MintNft is a paid mutator transaction binding the contract method 0xc2229fea.
//
// Solidity: function mintNft() returns()
func (_MoodNft *MoodNftSession) MintNft() (*types.Transaction, error) {
	return _MoodNft.Contract.MintNft(&_MoodNft.TransactOpts)
}

// MintNft is a paid mutator transaction binding the contract method 0xc2229fea.
//
// Solidity: function mintNft() returns()
func (_MoodNft *MoodNftTransactorSession) MintNft() (*types.Transaction, error) {
	return _MoodNft.Contract.MintNft(&_MoodNft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.SafeTransferFrom(&_MoodNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.SafeTransferFrom(&_MoodNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MoodNft *MoodNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MoodNft *MoodNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MoodNft.Contract.SafeTransferFrom0(&_MoodNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MoodNft *MoodNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MoodNft.Contract.SafeTransferFrom0(&_MoodNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MoodNft *MoodNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MoodNft *MoodNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MoodNft.Contract.SetApprovalForAll(&_MoodNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MoodNft *MoodNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MoodNft.Contract.SetApprovalForAll(&_MoodNft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.TransferFrom(&_MoodNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MoodNft *MoodNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MoodNft.Contract.TransferFrom(&_MoodNft.TransactOpts, from, to, tokenId)
}

// MoodNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MoodNft contract.
type MoodNftApprovalIterator struct {
	Event *MoodNftApproval // Event containing the contract specifics and raw log

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
func (it *MoodNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoodNftApproval)
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
		it.Event = new(MoodNftApproval)
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
func (it *MoodNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoodNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoodNftApproval represents a Approval event raised by the MoodNft contract.
type MoodNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MoodNftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MoodNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MoodNftApprovalIterator{contract: _MoodNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MoodNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MoodNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoodNftApproval)
				if err := _MoodNft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) ParseApproval(log types.Log) (*MoodNftApproval, error) {
	event := new(MoodNftApproval)
	if err := _MoodNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoodNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MoodNft contract.
type MoodNftApprovalForAllIterator struct {
	Event *MoodNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MoodNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoodNftApprovalForAll)
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
		it.Event = new(MoodNftApprovalForAll)
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
func (it *MoodNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoodNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoodNftApprovalForAll represents a ApprovalForAll event raised by the MoodNft contract.
type MoodNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MoodNft *MoodNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MoodNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MoodNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MoodNftApprovalForAllIterator{contract: _MoodNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MoodNft *MoodNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MoodNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MoodNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoodNftApprovalForAll)
				if err := _MoodNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MoodNft *MoodNftFilterer) ParseApprovalForAll(log types.Log) (*MoodNftApprovalForAll, error) {
	event := new(MoodNftApprovalForAll)
	if err := _MoodNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoodNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MoodNft contract.
type MoodNftTransferIterator struct {
	Event *MoodNftTransfer // Event containing the contract specifics and raw log

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
func (it *MoodNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoodNftTransfer)
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
		it.Event = new(MoodNftTransfer)
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
func (it *MoodNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoodNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoodNftTransfer represents a Transfer event raised by the MoodNft contract.
type MoodNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MoodNftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MoodNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MoodNftTransferIterator{contract: _MoodNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MoodNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MoodNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoodNftTransfer)
				if err := _MoodNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MoodNft *MoodNftFilterer) ParseTransfer(log types.Log) (*MoodNftTransfer, error) {
	event := new(MoodNftTransfer)
	if err := _MoodNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
