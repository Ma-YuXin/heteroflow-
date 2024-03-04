package flowmapmaker

import (
	"fmt"
	"testing"
)

func case1() {
	b := Build{map[string]*FuncFeatures{
		"A": {
			FuncName: "A",
			Callee: map[string]*FuncFeatures{
				"B": nil,
				"C": nil,
				"D": nil,
			},
		},
		"B": {
			FuncName: "B",
			Callee: map[string]*FuncFeatures{
				"E": nil,
				"F": nil,
			},
		},
		"C": {
			FuncName: "C",
			Callee: map[string]*FuncFeatures{
				"G": nil,
			},
		},
		"D": {
			FuncName: "D",
			Callee: map[string]*FuncFeatures{
				"H": nil,
			},
		},
		"E": {
			FuncName: "E",
			Callee: map[string]*FuncFeatures{
				"I": nil,
				"c": nil,
			},
		},
		"F": {
			FuncName: "F",
			Callee: map[string]*FuncFeatures{
				"I": nil,
			},
		},
		"G": {
			FuncName: "G",
			Callee: map[string]*FuncFeatures{
				"I": nil,
			},
		},
		"H": {
			FuncName: "H",
			Callee: map[string]*FuncFeatures{
				"I": nil,
				"A": nil,
			},
		},
		"I": {
			FuncName: "I",
			Callee: map[string]*FuncFeatures{
				"J": nil,
			},
		},
		"J": {
			FuncName: "J",
			Callee: map[string]*FuncFeatures{
				"L": nil,
			},
		},
		"L": {
			FuncName: "L",
			Callee:   map[string]*FuncFeatures{},
		}}}
	root := b.BuildGraphFromRoot("A")
	b.BFS(root)
}
func case2() {
	b := Build{map[string]*FuncFeatures{
		"A": {
			FuncName: "A",
			Callee: map[string]*FuncFeatures{
				"B": nil,
				"C": nil,
				"D": nil,
				"A": nil,
			},
		},
		"B": {
			FuncName: "B",
			Callee: map[string]*FuncFeatures{
				"E": nil,
				"F": nil,
			},
		},
		"C": {
			FuncName: "C",
			Callee: map[string]*FuncFeatures{
				"G": nil,
			},
		},
		"D": {
			FuncName: "D",
			Callee: map[string]*FuncFeatures{
				"H": nil,
			},
		},
		"E": {
			FuncName: "E",
			Callee: map[string]*FuncFeatures{
				"I": nil,
				"c": nil,
			},
		},
		"F": {
			FuncName: "F",
			Callee: map[string]*FuncFeatures{
				"I": nil,
			},
		},
		"G": {
			FuncName: "G",
			Callee: map[string]*FuncFeatures{
				"I": nil,
			},
		},
		"H": {
			FuncName: "H",
			Callee: map[string]*FuncFeatures{
				"I": nil,
				"A": nil,
			},
		},
		"I": {
			FuncName: "I",
			Callee: map[string]*FuncFeatures{
				"J": nil,
			},
		},
		"J": {
			FuncName: "J",
			Callee: map[string]*FuncFeatures{
				"L": nil,
			},
		},
		"L": {
			FuncName: "L",
			Callee:   map[string]*FuncFeatures{},
		}}}
	root := b.BuildGraphFromRoot("A")
	b.BFS(root)
}
func TestBFS(t *testing.T) {
	case1()
	fmt.Println()
	fmt.Println()
	case2()
}
