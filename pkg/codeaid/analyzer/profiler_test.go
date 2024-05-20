package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/slicer"
	"os"
	"path/filepath"
	"testing"
)

func TestMostSimilarProgramerByBinary(t *testing.T) {
	fmt.Println(MostSimilarProgramerByBinary("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"))
}
func TestMostSimilarProgramerByJson(t *testing.T) {
	fmt.Println(MostSimilarProgramerByJson("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"))
}
func TestSimilarityByCalculator(t *testing.T) {
	pro := "/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"
	path := "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl"
	proCalculator := slicer.Process(pro)
	SimilarityByCalculator(proCalculator, path)
}
func TestSimilarityByJson(t *testing.T) {
	path1 := "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"
	path2 := "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl"
	sim := SimilarityByJson(path1, path2)
	fmt.Println(path1, path2, sim)
}
func TestMostTotalInstruction(t *testing.T) {
	nowMax := 0
	err := filepath.Walk(def.JsonDatabase, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// fmt.Println(path)
			config := slicer.FetchCalculator(path)
			val := config.FileFeatures.Features()[0]
			if nowMax < val {
				t.Log(val)
				nowMax = val
			}
			// fmt.Println(path, sim)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	t.Log(nowMax)
}
