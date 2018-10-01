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
const ParentBridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"coinCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_ITERATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childBlocks\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHALLENGE_PERIOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockSubmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ExitRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExitFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNonFungible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositFungible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"exitTxData\",\"type\":\"bytes\"},{\"name\":\"exitProof\",\"type\":\"bytes\"},{\"name\":\"exitBlock\",\"type\":\"uint256\"},{\"name\":\"prevTxData\",\"type\":\"bytes\"},{\"name\":\"prevProof\",\"type\":\"bytes\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"claimTxData\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slotIds\",\"type\":\"uint64[]\"}],\"name\":\"finalizeMany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextDepositBlockIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"getCoinBySlotId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"getExitBySlotId\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"getCoinByCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ParentBridgeBin is the compiled bytecode used for deploying new contracts.
const ParentBridgeBin = `0x60806040526004805467ffffffffffffffff1916905534801561002157600080fd5b5060008054600160a060020a0319163317905561003c610084565b604051809103906000f080158015610058573d6000803e3d6000fd5b5060018054600160a060020a031916600160a060020a03929092169190911790556103e8600255610095565b6040516104f6806200224683390190565b6121a180620000a56000396000f3006080604052600436106101265763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662a3216c811461012b578063124135c9146101525780632791fd4c146102b1578063359bc19e146102db5780634b991bde146102fd5780634d88a8cf146103e35780635b1052f11461043f578063678fd289146104c85780636faf25ed146104fa578063715018a61461050f5780637a95f1e8146105245780638da5cb5b1461053957806397ef999b1461056a5780639a09a8ef1461059b578063a831fa07146105bd578063a98c7f2c146105d2578063baa47694146105e7578063bb2dc863146105ff578063c3a079ed14610617578063cb9a32511461062c578063d551c03514610681578063f2fde38b146106ab575b600080fd5b34801561013757600080fd5b506101406106cc565b60408051918252519081900360200190f35b34801561015e57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102af94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a90999401975091955091820193509150819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a9099940197509195509182019350915081908401838280828437509497506106fc9650505050505050565b005b3480156102bd57600080fd5b506102af600160a060020a0360043581169060243516604435610748565b3480156102e757600080fd5b506102af67ffffffffffffffff600435166107e5565b34801561030957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102af94823567ffffffffffffffff1694602480359536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750610afc9650505050505050565b3480156103ef57600080fd5b5061040567ffffffffffffffff60043516610dff565b60408051600160a060020a039687168152602081019590955292909416838301526060830152608082019290925290519081900360a00190f35b34801561044b57600080fd5b5061046167ffffffffffffffff60043516610e48565b6040518087600181111561047157fe5b60ff16815260200186600160a060020a0316600160a060020a031681526020018581526020018481526020018381526020018260028111156104af57fe5b60ff168152602001965050505050505060405180910390f35b3480156104d457600080fd5b506104dd610ea2565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561050657600080fd5b50610140610eb2565b34801561051b57600080fd5b506102af610eb7565b34801561053057600080fd5b50610140610f16565b34801561054557600080fd5b5061054e610f1c565b60408051600160a060020a039092168252519081900360200190f35b34801561057657600080fd5b50610582600435610f2b565b6040805192835260208301919091528051918290030190f35b3480156105a757600080fd5b5061046167ffffffffffffffff60043516610f44565b3480156105c957600080fd5b50610140610f89565b3480156105de57600080fd5b50610140610f8f565b3480156105f357600080fd5b506102af600435610f95565b34801561060b57600080fd5b5061058260043561103d565b34801561062357600080fd5b50610140611057565b34801561063857600080fd5b50604080516020600480358082013583810280860185019096528085526102af9536959394602494938501929182918501908490808284375094975061105c9650505050505050565b34801561068d57600080fd5b506102af600160a060020a03600435811690602435166044356110f3565b3480156106b757600080fd5b506102af600160a060020a036004351661111d565b60006106f76003546106eb6103e860025461114090919063ffffffff16565b9063ffffffff61115216565b905090565b6103e8850615610715576107108588611165565b61073f565b6000858152600560205260408082205484835291205461073f9187918a91908a908990898861132b565b50505050505050565b604080517f42842e0e000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152306024830152604482018490529151918516916342842e0e9160648082019260009290919082900301818387803b1580156107b957600080fd5b505af11580156107cd573d6000803e3d6000fd5b505050506107e0600184848460016115f6565b505050565b67ffffffffffffffff81166000908152600760205260408120908080806001600586015460ff16600281111561081757fe5b1461082157610af4565b600a85015460c89061083a90439063ffffffff61114016565b1161084457610af4565b60058501805460ff191660021790556006850154855474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a03909216919091021780865560009060ff16600181111561089957fe5b141561096e57600180860154865460048089015467ffffffffffffffff8b1660009081526007602081905260408220805474ffffffffffffffffffffffffffffffffffffffffff191681559687018054600160a060020a0319908116909155600288018390556003880183905593870182905560058701805460ff1916905560068701805485169055860181905560088601805490931690925560098501829055600a909401556101009004600160a060020a03908116965091945016915061096982858563ffffffff6118f516565b610aac565b50600180850154855460028088015467ffffffffffffffff8a166000908152600760208190526040808320805474ffffffffffffffffffffffffffffffffffffffffff191681559788018054600160a060020a031990811690915594880183905560038801839055600480890184905560058901805460ff191690556006890180548716905591880183905560088801805490951690945560098701829055600a90960181905582517f42842e0e0000000000000000000000000000000000000000000000000000000081523096810196909652610100909304600160a060020a03908116602487018190526044870183905292519298509096509092169283926342842e0e9260648084019391929182900301818387803b158015610a9357600080fd5b505af1158015610aa7573d6000803e3d6000fd5b505050505b8454604051610100909104600160a060020a03169067ffffffffffffffff8816907f9dfb93f0971f4e0ce2a7420e2cb0170877d6c1bf4c8726447b241d934bd2a94d90600090a35b505050505050565b610b04612116565b6000600167ffffffffffffffff881660009081526007602052604090206005015460ff166002811115610b3357fe5b14610b88576040805160e560020a62461bcd02815260206004820152601f60248201527f6f6e6c792065786974696e6720636f696e2063616e206368616c6c656e676500604482015290519081900360640190fd5b610b91856119a8565b67ffffffffffffffff808916600081815260076020526040902083519395509350911614610c09576040805160e560020a62461bcd02815260206004820152600c60248201527f696e76616c696420736c6f740000000000000000000000000000000000000000604482015290519081900360640190fd5b60088101546060830151600160a060020a0390911690610c2f908563ffffffff611abd16565b600160a060020a031614610c8d576040805160e560020a62461bcd02815260206004820152600b60248201527f696e76616c696420736967000000000000000000000000000000000000000000604482015290519081900360640190fd5b610c978787611b92565b15610ca157610d00565b6007810154602083015114610d00576040805160e560020a62461bcd02815260206004820152601160248201527f696e76616c6964207265666572656e6365000000000000000000000000000000604482015290519081900360640190fd5b600086815260056020526040902054610d1b90839086611bdf565b1515610d71576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b60058101805460ff1916905567ffffffffffffffff87166000818152600760208190526040808320600681018054600160a060020a031990811690915592810184905560088101805490931690925560098201839055600a909101829055513392917f172124f3161dc498f87d145ee7397b5063391ae141a8ed2e4170b71d1f1128d191a350505050505050565b67ffffffffffffffff16600090815260076020819052604090912060068101549181015460088201546009830154600a90930154600160a060020a039485169592949091169291565b67ffffffffffffffff1660009081526007602090815260408083208054600482015460038301548087526005958690529390952060010154939091015460ff80831696610100909304600160a060020a0316959492911690565b60045467ffffffffffffffff1681565b606481565b600054600160a060020a03163314610ece57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a260008054600160a060020a0319169055565b60025481565b600054600160a060020a031681565b6005602052600090815260409020805460019091015482565b67ffffffffffffffff808216600090815260066020526040812054909182918291829182918291610f759116610e48565b949c939b5091995097509550909350915050565b6103e881565b60035481565b600054600160a060020a03163314610fac57600080fd5b60408051808201825282815242602080830182815260028054600090815260058452869020945185559051600190940193909355915483519081529182018490528183015290517fc2b832dd3f70a2ac753587689afde9aff77239ca272e801b1577e290ca8bcf179181900360600190a1600254611032906103e863ffffffff61115216565b600255506001600355565b600090815260056020526040902080546001909101549091565b60c881565b600060648251111515156110ba576040805160e560020a62461bcd02815260206004820152600960248201527f676173206c696d69740000000000000000000000000000000000000000000000604482015290519081900360640190fd5b5060005b81518110156110ef576110e782828151811015156110d857fe5b906020019060200201516107e5565b6001016110be565b5050565b61110e600160a060020a03841683308463ffffffff611d0516565b6107e0600084846000856115f6565b600054600160a060020a0316331461113457600080fd5b61113d81611db0565b50565b60008282111561114c57fe5b50900390565b8181018281101561115f57fe5b92915050565b61116d612116565b6000611178836119a8565b805167ffffffffffffffff1660009081526007602052604090819020908201519193509150600160a060020a031633146111fc576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e206578697400000000000000000000000000604482015290519081900360640190fd5b60018101546002820154604080516c01000000000000000000000000338102602080840191909152600160a060020a03909516026034820152604880820193909352815180820390930183526068019081905281519192909182918401908083835b6020831061127d5780518252601f19909201916020918201910161125e565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008a8152600590925292902054909114925061130f915050576040805160e560020a62461bcd02815260206004820152601760248201527f6d656d6265727368697020636865636b206661696c6564000000000000000000604482015290519081900360640190fd5b6113258260000151836040015186600080611e20565b50505050565b611333612116565b61133b612116565b611344896119a8565b915061134f866119a8565b6040830151909150600160a060020a031633146113b6576040805160e560020a62461bcd02815260206004820152600d60248201527f696e76616c6964206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b8151815167ffffffffffffffff90811691161461141d576040805160e560020a62461bcd02815260206004820152600e60248201527f696e76616c696420736c6f744964000000000000000000000000000000000000604482015290519081900360640190fd5b6060820151611432908463ffffffff611abd16565b600160a060020a03168160400151600160a060020a03161415156114a0576040805160e560020a62461bcd02815260206004820152600b60248201527f696e76616c696420736967000000000000000000000000000000000000000000604482015290519081900360640190fd5b602080830151600090815260059091526040902054851461150b576040805160e560020a62461bcd02815260206004820152600f60248201527f696e76616c696420747820646174610000000000000000000000000000000000604482015290519081900360640190fd5b611516828989611bdf565b151561156c576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b611577818686611bdf565b15156115cd576040805160e560020a62461bcd02815260206004820152601660248201527f696e636c7573696f6e20636865636b206661696c656400000000000000000000604482015290519081900360640190fd5b6115ea826000015183604001518c84604001518660200151611e20565b50505050505050505050565b6000806000806103e860035410151561167f576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c7920616c6c6f7720757020746f2031303030206465706f73697473207060448201527f6572206368696c6420626c6f636b2e0000000000000000000000000000000000606482015290519081900360840190fd5b8688876040516020018084600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140183600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140182815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083106117205780518252601f199092019160209182019101611701565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350837801000000000000000000000000000000000000000000000000900492506117766106cc565b60408051808201825286815242602080830191825260008581526005909152929092209051815590516001918201556003549193506117bb919063ffffffff61115216565b6003555067ffffffffffffffff8216600090815260076020526040902080548990829060ff1916600183818111156117ef57fe5b0217905550805474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a03898116820292909217835560018084018054600160a060020a0319168c8516179055600284018990556003840185815560048086018a905560058601805460ff19169055805467ffffffffffffffff908116600090815260066020908152604091829020805467ffffffffffffffff199081168d86169081179092558554908116908516909701909316959095179092558654925482519081529151949092049094169390927f88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d929181900390910190a3505050505050505050565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561197157600080fd5b505af1158015611985573d6000803e3d6000fd5b505050506040513d602081101561199b57600080fd5b505115156107e057600080fd5b6119b0612116565b60606119ba612116565b6119d460036119c886611eec565b9063ffffffff611f0f16565b91506119f78260008151811015156119e857fe5b90602001906020020151611f9d565b67ffffffffffffffff1681528151611a1690839060019081106119e857fe5b60208201528151611a3d9083906002908110611a2e57fe5b90602001906020020151611fc1565b600160a060020a0316604080830191909152518451859190819060208401908083835b60208310611a7f5780518252601f199092019160209182019101611a60565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120606085015250919350839150505b5050919050565b60008060008084516041141515611ad75760009350611b89565b50505060208201516040830151606084015160001a601b60ff82161015611afc57601b015b8060ff16601b14158015611b1457508060ff16601c14155b15611b225760009350611b89565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015611b7c573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b67ffffffffffffffff821660009081526007602081905260408220015482108015611bd8575067ffffffffffffffff831660009081526007602052604090206009015482115b9392505050565b600154606084015184516040517ff586df65000000000000000000000000000000000000000000000000000000008152600481018381526024820187905267ffffffffffffffff83166044830152608060648301908152865160848401528651600096600160a060020a03169563f586df659590948a9491938a93909160a4019060208501908083838e5b83811015611c82578181015183820152602001611c6a565b50505050905090810190601f168015611caf5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015611cd157600080fd5b505af1158015611ce5573d6000803e3d6000fd5b505050506040513d6020811015611cfb57600080fd5b5051949350505050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b158015611d7957600080fd5b505af1158015611d8d573d6000803e3d6000fd5b505050506040513d6020811015611da357600080fd5b5051151561132557600080fd5b600160a060020a0381161515611dc557600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a031916600160a060020a0392909216919091179055565b67ffffffffffffffff85166000818152600760208181526040808420815160a081018352600160a060020a038b81168083529482018b9052898116828501819052606083018a9052436080909301839052600684018054600160a060020a0319908116909717908190559684018c90556008840180549096161790945560098201889055600a82015560058101805460ff191660011790559051909492909116927ff40a6c02d2711d3d98c0dcb225f64eec06734f556e141a616f79a18fe9fb359e91a3505050505050565b611ef461213d565b50805160408051808201909152602092830181529182015290565b6060611f19612154565b600083604051908082528060200260200182016040528015611f5557816020015b611f4261213d565b815260200190600190039081611f3a5790505b509250611f6185611fe6565b91505b83811015611f9557611f758261200b565b8382815181101515611f8357fe5b60209081029091010152600101611f64565b505092915050565b6000806000611fab8461203d565b90516020919091036101000a9004949350505050565b600080611fcd8361203d565b50516c0100000000000000000000000090049392505050565b611fee612154565b6000611ff9836120a0565b83519383529092016020820152919050565b61201361213d565b60208201516000612023826120e6565b828452602080850182905292019390910192909252919050565b805180516000918291821a9082608083101561205f5781945060019350612098565b60b883101561207d5760018660200151039350816001019450612098565b60b78303905080600187602001510303935080820160010194505b505050915091565b8051805160009190821a9060808210156120bd5760009250611ab6565b60b88210806120d8575060c082101580156120d8575060f882105b15611ab65760019250611ab6565b8051600090811a60808110156120ff5760019150612110565b60b881101561211057607e19810191505b50919050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b604080518082019091526000808252602082015290565b60606040519081016040528061216861213d565b81526020016000815250905600a165627a7a72305820e4a182ebfbbf2194385de78e02a82711f7872e21dd7790d98d5a555bfd3cf20b0029608060405234801561001057600080fd5b507f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56360005561004a6001604064010000000061004f810204565b61012c565b815b60ff8083169082161161012757600060ff6000198301166041811061007257fe5b0154600060ff6000198401166041811061008857fe5b0154604080516020808201949094528082019290925280518083038201815260609092019081905281519192909182918401908083835b602083106100de5780518252601f1990920191602091820191016100bf565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060008260ff1660418110151561011d57fe5b0155600101610051565b505050565b6103bb8061013b6000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166348419ad8811461005b578063509a7e5414610085578063f586df65146100ef575b600080fd5b34801561006757600080fd5b50610073600435610170565b60408051918252519081900360200190f35b34801561009157600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610073948235946024803567ffffffffffffffff16953695946064949201919081908401838280828437509497506101849650505050505050565b3480156100fb57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261015c94803594602480359567ffffffffffffffff6044351695369560849493019181908401838280828437509497506103759650505050505050565b604080519115158252519081900360200190f35b6000816041811061017d57fe5b0154905081565b6000806000806000806020600888510381151561019d57fe5b061580156101ae5750610808875111155b15156101b957600080fd5b50505050602083015185906008907801000000000000000000000000000000000000000000000000900460005b60408110156103685760018216151561020f576000816041811061020657fe5b01549450610230565b6020830192508261ffff1687511015151561022957600080fd5b8287015194505b6001881615156102bf57604080516020808201879052818301889052825180830384018152606090920192839052815191929182918401908083835b6020831061028b5780518252601f19909201916020918201910161026c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350610340565b604080516020808201889052818301879052825180830384018152606090920192839052815191929182918401908083835b602083106103105780518252601f1990920191602091820191016102f1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902093505b600267ffffffffffffffff8316049150600267ffffffffffffffff89160497506001016101e6565b5091979650505050505050565b600080610383868585610184565b909414959450505050505600a165627a7a723058205c5370b921fcf49a82addfb6c6361c436db0da74f1098d67b7f52bca25cb8d340029`

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

