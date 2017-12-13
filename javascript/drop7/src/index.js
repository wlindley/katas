const Drop7 = require('./drop7');
const ConsoleView = require('./console-view');

const cols = 4;
const rows = 4;

const app = new Drop7(rows, cols, new ConsoleView());
app.start();
