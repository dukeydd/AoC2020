const readFileInput = require('../helpers/readFileInput');
const data = readFileInput.getPasswords();

function getValidPasswords(data) {
    let valid = 0;
    data.forEach(entry => {
        // if (entry.pswd.charAt(entry.quant[0]-1) === entry.letter && entry.pswd.charAt(entry.quant[1]-1) === entry.letter)
        let letterCount = 0;
        for (let i = 0; i < entry.pswd.length; i++) {
            if (entry.pswd.charAt(i) === entry.letter) {
                letterCount++;
            }
        }
        if (letterCount <= entry.quant[1] && letterCount >= entry.quant[0]) {
            valid++;
        }
    });
    console.log(valid)
}

getValidPasswords(data);