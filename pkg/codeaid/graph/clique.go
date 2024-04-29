package graph

import (
	"fmt"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
)

const (
	MaxDegreeLocal  PivotSelection = iota
	MaxDegreeLocalX                = iota
)

type PivotSelection int
type Pair struct {
	Key   string
	Value string
}
type Vertex Pair

// type VertexSet map[Vertex]struct{}
// type VertexSets util.VertexSet[Vertex, struct{}]

type ProductGraph struct {
	adjacencies map[Vertex]util.VertexSet[Vertex, struct{}]
}

var (
	NumOfTopK = 40
)

func (pg ProductGraph) Edged(vex Vertex) util.VertexSet[Vertex, struct{}] {
	return pg.adjacencies[vex]
}

func (pg ProductGraph) Size() int {
	return len(pg.adjacencies)
}

func (pg ProductGraph) AllConnectedVertices() util.VertexSet[Vertex, struct{}] {
	result := make(util.VertexSet[Vertex, struct{}], 100)
	for v, neighbours := range pg.adjacencies {
		if !util.IsEmpty(neighbours) {
			util.Add(result, Vertex(v))
		}
		// if !neighbours.IsEmpty() {
		// 	result.Add(Vertex(v))
		// }
	}
	return result
}

func (pg ProductGraph) Degree(vex Vertex) int {
	return len(pg.adjacencies[vex])
}

// func (v VertexSet) Pop() (Vertex, error) {
// 	for k := range v {
// 		v.Remove(k)
// 		return k, nil
// 	}
// 	return Vertex{}, errors.New("vertex set is struct{}, can't pop a vertex")
// }

// func (v VertexSet) IsEmpty() bool {
// 	return len(v) == 0
// }

// func (v VertexSet) IsDisjoint(input VertexSet) bool {
// 	small, large := v, input
// 	if len(small) > len(large) {
// 		small, large = input, v
// 	}
// 	for v := range small {
// 		if large.Contains(v) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (v VertexSet) Contains(vex Vertex) bool {
// 	_, ok := v[vex]
// 	return ok
// }

// func (v VertexSet) Remove(vex Vertex) {
// 	delete(v, vex)
// }

// func (v VertexSet) Intersection(other VertexSet) VertexSet {
// 	small, large := v, other
// 	if len(small) > len(large) {
// 		small, large = other, v
// 	}
// 	result := make(VertexSet, len(small))
// 	for v := range small {
// 		if large.Contains(v) {
// 			result.Add(v)
// 		}
// 	}
// 	return result
// }

// func (v VertexSet) Difference(other VertexSet) VertexSet {
// 	result := make(VertexSet, len(v))
// 	for val := range v {
// 		if !other.Contains(val) {
// 			result.Add(val)
// 		}
// 	}
// 	return result
// }

// func (v VertexSet) Union(another VertexSet) VertexSet {
// 	result := make(VertexSet, len(v))
// 	for val := range v {
// 		result.Add(val)
// 	}
// 	for val := range another {
// 		result.Add(val)
// 	}
// 	return result
// }

// func (v VertexSet) Add(vex Vertex) {
// 	v[vex] = struct{}{}
// }

// func (v VertexSet) IntersectionLen(other VertexSet) int {
// 	small, large := v, other
// 	if len(small) > len(large) {
// 		small, large = other, v
// 	}
// 	result := 0
// 	for v := range small {
// 		if large.Contains(v) {
// 			result++
// 		}
// 	}
// 	return result
// }

// func (v VertexSet) Pick() (Vertex, error) {
// 	for v := range v {
// 		return v, nil
// 	}
// 	return Vertex{}, errors.New("attempt to pick from struct{} set")
// }

