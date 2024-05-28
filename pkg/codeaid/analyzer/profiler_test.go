package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/codeaid/slicer"
	"os"
	"path/filepath"
	"reflect"
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
	testCases := []struct {
		name     string  // 测试描述
		input1   string  // 输入值
		input2   string  // 输入值
		expected float64 // 期望结果
	}{
		// {"case1",
		// 	"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/libpfm4-4.9.0/pfmlib_amd64_fam10h.lo",
		// 	"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/clamav-0.101.2/clambc",
		// 	0.0,
		// },
		{"case2",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/getfacl",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/setfacl",
			0.0,
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := SimilarityByJson(tc.input1, tc.input2)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.name, tc.expected, output)
			}
		})
	}
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
