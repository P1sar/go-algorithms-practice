package graphs

import (
	"math/rand"
	"time"
)

func randomizedMinimumCuts(graph *UndirectedGraph) int {
	for len(graph.vertList) > 2 {
		v1 := selectRandomVertex(graph.vertList)
		v2 := selectRandomVertex(graph.vertList)
		graph.VerticesContraction(string(v1), string(v2))
	}
	return len(graph.adjacentList[graph.vertList[0]])

}

func selectRandomVertex(vertices []string) string {
	rand.Seed(time.Now().UnixNano())
	return vertices[rand.Intn(len(vertices))]
}