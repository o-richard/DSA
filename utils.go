package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/o-richard/DSA/bellmanford"
	"github.com/o-richard/DSA/bfs"
	"github.com/o-richard/DSA/dfs"
	"github.com/o-richard/DSA/dijkstra"
	"github.com/o-richard/DSA/floydwarshall"
	"github.com/o-richard/DSA/fordfulkerson"
	"github.com/o-richard/DSA/heap"
	"github.com/o-richard/DSA/huffman"
	"github.com/o-richard/DSA/kosaraju"
	"github.com/o-richard/DSA/kruskal"
	"github.com/o-richard/DSA/linkedlist"
	"github.com/o-richard/DSA/longestcommonsubsequence"
	"github.com/o-richard/DSA/prim"
	"github.com/o-richard/DSA/queue"
	"github.com/o-richard/DSA/robinkarp"
	"github.com/o-richard/DSA/search"
	"github.com/o-richard/DSA/sort"
	"github.com/o-richard/DSA/stack"
	"github.com/o-richard/DSA/tree"
)

var packageTemplates = map[string]map[string]string{
	"bellmanford":              bellmanford.Templates,
	"bfs":                      bfs.Templates,
	"dfs":                      dfs.Templates,
	"dijkstra":                 dijkstra.Templates,
	"floydwarshall":            floydwarshall.Templates,
	"fordfulkerson":            fordfulkerson.Templates,
	"heap":                     heap.Templates,
	"huffman":                  huffman.Templates,
	"kosaraju":                 kosaraju.Templates,
	"kruskal":                  kruskal.Templates,
	"linkedlist":               linkedlist.Templates,
	"longestcommonsubsequence": longestcommonsubsequence.Templates,
	"prim":                     prim.Templates,
	"queue":                    queue.Templates,
	"robinkarp":                robinkarp.Templates,
	"search":                   search.Templates,
	"sort":                     sort.Templates,
	"stack":                    stack.Templates,
	"tree":                     tree.Templates,
}

func writeToFile(dir, fileName, content string) {
	filePath := fmt.Sprintf("%v/%v.go", dir, fileName)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("there was an error in writing to the file path %v", filePath)
	}
}

// Generate the templates for the specified directories.
func generateTemplates(dirs []string) {
	for _, dir := range dirs {
		templates := packageTemplates[dir]
		for fileName, content := range templates {
			writeToFile(dir, fileName, content)
		}
	}
}

// Fetch the previous solutions from github for restoration
func fetchFileContent(dir, fileName string) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/o-richard/DSA/main/%v/%v.go", dir, fileName)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not fetch the solution of directory ", dir, "and file ", fileName, " from Github")
		return
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	writeToFile(dir, fileName, string(content))
}

// Restore the solutions for the specified directories.
func restoreSolutions(dirs []string) {
	for _, dir := range dirs {
		templates := packageTemplates[dir]
		for fileName := range templates {
			fetchFileContent(dir, fileName)
		}
	}
}
