package slicer

import (
	"bytes"
	"fmt"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/util"
	"os/exec"
	"strings"
)

func syscallAndLibs(filePath string) util.VertexSet[string, struct{}] {
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

func sharedLibs(filePath string) util.VertexSet[string, string] {
	cmd := exec.Command("bash", "-c", def.CommendOfSharedLibs+filePath)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		// fmt.Println("ldd", stderr.String())
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
			// fmt.Println("|"+string(after)+"|", "|"+string(path)+"|", found)
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
