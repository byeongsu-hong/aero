const Web3 = require('web3');

class Watcher {

    constructor(websocketEndpoint, callback, minConfimrations = 3) {
        this.web3 = new Web3(new Web3.providers.WebsocketProvider(websocketEndpoint));
        this.contracts = [];
        this.callback = callback;
        this.minConfimrations = minConfimrations;
    }

    /**
     * 
     * @param {String} address gateway contracts.
     * @param {Array} abi ABI of the gateway contract. 
     */
    addGatewayContract(address, abi) {
        const contract = new this.web3.eth.Contract(abi, address);
        this.contracts.push(contract);
    }

    start() {
        this.contracts.forEach(contract => {
            contract.events.allEvents()
                .on('data', event => {
                    // wait until confirmation— because some events (Deposit, Exit) needs finality.
                    this.callback
                })
                .on('error', console.error);
        });
    }

    stop() {
        // not implemented yet.
    }
}
