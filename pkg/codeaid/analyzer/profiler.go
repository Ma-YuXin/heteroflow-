package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/assemblyslicer"
	"heterflow/pkg/codeaid/graph"
)

var (
	fileweight, cfwweight = 0.7, 0.3
)

func Similarity() float64 {
	// config := assemblyslicer.NewConfig()
	// config.SegmentFile("/mnt/data/nfs/myx/tmp/heterflow-dis-intel")

	// config2 := assemblyslicer.NewConfig()
	// config2.SegmentFile("/mnt/data/nfs/myx/tmp/nginx-dis-intel")
	// var err error
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json", config)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/nginx.json", config2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return 0.0

	config1, config2 := getConfig()
	filesimilarity, err := proSimilarity(config1, config2)
	if err != nil {
		panic(err)
	}
	cfwsimilarity, err := cfwSimilarity(config1, config2)
	if err != nil {
		panic(err)
	}
	return calculate(filesimilarity, cfwsimilarity)
}
func calculate(filesimilarity, cfwsimilarity float64) float64 {
	fmt.Println("filesimilarity", filesimilarity, "cfwsimilarity", cfwsimilarity)
	return calculateWeightedSum(filesimilarity, cfwsimilarity, fileweight, cfwweight)
}

func calculateWeightedSum(a, b, wa, wb float64) float64 {
	return a*wa + b*wb
}

func proSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
	cos, err := graph.CosineSimilarity(config1.FileFeatures, config2.FileFeatures)
	fmt.Println("CosineSimilarity ", cos, err)
	pear, err := graph.Pearson(config1.FileFeatures, config2.FileFeatures)
	fmt.Println("Pearson ", pear, err)

	return pear, nil
}

func cfwSimilarity(config1 *assemblyslicer.Config, config2 *assemblyslicer.Config) (float64, error) {
	pg := graph.NewProductGraph(config1.Graph, config2.Graph)
	candidates := pg.AllConnectedVertices()
	if candidates.IsEmpty() {
		return 0.0, fmt.Errorf("no connected vertics")
	}
	excluded := make(graph.VertexSet, len(candidates))
	reporter := graph.NewRepoter(graph.MaxReporter)
	// graph.BronKerbosch2aGP(pg, reporter)
	// graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
	graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
	res2 := reporter.Report(pg.Size())
	return res2, nil
}

func getConfig() (*assemblyslicer.Config, *assemblyslicer.Config) {
	// config := assemblyslicer.NewConfig()
	// //读取并分析文件
	// config.SegmentFile("/mnt/data/nfs/myx/tmp/heterflow-dis-intel")

	// config2 := assemblyslicer.NewConfig()
	// config2.SegmentFile("/mnt/data/nfs/myx/tmp/nginx-dis-intel")

	// var err error
	// // 写入文件
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json", config)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/nginx.json", config2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// 从文件中读取结构体

	config1, err := assemblyslicer.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json")
	if err != nil {
		fmt.Println(err)
	}
	config2, err := assemblyslicer.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/nginx.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(config)
	// fmt.Println("--------------------------------")
	// for name, node := range config.Graph.Relation() {
	// 	fmt.Println(name, *node)
	// }
	// for name, node := range config2.Graph.Relation() {
	// 	fmt.Println(name, *node)
	// }
	return config1, config2
}
