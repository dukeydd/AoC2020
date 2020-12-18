const fs = require('fs');

function getNumbers() {
    const data = fs.readFileSync('input', 'utf-8');

    const numbers = data.split('\n').map(function(item) {
        return parseInt(item);
    });
    return numbers;
}


function getPasswords() {
    const pswdArray = [];
    const data = fs.readFileSync('input', 'utf-8')
    
    const lines = data.split('\n');
    lines.forEach(line => {
        const pswdDict = {};
        const pswdEntry = line.split(' ');
        pswdDict.pswd = pswdEntry[2];
        pswdDict.letter = pswdEntry[1].slice(0, -1);
        pswdDict.quant = pswdEntry[0].split('-').map((item) => {
            return parseInt(item);
        });
        pswdArray.push(pswdDict);
    });
    return pswdArray;
}

module.exports.getNumbers = getNumbers;
module.exports.getPasswords = getPasswords;