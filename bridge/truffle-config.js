
const mochaGasSettings = {
    reporter: 'eth-gas-reporter',
    reporterOptions : {
        currency: 'USD',
        gasPrice: 3
    }
};

const mocha = process.env.GAS_REPORTER ? mochaGasSettings : {};

module.exports = {
    networks: {
        child: {
            network_id: 2470,
            host: 'localhost',
            port: 7600,
            gas: 400000,
            gasPrice: 0,
        },
        coverage: {
            host: 'localhost',
            network_id: '*',
            port: 8555,
            gas: 0xffffffffff,
            gasPrice: 0x01,
        },
    },
    build: {},
    mocha,
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
};
