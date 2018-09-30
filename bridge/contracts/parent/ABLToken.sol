pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol";


contract ABLToken is MintableToken {
    string public name = "ABLTokenTest";
    string public symbol = "ABLT";
    uint8 public decimals = 18;

    constructor() public {
        mint(msg.sender, 400000 * 10 ** 18);
    }
}
