package prometheus

import (
	"fmt"

	"github.com/prometheus/common/model"
)

type BaseInfo struct {
	Disk    *model.Sample
	Memory  *model.Sample
	NetWork *model.Sample
	Cpu     *model.Sample
}

type ContainerResUsageInfo struct {
	BaseInfo
}

type NodeResUsageInfo struct {
	BaseInfo
}

// namespace="kube-system", node="kind-control-plane", operation="Total", pod="kube-apiserver-kind-control-plane",
func ContainerInfo(pod, namespace, node string) ContainerResUsageInfo {
	var diskqueryconstraint, memqueryconstraint, cpuqueryconstraint, netqueryconstraint []Constraint
	if node != "" {
		diskqueryconstraint = MakeConstraint("namespace", namespace, "node", node, "operation", "Total", "pod", pod)
		memqueryconstraint = MakeConstraint("namespace", namespace, "node", node, "pod", pod)
		cpuqueryconstraint = MakeConstraint("namespace", namespace, "node", node, "pod", pod)
		netqueryconstraint = MakeConstraint("namespace", namespace, "node", node, "pod", pod)
	} else {
		diskqueryconstraint = MakeConstraint("namespace", namespace, "operation", "Total", "pod", pod)
		memqueryconstraint = MakeConstraint("namespace", namespace, "pod", pod)
		cpuqueryconstraint = MakeConstraint("namespace", namespace, "pod", pod)
		netqueryconstraint = MakeConstraint("namespace", namespace, "pod", pod)
	}
	diskquery := Prom_Metric("container_blkio_device_usage_total", diskqueryconstraint...)
	memquery := Prom_Sum(Prom_Metric("container_memory_usage_bytes", memqueryconstraint...))
	cpuquery := Prom_Sum(Prom_Rate(Prom_Metric("container_cpu_usage_seconds_total", cpuqueryconstraint...), "5m"))
	speedquery := Prom_Metric("node_network_speed_bytes", MakeConstraint("instance", node)...)
	speedinfo := QueryPrometheus(speedquery)
	speed := float64(speedinfo[0].Value)
	netquery := fmt.Sprintf("(%s+%s)/%f", Prom_Rate(Prom_Metric("container_network_transmit_bytes_total", netqueryconstraint...), "1m"), Prom_Rate(Prom_Metric("container_network_transmit_bytes_total", netqueryconstraint...), "1m"), speed)
	// netquery := Prom_Metric("instance:node_cpu:ratio", netqueryconstraint...)
	disk := QueryPrometheus(diskquery)[0]
	mem := QueryPrometheus(memquery)[0]
	cpu := QueryPrometheus(cpuquery)[0]
	// fmt.Println(netquery)
	network := QueryPrometheus(netquery)[0]
	return ContainerResUsageInfo{
		BaseInfo: BaseInfo{
			Disk:    disk,
			Memory:  mem,
			NetWork: network,
			Cpu:     cpu,
		},
	}
}

func NodeInfo(node string) NodeResUsageInfo {
	diskquery := Prom_Metric("instance_device:node_disk_io_time_weighted_seconds:rate5m", MakeConstraint("instance", node)...)
	memquery := Prom_Metric("instance:node_memory_utilisation:ratio", MakeConstraint("instance", node)...)
	cpuquery := Prom_Metric("instance:node_cpu:ratio", MakeConstraint("instance", node)...)
	speedquery := Prom_Metric("node_network_speed_bytes", MakeConstraint("instance", node)...)
	speedinfo := QueryPrometheus(speedquery)
	speed := float64(speedinfo[0].Value)
	netquery := fmt.Sprintf("(%s+%s)/%f", Prom_Rate(Prom_Metric("node_network_transmit_bytes_total", MakeConstraint("device", "eth0", "instance", node)...), "1m"), Prom_Rate(Prom_Metric("node_network_receive_bytes_total", MakeConstraint("device", "eth0", "instance", node)...), "1m"), speed)
	disk := QueryPrometheus(diskquery)[0]
	mem := QueryPrometheus(memquery)[0]
	cpu := QueryPrometheus(cpuquery)[0]
	network := QueryPrometheus(netquery)[0]
	return NodeResUsageInfo{
		BaseInfo: BaseInfo{
			Disk:    disk,
			Memory:  mem,
			NetWork: network,
			Cpu:     cpu,
		},
	}
}
