package cfw

import (
	"heterflow/pkg/codeaid/def"
	"reflect"
	"testing"
)

func TestTechInstCounterClassify(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string              // 测试描述
		input    string              // 输入值
		expected def.TechInstManager // 期望结果
	}{
		{"case1", "clflushopt", def.TechInstManager{def.GP, def.GP_EXT}},
		{"case2", "cmpnlexadd", def.TechInstManager{def.GP, def.GP_EXT}},
		{"case3", "cmaafdd", nil},
	}
	tic := techInstCounter{}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := tic.classification(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}

func TestBehaviorInstCounterClassify(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name     string                  // 测试描述
		input    string                  // 输入值
		expected def.BehaviorInstManager // 期望结果
	}{
		{"case1",
			"andq",
			def.BehaviorInstManager{def.Logical},
		},
		{"case2",
			"cltd",
			def.BehaviorInstManager{def.Arithmetic},
		},
		{"case3",
			"cmaad",
			nil,
		},
	}
	bic := behaviorInstCounter{}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			output := bic.classification(tc.input)
			// 判断输出是否符合预期
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("For input %s, expected result is %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
