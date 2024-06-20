package cfw

import (
	"errors"
	"fmt"
	"heterflow/pkg/codeaid/def"
	"math"
	"strconv"
)

type Mapping[T any] interface {
	~int | ~float64 | ~int64
	Injection(NodeVectors) T
}

type Transform struct {
	nodeVector NodeVectors
	maximum    Vector[float64]
	minimum    Vector[float64]
	interval   Interval
	// numSplitsPerInterval Vector[int]
	// vector            StatisticalVectors
}

type NodeVectors map[string]Vector[float64]

type Vector[T comparable] []T

type StatisticalVectors interface {
	InnerProduct(StatisticalVectors) (float64, error)
	Insert(int)
	At(int) (float64, error)
	ForEach(func(int, float64))
	Len() int
}

type MapStatisticalVector map[int]float64

type SliceStatisticalVector Vector[float64]

type AggregatedFeature = ProgramFeatures

type value struct {
	// degree  float64
	*Node
}

type GraphKernels struct {
	// g           Graph
	// numSplitsPerInterval Vector[int]
	k           int //迭代次数
	adjacencies map[string]value
}

type IntervalType int

type Interval interface {
	//特征向量的长度
	// Buckets() int
	//将节点特征映射到特征向量的某一位置
	Hash(Vector[float64]) int
}

type FixedInterval struct {
	// buckets   int
	tran      *Transform
	intervals Vector[float64]
	base      []int
}

type DecreaseInterval struct {
	// buckets   int
	tran      *Transform
	intervals []Vector[float64]
	base      []int
}

const (
	FixedIntervalType IntervalType = iota
	DecreaseIntervalType
)

var (
	buckets              int
	numSplitsPerInterval Vector[int]
)

func init() {
	no := NewNode()
	numSplitsPerInterval = no.SegmentsPerInterval()
	if numSplitsPerInterval == nil || len(numSplitsPerInterval) == 0 {
		panic("numSplitsPerInterval is zero")
	}
	buckets = 1
	for _, v := range numSplitsPerInterval {
		buckets *= (v + 1)
	}
}
func FeatureVector(gra Graph, iterator int) StatisticalVectors {
	gk := NewGraphKernels(gra, iterator)
	if def.Debug {
		fmt.Println("successfully get GraphKernel")
	}
	transform := gk.Iterator()
	if def.Debug {
		fmt.Println("successfully Iterator")
	}
	if def.Debug {
		fmt.Println("next Injection")
	}
	return transform.Injection()
}

func NewMapStatisticalVector(hint int) MapStatisticalVector {
	if hint > 100 {
		hint = 100
	}
	return make(MapStatisticalVector, hint)
}

func NewSliceStatisticalVector(hint int) SliceStatisticalVector {
	return make(SliceStatisticalVector, hint)
}

func NewFixedInterval(t *Transform) *FixedInterval {
	maximum := t.maximum
	minimum := t.minimum
	length := len(maximum)
	divideHierarchies := numSplitsPerInterval
	base := make([]int, length)
	base[0] = 1
	for i := 1; i < length; i++ {
		base[i] = divideHierarchies[i-1] * base[i-1]
	}
	interval := make(Vector[float64], length)
	for i, v := range maximum {
		if minimum[i] == math.MaxFloat64 && maximum[i] == 0 {
			interval[i] = 0
		} else {
			interval[i] = (v - minimum[i]) / float64(divideHierarchies[i])
		}
	}
	num := 1
	for _, v := range numSplitsPerInterval {
		num *= (v + 1)
	}
	return &FixedInterval{tran: t, intervals: interval, base: base}
}

