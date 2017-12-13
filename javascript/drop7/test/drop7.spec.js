const chai = require('chai');
const sinon = require('sinon');
const expect = chai.expect;
chai.use(require('sinon-chai'));
const Drop7 = require('../src/drop7');
const Cell = require('../src/cell');

describe('Drop7', () => {
	let testObj, cols, rows, view, board;

	beforeEach(() => {
		board = null;
		cols = 2;
		rows = 2;
		view = {displayBoard: sinon.spy()};
		testObj = new Drop7(rows, cols, view);
	});

	describe('start', () => {
		it('calls displayBoard with empty board', () => {
			testObj.start();
			expect(view.displayBoard).to.have.been.calledWithMatch(sinon.match(b => {
				board = b;
				expectEmpty(0, 0);
				expectEmpty(0, 1);
				expectEmpty(1, 0);
				expectEmpty(1, 1);
				return board.rows == rows && board.cols == cols;
			}));
		});
	});

	function expectEmpty(row, col) {
		expect(board.at(row, col)).to.equal(Cell.Empty);
	}
});
