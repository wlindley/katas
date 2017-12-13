const Drop7 = require('./drop7');
const ConsoleView = require('./console-view');

const width = 4;
const height = 4;

const app = new Drop7(width, height, new ConsoleView());
app.start();
