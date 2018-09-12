const Web3 = require('web3');
const config = require('../config.json');

// parent chain
const parent = new Web3(new Web3.providers.HttpProvider());
