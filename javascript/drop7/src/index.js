const Drop7 = require('./drop7');
const ConsoleView = require('./console-view');
const readline = require('readline');

const cols = 4;
const rows = 4;

const app = new Drop7(rows, cols, new ConsoleView());
app.start();
const rl = readline.createInterface({
	input: process.stdin,
	output: process.stdout
})
rl.on('line', line => {
	
});