func NewDecreaseInterval(t *Transform) *DecreaseInterval {
	maximum := t.maximum
	minimum := t.minimum
	length := len(maximum)
	divideHierarchies := numSplitsPerInterval
	base := make([]int, length)
	base[0] = 1
	for i := 1; i < length; i++ {
		base[i] = divideHierarchies[i-1] * base[i-1]
	}
	interval := make([]Vector[float64], length)
	for i, v := range maximum {
		vec := make(Vector[float64], 0, divideHierarchies[i])
		if minimum[i] == math.MaxFloat64 && maximum[i] == 0 {
			vec = append(vec, 0)
		} else {
			total := (math.Pow(def.FeatureDividBase, float64(divideHierarchies[i])) - 1) / (def.FeatureDividBase - 1)
			dif := v - minimum[i]
			spac := dif / total
			for num := 0; num < divideHierarchies[i]; num++ {
				copies := math.Pow(def.FeatureDividBase, float64(num))
				vec = append(vec, copies*spac)
			}
			vec[len(vec)-1] = v
		}
		interval[i] = vec
	}
	return &DecreaseInterval{tran: t, intervals: interval, base: base}
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

func NewValue(node *Node) value {
	var val value
	ft := NewNode()
	if node != nil {
		ft.Add(node)
	}
	val.Node = ft
	return val
}

func NewGraphKernels(gra Graph, ittime int) GraphKernels {
	rel := gra.Relation()
	adj := make(map[string]value, len(gra.Relation()))
	for k, v := range rel {
		val := NewValue(v)
		adj[k] = val
	}
	if ittime == 0 {
		ittime = def.GraphKernalDefaultIteratorTimes
	}

	return GraphKernels{

		k:           ittime,
		adjacencies: adj,
	}
}

func (in1 *value) AddValue(in2 value) {
	in1.Add(in2.Node)
}

func (v value) DeepCopy() (val value) {
	val.Node = v.Node.DeepCopy()
	return
}

// 图核法，对每个节点聚合其邻居节点的信息，进行迭代，也就是感受野的大小
func (gk *GraphKernels) Iterator() Transform {
	for i := 0; i < gk.k; i++ {
		gk.aggregate()
	}
	return gk.Normalization()
}

// 具体的聚合操作，对每个节点聚合其邻居信息
func (gk *GraphKernels) aggregate() {
	rel := gk.adjacencies
	next := make(map[string]value, len(rel))
	for funcName, nodeval := range rel {
		tmp := nodeval.DeepCopy()
		for k := range nodeval.Callee {
			tmp.AddValue(gk.adjacencies[k])
		}
		next[funcName] = tmp
	}
	gk.adjacencies = next
}

// 由于节点向量每个特征的值的差异很大，因而需要归一化处理，又因为对于每个特征来说值的分布可能都集中最小值附近，
// 故使用间隔划分，统计落在每个区间的值的数量。
func (gk *GraphKernels) Normalization() Transform {
	sum := NewValue(nil)
	for _, v := range gk.adjacencies {
		sum.AddValue(v)
	}
	length := sum.MetricsNumber()
	nv := make(NodeVectors, len(gk.adjacencies))
	minimum, maximum := make(Vector[float64], length), make(Vector[float64], length)
	for i := range minimum {
		minimum[i] = math.MaxFloat64
	}
	// fmt.Printf("%f %+v\n", sum.degree, sum.feature)
	for k, v := range gk.adjacencies {
		// fmt.Printf("%f %+v\n\n", v.degree, v.feature)
		vec := make(Vector[float64], 0, length)
		feat := v.Features()
		sumfeat := sum.Features()
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
	res := Transform{}
	res.nodeVector = nv
	res.maximum = maximum
	res.minimum = minimum
	interval := NewInterval(DecreaseIntervalType, &res)
	res.interval = interval
	return res
}

// 图核法对每个节点向量进行映射，得到统计向量
func (tran *Transform) Injection() StatisticalVectors {
	// fmt.Println(lenth)

	// fmt.Println(lenth)
	res := NewMapStatisticalVector(buckets)
	// fmt.Println(len(tran.vector))
	for _, vector := range tran.nodeVector {
		idx := tran.interval.Hash(vector)
		// fmt.Printf("%d ", idx)
		// if idx == 0 {
		// fmt.Println()
		// fmt.Println("vector", vector)
		// fmt.Println("base", base)
		// fmt.Println("minimum", tran.minimum)
		// fmt.Println("maximum", tran.maximum)
		// fmt.Println("interval", tran.interval)
		// }

		res.Insert(idx)
	}
	// fmt.Println("\n")
	return res
}

// func (f *FixedInterval) Buckets() int {
// 	return f.buckets
// }

func (f *FixedInterval) Hash(vector Vector[float64]) int {
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

// func (d *DecreaseInterval) Buckets() int {
// 	return d.buckets
// }

func (d *DecreaseInterval) Hash(vector Vector[float64]) int {
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

func (sv SliceStatisticalVector) ForEach(f func(int, float64)) {
	for i, v := range sv {
		f(i, v)
	}
}

func (sv SliceStatisticalVector) Len() int {
	return len(sv)
}

func (msv MapStatisticalVector) InnerProduct(another StatisticalVectors) (float64, error) {
	intersectionlen := 0.0
	sum := 0.0
	for k, v := range msv {
		va, err := another.At(k)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		intersectionlen++
		sum += 1 / (math.Abs(v-va) + def.Alpha)
		// sum += (v * va)
	}
	sum = sum / (intersectionlen * (1.0 / def.Alpha))
	totallen := float64(buckets)
	ans := def.StatisticalVectorDisjointWeight*((totallen-intersectionlen)/totallen) + def.StatisticalVectorIntersectionWeight*sum
	return ans, nil
}

func (msv MapStatisticalVector) ForEach(f func(int, float64)) {
	for k, v := range msv {
		f(k, v)
	}
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

func (msv MapStatisticalVector) Len() int {
	return len(msv)
}
