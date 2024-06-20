package data

import "testing"

func TestExtract(t *testing.T) {
	// 定义测试案例结构
	testCases := []struct {
		name  string // 测试描述
		input string // 输入值
		// 期望结果
	}{
		// {"case1", "/mnt/data/nfs/myx/helloworld/cfp2017-results-20240614-040337.json"},
		// {"case2", "/mnt/data/nfs/myx/helloworld/cint2017-results-20240614-040247.json"},
		{"case3", "/mnt/data/nfs/myx/helloworld/rint2017-results-20240614-040318.json"},
	}

	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用要测试的函数并传入输入值
			FetchBenchmarkResult(tc.input)
		})
	}
}
