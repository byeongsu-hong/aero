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

// ChildBridgeABI is the input ABI used to generate the binding from.
const ChildBridgeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastBlockOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingTransactions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"createTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveWitness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBlockNumber\",\"type\":\"uint256\"}],\"name\":\"submitNewBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"parentToken\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"which\",\"type\":\"uint8\"}],\"name\":\"submitDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"exitor\",\"type\":\"address\"},{\"name\":\"parentToken\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"which\",\"type\":\"uint8\"}],\"name\":\"submitWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPendingTransactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChildBridgeBin is the compiled bytecode used for deploying new contracts.
const ChildBridgeBin = `0x608060405234801561001057600080fd5b506040516124fe3803806124fe833981016040528051602082015160008054600160a060020a03191633179055908201910181813061004d610171565b600160a060020a0382166040820152606080825284519082015283518190602080830191608084019188019080838360005b8381101561009757818101518382015260200161007f565b50505050905090810190601f1680156100c45780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156100f75781810151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b5095505050505050604051809103906000f080158015610148573d6000803e3d6000fd5b5060018054600160a060020a031916600160a060020a0392909216919091179055506101819050565b604051611449806110b583390190565b610f25806101906000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633efd0ac281146100c957806363a8374d146100f3578063642f2eaf1461010b578063715018a6146102045780638da5cb5b1461021b578063a560cd7e1461024c578063b3455f8914610289578063b3d3a508146102e7578063cafb00a01461031b578063ddbd5d3f14610333578063e023e75914610348578063f2fde38b14610385578063fc0c546a146103a6575b600080fd5b3480156100d557600080fd5b506100e16004356103bb565b60408051918252519081900360200190f35b3480156100ff57600080fd5b506100e16004356103cd565b34801561011757600080fd5b506101236004356103ec565b6040518087600181111561013357fe5b60ff1681526020018667ffffffffffffffff1667ffffffffffffffff16815260200185815260200184600160a060020a0316600160a060020a0316815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101c45781810151838201526020016101ac565b50505050905090810190601f1680156101f15780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561021057600080fd5b506102196104c3565b005b34801561022757600080fd5b5061023061052f565b60408051600160a060020a039092168252519081900360200190f35b34801561025857600080fd5b50610219600160a060020a036004358116906024351667ffffffffffffffff6044351660643560ff6084351661053e565b34801561029557600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102199583359536956044949193909101919081908401838280828437509497506106c29650505050505050565b3480156102f357600080fd5b506100e1600160a060020a036004358116906024351667ffffffffffffffff6044351661080a565b34801561032757600080fd5b50610219600435610ac1565b34801561033f57600080fd5b506100e1610ba0565b34801561035457600080fd5b50610219600160a060020a036004358116906024351667ffffffffffffffff6044351660643560ff60843516610ba7565b34801561039157600080fd5b50610219600160a060020a0360043516610cc3565b3480156103b257600080fd5b50610230610ce6565b60036020526000908152604090205481565b60048054829081106103db57fe5b600091825260209091200154905081565b60026020818152600092835260409283902080546001808301548386015460038501546004860180548a516101009682161587026000190190911699909904601f810189900489028a018901909a5289895260ff8616999490950467ffffffffffffffff16979296600160a060020a039283169692909116949193908301828280156104b95780601f1061048e576101008083540402835291602001916104b9565b820191906000526020600020905b81548152906001019060200180831161049c57829003601f168201915b5050505050905086565b600054600160a060020a031633146104da57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600080548190600160a060020a0316331461055857600080fd5b600083600181111561056657fe5b141561061057600154604080517fea922837000000000000000000000000000000000000000000000000000000008152600160a060020a038a81166004830152602482018890529151919092169350839163ea9228379160448083019260209291908290030181600087803b1580156105de57600080fd5b505af11580156105f2573d6000803e3d6000fd5b505050506040513d602081101561060857600080fd5b506106b99050565b50600154604080517f079cec9f000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015267ffffffffffffffff8816602483015291519190921691829163079cec9f916044808201926020929091908290030181600087803b15801561068c57600080fd5b505af11580156106a0573d6000803e3d6000fd5b505050506040513d60208110156106b657600080fd5b50505b50505050505050565b600082815260026020526040812054610100900467ffffffffffffffff16151561074d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4e6f205472616e73616374696f6e20494420666f756e642e0000000000000000604482015290519081900360640190fd5b5060008281526002602052604090206003810154600160a060020a031661077a848463ffffffff610cf516565b600160a060020a0316146107ef57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f5369676e6174757265206d69736d617463682e00000000000000000000000000604482015290519081900360640190fd5b81516108049060048301906020850190610e47565b50505050565b60015460009060609082908190600160a060020a0316331461088d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4469726563742063616c6c206973206e6f7420616c6c6f7765642e0000000000604482015290519081900360640190fd5b67ffffffffffffffff85166000818152600360209081526040918290205482517ff857000000000000000000000000000000000000000000000000000000000000818401527f88000000000000000000000000000000000000000000000000000000000000006022820152780100000000000000000000000000000000000000000000000090940260238501527fa000000000000000000000000000000000000000000000000000000000000000602b850152602c8401527f9400000000000000000000000000000000000000000000000000000000000000604c8401526c01000000000000000000000000600160a060020a038a1602604d840152815160418185030181526061909301918290528251929550859282918401908083835b602083106109cb5780518252601f1990920191602091820191016109ac565b518151600019602094850361010090810a9190910191821691199290921617909152604080519590930185900390942060008181526002808452848220805467ffffffffffffffff909f1697880268ffffffffffffffff0019909f169e909e178e559581526003928390529283205460018d810191909155948c018054600160a060020a039e8f1673ffffffffffffffffffffffffffffffffffffffff199182161790915591909b0180549d909c169c169b909b17909955600480549182018155909952505050507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9094018290555092915050565b6000805481908190600160a060020a03163314610add57600080fd5b600092505b600454831015610b94576004805484908110610afa57fe5b600091825260208083209190910154808352600282526040808420805460ff19166001178082556101009081900467ffffffffffffffff90811687526003958690528387208b9055825495830154935194985091965090930490921692600160a060020a039092169185917f80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a19190a4600190920191610ae2565b61080460046000610ec5565b6004545b90565b600080548190600160a060020a03163314610bc157600080fd5b6000836001811115610bcf57fe5b1415610c4757600154604080517f9470b0bd000000000000000000000000000000000000000000000000000000008152600160a060020a038a811660048301526024820188905291519190921693508391639470b0bd9160448083019260209291908290030181600087803b1580156105de57600080fd5b50600154604080517f0e723a37000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015267ffffffffffffffff88166024830152915191909216918291630e723a37916044808201926020929091908290030181600087803b15801561068c57600080fd5b600054600160a060020a03163314610cda57600080fd5b610ce381610dca565b50565b600154600160a060020a031681565b60008060008084516041141515610d0f5760009350610dc1565b50505060208201516040830151606084015160001a601b60ff82161015610d3457601b015b8060ff16601b14158015610d4c57508060ff16601c14155b15610d5a5760009350610dc1565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610db4573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600160a060020a0381161515610ddf57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e8857805160ff1916838001178555610eb5565b82800160010185558215610eb5579182015b82811115610eb5578251825591602001919060010190610e9a565b50610ec1929150610edf565b5090565b5080546000825590600052602060002090810190610ce391905b610ba491905b80821115610ec15760008155600101610ee55600a165627a7a723058208028174714dac4a0c8d04988773d0d72f2e9d085b96a40e4a713e3587d8c93e0002960806040523480156200001157600080fd5b50604051620014493803806200144983398101604090815281516020830151918301519083019291909101908282620000737f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006200019a810204565b620000a77f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006200019a810204565b620000db7f4f558e79000000000000000000000000000000000000000000000000000000006401000000006200019a810204565b8151620000f090600590602085019062000207565b5080516200010690600690602084019062000207565b506200013b7f780e9d63000000000000000000000000000000000000000000000000000000006401000000006200019a810204565b6200016f7f5b5e139f000000000000000000000000000000000000000000000000000000006401000000006200019a810204565b5050600c8054600160a060020a031916600160a060020a039290921691909117905550620002ac9050565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001ca57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024a57805160ff19168380011785556200027a565b828001600101855582156200027a579182015b828111156200027a5782518255916020019190600101906200025d565b50620002889291506200028c565b5090565b620002a991905b8082111562000288576000815560010162000293565b90565b61118d80620002bc6000396000f30060806040526004361061011c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461012157806306fdde0314610157578063079cec9f146101e1578063081812fc1461020f578063095ea7b3146102435780630e723a371461026957806318160ddd1461029757806319fa8f50146102be57806323b872dd146102f05780632f745c591461031a57806342842e0e1461033e5780634f558e79146103685780634f6ccce7146103805780636352211e1461039857806370a08231146103b057806395d89b41146103d1578063a22cb465146103e6578063b88d4fde1461040c578063c87b56dd1461047b578063e78cea9214610493578063e985e9c5146104a8575b600080fd5b34801561012d57600080fd5b50610143600160e060020a0319600435166104cf565b604080519115158252519081900360200190f35b34801561016357600080fd5b5061016c6104ee565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a657818101518382015260200161018e565b50505050905090810190601f1680156101d35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ed57600080fd5b50610143600160a060020a036004351667ffffffffffffffff60243516610585565b34801561021b57600080fd5b5061022760043561061e565b60408051600160a060020a039092168252519081900360200190f35b34801561024f57600080fd5b50610267600160a060020a0360043516602435610639565b005b34801561027557600080fd5b50610143600160a060020a036004351667ffffffffffffffff602435166106ef565b3480156102a357600080fd5b506102ac6107a5565b60408051918252519081900360200190f35b3480156102ca57600080fd5b506102d36107ab565b60408051600160e060020a03199092168252519081900360200190f35b3480156102fc57600080fd5b50610267600160a060020a03600435811690602435166044356107cf565b34801561032657600080fd5b506102ac600160a060020a036004351660243561088b565b34801561034a57600080fd5b50610267600160a060020a03600435811690602435166044356108d8565b34801561037457600080fd5b506101436004356108f9565b34801561038c57600080fd5b506102ac600435610916565b3480156103a457600080fd5b5061022760043561094b565b3480156103bc57600080fd5b506102ac600160a060020a0360043516610975565b3480156103dd57600080fd5b5061016c6109a8565b3480156103f257600080fd5b50610267600160a060020a03600435166024351515610a09565b34801561041857600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261026794600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750610a8d9650505050505050565b34801561048757600080fd5b5061016c600435610ab5565b34801561049f57600080fd5b50610227610b6a565b3480156104b457600080fd5b50610143600160a060020a0360043581169060243516610b79565b600160e060020a03191660009081526020819052604090205460ff1690565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561057a5780601f1061054f5761010080835404028352916020019161057a565b820191906000526020600020905b81548152906001019060200180831161055d57829003601f168201915b505050505090505b90565b600c54600090600160a060020a0316331461060157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4f6e6c79206272696467652063616e206d696e74206465706f73697473000000604482015290519081900360640190fd5b610615838367ffffffffffffffff16610ba7565b50600192915050565b600090815260026020526040902054600160a060020a031690565b60006106448261094b565b9050600160a060020a03838116908216141561065f57600080fd5b33600160a060020a038216148061067b575061067b8133610b79565b151561068657600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b600c54600090600160a060020a0316331461079157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4f6e6c79206f7261636c652063616e206275726e207769746864726177616c7360448201527f2e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610615838367ffffffffffffffff16610bf0565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6107da838383610cf7565b600c54604080517fb3d3a508000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152858116602483015267ffffffffffffffff851660448301529151919092169163b3d3a5089160648083019260209291908290030181600087803b15801561085a57600080fd5b505af115801561086e573d6000803e3d6000fd5b505050506040513d602081101561088457600080fd5b5050505050565b600061089683610975565b82106108a157600080fd5b600160a060020a03831660009081526007602052604090208054839081106108c557fe5b9060005260206000200154905092915050565b6108f48383836020604051908101604052806000815250610a8d565b505050565b600090815260016020526040902054600160a060020a0316151590565b60006109206107a5565b821061092b57600080fd5b600980548390811061093957fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561096f57600080fd5b92915050565b6000600160a060020a038216151561098c57600080fd5b50600160a060020a031660009081526003602052604090205490565b60068054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561057a5780601f1061054f5761010080835404028352916020019161057a565b600160a060020a038216331415610a1f57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610a988484846107cf565b610aa484848484610d9a565b1515610aaf57600080fd5b50505050565b6060610ac0826108f9565b1515610acb57600080fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610b5e5780601f10610b3357610100808354040283529160200191610b5e565b820191906000526020600020905b815481529060010190602001808311610b4157829003601f168201915b50505050509050919050565b600c54600160a060020a031681565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000610bb38383610f07565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b6000806000610bff8585610f97565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350610c3a90600163ffffffff61102d16565b600160a060020a038616600090815260076020526040902080549193509083908110610c6257fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610ca257fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490610cd9906000198301611124565b50600093845260086020526040808520859055908452909220555050565b610d01338261103f565b1515610d0c57600080fd5b600160a060020a0383161515610d2157600080fd5b600160a060020a0382161515610d3657600080fd5b610d40838261109e565b610d4a8382610bf0565b610d548282610ba7565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600080610daf85600160a060020a031661110f565b1515610dbe5760019150610efe565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610e51578181015183820152602001610e39565b50505050905090810190601f168015610e7e5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610ea057600080fd5b505af1158015610eb4573d6000803e3d6000fd5b505050506040513d6020811015610eca57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b600081815260016020526040902054600160a060020a031615610f2957600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610f7791611117565b600160a060020a0390921660009081526003602052604090209190915550565b81600160a060020a0316610faa8261094b565b600160a060020a031614610fbd57600080fd5b600160a060020a038216600090815260036020526040902054610fe790600163ffffffff61102d16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008282111561103957fe5b50900390565b60008061104b8361094b565b905080600160a060020a031684600160a060020a03161480611086575083600160a060020a031661107b8461061e565b600160a060020a0316145b8061109657506110968185610b79565b949350505050565b81600160a060020a03166110b18261094b565b600160a060020a0316146110c457600080fd5b600081815260026020526040902054600160a060020a03161561110b576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000903b1190565b8181018281101561096f57fe5b8154818355818111156108f4576000838152602090206108f491810190830161058291905b8082111561115d5760008155600101611149565b50905600a165627a7a7230582018c5cba0295671b144a2f9c5ac575bb1e4cfa1f54deaa727b32445bdb525de250029`

