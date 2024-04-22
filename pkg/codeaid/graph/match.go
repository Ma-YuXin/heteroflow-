package graph

import (
	"fmt"
	"math"
	"strconv"
)

const Inf = 1<<63 - 1

type MatcherType int

const (
	Hungarian MatcherType = iota
	KM
	GaleShapley
	threshold = 0.95
)

type Matcher interface {
	Match() float64
	Result() map[string]string
	Left() map[string]map[string]float64
	Right() map[string]map[string]float64
}

func NewMatcher(class MatcherType, adj *adjacentTable) Matcher {
	switch class {
	case Hungarian:
		return &hungarian{
			adjacentTable: adj,
			match:         map[string]string{},
			visited:       map[string]empty{},
		}
	case KM:
		return &kM{
			adjacentTable: adj,
			visitedleft:   make(map[string]bool, len(adj.leftPref)),
			visitedright:  make(map[string]bool, len(adj.rightPref)),
			weightleft:    make(map[string]float64, len(adj.leftPref)),
			weightright:   make(map[string]float64, len(adj.rightPref)),
			delta:         make(map[string]float64, len(adj.rightPref)),
			match:         make(map[string]string, len(adj.rightPref)),
		}
	case GaleShapley:
		return &galeshapley{
			adjacentTable: adj,
			match:         make(map[string]string, len(adj.rightPref)),
		}
	}
	return nil
}

type hungarian struct {
	*adjacentTable
	match   map[string]string
	visited map[string]empty
	// left, right map[Features]empty
}
type galeshapley struct {
	*adjacentTable
	match map[string]string
	// left, right map[Features]empty
}
type kM struct {
	*adjacentTable
	visitedleft  map[string]bool
	visitedright map[string]bool
	weightleft   map[string]float64
	weightright  map[string]float64
	delta        map[string]float64
	// left, right  map[Features]empty
	match map[string]string
}

type adjacentTable struct {
	leftPref  map[string]map[string]float64
	rightPref map[string]map[string]float64
	virtual   map[string]empty
	swaped    bool
}

func (adj *adjacentTable) init(left, right map[Features]empty) {
	for v1 := range left {
		for v2 := range right {
			similarity, err := Pearson(v1, v2)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("similarity  ", v1.Name(), v2.Name(), similarity)
			if similarity > threshold {
				if value, ok := adj.leftPref[v1.Name()]; ok {
					value[v2.Name()] = similarity
					adj.leftPref[v1.Name()] = value
				} else {
					adj.leftPref[v1.Name()] = map[string]float64{v2.Name(): similarity}
				}
				if value, ok := adj.rightPref[v2.Name()]; ok {
					value[v1.Name()] = similarity
					adj.rightPref[v2.Name()] = value
				} else {
					adj.rightPref[v2.Name()] = map[string]float64{v1.Name(): similarity}
				}
			}
		}
	}
}

// Find 尝试为顶点 u 找到匹配
func (table *hungarian) find(u string) bool {
	for v := range table.leftPref[u] {
		if _, ok := table.visited[v]; !ok {
			table.visited[v] = empty{}
			if table.match[v] == "" || table.find(table.match[v]) {
				table.match[v] = u
				return true
			}
		}
	}
	return false
}

func (table *hungarian) Left() map[string]map[string]float64 {
	return table.leftPref
}
func (table *hungarian) Right() map[string]map[string]float64 {
	return table.rightPref
}

// hungarian 返回二部图的最大匹配数
func (table *hungarian) Match() float64 {
	var result int
	for u := range table.leftPref {
		table.visited = make(map[string]empty)
		if table.find(u) {
			result++
		}
	}
	return float64(result)
}

func (table *hungarian) Result() map[string]string {
	return table.match
}

func (km *kM) prepare() {
	if len(km.leftPref) == len(km.rightPref) {
		return
	}
	if len(km.leftPref) < len(km.rightPref) {
		km.leftPref, km.rightPref = km.rightPref, km.leftPref
		km.swaped = true
	}
	different := len(km.leftPref) - len(km.rightPref)
	for i := 0; i < different; i++ {
		str := strconv.Itoa(i)
		mm := make(map[string]float64, len(km.leftPref))
		for left, lp := range km.leftPref {
			lp[str] = float64(0.000000000000000001)
			mm[left] = float64(0.000000000000000001)
		}
		km.rightPref[str] = mm
		km.virtual[str] = empty{}
	}
	for left, lp := range km.leftPref {
		for right, rp := range km.leftPref {
			if _, ok := lp[right]; !ok {
				lp[right] = math.Inf(-1)
				rp[left] = math.Inf(-1)
			}
		}
		km.leftPref[left] = lp
	}
	// if km.swaped {
	// 	km.leftPref, km.rightPref = km.rightPref, km.leftPref
	// }
}

func (km *kM) init() {
	km.prepare()
	for k := range km.leftPref {
		km.visitedleft[k] = false
		km.weightleft[k] = math.Inf(-1)
	}
	for k := range km.weightleft {
		neighbor := km.leftPref[k]
		maximam := math.Inf(-1)
		for _, val := range neighbor {
			maximam = max(maximam, val)
		}
		km.weightleft[k] = maximam
	}
	for k := range km.rightPref {
		km.visitedright[k] = false
		km.weightright[k] = 0
	}
}

