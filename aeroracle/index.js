const Web3 = require('web3');
const DepositService = require('./src/DepositService');

// hardcoded config
const config = {
    oraclePrivateKey: '0x17af9b432ced6167a4dba5cec6517a0d545abe118dbd494572e07e6a2847ff39',
    parentChainEndpoint: 'https://ropsten.infura.io/v3/e9783a0db20b4c85b1438019c6432dc3',
    childChainEndpoint: 'http://localhost:8645',
    deployments: [
        ParentChain: {
            "abiPath": './abi/ParentChain.abi',
            "deployedAt": '0x0000000000000000000000000000000000000000',
        },
        ChildChain: {
            "abiPath": './abi/ChildChain.abi',
            "deployedAt": '0x0000000000000000000000000000000000000000',
        },
    ],
};

async function main(config) {
    const parentChain = new Web3(new Web3.providers.HttpProvider(config.parentChainEndpoint));
    const chlidChain = new Web3(new Web3.providers.HttpProvider(childChainEndpoint));

    // set up oracle account
    const account = parentChain.eth.accounts.privateKeyToAccount(config.oraclePrivateKey);
    parentChain.eth.defaultAccount = account;
    childChain.eth.defaultAccount = account;

    const depositService = new DepositService(parentChain, childChain, config.deployments);
    depositService.start();
}

// run!
main(config)
    .catch(err => console.log(err.stack));
