pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

/**
 * @title OperatorRegistry
 * @dev Manages a list of operators and validators. 
 */
contract OperatorRegistry is Ownable {

    mapping(address => bool) public operators;
    mapping(address => bool) public validators;

    constructor() Ownable() public {
        operators[msg.sender] = true;
    }

    function toggleOperator(address _address) public {
        require(isOperator(msg.sender) || isOwner(), "Only operator can switch the operator.");
        operators[_address] = !operators[_address];
    }

    function toggleValidator(address _address) public {
        require(isValidator(msg.sender) || isOwner(), "Only operator can switch the operator.");
        validators[_address] = !validators[_address];
    }

    function isOperator(address _address) public view returns (bool) {
        return operators[_address];
    }

    function isValidator(address _address) public view returns (bool) {
        // TODO: ideally, separation of the role is needed.
        return validators[_address] || isOperator(_address);
    }
}
