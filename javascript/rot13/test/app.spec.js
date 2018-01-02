const expect = require('chai').expect;
const App = require('../src/app');
const FileFake = require('./file-fake');

describe('App', () => {
	let testObj, input, output, fileFake;

	beforeEach(() => {
		input = 'in.txt';
		output = 'out.txt';
		fileFake = new FileFake();
		testObj = new App({input, output});
		testObj._file = fileFake;
	});

	describe('execute', () => {
		it('writes rotated input to output', (done) => {
			fileFake.setReadResult('The dog barks at midnight.');
			testObj.execute(() => {
				expect(fileFake.writtenInfo.filename).to.equal(output);
				expect(fileFake.writtenInfo.data).to.equal('Gur qbt onexf ng zvqavtug.');
				done();
			});
		});
	});
});
