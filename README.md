Aero
=======

Sidechain framework used on Airbloc.  
Currently supports Plasma Cash with non-fungible tokens (ERC721).

**CURRENTLY UNDER DEVELOPMENT - NOT ELIGIBLE FOR PRODUCTION USE!**

## Design Principles

 * [Plasma](https://plasma.io)
 * Separation of Concerns : Avoid modifying the blockchain backend, and aim for being as layer.

## Glossaries

* Parent Chain : A parent chain that provides asset security to child chain. 
* Child Chain : A child chain or sidechain that the asset is pegged into the parent chain. The block headers of child chain is submitted (committed) to the parent chain regularly, and users can exit to the parent chain using the submitted data.
* Plasma Block : A light block containing the transaction data of assets on the child chain.
* Operator : Block producer of the child chain. Operator has a duty to commit the child chain blocks to the parent chain regularly.
* Bridge Contract : Contracts on each chain, bridging the transfer between two chains.
* Gateway Oracle : An (trustful) oracle that reports the other chain's information. (by Operator)
    * Aeroracle : Aero's gateway oracle.
* Enter (Deposit) : Deposit the funds to the child chain.
* Exit (Withdrawal) : Withdraw the funds from the child chain, however depending on root chain.

## Components

 - `aerod` - Aero sidechain daemon, running PoA Ethereum network
 - `operator` - Standalone operator instance
 - `contracts` - Contracts and utilities for both the child and the parent chain.

## Prerequisites

 - Go 1.9+
 - Node 10.0+
 - [Yarn](https://yarnpkg.com/lang/en/) - Package manager for Node
 - [Dep](http://github.com/golang/dep) - Dependency manager for Go
 - [Truffle](https://truffleframework.com/) - Contract deployment framework

## Installation

Executable files will be generated under `build/bin`

```
make
cd contracts
yarn
```

## Test

```
$ ./build/bin/aerod
$ node test/aero-coin-demo.js  # http://localhost:6672/
```

## Current Status

 - General
    - [x] PeggedERC721
    - [ ] PeggedERC20
    - [x] Operator
    - [ ] SDK
 - Plasma Cash
    - [x] ERC721 Deposit
    - [x] ERC721 Exit
    - Challenge
 - Plasma Debit
    - [ ] ERC20 (Capped Payment) Deposit
    - [ ] ERC20 (Capped Payment) Exit
 - Misc
    - [ ] Demo Frontend
