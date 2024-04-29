package util

import "errors"

type VertexSet[K comparable, V any] map[K]V

func Pop[K comparable, V any](v VertexSet[K, V]) (K, error) {
	for k := range v {
		Remove(v, k)
		return k, nil
	}
	var k K
	return k, errors.New("vertex set is empty, can't pop a vertex")
}

func IsEmpty[K comparable, V any](v VertexSet[K, V]) bool {
	return len(v) == 0
}
func IsDisjoint[K comparable, V any](input, v VertexSet[K, V]) bool {
	small, large := v, input
	if len(small) > len(large) {
		small, large = input, v
	}
	for v := range small {
		if Contains(large, v) {
			return false
		}
	}
	return true
}
func Contains[K comparable, V any](v VertexSet[K, V], vex K) bool {
	_, ok := v[vex]
	return ok
}
func Remove[K comparable, V any](v VertexSet[K, V], vex K) {
	delete(v, vex)
}
func Intersection[K comparable, V any](v, other VertexSet[K, V]) VertexSet[K, V] {
	small, large := v, other
	if len(small) > len(large) {
		small, large = other, v
	}
	result := make(VertexSet[K, V], len(small))
	for v := range small {
		if Contains(large, v) {
			Add(result, v)
		}
	}
	return result
}
func Difference[K comparable, V any](v VertexSet[K, V], other VertexSet[K, V]) VertexSet[K, V] {
	result := make(VertexSet[K, V], len(v))
	for val := range v {
		if !Contains(other, val) {
			Add(result, val)
		}
	}
	return result
}
func Union[K comparable, V any](v VertexSet[K, V], another VertexSet[K, V]) VertexSet[K, V] {
	result := make(VertexSet[K, V], len(v))
	for val := range v {
		Add(result, val)
	}
	for val := range another {
		Add(result, val)
	}
	return result
}

func Add[K comparable, V any](v VertexSet[K, V], vex K, val ...V) {
	var va V
	v[vex] = va
}
func IntersectionLen[K comparable, V any](v VertexSet[K, V], other VertexSet[K, V]) int {
	small, large := v, other
	if len(small) > len(large) {
		small, large = other, v
	}
	result := 0
	for v := range small {
		if Contains(large, v) {
			result++
		}
	}
	return result
}
func Pick[K comparable, V any](v VertexSet[K, V]) (k K, err error) {
	for v := range v {
		return v, nil
	}

	return k, errors.New("attempt to pick from empty set")
}
