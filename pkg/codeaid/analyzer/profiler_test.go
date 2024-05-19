package analyzer

import (
	"fmt"
	"heterflow/pkg/codeaid/assemblyslicer"
	"testing"
)

func TestMostSimilarity(t *testing.T) {
	fmt.Println(MostSimilarProgramer("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"))
}

func TestSimilarity(t *testing.T) {
	pro := "/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl"
	path := "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl"
	proConfig := assemblyslicer.Process(pro)
	sim := Similarity(proConfig, path)
	fmt.Println(pro, path, sim)
}
func TestWriteToFile(t *testing.T) {
	// config.Process("/mnt/data/nfs/myx/tmp/app/heterflow")
	// fmt.Println(config.DynamicLib)
	assemblyslicer.Process("/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/lib/libOpenImageDenoise_device_cuda.so.2.2.2")
	// config.Process("/mnt/data/nfs/myx/tmp/app/nginx-1.22.1/objs/nginx-1.22.1")
	// config.Process("/mnt/data/nfs/myx/tmp/app/nginx-1.24.0/objs/nginx-1.24.0")
	// config.Process("/mnt/data/nfs/myx/tmp/app/nginx-1.25.5/objs/nginx-1.25.5")
	// config.Process("/mnt/data/nfs/myx/tmp/app/nginx-1.26.0/objs/nginx-1.26.0")
	// config.Process("/usr/sbin/cron")
	// config1.SegmentFile("/mnt/data/nfs/myx/tmp/dis-intel/cron-dis-intel-all")
	// // config1.DynamicLib=assemblyslicer.DynamicLibs()

	// config2 := assemblyslicer.NewConfig()
	// config2.SegmentFile("/mnt/data/nfs/myx/tmp/dis-intel/systemd-dis-intel")
	// var err error
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/tmp/bin/cron.json", config1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = assemblyslicer.WriteJSONFile("/mnt/data/nfs/myx/tmp/bin/systemd.json", config2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// func TestSimilarity(t *testing.T) {
// 	config1 := getConfig("/mnt/data/nfs/myx/tmp/bin/nginx-1.25.5.json")
// 	config2 := getConfig("/mnt/data/nfs/myx/tmp/bin/nginx-1.26.0.json")
// 	config3 := getConfig("/mnt/data/nfs/myx/tmp/bin/heterflow.json")
// 	gk1 := graph.NewGraphKernels(config1.Graph, 2)
// 	t1 := gk1.Iterator()
// 	sv1 := t1.Injection()
// 	gk2 := graph.NewGraphKernels(config2.Graph, 2)
// 	t2 := gk2.Iterator()
// 	sv2 := t2.Injection()
// 	gk3 := graph.NewGraphKernels(config3.Graph, 2)
// 	t3 := gk3.Iterator()
// 	sv3 := t3.Injection()
// 	ans, err := sv1.InnerProduct(sv2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("nginx-1.25.5 nginx-1.26.0", ans)
// 	ans, err = sv2.InnerProduct(sv3)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("nginx-1.26.0 heterflow", ans)
// 	ans, err = sv1.InnerProduct(sv3)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("nginx-1.25.5 heterflow", ans)
// 	ans, err = sv1.InnerProduct(sv1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("nginx-1.25.5 nginx-1.25.5", ans)
// }
