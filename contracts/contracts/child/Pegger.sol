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

    enum Status {
        PENDING,
        CONFIRMED
    }

    struct Txn {
        Status status;

        // items that will be RLP-encoded by the operator
        uint256 tokenId; // slot
        uint256 prevBlock;
        address newOwner;

        // metadata
        address owner;
        bytes signature;
    }

    event ConfirmTransaction(
        bytes32 indexed txnHash,
        address indexed owner,
        uint256 indexed tokenId
    );

    PeggedERC721 token;

    mapping (bytes32 => Txn) public transactions;
    mapping (uint256 => uint256) public lastBlockOf;
    bytes32[] public pendingTransactions;

    constructor(PeggedERC721 _token) public {
        token = _token;
    }

    function createTransaction(address from, address to, uint256 tokenId) public returns (bytes32 txnHash) {
        require(msg.sender == address(token), "Direct call is not allowed.");

        // build transaction data
        Txn memory txn;
        txn.tokenId = tokenId;
        txn.prevBlock = lastBlockOf[tokenId];
        txn.newOwner = to;
        txn.owner = from;

        // calculate hash by constructing an naive RLP encoding of the transaction.
        bytes memory rlp = abi.encodePacked(
            bytes2(0xf857),
            bytes1(0xa0), txn.tokenId,
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
        require(transactions[txnHash].tokenId != 0, "No Transaction ID found.");

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
            lastBlockOf[txn.tokenId] = newBlockNumber;
            emit ConfirmTransaction(txnHash, txn.owner, txn.tokenId);
        }
        delete pendingTransactions;
    }

    function getPendingTransactionCount() public view returns (uint256) {
        return pendingTransactions.length;
    }
}