// MAXITERATION is a free data retrieval call binding the contract method 0x6faf25ed.
//
// Solidity: function MAX_ITERATION() constant returns(uint256)
func (_ParentBridge *ParentBridgeCaller) MAXITERATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "MAX_ITERATION")
	return *ret0, err
}

// MAXITERATION is a free data retrieval call binding the contract method 0x6faf25ed.
//
// Solidity: function MAX_ITERATION() constant returns(uint256)
func (_ParentBridge *ParentBridgeSession) MAXITERATION() (*big.Int, error) {
	return _ParentBridge.Contract.MAXITERATION(&_ParentBridge.CallOpts)
}

// MAXITERATION is a free data retrieval call binding the contract method 0x6faf25ed.
//
// Solidity: function MAX_ITERATION() constant returns(uint256)
func (_ParentBridge *ParentBridgeCallerSession) MAXITERATION() (*big.Int, error) {
	return _ParentBridge.Contract.MAXITERATION(&_ParentBridge.CallOpts)
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

// GetCoinByCount is a free data retrieval call binding the contract method 0x9a09a8ef.
//
// Solidity: function getCoinByCount(count uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeCaller) GetCoinByCount(opts *bind.CallOpts, count uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ParentBridge.contract.Call(opts, out, "getCoinByCount", count)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCoinByCount is a free data retrieval call binding the contract method 0x9a09a8ef.
//
// Solidity: function getCoinByCount(count uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeSession) GetCoinByCount(count uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _ParentBridge.Contract.GetCoinByCount(&_ParentBridge.CallOpts, count)
}

