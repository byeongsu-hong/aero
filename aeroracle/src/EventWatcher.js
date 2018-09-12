const Web3 = require('web3');

/**
 * EventWatcher watches event on the gateway contract of parent/child chain.
 * Since using the WebSocket to subscribe to event can occur consistency problem,
 * we can use polling instead.
 */
class EventWatcher {

    constructor(web3, contract, callback, minConfimrations = 3) {
        this.web3 = web3;
        this.contract = contract;
        this.callback = callback;
        this.minConfimrations = minConfimrations;

        this.looper = null;
        this.seenEvents = {};
    }

    start() {
        this.looper = setInterval(() => {
            // poll event for every 15 seconds
            this.loop().catch(err => console.error(err.stack));
        }, 15 * 1000);
    }

    async loop() {
        const currentBlock = await this.web3.eth.getBlockNumber();

        const events = await this.contract.getPastEvents('allEvents', {
            fromBlock: currentBlock - (this.minConfimrations * 2 + 1),
            toBlock: currentBlock - this.minConfimrations + 1
        });
        
        for (const event of events) {
            const eventId = `${event.transactionHash}-${event.logIndex}`;
            if (!this.seenEvents[eventId]) continue;

            this.seenEvents[eventId] = true;
            this.callback(event);
        }
    }

    stop() {
        if (this.looper) clearInterval(this.looper);
    }
}
