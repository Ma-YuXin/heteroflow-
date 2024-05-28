package data

import (
	"encoding/json"
	"fmt"
	"heterflow/pkg/codeaid/slicer"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// Define structs based on the JSON structure
type Instruction struct {
	Inst    string `json:"inst"`
	Op      string `json:"op"`
	Io      string `json:"io"`
	AltForm bool   `json:"altForm,omitempty"`
}

type Category struct {
	Category     string        `json:"category"`
	Instructions []Instruction `json:"data"`
}

type Instructions struct {
	Instructions []Category `json:"instructions"`
}

func PreprocessAssemblyFiles(root string) {
	// 使用带缓冲区的通道来限制并发数量
	sem := make(chan struct{}, 16)
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

func main() {
	// Open the JSON file
	file, err := os.Open("cmp.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the JSON file
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Unmarshal the JSON data
	var instructions Instructions
	err = json.Unmarshal(byteValue, &instructions)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	// Print the instructions to verify the deserialization
	for _, category := range instructions.Instructions {
		fmt.Printf("Category: %s\n", category.Category)
		for _, instruction := range category.Instructions {
			fmt.Printf("  Instruction: %s, Op: %s, Io: %s, AltForm: %t\n", instruction.Inst, instruction.Op, instruction.Io, instruction.AltForm)
		}
	}
}
