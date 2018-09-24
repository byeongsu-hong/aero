pragma solidity ^0.4.24;

import "../utils/ERC20.sol";
import "./ChildChain.sol";

/**
 * @title PeggedERC20
 * @dev Base ERC20 token contract for tokens on the child chain, which is pegged
 *   to original ERC20 on the parent chain.
 */
contract PeggedERC20 is ERC20 {

    // an address of the gateway contract on child chain (ChildChain.sol) 
    address public gateway;

    constructor(address gatewayAddress) {
        gateway = gatewayAddress;
    }

    /**
    * @dev Function to mint tokens for the deposit from the parent chain.
    */
    function mint(address to, uint256 amount) public returns (bool) {
        // PeggedERC20 should not create any tokens on the 
        require(msg.sender == gateway, "Only gateway can mint deposits");

        _mint(to, amount);
        return true;
    }

    /**
     * @dev Function to burn tokens for the withdrawal from this chain.
     */
    function burnFrom(address from, uint256 value) public returns (bool) {
        ChildChain gatewayContract = ChildChain(gateway);
        require(msg.sender == gatewayContract.oracle(), "Only oracle can burn.");

        _burn(from, value);
        return true;
    }
}