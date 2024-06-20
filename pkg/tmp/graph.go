package tmp

// import (
// 	"fmt"
// 	"heterflow/pkg/logger"
// )

// func AllFunctionName(g Graph) {
// 	rel := g.Relation()
// 	for name := range rel {
// 		fmt.Println(name)
// 	}
// }

// func FuncCalltimes(g Graph) {
// 	rel := g.Relation()
// 	for name, node := range rel {
// 		fmt.Println(name, node.CalledTimes)
// 	}
// }
// func buildGraphFromRoot(g Graph, root string) map[string]*Node {
// 	Relations := g.Relation()
// 	if fuc, ok := Relations[root]; ok {
// 		for calleeName := range fuc.Callee {
// 			// fmt.Printf("%s ", calleeName)
// 			if v := fuc.Callee[calleeName]; v == nil {
// 				fuc.Callee[calleeName] = Relations[calleeName]
// 				buildGraphFromRoot(g, calleeName)
// 			}
// 		}
// 		return map[string]*Node{
// 			fuc.FuncName: fuc,
// 		}
// 	} else {
// 		logger.Error("_" + root + "_" + " is not found ")
// 		return nil
// 	}
// }

// import (
// 	"fmt"
// 	"heterflow/pkg/codeaid/util"
// 	"heterflow/pkg/logger"
// )

// const (
// 	MaxDegreeLocal  PivotSelection = iota
// 	MaxDegreeLocalX                = iota
// )

// type PivotSelection int

// type VertexSet map[Vertex]struct{}
// type VertexSets util.VertexSet[Vertex, struct{}]

// type ProductGraph struct {
// 	adjacencies map[Vertex]util.VertexSet[Vertex, struct{}]
// }

// var (
// 	NumOfTopK = 40
// )

// func (pg ProductGraph) Edged(vex Vertex) util.VertexSet[Vertex, struct{}] {
// 	return pg.adjacencies[vex]
// }

// func (pg ProductGraph) Size() int {
// 	return len(pg.adjacencies)
// }

// func (pg ProductGraph) AllConnectedVertices() util.VertexSet[Vertex, struct{}] {
// 	result := make(util.VertexSet[Vertex, struct{}], 100)
// 	for v, neighbours := range pg.adjacencies {
// 		if !util.IsEmpty(neighbours) {
// 			util.Add(result, Vertex(v))
// 		}
// 		// if !neighbours.IsEmpty() {
// 		// 	result.Add(Vertex(v))
// 		// }
// 	}
// 	return result
// }

// func (pg ProductGraph) Degree(vex Vertex) int {
// 	return len(pg.adjacencies[vex])
// }

// // func (v VertexSet) Pop() (Vertex, error) {
// // 	for k := range v {
// // 		v.Remove(k)
// // 		return k, nil
// // 	}
// // 	return Vertex{}, errors.New("vertex set is struct{}, can't pop a vertex")
// // }

// // func (v VertexSet) IsEmpty() bool {
// // 	return len(v) == 0
// // }

// // func (v VertexSet) IsDisjoint(input VertexSet) bool {
// // 	small, large := v, input
// // 	if len(small) > len(large) {
// // 		small, large = input, v
// // 	}
// // 	for v := range small {
// // 		if large.Contains(v) {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }

// // func (v VertexSet) Contains(vex Vertex) bool {
// // 	_, ok := v[vex]
// // 	return ok
// // }

// // func (v VertexSet) Remove(vex Vertex) {
// // 	delete(v, vex)
// // }

// // func (v VertexSet) Intersection(other VertexSet) VertexSet {
// // 	small, large := v, other
// // 	if len(small) > len(large) {
// // 		small, large = other, v
// // 	}
// // 	result := make(VertexSet, len(small))
// // 	for v := range small {
// // 		if large.Contains(v) {
// // 			result.Add(v)
// // 		}
// // 	}
// // 	return result
// // }

