package slicer

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	Process("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl")
}
func TestRedirctedassembleToFile(t *testing.T) {
	RedirctedassembleToFile("/mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/blender")
}
func TestFetchCalculator(t *testing.T){
	cal:=FetchCalculator("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/chacl")
	fmt.Println(cal.DynamicLib)
	fmt.Println(cal.FileFeatures)
	fmt.Println(cal.Gpu)
	fmt.Println(cal.Graph)
	fmt.Println(cal.Vector)
	
}	