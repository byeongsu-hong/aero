pragma solidity ^0.4.24;

import "./utils/RLP.sol";

/**
 * @title Transaction
 * @dev Child chain transaction, represented in RLP.
 * From Loom Network Plasma Cash implementation, with a modification for Plasma Debit support.
 */
library Transaction {
    using RLP for bytes;
    using RLP for RLP.RLPItem;

    struct Tx {
        uint64 slotId;
        uint256 prevBlock;
        address newOwner;
        bytes32 hash;
    }

    function getTx(bytes memory txBytes) internal pure returns (Tx memory) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(3);
        Tx memory transaction;

        transaction.slotId = uint64(rlpTx[0].toUint());
        transaction.prevBlock = rlpTx[1].toUint();
        transaction.newOwner = rlpTx[2].toAddress();
        if (transaction.prevBlock == 0) {
            // deposit transaction
            transaction.hash = keccak256(abi.encodePacked(transaction.slotId));
        } else {
            transaction.hash = keccak256(txBytes);
        }
        return transaction;
    }

    function getHash(bytes memory txBytes) internal pure returns (bytes32 hash) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(3);
        uint64 slot = uint64(rlpTx[0].toUint());
        uint256 prevBlock = uint256(rlpTx[1].toUint());

        if (prevBlock == 0) {
            // deposit transaction
            hash = keccak256(abi.encodePacked(slot));
        } else {
            hash = keccak256(txBytes);
        }
    }

    function getOwner(bytes memory txBytes) internal pure returns (address newOwner) {
        RLP.RLPItem[] memory rlpTx = txBytes.toRLPItem().toList(3);
        newOwner = rlpTx[2].toAddress();
    }
}