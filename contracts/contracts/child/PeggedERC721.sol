pragma solidity ^0.4.24;

import "../utils/ERC721.sol";
import "./ChildChain.sol";

/**
 * @title PeggedERC721
 * @dev Base ERC20 token contract for tokens on the child chain, which is pegged
 *   to original ERC20 on the parent chain.
 */
contract PeggedERC721 is ERC721 {

    // an address of the gateway contract on child chain (ChildChain.sol) 
    address public gateway;

    constructor(address gatewayAddress) {
//        totalSupply = 0;
        gateway = gatewayAddress;
    }

    /**
    * @dev Function to mint tokens for the deposit from the parent chain.
    */
    function mint(address to, uint256 tokenId) public returns (bool) {
        // PeggedERC20 should not create any tokens on the 
        require(msg.sender == gateway, "Only gateway can mint deposits");

        _mint(to, tokenId);
        return true;
    }

    /**
     * @dev Function to burn tokens for the withdrawal from this chain.
     */
    function burnFrom(address from, uint256 tokenId) public returns (bool) {
        ChildChain gatewayContract = ChildChain(gateway);
        require(msg.sender == gatewayContract.oracle(), "Only oracle can burn withdrawals.");

        _burn(msg.sender, tokenId);
    }
}