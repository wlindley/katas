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
			(cb) => this._readFile(cb),
			(plaintext, cb) => rot13(plaintext, cb),
			(ciphertext, cb) => this._writeFile(ciphertext, cb)
		], callback);
	}

	_readFile(callback) {
		const modifiedPath = this._modifyPath(this._input);
		this._file.read(modifiedPath, callback);
	}

	_writeFile(ciphertext, callback) {
		const modifiedPath = this._modifyPath(this._output);
		this._file.write(modifiedPath, ciphertext, callback);
	}

	_modifyPath(filePath) {
		if (this._isDevelopment)
			return path.join('test', filePath);
		return filePath;
	}

	get _isDevelopment() {
		return 'development' === this._process.env.NODE_ENV;
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
	const alphaIndex = char.charCodeAt(0) - base;
	const rotatedIndex = (alphaIndex + 13) % 26;
	const rotatedCode = rotatedIndex + base;
	return String.fromCharCode(rotatedCode);
}

module.exports = App;
