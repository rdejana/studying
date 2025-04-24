package roman

import (
	"fmt"
	"log"
)

func InitMap() map[string]int {
	mymap := make(map[string]int)
	//I, V, X, L, C, D, and M, representing values 1, 5, 10, 50, 100, 500, and 1000
	mymap["I"] = 1
	mymap["V"] = 5
	mymap["X"] = 10
	mymap["L"] = 50
	mymap["C"] = 100
	mymap["D"] = 500
	mymap["M"] = 1000

	return mymap
}

func getValue(roman string, numberMap map[string]int) int {
	value, ok := numberMap[roman]
	if !ok {
		log.Fatal("No value for ", roman) //this should be an error
	}
	return value
}

func RomanToInt(input string) int {
	numberMap := InitMap() //could use global, init, etc

	length := len(input)
	sum := 0
	for i, v := range input {
		roman := string(v)
		value := getValue(roman, numberMap)
		//simple fun
		if i+1 < length {
			vNext := getValue(string(input[i+1]), numberMap)
			if vNext > value {
				//subtract
				sum -= value
			} else {
				sum += value
			}
		} else {
			sum += value
		}
	}

	fmt.Println(sum)
	return sum
}
