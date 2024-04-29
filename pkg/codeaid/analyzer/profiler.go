package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/assemblyslicer"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/codeaid/util"
	"math"
)

var (
	fileweight, cfwweight = 0.7, 0.3
)

func Similarity() float64 {
	// config.Process("/mnt/data/nfs/myx/heterflow/cmd/codeaid/main")
	// config1 := assemblyslicer.NewConfig()
	// config1.SegmentFile("/mnt/data/nfs/myx/tmp/dis-intel/nginx-1.25.5")

	// config2 := assemblyslicer.NewConfig()
	// config2.SegmentFile("/mnt/data/nfs/myx/tmp/dis-intel/nginx-1.26.0")
	// var err error
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/tmp/bin/nginx-1.25.5.json", config1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/tmp/bin/nginx-1.26.0.json", config2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return 0.0

	// config1, config2 := getConfig()
	// filesimilarity, err := proSimilarity(config1, config2)
	// if err != nil {
	// 	panic(err)
	// }
	// cfwsimilarity, err := cfwSimilarity(config1, config2)
	// if err != nil {
	// 	panic(err)
	// }
	// return calculate(filesimilarity, cfwsimilarity)

	config1 := FetchConfig("/mnt/data/nfs/myx/tmp/json/nginx-1.25.json")
	config2 := FetchConfig("/mnt/data/nfs/myx/tmp/json/nginx-1.26.json")
	config3 := FetchConfig("/mnt/data/nfs/myx/tmp/json/main.json")
	gk1 := graph.NewGraphKernels(config1.Graph, 2)
	t1 := gk1.Iterator()
	sv1 := t1.Injection()
	gk2 := graph.NewGraphKernels(config2.Graph, 2)
	t2 := gk2.Iterator()
	sv2 := t2.Injection()
	gk3 := graph.NewGraphKernels(config3.Graph, 2)
	t3 := gk3.Iterator()
	sv3 := t3.Injection()
	ans, err := sv1.InnerProduct(sv2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nginx-1.25.5 nginx-1.26.0", ans)
	ans, err = sv2.InnerProduct(sv3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nginx-1.26.0 heterflow", ans)
	ans, err = sv1.InnerProduct(sv3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nginx-1.25.5 heterflow", ans)
	ans, err = sv1.InnerProduct(sv1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nginx-1.25.5 nginx-1.25.5", ans)
	return 0.0
}
func calculate(filesimilarity, cfwsimilarity float64) float64 {
	fmt.Println("filesimilarity", filesimilarity, "cfwsimilarity", cfwsimilarity)
	return calculateWeightedSum(filesimilarity, cfwsimilarity, fileweight, cfwweight)
}

func calculateWeightedSum(a, b, wa, wb float64) float64 {
	return a*wa + b*wb
}

func programSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
	cos, err := graph.CosineSimilarity(config1.FileFeatures, config2.FileFeatures)
	if err != nil {
		fmt.Println(err)
	}
	pear, err := graph.Pearson(config1.FileFeatures, config2.FileFeatures)
	if err != nil {
		fmt.Println(err)
	}
	itslen := util.IntersectionLen(config1.DynamicLib, config2.DynamicLib)
	c1percent := float64(len(config1.DynamicLib)) / float64(itslen)
	c2percent := float64(len(config2.DynamicLib)) / float64(itslen)
	dymiclibsimilarity := 1 / (math.Abs(c1percent-c2percent) + 1)
	fmt.Println("Pearson ", pear, "CosineSimilarity ", cos, "dymiclibsimilarity", dymiclibsimilarity)
	return pear, nil
}

func CFGSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
	pg := graph.NewProductGraph(config1.Graph, config2.Graph)
	candidates := pg.AllConnectedVertices()
	if util.IsEmpty(candidates) {
		return 0.0, fmt.Errorf("no connected vertics")
	}
	excluded := make(util.VertexSet[graph.Vertex, struct{}], len(candidates))
	reporter := graph.NewRepoter(graph.MaxReporter)
	// graph.BronKerbosch2aGP(pg, reporter)
	// graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
	graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
	res2 := reporter.Report(pg.Size())
	return res2, nil
}

func FetchConfig(file string) *assemblyslicer.Config {
	config, err := assemblyslicer.ReadJSONFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
