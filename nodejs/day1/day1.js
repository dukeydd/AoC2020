const readFileInput = require('../helpers/readFileInput');

const dataByLine = readFileInput.getNumbers();

// Part 1
for (let i=0; i< dataByLine.length; i++) {
    for (let j=i+1; j<dataByLine.length; j++) {
        if (dataByLine[i] + dataByLine[j] == 2020) {
            console.log(dataByLine[i] * dataByLine[j])
        }
    }
}

// Part 2
for (let i=0; i< dataByLine.length; i++) {
    for (let j=i+1; j<dataByLine.length; j++) {
        for (let k=j+1; k<dataByLine.length; k++) {
            if (dataByLine[i] + dataByLine[j] +dataByLine[k]== 2020) {
                console.log(dataByLine[i] * dataByLine[j] * dataByLine[k])
            }
        }
        
    }
} 