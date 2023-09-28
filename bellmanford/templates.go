package bellmanford

var Templates = map[string]string{
	"bellmanFord": `package bellmanford

type graph struct {
	// Edit Here
}

func InitGraph() graph {
	// Edit here
}

func (g *graph) AddEdge(source, destination string, weight int) {
	// Edit here
}

// Returns the destinations and the respective distances
func (g *graph) BellmanFord(source string) (map[string]int, error) {
	// Edit here
}	
	`,
}
