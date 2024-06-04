package status

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func VirMemInfo() *mem.VirtualMemoryStat {
	// 获取 CPU 信息
	meminfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting VirtualMemory info:", err)
		return nil
	}
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", meminfo.Total, meminfo.Free, meminfo.UsedPercent)
	return meminfo
}

func SwapMemInfo() *mem.SwapMemoryStat {
	meminfo, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Error getting VirtualMemory info:", err)
		return nil
	}
	return meminfo
}

func CPUInfo() []cpu.InfoStat {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting CPU info:", err)
		return nil
	}

	fmt.Println("CPU Info:")
	for _, info := range cpuInfo {
		fmt.Printf("Model: %s, Cores: %d, MHz: %f\n", info.ModelName, info.Cores, info.Mhz)
	}

	// 获取 CPU 使用率
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
		return nil
	}

	fmt.Println("\nCPU Usage:")
	for i, p := range percent {
		fmt.Printf("CPU%d Usage: %.2f%%\n", i, p)
	}
	return cpuInfo
}

func HostInfo() *host.InfoStat {
	// 获取主机信息
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return nil
	}
	fmt.Println("\nHost Info:")
	fmt.Println("Hostname:", hostInfo.Hostname)
	fmt.Println("OS:", hostInfo.OS)
	fmt.Println("Platform:", hostInfo.Platform)
	fmt.Println("PlatformFamily:", hostInfo.PlatformFamily)
	fmt.Println("PlatformVersion:", hostInfo.PlatformVersion)
	fmt.Println("KernelVersion:", hostInfo.KernelVersion)
	return hostInfo
}
