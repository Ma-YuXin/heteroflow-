package graph

import (
	"fmt"
	"hash/fnv"
	"reflect"
)

type ReportType int

const (
	MapReporter ReportType = iota
	CollectingReporter
	CountingReporter
	ChannelReporter
)

type Reporter interface {
	Record(vertexSet []Vertex)
	Report()
}

type mapReporter struct {
	store map[uint64]empty
	collectingReporter
}

type collectingReporter struct {
	cliques [][]Vertex
}

type countingReporter struct {
	cliques int
}

type channelReporter struct {
	cliques chan []Vertex
}

func NewRepoter(class ReportType) Reporter {
	switch class {
	case MapReporter:
		return &mapReporter{
			store:              make(map[uint64]empty),
			collectingReporter: collectingReporter{make([][]Vertex, 0)},
		}
	case CollectingReporter:
		return &collectingReporter{make([][]Vertex, 0)}
	case CountingReporter:
		return &countingReporter{}
	case ChannelReporter:
		return &channelReporter{make(chan []Vertex)}
	default:
		return nil
	}
}
func (r *mapReporter) Record(vertexSet []Vertex) {
	hasher, err := hashSlice(vertexSet)
	if err != nil {
		fmt.Println(err)
	}
	if _, ok := r.store[hasher]; !ok {
		r.collectingReporter.Record(vertexSet)
	}
}
func (r *collectingReporter) Record(vertexSet []Vertex) {
	cc := make([]Vertex, len(vertexSet))
	copy(cc, vertexSet)
	r.cliques = append(r.cliques, cc)
}

func (r *countingReporter) Record(vertexSet []Vertex) {
	r.cliques += 1
}

func (r *channelReporter) Record(vertexSet []Vertex) {
	cc := make([]Vertex, len(vertexSet))
	copy(cc, vertexSet)
	r.cliques <- cc
}

func (r *mapReporter) Report() {
	fmt.Println("total :", len(r.cliques))
	for _, v := range r.cliques {
		fmt.Println(v)
	}
}

func (r *collectingReporter) Report() {
	fmt.Println("total :", len(r.cliques))
	for _, v := range r.cliques {
		fmt.Println(v)
	}
}

func (r *countingReporter) Report() {
	fmt.Println(r.cliques)
}

func (r *channelReporter) Report() {
	c := <-r.cliques
	fmt.Println(c)
}

// mapHash creates a hash of a map which can be used to compare if two maps are equal.
func mapHash(m map[string]interface{}) string {
	// Generate a string hash based on map content.
	// WARNING: This hash function is a simple demonstration and is not guaranteed to
	// be collision-free in all cases.
	var hash string
	keys := reflect.ValueOf(m).MapKeys()
	for _, k := range keys {
		hash += fmt.Sprintf("%v:%v|", k, m[k.String()])
	}
	return hash
}

// hashStructSlice 接收 MyStruct 类型的切片并返回它的哈希值。
func hashSlice(vertexSet []Vertex) (uint64, error) {
	var totalHash uint64 = 0
	for _, item := range vertexSet {
		// 将结构体字段转换为字节并写入哈希器
		totalHash += hashStruct(item)
	}
	// 返回计算出的哈希值
	return totalHash, nil
}
func hashStruct(v Vertex) uint64 {
	hasher := fnv.New64a()

	_, err := hasher.Write([]byte(v.Key))
	if err != nil {
		panic(err)
	}
	// 注意：这里为了简化代码，我们没有错误处理
	hasher.Write([]byte(v.Key))
	hasher.Write([]byte(v.Value))
	return hasher.Sum64()
}