// // func (v VertexSet) Difference(other VertexSet) VertexSet {
// // 	result := make(VertexSet, len(v))
// // 	for val := range v {
// // 		if !other.Contains(val) {
// // 			result.Add(val)
// // 		}
// // 	}
// // 	return result
// // }

// // func (v VertexSet) Union(another VertexSet) VertexSet {
// // 	result := make(VertexSet, len(v))
// // 	for val := range v {
// // 		result.Add(val)
// // 	}
// // 	for val := range another {
// // 		result.Add(val)
// // 	}
// // 	return result
// // }

// // func (v VertexSet) Add(vex Vertex) {
// // 	v[vex] = struct{}{}
// // }

// // func (v VertexSet) IntersectionLen(other VertexSet) int {
// // 	small, large := v, other
// // 	if len(small) > len(large) {
// // 		small, large = other, v
// // 	}
// // 	result := 0
// // 	for v := range small {
// // 		if large.Contains(v) {
// // 			result++
// // 		}
// // 	}
// // 	return result
// // }

// // func (v VertexSet) Pick() (Vertex, error) {
// // 	for v := range v {
// // 		return v, nil
// // 	}
// // 	return Vertex{}, errors.New("attempt to pick from struct{} set")
// // }

// func topCallTimes(rel map[string]*Node) map[Features]struct{} {
// 	times := NumOfTopK
// 	if len(rel) < NumOfTopK {
// 		times = len(rel)
// 	}
// 	return TopK(rel, func(feature Features) int {
// 		if feat, ok := feature.(*Node); ok {
// 			return feat.CalledTimes
// 		}
// 		return 0
// 	}, times)
// }
// func topInstructions(rel map[string]*Node) map[Features]struct{} {
// 	times := NumOfTopK
// 	if len(rel) < NumOfTopK {
// 		times = len(rel)
// 	}
// 	return TopK(rel, func(feature Features) int {
// 		if feat, ok := feature.(*Node); ok {
// 			return feat.TotalInstruction
// 		}
// 		return 0
// 	}, times)
// }
// func union(v1, v2 map[Features]struct{}) map[Features]struct{} {
// 	result := make(map[Features]struct{}, len(v1))
// 	for val := range v1 {
// 		result[val] = struct{}{}
// 	}
// 	for val := range v2 {
// 		result[val] = struct{}{}

// 	}
// 	return result
// }
// func NewProductGraph(g1, g2 Graph) ProductGraph {
// 	res := ProductGraph{make(map[Vertex]util.VertexSet[Vertex, struct{}])}
// 	if graph1, ok := g1.(*UndirectedGraph); ok {
// 		graph1.appendLinkInfo()
// 	}
// 	if graph2, ok := g2.(*UndirectedGraph); ok {
// 		graph2.appendLinkInfo()
// 	}
// 	rel1 := g1.Relation()
// 	rel2 := g2.Relation()
// 	// 得到排名前k的函数

// 	topcall1 := topCallTimes(rel1)

// 	topcall2 := topCallTimes(rel2)

// 	// topinstruct1 := topInstructions(rel1)
// 	// topinstruct2 := topInstructions(rel2)
// 	// total1 := union(topcall1, topinstruct1)
// 	// total2 := union(topcall2, topinstruct2)
// 	total1 := topcall1
// 	total2 := topcall2

// 	//分析之间的相关性系数
// 	table := adjacentTable{
// 		leftPref:  map[string]map[string]float64{},
// 		rightPref: map[string]map[string]float64{},
// 		virtual:   map[string]struct{}{},
// 	}
// 	table.init(total1, total2)

