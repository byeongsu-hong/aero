// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ParentBridgeABI is the input ABI used to generate the binding from.
const ParentBridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"coinCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childBlocks\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHALLENGE_PERIOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"depositBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockSubmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ExitRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExitFinalized\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNonFungible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositFungible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"exitTxData\",\"type\":\"bytes\"},{\"name\":\"exitProof\",\"type\":\"bytes\"},{\"name\":\"exitBlock\",\"type\":\"uint256\"},{\"name\":\"prevTxData\",\"type\":\"bytes\"},{\"name\":\"prevProof\",\"type\":\"bytes\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"claimTxData\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextDepositBlockIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ParentBridgeBin is the compiled bytecode used for deploying new contracts.
const ParentBridgeBin = `0x60806040526004805467ffffffffffffffff191690553480156200002257600080fd5b506200002d620000b9565b604051809103906000f0801580156200004a573d6000803e3d6000fd5b5060008054600160a060020a031916600160a060020a039290921691909117905562000075620000ca565b604051809103906000f08015801562000092573d6000803e3d6000fd5b5060018054600160a060020a031916600160a060020a0392909216919091179055620000db565b6040516105a18062001b1683390190565b6040516104f680620020b783390190565b611a2b80620000eb6000396000f3006080604052600436106100ce5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662a3216c81146100d3578063124135c9146100fa5780632791fd4c146102595780634b991bde146102835780634bb278f314610369578063678fd2891461037e5780637a95f1e8146103b057806397ef999b146103c5578063a831fa07146103f6578063a98c7f2c1461040b578063baa4769414610420578063bb2dc86314610438578063c3a079ed14610450578063d551c03514610465575b600080fd5b3480156100df57600080fd5b506100e861048f565b60408051918252519081900360200190f35b34801561010657600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261025794369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a90999401975091955091820193509150819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a9099940197509195509182019350915081908401838280828437509497506104bf9650505050505050565b005b34801561026557600080fd5b50610257600160a060020a036004358116906024351660443561051d565b34801561028f57600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261025794823567ffffffffffffffff1694602480359536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506105ba9650505050505050565b34801561037557600080fd5b506102576108be565b34801561038a57600080fd5b506103936108c0565b6040805167ffffffffffffffff9092168252519081900360200190f35b3480156103bc57600080fd5b506100e86108d0565b3480156103d157600080fd5b506103dd6004356108d6565b6040805192835260208301919091528051918290030190f35b34801561040257600080fd5b506100e86108ef565b34801561041757600080fd5b506100e86108f5565b34801561042c57600080fd5b506100e86004356108fb565b34801561044457600080fd5b506103dd600435610a26565b34801561045c57600080fd5b506100e8610a40565b34801561047157600080fd5b50610257600160a060020a0360043581169060243516604435610a45565b60006104ba6003546104ae6103e8600254610a6f90919063ffffffff16565b9063ffffffff610a8116565b905090565b6103e88506156104ea576000858152600560205260409020546104e59086908989610a94565b610514565b600085815260056020526040808220548483529120546105149187918a91908a9089908988610c0a565b50505050505050565b604080517f42842e0e000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152306024830152604482018490529151918516916342842e0e9160648082019260009290919082900301818387803b15801561058e57600080fd5b505af11580156105a2573d6000803e3d6000fd5b505050506105b560018484846001610ed5565b505050565b6105c26119a0565b6000600167ffffffffffffffff881660009081526006602052604090206005015460ff1660028111156105f157fe5b14610646576040805160e560020a62461bcd02815260206004820152601f60248201527f6f6e6c792065786974696e6720636f696e2063616e206368616c6c656e676500604482015290519081900360640190fd5b61064f856111ce565b67ffffffffffffffff8089166000818152600660205260409020835193955093509116146106c7576040805160e560020a62461bcd02815260206004820152600c60248201527f696e76616c696420736c6f740000000000000000000000000000000000000000604482015290519081900360640190fd5b60088101546060830151600160a060020a03909116906106ed908563ffffffff6113a216565b600160a060020a03161461074b576040805160e560020a62461bcd02815260206004820152600b60248201527f696e76616c696420736967000000000000000000000000000000000000000000604482015290519081900360640190fd5b6107558787611477565b1561075f576107be565b60078101546020830151146107be576040805160e560020a62461bcd02815260206004820152601160248201527f696e76616c6964207265666572656e6365000000000000000000000000000000604482015290519081900360640190fd5b6000868152600560205260409020546107d9908390866114c4565b151561082f576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b60058101805460ff1916905560068101805473ffffffffffffffffffffffffffffffffffffffff1990811690915560006007830181905560088301805490921690915560098201819055600a8201819055604051339167ffffffffffffffff8a16917f172124f3161dc498f87d145ee7397b5063391ae141a8ed2e4170b71d1f1128d19190a350505050505050565b565b60045467ffffffffffffffff1681565b60025481565b6005602052600090815260409020805460019091015482565b6103e881565b60035481565b60008054604080517f6d70f7ae0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691636d70f7ae9160248082019260209290919082900301818787803b15801561096157600080fd5b505af1158015610975573d6000803e3d6000fd5b505050506040513d602081101561098b57600080fd5b5051151561099857600080fd5b60408051808201825283815242602080830191825260028054600090815260059092529390209151825551600190910155546109dc906103e863ffffffff610a8116565b60025560016003556040805183815242602082015281517f79f92be16163b2592e15e92db55fab076ee3c8b88792ef547d89febb58170792929181900390910190a1505060025490565b600090815260056020526040902080546001909101549091565b60c881565b610a60600160a060020a03841683308463ffffffff6115ea16565b6105b560008484600085610ed5565b600082821115610a7b57fe5b50900390565b81810182811015610a8e57fe5b92915050565b610a9c6119a0565b610aa5836111ce565b6040810151909150600160a060020a03163314610b0c576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e206578697400000000000000000000000000604482015290519081900360640190fd5b602081015115610b8c576040805160e560020a62461bcd02815260206004820152602760248201527f7072657620626c6f636b2073686f756c64206265207a65726f20696e2064657060448201527f6f73697420747800000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610b978185846114c4565b1515610bed576040805160e560020a62461bcd02815260206004820152601760248201527f6d656d6265727368697020636865636b206661696c6564000000000000000000604482015290519081900360640190fd5b610c03816000015182604001518760008061169b565b5050505050565b610c126119a0565b610c1a6119a0565b610c23896111ce565b9150610c2e866111ce565b6040830151909150600160a060020a03163314610c95576040805160e560020a62461bcd02815260206004820152600d60248201527f696e76616c6964206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b8151815167ffffffffffffffff908116911614610cfc576040805160e560020a62461bcd02815260206004820152600e60248201527f696e76616c696420736c6f744964000000000000000000000000000000000000604482015290519081900360640190fd5b6060820151610d11908463ffffffff6113a216565b600160a060020a03168160400151600160a060020a0316141515610d7f576040805160e560020a62461bcd02815260206004820152600b60248201527f696e76616c696420736967000000000000000000000000000000000000000000604482015290519081900360640190fd5b6020808301516000908152600590915260409020548514610dea576040805160e560020a62461bcd02815260206004820152600f60248201527f696e76616c696420747820646174610000000000000000000000000000000000604482015290519081900360640190fd5b610df58289896114c4565b1515610e4b576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b610e568186866114c4565b1515610eac576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b610ec9826000015183604001518c8460400151866020015161169b565b50505050505050505050565b6000806000806103e8600354101515610f5e576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c7920616c6c6f7720757020746f2031303030206465706f73697473207060448201527f6572206368696c6420626c6f636b2e0000000000000000000000000000000000606482015290519081900360840190fd5b60045460408051600160a060020a038a81166c01000000000000000000000000908102602080850191909152918d1602603483015267ffffffffffffffff909316780100000000000000000000000000000000000000000000000002604882015281518082036030018152605090910191829052805190928291908401908083835b60208310610fff5780518252601f199092019160209182019101610fe0565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902093508378010000000000000000000000000000000000000000000000009004925061105561048f565b604080518082018252868152426020808301918252600085815260059091529290922090518155905160019182015560035491935061109a919063ffffffff610a8116565b6003555067ffffffffffffffff8216600090815260066020526040902080548990829060ff1916600183818111156110ce57fe5b0217905550805474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a0389811682029290921783556001808401805473ffffffffffffffffffffffffffffffffffffffff19168c8516178155600285018a90556003850186815560048087018b815560058801805460ff19169055815467ffffffffffffffff19811667ffffffffffffffff91821690960181169590951790915590548654925491546040805193881684526020840191909152805191969590930490941693928816927f673daafb22e46978f66bf78c4b57c5ca0f03f8af2c773103438514ec289a0bc9929181900390910190a4505050505050505050565b6111d66119a0565b60606111e06119a0565b6111fa60036111ee86611776565b9063ffffffff61179916565b915061121d82600081518110151561120e57fe5b90602001906020020151611827565b67ffffffffffffffff168152815161123c908390600190811061120e57fe5b60208201528151611263908390600290811061125457fe5b9060200190602002015161184b565b600160a060020a0316604082015260208101511515611335578060000151604051602001808267ffffffffffffffff1667ffffffffffffffff1678010000000000000000000000000000000000000000000000000281526008019150506040516020818303038152906040526040518082805190602001908083835b602083106112fe5780518252601f1990920191602091820191016112df565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120606085015250611397915050565b836040518082805190602001908083835b602083106113655780518252601f199092019160209182019101611346565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060608501525050505b8092505b5050919050565b600080600080845160411415156113bc576000935061146e565b50505060208201516040830151606084015160001a601b60ff821610156113e157601b015b8060ff16601b141580156113f957508060ff16601c14155b15611407576000935061146e565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015611461573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b67ffffffffffffffff8216600090815260066020526040812060070154821080156114bd575067ffffffffffffffff831660009081526006602052604090206009015482115b9392505050565b600154606084015184516040517ff586df65000000000000000000000000000000000000000000000000000000008152600481018381526024820187905267ffffffffffffffff83166044830152608060648301908152865160848401528651600096600160a060020a03169563f586df659590948a9491938a93909160a4019060208501908083838e5b8381101561156757818101518382015260200161154f565b50505050905090810190601f1680156115945780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156115b657600080fd5b505af11580156115ca573d6000803e3d6000fd5b505050506040513d60208110156115e057600080fd5b5051949350505050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b15801561165e57600080fd5b505af1158015611672573d6000803e3d6000fd5b505050506040513d602081101561168857600080fd5b5051151561169557600080fd5b50505050565b67ffffffffffffffff85166000818152600660208181526040808420815160a081018352600160a060020a038b81168083529482018b9052898116828501819052606083018a9052436080909301839052958301805473ffffffffffffffffffffffffffffffffffffffff1990811690961790819055600784018c905560088401805490961690961790945560098201889055600a82015560058101805460ff191660011790559051909492909116927ff40a6c02d2711d3d98c0dcb225f64eec06734f556e141a616f79a18fe9fb359e91a3505050505050565b61177e6119c7565b50805160408051808201909152602092830181529182015290565b60606117a36119de565b6000836040519080825280602002602001820160405280156117df57816020015b6117cc6119c7565b8152602001906001900390816117c45790505b5092506117eb85611870565b91505b8381101561181f576117ff82611895565b838281518110151561180d57fe5b602090810290910101526001016117ee565b505092915050565b6000806000611835846118c7565b90516020919091036101000a9004949350505050565b600080611857836118c7565b50516c0100000000000000000000000090049392505050565b6118786119de565b60006118838361192a565b83519383529092016020820152919050565b61189d6119c7565b602082015160006118ad82611970565b828452602080850182905292019390910192909252919050565b805180516000918291821a908260808310156118e95781945060019350611922565b60b88310156119075760018660200151039350816001019450611922565b60b78303905080600187602001510303935080820160010194505b505050915091565b8051805160009190821a906080821015611947576000925061139b565b60b8821080611962575060c08210158015611962575060f882105b1561139b576001925061139b565b8051600090811a6080811015611989576001915061199a565b60b881101561199a57607e19810191505b50919050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b604080518082019091526000808252602082015290565b6060604051908101604052806119f26119c7565b81526020016000815250905600a165627a7a72305820f48ebaf3986a197e0ad0d794ded800a4028317c963f650f99341ffdaf0c82ab90029608060405234801561001057600080fd5b5060008054600160a060020a0319163390811782558152600160208190526040909120805460ff191690911790556105548061004d6000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631124e56f811461009d57806313e7c9d8146100c05780633fe529f0146100f55780636d70f7ae14610116578063715018a6146101375780638da5cb5b1461014c578063f2fde38b1461017d578063fa52c7d81461019e578063facd743b146101bf575b600080fd5b3480156100a957600080fd5b506100be600160a060020a03600435166101e0565b005b3480156100cc57600080fd5b506100e1600160a060020a03600435166102ba565b604080519115158252519081900360200190f35b34801561010157600080fd5b506100be600160a060020a03600435166102cf565b34801561012257600080fd5b506100e1600160a060020a03600435166103a9565b34801561014357600080fd5b506100be6103c7565b34801561015857600080fd5b50610161610433565b60408051600160a060020a039092168252519081900360200190f35b34801561018957600080fd5b506100be600160a060020a0360043516610442565b3480156101aa57600080fd5b506100e1600160a060020a0360043516610465565b3480156101cb57600080fd5b506100e1600160a060020a036004351661047a565b6101e93361047a565b806101fe5750600054600160a060020a031633145b151561029157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f6e6c79206f70657261746f722063616e2073776974636820746865206f706560448201527f7261746f722e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03166000908152600260205260409020805460ff19811660ff90911615179055565b60016020526000908152604090205460ff1681565b6102d8336103a9565b806102ed5750600054600160a060020a031633145b151561038057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f6e6c79206f70657261746f722063616e2073776974636820746865206f706560448201527f7261746f722e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03166000908152600160205260409020805460ff19811660ff90911615179055565b600160a060020a031660009081526001602052604090205460ff1690565b600054600160a060020a031633146103de57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461045957600080fd5b610462816104ab565b50565b60026020526000908152604090205460ff1681565b600160a060020a03811660009081526002602052604081205460ff16806104a557506104a5826103a9565b92915050565b600160a060020a03811615156104c057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582007a43123081fe42a5eefabb1d410664a8b19bb1899d0ba076a6bf2e3d92115140029608060405234801561001057600080fd5b507f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56360005561004a6001604064010000000061004f810204565b61012c565b815b60ff8083169082161161012757600060ff6000198301166041811061007257fe5b0154600060ff6000198401166041811061008857fe5b0154604080516020808201949094528082019290925280518083038201815260609092019081905281519192909182918401908083835b602083106100de5780518252601f1990920191602091820191016100bf565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060008260ff1660418110151561011d57fe5b0155600101610051565b505050565b6103bb8061013b6000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166348419ad8811461005b578063509a7e5414610085578063f586df65146100ef575b600080fd5b34801561006757600080fd5b50610073600435610170565b60408051918252519081900360200190f35b34801561009157600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610073948235946024803567ffffffffffffffff16953695946064949201919081908401838280828437509497506101849650505050505050565b3480156100fb57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261015c94803594602480359567ffffffffffffffff6044351695369560849493019181908401838280828437509497506103759650505050505050565b604080519115158252519081900360200190f35b6000816041811061017d57fe5b0154905081565b6000806000806000806020600888510381151561019d57fe5b061580156101ae5750610808875111155b15156101b957600080fd5b50505050602083015185906008907801000000000000000000000000000000000000000000000000900460005b60408110156103685760018216151561020f576000816041811061020657fe5b01549450610230565b6020830192508261ffff1687511015151561022957600080fd5b8287015194505b6001881615156102bf57604080516020808201879052818301889052825180830384018152606090920192839052815191929182918401908083835b6020831061028b5780518252601f19909201916020918201910161026c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350610340565b604080516020808201889052818301879052825180830384018152606090920192839052815191929182918401908083835b602083106103105780518252601f1990920191602091820191016102f1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902093505b600267ffffffffffffffff8316049150600267ffffffffffffffff89160497506001016101e6565b5091979650505050505050565b600080610383868585610184565b909414959450505050505600a165627a7a72305820cb95e8633f20ff9e97eca502ac61407e0353f7c129ebdbbe0c9ed7b96cf2664f0029`

