// test.js

const assert = require('assert');

// Simple function to add two numbers
function addNumbers(a, b) {
    return a + b;
}

// Mocha test suite
describe('Addition', function () {
    // Mocha test case
    it('should return the sum of two numbers', function () {
        assert.strictEqual(addNumbers(5, 7), 12);
    });
});
