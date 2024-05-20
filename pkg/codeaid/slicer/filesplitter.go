package slicer

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"heterflow/pkg/codeaid/cfw"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// type encodedFeature struct {
// 	// Type string          `json:"type"`
// 	Data json.RawMessage `json:"data"`
// }

type encodedGraph struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type encodedStatisticalVectors struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type configEncoded struct {
	Feature    cfw.ProgramFeatures              `json:"programfeatures"`
	Graph      encodedGraph                     `json:"graph"`
	Vector     encodedStatisticalVectors        `json:"statisticalvectors"`
	DynamicLib util.VertexSet[string, struct{}] `json:"dynamicLib"`
	GPU        bool                             `json:"gpu"`
}

type Calculator struct {
	Graph        cfw.Graph
	FileFeatures *cfw.ProgramFeatures
	DynamicLib   util.VertexSet[string, struct{}]
	Gpu          bool
	Vector       cfw.StatisticalVectors
}

var (
	extract  = intelExtract{}
	cudaFunc = map[string]struct{}{
		"cudaFree":   {},
		"cudaMemcpy": {},
		"cudaMalloc": {},
	}
)

func Process(filepath string) *Calculator {
	asmPath, filename := RedirctedassembleToFile(filepath)
	filefeature, gra := ReadEntireAssembly(asmPath)
	libs, gpu := allLibsAndGpu(filepath, gra)
	sv := cfw.FeatureVector(gra, 2)
	cal := Calculator{
		Graph:        gra,
		FileFeatures: filefeature,
		DynamicLib:   libs,
		Gpu:          gpu,
		Vector:       sv,
	}
	err := WriteCalculator(filename, cal)
	if err != nil {
		fmt.Println(err)
	}
	return &cal
}

func ReadEntireAssembly(path string) (*cfw.ProgramFeatures, cfw.Graph) {
	// start := time.Now()
	defer func() {
		// fmt.Println("used action ", tools.UsedAction)
		if len(MissedAction) != 0 {
			fmt.Println("missed action", MissedAction)
		}
		// fmt.Println("total time used", time.Since(start))
	}()
	f, err := os.Open(path)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	filenameWithExt := filepath.Base(path)
	ff, gra := readAssemblyFunction(scan)
	ff.AddInfo(filenameWithExt)
	return ff, gra
}

func readAssemblyFunction(scan *bufio.Scanner) (*cfw.ProgramFeatures, cfw.Graph) {
	gra := cfw.NewUndirectedGraph()
	extract.removeleading(scan)
	filefeature := cfw.NewProgramFeatures()
	for {
		ff, isEnd := extract.readAsmFunc(scan)
		gra.AddNode(ff)
		filefeature.AddInfo(ff)
		if isEnd {
			break
		}
	}
	return filefeature, gra

}

func allLibsAndGpu(filepath string, gra cfw.Graph) (util.VertexSet[string, struct{}], bool) {
	sharelib := sharedLibs(filepath)
	syscall := syscallAndLibs(filepath)
	gpu := isGPUUsed(gra, sharelib)
	return util.UnionKey(syscall, sharelib), gpu
}

func isGPUUsed(gra cfw.Graph, sharedlib util.VertexSet[string, string]) bool {
	for _, path := range sharedlib {
		if len(path) == 0 {
			continue
		}
		// fmt.Println(path)
		// buf.WriteString(path)
		cmd := exec.Command("grep", "-Ec", def.CudaFlags, path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			// fmt.Println("error:", err)
			continue
		}
		if !bytes.Equal(out, []byte("0\n")) {
			fmt.Printf("%q", string(out))
			// fmt.Println("out:", string(out), len(out))
			fmt.Println(cmd.String())
			return true
		}
	}
	rel := gra.Relation()
	for k := range cudaFunc {
		if _, ok := rel[k]; ok {
			return true
		}
	}
	return false
}

