package assemblyslicer

import (
	"heterflow/pkg/codeaid/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DynamicLibs(filePath string) util.VertexSet[string, struct{}] {
	cmd := exec.Command("bash", "-c", "nm -D "+filePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%q", string(out))
	sets := strings.Split(string(out), "\n")
	res := make(util.VertexSet[string, struct{}], len(sets))
	for _, v := range sets {
		if len(v) == 0 {
			continue
		}
		res[v[19:]] = struct{}{}
	}
	// for k := range res {
	// 	fmt.Println(k, len(k))
	// }
	return nil
}

func RedirctedassembleToFile(path string) (string, string) {
	filenameWithExt := filepath.Base(path)
	filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(filenameWithExt))
	cmd := exec.Command("bash", "-c", "objdump -M intel -d "+path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(filename)
	outpath := "/mnt/data/nfs/myx/tmp/assem/" + filename + ".txt"
	err = os.WriteFile(outpath, out, 0644)
	if err != nil {
		// 如果出现错误，记录错误并退出
		log.Fatal(err)
	}
	return outpath, filename
}
