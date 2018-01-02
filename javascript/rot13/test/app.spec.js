const expect = require('chai').expect;
const App = require('../src/app');
const FileFake = require('./file-fake');
const path = require('path');

describe('App', () => {
	let testObj, input, output, fileFake, processFake;

	beforeEach(() => {
		input = 'in.txt';
		output = 'out.txt';
		fileFake = new FileFake();
		processFake = {env: {}};
		testObj = new App({input, output});
		testObj._file = fileFake;
		testObj._process = processFake;
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

		it('writes to test directory in development', (done) => {
			fileFake.setReadResult('');
			processFake.env.NODE_ENV = 'development';
			testObj.execute(() => {
				expect(fileFake.readInfo.filename).to.equal(path.join('test', input));
				expect(fileFake.writtenInfo.filename).to.equal(path.join('test', output));
				done();
			});
		});

		it('writes to current directory in production', (done) => {
			fileFake.setReadResult('');
			processFake.env.NODE_ENV = 'production';
			testObj.execute(() => {
				expect(fileFake.readInfo.filename).to.equal(input);
				expect(fileFake.writtenInfo.filename).to.equal(output);
				done();
			});
		});
	});
});
