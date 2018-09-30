pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol";

contract ABLToken is MintableToken {
    constructor() public {
        mint(msg.sender, 400000 * 10**18);
    }
}