package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func getIDs(input string) (string, string) {
	rowID := string(input[0:7])
	columnID := string(input[7:])
	return rowID, columnID
}

func getNewRange(upper int, lower int, s string) (int, int) {
	// this needs to be changed to add or subtract half difference of upper-lower not /2
	if s == "F" {
		return upper, int(lower/2)
	} else if s == "B" {
		return int(upper/2), lower
	} else {
		fmt.Println("input incorrect")
		panic("ahh")
	}
}

// func calculateSeatLoc(id string, columnId string) int {
// 	for i:=0; i<len(id)
// }

func main() {
	dataList, dataRange := getInput()
	for i:=0; i<dataRange; i++ {
		rowID, columnID := getIDs(dataList[i])
		newUp, newLow := getNewRange(0, 127, "B")
		fmt.Println(newUp, newLow)
	}
}