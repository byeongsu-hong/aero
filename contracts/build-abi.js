/**
 * Builds ABI for aeroracle.
 */
const fs = require('fs');

const TARGETS = ['ParentChain', 'ChildChain', 'Pegger'];
const TARGETS_DIR = './build';
const DEST = '../aeroracle/abi';

if (!fs.existsSync(TARGETS_DIR)) {
    console.log('Please build the contracts first, using "npm run compile".');
    process.exit(1);
}

try {
    fs.mkdirSync(DEST);
} catch (err) {
    // ignored
}

TARGETS.forEach(target => {
    const truffleBuild = require(`${TARGETS_DIR}/contracts/${target}.json`);
    const abi = JSON.stringify(truffleBuild.abi);
    fs.writeFileSync(`${DEST}/${target}.abi`, abi);
});
