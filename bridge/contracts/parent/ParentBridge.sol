pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/MerkleProof.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";
import "../OperatorRegistry.sol";
import "../Transaction.sol";
import "./SparseMerkleTree.sol";

/**
 * @title ParentChain
 * @dev A gateway contract on the parent chain, bridging parent chain with the child chain.
 */
contract ParentBridge {
    using SafeERC20 for ERC20;
    using SafeMath for uint256;
    using Transaction for bytes;
    using ECRecovery for bytes32;

    event Deposit(
        uint64 indexed slotId,
        address indexed depositor,
        uint256 indexed depositBlockNumber,
        address token,
        uint256 amount
    );
    event BlockSubmit(bytes32 root, uint256 timestamp);
    event ExitStarted(uint64 indexed slotId, address indexed owner);
    event ExitRejected(uint64 indexed slotId, address indexed claimer);
    event ExitFinalized(
        uint64 indexed slotId,
        address indexed owner,
        uint256 indexed amount
    );

    enum Type {
        ERC20,
        ERC721
    }

    enum State {
        DEPOSITED,
        EXITING,
        EXITED
    }

    struct Exit {
        address owner;
        uint256 exitBlock;
        address prevOwner; // previous owner of coin
        uint256 prevBlock;
        uint256 createdAt;
    }

    struct Coin {
        Type typ;
        address owner;
        address token;
        uint256 uid;
        uint256 depositBlock;
        uint256 value;
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
    uint256 public constant CHALLENGE_PERIOD = 200; // no timestamp. no.

    OperatorRegistry operatorRegistry;
    SparseMerkleTree merkleTree;

    uint256 public currentChildBlock;
    uint256 public currentDepositBlock;
    uint64 public coinCount = 0;

    mapping (uint256 => ChildBlock) public childBlocks;
    mapping (uint64 => Coin) coins;

    modifier onlyOperator() { require(operatorRegistry.isOperator(msg.sender)); _;}
    modifier onlyValidator() { require(operatorRegistry.isValidator(msg.sender)); _;}

    constructor() public {
        operatorRegistry = new OperatorRegistry();
        merkleTree = new SparseMerkleTree();
    }

    // User-side actions
    // TODO: check approved balance
    /**
     * @dev Deposit non-fungible token (ERC721) into the child chain.
     */
    function depositNonFungible(
        ERC721 token,
        address from,
        uint256 tokenId
    ) public {
        token.safeTransferFrom(from, address(this), tokenId);
        deposit(Type.ERC721, token, from, tokenId, 1);
    }

    /**
     * @dev Deposit fungible token (ERC20) into the child chain.
     */
    function depositFungible(
        ERC20 token,
        address from,
        uint256 value
    ) public {
        token.safeTransferFrom(from, address(this), value);
        deposit(Type.ERC20, token, from, 0, value);
    }

    function deposit(
        Type typ,
        address token,
        address from,
        uint256 tokenId,
        uint256 amount
    ) private {
        require(
            currentDepositBlock < CHILD_BLOCK_INTERVAL,
            "Only allow up to 1000 deposits per child block.");

        bytes32 depositId = keccak256(abi.encodePacked(from, address(token), coinCount));
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
        coin.typ = typ;
        coin.owner = from;
        coin.token = token;
        coin.uid = tokenId;
        coin.depositBlock = depositBlock;
        coin.value = amount;
        coin.state = State.DEPOSITED;

        coinCount += 1;
        emit Deposit(
            slotId,
            coin.owner,
            coin.depositBlock,
            coin.token,
            coin.value
        );
    }

    function exit(
        bytes exitTxData, bytes exitProof, uint256 exitBlock,
        bytes prevTxData, bytes prevProof, uint256 prevBlock,
        bytes sign
    ) public {
        if(exitBlock % CHILD_BLOCK_INTERVAL != 0) {
            depositExit(
                exitBlock,
                childBlocks[exitBlock].root,
                exitTxData,
                exitProof
            );
        } else {
            defaultExit(
                exitBlock,
                exitTxData, childBlocks[exitBlock].root, exitProof,
                prevTxData, childBlocks[prevBlock].root, prevProof,
                sign
            );
        }
    }

    function depositExit(
        uint256 blockNumber,
        bytes32 root,
        bytes exitTxData,
        bytes exitProof
    ) private {
        Transaction.Tx memory exitTx = exitTxData.getTx();

        require(exitTx.newOwner == msg.sender, "only owner can exit");
        require(exitTx.prevBlock == 0, "prev block should be zero in deposit tx");
        require(inclusionCheck(exitTx, root, exitProof), "membership check failed");

        submitExit(
            exitTx.slotId,
            exitTx.newOwner, blockNumber,
            address(0), 0);
    }

    function defaultExit(
        uint256 blockNumber,
        bytes exitTxData, bytes32 exitRoot, bytes exitProof,
        bytes prevTxData, bytes32 prevRoot, bytes prevProof,
        bytes sign
    ) private {
        Transaction.Tx memory exitTx = exitTxData.getTx();
        Transaction.Tx memory prevTx = prevTxData.getTx();

        require(exitTx.newOwner == msg.sender, "invalid owner");
        require(prevTx.slotId == exitTx.slotId, "invalid slotId");
        require(prevTx.newOwner == exitTx.hash.recover(sign), "invalid sig");
        require(prevRoot == childBlocks[exitTx.prevBlock].root, "invalid tx data");
        require(inclusionCheck(exitTx, exitRoot, exitProof), "inclusion check failed");
        require(inclusionCheck(prevTx, prevRoot, prevProof), "inclusion check failed");

        submitExit(
            exitTx.slotId,
            exitTx.newOwner, blockNumber,
            prevTx.newOwner, exitTx.prevBlock);
    }

    function submitExit(
        uint64 slotId,
        address exitOwner, uint256 exitBlock,
        address prevOwner, uint256 prevBlock
    ) private {
        Coin storage coin = coins[slotId];
        coin.exit = Exit({
            owner: exitOwner,
            exitBlock: exitBlock,
            prevOwner: prevOwner,
            prevBlock: prevBlock,
            createdAt: block.number
        });
        coin.state = State.EXITING;
        emit ExitStarted(slotId, coin.exit.owner);
    }

    // TODO list
    // before  : x
    // between : v
    // after   : v
    function challenge(
        uint64 slotId,
        uint256 blockNumber,
        bytes claimTxData,
        bytes proof,
        bytes sign
    ) public {
        require(
            coins[slotId].state == State.EXITING,
            "only exiting coin can challenge");

        Transaction.Tx memory Tx = claimTxData.getTx();
        Coin storage coin = coins[slotId];
        require(Tx.slotId == slotId, "invalid slot");
        require(Tx.hash.recover(sign) == coin.exit.prevOwner, "invalid sig");

        if(getChallengeType(slotId, blockNumber)) {
            // between

        } else {
            // after
            require(Tx.prevBlock == coin.exit.exitBlock, "invalid reference");
        }

        require(
            inclusionCheck(Tx, childBlocks[blockNumber].root, proof),
            "inclusion check failed");

        // reject exit
        coin.state = State.DEPOSITED;
        delete coin.exit;
        emit ExitRejected(slotId, msg.sender);
    }

    function getChallengeType(uint64 slotId, uint256 blockNumber)
        private
        view
        returns (bool)
    {
        return (
            coins[slotId].exit.exitBlock > blockNumber &&
            coins[slotId].exit.prevBlock < blockNumber
        );
    }

    function inclusionCheck(
        Transaction.Tx memory Tx,
        bytes32 root,
        bytes proof
    ) private view returns (bool) {
        return merkleTree.checkMembership(Tx.hash, root, Tx.slotId, proof);
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

    function isDepositBlock(uint256 _blockNumber) internal pure returns (bool) {
        return (_blockNumber % CHILD_BLOCK_INTERVAL == 0);
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