// DeployParentBridge deploys a new Ethereum contract, binding an instance of ParentBridge to it.
func DeployParentBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ParentBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParentBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ParentBridge{ParentBridgeCaller: ParentBridgeCaller{contract: contract}, ParentBridgeTransactor: ParentBridgeTransactor{contract: contract}, ParentBridgeFilterer: ParentBridgeFilterer{contract: contract}}, nil
}

// ParentBridge is an auto generated Go binding around an Ethereum contract.
type ParentBridge struct {
	ParentBridgeCaller     // Read-only binding to the contract
	ParentBridgeTransactor // Write-only binding to the contract
	ParentBridgeFilterer   // Log filterer for contract events
}

// ParentBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParentBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParentBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParentBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParentBridgeSession struct {
	Contract     *ParentBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParentBridgeCallerSession struct {
	Contract *ParentBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ParentBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParentBridgeTransactorSession struct {
	Contract     *ParentBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ParentBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParentBridgeRaw struct {
	Contract *ParentBridge // Generic contract binding to access the raw methods on
}

// ParentBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParentBridgeCallerRaw struct {
	Contract *ParentBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ParentBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParentBridgeTransactorRaw struct {
	Contract *ParentBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParentBridge creates a new instance of ParentBridge, bound to a specific deployed contract.
func NewParentBridge(address common.Address, backend bind.ContractBackend) (*ParentBridge, error) {
	contract, err := bindParentBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParentBridge{ParentBridgeCaller: ParentBridgeCaller{contract: contract}, ParentBridgeTransactor: ParentBridgeTransactor{contract: contract}, ParentBridgeFilterer: ParentBridgeFilterer{contract: contract}}, nil
}

// NewParentBridgeCaller creates a new read-only instance of ParentBridge, bound to a specific deployed contract.
func NewParentBridgeCaller(address common.Address, caller bind.ContractCaller) (*ParentBridgeCaller, error) {
	contract, err := bindParentBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeCaller{contract: contract}, nil
}

// NewParentBridgeTransactor creates a new write-only instance of ParentBridge, bound to a specific deployed contract.
func NewParentBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ParentBridgeTransactor, error) {
	contract, err := bindParentBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeTransactor{contract: contract}, nil
}

// NewParentBridgeFilterer creates a new log filterer instance of ParentBridge, bound to a specific deployed contract.
func NewParentBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ParentBridgeFilterer, error) {
	contract, err := bindParentBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeFilterer{contract: contract}, nil
}

// bindParentBridge binds a generic wrapper to an already deployed contract.
func bindParentBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParentBridge *ParentBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParentBridge.Contract.ParentBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParentBridge *ParentBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentBridge.Contract.ParentBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParentBridge *ParentBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParentBridge.Contract.ParentBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParentBridge *ParentBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParentBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParentBridge *ParentBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParentBridge *ParentBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParentBridge.Contract.contract.Transact(opts, method, params...)
}

