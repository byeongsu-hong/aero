pragma solidity ^0.4.24;

import "./EthereumTransaction.sol";

contract EthereumTransactionValidator {
    using EthereumTransaction for bytes;

    function submit(bytes tx) {

        EthereumTransaction.TX memory txData = tx.getTx();


    }
}
