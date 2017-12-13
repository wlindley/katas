const chai = require('chai');
const sinon = require('sinon');
const expect = chai.expect;
chai.use(require('sinon-chai'));
const Drop7 = require('../src/drop7');

describe('Drop7', () => {
	let testObj, width, height, view;

	beforeEach(() => {
		width = 2;
		height = 2;
		view = {displayBoard: sinon.spy()};
		testObj = new Drop7(width, height, view);
	});

	describe('start', () => {
		it('calls displayBoard with 2D array of expected size', () => {
			testObj.start();
			expect(view.displayBoard).to.have.been.calledWithMatch(sinon.match(value => {
				return value.length == height && value[0].length == width;
			}));
		});
	});
});
