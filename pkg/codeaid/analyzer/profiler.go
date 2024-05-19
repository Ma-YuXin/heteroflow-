package analyzer

import (
	"errors"
	"fmt"
	"heterflow/pkg/codeaid/assemblyslicer"
	"heterflow/pkg/codeaid/definition"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/codeaid/util"
	"math"
)

var (
	maxTotalInstDiff = 1000000
	wig              = []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
)

// 将已经处理好的一个程序（也就是刚提交的那一个），与pro2去对比，pro2将从proJsonPath文件反序列化出来
func Similarity(config1 *assemblyslicer.Config, proJsonPath string) float64 {
	// prepare.PrepareAlljsonFiles("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86")
	// config1 := FetchConfig("/mnt/data/nfs/myx/tmp/json/nginx-1.25.json")
	// config2 := FetchConfig("/mnt/data/nfs/myx/tmp/json/nginx-1.26.json")
	config2 := FetchConfig(proJsonPath)
	prosim := programmerSimilarity(config1, config2)
	libsim := libSimilarity(config1, config2)
	cfwsim := cfwSimilarity(config1, config2)
	vec := []float64{prosim, libsim, cfwsim}
	sim, err := calculateWeightedSum(vec, wig)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(config1.FileFeatures.Name(), proJsonPath, "cfwsim", cfwsim, "libsim", libsim, "prosim", prosim, "sim", sim)
	fmt.Printf("%-70v %-100v %-7v %-17v %-7v %-17v %-7v %-17v %-5v %-17v\n", config1.FileFeatures.Name(), proJsonPath, "cfwsim", cfwsim, "libsim", libsim, "prosim", prosim, "sim", sim)
	if config1.Gpu != config2.Gpu {
		sim *= definition.PercentageDecline
	}
	return sim
}

func programmerSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) float64 {
	// cos, err := util.CosineSimilarity(config1.FileFeatures.Features(), config2.FileFeatures.Features())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// pear, err := util.Pearson(config1.FileFeatures.Features(), config1.FileFeatures.Features())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	f1 := config1.FileFeatures.Features()
	f2 := config2.FileFeatures.Features()
	length := len(f1)
	ratio1 := make([]float64, length)
	ratio2 := make([]float64, length)
	total1 := f1[0]
	total2 := f2[0]
	for i := 0; i < length; i++ {
		ratio1[i] = float64(f1[i]) / float64(total1)
		ratio2[i] = float64(f2[i]) / float64(total2)
	}
	sim := 0.0
	for i := 0; i < length; i++ {
		// fmt.Println(sim)
		sim += 1 / (math.Abs(ratio1[i]-ratio2[i]) + definition.Delta)
	}
	// fmt.Println("Pearson ", pear, "CosineSimilarity ", cos, "Sim", sim)
	sim = sim / (float64(length) * (1 / definition.Delta))
	return sim
}

func libSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) float64 {
	itslen := util.IntersectionLen(config1.DynamicLib, config2.DynamicLib)
	var dymiclibsimilarity float64
	if itslen != 0 {
		c1percent := float64(len(config1.DynamicLib)) / float64(itslen)
		c2percent := float64(len(config2.DynamicLib)) / float64(itslen)
		// fmt.Println(itslen, c1percent, c2percent)
		dymiclibsimilarity = 1 / (math.Abs(c1percent-c2percent) + 1)
	}
	return dymiclibsimilarity
}

func cfwSimilarity(config1, config2 *assemblyslicer.Config) float64 {
	gk1 := graph.NewGraphKernels(config1.Graph, 2)
	t1 := gk1.Iterator()
	sv1 := t1.Injection()
	gk2 := graph.NewGraphKernels(config2.Graph, 2)
	t2 := gk2.Iterator()
	sv2 := t2.Injection()
	ans, err := sv1.InnerProduct(sv2)
	if err != nil {
		fmt.Println(err)
	}
	max := 0.0
	sv1.ForEach(func(pos int, val float64) {
		v, err := sv2.At(pos)
		if err != nil {
			// fmt.Println(err)
			return
		}
		if math.Abs(val-v) > max {
			max = math.Abs(val - v)
		}
	})
	// fmt.Printf("%f ", max)
	// fmt.Println(config1.FileFeatures.Name(), config2.FileFeatures.Name(), "inner Production", ans)
	// fmt.Println()
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

func FetchConfig(file string) *assemblyslicer.Config {
	config, err := assemblyslicer.ReadJSONFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

// func CFGSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
// 	pg := graph.NewProductGraph(config1.Graph, config2.Graph)
// 	candidates := pg.AllConnectedVertices()
// 	if util.IsEmpty(candidates) {
// 		return 0.0, fmt.Errorf("no connected vertics")
// 	}
// 	excluded := make(util.VertexSet[graph.Vertex, struct{}], len(candidates))
// 	reporter := graph.NewRepoter(graph.MaxReporter)
// 	// graph.BronKerbosch2aGP(pg, reporter)
// 	// graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
// 	graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
// 	res2 := reporter.Report(pg.Size())
// 	return res2, nil
// }
