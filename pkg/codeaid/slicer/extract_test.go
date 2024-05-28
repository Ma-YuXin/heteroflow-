package slicer

import (
	"reflect"
	"testing"
)

func TestCallArgs(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			" 80487b8:	e8 25 02 00 00       	call   80489e2 <__get_pc_thunk_bx>",
			"__get_pc_thunk_bx"},
		{"case2",
			"  400e1c:	ff d0                	call   rax",
			`inst doesn't has callee `},
		{"case3",
			"  400e1e:	eb 01                	jmp    400e21 <deregister_tm_clones+0x47>",
			"deregister_tm_clones"},
		{"case4",
			"  e6:	e8 00 00 00 00       	call   eb <el_substandrun_str+0xeb>",
			"el_substandrun_str"},
		{"case5",
			"  e6:	e8 00 00 00 00       	call   eb <el_substandrun_str>",
			"el_substandrun_str"},
		{"case6",
			"	...",
			"inst is not illegal 	..."},
		{"case7",
			"  404667:       ff 15 0b c0 21 00       call   QWORD PTR [rip+0x21c00b]        # 620678 <set_dcd_val>",
			"set_dcd_val"},
	}
	ie := intelExtract{}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output, err := ie.callInstArgs(tc.input)
			if err != nil {
				output = err.Error()
			}
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}
func TestFunctionName(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"0000000000003020 <.plt>:",
			".plt",
		},
		{"case2",
			"0000000000401289 <usage>:",
			"usage",
		},
		{"case3",
			"Disassembly of section .plt:",
			"not proper function header instruction",
		},
		{"case4",
			"0000000000000000 <el_substandrun_str>:",
			"el_substandrun_str",
		},
	}
	ie := intelExtract{}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output, err := ie.functionName(tc.input)
			if err != nil {
				output = err.Error()
			}
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}

func TestVerb(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string // 测试描述
		input    string // 输入值
		expected string // 期望结果
	}{
		{"case1",
			"	...",
			"inst is not illegal 	...",
		},
		{"case2",
			"  402090:	ff 25 b2 5f 20 00    	jmp    QWORD PTR [rip+0x205fb2]        # 608048 <snd_config_iterator_first>",
			"jmp"},
		{"case3",
			"  401860:	48 63 d0             	movsxd rdx,eax",
			"movsxd"},
		{"case4",
			"  8d:	41 51                	push   r9",
			"push"},
		{"case5",
			"  98:	e8 00 00 00 00       	call   9d <el_substandrun_str+0x9d>",
			"call"},
		{"case6",
			"  9d:	48 83 c4 10          	add    rsp,0x10",
			"add"},
		{"case7",
			" 118:	48 63 d0             	movsxd rdx,eax",
			"movsxd"},
		{"case8",
			" 80498b2:	90                   	nop",
			"nop"},
		{"case9",
			" 804b6d0:	00 00 00 ",
			"inst doesn't has verb 	00 00 00 "},
	}
	ie := intelExtract{}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output, err := ie.verb(tc.input)
			if err != nil {
				output = err.Error()
			}
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got \"%v\"",
					tc.input, tc.expected, output)
			}
		})
	}
}
