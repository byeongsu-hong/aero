pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./OperatorRegistry.sol";
import "./ERC20.sol";

contract ChildChain {
    using SafeMath for uint256;

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
        require(token.getGateway() == address(this), "Given token uses different gateway.");

        parentToChildContracts[addressOnParent] = address(token);
    }

    /**
     * @dev Submit deposit event of ERC20 token from the parent chain.
     */
    function submitDeposit(
        address depositor,
        address parentToken,
        uint256 value)
        public
        onlyOracle
    {
        require(
            parentToChildContracts[parentToken] != 0x0, 
            "Only allow up to 1000 deposits per child block.");
        
        // mint deposits to the depositor.
        PeggedERC20 token = PeggedERC20(parentToChildContracts[parentToken]);
        token.mint(address(depositor), value);
    }
}
