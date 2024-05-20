package util

import (
	"fmt"
	"heterflow/pkg/codeaid/def"
	"math"
)

// 计算两个向量的相似度，先将向量的各个维度转换为比率，然后通过反比例函数算出相应值
func VectorApproximationRate(vectorx, vectory []int) (float64, error) {
	length := len(vectorx)
	ratio1 := make([]float64, length)
	ratio2 := make([]float64, length)
	total1 := vectorx[0]
	total2 := vectory[0]
	for i := 0; i < length; i++ {
		ratio1[i] = float64(vectorx[i]) / float64(total1)
		ratio2[i] = float64(vectory[i]) / float64(total2)
	}
	sim := 0.0
	for i := 0; i < length; i++ {
		// fmt.Println(sim)
		sim += 1 / (Absdif(ratio1[i], ratio2[i]) + def.Delta)
	}
	// fmt.Println("Pearson ", pear, "CosineSimilarity ", cos, "Sim", sim)
	sim = sim / (float64(length) * (1 / def.Delta))
	return sim, nil
}

// 计算皮尔逊相关系数
func Pearson(vectorx, vectory []int) (float64, error) {
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
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
func dotProduct(vectorx, vectory []int) (float64, error) {
	var product float64
	if len(vectorx) != len(vectory) {
		return 0, fmt.Errorf("slices x and y have different lengths")
	}
	for i, value := range vectorx {
		product += float64(value) * float64(vectory[i])
	}
	return product, nil
}

// vectorNorm 计算向量的范数（长度）
func vectorNorm(vectorx []int) float64 {
	var sum float64

	for _, value := range vectorx {
		sum += float64(value) * float64(value)
	}
	return math.Sqrt(sum)
}

// cosineSimilarity 计算两个向量的余弦相似度
func CosineSimilarity(v1, v2 []int) (float64, error) {
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
