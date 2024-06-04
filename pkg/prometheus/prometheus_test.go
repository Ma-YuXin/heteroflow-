package prometheus

import (
	"testing"
	"time"
)

func TestRangeQuery(t *testing.T) {
	testCases := []struct {
		name  string
		query string
		start time.Time
		end   time.Time
		step  time.Duration
	}{
		{"container_cpu_usage",
			func() string {
				query := Prom_Metric("container_cpu_usage_seconds_total", "pod", "kube-apiserver-kind-control-plane")
				return Prom_Rate(query, "1m")
			}(),
			time.Now().Add(-2 * time.Hour),
			time.Now().Add(-1 * time.Hour),
			time.Minute,
		},
		{"node_cpu_usage",
			`rate(container_cpu_usage_seconds_total{pod="kube-apiserver-kind-control-plane"}[1m])`,
			time.Now().Add(-2 * time.Hour),
			time.Now().Add(-1 * time.Hour),
			time.Minute,
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := QueryRangePrometheus(tc.query, tc.start, tc.end, tc.step)
			// Print the results
			PrintRes(res)
		})
	}
}

func TestQuery(t *testing.T) {
	testCases := []struct {
		name  string // 测试描述
		query string // 输入值
	}{
		{"node_cpu_usage",
			`node:node_cpu_utilization:ratio_rate5m`,
		},
		// {"node_disk_info",
		// 	`node_disk_info`,
		// 	"",
		// 	"",
		// },
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := QueryPrometheus(tc.query)
			// Print the results
			PrintRes(res)
		})
	}
}
