const os = require('node:os')
const fs = require('node:fs');

console.log("User info:")
console.log(os.userInfo());

console.log("Hostname:")
console.log(os.hostname());

console.log("Env:")
console.log(process.env);

console.log(`CWD: ${process.cwd()}`);

const currentCommand = process.argv.join(' ');
console.log(`Current running command: ${currentCommand}`);

const data = fs.readFileSync('/etc/os-release', 'utf8');
console.log("/etc/os-release:")
console.log(data);
