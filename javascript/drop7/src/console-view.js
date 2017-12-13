class ConsoleView {
	displayBoard(board) {
		let display = '';
		for (let r = 0; r < board.rows; r++) {
			for (let c = 0; c < board.cols; c++)
				display += board.at(r, c);
			display += "\n";
		}
		console.log(display);
	}
}
module.exports = ConsoleView;
