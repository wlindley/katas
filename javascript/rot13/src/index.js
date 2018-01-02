const App = require('./app');

const config = {
	input: process.argv[2],
	output: process.argv[3]
};
const app = new App(config);
app.execute(() => {});