// 	// 运行匈牙利算法
// 	fmt.Printf("left length is : %d  right length is : %d\n", len(total1), len(total2))
// 	matcher := NewMatcher(GaleShapley, &table)
// 	fmt.Printf("leftPref length is : %d  rightPref length is : %d\n", len(table.leftPref), len(table.rightPref))
// 	result := matcher.Match()
// 	matched := matcher.Result()
// 	// 输出最大匹配数及匹配情况
// 	fmt.Printf("Max Matching: %f match num is %d \n", result, len(matched))
// 	// for v := range matcher.Right() {
// 	// 	if md, ok := matched[v]; ok {
// 	// 		fmt.Printf("Vertex %s in Left is matched with Vertex %s in Right\n", md, v)
// 	// 	}
// 	// }
// 	// for v := range topcall1 {
// 	// 	if ff, ok := v.(*Node); ok {
// 	// 		fmt.Println(ff.FuncName, ff.CalledTimes, ff.TotalInstruction)
// 	// 	}
// 	// }
// 	// for v := range topcall2 {
// 	// 	if ff, ok := v.(*Node); ok {
// 	// 		fmt.Println(ff.FuncName, ff.CalledTimes)
// 	// 	}
// 	// }
// 	// fmt.Println("-----------------------------------------------")
// 	// for _, v := range rel1 {
// 	// 	fmt.Println(v.FuncName, v.CalledTimes)
// 	// }
// 	// fmt.Println("-----------------------------------------------")
// 	// for _, v := range rel2 {
// 	// 	fmt.Println(v.FuncName, v.CalledTimes)
// 	// }
// 	// fmt.Println("-----------------------------------------------")

// 	for right1, left1 := range matched {
// 		for right2, left2 := range matched {
// 			if right1 == right2 && left1 == left2 {
// 				continue
// 			}
// 			// if _, kk := rel2[right1]; !kk {
// 			// 	fmt.Println(table.swaped, right1, " is not exist in rel2")
// 			// } else {
// 			// 	fmt.Println(table.swaped, right1, "exist in rel2")
// 			// }
// 			// if _, kk := rel1[left1]; !kk {
// 			// 	fmt.Println(table.swaped, left1, " is not exist in rel1")
// 			// } else {
// 			// 	fmt.Println(table.swaped, left1, "exist in rel1")
// 			// }
// 			_, ok1 := rel2[right1].Callee[right2]
// 			_, ok2 := rel1[left1].Callee[left2]
// 			if ok1 == ok2 {
// 				if val, ok := res.adjacencies[Vertex{right1, left1}]; ok {
// 					val[Vertex{right2, left2}] = struct{}{}
// 				} else {
// 					res.adjacencies[Vertex{right1, left1}] = util.VertexSet[Vertex, struct{}]{Vertex{right2, left2}: struct{}{}}
// 				}
// 				if val, ok := res.adjacencies[Vertex{right2, left2}]; ok {
// 					val[Vertex{right1, left1}] = struct{}{}
// 				} else {
// 					res.adjacencies[Vertex{right2, left2}] = util.VertexSet[Vertex, struct{}]{Vertex{right1, left1}: struct{}{}}
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("Product graph construct complete ")
// 	// for k, v := range res.adjacencies {
// 	// 	fmt.Println(k, "    ", v)
// 	// }
// 	// fmt.Println("Product graph ------------------------------------")
// 	return res
// }

// func BronKerbosch(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
// 	if util.IsEmpty(p) && util.IsEmpty(x) {
// 		fmt.Println(r)
// 		reporter.Record(r)
// 	}
// 	for !util.IsEmpty(p) {
// 		v, err := util.Pop(p)
// 		if err != nil {
// 			logger.Fatal(err.Error())
// 		}
// 		neighbors := pg.Edged(v)
// 		BronKerbosch(pg, reporter, append(r, v), util.Intersection(p, neighbors), util.Intersection(x, neighbors))

// 		util.Add(x, v)
// 	}
// }

// func BronKerboschPivot(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
// 	if util.IsEmpty(p) && util.IsEmpty(x) {
// 		fmt.Println(r)
// 		reporter.Record(r)
// 	}
// 	u, err := util.Pop(util.Union(p, x))

// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}
// 	candidates := util.Difference(p, pg.Edged(u))
// 	for v := range candidates {
// 		neighbors := pg.Edged(v)
// 		BronKerboschPivot(pg, reporter, append(r, v), util.Intersection(p, neighbors), util.Intersection(x, neighbors))
// 		util.Remove(p, u)
// 		util.Add(x, v)
// 	}
// }

