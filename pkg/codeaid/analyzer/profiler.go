package analyzer

import (
	"container/heap"
	"errors"
	"fmt"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/slicer"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
	"math"
	"os"
	"path/filepath"
	"sync"
)

// 定义一个函数类型，用于处理文件
type fileProcessor func(pro interface{}, path string) result

// 用来帮助在maxchan管道中传递数据
type value struct {
	result
	name string
}
type result struct {
	prosim float64
	libsim float64
	cfwsim float64
	sim    float64
}

// FloatHeap is a min-heap of float64s.
type FloatHeap []value

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i].sim < h[j].sim }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(value))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSim(maxchan chan value) []value {
	nowMax := value{}
	for v := range maxchan {
		if v.sim > nowMax.sim {
			nowMax = v
		}
		if v.sim > 0.9 {
			str := fmt.Sprintf("%v", v)
			logger.Info(str)
		}
	}
	return []value{
		nowMax,
	}
}

func maxSimN(maxchan chan value) []value {
	n := 10 // Number of largest elements to find10
	floatHeap := &FloatHeap{}
	heap.Init(floatHeap)

	for value := range maxchan {
		if floatHeap.Len() < n {
			heap.Push(floatHeap, value)
		} else if value.sim > (*floatHeap)[0].sim {
			heap.Pop(floatHeap)
			heap.Push(floatHeap, value)
		}
	}

	// Extract the elements from the heap and sort them in descending order.
	result := make([]value, floatHeap.Len())
	for i := floatHeap.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(floatHeap).(value)
	}

	return result
}

// 将已经处理好的一个程序（也就是刚提交的那一个），与pro去对比，pro将从proJsonPath文件反序列化出来
func SimilarityByCalculator(cal interface{}, proJsonPath string) result {
	calculator1 := cal.(*slicer.Calculator)
	calculator2, err := slicer.FetchCalculator(proJsonPath)
	if err != nil {
		fmt.Println(err)
		return result{}
	}
	// fmt.Println(calculator2.Vector)
	return calsimilarity(calculator1, calculator2)
}

// 将已经序列化为json的程序，与pro去对比，pro将从proJsonPath文件反序列化出来
func SimilarityByJson(proJsonPath1 interface{}, proJsonPath2 string) result {
	p1, _ := proJsonPath1.(string)
	calculator1, err := slicer.FetchCalculator(p1)
	if err != nil {
		fmt.Println(err)
		return result{}
	}
	calculator2, err := slicer.FetchCalculator(proJsonPath2)
	if err != nil {
		fmt.Println(err)
		return result{}
	}
	return calsimilarity(calculator1, calculator2)
}

func calsimilarity(calculator1, calculator2 *slicer.Calculator) result {
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
	// fmt.Printf("%-10v %-20v %-7v %-20v %-7v %-20v %-7v %-20v %-5v %-20v\n", calculator1.FileFeatures.Name(), calculator2.FileFeatures.Name(), "cfwsim", cfwsim, "libsim", libsim, "prosim", prosim, "sim", sim)
	// fmt.Printf("%-20v %-7v %-20v %-7v %-20v %-7v %-20v %-5v %-20v\n", calculator2.FileFeatures.Name(), "cfwsim", cfwsim, "libsim", libsim, "prosim", prosim, "sim", sim)
	if calculator1.Gpu != calculator2.Gpu {
		sim *= def.PercentageDecline
	}
	return result{
		prosim: prosim,
		libsim: libsim,
		cfwsim: cfwsim,
		sim:    sim,
	}
}

// 处理文件的通用函数
func processFiles(pro interface{}, maxchan chan value, processor fileProcessor) error {
	var wg sync.WaitGroup
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
			wg.Add(1)
			go func() {
				defer func() {
					<-sem
					wg.Done()
				}() // 释放槽位
				sim := processor(pro, path)
				if err != nil {
					fmt.Printf("处理文件%s时出错: %v\n", path, err)
					return
				}
				maxchan <- value{
					result: sim,
					name:   path,
				}
			}()
		}
		return nil
	})
	wg.Wait() // 等待所有协程完成
	return err
}

func MostSimilarMatrix() {
	sem := make(chan struct{}, def.MaxGoroutines)
	filepath.Walk(def.JsonDatabase, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// fmt.Println(path)
			sem <- struct{}{}
			go func() {
				defer func() { <-sem }() // 释放槽位
				ans := MostSimilarProgramerByJson(path)
				fmt.Print(path, " {")
				for _, p := range ans {
					fmt.Printf("%+v ", p)
				}
				fmt.Println("}")
			}()
		}
		return nil
	})
}

// 将pro与json库中的程序做对比，找出最相似的那个,这里的pro是指已经序列化好的程序json路径
func MostSimilarProgramerByJson(pro string) []value {
	maxchan := make(chan value, def.MaxGoroutines)
	go func() {
		err := processFiles(pro, maxchan, SimilarityByJson)
		if err != nil {
			panic(err)
		}
	}()
	return maxSimN(maxchan)
}

// 将pro与json库中的程序做对比，找出最相似的那个,这里的pro是指二进制程序的路径
func MostSimilarProgramerByBinary(pro string) []value {
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
	// itslen := util.IntersectionLen(calculator1.DynamicLib, calculator2.DynamicLib)
	// var dymiclibsimilarity float64
	// if itslen != 0 {
	// 	c1percent := float64(len(calculator1.DynamicLib)) / float64(itslen)
	// 	c2percent := float64(len(calculator2.DynamicLib)) / float64(itslen)
	// 	// fmt.Println(itslen, c1percent, c2percent)
	// 	dymiclibsimilarity = 1 / (util.Absdif(c1percent, c2percent) + 1)
	// }
	dymiclibsimilarity := float64(util.IntersectionLen(calculator1.DynamicLib, calculator2.DynamicLib)) / float64(len(util.Union(calculator1.DynamicLib, calculator2.DynamicLib)))
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
