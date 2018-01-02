const EventEmitter = require('events');
const fs = require('fs');

class File {
	read(filename) {
		return new ReadStream(filename);
	}

	write(filename) {
		return new WriteStream(filename);
	}
}

class ReadStream extends EventEmitter {
	constructor(filePath) {
		super();
		this._stream = fs.createReadStream(filePath, {encoding: 'utf-8'});
		this._stream.on('data', (chunk) => this.emit('data', chunk));
		this._stream.on('close', () => this.emit('close'));
	}
}

class WriteStream {
	constructor(filePath) {
		this._stream = fs.createWriteStream(filePath);
	}

	write(data) {
		this._stream.write(data);
	}

	close(callback) {
		this._stream.end(callback);
	}
}

module.exports = File;