// func BronKerbosch2(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
// 	if util.IsEmpty(p) && util.IsEmpty(x) {
// 		fmt.Println(r)
// 		reporter.Record(r)
// 	}
// 	for !util.IsEmpty(p) {
// 		v, err := util.Pop(p)
// 		if err != nil {
// 			logger.Fatal(err.Error())
// 		}
// 		// fmt.Println("to find clique which include ", v)
// 		neighbors := pg.Edged(v)
// 		candidate := util.Intersection(p, neighbors)
// 		for node := range candidate {
// 			if pg.Degree(node) < len(r) {
// 				util.Remove(candidate, node)
// 				// candidate.Remove(node)
// 			}
// 		}
// 		BronKerbosch2(pg, reporter, append(r, v), candidate, util.Intersection(x, neighbors))
// 		util.Add(x, v)

// 	}
// }

// func BronKerbosch2aGP(graph *ProductGraph, reporter Reporter) {
// 	// Bron-Kerbosch algorithm with pivot of highest degree within remaining candidates
// 	// chosen from candidates only (IK_GP).
// 	candidates := graph.AllConnectedVertices()
// 	if util.IsEmpty(candidates) {
// 		return
// 	}
// 	excluded := make(util.VertexSet[Vertex, struct{}], len(candidates))
// 	visit(graph, reporter, MaxDegreeLocal, nil, candidates, excluded)
// }

// func visit(graph *ProductGraph, reporter Reporter, pivotSelection PivotSelection, clique []Vertex, candidates, excluded util.VertexSet[Vertex, struct{}]) {
// 	if len(candidates) == 1 {
// 		for v := range candidates {
// 			// Same logic as below, stripped down
// 			neighbours := graph.Edged(v)
// 			if util.IsDisjoint(excluded, neighbours) {
// 				fmt.Println(append(clique, v))
// 				reporter.Record(append(clique, v))
// 			}
// 		}
// 		return
// 	}
// 	var pivot Vertex
// 	remainingCandidates := make([]Vertex, 0, len(candidates))
// 	// Quickly handle locally unconnected candidates while finding pivot
// 	seenLocalDegree := 0
// 	for v := range candidates {
// 		neighbours := graph.Edged(v)

// 		localDegree := util.IntersectionLen(neighbours, candidates)
// 		if localDegree == 0 {
// 			// Same logic as below, stripped down
// 			if util.IsDisjoint(neighbours, excluded) {
// 				fmt.Println(append(clique, v))
// 				reporter.Record(append(clique, v))
// 			}
// 		} else {
// 			if seenLocalDegree < localDegree {
// 				seenLocalDegree = localDegree
// 				pivot = v
// 			}
// 			remainingCandidates = append(remainingCandidates, v)
// 		}
// 	}
// 	if seenLocalDegree == 0 {
// 		return
// 	}
// 	if pivotSelection == MaxDegreeLocalX {
// 		for v := range excluded {
// 			neighbours := graph.Edged(v)
// 			localDegree := util.IntersectionLen(neighbours, candidates)

// 			if seenLocalDegree < localDegree {
// 				seenLocalDegree = localDegree
// 				pivot = v
// 			}
// 		}
// 	}

// 	for _, v := range remainingCandidates {
// 		neighbours := graph.Edged(v)
// 		if util.Contains(neighbours, pivot) {
// 			continue
// 		}
// 		util.Remove(candidates, v)
// 		// candidates.Remove(v)

// 		neighbouringCandidates := util.Intersection(neighbours, candidates)
// 		if !util.IsEmpty(neighbouringCandidates) {
// 			neighbouringExcluded := util.Intersection(neighbours, excluded)
// 			visit(
// 				graph, reporter,
// 				pivotSelection,
// 				append(clique, v),
// 				neighbouringCandidates,
// 				neighbouringExcluded)
// 		} else {
// 			if util.IsDisjoint(neighbours, excluded) {
// 				fmt.Println(append(clique, v))
// 				reporter.Record(append(clique, v))
// 			}
// 		}
// 		util.Add(excluded, v)
// 		// excluded.Add(v)
// 	}
// }

