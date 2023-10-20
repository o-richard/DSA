```go
import "github.com/o-richard/DSA/floydwarshall"
                        .
                        .
                        .
graph := make([][]uint, 4)
graph[0] = []uint{0, 3, math.MaxInt, 5}
graph[1] = []uint{2, 0, math.MaxInt, 4}
graph[2] = []uint{math.MaxInt, 1, 0, math.MaxInt}
graph[3] = []uint{math.MaxInt, math.MaxInt,2, 0}
fmt.Println(floydwarshall.FloydWarshall(graph, 4))
```
