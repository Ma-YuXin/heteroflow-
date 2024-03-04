package assemblyslicer

import (
	"bufio"
	"fmt"
	"heteroflow/pkg/codeaid/flowmapmaker"
	"heteroflow/pkg/codeaid/util"
	"heteroflow/pkg/logger"
	"os"
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
	defer func() {
		fmt.Println("used action ", util.UsedAction)
		fmt.Println("missed action", util.MissedAction)
	}()
	f, err := os.Open(filepath)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	c.extract.removeleading(scan)
	c.readSegment(scan)
	// root := c.build.BuildGraphFromRoot("main@@Base-0x50")
	// c.build.AllFunctionName()
	// c.build.BFS(root)
	roots:=c.build.BuildGraph()
	for _,funcfeature:=range roots{
		c.build.BFS(funcfeature)
		fmt.Println()
		fmt.Println(`********************************************************************************`)
		fmt.Println()
	}

}
func (c *Config) readSegment(scan *bufio.Scanner) {
	for {
		ff, isEnd := c.extract.segmentFeatures(scan)
		c.build.AddNode(ff)
		if isEnd {
			break
		}
	}
}
