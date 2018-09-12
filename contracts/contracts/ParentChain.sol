pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/MerkleProof.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "./OperatorRegistry.sol";
import "./ChallengeLib.sol";
import "./Transaction.sol";
import "./ERC20.sol";

/**
 * @title ParentChain
 * @dev A gateway contract on the parent chain, bridging parent chain with the child chain.
 */
contract ParentChain {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    using Transaction for bytes;
    using ChallengeLib for ChallengeLib.Challenge[];

    event Deposit(
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

    enum State {
        DEPOSITED,
        EXITING,
        EXITED
    }

    struct Coin {
        address owner;
        address token;
        uint256 uid;
        uint256 depositBlock;

        // can represent v, a on Plasma Debit
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

    uint256 public currentChildBlock;
    uint256 public currentDepositBlock;
    uint64 public coinCount = 0;

    mapping (uint256 => ChildBlock) public childBlocks;
    mapping (uint64 => Exit) public exits;
    mapping (uint64 => Coin) public coins;
    mapping (uint64 => ChallengeLib.Challenge[]) public challenges;

    modifier onlyOperator() { require(operatorRegistry.isOperator(msg.sender)); _;}
    modifier onlyValidator() { require(operatorRegistry.isValidator(msg.sender)); _;}

    constructor() public {
        operatorRegistry = new OperatorRegistry();
    }

    /**
     * @dev Deposit ERC20 token into the child chain.
     */
    function deposit(
        address from,
        IERC20 token,
        uint256 value)
        public
    {
        require(
            currentDepositBlock < CHILD_BLOCK_INTERVAL, 
            "Only allow up to 1000 deposits per child block.");
        
        token.safeTransferFrom(msg.sender, address(this), value);

        // create a deposit block
        bytes memory data = abi.encodePacked(msg.sender, address(token), value);
        bytes32 root = keccak256(data);
        
        uint256 depositBlock = getNextDepositBlockIndex();
        childBlocks[depositBlock] = ChildBlock({
            root: root,
            timestamp: block.timestamp
        });
        currentDepositBlock = currentDepositBlock.add(1);

        emit Deposit(msg.sender, depositBlock, address(token), value);
    }

    /**
     * @dev Submit a merkle root of child chain blocks, by the operator.
     */
    function submitBlock(bytes32 _root) public onlyOperator {
        childBlocks[currentChildBlock] = ChildBlock({
            root: _root,
            timestamp: block.timestamp
        });

        // update index
        currentChildBlock = currentChildBlock.add(CHILD_BLOCK_INTERVAL);
        currentDepositBlock = 1;
        
        emit BlockSubmit(_root, block.timestamp);
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