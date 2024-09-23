console.log(process.env);

console.log(`CWD: ${process.cwd()}`);

const currentCommand = process.argv.join(' ');
console.log(`Current running command: ${currentCommand}`);

// console.log("exit 1");
// process.exit(1);
