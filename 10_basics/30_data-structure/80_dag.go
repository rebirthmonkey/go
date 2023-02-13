package main

import (
	"fmt"

	"github.com/goombaio/dag"
)

func main() {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)
	vertex4 := dag.NewVertex("4", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddVertex(vertex2)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddVertex(vertex3)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddVertex(vertex4)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	// Edges
	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddEdge(vertex2, vertex4)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddEdge(vertex3, vertex4)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	fmt.Println(dag1.String())

	v4, _ := dag1.GetVertex("4")
	fmt.Println("the vertex4 is: ", v4)

	vx, _ := dag1.Predecessors(v4)
	for _, pv := range vx {
		fmt.Println("the prodecessor of vertex4 is: ", pv)
	}

	v2, _ := dag1.GetVertex("2")
	vx, _ = dag1.Successors(v2)
	for _, sv := range vx {
		fmt.Println("the successor of vertex2 is: ", sv)
	}

	sinkVertexes := dag1.SinkVertices()
	for _, sinkv := range sinkVertexes {
		fmt.Println("the sink vertex is: ", sinkv)
	}

	sourceVertexes := dag1.SourceVertices()
	for _, sourcev := range sourceVertexes {
		fmt.Println("the source vertex is: ", sourcev)
	}

}
