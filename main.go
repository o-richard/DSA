package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	var commandType string
	flag.StringVar(&commandType, "command", "generate", "Choices: generate or restore. It specifies whether to generate a template for the specified file path or to restore the previous solution. Default choice is generate.")
	var directory string
	flag.StringVar(&directory, "directory", "", "(Required). A list of comma separed directories to perform the specified command. Examples bfs or bfs,dfs,prim. A shorthand notation in case of the whole project is to use a period (.) The shorthand notation for the whole project has a higher priority")
	var fileName string
	flag.StringVar(&fileName, "file", "", "(Optional). The name of the target file in case of a single directory. Examples bfs.go, dfs.go")

	flag.Parse()

	// Verify the command type
	if slices.Index([]string{"restore", "generate"}, commandType) == -1 {
		flag.Usage()
		os.Exit(1)
	}
	// Verify the directory
	directorySlice := strings.Split(directory, ",")
	if len(directorySlice) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if slices.Index(directorySlice, ".") != -1 {
		directorySlice = make([]string, 0, len(packageTemplates))
		for k := range packageTemplates {
			directorySlice = append(directorySlice, k)
		}
	} else {
		for _, dir := range directorySlice {
			if _, ok := packageTemplates[dir]; !ok {
				fmt.Println("could not locate the directory: ", dir)
				os.Exit(1)
			}
		}
	}
	// Verify the file
	if fileName != "" && len(directorySlice) > 1 {
		flag.Usage()
		os.Exit(1)
	}
	if fileName != "" {
		tempName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		_, ok := packageTemplates[directorySlice[0]][tempName]
		if !ok {
			fmt.Println("could not locate the file ", fileName, "in the directory ", directorySlice[0])
			os.Exit(1)
		}
		fileName = tempName
	}

	if commandType == "generate" {
		if fileName != "" {
			writeToFile(directorySlice[0], fileName, packageTemplates[directorySlice[0]][fileName])
		} else {
			generateTemplates(directorySlice)
		}
	} else {
		if fileName != "" {
			fetchFileContent(directorySlice[0], fileName)
		} else {
			restoreSolutions(directorySlice)
		}
	}
}
