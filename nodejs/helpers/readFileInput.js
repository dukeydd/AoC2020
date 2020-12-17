const fs = require('fs');

function getNumbers() {
    const data = fs.readFileSync('input', 'utf-8');

    const numbers = data.split('\n').map(function(item) {
        return parseInt(item);
    });
    return numbers
}

module.exports.getNumbers = getNumbers;