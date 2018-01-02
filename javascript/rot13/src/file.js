const fs = require('fs');

class File {
	read(filename, callback) {
		fs.readFile(filename, 'utf-8', callback);
	}

	write(filename, data, callback) {
		fs.writeFile(filename, data, callback);
	}
}
module.exports = File;