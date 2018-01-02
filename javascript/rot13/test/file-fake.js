const async = require('async');

class FileFake {
	read(filename, callback) {
		this.readInfo = {filename};
		async.nextTick(() => callback(null, this._readData));
	}

	write(filename, data, callback) {
		this.writtenInfo = {filename, data};
		async.nextTick(callback);
	}

	setReadResult(data) {
		this._readData = data;
	}
}
module.exports = FileFake;
