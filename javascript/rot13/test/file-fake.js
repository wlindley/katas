const async = require('async');
const EventEmitter = require('events');

class FileFake {
	read(filename, callback) {
		this.readInfo = {filename};
		return new ReadStreamFake(this._readData);
	}

	write(filename) {
		this.writtenInfo = {filename, data: '', wasClosed: false};
		return new WriteStreamFake(this.writtenInfo);
	}

	setReadResult(data) {
		this._readData = data;
	}
}

class ReadStreamFake extends EventEmitter {
	constructor(fakeData) {
		super();
		async.nextTick(() => {
			this.emit('data', fakeData);
			async.nextTick(() => this.emit('close'));
		});
	}
}

class WriteStreamFake {
	constructor(info) {
		this._info = info;
	}

	write(data) {
		this._info.data += data;
	}

	close(callback) {
		this._info.wasClosed = true;
		async.nextTick(callback);
	}
}

module.exports = FileFake;
