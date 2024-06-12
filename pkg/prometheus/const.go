package prometheus

const (
	PROM_URL = "http://127.0.0.1:9090"
)

var(
	// 机器层面:
	// //节点网络最大速率
	// node_network_speed_bytes
	// //节点网络利用率 1m表示在一分钟内的平均每秒的速率，更倾向与表达一个瞬时的值，如果希望能够表达更为平滑的值应增大时间间隔，例如1h。
	// (rate(node_network_transmit_bytes_total{device="eth0"}[1m]) + rate(node_network_receive_bytes_total{device="eth0"}[1m])) / 1250000000
	// //节点磁盘利用率
	// instance_device:node_disk_io_time_weighted_seconds:rate5m
	// //节点cpu使用率
	// instance:node_cpu:ratio{instance="kind-control-plane"}
	// //节点内存利用率
	// instance:node_memory_utilisation:ratio
	// //节点剩余内存，一个更为积极的值，它考虑了内核使用的缓存和缓冲区内存，以及一些可能回收的内存，如临时文件等。
	// node_memory_MemAvailable_bytes{instance="kind-control-plane"}
	// //节点剩余内存，一个更为保守的值，类似于htop的值
	// node_memory_MemFree_bytes{instance="kind-control-plane"}

	// 容器层面:
	// 特定容器的块设备 I/O 操作的总量
	// container_blkio_device_usage_total{ instance="172.18.0.2:10250", job="kubelet", namespace="kube-system", node="kind-control-plane", operation="Read", pod="kube-scheduler-kind-control-plane", service="kubelet"}
	// 特定容器的内存使用总量
	// container_memory_usage_bytes{pod="kube-scheduler-kind-control-plane"}
	// 使用 container_cpu_usage_seconds_total 指标来获取容器的 CPU 使用量
	// rate(container_cpu_usage_seconds_total{container_name="your_container_name", pod_name="your_pod_name"}[5m])
	// 容器cpu使用百分比
	// sum(rate(container_cpu_usage_seconds_total{container="kube-scheduler",namespace="kube-system",pod="kube-scheduler-kind-control-plane"}[5m]))  
	// sum(container_memory_usage_bytes{container="kube-scheduler",namespace="kube-system",pod="kube-scheduler-kind-control-plane"})/	sum(node_memory_MemTotal_bytes{instance="kind-control-plane"})
)