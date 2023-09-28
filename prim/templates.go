package prim

var Templates = map[string]string{
	"prim": `package prim

type Edge struct {
	Source, Destination string
	Weight int
}

func Prim(graph map[string]map[string]int) ([]Edge, int) {
	// Edit here
}	
	`,
}