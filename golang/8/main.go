package main

import (
	"advent-of-code/8/common"
	"fmt"
	"strconv"
	"strings"
)

type Tuple[T int | string, K int | string] struct {
	a T
	b K
}

func unique(intSlice []Tuple[int, string]) []Tuple[int, string] {
	keys := make(map[string]bool)
	var list []Tuple[int, string]
	for _, entry := range intSlice {
		if _, value := keys[entry.b]; !value {
			keys[entry.b] = true
			list = append(list, entry)
		}
	}
	return list
}

func reverseArray[T int | string | Tuple[int, string]](arr []T) []T {
	var reversed []T

	for i, _ := range arr {
		reversed = append(reversed, arr[len(arr)-1-i])
	}

	return reversed
}

func any(array []Tuple[int, string], validator func(item Tuple[int, string]) bool) bool {
	for _, item := range array {
		result := validator(item)
		if result {
			return result
		}
	}
	return false
}

func all(array []Tuple[int, string], validator func(item Tuple[int, string]) bool) bool {
	for _, item := range array {
		result := validator(item)
		if !result {
			return false
		}
	}

	return true
}

// testing generics
func getColumn[T int | string | Tuple[int, string]](colNumber int, data [][]T) []T {

	var extractedColumn []T

	for _, item := range data {
		extractedColumn = append(extractedColumn, item[colNumber])
	}

	return extractedColumn
}

func determineVisibilityFromTheTop(data [][]Tuple[int, string], column int) []Tuple[int, string] {
	var visitedTrees []Tuple[int, string]
	var visibleTrees []Tuple[int, string]

	extractedColumn := getColumn(column, data)

	for _, item := range extractedColumn {
		if len(visitedTrees) == 0 || all(visitedTrees, func(element Tuple[int, string]) bool {
			result := item.a > element.a
			return result
		}) {
			visibleTrees = append(visibleTrees, item)
		}
		visitedTrees = append(visitedTrees, item)
	}

	return visibleTrees
}

func determineVisibilityFromTheBottom(data [][]Tuple[int, string], column int) []Tuple[int, string] {
	var visitedTrees []Tuple[int, string]
	var visibleTrees []Tuple[int, string]

	extractedColumn := reverseArray(getColumn(column, data))

	for _, item := range extractedColumn {
		if len(visitedTrees) == 0 || all(visitedTrees, func(element Tuple[int, string]) bool {
			result := item.a > element.a
			return result
		}) {
			visibleTrees = append(visibleTrees, item)
		}
		visitedTrees = append(visitedTrees, item)
	}

	return visibleTrees
}

func determineVisibilityFromTheLeft(data [][]Tuple[int, string], row int) []Tuple[int, string] {
	var visitedTrees []Tuple[int, string]
	var visibleTrees []Tuple[int, string]

	extractedRow := data[row]

	for _, item := range extractedRow {
		if len(visitedTrees) == 0 || all(visitedTrees, func(element Tuple[int, string]) bool {
			result := item.a > element.a
			return result
		}) {
			visibleTrees = append(visibleTrees, item)
		}
		visitedTrees = append(visitedTrees, item)
	}

	return visibleTrees
}

func determineVisibilityFromTheRight(data [][]Tuple[int, string], row int) []Tuple[int, string] {
	var visitedTrees []Tuple[int, string]
	var visibleTrees []Tuple[int, string]

	extractedRow := reverseArray(data[row])

	for _, item := range extractedRow {
		if len(visitedTrees) == 0 || all(visitedTrees, func(element Tuple[int, string]) bool {
			result := item.a > element.a
			return result
		}) {
			visibleTrees = append(visibleTrees, item)
		}
		visitedTrees = append(visitedTrees, item)
	}

	return visibleTrees
}

func handleCountVisibleTrees(data [][]Tuple[int, string]) []Tuple[int, string] {

	var visibleTrees []Tuple[int, string]

	for i, _ := range data {
		visibleFromRight := determineVisibilityFromTheRight(data, i)
		visibleFromLeft := determineVisibilityFromTheLeft(data, i)

		visibleTrees = append(visibleTrees, visibleFromRight...)
		visibleTrees = append(visibleTrees, visibleFromLeft...)
	}

	for i := 0; i <= len(data[0])-1; i++ {
		visibleFromTop := determineVisibilityFromTheTop(data, i)
		visibleFromBottom := determineVisibilityFromTheBottom(data, i)

		visibleTrees = append(visibleTrees, visibleFromTop...)
		visibleTrees = append(visibleTrees, visibleFromBottom...)
	}

	return visibleTrees
}

func goUp(data [][]Tuple[int, string], row int, col int) {
	column := getColumn(col, data)
	var visibleTrees []Tuple[int, string]
	for i := 0; i <= len(column); i++ {
		if (column[row-i].a > column[row].a) {
			visibleTrees
		}

	}
}

func goDown(data [][]Tuple[int, string], row int, col int) {

}

func goLeft(data [][]Tuple[int, string], row int, col int) {

}

func goRight(data [][]Tuple[int, string], row int, col int) {

}

func main() {
	scanner := common.GetScanner("./8/test.txt")
	var data [][]Tuple[int, string]

	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		var convertedAndKeyed []Tuple[int, string]

		for i, item := range strings.Split(line, "") {
			convertedItem, _ := strconv.Atoi(item)
			key := "c" + strconv.Itoa(i) + "r" + strconv.Itoa(lineIndex)
			tuple := &Tuple[int, string]{convertedItem, key}
			convertedAndKeyed = append(convertedAndKeyed, *tuple)
		}

		data = append(data, convertedAndKeyed)
		lineIndex += 1
	}

	fmt.Println(handleCountVisibleTrees(data))
	fmt.Println(unique(handleCountVisibleTrees(data)))
	fmt.Println(len(unique(handleCountVisibleTrees(data))))
}
