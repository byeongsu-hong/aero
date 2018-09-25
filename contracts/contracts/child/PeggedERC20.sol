pragma solidity ^0.4.24;

import "../utils/ERC20.sol";
import "./Pegger.sol";

/**
 * @title PeggedERC20
 * @dev Base ERC20 token contract for tokens on the child chain, which is pegged
 *   to original ERC20 on the parent chain.
 */
contract PeggedERC20 is ERC20 {

    // an address of the bridge contract on child chain (Pegger.sol)
    Pegger public bridge;

    constructor(address _bridge) public {
        bridge = _bridge;
    }

    /**
     * @dev Function to mint tokens for the deposit from the parent chain.
     */
    function addDepositTo(address to, uint256 amount) public returns (bool) {
        require(msg.sender == address(bridge), "Only the bridge can mint deposits");

        _mint(to, amount);
        return true;
    }

    /**
     * @dev Function to burn tokens for the withdrawal from this chain.
     */
    function withdrawFrom(address from, uint256 value) public returns (bool) {
        require(msg.sender == address(bridge), "Only the bridge can burn withdrawn tokens.");

        _burn(from, value);
        return true;
    }
}