// CHALLENGEPERIOD is a free data retrieval call binding the contract method 0xc3a079ed.
//
// Solidity: function CHALLENGE_PERIOD() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) CHALLENGEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "CHALLENGE_PERIOD")
	return *ret0, err
}

// CHALLENGEPERIOD is a free data retrieval call binding the contract method 0xc3a079ed.
//
// Solidity: function CHALLENGE_PERIOD() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) CHALLENGEPERIOD() (*big.Int, error) {
	return _ParentBridge.Contract.CHALLENGEPERIOD(&_ParentBridge.CallOpts)
}

// CHALLENGEPERIOD is a free data retrieval call binding the contract method 0xc3a079ed.
//
// Solidity: function CHALLENGE_PERIOD() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) CHALLENGEPERIOD() (*big.Int, error) {
	return _ParentBridge.Contract.CHALLENGEPERIOD(&_ParentBridge.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) CHILDBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "CHILD_BLOCK_INTERVAL")
	return *ret0, err
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _ParentBridge.Contract.CHILDBLOCKINTERVAL(&_ParentBridge.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _ParentBridge.Contract.CHILDBLOCKINTERVAL(&_ParentBridge.CallOpts)
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeCaller) ChildBlocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _ParentBridge.contract.Call(opts, out, "childBlocks", arg0)
	return *ret, err
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeSession) ChildBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _ParentBridge.Contract.ChildBlocks(&_ParentBridge.CallOpts, arg0)
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeCallerSession) ChildBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _ParentBridge.Contract.ChildBlocks(&_ParentBridge.CallOpts, arg0)
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentBridge *ParentBridgeCaller) CoinCount(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "coinCount")
	return *ret0, err
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentBridge *ParentBridgeSession) CoinCount() (uint64, error) {
	return _ParentBridge.Contract.CoinCount(&_ParentBridge.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentBridge *ParentBridgeCallerSession) CoinCount() (uint64, error) {
	return _ParentBridge.Contract.CoinCount(&_ParentBridge.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) CurrentChildBlock() (*big.Int, error) {
	return _ParentBridge.Contract.CurrentChildBlock(&_ParentBridge.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _ParentBridge.Contract.CurrentChildBlock(&_ParentBridge.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) CurrentDepositBlock() (*big.Int, error) {
	return _ParentBridge.Contract.CurrentDepositBlock(&_ParentBridge.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _ParentBridge.Contract.CurrentDepositBlock(&_ParentBridge.CallOpts)
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentBridge *ParentBridgeCaller) GetChildBlock(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ParentBridge.contract.Call(opts, out, "getChildBlock", _blockNumber)
	return *ret0, *ret1, err
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentBridge *ParentBridgeSession) GetChildBlock(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _ParentBridge.Contract.GetChildBlock(&_ParentBridge.CallOpts, _blockNumber)
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentBridge *ParentBridgeCallerSession) GetChildBlock(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _ParentBridge.Contract.GetChildBlock(&_ParentBridge.CallOpts, _blockNumber)
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) GetNextDepositBlockIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "getNextDepositBlockIndex")
	return *ret0, err
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) GetNextDepositBlockIndex() (*big.Int, error) {
	return _ParentBridge.Contract.GetNextDepositBlockIndex(&_ParentBridge.CallOpts)
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) GetNextDepositBlockIndex() (*big.Int, error) {
	return _ParentBridge.Contract.GetNextDepositBlockIndex(&_ParentBridge.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x4b991bde.
//
// Solidity: function challenge(slotId uint64, blockNumber uint256, claimTxData bytes, proof bytes, sign bytes) returns()
func (_ParentBridge *ParentBridgeTransactor) Challenge(opts *bind.TransactOpts, slotId uint64, blockNumber *big.Int, claimTxData []byte, proof []byte, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "challenge", slotId, blockNumber, claimTxData, proof, sign)
}

// Challenge is a paid mutator transaction binding the contract method 0x4b991bde.
//
// Solidity: function challenge(slotId uint64, blockNumber uint256, claimTxData bytes, proof bytes, sign bytes) returns()
func (_ParentBridge *ParentBridgeSession) Challenge(slotId uint64, blockNumber *big.Int, claimTxData []byte, proof []byte, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.Challenge(&_ParentBridge.TransactOpts, slotId, blockNumber, claimTxData, proof, sign)
}

// Challenge is a paid mutator transaction binding the contract method 0x4b991bde.
//
// Solidity: function challenge(slotId uint64, blockNumber uint256, claimTxData bytes, proof bytes, sign bytes) returns()
func (_ParentBridge *ParentBridgeTransactorSession) Challenge(slotId uint64, blockNumber *big.Int, claimTxData []byte, proof []byte, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.Challenge(&_ParentBridge.TransactOpts, slotId, blockNumber, claimTxData, proof, sign)
}

// DepositFungible is a paid mutator transaction binding the contract method 0xd551c035.
//
// Solidity: function depositFungible(token address, from address, value uint256) returns()
func (_ParentBridge *ParentBridgeTransactor) DepositFungible(opts *bind.TransactOpts, token common.Address, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "depositFungible", token, from, value)
}

// DepositFungible is a paid mutator transaction binding the contract method 0xd551c035.
//
// Solidity: function depositFungible(token address, from address, value uint256) returns()
func (_ParentBridge *ParentBridgeSession) DepositFungible(token common.Address, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentBridge.Contract.DepositFungible(&_ParentBridge.TransactOpts, token, from, value)
}

// DepositFungible is a paid mutator transaction binding the contract method 0xd551c035.
//
// Solidity: function depositFungible(token address, from address, value uint256) returns()
func (_ParentBridge *ParentBridgeTransactorSession) DepositFungible(token common.Address, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentBridge.Contract.DepositFungible(&_ParentBridge.TransactOpts, token, from, value)
}

// DepositNonFungible is a paid mutator transaction binding the contract method 0x2791fd4c.
//
// Solidity: function depositNonFungible(token address, from address, tokenId uint256) returns()
func (_ParentBridge *ParentBridgeTransactor) DepositNonFungible(opts *bind.TransactOpts, token common.Address, from common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "depositNonFungible", token, from, tokenId)
}

// DepositNonFungible is a paid mutator transaction binding the contract method 0x2791fd4c.
//
// Solidity: function depositNonFungible(token address, from address, tokenId uint256) returns()
func (_ParentBridge *ParentBridgeSession) DepositNonFungible(token common.Address, from common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParentBridge.Contract.DepositNonFungible(&_ParentBridge.TransactOpts, token, from, tokenId)
}

// DepositNonFungible is a paid mutator transaction binding the contract method 0x2791fd4c.
//
// Solidity: function depositNonFungible(token address, from address, tokenId uint256) returns()
func (_ParentBridge *ParentBridgeTransactorSession) DepositNonFungible(token common.Address, from common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParentBridge.Contract.DepositNonFungible(&_ParentBridge.TransactOpts, token, from, tokenId)
}

// Exit is a paid mutator transaction binding the contract method 0x124135c9.
//
// Solidity: function exit(exitTxData bytes, exitProof bytes, exitBlock uint256, prevTxData bytes, prevProof bytes, prevBlock uint256, sign bytes) returns()
func (_ParentBridge *ParentBridgeTransactor) Exit(opts *bind.TransactOpts, exitTxData []byte, exitProof []byte, exitBlock *big.Int, prevTxData []byte, prevProof []byte, prevBlock *big.Int, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "exit", exitTxData, exitProof, exitBlock, prevTxData, prevProof, prevBlock, sign)
}

// Exit is a paid mutator transaction binding the contract method 0x124135c9.
//
// Solidity: function exit(exitTxData bytes, exitProof bytes, exitBlock uint256, prevTxData bytes, prevProof bytes, prevBlock uint256, sign bytes) returns()
func (_ParentBridge *ParentBridgeSession) Exit(exitTxData []byte, exitProof []byte, exitBlock *big.Int, prevTxData []byte, prevProof []byte, prevBlock *big.Int, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.Exit(&_ParentBridge.TransactOpts, exitTxData, exitProof, exitBlock, prevTxData, prevProof, prevBlock, sign)
}

// Exit is a paid mutator transaction binding the contract method 0x124135c9.
//
// Solidity: function exit(exitTxData bytes, exitProof bytes, exitBlock uint256, prevTxData bytes, prevProof bytes, prevBlock uint256, sign bytes) returns()
func (_ParentBridge *ParentBridgeTransactorSession) Exit(exitTxData []byte, exitProof []byte, exitBlock *big.Int, prevTxData []byte, prevProof []byte, prevBlock *big.Int, sign []byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.Exit(&_ParentBridge.TransactOpts, exitTxData, exitProof, exitBlock, prevTxData, prevProof, prevBlock, sign)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_ParentBridge *ParentBridgeTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_ParentBridge *ParentBridgeSession) Finalize() (*types.Transaction, error) {
	return _ParentBridge.Contract.Finalize(&_ParentBridge.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_ParentBridge *ParentBridgeTransactorSession) Finalize() (*types.Transaction, error) {
	return _ParentBridge.Contract.Finalize(&_ParentBridge.TransactOpts)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns(uint256)
func (_ParentBridge *ParentBridgeTransactor) SubmitBlock(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "submitBlock", _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns(uint256)
func (_ParentBridge *ParentBridgeSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.SubmitBlock(&_ParentBridge.TransactOpts, _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns(uint256)
func (_ParentBridge *ParentBridgeTransactorSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.SubmitBlock(&_ParentBridge.TransactOpts, _root)
}

// ParentBridgeBlockSubmitIterator is returned from FilterBlockSubmit and is used to iterate over the raw logs and unpacked data for BlockSubmit events raised by the ParentBridge contract.
type ParentBridgeBlockSubmitIterator struct {
	Event *ParentBridgeBlockSubmit // Event containing the contract specifics and raw log

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
func (it *ParentBridgeBlockSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeBlockSubmit)
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
		it.Event = new(ParentBridgeBlockSubmit)
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
func (it *ParentBridgeBlockSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeBlockSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeBlockSubmit represents a BlockSubmit event raised by the ParentBridge contract.
type ParentBridgeBlockSubmit struct {
	Root      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmit is a free log retrieval operation binding the contract event 0x79f92be16163b2592e15e92db55fab076ee3c8b88792ef547d89febb58170792.
//
// Solidity: event BlockSubmit(root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeFilterer) FilterBlockSubmit(opts *bind.FilterOpts) (*ParentBridgeBlockSubmitIterator, error) {

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "BlockSubmit")
	if err != nil {
		return nil, err
	}
	return &ParentBridgeBlockSubmitIterator{contract: _ParentBridge.contract, event: "BlockSubmit", logs: logs, sub: sub}, nil
}

// WatchBlockSubmit is a free log subscription operation binding the contract event 0x79f92be16163b2592e15e92db55fab076ee3c8b88792ef547d89febb58170792.
//
// Solidity: event BlockSubmit(root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeFilterer) WatchBlockSubmit(opts *bind.WatchOpts, sink chan<- *ParentBridgeBlockSubmit) (event.Subscription, error) {

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "BlockSubmit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeBlockSubmit)
				if err := _ParentBridge.contract.UnpackLog(event, "BlockSubmit", log); err != nil {
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

// ParentBridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ParentBridge contract.
type ParentBridgeDepositIterator struct {
	Event *ParentBridgeDeposit // Event containing the contract specifics and raw log

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
func (it *ParentBridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeDeposit)
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
		it.Event = new(ParentBridgeDeposit)
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
func (it *ParentBridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeDeposit represents a Deposit event raised by the ParentBridge contract.
type ParentBridgeDeposit struct {
	SlotId             uint64
	Depositor          common.Address
	DepositBlockNumber *big.Int
	Token              common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x673daafb22e46978f66bf78c4b57c5ca0f03f8af2c773103438514ec289a0bc9.
//
// Solidity: event Deposit(slotId indexed uint64, depositor indexed address, depositBlockNumber indexed uint256, token address, amount uint256)
func (_ParentBridge *ParentBridgeFilterer) FilterDeposit(opts *bind.FilterOpts, slotId []uint64, depositor []common.Address, depositBlockNumber []*big.Int) (*ParentBridgeDepositIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockNumberRule []interface{}
	for _, depositBlockNumberItem := range depositBlockNumber {
		depositBlockNumberRule = append(depositBlockNumberRule, depositBlockNumberItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "Deposit", slotIdRule, depositorRule, depositBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeDepositIterator{contract: _ParentBridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x673daafb22e46978f66bf78c4b57c5ca0f03f8af2c773103438514ec289a0bc9.
//
// Solidity: event Deposit(slotId indexed uint64, depositor indexed address, depositBlockNumber indexed uint256, token address, amount uint256)
func (_ParentBridge *ParentBridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ParentBridgeDeposit, slotId []uint64, depositor []common.Address, depositBlockNumber []*big.Int) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockNumberRule []interface{}
	for _, depositBlockNumberItem := range depositBlockNumber {
		depositBlockNumberRule = append(depositBlockNumberRule, depositBlockNumberItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "Deposit", slotIdRule, depositorRule, depositBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeDeposit)
				if err := _ParentBridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParentBridgeExitFinalizedIterator is returned from FilterExitFinalized and is used to iterate over the raw logs and unpacked data for ExitFinalized events raised by the ParentBridge contract.
type ParentBridgeExitFinalizedIterator struct {
	Event *ParentBridgeExitFinalized // Event containing the contract specifics and raw log

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
func (it *ParentBridgeExitFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeExitFinalized)
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
		it.Event = new(ParentBridgeExitFinalized)
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
func (it *ParentBridgeExitFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeExitFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeExitFinalized represents a ExitFinalized event raised by the ParentBridge contract.
type ParentBridgeExitFinalized struct {
	SlotId uint64
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitFinalized is a free log retrieval operation binding the contract event 0x6a9dd51c34b2e993c5129cce29598c784eb6f8b1f0832b9bd1e7f09f3daba156.
//
// Solidity: event ExitFinalized(slotId indexed uint64, owner indexed address, amount indexed uint256)
func (_ParentBridge *ParentBridgeFilterer) FilterExitFinalized(opts *bind.FilterOpts, slotId []uint64, owner []common.Address, amount []*big.Int) (*ParentBridgeExitFinalizedIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "ExitFinalized", slotIdRule, ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeExitFinalizedIterator{contract: _ParentBridge.contract, event: "ExitFinalized", logs: logs, sub: sub}, nil
}

// WatchExitFinalized is a free log subscription operation binding the contract event 0x6a9dd51c34b2e993c5129cce29598c784eb6f8b1f0832b9bd1e7f09f3daba156.
//
// Solidity: event ExitFinalized(slotId indexed uint64, owner indexed address, amount indexed uint256)
func (_ParentBridge *ParentBridgeFilterer) WatchExitFinalized(opts *bind.WatchOpts, sink chan<- *ParentBridgeExitFinalized, slotId []uint64, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "ExitFinalized", slotIdRule, ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeExitFinalized)
				if err := _ParentBridge.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
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

// ParentBridgeExitRejectedIterator is returned from FilterExitRejected and is used to iterate over the raw logs and unpacked data for ExitRejected events raised by the ParentBridge contract.
type ParentBridgeExitRejectedIterator struct {
	Event *ParentBridgeExitRejected // Event containing the contract specifics and raw log

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
func (it *ParentBridgeExitRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeExitRejected)
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
		it.Event = new(ParentBridgeExitRejected)
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
func (it *ParentBridgeExitRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeExitRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeExitRejected represents a ExitRejected event raised by the ParentBridge contract.
type ParentBridgeExitRejected struct {
	SlotId  uint64
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitRejected is a free log retrieval operation binding the contract event 0x172124f3161dc498f87d145ee7397b5063391ae141a8ed2e4170b71d1f1128d1.
//
// Solidity: event ExitRejected(slotId indexed uint64, claimer indexed address)
func (_ParentBridge *ParentBridgeFilterer) FilterExitRejected(opts *bind.FilterOpts, slotId []uint64, claimer []common.Address) (*ParentBridgeExitRejectedIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "ExitRejected", slotIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeExitRejectedIterator{contract: _ParentBridge.contract, event: "ExitRejected", logs: logs, sub: sub}, nil
}

// WatchExitRejected is a free log subscription operation binding the contract event 0x172124f3161dc498f87d145ee7397b5063391ae141a8ed2e4170b71d1f1128d1.
//
// Solidity: event ExitRejected(slotId indexed uint64, claimer indexed address)
func (_ParentBridge *ParentBridgeFilterer) WatchExitRejected(opts *bind.WatchOpts, sink chan<- *ParentBridgeExitRejected, slotId []uint64, claimer []common.Address) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "ExitRejected", slotIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeExitRejected)
				if err := _ParentBridge.contract.UnpackLog(event, "ExitRejected", log); err != nil {
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

// ParentBridgeExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the ParentBridge contract.
type ParentBridgeExitStartedIterator struct {
	Event *ParentBridgeExitStarted // Event containing the contract specifics and raw log

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
func (it *ParentBridgeExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeExitStarted)
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
		it.Event = new(ParentBridgeExitStarted)
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
func (it *ParentBridgeExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeExitStarted represents a ExitStarted event raised by the ParentBridge contract.
type ParentBridgeExitStarted struct {
	SlotId uint64
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xf40a6c02d2711d3d98c0dcb225f64eec06734f556e141a616f79a18fe9fb359e.
//
// Solidity: event ExitStarted(slotId indexed uint64, owner indexed address)
func (_ParentBridge *ParentBridgeFilterer) FilterExitStarted(opts *bind.FilterOpts, slotId []uint64, owner []common.Address) (*ParentBridgeExitStartedIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "ExitStarted", slotIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeExitStartedIterator{contract: _ParentBridge.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xf40a6c02d2711d3d98c0dcb225f64eec06734f556e141a616f79a18fe9fb359e.
//
// Solidity: event ExitStarted(slotId indexed uint64, owner indexed address)
func (_ParentBridge *ParentBridgeFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *ParentBridgeExitStarted, slotId []uint64, owner []common.Address) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "ExitStarted", slotIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeExitStarted)
				if err := _ParentBridge.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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
