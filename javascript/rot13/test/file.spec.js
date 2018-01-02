const File = require('../src/file');
const fs = require('fs');
const expect = require('chai').expect;

describe('File', () => {
	let testObj;

	beforeEach((done) => {
		testObj = new File();
		fs.unlink('test/out.txt', () => done());
	});

	afterEach((done) => {
		fs.unlink('test/out.txt', () => done());
	});

	describe('read', () => {
		it('calls back with contents of file as string', (done) => {
			testObj.read('test/in.txt', (err, contents) => {
				expect(err).not.to.exist;
				expect(contents).to.equal("The dog barks at midnight");
				done();
			});
		});
	});

	describe('write', () => {
		it('writes data to given file', (done) => {
			const filePath = 'test/out.txt';
			const data = 'hello, world! this is some data';
			testObj.write(filePath, data, () => {
				testObj.read(filePath, (err, contents) => {
					expect(contents).to.equal(data);
					done();
				});
			});
		});
	});
});
