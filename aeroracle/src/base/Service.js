const fs = require('fs');

class Service {

    constructor(minConfirmations = 2) {
        this.minConfirmations = minConfirmations;
    }

    loadContract(web3, deployment) {
        const abi = JSON.parse(fs.readFileSync(deployment.abiPath));
        return new web3.eth.Contract(abi, deployment.address);
    }
    
    async sendAndWaitForTx(tx) {
        return new Promise((resolve, reject) => {
            tx.send()
                .on('error', reject)
                .on('confirmation', (confirmNumber, receipt) => {
                    if (confirmNumber >= this.minConfirmations) return resolve(receipt);
                });
        });
    }
}

module.exports = Service;
