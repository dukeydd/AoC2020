package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInput() string {  
    data, err := ioutil.ReadFile("input1")
    if err != nil {
        fmt.Println("File reading error", err)
	}
    return string(data)
}

func main() {
	data := getInput()
	dataList := strings.Split(data, "\n")
	dataRange := len(dataList)
	for i:=0; i<dataRange; i++ {
		for j:=i+1; j<dataRange; j++ {
			for k:=j+1; k<dataRange; k++ {
				iInt, err := strconv.Atoi(dataList[i])
				if err != nil {
					fmt.Println("Error converting type", err)
				}
				jInt, err := strconv.Atoi(dataList[j])
				if err != nil {
					fmt.Println("Error converting type", err)
				}
				kInt, err := strconv.Atoi(dataList[k])
				if err != nil {
					fmt.Println("Error converting type", err)
				}
				if  iInt + jInt +kInt == 2020 {
					fmt.Println(iInt*jInt*kInt)
				}
			}
		}
	}
}