// GetCoinByCount is a free data retrieval call binding the contract method 0x9a09a8ef.
//
// Solidity: function getCoinByCount(count uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeCallerSession) GetCoinByCount(count uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _ParentBridge.Contract.GetCoinByCount(&_ParentBridge.CallOpts, count)
}

// GetCoinBySlotId is a free data retrieval call binding the contract method 0x5b1052f1.
//
// Solidity: function getCoinBySlotId(slotId uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeCaller) GetCoinBySlotId(opts *bind.CallOpts, slotId uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ParentBridge.contract.Call(opts, out, "getCoinBySlotId", slotId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCoinBySlotId is a free data retrieval call binding the contract method 0x5b1052f1.
//
// Solidity: function getCoinBySlotId(slotId uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeSession) GetCoinBySlotId(slotId uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _ParentBridge.Contract.GetCoinBySlotId(&_ParentBridge.CallOpts, slotId)
}

// GetCoinBySlotId is a free data retrieval call binding the contract method 0x5b1052f1.
//
// Solidity: function getCoinBySlotId(slotId uint64) constant returns(uint8, address, uint256, uint256, uint256, uint8)
func (_ParentBridge *ParentBridgeCallerSession) GetCoinBySlotId(slotId uint64) (uint8, common.Address, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _ParentBridge.Contract.GetCoinBySlotId(&_ParentBridge.CallOpts, slotId)
}

