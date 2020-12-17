package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInput() ([]string, int) {  
    data, err := ioutil.ReadFile("input")
    if err != nil {
        fmt.Println("File reading error", err)
	}
	dataList := strings.Split(string(data), "\n")
	dataRange := len(dataList)
    return dataList, dataRange
}

func splitInputRow(inputLine string) ([]string, string, string) {
	inputArray := strings.Split(inputLine, " ")
	pswd := inputArray[2]
	pswdLetter := strings.Trim(inputArray[1], ":")
	pswdQuantity := strings.Split(inputArray[0], "-")
	return pswdQuantity, pswdLetter, pswd
}

func main() {
	dataList, dataRange := getInput()
	valid:=0

	for i := 0; i < dataRange; i++ {
		pswdQuantity, pswdLetter, pswd := splitInputRow(string(dataList[i]))
		// numReqLetter := strings.Count(pswd, pswdLetter)
		pswdPositionOne, err := strconv.Atoi(pswdQuantity[0])
		if err != nil {
			fmt.Println(err)
		}
		pswdPositionTwo, err := strconv.Atoi(pswdQuantity[1])
		if err != nil {
			fmt.Println(err)
		}

		pswdArray := strings.Split(pswd, "")
		correctIndexes := 0

		if len(pswd) >= pswdPositionOne {
			if pswdArray[pswdPositionOne-1] == pswdLetter {
				correctIndexes++
			}
		}

		if len(pswd) >= pswdPositionTwo {
			if pswdArray[pswdPositionTwo-1] == pswdLetter {
				correctIndexes++
			}
		} 
		
		// only one allowed
		if correctIndexes == 1 {
			valid++
		}
	}

	fmt.Println(valid)
}