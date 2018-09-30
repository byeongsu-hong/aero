pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/DetailedERC20.sol";


contract PeggedABLToken is DetailedERC20 {
    constructor() DetailedERC20("ABLTokenTest", "ABLT", 10 ** 18) public {
    }
}
