package assemblyslicer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/codeaid/tools"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
	"io"
	"os"
	"time"
)

var (
	extract = IntelExtract{}
)

type Config struct {
	Graph        graph.Graph
	FileFeatures graph.Features
	DynamicLib   util.VertexSet[string, struct{}]
}

func (c *Config) Process(filepath string) {
	c.Graph = graph.NewGraph(graph.Undirected)
	c.DynamicLib = DynamicLibs(filepath)
	path, filename := RedirctedassembleToFile(filepath)
	c.SegmentFile(path)
	err := WriteJSONFile("/mnt/data/nfs/myx/tmp/json/"+filename+".json", *c)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Config) SegmentFile(filepath string) {
	start := time.Now()
	defer func() {
		// fmt.Println("used action ", tools.UsedAction)
		fmt.Println("missed action", tools.MissedAction)
		fmt.Println("total time used", time.Since(start))
	}()
	f, err := os.Open(filepath)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	extract.removeleading(scan)
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
		ff, isEnd := extract.segmentFeatures(scan)
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

func WriteJSONFile(filename string, t Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	enc := &configEncoded{
		Feature:    encodedFeature{Type: "", Data: nil},
		Graph:      encodedGraph{Type: "", Data: nil},
		DynamicLib: t.DynamicLib,
	}
	switch t.FileFeatures.(type) {
	case *graph.Node:
		enc.Feature.Type = "Node"
		enc.Feature.Data, err = json.Marshal(t.FileFeatures)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file node", shape)
	case *graph.Programmer:
		enc.Feature.Type = "Programmer"
		enc.Feature.Data, err = json.Marshal(t.FileFeatures)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file pro", shape)
	default:
		fmt.Println("Unknown type")
	}
	switch t.Graph.(type) {
	case *graph.UndirectedGraph:
		enc.Graph.Type = "UndirectedGraph"
		enc.Graph.Data, err = json.Marshal(t.Graph)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file undir", shape)
	case *graph.DirectedGraph:
		enc.Graph.Type = "DirectedGraph"
		enc.Graph.Data, err = json.Marshal(t.Graph)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file dir", shape)
	default:
		fmt.Println("unknown type ,can't to marsh")
	}
	return encoder.Encode(enc)
}

func ReadJSONFile(filename string) (*Config, error) {
	data := &Config{}
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var enc configEncoded
	if err := decoder.Decode(&enc); err != nil && err != io.EOF {
		return data, err
	}
	data.DynamicLib = enc.DynamicLib
	// fmt.Println("enc", enc.Feature.Type, enc.Graph.Type)
	// fmt.Println("enc", enc.Feature.Data, enc.Graph.Data)
	switch enc.Feature.Type {
	case "Node":
		var c graph.Node
		if err := json.Unmarshal(enc.Feature.Data, &c); err != nil {
			return data, err
		}
		data.FileFeatures = &c
	case "Programmer":
		var c graph.Programmer
		if err := json.Unmarshal(enc.Feature.Data, &c); err != nil {
			return data, err
		}
		// fmt.Println(c)
		data.FileFeatures = &c
	default:
		fmt.Println("unknown type ,can't to unmarsh in feature ", enc.Feature.Type)
	}

	switch enc.Graph.Type {
	case "UndirectedGraph":
		var r graph.UndirectedGraph
		if err := json.Unmarshal(enc.Graph.Data, &r); err != nil {
			return data, err
		}
		data.Graph = &r
	case "DirectedGraph":
		var r graph.DirectedGraph
		if err := json.Unmarshal(enc.Graph.Data, &r); err != nil {
			return data, err
		}
		data.Graph = &r
	default:
		fmt.Println("unknown type ,can't to unmarsh", enc.Graph.Type)
	}
	return data, err
}

type encodedFeature struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type encodedGraph struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type configEncoded struct {
	Feature    encodedFeature                   `json:"feature"`
	Graph      encodedGraph                     `json:"graph"`
	DynamicLib util.VertexSet[string, struct{}] `json:"dynamicLib"`
}
