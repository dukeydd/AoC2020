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


func findNumTrees(xDiff int, yDiff int) int {
	dataList, dataRange := getInput()
	// fmt.Println(reflect.TypeOf(dataList[0]))
	var coords [2]int
	coords[0], coords[1] = 0, 0
	xMax := len(dataList[0])

	numTrees:=0
	for i:=0; i<dataRange; i=i+yDiff {
		fmt.Println(coords, dataList[i])

		if coords[0] >= xMax {
			coords[0]-= xMax
		}
		if dataList[i][coords[0]] == '#' {
			numTrees++
		}

		coords[0]+=xDiff
		coords[1]+=yDiff
	}
	fmt.Println(numTrees)
	return numTrees
}

func main() {
	scenario1 := findNumTrees(1,1)
	scenario2 := findNumTrees(3,1)
	scenario3 := findNumTrees(5,1)
	scenario4 := findNumTrees(7,1)
	scenario5 := findNumTrees(1,2)
	multiplyTrees := scenario1 * scenario2 * scenario3 * scenario4 * scenario5
	fmt.Println(multiplyTrees)
}
