package graph

import "fmt"

type PivotSelection int

const (
	MaxDegreeLocal  PivotSelection = iota
	MaxDegreeLocalX                = iota
)

// func bronKerbosch2bGP(graph *ProductGraph, reporter Reporter) {
// 	// Bron-Kerbosch algorithm with pivot of highest degree within remaining candidates
// 	// chosen from candidates only (IK_GP).
// 	order := graph.Order()
// 	if order == 0 {
// 		return
// 	}
// 	pivot := graph.maxDegreeVertex()
// 	excluded := make(VertexSet, order)
// 	for i := 0; i < order; i++ {
// 		v := Vertex(i)
// 		neighbours := graph.neighbours(v)
// 		if !neighbours.Contains(pivot) {
// 			neighbouringExcluded := neighbours.Intersection(excluded)
// 			if len(neighbouringExcluded) < len(neighbours) {
// 				neighbouringCandidates := neighbours.Difference(neighbouringExcluded)
// 				visit(
// 					graph, reporter,
// 					MaxDegreeLocal,
// 					neighbouringCandidates,
// 					neighbouringExcluded,
// 					[]Vertex{v})
// 			}
// 			excluded.Add(v)
// 		}
// 	}
// }

func BronKerbosch2aGP(graph *ProductGraph, reporter Reporter) {
	// Bron-Kerbosch algorithm with pivot of highest degree within remaining candidates
	// chosen from candidates only (IK_GP).
	candidates := graph.AllConnectedVertices()
	if candidates.IsEmpty() {
		return
	}
	excluded := make(VertexSet, len(candidates))
	visit(graph, reporter, MaxDegreeLocal, nil, candidates, excluded)
}
func visit(graph *ProductGraph, reporter Reporter, pivotSelection PivotSelection, clique []Vertex, candidates VertexSet, excluded VertexSet) {
	if len(candidates) == 1 {
		for v := range candidates {
			// Same logic as below, stripped down
			neighbours := graph.Edged(v)
			if excluded.IsDisjoint(neighbours) {
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
		localDegree := neighbours.IntersectionLen(candidates)
		if localDegree == 0 {
			// Same logic as below, stripped down
			if neighbours.IsDisjoint(excluded) {
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
			localDegree := neighbours.IntersectionLen(candidates)
			if seenLocalDegree < localDegree {
				seenLocalDegree = localDegree
				pivot = v
			}
		}
	}

	for _, v := range remainingCandidates {
		neighbours := graph.Edged(v)
		if neighbours.Contains(pivot) {
			continue
		}
		candidates.Remove(v)
		neighbouringCandidates := neighbours.Intersection(candidates)
		if !neighbouringCandidates.IsEmpty() {
			neighbouringExcluded := neighbours.Intersection(excluded)
			visit(
				graph, reporter,
				pivotSelection,
				append(clique, v),
				neighbouringCandidates,
				neighbouringExcluded)
		} else {
			if neighbours.IsDisjoint(excluded) {
				fmt.Println(append(clique, v))
				reporter.Record(append(clique, v))
			}
		}
		excluded.Add(v)
	}
}
