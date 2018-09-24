pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "../OperatorRegistry.sol";
import "../Transaction.sol";
import "./PeggedERC20.sol";
import "./PeggedERC721.sol";

contract ChildChain {
    using SafeMath for uint256;

    // type of transaction.
    enum Mode {
        ERC20,
        ERC721
    }

    // child containing plasma blocks (RLP-encoded light block)
    mapping (uint256 => mapping (bytes32 => bytes)) blocks;

    mapping (address => address) public parentToChildContracts;
    address public oracle;

    // the identical contract is also used on the child chain.
    OperatorRegistry operatorRegistry;

    modifier onlyOperator() { require(operatorRegistry.isOperator(msg.sender)); _; }
    modifier onlyOracle() { require(msg.sender == oracle); _; }

    constructor(address _oracle) public {
        oracle = _oracle;
        operatorRegistry = new OperatorRegistry();
    }

    /**
     * @dev Register ERC20 token contract which is pegged to the parent chain.
     * @param token ERC20 contract instance inheriting PeggedERC20
     * @param addressOnParent An address of the parent token corresponding to the given token.
     */
    function registerERC20(
        PeggedERC20 token,
        address addressOnParent)
        public
        onlyOperator
    {
        require(parentToChildContracts[addressOnParent] == 0x0, "Already registered.");
        require(token.gateway() == address(this), "Given token uses different gateway.");

        parentToChildContracts[addressOnParent] = address(token);
    }

    /**
     * @dev Register ERC721 token contract which is pegged to the parent chain.
     * @param token ERC721 contract instance inheriting PeggedERC721
     * @param addressOnParent An address of the parent token corresponding to the given token.
     */
    function registerERC721(
        PeggedERC721 token,
        address addressOnParent)
        public
        onlyOperator
    {
        require(parentToChildContracts[addressOnParent] == 0x0, "Already registered.");
        require(token.gateway() == address(this), "Given token uses different gateway.");

        parentToChildContracts[addressOnParent] = address(token);
    }

    /**
     * @dev Submit deposit event of ERC20 token from the parent chain,
     * and creates a deposit block.
     */
    function submitDeposit(
        address depositor,
        address parentToken,
        uint256 amount,
        Mode which)
        public
        onlyOracle
    {
        require(
            parentToChildContracts[parentToken] != 0x0,
            "Unregistered token.");
        address tokenAddress = parentToChildContracts[parentToken];

        // mint deposits to the depositor.
        if (which == Mode.ERC20) {
           PeggedERC20 token = PeggedERC20(tokenAddress);
           token.mint(depositor, amount);
        } else {
            uint256 tokenId = amount;
            PeggedERC721 token = PeggedERC721(tokenAddress);
            token.mint(depositor, tokenId);
        }
    }

    /**
     * @dev Submit withdrawal event of ERC20 token from the parent chain,
     * which updates the state of child chain by burning the withdrawn coin.
     */
    function submitWithdrawal(
        address withdrawer,
        address parentToken,
        uint256 amount,
        Mode which)
        public
        onlyOracle
    {
        require(
            parentToChildContracts[parentToken] != 0x0,
            "Unregistered token.");
        address tokenAddress = parentToChildContracts[parentToken];

        // update the child state by burning.
        if (which == Mode.ERC20) {
            PeggedERC20 token = PeggedERC20(tokenAddress);
            token.burnFrom(withdrawer, amount);

        } else if (which == Mode.ERC721) {
            uint256 tokenId = amount;
            PeggedERC721 token = PeggedERC721(tokenAddress);
            token.burnFrom(withdrawer, tokenId);
        }
    }
}
