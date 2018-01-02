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
			trainFileContents('The dog barks at midnight.');
			testObj.execute(() => {
				verifyWrittenFilename(output);
				verifyWrittenData('Gur qbt onexf ng zvqavtug.');
				done();
			});
		});

		it('closes output file when complete', (done) => {
			trainFileContents();
			testObj.execute(() => {
				verifyOutputFileClosed();
				done();
			});
		});

		it('writes to test directory in development', (done) => {
			trainFileContents();
			trainEnvironment('development');
			testObj.execute(() => {
				verifyReadFilename(path.join('test', input));
				verifyWrittenFilename(path.join('test', output));
				done();
			});
		});

		it('writes to current directory in production', (done) => {
			trainFileContents();
			trainEnvironment('production');
			testObj.execute(() => {
				verifyReadFilename(input);
				verifyWrittenFilename(output);
				done();
			});
		});
	});

	function trainFileContents(contents='') {
		fileFake.setReadResult(contents);
	}

	function trainEnvironment(environment) {
		processFake.env.NODE_ENV = environment;
	}

	function verifyWrittenFilename(expected) {
		expect(fileFake.writtenInfo.filename).to.equal(expected);
	}

	function verifyWrittenData(expected) {
		expect(fileFake.writtenInfo.data).to.equal(expected);
	}

	function verifyReadFilename(expected) {
		expect(fileFake.readInfo.filename).to.equal(expected);
	}

	function verifyOutputFileClosed() {
		expect(fileFake.writtenInfo.wasClosed).to.be.true;
	}
});
