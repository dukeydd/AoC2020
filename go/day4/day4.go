package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "reflect"
	"strconv"
	"regexp"
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

func formatPassport(passport string) map[string]string {
	m := make(map[string]string)

	lines := strings.Split(passport, "\n")
	for _, line := range(lines) {
		records := strings.Split(line, " ")
		
		for _, record := range(records) {
			finally := strings.Split(record, ":")
			m[finally[0]] = finally[1]
		}
	}
	return m
}

func validatePassport(passport map[string]string,  valid *int) {
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return
	}
	fmt.Println(byr)
	length := len(passport["byr"])
	if passport["byr"] != "" && length == 4 && byr >= 1920 && byr <= 2002 {
		
		iyr, err := strconv.Atoi(passport["iyr"])
		if err != nil {
			return
		}
		fmt.Println(iyr)
		length := len(passport["iyr"])
		if passport["iyr"] != "" && length == 4 && iyr >= 2010 && iyr <= 2020 {
		
			eyr, err := strconv.Atoi(passport["eyr"])
			if err != nil {
				return
			}
			fmt.Println(eyr)
			length := len(passport["eyr"])
			if passport["eyr"] != "" && length == 4 && eyr >= 2020 && eyr <= 2030 {
		
				if passport["hgt"] != "" {
					if strings.HasSuffix(passport["hgt"], "cm") == true {
						height, err := strconv.Atoi(strings.Trim(passport["hgt"], "cm"))
						if err != nil {
							return
						}
						fmt.Println(height)
						if height < 150 || height > 193 {
							return
						}
					} else if strings.HasSuffix(passport["hgt"], "in") == true {
						height, err := strconv.Atoi(strings.Trim(passport["hgt"], "in"))
						if err != nil {
							return
						}
						fmt.Println(height)
						if height < 59 || height > 76 {
							return
						}
					} else {
						return
					}

					if passport["hcl"] != "" {
						matched, err := regexp.Match(`^#[a-f0-9]{6}$`, []byte(passport["hcl"]))
						if err != nil {
							return
						}
						if matched == false {
							return
						}
						fmt.Println(passport["hcl"], matched)
						if passport["ecl"] != "" && (passport["ecl"] == "amb" || passport["ecl"] == "blu" || passport["ecl"] == "brn" || passport["ecl"] == "gry" || passport["ecl"] == "grn" || passport["ecl"] == "hzl" || passport["ecl"] == "oth") {
							fmt.Println(passport["ecl"])
							if passport["pid"] != "" {
								matched, err := regexp.Match(`^[0-9]{9}$`, []byte(passport["pid"]))
								if err != nil {
									return
								}
								if matched == true {
									*valid++
								}
								fmt.Println(passport["pid"])
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	dataList, _ := getInput()
	valid := 0
	for _, data := range(dataList) {
		passport := formatPassport(data)
		fmt.Println(passport)
		validatePassport(passport, &valid)
	}
	fmt.Println(valid)

}