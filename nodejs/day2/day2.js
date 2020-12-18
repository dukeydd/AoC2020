const readFileInput = require('../helpers/readFileInput');
const data = readFileInput.getPasswords();

// Part 1
function getValidPasswords(data) {
    let valid = 0;
    data.forEach(entry => {
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

// Part 2
function getValidPasswordsPart2(data) {
    let valid = 0;
    data.forEach(entry => {
        let correctIndexes = 0;
        if (entry.pswd.charAt(entry.quant[0]-1) === entry.letter) {
            correctIndexes++;
        }
        if (entry.pswd.charAt(entry.quant[1]-1) === entry.letter) {
            correctIndexes++;
        }
        if (correctIndexes == 1) {
            valid++;
        }
    });
    console.log(valid)
}
getValidPasswordsPart2(data);