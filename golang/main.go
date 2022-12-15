package main

import (
	"./common"
	"fmt"
)

const (
	ASCII_UPPERCASE_SUBSTRACT = 38
	ASCII_LOWERCASE_SUBSTRACT = 96
	ASCII_UPPERCASE_START     = 65
	ASCII_LOWERCASE_START     = 97
)

func getCommonItemBetweenCompartments(firstCompartment string, secondCompartment string) (commonItem int32) {
	for _, fItem := range firstCompartment {
		for _, sItem := range secondCompartment {
			if sItem == fItem {
				commonItem = fItem
				break
			}
		}
	}

	return
}

func getItemWeight(letterAsciiCode int32) (weight int32) {
	if letterAsciiCode >= ASCII_LOWERCASE_START {
		weight = letterAsciiCode - ASCII_LOWERCASE_SUBSTRACT
	} else {
		weight = letterAsciiCode - ASCII_UPPERCASE_SUBSTRACT
	}

	return
}

func main() {
	var sumOfWeights int32 = 0
	var data []string
	scanner := common.GetScanner("./input3.txt")

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for _, item := range data {
		firstCompartment, secondCompartment := item[0:(len(item)/2)], item[len(item)/2:len(item)]

		commonChar := getCommonItemBetweenCompartments(firstCompartment, secondCompartment)

		sumOfWeights = sumOfWeights + getItemWeight(commonChar)
	}

	fmt.Println(sumOfWeights)

}
