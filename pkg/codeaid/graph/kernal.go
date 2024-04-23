package graph

import (
	"errors"
	"math"
)

const (
	defaultIteratorTimes = 3
	metricsNum           = 11
)

type Mapping[T any] interface {
	int | float64 | int64
	Injection(NodeVectors) T
}

type Transform struct {
	nodeVector        NodeVectors
	maximum           Vector
	minimum           Vector
	interval          Vector
	divideHierarchies []int
	// vector            StatisticalVectors
}

type NodeVectors map[string]Vector

type Vector []float64
type StatisticalVectors []int

type AggregatedFeature = Programmer

type value struct {
	degree  float64
	feature Features
}

type GraphKernels struct {
	g           Graph
	K           int //迭代次数
	adjacencies map[string]value
}

func NewTransfrm() Transform {
	hierarchies := []int{6, 3, 1, 3, 3, 3, 4, 2, 1, 3, 4}

	return Transform{divideHierarchies: hierarchies}
}

func NewValue(deg int, node *Node) value {
	var val value
	val.degree = float64(deg)
	ft := NewFeatures(ProgrammerFeature)
	if node != nil {
		ft.AddInfo(node)
	}
	val.feature = ft
	return val
}

func NewGraphKernels(gra Graph, ittime int) GraphKernels {
	rel := gra.Relation()
	adj := make(map[string]value, len(gra.Relation()))
	for k, v := range rel {
		val := NewValue(len(v.Callee), v)
		adj[k] = val
	}
	if ittime == 0 {
		ittime = defaultIteratorTimes
	}
	return GraphKernels{
		g:           gra,
		K:           ittime,
		adjacencies: adj,
	}
}

func (in1 *value) AddValue(in2 value) {
	in1.degree += in2.degree
	in1.feature.AddInfo(in2.feature)
}

func (v value) DeepCopy() (val value) {
	ft := NewFeatures(ProgrammerFeature)
	ft.AddInfo(v.feature)
	val.feature = ft
	val.degree = v.degree
	return
}

func (gk *GraphKernels) Iterator() Transform {
	for i := 0; i < gk.K; i++ {
		gk.aggregate()
	}
	return gk.Normalization()
}

func (gk *GraphKernels) aggregate() {
	rel := gk.g.Relation()
	next := make(map[string]value, len(rel))
	for funcName, funcNode := range rel {
		tmp := gk.adjacencies[funcName].DeepCopy()
		for k := range funcNode.Callee {
			tmp.AddValue(gk.adjacencies[k])
		}
		next[funcName] = tmp
	}
	gk.adjacencies = next
}

func (gk *GraphKernels) Normalization() Transform {
	sum := NewValue(0, nil)
	for _, v := range gk.adjacencies {
		sum.AddValue(v)
	}
	res := NewTransfrm()
	nv := make(NodeVectors, len(gk.adjacencies))
	minimum, maximum, interval := make(Vector, metricsNum), make(Vector, metricsNum), make(Vector, metricsNum)
	for i := range minimum {
		minimum[i] = math.MaxFloat64
	}
	// fmt.Printf("%f %+v\n", sum.degree, sum.feature)
	for k, v := range gk.adjacencies {
		// fmt.Printf("%f %+v\n\n", v.degree, v.feature)

		vec := make(Vector, 0, metricsNum)
		feat := append(v.feature.Features(), int(v.degree))
		sumfeat := append(sum.feature.Features(), int(sum.degree))
		for i, v := range feat {
			if v == 0 {
				vec = append(vec, 0.0)
				// minimum[i] = 0
			} else {
				num := float64(sumfeat[i]) / float64(v)
				vec = append(vec, num)
				minimum[i] = min(minimum[i], num)
				maximum[i] = max(maximum[i], num)
			}
			// fmt.Println(v, sumfeat[i],float64(sumfeat[i])/float64(v))
		}
		// fmt.Println(len(vec))
		nv[k] = vec
	}
	// fmt.Println(minimum, maximum)
	for i, v := range maximum {
		if minimum[i] == math.MaxFloat64 && maximum[i] == 0 {
			interval[i] = 0
		} else {
			interval[i] = (v - minimum[i]) / float64(res.divideHierarchies[i])
		}
	}
	res.nodeVector = nv
	res.maximum = maximum
	res.minimum = minimum
	res.interval = interval
	return res
}

func (tran *Transform) Injection() StatisticalVectors {
	lenth := 1
	// fmt.Println(tran.divideHierarchies)
	for _, v := range tran.divideHierarchies {
		lenth *= (v + 1)
	}
	// fmt.Println(lenth)

	res := make(StatisticalVectors, lenth)
	// fmt.Println(len(tran.vector))
	base := make([]int, metricsNum)
	base[0] = 1
	for i := 1; i < metricsNum; i++ {
		base[i] = tran.divideHierarchies[i-1] * base[i-1]
		// fmt.Println("base[i]: ", base[i], tran.divideHierarchies[i-1])
	}
	for _, vector := range tran.nodeVector {
		idx := hash(tran.minimum, tran.interval, vector, base)
		// fmt.Printf("%d ", idx)
		if idx == 0 {
			// fmt.Println()
			// fmt.Println("vector", vector)
			// fmt.Println("base", base)
			// fmt.Println("minimum", tran.minimum)
			// fmt.Println("maximum", tran.maximum)
			// fmt.Println("interval", tran.interval)
		}
		res[idx]++
	}
	return res
}

func hash(minimum, interval, vector Vector, base []int) int {
	idx := 0
	// fmt.Println("+++++++++++++++++++++++++++++")
	for pos, val := range vector {
		dif := val - minimum[pos]
		if interval[pos] != 0 {
			mul := int(dif / interval[pos])
			idx += mul * base[pos]
			// fmt.Println("idx: ", idx, "val: ", val, "minval:", minimum[pos], "dif: ", dif, "interval[", pos, "] ", interval[pos], "mul: ", mul, "base", base[pos])
		}
		//  else {
		// fmt.Println("interval is zero ", idx, dif, interval[pos])		// }
	}
	// fmt.Println("+++++++++++++++++++++++++++++")
	return idx
}

func (sv StatisticalVectors) InnerProduct(another StatisticalVectors) (sum int, err error) {
	if len(sv) != len(another) {
		err = errors.New("len of two vector is different")
		return
	}
	for i, v := range sv {
		sum += (v * another[i])
	}
	return
}
