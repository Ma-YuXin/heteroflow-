package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/assemblyslicer"
	"heterflow/pkg/codeaid/definition"
	"os"
	"path/filepath"
)

func PrepareAlljsonFiles(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
			assemblyslicer.Process(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// 将pro与json库中的程序做对比，找出最相似的那个
func MostSimilarProgramer(pro string) string {
	proConfig := assemblyslicer.Process(pro)
	mostSimilarity := ""
	nowMax := 0.0
	err := filepath.Walk(definition.JsonDatabase, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// fmt.Println(path)
			sim := Similarity(proConfig, path)
			if nowMax < sim {
				mostSimilarity = path
				nowMax = sim
			}
			// fmt.Println(path, sim)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return mostSimilarity
}
