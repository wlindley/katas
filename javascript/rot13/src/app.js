const async = require('async');
const process = require('process');
const path = require('path');
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
		this._process = process;
	}

	execute(callback) {
		async.waterfall([
			(cb) => this._file.read(this._modifyPath(this._input), cb),
			(plaintext, cb) => rot13(plaintext, cb),
			(ciphertext, cb) => this._file.write(this._modifyPath(this._output), ciphertext, cb)
		], callback);
	}

	_modifyPath(filePath) {
		if ('development' === this._process.env.NODE_ENV)
			return path.join('test', filePath);
		return filePath;
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
