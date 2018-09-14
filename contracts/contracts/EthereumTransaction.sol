pragma solidity ^0.4.24;

import "./RLP.sol";

/**
 * @title EthereumTransaction
 */
library EthereumTransaction {
    using RLP for bytes;
    using RLP for RLP.RLPItem;

    struct TX {
        uint64 nonce;
        uint256 gasPrice;
        uint256 gasLimit;
        address to;
        uint256 value;
        bytes data;

        // signatures
        uint256 v;
        bytes r;
        bytes s;

        // additional information for Plasma validation
        address newOwner;
        uint256 slot;
        uint256 denomination;
    }

    function getTx(bytes memory txBytes) internal pure returns (TX memory) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(9);
        TX memory transaction;

        transaction.nonce = uint64(rlpTx[0].toUint());
        transaction.gasPrice = rlpTx[1].toUint();
        transaction.gasLimit = rlpTx[2].toUint();
        transaction.to = rlpTx[3].toAddress();
        transaction.value = rlpTx[4].toUint();
        transaction.data = rlpTx[5].toBytes();
        transaction.v = rlpTx[6].toUint();
        transaction.r = rlpTx[7].toUint();
        transaction.s = rlpTx[8].toUint();

        // parse from ERC20 / ERC721 transaction data
        bytes memory txData = transaction.data;
        uint256 param1;
        uint256 param2;
        assembly {
            param1 := div(add(txData, 4), add(txData, 36))
            param2 := div(add(txData, 36), add(txData, 68))
        }

        transaction.newOwner = address(param1);
        transaction.slot = address(param2);

        if (transaction.prevBlock == 0) {
            // deposit transaction
            transaction.hash = keccak256(abi.encodePacked(transaction.slot));
        } else {
            transaction.hash = keccak256(txBytes);
        }
        return transaction;
    }

    function getHash(bytes memory txBytes) internal pure returns (bytes32 hash) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(5);
        uint64 slot = uint64(rlpTx[0].toUint());
        uint256 prevBlock = uint256(rlpTx[1].toUint());

        if (prevBlock == 0) {
            // deposit transaction
            hash = keccak256(abi.encodePacked(slot));
        } else {
            hash = keccak256(txBytes);
        }
    }

    function getOwner(bytes memory txBytes) internal pure returns (address owner) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(5);
        owner = rlpTx[2].toAddress();
    }
}