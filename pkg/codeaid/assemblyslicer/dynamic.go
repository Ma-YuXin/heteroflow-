package assemblyslicer

import (
	"bytes"
	"fmt"
	"heterflow/pkg/codeaid/definition"
	"heterflow/pkg/codeaid/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SyscallAndLibs(filePath string) util.VertexSet[string, struct{}] {
	cmd := exec.Command("bash", "-c", "nm -D "+filePath)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("nm -D ", stderr.String())
		return util.VertexSet[string, struct{}]{}
	}
	out := stdout.Bytes()
	// fmt.Printf("%q", string(out))
	sets := strings.Split(string(out), "\n")
	res := make(util.VertexSet[string, struct{}], len(sets))
	for _, v := range sets {
		if len(v) == 0 {
			continue
		}
		strs := strings.Split(v, " ")
		res[strs[len(strs)-1]] = struct{}{}
	}
	// for k := range res {
	// 	fmt.Println(k, len(k))
	// }
	return res
}

func SharedLibs(filePath string) util.VertexSet[string, string] {
	cmd := exec.Command("bash", "-c", definition.CommendOfSharedLibs+filePath)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("ldd", stderr.String())
		return util.VertexSet[string, string]{}
	}
	out := stdout.Bytes()
	// fmt.Printf("%q", string(out))
	sets := bytes.Split(out, []byte("\n"))
	res := make(util.VertexSet[string, string], len(sets))
	for _, v := range sets {
		if len(v) == 0 {
			continue
		}
		ctrs := bytes.TrimSpace(v)
		before, after, found := bytes.Cut(ctrs, []byte(" => "))
		// fmt.Println("|"+string(before)+"|", "|"+string(after)+"|", found)
		if found {
			path := bytes.Split(after, []byte(" "))[0]
			fmt.Println("|"+string(after)+"|", "|"+string(path)+"|", found)
			// fmt.Println(string(after))
			res[string(before)] = ""
			if string(path) != "not" {
				res[string(before)] = string(path)
			}
		} else {
			lib := bytes.Split(before, []byte(" "))[0]
			// fmt.Println("|"+string(before)+"|", "|"+string(lib)+"|", found)
			// fmt.Println(string(before))
			res[string(lib)] = ""
		}
	}
	// for k, v := range res {
	// 	fmt.Println(k, v)
	// }
	return res
}

func createDirIfNotExist(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 将指定路径的二进制文件反汇编并将汇编代码写入到指定文件夹中
func RedirctedassembleToFile(path string) (string, string) {
	// filenameWithExt := filepath.Base(path)
	// dir := filepath.Dir(path)
	// filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(filenameWithExt))
	filename, ok := strings.CutPrefix(path, definition.BinaryFilePathPrefix)
	if !ok {
		// fmt.Println("base is wrong")
		// return "", ""
	}
	outpath := definition.BasePath + "assem/" + filename
	cmd := exec.Command("bash", "-c", definition.CommendOfDisassembly+path)
	err := createDirIfNotExist(outpath)
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
