package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
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

func getNewRange(front int, back int, s string) (int, int) {
	// this needs to be changed to add or subtract half difference of upper-lower not /2
	diff := back - front + 1
	
	if s == "F" || s == "L" {
		return front, (back - diff/2)
	} else if s == "B" || s == "R" {
		return (front + diff/2), back
	} else {
		fmt.Println("input incorrect")
		panic("ahh")
	}	
}

func getNums(rowID string, columnID string) (int, int) {
	var rowNum int
	var columnNum int
	// find row number
	currentFront, currentBack := 0, 127
	for j:=0; j<len(rowID); j++ {
		currentFront, currentBack = getNewRange(currentFront, currentBack, string(rowID[j]))
	}
	// check they are the same
	if currentBack == currentFront {
		rowNum = currentFront
	} else {
		fmt.Println("something went wrong")
	}

	// find column number
	currentLeft, currentRight := 0, 7
	for k:=0; k<len(columnID); k++ {
		currentLeft, currentRight = getNewRange(currentLeft, currentRight, string(columnID[k]))
	}
	// check they are the same
	if currentLeft == currentRight {
		columnNum = currentLeft
	} else {
		fmt.Println("something went wrong")
	}
	return rowNum, columnNum
}

func calculateSeatID(rowNum int, columnNum int) int {
	return rowNum*8 + columnNum
}

func getMaxInt(nums []int) int {
	var m int
	for i, e := range nums {
		if i==0 || e > m {
			m = e
		}
	}
	return m 
}

func main() {
	dataList, dataRange := getInput()

	var seatIDs []int 
	for i:=0; i<dataRange; i++ {
		rowID, columnID := getIDs(dataList[i])
		rowNum, columnNum := getNums(rowID, columnID)
		seatIDs = append(seatIDs, calculateSeatID(rowNum, columnNum))
	}

	sort.Ints(seatIDs)
	// fmt.Println(seatIDs)
	for j:=0; j<len(seatIDs); j++ {
		if j != 0 {
			if seatIDs[j+1] != seatIDs[j] +1 {
				fmt.Println(seatIDs[j] + 1)
				return
			}
		}
	}
}