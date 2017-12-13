class ConsoleView {
	displayBoard(board) {
		let display = '';
		for (let y = 0; y < board.length; y++) {
			for (let x = 0; x < board[y].length; x++)
				display += board[y][x];
			display += "\n";
		}
		console.log(display);
	}
}
module.exports = ConsoleView;
