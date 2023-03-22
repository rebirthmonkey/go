package main

import (
	"fmt"

	"github.com/rebirthmonkey/go/pkg/workflow"
	"github.com/rebirthmonkey/go/pkg/workflow/activity"
)

func main() {

	template := workflow.NewTemplate()

	e1 := workflow.NewEvent("Start")
	e2 := workflow.NewEvent("Doing")
	e3 := workflow.NewEvent("End")
	template.AddVertex(e1)
	template.AddVertex(e2)
	template.AddVertex(e3)

	template.AddEdge(e1, e2)

	a1 := activity.NewBasic("a1")
	a2 := activity.NewShell("a2", "/tmp/test.sh", "aaa", "bbb")
	a3 := activity.NewGolang("a3", func(s interface{}) interface{} {
		fmt.Println("------------")
		return s.(string)
	}, "ccc")
	template.AddVertex(a1)
	template.AddVertex(a2)
	template.AddVertex(a3)

	template.AddEdge(e2, a1)
	template.AddEdge(a1, a2)
	template.AddEdge(a1, a3)

	g1 := workflow.NewGateway("g1")
	template.AddVertex(g1)

	template.AddEdge(a2, g1)
	template.AddEdge(a3, g1)
	template.AddEdge(g1, e3)

	engine := workflow.NewEngine(template)
	err := engine.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
