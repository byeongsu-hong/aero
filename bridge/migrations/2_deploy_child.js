const OperatorRegistry = artifacts.require('OperatorRegistry');

module.exports = async function (deployer, network, accounts) {

    if (network !== 'child') {
        // skip parent chain
        return;
    }

    const operatorRegistry = await deployer.deploy(OperatorRegistry);
    const isMeOperator = await operatorRegistry.isOperator.call(accounts[0]);
    console.log(`The address ${accounts[0]} is operator? ${isMeOperator}!`);
};