// GetExitBySlotId is a free data retrieval call binding the contract method 0x4d88a8cf.
//
// Solidity: function getExitBySlotId(slotId uint64) constant returns(address, uint256, address, uint256, uint256)
func (_ParentBridge *ParentBridgeCaller) GetExitBySlotId(opts *bind.CallOpts, slotId uint64) (common.Address, *big.Int, common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _ParentBridge.contract.Call(opts, out, "getExitBySlotId", slotId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetExitBySlotId is a free data retrieval call binding the contract method 0x4d88a8cf.
//
// Solidity: function getExitBySlotId(slotId uint64) constant returns(address, uint256, address, uint256, uint256)
func (_ParentBridge *ParentBridgeSession) GetExitBySlotId(slotId uint64) (common.Address, *big.Int, common.Address, *big.Int, *big.Int, error) {
	return _ParentBridge.Contract.GetExitBySlotId(&_ParentBridge.CallOpts, slotId)
}

// GetExitBySlotId is a free data retrieval call binding the contract method 0x4d88a8cf.
//
// Solidity: function getExitBySlotId(slotId uint64) constant returns(address, uint256, address, uint256, uint256)
func (_ParentBridge *ParentBridgeCallerSession) GetExitBySlotId(slotId uint64) (common.Address, *big.Int, common.Address, *big.Int, *big.Int, error) {
	return _ParentBridge.Contract.GetExitBySlotId(&_ParentBridge.CallOpts, slotId)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ParentBridge *ParentBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ParentBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ParentBridge *ParentBridgeSession) Owner() (common.Address, error) {
	return _ParentBridge.Contract.Owner(&_ParentBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ParentBridge *ParentBridgeCallerSession) Owner() (common.Address, error) {
	return _ParentBridge.Contract.Owner(&_ParentBridge.CallOpts)
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

// Finalize is a paid mutator transaction binding the contract method 0x359bc19e.
//
// Solidity: function finalize(slotId uint64) returns()
func (_ParentBridge *ParentBridgeTransactor) Finalize(opts *bind.TransactOpts, slotId uint64) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "finalize", slotId)
}

// Finalize is a paid mutator transaction binding the contract method 0x359bc19e.
//
// Solidity: function finalize(slotId uint64) returns()
func (_ParentBridge *ParentBridgeSession) Finalize(slotId uint64) (*types.Transaction, error) {
	return _ParentBridge.Contract.Finalize(&_ParentBridge.TransactOpts, slotId)
}

// Finalize is a paid mutator transaction binding the contract method 0x359bc19e.
//
// Solidity: function finalize(slotId uint64) returns()
func (_ParentBridge *ParentBridgeTransactorSession) Finalize(slotId uint64) (*types.Transaction, error) {
	return _ParentBridge.Contract.Finalize(&_ParentBridge.TransactOpts, slotId)
}

// FinalizeMany is a paid mutator transaction binding the contract method 0xcb9a3251.
//
// Solidity: function finalizeMany(slotIds uint64[]) returns()
func (_ParentBridge *ParentBridgeTransactor) FinalizeMany(opts *bind.TransactOpts, slotIds []uint64) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "finalizeMany", slotIds)
}

