pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./PeggedERC20.sol";
import "./PeggedERC721.sol";

/**
 * @dev A ERC721 Pegger,
 * which is a proof-of-concept implementation of the higher-level Plasma Cash.
 */
contract Pegger is Ownable {
    using SafeMath for uint256;
    using ECRecovery for bytes32;

    // type of transaction.
    enum Mode {
        ERC20,
        ERC721
    }

    enum Status {
        PENDING,
        CONFIRMED
    }

    struct Txn {
        Status status;

        // items that will be RLP-encoded by the operator
        uint64 slotId;
        uint256 prevBlock;
        address newOwner;

        // metadata
        address owner;
        bytes signature;
    }

    event ConfirmTransaction(
        bytes32 indexed txnHash,
        address indexed owner,
        uint64 indexed slotId
    );

    PeggedERC721 token;

    mapping (bytes32 => Txn) public transactions;
    mapping (uint256 => uint256) public lastBlockOf;
    bytes32[] public pendingTransactions;

    constructor(PeggedERC721 _token) public {
        token = _token;
    }

    function createTransaction(address from, address to, uint64 slotId) public returns (bytes32 txnHash) {
        // TODO: max transaction limit
        require(msg.sender == address(token), "Direct call is not allowed.");

        // build transaction data
        Txn memory txn;
        txn.slotId = slotId;
        txn.prevBlock = lastBlockOf[slotId];
        txn.newOwner = to;
        txn.owner = from;

        // calculate hash by constructing an naive RLP encoding of the transaction.
        bytes memory rlp = abi.encodePacked(
            bytes2(0xf857),
            bytes1(0x88), txn.slotId,
            bytes1(0xa0), txn.prevBlock,
            bytes1(0x94), txn.newOwner
        );
        txnHash = keccak256(rlp);

        // save the transaction.
        // NOTE THAT txn.signature could be further provided by the client.
        transactions[txnHash] = txn;
        pendingTransactions.push(txnHash);
    }

    function saveWitness(bytes32 txnHash, bytes signature) public {
        require(transactions[txnHash].slotId != 0, "No Transaction ID found.");

        // TODO: do we really need signature verification here?
        Txn storage txn = transactions[txnHash];
        require(txnHash.recover(signature) == txn.owner, "Signature mismatch.");

        txn.signature = signature;
    }

    function submitNewBlock(uint256 newBlockNumber) public onlyOwner {
        for (uint256 i = 0; i < pendingTransactions.length; i++) {
            bytes32 txnHash = pendingTransactions[i];
            Txn storage txn = transactions[txnHash];

            txn.status = Status.CONFIRMED;
            lastBlockOf[txn.slotId] = newBlockNumber;
            emit ConfirmTransaction(txnHash, txn.owner, txn.slotId);
        }
        delete pendingTransactions;
    }

    /**
     * @dev Submit deposit event of ERC20 / ERC721 token from the parent chain,
     * and creates a deposit block.
     */
    function submitDeposit(address depositor, address parentToken, uint64 slotId, uint256 amount, Mode which) public onlyOwner {
        require(
            parentToken == address(token),
            "Unregistered token.");

        // mint deposits to the depositor.
        if (which == Mode.ERC20) {
            PeggedERC20 token20 = PeggedERC20(parentToken);
            token20.mint(depositor, amount);
        } else {
            PeggedERC721 token721 = PeggedERC721(parentToken);
            token721.mint(depositor, slotId);
        }
    }

    function getPendingTransactionCount() public view returns (uint256) {
        return pendingTransactions.length;
    }
}
