package graph

import (
	"errors"
	"fmt"
	"heterflow/pkg/logger"
)

type Pair struct {
	Key   string
	Value string
}
type Vertex Pair

type VertexSet map[Vertex]empty

type ProductGraph struct {
	adjacencies map[Vertex]VertexSet
}
type empty struct{}

var (
	NumOfTopK = 100
)

func (pg *ProductGraph) Edged(vex Vertex) VertexSet {
	return pg.adjacencies[vex]
}

func (pg *ProductGraph) AllConnectedVertices() VertexSet {
	result := make(VertexSet)
	for v, neighbours := range pg.adjacencies {
		if !neighbours.IsEmpty() {
			result.Add(Vertex(v))
		}
	}
	return result
}

func (pg *ProductGraph) Degree(vex Vertex) int {
	return len(pg.adjacencies[vex])
}

func (v VertexSet) Pop() (Vertex, error) {
	for k := range v {
		v.Remove(k)
		return k, nil
	}
	return Vertex{}, errors.New("vertex set is empty, can't pop a vertex")
}

func (v VertexSet) IsEmpty() bool {
	return len(v) == 0
}

func (v VertexSet) IsDisjoint(input VertexSet) bool {
	small, large := v, input
	if len(small) > len(large) {
		small, large = input, v
	}
	for v := range small {
		if large.Contains(v) {
			return false
		}
	}
	return true
}

func (v VertexSet) Contains(vex Vertex) bool {
	_, ok := v[vex]
	return ok
}

func (v VertexSet) Remove(vex Vertex) {
	delete(v, vex)
}

func (v VertexSet) Intersection(other VertexSet) VertexSet {
	small, large := v, other
	if len(small) > len(large) {
		small, large = other, v
	}
	result := make(VertexSet, len(small))
	for v := range small {
		if large.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (v VertexSet) Difference(other VertexSet) VertexSet {
	result := make(VertexSet, len(v))
	for val := range v {
		if !other.Contains(val) {
			result.Add(val)
		}
	}
	return result
}

func (v VertexSet) Union(another VertexSet) VertexSet {
	result := make(VertexSet, len(v))
	for val := range v {
		result.Add(val)
	}
	for val := range another {
		result.Add(val)
	}
	return result
}

func (v VertexSet) Add(vex Vertex) {
	v[vex] = struct{}{}
}

func (v VertexSet) IntersectionLen(other VertexSet) int {
	small, large := v, other
	if len(small) > len(large) {
		small, large = other, v
	}
	result := 0
	for v := range small {
		if large.Contains(v) {
			result++
		}
	}
	return result
}

func (v VertexSet) Pick() (Vertex, error) {
	for v := range v {
		return v, nil
	}
	return Vertex{}, errors.New("attempt to pick from empty set")
}

func topCallTimes(rel map[string]*Node) map[Features]empty {
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
func topInstructions(rel map[string]*Node) map[Features]empty {
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
func union(v1, v2 map[Features]empty) map[Features]empty {
	result := make(map[Features]empty, len(v1))
	for val := range v1 {
		result[val] = empty{}
	}
	for val := range v2 {
		result[val] = empty{}

	}
	return result
}
func NewProductGraph(g1, g2 Graph) *ProductGraph {
	res := ProductGraph{make(map[Vertex]VertexSet)}
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
		virtual:   map[string]empty{},
	}
	table.init(total1, total2)

	// 运行匈牙利算法
	fmt.Printf("left length is : %d  right length is : %d\n", len(total1), len(total2))
	matcher := NewMatcher(KM, &table)
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
					val[Vertex{right2, left2}] = empty{}
				} else {
					res.adjacencies[Vertex{right1, left1}] = VertexSet{Vertex{right2, left2}: empty{}}
				}
				if val, ok := res.adjacencies[Vertex{right2, left2}]; ok {
					val[Vertex{right1, left1}] = empty{}
				} else {
					res.adjacencies[Vertex{right2, left2}] = VertexSet{Vertex{right1, left1}: empty{}}
				}
			}
		}
	}
	fmt.Println("Product graph construct complete ")
	// for k, v := range res.adjacencies {
	// 	fmt.Println(k, "    ", v)
	// }
	// fmt.Println("Product graph ------------------------------------")
	return &res
}

func BronKerbosch(pg *ProductGraph, reporter Reporter, r []Vertex, p, x VertexSet) {
	if p.IsEmpty() && x.IsEmpty() {
		fmt.Println(r)
		reporter.Record(r)
	}
	for !p.IsEmpty() {
		v, err := p.Pop()
		if err != nil {
			logger.Fatal(err.Error())
		}
		neighbors := pg.Edged(v)
		BronKerbosch(pg, reporter, append(r, v), p.Intersection(neighbors), x.Intersection(neighbors))
		x.Add(v)
	}
}

func BronKerboschPivot(pg *ProductGraph, reporter Reporter, r []Vertex, p, x VertexSet) {
	if p.IsEmpty() && x.IsEmpty() {
		fmt.Println(r)
		reporter.Record(r)
	}
	u, err := p.Union(x).Pop()
	if err != nil {
		logger.Fatal(err.Error())
	}
	candidates := p.Difference(pg.Edged(u))
	for v := range candidates {
		neighbors := pg.Edged(v)
		BronKerboschPivot(pg, reporter, append(r, v), p.Intersection(neighbors), x.Intersection(neighbors))
		p.Remove(u)
		x.Add(v)
	}
}

func BronKerbosch2(pg *ProductGraph, reporter Reporter, r []Vertex, p, x VertexSet) {
	if p.IsEmpty() && x.IsEmpty() {
		fmt.Println(r)
		reporter.Record(r)
	}
	for !p.IsEmpty() {
		v, err := p.Pop()
		if err != nil {
			logger.Fatal(err.Error())
		}
		fmt.Println("to find clique which include ", v)
		neighbors := pg.Edged(v)
		candidate := p.Intersection(neighbors)
		for node := range candidate {
			if pg.Degree(node) < len(r) {
				candidate.Remove(node)
			}
		}
		BronKerbosch2(pg, reporter, append(r, v), candidate, x.Intersection(neighbors))
		x.Add(v)
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
