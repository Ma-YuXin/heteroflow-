package graph

import (
	"fmt"
	"heterflow/pkg/logger"
)

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
func buildGraphFromRoot(g Graph, root string) map[string]*Node {
	Relations := g.Relation()
	if fuc, ok := Relations[root]; ok {
		for calleeName := range fuc.Callee {
			// fmt.Printf("%s ", calleeName)
			if v := fuc.Callee[calleeName]; v == nil {
				fuc.Callee[calleeName] = Relations[calleeName]
				buildGraphFromRoot(g, calleeName)
			}
		}
		return map[string]*Node{
			fuc.FuncName: fuc,
		}
	} else {
		logger.Error("_" + root + "_" + " is not found ")
		return nil
	}
}

func buildGraph(g Graph) map[string]*Node {
	children := make(map[string]bool)
	Relations := g.Relation()
	for _, parent := range Relations {
		for childName := range parent.Callee {
			if _, exits := Relations[childName]; !exits {
				logger.Error("child " + childName + " is not found")
			}
			parent.Callee[childName] = Relations[childName]
			if childName != parent.FuncName {
				children[childName] = true
			}
		}
	}
	roots := make(map[string]*Node)
	for name, node := range Relations {
		if !children[name] {
			roots[name] = node
		}
	}
	return roots
}
