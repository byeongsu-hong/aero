const fs = require('fs');
const Service = require('./base/Service');
const EventWatcher = require('./EventWatcher');

class DepositService extends Service {

    constructor(parentChain, childChain, deployments) {
        super();

        this.parentChain = parentChain;
        this.childChain = childChain;

        this.parentContract = this.loadContract(parentChain, deployments['ParentChain']);
        this.childContract = this.loadContract(childChain, deployments['ChildChain']);

        // watch deposit event on parent chain
        this.watcher = new EventWatcher(
            parentChain,
            this.parentContract,
            (event) => this.onDeposit(event));
    }

    start() {
        this.watcher.start();
    }

    onDeposit(eventData) {
        const event = {
            name: eventData.event,
            data: eventData.returnValues,
        };
        if (event.name === 'DepositEvent') {
            // submit block to child chain.
            await this.deposit(
                eventData.depositor,
                eventData.depositBlockNumber,
                eventData.token,
                eventData.amount
            );
        }
    }
    
    async deposit(depositorAddress, depositBlockNumber, tokenAddress, amount) {
        const tx = this.childContract.methods.submitDeposit(
                depositorAddress,
                depositBlockNumber,
                tokenAddress,
                amount
        );        
        await this.sendAndWaitForTx(tx);
    }
}