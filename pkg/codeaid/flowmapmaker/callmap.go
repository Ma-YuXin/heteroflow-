package flowmapmaker

import (
	"fmt"
	"heterflow/pkg/logger"
)

var ()

type Build struct {
	Relations map[string]*FuncFeatures
}

func (b *Build) AddNode(ff *FuncFeatures) {
	b.Relations[ff.FuncName] = ff
}
func (b *Build) BuildGraph() map[string]*FuncFeatures {
	children := make(map[string]bool)
	for _, parent := range b.Relations {
		for childName := range parent.Callee {
			if _, exits := b.Relations[childName]; !exits {
				logger.Error("child " + childName + " is not found")
			}
			parent.Callee[childName] = b.Relations[childName]
			if childName != parent.FuncName {
				children[childName] = true
			}
		}
	}
	roots := make(map[string]*FuncFeatures)
	for name, node := range b.Relations {
		if !children[name] {
			roots[name] = node
		}
	}
	return roots
}
func (b *Build) BuildGraphFromRoot(root string) *FuncFeatures {
	if fuc, ok := b.Relations[root]; ok {
		for calleeName := range fuc.Callee {
			// fmt.Printf("%s ", calleeName)
			if v := fuc.Callee[calleeName]; v == nil {
				fuc.Callee[calleeName] = b.Relations[calleeName]
				b.BuildGraphFromRoot(calleeName)
			}
		}
		return fuc
	} else {
		logger.Error("_" + root + "_" + " is not found ")
		return nil
	}
}
func (b *Build) AllFunctionName() {
	for name := range b.Relations {
		fmt.Println(name)
	}
}

func BFS(root *FuncFeatures) {
	// for k, v := range b.Relations {
	// 	fmt.Printf("%s %+v \n", k, v.Callee)
	// }
	// fmt.Println()
	history := map[*FuncFeatures]struct{}{root: {}}
	queue := []*FuncFeatures{root}
	last := root
	for len(queue) != 0 {
		cur := queue[0]
		if cur != nil {
			fmt.Printf("%s  ", cur.FuncName)
			for _, v := range cur.Callee {
				if _, ok := history[v]; !ok {
					queue = append(queue, v)
					history[v] = struct{}{}
				}
			}
		}
		if cur == last {
			fmt.Println("\n--------------------------------------------------")
			last = queue[len(queue)-1]
		}
		queue = queue[1:]
	}
}
