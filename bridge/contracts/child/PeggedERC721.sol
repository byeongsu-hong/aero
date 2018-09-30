pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721Token.sol";
import "./ChildBridge.sol";

/**
 * @title PeggedERC721
 * @dev Base ERC20 token contract for tokens on the child chain, which is pegged
 *   to original ERC20 on the parent chain.
 */
contract PeggedERC721 is ERC721Token {

    // an address of the bridge contract on child chain (Pegger.sol)
    ChildBridge public bridge;

    constructor(string _name, string _symbol, ChildBridge _bridge) ERC721Token(_name, _symbol) public {
        bridge = _bridge;
        // TODO: is registered?
    }

    /**
     * @dev Transfers the ownership of a given token ID to another address,
     * and report the transaction to the parent chain.
     * @param _from current owner of the token
     * @param _to address to receive the ownership of the given token ID
     * @param _tokenId uint256 ID of the token to be transferred
    */
    function transferFrom(address _from, address _to, uint256 _tokenId) public {
        super.transferFrom(_from, _to, _tokenId);

        // report transaction to the parent by creating Plasma TX.
        bridge.createTransaction(_from, _to, uint64(_tokenId));
    }

    /**
     * @dev Function to mint tokens for the deposit from the parent chain.
     */
    function addDepositTo(address to, uint64 slotId) public returns (bool) {
        require(msg.sender == address(bridge), "Only bridge can mint deposits");

        addTokenTo(to, uint256(slotId));
        return true;
    }

    /**
     * @dev Function to burn tokens for the withdrawal from this chain.
     */
    function withdrawFrom(address from, uint64 slotId) public returns (bool) {
        require(msg.sender == address(bridge), "Only oracle can burn withdrawals.");

        removeTokenFrom(from, uint256(slotId));
        return true;
    }
}