func topCallTimes(rel map[string]*Node) map[Features]struct{} {
	times := NumOfTopK
	if len(rel) < NumOfTopK {
		times = len(rel)
	}
	return TopK(rel, func(feature Features) int {
		if feat, ok := feature.(*Node); ok {
			return feat.CalledTimes
		}
		return 0
	}, times)
}
func topInstructions(rel map[string]*Node) map[Features]struct{} {
	times := NumOfTopK
	if len(rel) < NumOfTopK {
		times = len(rel)
	}
	return TopK(rel, func(feature Features) int {
		if feat, ok := feature.(*Node); ok {
			return feat.TotalInstruction
		}
		return 0
	}, times)
}
func union(v1, v2 map[Features]struct{}) map[Features]struct{} {
	result := make(map[Features]struct{}, len(v1))
	for val := range v1 {
		result[val] = struct{}{}
	}
	for val := range v2 {
		result[val] = struct{}{}

	}
	return result
}
func NewProductGraph(g1, g2 Graph) ProductGraph {
	res := ProductGraph{make(map[Vertex]util.VertexSet[Vertex, struct{}])}
	if graph1, ok := g1.(*UndirectedGraph); ok {
		graph1.appendLinkInfo()
	}
	if graph2, ok := g2.(*UndirectedGraph); ok {
		graph2.appendLinkInfo()
	}
	rel1 := g1.Relation()
	rel2 := g2.Relation()
	// 得到排名前k的函数

	topcall1 := topCallTimes(rel1)

	topcall2 := topCallTimes(rel2)

	// topinstruct1 := topInstructions(rel1)
	// topinstruct2 := topInstructions(rel2)
	// total1 := union(topcall1, topinstruct1)
	// total2 := union(topcall2, topinstruct2)
	total1 := topcall1
	total2 := topcall2

	//分析之间的相关性系数
	table := adjacentTable{
		leftPref:  map[string]map[string]float64{},
		rightPref: map[string]map[string]float64{},
		virtual:   map[string]struct{}{},
	}
	table.init(total1, total2)

	// 运行匈牙利算法
	fmt.Printf("left length is : %d  right length is : %d\n", len(total1), len(total2))
	matcher := NewMatcher(GaleShapley, &table)
	fmt.Printf("leftPref length is : %d  rightPref length is : %d\n", len(table.leftPref), len(table.rightPref))
	result := matcher.Match()
	matched := matcher.Result()
	// 输出最大匹配数及匹配情况
	fmt.Printf("Max Matching: %f match num is %d \n", result, len(matched))
	// for v := range matcher.Right() {
	// 	if md, ok := matched[v]; ok {
	// 		fmt.Printf("Vertex %s in Left is matched with Vertex %s in Right\n", md, v)
	// 	}
	// }
	// for v := range topcall1 {
	// 	if ff, ok := v.(*Node); ok {
	// 		fmt.Println(ff.FuncName, ff.CalledTimes, ff.TotalInstruction)
	// 	}
	// }
	// for v := range topcall2 {
	// 	if ff, ok := v.(*Node); ok {
	// 		fmt.Println(ff.FuncName, ff.CalledTimes)
	// 	}
	// }
	// fmt.Println("-----------------------------------------------")
	// for _, v := range rel1 {
	// 	fmt.Println(v.FuncName, v.CalledTimes)
	// }
	// fmt.Println("-----------------------------------------------")
	// for _, v := range rel2 {
	// 	fmt.Println(v.FuncName, v.CalledTimes)
	// }
	// fmt.Println("-----------------------------------------------")

	for right1, left1 := range matched {
		for right2, left2 := range matched {
			if right1 == right2 && left1 == left2 {
				continue
			}
			// if _, kk := rel2[right1]; !kk {
			// 	fmt.Println(table.swaped, right1, " is not exist in rel2")
			// } else {
			// 	fmt.Println(table.swaped, right1, "exist in rel2")
			// }
			// if _, kk := rel1[left1]; !kk {
			// 	fmt.Println(table.swaped, left1, " is not exist in rel1")
			// } else {
			// 	fmt.Println(table.swaped, left1, "exist in rel1")
			// }
			_, ok1 := rel2[right1].Callee[right2]
			_, ok2 := rel1[left1].Callee[left2]
			if ok1 == ok2 {
				if val, ok := res.adjacencies[Vertex{right1, left1}]; ok {
					val[Vertex{right2, left2}] = struct{}{}
				} else {
					res.adjacencies[Vertex{right1, left1}] = util.VertexSet[Vertex, struct{}]{Vertex{right2, left2}: struct{}{}}
				}
				if val, ok := res.adjacencies[Vertex{right2, left2}]; ok {
					val[Vertex{right1, left1}] = struct{}{}
				} else {
					res.adjacencies[Vertex{right2, left2}] = util.VertexSet[Vertex, struct{}]{Vertex{right1, left1}: struct{}{}}
				}
			}
		}
	}
	fmt.Println("Product graph construct complete ")
	// for k, v := range res.adjacencies {
	// 	fmt.Println(k, "    ", v)
	// }
	// fmt.Println("Product graph ------------------------------------")
	return res
}

func BronKerbosch(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
	if util.IsEmpty(p) && util.IsEmpty(x) {
		fmt.Println(r)
		reporter.Record(r)
	}
	for !util.IsEmpty(p) {
		v, err := util.Pop(p)
		if err != nil {
			logger.Fatal(err.Error())
		}
		neighbors := pg.Edged(v)
		BronKerbosch(pg, reporter, append(r, v), util.Intersection(p, neighbors), util.Intersection(x, neighbors))

		util.Add(x, v)
	}
}

func BronKerboschPivot(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
	if util.IsEmpty(p) && util.IsEmpty(x) {
		fmt.Println(r)
		reporter.Record(r)
	}
	u, err := util.Pop(util.Union(p, x))

	if err != nil {
		logger.Fatal(err.Error())
	}
	candidates := util.Difference(p, pg.Edged(u))
	for v := range candidates {
		neighbors := pg.Edged(v)
		BronKerboschPivot(pg, reporter, append(r, v), util.Intersection(p, neighbors), util.Intersection(x, neighbors))
		util.Remove(p, u)
		util.Add(x, v)
	}
}

