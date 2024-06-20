package app

import (
	"fmt"
	"heterflow/pkg/codeaid/analyzer"
	"heterflow/pkg/codeaid/def"
	"os"

	"github.com/spf13/cobra"
)

func NewCodeAidCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "CodeAid",
		Long:  ` `,
		Short: ` `,
		RunE: func(cmd *cobra.Command, args []string) error {
			nums := []int{3, 4}
			for j := 0; j < len(nums); j++ {
				for i := 0; i < len(def.TechInstCounterHierarchies); i++ {
					def.TechInstCounterHierarchies[i] = nums[j]
					name := fmt.Sprintf("%v", def.TechInstCounterHierarchies)
					filename := "/mnt/data/myx/heterflow/logs/result/buildroot" + name
					file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
					if err != nil {
						return err
					}
					// 改变标准输出
					os.Stdout = file
					// 改变标准错误输出
					os.Stderr = file
					analyzer.MostSimilarMatrix()
					file.Close()
				}
			}
			//
			// analyzer.PreprocessAssemblyFiles("/mnt/data/nfs/myx/tmp/datasets/gnu_debug")
			// analyzer.SimilarityByJson("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64/O0/acl-2.2.53/getfacl", "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/Os/xapian-1.4.9/xapian-delve")
			// analyzer.MostSimilarProgramerByJson("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc")
			// analyzer.MostSimilarProgramerByBinary("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc")
			// analyzer.MostSimilarProgramerByBinary("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc")
			// fmt.Println(analyzer.MostSimilarProgramer("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"))
			// analyzer.PreprocessAssemblyFiles("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86")
			// slicer.Process("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl")
			// pro := "/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"
			// path := "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/acl-2.2.53/setfacl"
			// proCalculator := slicer.Process(pro)
			// analyzer.SimilarityByCalculator(proCalculator, path)
			// err := filepath.Walk(def.JsonDatabase, func(path string, info os.FileInfo, err error) error {
			// 	if err != nil {
			// 		fmt.Printf("访问文件/目录%s时出错: %v\n", path, err)
			// 		return err
			// 	}
			// 	if !info.IsDir() {
			// 		fmt.Println(path)
			// 	}
			// 	return nil
			// })
			// if err != nil {
			// 	panic(err)
			// }
			return nil
		}}
	return cmd
}

func Run(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
