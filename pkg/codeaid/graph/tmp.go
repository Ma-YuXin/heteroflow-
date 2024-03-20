package graph

import "fmt"

func AllFunctionName(g Graph) {
	rel := g.Relation()
	for name := range rel {
		fmt.Println(name)
	}
}

func FuncCalltimes(g Graph) {
	rel := g.Relation()
	for name, node := range rel {
		fmt.Println(name, node.CalledTimes)
	}
}
