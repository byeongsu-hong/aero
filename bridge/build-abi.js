/**
 * Builds ABI for aeroracle.
 */
const fs = require('fs');
const spawn = require('child_process').spawn;
const uncamelize = str => str.replace('ABLToken', 'abl_token').replace('ERC', '_erc').replace(/(?:^|\.?)([A-Z])/g, (x,y) => "_" + y.toLowerCase()).replace(/^_/, "");

const TARGETS = ['ParentBridge', 'ChildBridge', 'ABLToken', "PeggedERC20", "PeggedERC721"];
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

async function abigen(targets, pkgName = 'contracts') {
    for (const target of targets) {
        const args = `-pkg ${pkgName} `
            + `-type ${target} `
            + `-abi ${DEST}/${target}.abi `
            + `-bin ${DEST}/${target}.bin `
            + `-out binds/${uncamelize(target)}.go`;

        await new Promise((resolve, reject) => {
            spawn('abigen', args.split(' '), { stdio: 'inherit' })
                .on('exit', code => {
                    if (code !== 0) reject(new Error(`abigen exited with code ${code}`));
                    resolve();
                });
        });
    }
}

abigen(TARGETS).catch(err => console.error(err.stack));

