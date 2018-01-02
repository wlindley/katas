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
		it('returns stream that can be read from', (done) => {
			const stream = testObj.read('test/in.txt');
			let contents = '';
			stream.on('data', (chunk) => contents += chunk);
			stream.on('close', () => {
				expect(contents).to.equal("The dog barks at midnight");
				done();
			});
		});
	});

	describe('write', () => {
		it('writes data to given file', (done) => {
			const filePath = 'test/out.txt';
			const stream = testObj.write(filePath);
			stream.write('hello, world!');
			stream.write(' this is some data');
			stream.close(() => {
				fs.readFile(filePath, 'utf-8', (err, contents) => {
					expect(contents).to.equal('hello, world! this is some data');
					done();
				});
			});
		});
	});
});
