package prometheus

import (
	"fmt"
	"testing"
)

func TestNodeInfo(t *testing.T) {
	testCases := []struct {
		name string
		node string // 测试描述

	}{
		{
			"master-node",
			"kind-control-plane",
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := NodeInfo(tc.node)
			fmt.Printf("%+v", res)
		})
	}

}

func TestContainerInfo(t *testing.T) {
	testCases := []struct {
		name      string
		pod       string
		namespace string
		node      string
	}{
		{
			"schedule-info",
			"kube-scheduler-kind-control-plane",
			"kube-system",
			"kind-control-plane",
		},
	}
	// 迭代测试案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainerInfo(tc.pod, tc.namespace, tc.node)
			// fmt.Printf("%+v", res)
			fmt.Println(*res.Disk)
			fmt.Println()
			fmt.Println(*res.Memory)
			fmt.Println()
			fmt.Println(*res.NetWork)
			fmt.Println()
			fmt.Println(*res.Cpu)

		})
	}

}