// DeployChildBridge deploys a new Ethereum contract, binding an instance of ChildBridge to it.
func DeployChildBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *ChildBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChildBridgeBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChildBridge{ChildBridgeCaller: ChildBridgeCaller{contract: contract}, ChildBridgeTransactor: ChildBridgeTransactor{contract: contract}, ChildBridgeFilterer: ChildBridgeFilterer{contract: contract}}, nil
}

// ChildBridge is an auto generated Go binding around an Ethereum contract.
type ChildBridge struct {
	ChildBridgeCaller     // Read-only binding to the contract
	ChildBridgeTransactor // Write-only binding to the contract
	ChildBridgeFilterer   // Log filterer for contract events
}

// ChildBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildBridgeSession struct {
	Contract     *ChildBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildBridgeCallerSession struct {
	Contract *ChildBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChildBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildBridgeTransactorSession struct {
	Contract     *ChildBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChildBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildBridgeRaw struct {
	Contract *ChildBridge // Generic contract binding to access the raw methods on
}

// ChildBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildBridgeCallerRaw struct {
	Contract *ChildBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ChildBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildBridgeTransactorRaw struct {
	Contract *ChildBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildBridge creates a new instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridge(address common.Address, backend bind.ContractBackend) (*ChildBridge, error) {
	contract, err := bindChildBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildBridge{ChildBridgeCaller: ChildBridgeCaller{contract: contract}, ChildBridgeTransactor: ChildBridgeTransactor{contract: contract}, ChildBridgeFilterer: ChildBridgeFilterer{contract: contract}}, nil
}

// NewChildBridgeCaller creates a new read-only instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeCaller(address common.Address, caller bind.ContractCaller) (*ChildBridgeCaller, error) {
	contract, err := bindChildBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeCaller{contract: contract}, nil
}

// NewChildBridgeTransactor creates a new write-only instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildBridgeTransactor, error) {
	contract, err := bindChildBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeTransactor{contract: contract}, nil
}

// NewChildBridgeFilterer creates a new log filterer instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildBridgeFilterer, error) {
	contract, err := bindChildBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeFilterer{contract: contract}, nil
}

