pragma solidity ^0.4.24;

import "./ERC20.sol";

/**
 * @title PeggedERC20
 * @dev Base ERC20 token contract for tokens on the child chain, which is pegged
 *   to original ERC20 on the parent chain.
 */
contract PeggedERC20 is ERC20 {

    // an address of the gateway contract on child chain (ChildChain.sol) 
    address private gateway;

    /**
     * @return An address of the gateway
     */
    function getGateway() public view returns (address) {
        return gateway;
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
     * @dev Burns a specific amount of tokens.
     */
    function burn(uint256 value) public {
        _burn(msg.sender, value);
    }
}