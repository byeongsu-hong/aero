pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/MerkleProof.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "./OperatorRegistry.sol";
import "./ERC20.sol";

/**
 * @title ParentChain
 * @dev A gateway contract on the parent chain, bridging parent chain with the child chain.
 */
contract ParentChain {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

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
        address owner;
        address token;
        uint256 amount;
    }

    struct ChildBlock {
        bytes32 root;
        uint256 timestamp;
    }

    // constraints from Plasma MVP by OmiseGO
    uint256 public constant CHILD_BLOCK_INTERVAL = 1000;

    OperatorRegistry operatorRegistry;

    uint256 public currentChildBlock;
    uint256 public currentDepositBlock;
    uint256 public currentFeeExit;

    mapping (uint256 => ChildBlock) public childBlocks;
    mapping (uint256 => Exit) public exits;
    mapping (address => address) public exitsQueues;

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

    /**
     * @dev Get block of the child chain.
     */
    function getChildBlock(uint256 _blockNumber) 
        public 
        view
        returns (bytes32, uint256) {
        return (childBlocks[_blockNumber].root, childBlocks[_blockNumber].timestamp);
    }

    function getNextDepositBlockIndex()
        public
        view
        returns (uint256)
    {
        return currentChildBlock.sub(CHILD_BLOCK_INTERVAL).add(currentDepositBlock);
    }
}