func BronKerbosch2(pg ProductGraph, reporter Reporter, r []Vertex, p, x util.VertexSet[Vertex, struct{}]) {
	if util.IsEmpty(p) && util.IsEmpty(x) {
		fmt.Println(r)
		reporter.Record(r)
	}
	for !util.IsEmpty(p) {
		v, err := util.Pop(p)
		if err != nil {
			logger.Fatal(err.Error())
		}
		// fmt.Println("to find clique which include ", v)
		neighbors := pg.Edged(v)
		candidate := util.Intersection(p, neighbors)
		for node := range candidate {
			if pg.Degree(node) < len(r) {
				util.Remove(candidate, node)
				// candidate.Remove(node)
			}
		}
		BronKerbosch2(pg, reporter, append(r, v), candidate, util.Intersection(x, neighbors))
		util.Add(x, v)

	}
}

func BronKerbosch2aGP(graph *ProductGraph, reporter Reporter) {
	// Bron-Kerbosch algorithm with pivot of highest degree within remaining candidates
	// chosen from candidates only (IK_GP).
	candidates := graph.AllConnectedVertices()
	if util.IsEmpty(candidates) {
		return
	}
	excluded := make(util.VertexSet[Vertex, struct{}], len(candidates))
	visit(graph, reporter, MaxDegreeLocal, nil, candidates, excluded)
}

func visit(graph *ProductGraph, reporter Reporter, pivotSelection PivotSelection, clique []Vertex, candidates, excluded util.VertexSet[Vertex, struct{}]) {
	if len(candidates) == 1 {
		for v := range candidates {
			// Same logic as below, stripped down
			neighbours := graph.Edged(v)
			if util.IsDisjoint(excluded, neighbours) {
				fmt.Println(append(clique, v))
				reporter.Record(append(clique, v))
			}
		}
		return
	}
	var pivot Vertex
	remainingCandidates := make([]Vertex, 0, len(candidates))
	// Quickly handle locally unconnected candidates while finding pivot
	seenLocalDegree := 0
	for v := range candidates {
		neighbours := graph.Edged(v)

		localDegree := util.IntersectionLen(neighbours, candidates)
		if localDegree == 0 {
			// Same logic as below, stripped down
			if util.IsDisjoint(neighbours, excluded) {
				fmt.Println(append(clique, v))
				reporter.Record(append(clique, v))
			}
		} else {
			if seenLocalDegree < localDegree {
				seenLocalDegree = localDegree
				pivot = v
			}
			remainingCandidates = append(remainingCandidates, v)
		}
	}
	if seenLocalDegree == 0 {
		return
	}
	if pivotSelection == MaxDegreeLocalX {
		for v := range excluded {
			neighbours := graph.Edged(v)
			localDegree := util.IntersectionLen(neighbours, candidates)

			if seenLocalDegree < localDegree {
				seenLocalDegree = localDegree
				pivot = v
			}
		}
	}

	for _, v := range remainingCandidates {
		neighbours := graph.Edged(v)
		if util.Contains(neighbours, pivot) {
			continue
		}
		util.Remove(candidates, v)
		// candidates.Remove(v)

		neighbouringCandidates := util.Intersection(neighbours, candidates)
		if !util.IsEmpty(neighbouringCandidates) {
			neighbouringExcluded := util.Intersection(neighbours, excluded)
			visit(
				graph, reporter,
				pivotSelection,
				append(clique, v),
				neighbouringCandidates,
				neighbouringExcluded)
		} else {
			if util.IsDisjoint(neighbours, excluded) {
				fmt.Println(append(clique, v))
				reporter.Record(append(clique, v))
			}
		}
		util.Add(excluded, v)
		// excluded.Add(v)
	}
}

// package main

// import "fmt"

// // Finds index of string in slice.
// func indexOf(s []string, val string) int {
//     for i, v := range s {
//         if v == val {
//             return i
//         }
//     }
//     return -1
// }

// // GaleShapley algorithm to find stable matches.
// func GaleShapley(menPreferences map[string][]string, womenPreferences map[string][]string) (matches map[string]string) {
//     // Initialize men and free men list.
//     matches = make(map[string]string)
//     freeMen := make([]string, 0, len(menPreferences))
//     for man := range menPreferences {
//         freeMen = append(freeMen, man)
//     }

//     // While there are free men who still have a woman to propose to.
//     for len(freeMen) > 0 {
//         man := freeMen[0] // Get the first free man.
//         manPrefs := menPreferences[man]

//         for _, woman := range manPrefs {
//             if partner, ok := matches[woman]; !ok {
//                 // If the woman is not yet engaged, engage her with the man.
//                 matches[woman] = man
//                 freeMen = freeMen[1:]
//                 break
//             } else if indexOf(womenPreferences[woman], man) < indexOf(womenPreferences[woman], partner) {
//                 // If the woman prefers this man to her current partner.
//                 matches[woman] = man
//                 freeMen = freeMen[1:]
//                 freeMen = append(freeMen, partner) // The previous partner is now free.
//                 break
//             }
//         }
//     }

//     // Reverse matches to show matches from men's perspective.
//     reversedMatches := make(map[string]string)
//     for woman, man := range matches {
//         reversedMatches[man] = woman
//     }

//     return reversedMatches
// }
