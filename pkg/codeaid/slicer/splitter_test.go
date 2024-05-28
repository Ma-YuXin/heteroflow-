package slicer

import (
	"heterflow/pkg/codeaid/def"
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl",
			"setfacl",
		},
		{"case2",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O0/uboot-tools-2018.07/dumpimage",
			"dumpimage",
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := Process(tc.input).FileFeatures.Name()
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %s, but got %s",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestRedirctedassembleToFile(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O3/acl-2.2.53/chacl",
			"chacl",
		},
		{"case2",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O3/atf-0.21/application_test",
			"application_test",
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, filename := RedirctedassembleToFile(tc.input)
			path := def.BasePath + "json/" + filename + ".json"
			// 调用要测试的函数并传入输入值
			cal := FetchCalculator(path)
			// fmt.Println(cal.DynamicLib)
			// fmt.Println(cal.FileFeatures)
			// fmt.Println(cal.Gpu)
			// fmt.Println(cal.Graph)
			// fmt.Println(cal.Vector)
			output := cal.FileFeatures.Name()
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %s, but got %s",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestFetchCalculatorAndPrint(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl.json",
			"setfacl",
		},
		{"case2",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl.json",
			"chacl",
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			cal := FetchCalculator(tc.input)
			// fmt.Println(cal.DynamicLib)
			// fmt.Println(cal.FileFeatures)
			// fmt.Println(cal.Gpu)
			// fmt.Println(cal.Graph)
			// fmt.Println(cal.Vector)
			output := cal.FileFeatures.Name()
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %s, but got %s",
					tc.input, tc.expected, output)
			}
		})
	}
	// cal := FetchCalculator("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl")

}
func TestFetchCalculator(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1", "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl", "chacl"},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := FetchCalculator(tc.input).FileFeatures.Name()
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