// 从json文件中读取Calculator,file为文件路径
func FetchCalculator(file string) *Calculator {
	cal, err := readJSONFile(file)
	if err != nil {
		fmt.Println(err)
	}
	if def.Debug {
		fmt.Println("successfully get Calculator")
	}
	return cal
}

// 将指定路径的二进制文件反汇编并将汇编代码写入到指定文件夹中
func RedirctedassembleToFile(path string) (string, string) {
	// filenameWithExt := filepath.Base(path)
	// dir := filepath.Dir(path)
	// filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(filenameWithExt))
	filename, ok := strings.CutPrefix(path, def.BinaryFilePathPrefix)
	if !ok {
		fmt.Println("base is wrong")
		return "", ""
	}
	outpath := def.BasePath + "assem/" + filename
	cmd := exec.Command("bash", "-c", def.CommendOfDisassembly+path)
	err := util.CreateDirIfNotExist(outpath)
	if err != nil {
		fmt.Println(err)
	}
	outputFile, err := os.OpenFile(outpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	// 将命令的输出和错误重定向到文件
	cmd.Stdout = outputFile
	cmd.Stderr = outputFile
	// 执行命令
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	return outpath, filename
}

// 将Calculator写入文件，filename为文件名，t为要写入的Calculator
func WriteCalculator(filename string, t Calculator) error {
	path := def.BasePath + "json/" + filename
	util.CreateDirIfNotExist(path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	enc := &configEncoded{
		Feature:    *t.FileFeatures,
		Graph:      encodedGraph{Type: "", Data: nil},
		Vector:     encodedStatisticalVectors{Type: "", Data: nil},
		DynamicLib: t.DynamicLib,
		GPU:        t.Gpu,
	}
	switch t.Graph.(type) {
	case *cfw.UndirectedGraph:
		enc.Graph.Type = "UndirectedGraph"
		enc.Graph.Data, err = json.Marshal(t.Graph)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file undir", shape)
	case *cfw.DirectedGraph:
		enc.Graph.Type = "DirectedGraph"
		enc.Graph.Data, err = json.Marshal(t.Graph)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file dir", shape)
	default:
		fmt.Println("unknown type ,can't to marsh Graph")
	}
	switch t.Vector.(type) {
	case cfw.MapStatisticalVector:
		enc.Vector.Type = "MapStatisticalVector"
		enc.Vector.Data, err = json.Marshal(t.Vector)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file undir", shape)
	case cfw.SliceStatisticalVector:
		enc.Vector.Type = "SliceStatisticalVector"
		enc.Vector.Data, err = json.Marshal(t.Vector)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("marsh to file dir", shape)
	default:
		fmt.Println("unknown type ,can't to marsh SliceStatisticalVector")
	}
	return encoder.Encode(enc)
}

func readJSONFile(filename string) (*Calculator, error) {
	data := &Calculator{}
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
	data.FileFeatures = &enc.Feature
	data.Gpu = enc.GPU
	switch enc.Graph.Type {
	case "UndirectedGraph":
		var r cfw.UndirectedGraph
		if err := json.Unmarshal(enc.Graph.Data, &r); err != nil {
			return data, err
		}
		data.Graph = &r
	case "DirectedGraph":
		var r cfw.DirectedGraph
		if err := json.Unmarshal(enc.Graph.Data, &r); err != nil {
			return data, err
		}
		data.Graph = &r
	default:
		fmt.Println("unknown type ,can't to unmarsh", enc.Graph.Type)
	}
	switch enc.Vector.Type {
	case "MapStatisticalVector":
		var r cfw.MapStatisticalVector
		if err := json.Unmarshal(enc.Vector.Data, &r); err != nil {
			return data, err
		}
		data.Vector = &r
	case "SliceStatisticalVector":
		var r cfw.SliceStatisticalVector
		if err := json.Unmarshal(enc.Vector.Data, &r); err != nil {
			return data, err
		}
		data.Vector = &r
	default:
		fmt.Println("unknown type ,can't to unmarsh", enc.Vector.Type)
	}
	return data, err
}
