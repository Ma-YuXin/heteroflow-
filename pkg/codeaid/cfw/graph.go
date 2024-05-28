package cfw

type Vertex Pair[string, string]

type Pair[K, V any] struct {
	Key   K
	Value V
}

type Graph interface {
	// BuildGraph() map[string]*Node
	Relation() map[string]*Node
	AddNode(*Node)
	NodeNum() int
	SetNodeOutDegree()
}

type DirectedGraph struct {
	Relations map[string]*Node
}

type UndirectedGraph struct {
	DirectedGraph
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Relations: make(map[string]*Node),
	}
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		DirectedGraph: DirectedGraph{
			Relations: make(map[string]*Node),
		},
	}
}

func addNode(g Graph, node *Node) {
	if node.Name() == "" {
		return
	}
	rel := g.Relation()
	if v, ok := rel[node.Name()]; ok {
		if v.flag {
			node.CalledTimes = v.CalledTimes
			rel[node.Name()] = node
		}
		// } else {
		// 	fmt.Println("readd function : " + node.Name())
		// }
	} else {
		node.CalledTimes = 1
		rel[node.Name()] = node
	}
	for name := range node.Callee {
		if v, ok := rel[name]; ok {
			v.CalledTimes += 1
		} else {
			node := NewNode()
			node.CalledTimes = 1
			node.flag = true
			rel[name] = node
		}
	}
}

func (g *DirectedGraph) AddNode(node *Node) {
	addNode(g, node)
}

func (g *DirectedGraph) NodeNum() int {
	return len(g.Relations)
}

func (g *DirectedGraph) Relation() map[string]*Node {
	return g.Relations
}

func (g *DirectedGraph) SetNodeOutDegree() {
	rel := g.Relations
	for _, node := range rel {
		node.OutDegree = len(node.Callee)
	}
}

func (g *UndirectedGraph) appendLinkInfo() {
	for name, node := range g.Relations {
		for callee := range node.Callee {
			g.Relations[callee].AddCallee(name)
		}
	}
}

// func (g *DirectedGraph) BuildGraph() map[string]*Node {
// 	return buildGraph(g)
// }

// func (g *UndirectedGraph) AddNode(node *Node) {
// 	addNode(g, node)
// }

// func (g *UndirectedGraph) BuildGraph() map[string]*Node {
// 	g.appendLinkInfo()
// 	return buildGraph(g)
// }

// func (g *UndirectedGraph) Relation() map[string]*Node {
// 	return g.Relations
// }

// func (g *UndirectedGraph) NodeNum() int {
// 	return len(g.Relations)
// }

// func topK(m map[string]*Node, featureSelect func(*Node) int, k int) map[*Node]struct{} {
// 	if len(m) < k {
// 		topk := make(map[*Node]struct{}, k)
// 		for _, value := range m {
// 			topk[value] = struct{}{}
// 		}
// 		return topk
// 	}
// 	return findTopKValues(m, featureSelect, k)
// }

// func findTopKValues(m map[string]*Node, featureSelect func(*Node) int, k int) map[*Node]struct{} {
// 	type kv struct {
// 		Key   string
// 		Value int
// 	}

// 	var ss []kv
// 	for k, v := range m {
// 		ss = append(ss, kv{k, featureSelect(v)})
// 	}

// 	sort.Slice(ss, func(i, j int) bool {
// 		return ss[i].Value > ss[j].Value // 降序排序
// 	})

// 	topK := make(map[*Node]struct{}, k)
// 	for i := 0; i < k; i++ {
// 		topK[m[ss[i].Key]] = struct{}{}
// 	}

// 	return topK
// }
// func bFS(root *Node) {
// 	// for k, v := range b.Relations {
// 	// 	fmt.Printf("%s %+v \n", k, v.Callee)
// 	// }
// 	// fmt.Println()
// 	history := map[*Node]struct{}{root: {}}
// 	queue := []*Node{root}
// 	var last *Node = root
// 	for len(queue) != 0 {
// 		cur := queue[0]
// 		if cur != nil {
// 			fmt.Printf("%s  ", cur.Name())
// 			for _, v := range cur.Neighbors() {
// 				if _, ok := history[v]; !ok {
// 					queue = append(queue, v)
// 					history[v] = struct{}{}
// 				}
// 			}
// 		}
// 		if cur == last {
// 			fmt.Println("\n--------------------------------------------------")
// 			last = queue[len(queue)-1]
// 		}
// 		queue = queue[1:]
// 	}
// }

// func buildGraph(g Graph) map[string]*Node {
// 	children := make(map[string]bool)
// 	relations := g.Relation()
// 	for _, parent := range relations {
// 		for childName := range parent.Callee {
// 			if _, exits := relations[childName]; !exits {
// 				logger.Error("child " + childName + " is not found")
// 			}
// 			parent.Callee[childName] = relations[childName]
// 			if childName != parent.FuncName {
// 				children[childName] = true
// 			}
// 		}
// 	}
// 	roots := make(map[string]*Node)
// 	for name, node := range relations {
// 		if !children[name] {
// 			roots[name] = node
// 		}
// 	}
// 	return roots
// }