// bindChildBridge binds a generic wrapper to an already deployed contract.
func bindChildBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildBridge *ChildBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildBridge.Contract.ChildBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildBridge *ChildBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.Contract.ChildBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildBridge *ChildBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildBridge.Contract.ChildBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildBridge *ChildBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildBridge *ChildBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildBridge *ChildBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildBridge.Contract.contract.Transact(opts, method, params...)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeCaller) GetPendingTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "getPendingTransactionCount")
	return *ret0, err
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildBridge.Contract.GetPendingTransactionCount(&_ChildBridge.CallOpts)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeCallerSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildBridge.Contract.GetPendingTransactionCount(&_ChildBridge.CallOpts)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeCaller) LastBlockOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "lastBlockOf", arg0)
	return *ret0, err
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildBridge.Contract.LastBlockOf(&_ChildBridge.CallOpts, arg0)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeCallerSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildBridge.Contract.LastBlockOf(&_ChildBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeSession) Owner() (common.Address, error) {
	return _ChildBridge.Contract.Owner(&_ChildBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeCallerSession) Owner() (common.Address, error) {
	return _ChildBridge.Contract.Owner(&_ChildBridge.CallOpts)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeCaller) PendingTransactions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "pendingTransactions", arg0)
	return *ret0, err
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildBridge.Contract.PendingTransactions(&_ChildBridge.CallOpts, arg0)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeCallerSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildBridge.Contract.PendingTransactions(&_ChildBridge.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChildBridge *ChildBridgeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChildBridge *ChildBridgeSession) Token() (common.Address, error) {
	return _ChildBridge.Contract.Token(&_ChildBridge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChildBridge *ChildBridgeCallerSession) Token() (common.Address, error) {
	return _ChildBridge.Contract.Token(&_ChildBridge.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	ret := new(struct {
		Status    uint8
		SlotId    uint64
		PrevBlock *big.Int
		NewOwner  common.Address
		Owner     common.Address
		Signature []byte
	})
	out := ret
	err := _ChildBridge.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildBridge.Contract.Transactions(&_ChildBridge.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeCallerSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildBridge.Contract.Transactions(&_ChildBridge.CallOpts, arg0)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeTransactor) CreateTransaction(opts *bind.TransactOpts, from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "createTransaction", from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.Contract.CreateTransaction(&_ChildBridge.TransactOpts, from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeTransactorSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.Contract.CreateTransaction(&_ChildBridge.TransactOpts, from, to, slotId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildBridge.Contract.RenounceOwnership(&_ChildBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildBridge.Contract.RenounceOwnership(&_ChildBridge.TransactOpts)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeTransactor) SaveWitness(opts *bind.TransactOpts, txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "saveWitness", txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.Contract.SaveWitness(&_ChildBridge.TransactOpts, txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.Contract.SaveWitness(&_ChildBridge.TransactOpts, txnHash, signature)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactor) SubmitDeposit(opts *bind.TransactOpts, depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "submitDeposit", depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitDeposit(&_ChildBridge.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitDeposit(&_ChildBridge.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeTransactor) SubmitNewBlock(opts *bind.TransactOpts, newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "submitNewBlock", newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitNewBlock(&_ChildBridge.TransactOpts, newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitNewBlock(&_ChildBridge.TransactOpts, newBlockNumber)
}

// SubmitWithdraw is a paid mutator transaction binding the contract method 0xe023e759.
//
// Solidity: function submitWithdraw(exitor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactor) SubmitWithdraw(opts *bind.TransactOpts, exitor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "submitWithdraw", exitor, parentToken, slotId, amount, which)
}

// SubmitWithdraw is a paid mutator transaction binding the contract method 0xe023e759.
//
// Solidity: function submitWithdraw(exitor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeSession) SubmitWithdraw(exitor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitWithdraw(&_ChildBridge.TransactOpts, exitor, parentToken, slotId, amount, which)
}

// SubmitWithdraw is a paid mutator transaction binding the contract method 0xe023e759.
//
// Solidity: function submitWithdraw(exitor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SubmitWithdraw(exitor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitWithdraw(&_ChildBridge.TransactOpts, exitor, parentToken, slotId, amount, which)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.Contract.TransferOwnership(&_ChildBridge.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.Contract.TransferOwnership(&_ChildBridge.TransactOpts, _newOwner)
}

// ChildBridgeConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the ChildBridge contract.
type ChildBridgeConfirmTransactionIterator struct {
	Event *ChildBridgeConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *ChildBridgeConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeConfirmTransaction)
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
		it.Event = new(ChildBridgeConfirmTransaction)
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
func (it *ChildBridgeConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeConfirmTransaction represents a ConfirmTransaction event raised by the ChildBridge contract.
type ChildBridgeConfirmTransaction struct {
	TxnHash [32]byte
	Owner   common.Address
	SlotId  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: event ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildBridge *ChildBridgeFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, txnHash [][32]byte, owner []common.Address, slotId []uint64) (*ChildBridgeConfirmTransactionIterator, error) {

	var txnHashRule []interface{}
	for _, txnHashItem := range txnHash {
		txnHashRule = append(txnHashRule, txnHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeConfirmTransactionIterator{contract: _ChildBridge.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: event ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildBridge *ChildBridgeFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *ChildBridgeConfirmTransaction, txnHash [][32]byte, owner []common.Address, slotId []uint64) (event.Subscription, error) {

	var txnHashRule []interface{}
	for _, txnHashItem := range txnHash {
		txnHashRule = append(txnHashRule, txnHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeConfirmTransaction)
				if err := _ChildBridge.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ChildBridgeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ChildBridge contract.
type ChildBridgeOwnershipRenouncedIterator struct {
	Event *ChildBridgeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ChildBridgeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeOwnershipRenounced)
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
		it.Event = new(ChildBridgeOwnershipRenounced)
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
func (it *ChildBridgeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeOwnershipRenounced represents a OwnershipRenounced event raised by the ChildBridge contract.
type ChildBridgeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ChildBridgeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeOwnershipRenouncedIterator{contract: _ChildBridge.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ChildBridgeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeOwnershipRenounced)
				if err := _ChildBridge.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ChildBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChildBridge contract.
type ChildBridgeOwnershipTransferredIterator struct {
	Event *ChildBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChildBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeOwnershipTransferred)
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
		it.Event = new(ChildBridgeOwnershipTransferred)
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
func (it *ChildBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ChildBridge contract.
type ChildBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChildBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeOwnershipTransferredIterator{contract: _ChildBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChildBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeOwnershipTransferred)
				if err := _ChildBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
