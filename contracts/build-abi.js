/**
 * Builds ABI for aeroracle.
 */
const fs = require('fs');
const spawn = require('child_process').spawn;

const TARGETS = ['ParentChain', 'Pegger'];
const TARGETS_DIR = './build';
const DEST = './build/abi';

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

// WARN: hardcoded build script
const parentCommand = `-pkg contracts -type ParentChain -abi ${DEST}/ParentChain.abi -out binds/parent_chain.go`;
const childCommand = `-pkg contracts -type ChildChain -abi ${DEST}/Pegger.abi -out binds/child_chain.go`;

spawn('abigen', parentCommand.split(' '), { stdio: 'inherit' }).on('exit', (code) => {
    if (code !== 0) return;
    spawn('abigen', childCommand.split(' '), { stdio: 'inherit' }).on('exit', process.exit);
});
