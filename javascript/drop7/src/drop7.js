class Drop7 {
	constructor(width, height, view) {
		this._board = [];
		this._initBoard(height, width);
		this._view = view;
	}

	start() {
		this._view.displayBoard(this._board);
	}

	_initBoard(rows, cols) {
		for (let r = 0; r < rows; r++) {
			this._board.push([]);
			for (let c = 0; c < cols; c++)
				this._board[r].push('.');
		}
	}
}
module.exports = Drop7;
