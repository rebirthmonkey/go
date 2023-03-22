package workflow

import (
	"fmt"

	"github.com/goombaio/dag"
)

type Template struct {
	DAG *dag.DAG

	object2Vertex map[Object]*dag.Vertex
	vertex2Object map[*dag.Vertex]Object
}

func NewTemplate() *Template {

	return &Template{
		DAG: dag.NewDAG(),

		object2Vertex: make(map[Object]*dag.Vertex),
		vertex2Object: make(map[*dag.Vertex]Object),
	}
}

func (t *Template) AddVertex(obj interface{}) {
	vertex := dag.NewVertex(obj.(Object).GetName(), obj)
	t.object2Vertex[obj.(Object)] = vertex
	t.vertex2Object[vertex] = obj.(Object)
	err := t.DAG.AddVertex(vertex)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Template) AddEdge(obj1, obj2 interface{}) {
	err := t.DAG.AddEdge(t.object2Vertex[obj1.(Object)], t.object2Vertex[obj2.(Object)])
	if err != nil {
		fmt.Println(err)
	}
}

type Object interface {
	GetName() string
	Execute() error
}
