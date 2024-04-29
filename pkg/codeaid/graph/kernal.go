package graph

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	FixedIntervalType IntervalType = iota
	DecreaseIntervalType
	defaultIteratorTimes         = 2
	metricsNum                   = 11
	divideBase           float64 = 3
)

type Mapping[T any] interface {
	int | float64 | int64
	Injection(NodeVectors) T
}

type Transform struct {
	nodeVector        NodeVectors
	maximum           Vector
	minimum           Vector
	interval          Interval
	divideHierarchies []int
	// vector            StatisticalVectors
}

type NodeVectors map[string]Vector

type Vector []float64

type StatisticalVectors interface {
	InnerProduct(StatisticalVectors) (float64, error)
	Insert(int)
	At(int) (float64, error)
}

type MapStatisticalVector map[int]float64

type SliceStatisticalVector []float64

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

type IntervalType int

type Interval interface {
	Buckets() int
	Hash(Vector) int
}

type FixedInterval struct {
	tran      *Transform
	intervals Vector
	base      []int
}
type DecreaseInterval struct {
	tran      *Transform
	intervals []Vector
	base      []int
}

func NewMapStatisticalVector(hint int) MapStatisticalVector {
	if hint < 4 {
		hint = 4
	}
	return make(MapStatisticalVector, hint>>2)
}

func NewSliceStatisticalVector(hint int) SliceStatisticalVector {
	return make(SliceStatisticalVector, hint)
}
func NewFixedInterval(t *Transform) *FixedInterval {
	maximum := t.maximum
	minimum := t.minimum
	divideHierarchies := t.divideHierarchies
	base := make([]int, metricsNum)
	base[0] = 1
	for i := 1; i < metricsNum; i++ {
		base[i] = divideHierarchies[i-1] * base[i-1]
		// fmt.Println("base[i]: ", base[i], tran.divideHierarchies[i-1])
	}
	interval := make(Vector, metricsNum)
	for i, v := range maximum {
		if minimum[i] == math.MaxFloat64 && maximum[i] == 0 {
			interval[i] = 0
		} else {
			interval[i] = (v - minimum[i]) / float64(divideHierarchies[i])
		}
	}
	return &FixedInterval{tran: t, intervals: interval, base: base}
}

func NewDecreaseInterval(t *Transform) *DecreaseInterval {
	maximum := t.maximum
	minimum := t.minimum
	divideHierarchies := t.divideHierarchies
	base := make([]int, metricsNum)
	base[0] = 1
	for i := 1; i < metricsNum; i++ {
		base[i] = divideHierarchies[i-1] * base[i-1]
		// fmt.Println("base[i]: ", base[i], tran.divideHierarchies[i-1])
	}
	interval := make([]Vector, metricsNum)
	for i, v := range maximum {
		vec := make(Vector, 0, divideHierarchies[i])
		if minimum[i] == math.MaxFloat64 && maximum[i] == 0 {
			vec = append(vec, 0)
		} else {
			total := (math.Pow(divideBase, float64(divideHierarchies[i])) - 1) / (divideBase - 1)
			dif := v - minimum[i]
			spac := dif / total
			for num := 0; num < divideHierarchies[i]; num++ {
				copies := math.Pow(divideBase, float64(num))
				vec = append(vec, copies*spac)
			}
			vec[len(vec)-1] = v
			// total := (1 << (divideHierarchies[i])) - 1
			// dif := v - minimum[i]
			// spac := dif / float64(total)
			// for num := 0; num < divideHierarchies[i]; num++ {
			// 	copies := 1 << num
			// 	vec = append(vec, float64(copies)*spac)
			// }
			// vec[len(vec)-1] = v
		}
		interval[i] = vec
	}
	return &DecreaseInterval{tran: t, intervals: interval, base: base}
}

func NewTransfrm() Transform {
	// hierarchies := []int{6, 3, 1, 3, 3, 3, 4, 2, 1, 3, 4}
	hierarchies := []int{12, 6, 2, 6, 6, 6, 8, 4, 2, 6, 4}
	return Transform{divideHierarchies: hierarchies}
}