// FinalizeMany is a paid mutator transaction binding the contract method 0xcb9a3251.
//
// Solidity: function finalizeMany(slotIds uint64[]) returns()
func (_ParentBridge *ParentBridgeSession) FinalizeMany(slotIds []uint64) (*types.Transaction, error) {
	return _ParentBridge.Contract.FinalizeMany(&_ParentBridge.TransactOpts, slotIds)
}

// FinalizeMany is a paid mutator transaction binding the contract method 0xcb9a3251.
//
// Solidity: function finalizeMany(slotIds uint64[]) returns()
func (_ParentBridge *ParentBridgeTransactorSession) FinalizeMany(slotIds []uint64) (*types.Transaction, error) {
	return _ParentBridge.Contract.FinalizeMany(&_ParentBridge.TransactOpts, slotIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParentBridge *ParentBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParentBridge *ParentBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParentBridge.Contract.RenounceOwnership(&_ParentBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParentBridge *ParentBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParentBridge.Contract.RenounceOwnership(&_ParentBridge.TransactOpts)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentBridge *ParentBridgeTransactor) SubmitBlock(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "submitBlock", _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentBridge *ParentBridgeSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.SubmitBlock(&_ParentBridge.TransactOpts, _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentBridge *ParentBridgeTransactorSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentBridge.Contract.SubmitBlock(&_ParentBridge.TransactOpts, _root)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ParentBridge *ParentBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ParentBridge.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ParentBridge *ParentBridgeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ParentBridge.Contract.TransferOwnership(&_ParentBridge.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ParentBridge *ParentBridgeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ParentBridge.Contract.TransferOwnership(&_ParentBridge.TransactOpts, _newOwner)
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
	BlockNumber *big.Int
	Root        [32]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmit is a free log retrieval operation binding the contract event 0xc2b832dd3f70a2ac753587689afde9aff77239ca272e801b1577e290ca8bcf17.
//
// Solidity: e BlockSubmit(blockNumber uint256, root bytes32, timestamp uint256)
func (_ParentBridge *ParentBridgeFilterer) FilterBlockSubmit(opts *bind.FilterOpts) (*ParentBridgeBlockSubmitIterator, error) {

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "BlockSubmit")
	if err != nil {
		return nil, err
	}
	return &ParentBridgeBlockSubmitIterator{contract: _ParentBridge.contract, event: "BlockSubmit", logs: logs, sub: sub}, nil
}

// WatchBlockSubmit is a free log subscription operation binding the contract event 0xc2b832dd3f70a2ac753587689afde9aff77239ca272e801b1577e290ca8bcf17.
//
// Solidity: e BlockSubmit(blockNumber uint256, root bytes32, timestamp uint256)
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
	SlotId      uint64
	Owner       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d.
//
// Solidity: e Deposit(slotId indexed uint64, owner indexed address, blockNumber uint256)
func (_ParentBridge *ParentBridgeFilterer) FilterDeposit(opts *bind.FilterOpts, slotId []uint64, owner []common.Address) (*ParentBridgeDepositIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "Deposit", slotIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeDepositIterator{contract: _ParentBridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d.
//
// Solidity: e Deposit(slotId indexed uint64, owner indexed address, blockNumber uint256)
func (_ParentBridge *ParentBridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ParentBridgeDeposit, slotId []uint64, owner []common.Address) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "Deposit", slotIdRule, ownerRule)
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
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitFinalized is a free log retrieval operation binding the contract event 0x9dfb93f0971f4e0ce2a7420e2cb0170877d6c1bf4c8726447b241d934bd2a94d.
//
// Solidity: e ExitFinalized(slotId indexed uint64, owner indexed address)
func (_ParentBridge *ParentBridgeFilterer) FilterExitFinalized(opts *bind.FilterOpts, slotId []uint64, owner []common.Address) (*ParentBridgeExitFinalizedIterator, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "ExitFinalized", slotIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeExitFinalizedIterator{contract: _ParentBridge.contract, event: "ExitFinalized", logs: logs, sub: sub}, nil
}

// WatchExitFinalized is a free log subscription operation binding the contract event 0x9dfb93f0971f4e0ce2a7420e2cb0170877d6c1bf4c8726447b241d934bd2a94d.
//
// Solidity: e ExitFinalized(slotId indexed uint64, owner indexed address)
func (_ParentBridge *ParentBridgeFilterer) WatchExitFinalized(opts *bind.WatchOpts, sink chan<- *ParentBridgeExitFinalized, slotId []uint64, owner []common.Address) (event.Subscription, error) {

	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "ExitFinalized", slotIdRule, ownerRule)
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
// Solidity: e ExitRejected(slotId indexed uint64, claimer indexed address)
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
// Solidity: e ExitRejected(slotId indexed uint64, claimer indexed address)
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
// Solidity: e ExitStarted(slotId indexed uint64, owner indexed address)
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
// Solidity: e ExitStarted(slotId indexed uint64, owner indexed address)
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

// ParentBridgeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ParentBridge contract.
type ParentBridgeOwnershipRenouncedIterator struct {
	Event *ParentBridgeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ParentBridgeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeOwnershipRenounced)
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
		it.Event = new(ParentBridgeOwnershipRenounced)
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
func (it *ParentBridgeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeOwnershipRenounced represents a OwnershipRenounced event raised by the ParentBridge contract.
type ParentBridgeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ParentBridge *ParentBridgeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ParentBridgeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeOwnershipRenouncedIterator{contract: _ParentBridge.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ParentBridge *ParentBridgeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ParentBridgeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeOwnershipRenounced)
				if err := _ParentBridge.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParentBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParentBridge contract.
type ParentBridgeOwnershipTransferredIterator struct {
	Event *ParentBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ParentBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBridgeOwnershipTransferred)
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
		it.Event = new(ParentBridgeOwnershipTransferred)
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
func (it *ParentBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ParentBridge contract.
type ParentBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ParentBridge *ParentBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParentBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParentBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParentBridgeOwnershipTransferredIterator{contract: _ParentBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ParentBridge *ParentBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParentBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParentBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBridgeOwnershipTransferred)
				if err := _ParentBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
