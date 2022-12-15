package main

import (
	"./common"
	"fmt"
	"strings"
)

const (
	CD       = "$ cd"
	BACK     = "$ cd .."
	LIST     = "$ ls"
	ROOT_CMD = "$ cd /"
)

type Directory struct {
	Parent  *Directory
	name    string
	content []string
}

var visitedNodes []*Directory

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
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
	scanner := common.GetScanner("./input.txt")

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

	fmt.Println(visitedNodes)

}
