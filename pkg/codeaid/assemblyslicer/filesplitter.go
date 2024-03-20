package assemblyslicer

import (
	"bufio"
	"fmt"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/codeaid/tools"
	"heterflow/pkg/logger"
	"os"
	"time"
)

type Config struct {
	Graph        graph.Graph
	extract      Extract
	FileFeatures graph.Features
}

func NewConfig() *Config {
	return &Config{
		Graph:   graph.NewGraph(graph.Undirected),
		extract: &IntelExtract{},
	}
}

func (c *Config) SegmentFile(filepath string) {
	start := time.Now()
	defer func() {
		fmt.Println("used action ", tools.UsedAction)
		fmt.Println("missed action", tools.MissedAction)
		fmt.Println("total time used", time.Since(start))
	}()
	f, err := os.Open(filepath)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	c.extract.removeleading(scan)
	filefeature := c.readSegment(scan)
	// root := c.build.GraphGraphFromRoot("main@@Base-0x50")
	// c.build.AllFunctionName()
	// c.build.BFS(root)
	// roots := c.Graph.BuildGraph()
	// filefeature.ControlFlowGraphRoots = roots
	c.FileFeatures = filefeature
	// graph.FuncCalltimes(c.Graph)
	// for _, node := range roots {
	// 	graph.BFS(node.FuncFeatures)
	// 	fmt.Println()
	// 	fmt.Println(`***************************************************`)
	// 	fmt.Println()
	// }
}
func (c *Config) readSegment(scan *bufio.Scanner) graph.Features {
	filefeature := graph.NewFeatures(graph.ProgrammerFeature)
	for {
		ff, isEnd := c.extract.segmentFeatures(scan)
		if v, ok := ff.(*graph.Node); ok {
			c.Graph.AddNode(v)
		}

		filefeature.AddInfo(ff)
		if isEnd {
			break
		}
	}
	return filefeature
}
