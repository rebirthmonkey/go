package workflow

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/goombaio/dag"
)

type Engine struct {
	template *Template
	phase    string
	entries  *queue.Queue
	entrySet *set.Set
}

func NewEngine(t *Template) *Engine {

	return &Engine{
		template: t,
		phase:    InitPhase,
		entries:  queue.New(),
		entrySet: set.New(),
	}
}

func (e *Engine) Execute() error {
	switch e.phase {
	case InitPhase:
		fmt.Println("================ InitPhase")

		sourceVertexes := e.template.DAG.SourceVertices()
		for _, sourcev := range sourceVertexes {
			fmt.Println("the source vertex is: ", sourcev.ID)
			e.entrySet.Insert(sourcev)
			e.entries.Enqueue(sourcev)
		}

		e.phase = ExecPhase
		e.Execute()
	case ExecPhase:
		fmt.Println("================ ExecPhase")

		for e.entries.Len() != 0 {
			val := e.entries.Dequeue()
			e.entrySet.Remove(val)

			e.template.vertex2Object[val.(*dag.Vertex)].Execute()

			vx, _ := e.template.DAG.Successors(val.(*dag.Vertex))
			for _, pv := range vx {
				if e.entrySet.Has(pv) == false {
					e.entries.Enqueue(pv)
					e.entrySet.Insert(pv)
				}
			}
		}

		e.phase = FinalPhase
		e.Execute()
	case FinalPhase:
		fmt.Println("================ FinalPhase")
	}
	return nil
}
