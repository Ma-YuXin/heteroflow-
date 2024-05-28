package slicer

import (
	"heterflow/pkg/codeaid/util"
	"reflect"
	"testing"
)

func TestSystemCallAndLibs(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string                           // 测试描述
		input    string                           // 输入值
		expected util.VertexSet[string, struct{}] // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/chacl",
			util.VertexSet[string, struct{}]{
				"__bss_start":         {},
				"__errno_location":    {},
				"__uClibc_main":       {},
				"__xpg_basename":      {},
				"_edata":              {},
				"_end":                {},
				"acl_check":           {},
				"acl_delete_def_file": {},
				"acl_delete_entry":    {},
				"acl_entries":         {},
				"acl_error":           {},
				"acl_free":            {},
				"acl_from_text":       {},
				"acl_get_entry":       {},
				"acl_get_file":        {},
				"acl_get_tag_type":    {},
				"acl_set_file":        {},
				"acl_to_any_text":     {},
				"closedir":            {},
				"exit":                {},
				"fprintf":             {},
				"free":                {},
				"fwrite":              {},
				"getopt":              {},
				"malloc":              {},
				"opendir":             {},
				"optind":              {},
				"printf":              {},
				"readdir64":           {},
				"setlocale":           {},
				"sprintf":             {},
				"stderr":              {},
				"strcmp":              {},
				"strerror":            {},
				"strlen":              {},
			},
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := syscallAndLibs(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestSharedLibs(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string                         // 测试描述
		input    string                         // 输入值
		expected util.VertexSet[string, string] // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/chacl",
			util.VertexSet[string, string]{
				"linux-vdso.so.1":       "",
				"libacl.so.1":           "/lib/x86_64-linux-gnu/libacl.so.1",
				"libattr.so.1":          "/lib/x86_64-linux-gnu/libattr.so.1",
				"libc.so.0":             "",
				"libc.so.6":             "/lib/x86_64-linux-gnu/libc.so.6",
				"/lib/ld64-uClibc.so.0": "/lib64/ld-linux-x86-64.so.2",
			},
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := sharedLibs(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
