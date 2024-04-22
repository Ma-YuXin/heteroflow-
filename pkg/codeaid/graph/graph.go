package graph

import (
	"heterflow/pkg/logger"
	"sort"
)

type Graph interface {
	BuildGraph() map[string]*Node
	Relation() map[string]*Node
	AddNode(*Node)
	NodeNum() int
}
type DirectedGraph struct {
	Relations map[string]*Node
}

type UndirectedGraph struct {
	Relations map[string]*Node
}

func NewGraph(class GraphType) Graph {
	switch class {
	case Directedh:
		return &DirectedGraph{
			Relations: make(map[string]*Node),
		}
	case Undirected:
		return &UndirectedGraph{
			Relations: make(map[string]*Node),
		}
	default:
		return nil
	}
}

func (g *DirectedGraph) AddNode(ff *Node) {
	addNode(g, ff)
}

func (g *DirectedGraph) BuildGraph() map[string]*Node {
	return buildGraph(g)
}

func (g *DirectedGraph) NodeNum() int {
	return len(g.Relations)
}

func (g *DirectedGraph) Relation() map[string]*Node {
	return g.Relations
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

func (g *UndirectedGraph) AddNode(ff *Node) {
	addNode(g, ff)
}

func addNode(g Graph, ff *Node) {
	rel := g.Relation()
	if v, ok := rel[ff.FuncName]; ok {
		if v.flag {
			ff.CalledTimes = v.CalledTimes
			rel[ff.FuncName] = ff
		} else {
			logger.Info("re-add function : " + ff.FuncName)
		}
	} else {
		ff.CalledTimes = 1
		rel[ff.FuncName] = ff
	}
	for name := range ff.Callee {
		if v, ok := rel[name]; ok {
			v.CalledTimes += 1
		} else {
			if nf, ok := NewFeatures(FuncFeatures).(*Node); ok {
				nf.CalledTimes = 1
				nf.flag = true
				rel[name] = nf
			}
		}
	}
}

func (g *UndirectedGraph) BuildGraph() map[string]*Node {
	g.appendLinkInfo()
	return buildGraph(g)
}

func (g *UndirectedGraph) appendLinkInfo() {
	for name, node := range g.Relations {
		for callee := range node.Callee {
			g.Relations[callee].AddCallee(name)
		}
	}
}

func (g *UndirectedGraph) Relation() map[string]*Node {
	return g.Relations
}

func (g *UndirectedGraph) NodeNum() int {
	return len(g.Relations)
}

func TopK(m map[string]*Node, featureSelect func(Features) int, k int) map[Features]empty {
	if len(m) < k {
		topk := make(map[Features]empty, k)
		for _, value := range m {
			topk[value] = empty{}
		}
		return topk
	}
	return findTopKValues(m, featureSelect, k)
}

func findTopKValues(m map[string]*Node, featureSelect func(Features) int, k int) map[Features]empty {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, featureSelect(v)})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value // 降序排序
	})

	topK := make(map[Features]empty, k)
	for i := 0; i < k; i++ {
		topK[m[ss[i].Key]] = empty{}
	}

	return topK
}
