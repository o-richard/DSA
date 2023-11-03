package kosaraju

var Templates = map[string]string{
	"kosaraju": `
package kosaraju

type graph struct {
	// Edit here
}

func InitGraph() graph {
	// Edit here
}

func (g *graph) AddEdge(source, destination string) {
	// Edit here
}

func (g *graph) SCC() [][]string {
	// Edit here
}
	
	`,
}