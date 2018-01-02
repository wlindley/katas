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
		const inputStream = this._getReadStream();
		const outputStream = this._getWriteStream();
		inputStream.on('data', (chunk) => {
			outputStream.write(rot13(chunk));
		});
		inputStream.on('close', () => {
			outputStream.close(callback);
		});
	}

	_getReadStream() {
		const modifiedPath = this._modifyPath(this._input);
		return this._file.read(modifiedPath);
	}

	_getWriteStream() {
		const modifiedPath = this._modifyPath(this._output);
		return this._file.write(modifiedPath);
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

function rot13(plaintext) {
	return Array.from(plaintext).map(rot13Char).join('');
}

function rot13Char(char) {
	if (isLowerCase(char))
		return rotate(char, charCodes['a']);
	if (isUpperCase(char))
		return rotate(char, charCodes['A']);
	return char;
}

function isLowerCase(char) {
	return 'a' <= char && char <= 'z';
}

function isUpperCase(char) {
	return 'A' <= char && char <= 'Z';
}

function rotate(char, base) {
	const alphaIndex = char.charCodeAt(0) - base;
	const rotatedIndex = (alphaIndex + 13) % 26;
	const rotatedCode = rotatedIndex + base;
	return String.fromCharCode(rotatedCode);
}

module.exports = App;