func NewInterval(it IntervalType, t *Transform) Interval {
	switch it {
	case FixedIntervalType:
		return NewFixedInterval(t)
	case DecreaseIntervalType:
		return NewDecreaseInterval(t)
	default:
		return nil
	}
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

	nv := make(NodeVectors, len(gk.adjacencies))
	minimum, maximum := make(Vector, metricsNum), make(Vector, metricsNum)
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
	res := NewTransfrm()
	res.nodeVector = nv
	res.maximum = maximum
	res.minimum = minimum
	interval := NewInterval(DecreaseIntervalType, &res)
	res.interval = interval
	return res
}

func (tran *Transform) Injection() StatisticalVectors {

	// fmt.Println(lenth)
	lenth := tran.interval.Buckets()
	res := NewMapStatisticalVector(lenth)
	// fmt.Println(len(tran.vector))

	for _, vector := range tran.nodeVector {
		idx := tran.interval.Hash(vector)
		fmt.Printf("%d ", idx)
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
	fmt.Println("\n")
	return res
}

func (f *FixedInterval) Buckets() int {
	length := 1
	// fmt.Println(tran.divideHierarchies)
	for _, v := range f.tran.divideHierarchies {
		length *= (v + 1)
	}
	return length
}

func (f *FixedInterval) Hash(vector Vector) int {
	idx := 0
	for pos, val := range vector {
		dif := val - f.tran.minimum[pos]
		// if dif < 0 && val != 0 {
		// 	fmt.Println(dif)
		// }
		if f.intervals[pos] != 0 {
			mul := int(dif / f.intervals[pos])
			idx += mul * f.base[pos]
			// fmt.Println("idx: ", idx, "val: ", val, "minval:", minimum[pos], "dif: ", dif, "interval[", pos, "] ", interval[pos], "mul: ", mul, "base", base[pos])
		}
		//  else {
		// fmt.Println("interval is zero ", idx, dif, interval[pos])		// }
	}
	return idx
}

func (d *DecreaseInterval) Buckets() int {
	length := 1
	// fmt.Println(tran.divideHierarchies)
	for _, v := range d.tran.divideHierarchies {
		length *= v
	}
	return length
}

func (d *DecreaseInterval) Hash(vector Vector) int {
	idx := 0
	for pos, val := range vector {
		dif := val - d.tran.minimum[pos]
		mul := d.bsearch(pos, dif)
		idx += mul * d.base[pos]
	}
	return idx
}

func (d *DecreaseInterval) bsearch(pos int, val float64) int {
	inter := d.intervals[pos]
	left, right := 0, len(inter)-1
	for left <= right {
		mid := (left + right) >> 1
		if inter[mid] >= val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func (sv SliceStatisticalVector) InnerProduct(another StatisticalVectors) (sum float64, err error) {
	for i, v := range sv {
		va, err := another.At(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum += (v * va)
	}
	return
}

func (sv SliceStatisticalVector) At(pos int) (sum float64, err error) {
	if pos > len(sv) {
		return 0, errors.New("pos greater than sv's length")
	}
	return sv[pos], nil
}

func (sv SliceStatisticalVector) Insert(pos int) {
	sv[pos]++
}

func (msv MapStatisticalVector) InnerProduct(another StatisticalVectors) (sum float64, err error) {
	for k, v := range msv {
		va, err := another.At(k)
		if err != nil {
			// fmt.Println(err)
			continue
		}

		sum += 1 / (math.Abs(v-va) + 1)

		// sum += (v * va)
	}
	return
}

func (msv MapStatisticalVector) Insert(pos int) {
	msv[pos]++
}
func (msv MapStatisticalVector) At(pos int) (sum float64, err error) {
	if v, ok := msv[pos]; ok {
		return v, nil
	}
	return 0, errors.New(strconv.Itoa(pos) + "doesn't exist")
}
