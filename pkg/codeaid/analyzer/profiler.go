package analyzer

import (
	"errors"
	"fmt"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/slicer"
	"heterflow/pkg/codeaid/util"
	"math"
	"os"
	"path/filepath"
	"sync"
)

// 定义一个函数类型，用于处理文件
type fileProcessor func(pro interface{}, path string) float64

// 用来帮助在maxchan管道中传递数据
type value struct {
	num  float64
	name string
}

// 将已经处理好的一个程序（也就是刚提交的那一个），与pro去对比，pro将从proJsonPath文件反序列化出来
func SimilarityByCalculator(cal interface{}, proJsonPath string) float64 {
	calculator1 := cal.(*slicer.Calculator)
	calculator2 := slicer.FetchCalculator(proJsonPath)
	// fmt.Println(calculator2.Vector)
	return calsimilarity(calculator1, calculator2)
}

// 将已经序列化为json的程序，与pro去对比，pro将从proJsonPath文件反序列化出来
func SimilarityByJson(proJsonPath1 interface{}, proJsonPath2 string) float64 {
	p1, _ := proJsonPath1.(string)
	calculator1 := slicer.FetchCalculator(p1)
	calculator2 := slicer.FetchCalculator(proJsonPath2)
	return calsimilarity(calculator1, calculator2)
}
func calsimilarity(calculator1, calculator2 *slicer.Calculator) float64 {
	prosim := programmerSimilarity(calculator1, calculator2)
	if def.Debug {
		fmt.Println("successfully arrive prosim")
	}
	libsim := libSimilarity(calculator1, calculator2)
	if def.Debug {
		fmt.Println("successfully arrive libsim")
	}
	cfwsim := cfwSimilarity(calculator1, calculator2)
	if def.Debug {
		fmt.Println("successfully arrive cfwsim")
	}
	vec := []float64{prosim, libsim, cfwsim}
	sim, err := calculateWeightedSum(vec, def.Weight)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(calculator1.FileFeatures.Name(), calculator2.FileFeatures.Name())
	fmt.Printf("%-10v %-20v %-7v %-20v %-7v %-20v %-7v %-20v %-5v %-20v\n", calculator1.FileFeatures.Name(), calculator2.FileFeatures.Name(), "cfwsim", cfwsim, "libsim", libsim, "prosim", prosim, "sim", sim)
	if calculator1.Gpu != calculator2.Gpu {
		sim *= def.PercentageDecline
	}
	return sim
}

func maxSim(maxchan chan value) string {
	mostSimilarity := ""
	nowMax := 0.0
	for v := range maxchan {
		if v.num > nowMax {
			nowMax = v.num
			mostSimilarity = v.name
		}
		if v.num > 0.9 {
			// fmt.Println(v.num, v.name)
		}
	}
	return mostSimilarity
}

// 处理文件的通用函数
func processFiles(pro interface{}, maxchan chan value, processor fileProcessor) error {
	defer close(maxchan)
	sem := make(chan struct{}, def.MaxGoroutines)
	err := filepath.Walk(def.JsonDatabase, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// fmt.Println(path)
			sem <- struct{}{}
			go func() {
				defer func() { <-sem }() // 释放槽位
				sim := processor(pro, path)
				if err != nil {
					fmt.Printf("处理文件%s时出错: %v\n", path, err)
					return
				}
				maxchan <- value{
					num:  sim,
					name: path,
				}
			}()
		}
		return nil
	})
	return err
}

// 将pro与json库中的程序做对比，找出最相似的那个,这里的pro是指已经序列化好的程序json路径
func MostSimilarProgramerByJson(pro string) string {
	maxchan := make(chan value, def.MaxGoroutines)
	go func() {
		err := processFiles(pro, maxchan, SimilarityByJson)
		if err != nil {
			panic(err)
		}
	}()
	return maxSim(maxchan)
}

// 将pro与json库中的程序做对比，找出最相似的那个,这里的pro是指二进制程序的路径
func MostSimilarProgramerByBinary(pro string) string {
	proCalculator := slicer.Process(pro)
	maxchan := make(chan value, def.MaxGoroutines)
	go func() {
		err := processFiles(proCalculator, maxchan, SimilarityByCalculator)
		if err != nil {
			panic(err)
		}
	}()
	return maxSim(maxchan)
}

func programmerSimilarity(calculator1 *slicer.Calculator, calculator2 *slicer.Calculator) float64 {
	// cos, err := util.CosineSimilarity(calculator1.FileFeatures.Features(), calculator2.FileFeatures.Features())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// pear, err := util.Pearson(calculator1.FileFeatures.Features(), calculator1.FileFeatures.Features())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(calculator1.FileFeatures.Features())
	// fmt.Println(calculator2.FileFeatures.Features())
	// fmt.Println(calculator1.FileFeatures.FeaturesWeight())
	ar, err := util.VectorApproximationRate(calculator1.FileFeatures.Features(), calculator2.FileFeatures.Features(), calculator1.FileFeatures.FeaturesWeight())
	if err != nil {
		fmt.Println(err)
	}
	t1 := calculator1.FileFeatures.Features()[0]
	t2 := calculator2.FileFeatures.Features()[0]
	dim := util.Absdif(t1, t2)
	var tir float64
	if dim < 1 {
		tir = 1.0
	} else {
		tir = 1 - math.Log10(float64(dim))/math.Log10(def.MaxTotalInstDiff)
		// float64(dim)/def.MaxTotalInstDiff
	}
	// fmt.Println("dim", dim, "tir", tir, "ar", ar, math.Log10(float64(dim)), math.Log10(def.MaxTotalInstDiff))
	return def.TotalInstWeight*tir + def.ProgramFeatureWeight*ar
}

func libSimilarity(calculator1 *slicer.Calculator, calculator2 *slicer.Calculator) float64 {
	itslen := util.IntersectionLen(calculator1.DynamicLib, calculator2.DynamicLib)
	var dymiclibsimilarity float64
	if itslen != 0 {
		c1percent := float64(len(calculator1.DynamicLib)) / float64(itslen)
		c2percent := float64(len(calculator2.DynamicLib)) / float64(itslen)
		// fmt.Println(itslen, c1percent, c2percent)
		dymiclibsimilarity = 1 / (util.Absdif(c1percent, c2percent) + 1)
	}
	return dymiclibsimilarity
}

func cfwSimilarity(calculator1, calculator2 *slicer.Calculator) float64 {
	sv1 := calculator1.Vector
	sv2 := calculator2.Vector
	ans, err := sv1.InnerProduct(sv2)
	if err != nil {
		fmt.Println(err)
	}
	if def.Debug {
		fmt.Println("successfully get ans")
	}
	return ans
}

func calculateWeightedSum(vec, weight []float64) (sum float64, err error) {
	if len(vec) != len(weight) {
		return 0.0, errors.New("vector len is not equal can't calculateWeightedSum")
	}
	for i := 0; i < len(vec); i++ {
		sum += (vec[i] * weight[i])
	}
	return
}
func PreprocessAssemblyFiles(root string) {
	// 使用带缓冲区的通道来限制并发数量
	sem := make(chan struct{}, 32)
	var wg sync.WaitGroup
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			sem <- struct{}{}
			go func() {
				defer wg.Done()
				defer func() { <-sem }() // 释放槽位
				fmt.Println(path)
				slicer.Process(path)
			}()

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