// // package main

// // import "fmt"

// // // Finds index of string in slice.
// // func indexOf(s []string, val string) int {
// //     for i, v := range s {
// //         if v == val {
// //             return i
// //         }
// //     }
// //     return -1
// // }

// // // GaleShapley algorithm to find stable matches.
// // func GaleShapley(menPreferences map[string][]string, womenPreferences map[string][]string) (matches map[string]string) {
// //     // Initialize men and free men list.
// //     matches = make(map[string]string)
// //     freeMen := make([]string, 0, len(menPreferences))
// //     for man := range menPreferences {
// //         freeMen = append(freeMen, man)
// //     }

// //     // While there are free men who still have a woman to propose to.
// //     for len(freeMen) > 0 {
// //         man := freeMen[0] // Get the first free man.
// //         manPrefs := menPreferences[man]

// //         for _, woman := range manPrefs {
// //             if partner, ok := matches[woman]; !ok {
// //                 // If the woman is not yet engaged, engage her with the man.
// //                 matches[woman] = man
// //                 freeMen = freeMen[1:]
// //                 break
// //             } else if indexOf(womenPreferences[woman], man) < indexOf(womenPreferences[woman], partner) {
// //                 // If the woman prefers this man to her current partner.
// //                 matches[woman] = man
// //                 freeMen = freeMen[1:]
// //                 freeMen = append(freeMen, partner) // The previous partner is now free.
// //                 break
// //             }
// //         }
// //     }

// //     // Reverse matches to show matches from men's perspective.
// //     reversedMatches := make(map[string]string)
// //     for woman, man := range matches {
// //         reversedMatches[man] = woman
// //     }

// //     return reversedMatches
// // }

// import (
// 	"fmt"
// 	"heterflow/pkg/codeaid/util"
// 	"math"
// 	"strconv"
// )

// const Inf = 1<<63 - 1

// type MatcherType int

// const (
// 	Hungarian MatcherType = iota
// 	KM
// 	GaleShapley
// 	threshold = 0.95
// )

// type Matcher interface {
// 	Match() float64
// 	Result() map[string]string
// 	Left() map[string]map[string]float64
// 	Right() map[string]map[string]float64
// }

// func NewMatcher(class MatcherType, adj *adjacentTable) Matcher {
// 	switch class {
// 	case Hungarian:
// 		return &hungarian{
// 			adjacentTable: adj,
// 			match:         map[string]string{},
// 			visited:       map[string]struct{}{},
// 		}
// 	case KM:
// 		return &kM{
// 			adjacentTable: adj,
// 			visitedleft:   make(map[string]bool, len(adj.leftPref)),
// 			visitedright:  make(map[string]bool, len(adj.rightPref)),
// 			weightleft:    make(map[string]float64, len(adj.leftPref)),
// 			weightright:   make(map[string]float64, len(adj.rightPref)),
// 			delta:         make(map[string]float64, len(adj.rightPref)),
// 			match:         make(map[string]string, len(adj.rightPref)),
// 		}
// 	case GaleShapley:
// 		return &galeshapley{
// 			adjacentTable: adj,
// 			match:         make(map[string]string, len(adj.rightPref)),
// 		}
// 	}
// 	return nil
// }

// type hungarian struct {
// 	*adjacentTable
// 	match   map[string]string
// 	visited map[string]struct{}
// 	// left, right map[Features]struct{}
// }
// type galeshapley struct {
// 	*adjacentTable
// 	match map[string]string
// 	// left, right map[Features]struct{}
// }
// type kM struct {
// 	*adjacentTable
// 	visitedleft  map[string]bool
// 	visitedright map[string]bool
// 	weightleft   map[string]float64
// 	weightright  map[string]float64
// 	delta        map[string]float64
// 	// left, right  map[Features]struct{}
// 	match map[string]string
// }

