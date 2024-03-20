package graph

import (
	"fmt"
	"math"
)


func BFS(root Features) {
	// for k, v := range b.Relations {
	// 	fmt.Printf("%s %+v \n", k, v.Callee)
	// }
	// fmt.Println()
	history := map[Features]struct{}{root: {}}
	queue := []Features{root}
	var last Features = root
	for len(queue) != 0 {
		cur := queue[0]
		if cur != nil {
			fmt.Printf("%s  ", cur.Name())
			for _, v := range cur.Nodes() {
				if _, ok := history[v]; !ok {
					queue = append(queue, v)
					history[v] = struct{}{}
				}
			}
		}
		if cur == last {
			fmt.Println("\n--------------------------------------------------")
			last = queue[len(queue)-1]
		}
		queue = queue[1:]
	}
}

// 计算皮尔逊相关系数
func Pearson(x, y Features) (float64, error) {
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
	vectorx := x.Features()
	vectory := y.Features()
	n := float64(len(vectorx))
	if len(vectorx) != len(vectory) {
		return 0, fmt.Errorf("slices x and y have different lengths")
	}
	for i := 0; i < len(vectorx); i++ {
		sumX += float64(vectorx[i])
		sumY += float64(vectory[i])
		sumXY += float64(vectorx[i]) * float64(vectory[i])
		sumX2 += float64(vectorx[i]) * float64(vectorx[i])
		sumY2 += float64(vectory[i]) * float64(vectory[i])
	}
	numerator := (n*sumXY - sumX*sumY)
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	if denominator == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return numerator / denominator, nil
}

// dotProduct 计算两个向量的点积
func dotProduct(x, y Features) (float64, error) {
	var product float64
	vectorx := x.Features()
	vectory := y.Features()
	if len(vectorx) != len(vectory) {
		return 0, fmt.Errorf("slices x and y have different lengths")
	}
	for i, value := range vectorx {
		product += float64(value) * float64(vectory[i])
	}
	return product, nil
}

// vectorNorm 计算向量的范数（长度）
func vectorNorm(x Features) float64 {
	var sum float64
	vectorx := x.Features()
	for _, value := range vectorx {
		sum += float64(value) * float64(value)
	}
	return math.Sqrt(sum)
}

// cosineSimilarity 计算两个向量的余弦相似度
func CosineSimilarity(v1, v2 Features) (float64, error) {
	product, err := dotProduct(v1, v2)
	if err != nil {
		fmt.Println(err)
	}
	normV1 := vectorNorm(v1)
	normV2 := vectorNorm(v2)
	// 防止除以0
	if normV1 == 0 || normV2 == 0 {
		return 0, fmt.Errorf("norm of vector must not be zero")
	}
	return product / (normV1 * normV2), nil
}

func max(i, j float64) float64 {
	if i < j {
		return j
	}
	return i
}
func min(i,j float64)float64{
	if i<j{
		return i
	}
	return j
}
