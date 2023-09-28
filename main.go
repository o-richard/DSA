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
	var commandType, directory, fileName string
	flag.StringVar(&commandType, "command", "generate", "Choices: generate or restore. It specifies whether to generate a template for the specified file path or to restore the previous solution. Default choice is generate.")
	flag.StringVar(&directory, "directory", "", "(Required). A list of comma separed directories to perform the specified command. Examples bfs or bfs,dfs,prim. A shorthand notation in case of the whole project is to use a period (.) The shorthand notation for the whole project has a higher priority")
	flag.StringVar(&fileName, "file", "", "(Optional). The name of the target file in case of a single directory. Examples bfs.go, dfs.go")
	flag.Parse()

	if commandType != "restore" && commandType != "generate" {
		flag.Usage()
		os.Exit(1)
	}

	directories := strings.Split(directory, ",")
	if len(directories) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if slices.Index(directories, ".") != -1 {
		directories = make([]string, 0, len(packageTemplates))
		for k := range packageTemplates {
			directories = append(directories, k)
		}
	} else {
		for _, dir := range directories {
			if _, ok := packageTemplates[dir]; !ok {
				fmt.Println("could not locate the directory: ", dir)
				os.Exit(1)
			}
		}
	}

	if fileName != "" && len(directories) > 1 {
		flag.Usage()
		os.Exit(1)
	}
	if fileName != "" {
		tempName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		if _, ok := packageTemplates[directories[0]][tempName]; !ok {
			fmt.Println("could not locate the file ", fileName, "in the directory ", directories[0])
			os.Exit(1)
		}
		fileName = tempName
	}

	if commandType == "generate" {
		if fileName != "" {
			writeToFile(directories[0], fileName, packageTemplates[directories[0]][fileName])
		} else {
			generateTemplates(directories)
		}
	} else {
		if fileName != "" {
			fetchFileContent(directories[0], fileName)
		} else {
			restoreSolutions(directories)
		}
	}
}
