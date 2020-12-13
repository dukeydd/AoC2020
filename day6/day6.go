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
	dataList := strings.Split(string(data), "\n\n")
	dataRange := len(dataList)
    return dataList, dataRange
}

func findNumQuestions(entry string) int {
	set := make(map[string]bool)
	for i:=0; i<len(entry); i++ {
		set[string(entry[i])] = true
	}
	delete(set, "\n")
	return len(set)
}

func main() {
	dataList, dataRange := getInput()

	sum := 0
	// fmt.Println(dataList[0])
	for i:=0; i<dataRange; i++ {
		sum += findNumQuestions(string(dataList[i]))
	}
	fmt.Println(sum)
}