// type adjacentTable struct {
// 	leftPref  map[string]map[string]float64
// 	rightPref map[string]map[string]float64
// 	virtual   map[string]struct{}
// 	swaped    bool
// }

// func (adj *adjacentTable) init(left, right map[Features]struct{}) {
// 	for v1 := range left {
// 		for v2 := range right {
// 			similarity, err := util.Pearson(v1.Features(), v2.Features())
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			// fmt.Println("similarity  ", v1.Name(), v2.Name(), similarity)
// 			if similarity > threshold {
// 				if value, ok := adj.leftPref[v1.Name()]; ok {
// 					value[v2.Name()] = similarity
// 					adj.leftPref[v1.Name()] = value
// 				} else {
// 					adj.leftPref[v1.Name()] = map[string]float64{v2.Name(): similarity}
// 				}
// 				if value, ok := adj.rightPref[v2.Name()]; ok {
// 					value[v1.Name()] = similarity
// 					adj.rightPref[v2.Name()] = value
// 				} else {
// 					adj.rightPref[v2.Name()] = map[string]float64{v1.Name(): similarity}
// 				}
// 			}
// 		}
// 	}
// }

// // Find 尝试为顶点 u 找到匹配
// func (table *hungarian) find(u string) bool {
// 	for v := range table.leftPref[u] {
// 		if _, ok := table.visited[v]; !ok {
// 			table.visited[v] = struct{}{}
// 			if table.match[v] == "" || table.find(table.match[v]) {
// 				table.match[v] = u
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func (table *hungarian) Left() map[string]map[string]float64 {
// 	return table.leftPref
// }
// func (table *hungarian) Right() map[string]map[string]float64 {
// 	return table.rightPref
// }

// // hungarian 返回二部图的最大匹配数
// func (table *hungarian) Match() float64 {
// 	var result int
// 	for u := range table.leftPref {
// 		table.visited = make(map[string]struct{})
// 		if table.find(u) {
// 			result++
// 		}
// 	}
// 	return float64(result)
// }

// func (table *hungarian) Result() map[string]string {
// 	return table.match
// }

// func (km *kM) prepare() {
// 	if len(km.leftPref) == len(km.rightPref) {
// 		return
// 	}
// 	if len(km.leftPref) < len(km.rightPref) {
// 		km.leftPref, km.rightPref = km.rightPref, km.leftPref
// 		km.swaped = true
// 	}
// 	different := len(km.leftPref) - len(km.rightPref)
// 	for i := 0; i < different; i++ {
// 		str := strconv.Itoa(i)
// 		mm := make(map[string]float64, len(km.leftPref))
// 		for left, lp := range km.leftPref {
// 			lp[str] = float64(0.000000000000000001)
// 			mm[left] = float64(0.000000000000000001)
// 		}
// 		km.rightPref[str] = mm
// 		km.virtual[str] = struct{}{}
// 	}
// 	for left, lp := range km.leftPref {
// 		for right, rp := range km.leftPref {
// 			if _, ok := lp[right]; !ok {
// 				lp[right] = math.Inf(-1)
// 				rp[left] = math.Inf(-1)
// 			}
// 		}
// 		km.leftPref[left] = lp
// 	}
// 	// if km.swaped {
// 	// 	km.leftPref, km.rightPref = km.rightPref, km.leftPref
// 	// }
// }

// func (km *kM) init() {
// 	km.prepare()
// 	for k := range km.leftPref {
// 		km.visitedleft[k] = false
// 		km.weightleft[k] = math.Inf(-1)
// 	}
// 	for k := range km.weightleft {
// 		neighbor := km.leftPref[k]
// 		maximam := math.Inf(-1)
// 		for _, val := range neighbor {
// 			maximam = max(maximam, val)
// 		}
// 		km.weightleft[k] = maximam
// 	}
// 	for k := range km.rightPref {
// 		km.visitedright[k] = false
// 		km.weightright[k] = 0
// 	}
// }

// func (km *kM) find(x string) bool {
// 	km.visitedleft[x] = true
// 	// fmt.Println(x, "'s adjacent is ", km.leftPref[x])
// 	for k := range km.leftPref[x] {
// 		// fmt.Println("judge ", k, " weight is ", km.leftPref[x][k])
// 		if v := km.visitedright[k]; !v {
// 			// 0.000000000000000003
// 			fmt.Println("value is  ", x, km.weightleft[x], k, km.weightright[k], km.leftPref[x][k], " km.weightleft[x]+km.weightright[k]-km.leftPref[x][k] is ", km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
// 			if km.weightleft[x]+km.weightright[k]-km.leftPref[x][k] == 0 {
// 				km.visitedright[k] = true
// 				if km.match[k] == "" || km.find(km.match[k]) {
// 					// fmt.Println(k, "is match to ", x)
// 					km.match[k] = x
// 					return true
// 				}
// 			} else {
// 				// fmt.Println("set delte ", k, " ", km.weightleft[x], km.weightright[k], km.leftPref[x][k], km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
// 				km.delta[k] = min(km.delta[k], km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
// 			}
// 		}
// 	}
// 	return false
// }

// func (km *kM) Match() float64 {
// 	// for left, val := range km.adj {
// 	// 	for right, weight := range val {
// 	// 		fmt.Println(left, " between", right, " is ", weight)
// 	// 	}
// 	// }
// 	km.init()
// 	fmt.Printf("leftPref length is : %d  rightPref length is : %d\n", len(km.leftPref), len(km.rightPref))
// 	for x := range km.leftPref {
// 		fmt.Println("to find ", x, " pair")
// 		for {
// 			km.reset()
// 			if km.find(x) {
// 				break
// 			}
// 			mindelta := math.Inf(1)
// 			for k, v := range km.visitedright {
// 				// fmt.Println("visited right ", k, "  mindelta ", km.delta[k])
// 				if !v {
// 					mindelta = min(mindelta, km.delta[k])
// 				}
// 			}
// 			for k := range km.leftPref {
// 				if km.visitedleft[k] {
// 					km.weightleft[k] -= mindelta
// 				}
// 			}
// 			for k := range km.rightPref {
// 				if km.visitedright[k] {
// 					km.weightright[k] += mindelta
// 				}
// 			}
// 			// fmt.Println("mindelta ", mindelta)
// 		}
// 	}
// 	res := 0.0
// 	newmatcher := make(map[string]string, len(km.match))
// 	for right, left := range km.match {
// 		if _, ok := km.virtual[right]; ok {
// 			continue
// 		}
// 		if _, ok := km.virtual[left]; ok {
// 			continue
// 		}
// 		if km.swaped {
// 			newmatcher[left] = right
// 		} else {
// 			newmatcher[right] = left
// 		}
// 	}
// 	km.match = newmatcher
// 	for k, v := range km.match {
// 		res += km.leftPref[k][v]
// 		fmt.Println("res+=", "pair:", v, k, km.leftPref[k][v])
// 	}
// 	return res
// }

// func (km *kM) reset() {
// 	for k := range km.visitedleft {
// 		km.visitedleft[k] = false
// 	}
// 	for k := range km.visitedright {
// 		km.visitedright[k] = false
// 	}
// 	for k := range km.delta {
// 		km.delta[k] = math.Inf(1)
// 	}
// }
// func (km *kM) Left() map[string]map[string]float64 {
// 	return km.leftPref
// }
// func (km *kM) Right() map[string]map[string]float64 {
// 	return km.rightPref
// }

// func (km *kM) Result() map[string]string {
// 	return km.match
// }

// func (gs *galeshapley) Left() map[string]map[string]float64 {
// 	return gs.leftPref
// }

// func (gs *galeshapley) Right() map[string]map[string]float64 {
// 	return gs.rightPref
// }

// func (gs *galeshapley) Result() map[string]string {
// 	return gs.match
// }

// func (gs *galeshapley) prepare() {
// 	if len(gs.leftPref) == len(gs.rightPref) {
// 		return
// 	}
// 	if len(gs.leftPref) < len(gs.rightPref) {
// 		gs.leftPref, gs.rightPref = gs.rightPref, gs.leftPref
// 		gs.swaped = true
// 	}
// 	different := len(gs.leftPref) - len(gs.rightPref)
// 	for _, v := range gs.leftPref {
// 		for right := range gs.rightPref {
// 			if _, ok := v[right]; !ok {
// 				v[right] = math.Inf(-1)
// 			}
// 		}
// 		for i := 0; i < different; i++ {
// 			v[strconv.Itoa(i)] = math.Inf(-1)
// 		}
// 	}
// 	for _, v := range gs.rightPref {
// 		for left := range gs.leftPref {
// 			if _, ok := v[left]; !ok {
// 				v[left] = math.Inf(-1)
// 			}
// 		}
// 	}
// 	for i := 0; i < different; i++ {
// 		str := strconv.Itoa(i)
// 		m := make(map[string]float64, len(gs.leftPref))
// 		for left := range gs.leftPref {
// 			m[left] = math.Inf(-1)
// 		}
// 		gs.rightPref[str] = m
// 		gs.virtual[str] = struct{}{}
// 	}
// }

// func (gs *galeshapley) Match() float64 {
// 	gs.prepare()
// 	freeLeft := make([]string, 0, len(gs.leftPref))
// 	for left := range gs.leftPref {
// 		freeLeft = append(freeLeft, left)
// 	}
// 	proposals := make(map[string]string, len(gs.leftPref))
// 	res := 0.0
// 	for len(freeLeft) > 0 {
// 		curLeft := freeLeft[0]
// 		// fmt.Println("match for ", curLeft)
// 		lp := gs.leftPref[curLeft]
// 		for right, v := range lp {
// 			currentPartner, ok1 := proposals[right]
// 			if !ok1 {
// 				// fmt.Println("match", right, "to", curLeft)
// 				proposals[right] = curLeft
// 				res += v
// 				freeLeft = freeLeft[1:]
// 				break
// 			} else {
// 				rp := gs.rightPref[right]
// 				if rp[curLeft] > rp[currentPartner] {
// 					// fmt.Println("dismatch", right, "with", currentPartner)
// 					// fmt.Println("match", right, "to", curLeft)
// 					res += rp[curLeft]
// 					res -= rp[currentPartner]
// 					proposals[right] = curLeft
// 					freeLeft = freeLeft[1:]
// 					freeLeft = append(freeLeft, currentPartner)
// 					break
// 				}
// 			}
// 		}
// 	}
// 	for right, left := range proposals {
// 		if _, ok := gs.virtual[right]; ok {
// 			continue
// 		}
// 		if _, ok := gs.virtual[left]; ok {
// 			continue
// 		}
// 		if gs.swaped {
// 			gs.match[left] = right
// 		} else {
// 			gs.match[right] = left
// 		}
// 	}
// 	return res
// }

// func CFGSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
// 	pg := graph.NewProductGraph(config1.Graph, config2.Graph)
// 	candidates := pg.AllConnectedVertices()
// 	if util.IsEmpty(candidates) {
// 		return 0.0, fmt.Errorf("no connected vertics")
// 	}
// 	excluded := make(util.VertexSet[graph.Vertex, struct{}], len(candidates))
// 	reporter := graph.NewRepoter(graph.MaxReporter)
// 	// graph.BronKerbosch2aGP(pg, reporter)
// 	// graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
// 	graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
// 	res2 := reporter.Report(pg.Size())
// 	return res2, nil
// }

// root := c.build.GraphGraphFromRoot("main@@Base-0x50")
// c.build.AllFunctionName()
// c.build.BFS(root)
// roots := c.Graph.BuildGraph()
// filefeature.ControlFlowGraphRoots = roots
// graph.FuncCalltimes(c.Graph)
// for _, node := range roots {
// 	graph.BFS(node.FuncFeatures)
// 	fmt.Println()
// 	fmt.Println(`***************************************************`)
// 	fmt.Println()
// }
