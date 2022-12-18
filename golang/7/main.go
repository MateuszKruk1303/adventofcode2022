package main

import (
	"advent-of-code/7/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	BACK = "$ cd .."
	LIST = "$ ls"
)

type Directory struct {
	Parent  *Directory
	name    string
	content []string
}

var visitedNodes []*Directory
var totalSumOfContent int64

func filterAndSum() {

	for _, node := range visitedNodes {
		entryContent := node.content
		var sumOfContent int64 = 0
		fmt.Println("=======================================================================")
		fmt.Println(node.name)
		for _, item := range entryContent {
			parsedItem := strings.Split(item, " ")[0]
			if regexp.MustCompile(`\d`).MatchString(parsedItem) {
				fmt.Println(parsedItem)
				converted, _ := strconv.ParseInt(parsedItem, 0, 64)
				sumOfContent += converted
			}
		}
		fmt.Println("=======================================================================")
		if sumOfContent <= 100000 {
			fmt.Println("SUM")
			fmt.Println(sumOfContent)
			fmt.Println(totalSumOfContent)
			fmt.Println(sumOfContent + totalSumOfContent)
			totalSumOfContent = sumOfContent + totalSumOfContent
		}
		fmt.Println("=======================================================================")

		sumOfContent = 0
	}
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func processFilesAndDirectories(sample string, currentNode *Directory) {
	currentNode.content = unique(append(currentNode.content, sample))
}

func executeCommand(command string, currentNode *Directory) *Directory {
	switch command {
	case BACK:
		return currentNode.Parent
	case LIST:
		return currentNode
	default:
		dirname := command[5:len(command)]
		var newNode = &Directory{Parent: currentNode, name: dirname}
		visitedNodes = append(visitedNodes, newNode)
		return newNode
	}

	return currentNode
}

func processDataSample(sample string, currentNode *Directory) *Directory {
	isCommand := strings.Contains(sample, "$")
	if isCommand {
		return executeCommand(sample, currentNode)
	} else {
		processFilesAndDirectories(sample, currentNode)
		return currentNode
	}
}

func main() {
	scanner := common.GetScanner("./7/test.txt")

	var currentNode = &Directory{Parent: nil, name: "/"}
	visitedNodes = append(visitedNodes, currentNode)

	var data []string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for _, item := range data {
		currentNode = processDataSample(item, currentNode)
	}

	filterAndSum()
	fmt.Println(totalSumOfContent)

}
