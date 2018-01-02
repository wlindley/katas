const async = require('async');
const File = require('./file');

const charCodes = {
	'a': 'a'.charCodeAt(0),
	'A': 'A'.charCodeAt(0)
};

class App {
	constructor(config) {
		this._input = config.input;
		this._output = config.output;
		this._file = new File();
	}

	execute(callback) {
		async.waterfall([
			(cb) => this._file.read(this._input, cb),
			(plaintext, cb) => rot13(plaintext, cb),
			(ciphertext, cb) => this._file.write(this._output, ciphertext, cb)
		], callback);
	}
}

function rot13(plaintext, callback) {
	const ciphertext = Array.from(plaintext).map(rot13Char).join('');
	callback(null, ciphertext);
}

function rot13Char(char) {
	if ('a' <= char && char <= 'z')
		return rotate(char, charCodes['a']);
	if ('A' <= char && char <= 'Z')
		return rotate(char, charCodes['A']);
	return char;
}

function rotate(char, base) {
	return String.fromCharCode((((char.charCodeAt(0) - base) + 13) % 26) + base);
}

module.exports = App;
