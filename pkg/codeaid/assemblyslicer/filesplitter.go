package assemblyslicer

import (
	"bufio"
	"fmt"
	"heterflow/pkg/codeaid/flowmapmaker"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
	"os"
	"time"
)

type Config struct {
	build   *flowmapmaker.Build
	extract Extract
}

func NewConfig() *Config {
	return &Config{
		build: &flowmapmaker.Build{
			Relations: make(map[string]*flowmapmaker.FuncFeatures),
		},
		extract: &IntelExtract{},
	}
}

func (c *Config) SegmentFile(filepath string) {
	start := time.Now()
	defer func() {
		fmt.Println("used action ", util.UsedAction)
		fmt.Println("missed action", util.MissedAction)
		fmt.Println("total time used", time.Since(start))
	}()
	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 3)
	// 		bToMb := func(alloc uint64) uint64 { return alloc / 1024 }
	// 		var m runtime.MemStats
	// 		runtime.ReadMemStats(&m)
	// 		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	// 		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	// 		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	// 		fmt.Printf("\tNumGC = %v\n", m.NumGC)
	// 	}
	// }()
	f, err := os.Open(filepath)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	c.extract.removeleading(scan)
	filefeature := c.readSegment(scan)
	// root := c.build.BuildGraphFromRoot("main@@Base-0x50")
	// c.build.AllFunctionName()
	// c.build.BFS(root)
	roots := c.build.BuildGraph()
	filefeature.ControlFlowGraphRoots = roots
	// for _, funcfeature := range roots {
	// 	flowmapmaker.BFS(funcfeature)
	// 	fmt.Println()
	// 	fmt.Println(`********************************************************************************`)
	// 	fmt.Println()
	// }
}
func (c *Config) readSegment(scan *bufio.Scanner) *flowmapmaker.FileFeatures {
	filefeature := flowmapmaker.NewFileFeatures()
	for {
		ff, isEnd := c.extract.segmentFeatures(scan)
		c.build.AddNode(ff)
		filefeature.AddFileFeatures(ff)
		if isEnd {
			break
		}
	}
	return filefeature
}
