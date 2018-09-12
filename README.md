Aero
=======

Sidechain framework used on Airbloc. 

## Glossaries

* Parent Chain : A parent chain that provides asset security to child chain. 
* Child Chain : A child chain or sidechain that the asset is pegged into the parent chain. The block headers of child chain is submitted (committed) to the parent chain regularly, and users can exit to the parent chain using the submitted data.
* Operator : Block producer of the child chain. Operator has a duty to commit the child chain blocks to the parent chain regularly.
* Gateway Contract : Contracts on each chain, bridging the transfer between two chains.
* Gateway Oracle : An (trustful) oracle that reports the other chain's information. 
    * Aeroracle : Aero's gateway oracle.
* Enter (Deposit) : Deposit the funds to the child chain.
* Exit (Withdrawal) : Withdraw the funds from the child chain, however depending on root chain.

## Components

 - `aerod` - Aero sidechain daemon, running PoA Ethereum network
 - `aeroracle` - Gateway oracle that connects between two chain.
 - `contracts` - Contracts and utilities for both the child and the parent chain.

## Prerequisites

 - Go 1.9+
 - Node 10.0+
 - [Yarn](https://yarnpkg.com/lang/en/) - Package manager for Node
 - [Dep](http://github.com/golang/dep) - Dependency manager for Go
 - [Truffle](https://truffleframework.com/) - Contract deployment framework

## Installation

```
make
```


## Current Status

 - Plasma MVP
    - [x] ERC20 Deposit
    - [ ] Block Submit
    - [ ] ERC20 Exit
    - [ ] Gateway Oracle
 - Plasma Cash
    - [ ] ERC721 Deposit
    - [ ] ERC721 Exit
 - Plasma Debit
    - [ ] ERC721 -> ERC20 Deposit
    - [ ] ERC721 -> ERC20 Exit
    - [ ] CAS Challenge
 - Misc
    - [ ] Demo Frontend
