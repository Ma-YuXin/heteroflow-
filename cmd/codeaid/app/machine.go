package app

import (
	"fmt"
	"heterflow/pkg/codeaid/graph"
	"os"

	"github.com/spf13/cobra"
)

func NewCodeAidCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "CodeAid",
		Long:  ` `,
		Short: ` `,
		RunE: func(cmd *cobra.Command, args []string) error {
			// config := assemblyslicer.NewConfig()
			// //读取并分析文件
			// config.SegmentFile("/mnt/data/nfs/myx/tmp/heterflow-dis-intel")

			// config2 := assemblyslicer.NewConfig()
			// config2.SegmentFile("/mnt/data/nfs/myx/tmp/nginx-dis-intel")
			// // config2.SegmentFile("/mnt/data/nfs/myx/tmp/heterflow-dis-intel")
			// cos, err := graph.CosineSimilarity(config.FileFeatures, config2.FileFeatures)
			// fmt.Println("CosineSimilarity ", cos, err)
			// pear, err := graph.Pearson(config.FileFeatures, config2.FileFeatures)
			// fmt.Println("Pearson ", pear, err)

			// var err error
			// // 写入文件
			// err = graph.WriteJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json", config.Graph)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			//从文件中读取结构体
			// config.Graph, err = graph.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// config2.Graph, err = graph.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/nginx.json")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// for name, node := range config.Graph.Relation() {
			// 	fmt.Println(name, *node)
			// }
			g1, err := graph.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/heterflow.json")
			if err != nil {
				fmt.Println(err)
			}
			g2, err := graph.ReadJSONFile("/mnt/data/nfs/myx/heterflow/cmd/codeaid/nginx.json")
			if err != nil {
				fmt.Println(err)
			}
			pg := graph.NewProductGraph(g1, g2)
			candidates := pg.AllConnectedVertices()
			if candidates.IsEmpty() {
				return nil
			}
			excluded := make(graph.VertexSet, len(candidates))
			reporter := graph.NewRepoter(graph.CollectingReporter)
			// graph.BronKerbosch2aGP(pg, reporter)
			// graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
			graph.BronKerbosch2(pg, reporter, nil, candidates, excluded)
			fmt.Println("3333333333333333333333333&&&&&&&&&&&&&&&&&&")
			reporter.Report()
			// for _, v := range reporter.Cliques {
			// 	fmt.Println(v)
			// }
			fmt.Println("444444444444444444444&&&&&&&&&&&&&&&&")
			return nil
		}}
	return cmd
}
func Run(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
