package graphs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//
//func TestGraph(t *testing.T) {
//	file, err := os.Open("../test-data/kargerMinCut.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	graph := NewUndirectedGraphFromAdjacencyList(file)
//	if graph.Degree("60") != 24 {
//			t.Fail()
//	}
//
//	if len(graph.vertList) != 200 {
//		t.Fail()
//	}
//
//	if graph.Adjacent("60", "1") {
//		t.Fail()
//	}
//	if !graph.Adjacent("60", "86") {
//		t.Fail()
//	}
//
//	graph.RemoveEdge("60", "86")
//
//	if graph.Adjacent("60", "86") {
//		t.Fail()
//	}
//	if graph.Adjacent("86", "60") {
//		t.Fail()
//	}
//
//	graph.VerticesContraction("86", "60")
//
//	//if v, ok := graph.vertices["86"]; ok {
//	//	t.Fail()
//	//	fmt.Println(v)
//	//}
//	//if _, ok := graph.vertices["60"]; ok {
//	//	t.Fail()
//	//}
//	//if _, ok := graph.vertices["8660"]; !ok {
//	//	t.Fail()
//	//}
//}
//
//func TestUndirectedGraph_Adjacent(t *testing.T) {
//	file, err := os.Open("../test-data/kargerMinCut.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	graph := NewUndirectedGraphFromAdjacencyList(file)
//	if graph.Degree("60") != 24 {
//		t.Fail()
//	}
//
//
//	if !graph.Adjacent("60", "86") {
//		t.Fail()
//	}
//
//	newVertexName := graph.VerticesContraction("86", "60")
//
//	if _, ok := graph.adjacentList["86"]; ok {
//		t.Fail()
//		//fmt.Println(v)
//	}
//	if _, ok := graph.adjacentList["60"]; ok {
//		t.Fail()
//	}
//	if _, ok := graph.adjacentList["8660"]; !ok {
//		t.Fail()
//	}
//
//
//	if !graph.Adjacent(newVertexName, "112") {
//		t.Fail()
//	}
//
//	if len(graph.vertList) != 199 {
//		t.Fail()
//		fmt.Println(graph.vertList)
//	}
//}
//
func TestRandomizedMinimumCuts(t *testing.T) {
	minCut := 1 << 31

	file, err := os.Open("../test-data/kargerMinCut.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	for i:=0;i<=400; i++ {
		graph := NewUndirectedGraphFromAdjacencyList(bytes.NewReader(b))
		tmp := randomizedMinimumCuts(graph)
		if tmp <= minCut {
			minCut = tmp
		}
	}
	fmt.Println(minCut)
}


func BenchmarkRandomizedMinimumCuts(b *testing.B) {
	file, err := os.Open("../test-data/kargerMinCut.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	by, _ := ioutil.ReadAll(file)
	for n := 0; n < b.N; n++ {
		graph := NewUndirectedGraphFromAdjacencyList(bytes.NewReader(by))
		randomizedMinimumCuts(graph)
	}
	//fmt.Println(res)
}




func TestUndirectedGraph_VerticesContraction(t *testing.T) {
	file, err := os.Open("../test-data/smallGraph.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	graph := NewUndirectedGraphFromAdjacencyList(file)

	if graph.VerticesContraction("1", "5") != "15" {
		t.Error("not 15")
	}
	if len(graph.vertList) != 7 {
		t.Errorf("not 7 but %v", graph.vertList)
	}

	if graph.VerticesContraction("15", "6") != "156" {
		t.Error("Not 156")
	}

	for k,v := range graph.adjacentList {
		fmt.Println(k)
		fmt.Println(v)
	}
}