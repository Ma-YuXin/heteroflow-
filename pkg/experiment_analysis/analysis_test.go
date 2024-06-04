package experiment_analysis

import (
	"reflect"
	"testing"
)

func TestSplitline(t *testing.T) {
	testCases := []struct {
		name     string   // 测试描述
		input    string   // 输入值
		expected []string // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr {{result:{prosim:1 libsim:1 cfwsim:0.9999999999985583 sim:0.9999999999995194} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr} {result:{prosim:0.8271811037024095 libsim:0.9777777777777777 cfwsim:0.8291666666653559 sim:0.8780418493818477} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/attr-2.4.48/getfattr} {result:{prosim:0.840433926176662 libsim:1 cfwsim:0.7833333333320226 sim:0.8745890865028948} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O2/attr-2.4.48/getfattr} {result:{prosim:0.8606655304405717 libsim:1 cfwsim:0.7433333333320226 sim:0.8679996212575314} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O3/attr-2.4.48/getfattr} {result:{prosim:0.8214451926302497 libsim:0.9777777777777777 cfwsim:0.8016666666653559 sim:0.8669632123577944} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/Os/attr-2.4.48/getfattr} {result:{prosim:0.8447312600408259 libsim:0.5087719298245614 cfwsim:0.7803030303015885 sim:0.7112687400556585} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/setfattr} {result:{prosim:0.8734049195004121 libsim:0.4857142857142857 cfwsim:0.7657407407395611 sim:0.7082866486514197} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O2/acl-2.2.53/getfacl} {result:{prosim:0.8160520715787617 libsim:0.48214285714285715 cfwsim:0.8083333333320226 sim:0.7021760873512137} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/attr-2.4.48/setfattr} {result:{prosim:0.8650849960838193 libsim:0.4714285714285714 cfwsim:0.7662499999986893 sim:0.70092118917036} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/acl-2.2.53/getfacl} {result:{prosim:0.9552494573842691 libsim:0.36764705882352944 cfwsim:0.7797619047608562 sim:0.7008861403228849} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O3/gdbm-1.13/gdbm_dump} }",
			[]string{
				"{result:{prosim:1 libsim:1 cfwsim:0.9999999999985583 sim:0.9999999999995194} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr}",
				"{result:{prosim:0.8271811037024095 libsim:0.9777777777777777 cfwsim:0.8291666666653559 sim:0.8780418493818477} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/attr-2.4.48/getfattr}",
				"{result:{prosim:0.840433926176662 libsim:1 cfwsim:0.7833333333320226 sim:0.8745890865028948} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O2/attr-2.4.48/getfattr}",
				"{result:{prosim:0.8606655304405717 libsim:1 cfwsim:0.7433333333320226 sim:0.8679996212575314} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O3/attr-2.4.48/getfattr}",
				"{result:{prosim:0.8214451926302497 libsim:0.9777777777777777 cfwsim:0.8016666666653559 sim:0.8669632123577944} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/Os/attr-2.4.48/getfattr}",
				"{result:{prosim:0.8447312600408259 libsim:0.5087719298245614 cfwsim:0.7803030303015885 sim:0.7112687400556585} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/setfattr}",
				"{result:{prosim:0.8734049195004121 libsim:0.4857142857142857 cfwsim:0.7657407407395611 sim:0.7082866486514197} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O2/acl-2.2.53/getfacl}",
				"{result:{prosim:0.8160520715787617 libsim:0.48214285714285715 cfwsim:0.8083333333320226 sim:0.7021760873512137} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/attr-2.4.48/setfattr}",
				"{result:{prosim:0.8650849960838193 libsim:0.4714285714285714 cfwsim:0.7662499999986893 sim:0.70092118917036} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O1/acl-2.2.53/getfacl}",
				"{result:{prosim:0.9552494573842691 libsim:0.36764705882352944 cfwsim:0.7797619047608562 sim:0.7008861403228849} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O3/gdbm-1.13/gdbm_dump}",
			}},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			_, output := splitline(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"{result:{prosim:1 libsim:1 cfwsim:0.9999999999985583 sim:0.9999999999995194} name:/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr}",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr",
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := extract(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestCollectInfoFromDir(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected result // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/attr-2.4.48/getfattr",
			result{
				optimizationlevel: "O0",
				architecture:      "X64",
				name:              "attr-2.4.48_getfattr",
			},
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := collectInfoFromDir(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestCollectInfoFromName(t *testing.T) {
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected result // 期望结果
	}{
		{"case1",
			"/mnt/data/nfs/myx/tmp/datasets/gnu_debug/a2ps/a2ps-4.14_clang-4.0_arm_32_O0_a2ps.elf",
			result{
				optimizationlevel: "O0",
				architecture:      "arm",
				name:              "a2ps-4.14_a2ps",
				complier:          "clang-4.0",
				bitnum:            32,
			},
		},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := collectInfoFromName(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}
