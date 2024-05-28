package app

import (
	"fmt"
	"heterflow/pkg/codeaid/analyzer"
	"os"

	"github.com/spf13/cobra"
)

func NewCodeAidCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "CodeAid",
		Long:  ` `,
		Short: ` `,
		RunE: func(cmd *cobra.Command, args []string) error {
			// analyzer.MostSimilarProgramerByBinary("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc")
			// fmt.Println(analyzer.MostSimilarProgramer("/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86/O0/clamav-0.101.2/clambc"))
			analyzer.PreprocessAssemblyFiles("/mnt/data/nfs/myx/tmp/datasets/Asteria-Pro/buildroot-elf-5arch/X86")
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
