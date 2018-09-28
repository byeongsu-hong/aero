pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/MerkleProof.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";
import "../OperatorRegistry.sol";
import "../utils/ChallengeLib.sol";
import "../Transaction.sol";
import "./SparseMerkleTree.sol";

/**
 * @title ParentChain
 * @dev A gateway contract on the parent chain, bridging parent chain with the child chain.
 */
contract ParentChain {
    using SafeERC20 for ERC20;
    using SafeMath for uint256;
    using Transaction for bytes;
    using ECRecovery for bytes32;
    using ChallengeLib for ChallengeLib.Challenge[];

    event Deposit(
        uint64 indexed slotId,
        address indexed depositor,
        uint256 indexed depositBlockNumber,
        address token,
        uint256 amount
    );

    event BlockSubmit(
        bytes32 root,
        uint256 timestamp
    );

    struct Exit {
        address prevOwner; // previous owner of coin
        address owner;
        uint256 amount;
        uint256 createdAt;
        uint256 bond;
        uint256 prevBlock;
        uint256 exitBlock;
    }

    enum Type {
        ERC20,
        ERC721
    }

    enum State {
        DEPOSITED,
        EXITING,
        EXITED
    }

    struct Coin {
        Type type;

        address owner;
        address token;
        uint256 uid;
        uint256 depositBlock;

        // can represent v, a on Plasma Debit - unused for now.
        uint256 value;
        uint256 amount;

        // exit state of the coin
        State state;
        Exit exit;
    }

    struct ChildBlock {
        bytes32 root;
        uint256 timestamp;
    }

    // Track owners of txs that are pending a response
    struct Challenge {
        address owner;
        uint256 blockNumber;
    }

    // constraints from Plasma MVP by OmiseGO
    uint256 public constant CHILD_BLOCK_INTERVAL = 1000;

    OperatorRegistry operatorRegistry;
    SparseMerkleTree merkleTree;

    uint256 public currentChildBlock;
    uint256 public currentDepositBlock;
    uint64 public coinCount = 0;

    mapping (uint256 => ChildBlock) public childBlocks;
    mapping (uint64 => Exit) public exits;
    mapping (uint64 => Coin) coins;
    mapping (uint64 => ChallengeLib.Challenge[]) public challenges;

    modifier onlyOperator() { require(operatorRegistry.isOperator(msg.sender)); _;}
    modifier onlyValidator() { require(operatorRegistry.isValidator(msg.sender)); _;}

    constructor() public {
        operatorRegistry = new OperatorRegistry();
        merkleTree = new SparseMerkleTree();
    }

    // User-side actions

    /**
     * @dev Deposit non-fungible token (ERC721) into the child chain.
     */
    function depositNonFungible(ERC721 token, address from, uint256 tokenId) public {
        token.safeTransferFrom(from, address(this), tokenId);
        deposit(Type.ERC721, token, from, tokenId, 1);
    }

    /**
     * @dev Deposit fungible token (ERC20) into the child chain.
     */
    function depositFungible(ERC20 token, address from, uint256 value) {
        token.safeTransferFrom(from, address(this), value);
        deposit(Type.ERC20, token, from, 0, value);
    }

    function deposit(Type type, address token, address from, uint256 tokenId, uint256 amount) private {
        require(
            currentDepositBlock < CHILD_BLOCK_INTERVAL,
            "Only allow up to 1000 deposits per child block.");

        bytes memory depositId = keccak256(abi.encodePacked(msg.sender, address(token), coinCount));
        uint64 slotId = uint64(bytes8(depositId));

        // generate deposit block
        uint256 depositBlock = getNextDepositBlockIndex();
        childBlocks[depositBlock] = ChildBlock({
            root: depositId,
            timestamp: block.timestamp
        });
        currentDepositBlock = currentDepositBlock.add(1);

        // create coin. we use slot like Loom Network implementation
        Coin storage coin = coins[slotId];
        coin.type = type;
        coin.owner = msg.sender;
        coin.token = token;
        coin.uid = tokenId;
        coin.depositBlock = depositBlock;
        coin.value = amount; // this needs to be modified on Plasma Debit
        coin.amount = amount;
        coin.state = State.DEPOSITED;

        coinCount += 1;
        emit Deposit(
            slotId,
            msg.sender,
            depositBlock,
            token,
            amount
        );
    }

    function exit(
        bytes exitTxData, bytes exitProof, uint256 exitBlock,
        bytes prevTxData, bytes prevProof, uint256 prevBlock,
        bytes sign
    ) public {
        if(prevBlock == 0) {
            depositExit(
                exitBlock,
                exitTxData,
                exitProof,
                sign
            );
        } else {
            defaultExit(
                exitBlock,
                exitTxData, exitProof,
                prevTxData, prevProof,
                sign
            );
        }
    }

    function depositExit(
        uint256 blockNumber,
        bytes exitTxData,
        bytes exitProof,
        bytes sign
    ) private {
        Transaction.Tx memory exitTx = exitTxData.getTx();

        require(exitTx.owner == msg.owner, "only owner can exit");
        require(exitTx.prevBlock == 0, "prev block should be zero in deposit tx");

        require(
            merkleTree.checkMembership(
                exitTx.hash,
                childBlocks[blockNumber].root,
                exitTx.slot,
                exitProof
            ),
            "check membership failed"
        );
    }

    function defaultExit(
        uint256 blockNumber,
        bytes exitTxData, bytes exitProof,
        bytes prevTxData, bytes prevProof,
        bytes sign
    ) private {
        Transaction.Tx memory exitTx = exitTxData.getTx();
        Transaction.Tx memory prevTx = prevTxData.getTx();

    }

    function challenge() public {

    }

    function finalize() public {

    }

    // Operator-side actions

    /**
     * @dev Submit a merkle root of child chain blocks, by the operator.
     */
    function submitBlock(bytes32 _root) external onlyOperator returns (uint256) {
        childBlocks[currentChildBlock] = ChildBlock({
            root: _root,
            timestamp: block.timestamp
        });

        // update index
        currentChildBlock = currentChildBlock.add(CHILD_BLOCK_INTERVAL);
        currentDepositBlock = 1;
        
        emit BlockSubmit(_root, block.timestamp);
        return currentChildBlock;
    }

    function isDepositBlock(uint256 _blockNumber) internal view returns (bool) {
        return _blockNumber % CHILD_BLOCK_INTERVAL == 0;
    }

    function getNextDepositBlockIndex()
        public
        view
        returns (uint256)
    {
        return currentChildBlock.sub(CHILD_BLOCK_INTERVAL).add(currentDepositBlock);
    }

    /**
     * @dev Get block of the child chain.
     */
    function getChildBlock(uint256 _blockNumber) 
        public 
        view
        returns (bytes32, uint256) {
        return (childBlocks[_blockNumber].root, childBlocks[_blockNumber].timestamp);
    }
}