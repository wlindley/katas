class Drop7 {
	constructor(rows, cols, view) {
		this._board = new Board(rows, cols);
		this._view = view;
	}

	start() {
		this._view.displayBoard(this._board);
	}
}

class Board {
	constructor(rows, cols) {
		this._rows = rows;
		this._cols = cols;
	}

	get cols() {
		return this._cols;
	}

	get rows() {
		return this._rows;
	}

	at(row, col) {
		return '.';
	}
}

module.exports = Drop7;
