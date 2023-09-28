package kruskal

var Templates = map[string]string{
	"kruskal": `package kruskal

type Edge struct {
	Source, Destination string
	Weight int
}

type graph struct {
	// Edit here
}

func InitGraph() graph {
	// Edit here
}

func (g *graph) AddEdge(source, destination string, weight int) {
	// Edit here
}

// Returns the edges in the minimum spanning tree and the minimum cost
func (g *graph) Kruskal() ([]Edge, int) {
	// Edit here
}
	`,
}
