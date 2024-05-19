package assemblyslicer

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestSystemCallAndLibs(t *testing.T) {
	SyscallAndLibs("/mnt/data/nfs/myx/heterflow/cmd/codeaid/main")
}
func TestRedirctedassembleToFile(t *testing.T) {
	RedirctedassembleToFile("/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/blender")
}

func TestSharedLibs(t *testing.T) {
	fmt.Println(SharedLibs("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/chacl"))
}
func TestUnion(t *testing.T) {
	fmt.Println("------------------------------")
	sharedlib := SharedLibs("/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/blender")
	// sharedlib := map[string]string{
	// 	"":  "/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/lib/libOpenImageDenoise_device_cuda.so.2.2.2",
	// 	"d": "/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/lib/libOpenImageDenoise_device_cpu.so.2.2.2",
	// }
	// syscall := SyscallAndLibs("/mnt/data/nfs/myx/tmp/app/heterflow")
	// total := util.UnionKey(sharelib, syscall)
	for _, path := range sharedlib {
		if len(path) == 0 {
			continue
		}
		// fmt.Println(path)
		// buf.WriteString(path)
		cmd := exec.Command("grep", "-Ec", "cudaFree|cudaMemcpy|cudaMalloc", path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			// fmt.Println("error:", err)
			continue
		}
		if !bytes.Equal(out, []byte("0\n")) {
			fmt.Printf("%q", string(out))
			// fmt.Println("out:", string(out), len(out))
			fmt.Println(cmd.String())
			fmt.Println(true)
		}
	}
	fmt.Println(false)
	// buf.WriteString("/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/lib/libOpenImageDenoise_device_cuda.so.2.2.2")
	fmt.Println("------------------------------")
}
