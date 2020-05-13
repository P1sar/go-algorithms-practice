package graphs

import (
	"bufio"
	"io"
	"strings"
)

func NewUndirectedGraphFromAdjacencyList(r io.Reader) *UndirectedGraph {
	scanner := bufio.NewScanner(r)
	var lines [][]string
	for scanner.Scan() {
		arrS := scanner.Text()
		arrS = strings.TrimSpace(arrS)
		arr := strings.Split(arrS,"	")
		lines = append(lines, arr)
	}
	adjacentList := make(map[string][]string)
	vertList := make([]string, len(lines))
	for i, line := range lines {
		adjacentVertices := line[1:]
		vertList[i] = line[0]
		adjacentList[line[0]] = adjacentVertices
	}
	return &UndirectedGraph{
		adjacentList: adjacentList,
		vertList: vertList,
	}
}


type UndirectedGraph struct {
	adjacentList map[string][]string // AdjacentList
	vertList []string // list of all vertices
	//edges ?
}

func (g *UndirectedGraph) RandomizedMinimumCuts() {

}

func (g *UndirectedGraph) VerticesContraction(v1, v2 string) string {
	if !g.Adjacent(v1, v2) {
		return ""
	}
	newVertexName := v1+v2
	// concating all edges of contracted vertices
	edges := make([]string,0)
	edges = append(g.adjacentList[v1], g.adjacentList[v2]...)

	// removing self loops
	edges = removeElementFromArr(edges, v1)
	edges = removeElementFromArr(edges, v2)

	// updating all old vertices to newName
	for _, v := range edges {
		for i, vert := range g.adjacentList[v] {
			if vert == v1 || vert == v2 {
				g.adjacentList[v][i] = newVertexName
			}
		}
	}

	delete(g.adjacentList, v1)
	delete(g.adjacentList, v2)
	g.adjacentList[newVertexName] = edges

	g.vertList = removeElementFromArr(g.vertList, v1)
	g.vertList = removeElementFromArr(g.vertList, v2)

	g.vertList = append(g.vertList, newVertexName)
	return newVertexName
}


//Does two vertices connected (adjacent). Does in O(n) time
func (g *UndirectedGraph) Adjacent(v1, v2 string) bool {
	for _, v := range g.adjacentList[v1] {
		if v == v2 {
			return true
		}
	}
	return false
}

func removeElementFromArr(a []string, e string) []string {
	b := a[:0]
	for _, x := range a {
		if x!=e {
			b = append(b, x)
		}
	}
	return b
}
