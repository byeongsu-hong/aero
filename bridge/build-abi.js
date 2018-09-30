/**
 * Builds ABI for aeroracle.
 */
const fs = require('fs');
const spawn = require('child_process').spawn;

const TARGETS = ['ParentBridge', 'ChildBridge'];
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
    const { abi, bytecode } = truffleBuild;
    fs.writeFileSync(`${DEST}/${target}.abi`, JSON.stringify(abi));
    fs.writeFileSync(`${DEST}/${target}.bin`, bytecode);
});

// WARN: hardcoded build script
const parentCommand = `-pkg contracts -type ParentBridge -abi ${DEST}/ParentBridge.abi -bin ${DEST}/ParentBridge.bin -out binds/parent_bridge.go`;
const childCommand = `-pkg contracts -type ChildBridge -abi ${DEST}/ChildBridge.abi -bin ${DEST}/ChildBridge.bin -out binds/child_bridge.go`;

spawn('abigen', parentCommand.split(' '), { stdio: 'inherit' }).on('exit', (code) => {
    if (code !== 0) return;
    spawn('abigen', childCommand.split(' '), { stdio: 'inherit' }).on('exit', process.exit);
});
