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

//
//func (g *UndirectedGraph) VerticesContraction123(v1, v2 string) string {
//	fmt.Printf("Contracting %v %v\n", v1, v2)
//	if !g.Adjacent(v1, v2) {
//		fmt.Printf("Not adjacent %v %v! list is %v\n", v1, v2, g.vertList)
//
//		return ""
//	}
//	newVertexName := v1+v2
//	newVerticesList := g.slamVerticesAdjacentVertices(v1, v2)
//	delete(g.vertices, v1)
//	delete(g.vertices, v2)
//	g.vertices[newVertexName] = make([]string,0)
//
//	for _, v := range newVerticesList {
//		g.RemoveEdge(v, v1)
//		g.RemoveEdge(v, v2)
//		g.AddEdge(v, newVertexName)
//	}
//	for i:=0;i<len(g.vertList); i++ {
//		if g.vertList[i] == v1 {
//			// Moving last arr element on position of found element
//			g.vertList[i] = g.vertList[len(g.vertList)-1]
//			// Remove last element
//			g.vertList = g.vertList[:len(g.vertList)-1]
//		}
//	}
//	for i:=0;i<len(g.vertList); i++ {
//		if g.vertList[i] == v2 {
//			// Moving last arr element on position of found element
//			g.vertList[i] = g.vertList[len(g.vertList)-1]
//			// Remove last element
//			g.vertList = g.vertList[:len(g.vertList)-1]
//		}
//	}
//
//	g.vertList = append(g.vertList, newVertexName)
//	fmt.Printf("Vertex list %v\n", g.vertList)
//	fmt.Printf("edges list %v\n", g.vertices[newVertexName])
//	return newVertexName
//}
//
//func (g *UndirectedGraph) slamVerticesAdjacentVertices(v1, v2 string) []string {
//	newAdjacentVertices := make([]string, 0)
//
//	for _, v := range g.vertices[v1] {
//		if v == v2 {
//			continue
//		}
//		newAdjacentVertices = append(newAdjacentVertices,v)
//	}
//	for _, v := range g.vertices[v2] {
//		if v == v1 {
//			continue
//		}
//		newAdjacentVertices = append(newAdjacentVertices,v)
//	}
//	return newAdjacentVertices
//}
//
//func (g *UndirectedGraph) addNewVertex(v string, adjacentVertices []string) {
//	g.vertices[v] = adjacentVertices
//}
//
//func (g *UndirectedGraph) AddEdge(v1, v2 string) {
//	if g.Adjacent(v1, v2) {
//		return
//	}
//	g.vertices[v1] = append(g.vertices[v1], v2)
//	g.vertices[v2] = append(g.vertices[v2], v1)
//}
//
//func (g *UndirectedGraph) RemoveEdge(v1, v2 string) {
//	for i, v := range g.vertices[v1] {
//		if v == v2 {
//			// Moving last arr element on position of found element
//			g.vertices[v1][i] = g.vertices[v1][len(g.vertices[v1])-1]
//			// Remove last element
//			g.vertices[v1] = g.vertices[v1][:len(g.vertices[v1])-1]
//		}
//	}
//	for i, v := range g.vertices[v2] {
//		if v == v1 {
//			// Moving last arr element on position of found element
//			g.vertices[v2][i] = g.vertices[v2][len(g.vertices[v2])-1]
//			// Remove last element
//			g.vertices[v2] = g.vertices[v2][:len(g.vertices[v2])-1]
//		}
//	}
//}


//func (g *UndirectedGraph) Degree(v2 string) int {
//	return len(g.vertices[v2])
//}

//func (g *UndirectedGraph) FindVerticesOf()