func (km *kM) find(x string) bool {
	km.visitedleft[x] = true
	// fmt.Println(x, "'s adjacent is ", km.leftPref[x])
	for k := range km.leftPref[x] {
		// fmt.Println("judge ", k, " weight is ", km.leftPref[x][k])
		if v := km.visitedright[k]; !v {
			// 0.000000000000000003
			fmt.Println("value is  ", x, km.weightleft[x], k, km.weightright[k], km.leftPref[x][k], " km.weightleft[x]+km.weightright[k]-km.leftPref[x][k] is ", km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
			if km.weightleft[x]+km.weightright[k]-km.leftPref[x][k] == 0 {
				km.visitedright[k] = true
				if km.match[k] == "" || km.find(km.match[k]) {
					// fmt.Println(k, "is match to ", x)
					km.match[k] = x
					return true
				}
			} else {
				// fmt.Println("set delte ", k, " ", km.weightleft[x], km.weightright[k], km.leftPref[x][k], km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
				km.delta[k] = min(km.delta[k], km.weightleft[x]+km.weightright[k]-km.leftPref[x][k])
			}
		}
	}
	return false
}

func (km *kM) Match() float64 {
	// for left, val := range km.adj {
	// 	for right, weight := range val {
	// 		fmt.Println(left, " between", right, " is ", weight)
	// 	}
	// }
	km.init()
	fmt.Printf("leftPref length is : %d  rightPref length is : %d\n", len(km.leftPref), len(km.rightPref))
	for x := range km.leftPref {
		fmt.Println("to find ", x, " pair")
		for {
			km.reset()
			if km.find(x) {
				break
			}
			mindelta := math.Inf(1)
			for k, v := range km.visitedright {
				// fmt.Println("visited right ", k, "  mindelta ", km.delta[k])
				if !v {
					mindelta = min(mindelta, km.delta[k])
				}
			}
			for k := range km.leftPref {
				if km.visitedleft[k] {
					km.weightleft[k] -= mindelta
				}
			}
			for k := range km.rightPref {
				if km.visitedright[k] {
					km.weightright[k] += mindelta
				}
			}
			// fmt.Println("mindelta ", mindelta)
		}
	}
	res := 0.0
	newmatcher := make(map[string]string, len(km.match))
	for right, left := range km.match {
		if _, ok := km.virtual[right]; ok {
			continue
		}
		if _, ok := km.virtual[left]; ok {
			continue
		}
		if km.swaped {
			newmatcher[left] = right
		} else {
			newmatcher[right] = left
		}
	}
	km.match = newmatcher
	for k, v := range km.match {
		res += km.leftPref[k][v]
		fmt.Println("res+=", "pair:", v, k, km.leftPref[k][v])
	}
	return res
}

func (km *kM) reset() {
	for k := range km.visitedleft {
		km.visitedleft[k] = false
	}
	for k := range km.visitedright {
		km.visitedright[k] = false
	}
	for k := range km.delta {
		km.delta[k] = math.Inf(1)
	}
}
func (km *kM) Left() map[string]map[string]float64 {
	return km.leftPref
}
func (km *kM) Right() map[string]map[string]float64 {
	return km.rightPref
}

func (km *kM) Result() map[string]string {
	return km.match
}

func (gs *galeshapley) Left() map[string]map[string]float64 {
	return gs.leftPref
}

func (gs *galeshapley) Right() map[string]map[string]float64 {
	return gs.rightPref
}

func (gs *galeshapley) Result() map[string]string {
	return gs.match
}

func (gs *galeshapley) prepare() {
	if len(gs.leftPref) == len(gs.rightPref) {
		return
	}
	if len(gs.leftPref) < len(gs.rightPref) {
		gs.leftPref, gs.rightPref = gs.rightPref, gs.leftPref
		gs.swaped = true
	}
	different := len(gs.leftPref) - len(gs.rightPref)
	for _, v := range gs.leftPref {
		for right := range gs.rightPref {
			if _, ok := v[right]; !ok {
				v[right] = math.Inf(-1)
			}
		}
		for i := 0; i < different; i++ {
			v[strconv.Itoa(i)] = math.Inf(-1)
		}
	}
	for _, v := range gs.rightPref {
		for left := range gs.leftPref {
			if _, ok := v[left]; !ok {
				v[left] = math.Inf(-1)
			}
		}
	}
	for i := 0; i < different; i++ {
		str := strconv.Itoa(i)
		m := make(map[string]float64, len(gs.leftPref))
		for left := range gs.leftPref {
			m[left] = math.Inf(-1)
		}
		gs.rightPref[str] = m
		gs.virtual[str] = empty{}
	}
}

func (gs *galeshapley) Match() float64 {
	gs.prepare()
	freeLeft := make([]string, 0, len(gs.leftPref))
	for left := range gs.leftPref {
		freeLeft = append(freeLeft, left)
	}
	proposals := make(map[string]string, len(gs.leftPref))
	res := 0.0
	for len(freeLeft) > 0 {
		curLeft := freeLeft[0]
		// fmt.Println("match for ", curLeft)
		lp := gs.leftPref[curLeft]
		for right, v := range lp {
			currentPartner, ok1 := proposals[right]
			if !ok1 {
				// fmt.Println("match", right, "to", curLeft)
				proposals[right] = curLeft
				res += v
				freeLeft = freeLeft[1:]
				break
			} else {
				rp := gs.rightPref[right]
				if rp[curLeft] > rp[currentPartner] {
					// fmt.Println("dismatch", right, "with", currentPartner)
					// fmt.Println("match", right, "to", curLeft)
					res += rp[curLeft]
					res -= rp[currentPartner]
					proposals[right] = curLeft
					freeLeft = freeLeft[1:]
					freeLeft = append(freeLeft, currentPartner)
					break
				}
			}
		}
	}
	for right, left := range proposals {
		if _, ok := gs.virtual[right]; ok {
			continue
		}
		if _, ok := gs.virtual[left]; ok {
			continue
		}
		if gs.swaped {
			gs.match[left] = right
		} else {
			gs.match[right] = left
		}
	}
	return res
}
