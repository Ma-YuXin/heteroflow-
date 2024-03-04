package app

import (
	"fmt"
	"heteroflow/pkg/codeaid/assemblyslicer"
	"os"

	"github.com/spf13/cobra"
)

func NewCodeAidCommand() *cobra.Command {
	config := assemblyslicer.NewConfig()
	cmd := &cobra.Command{
		Use:   "CodeAid",
		Long:  ` `,
		Short: ` `,
		RunE: func(cmd *cobra.Command, args []string) error {
			config.SegmentFile("/mnt/data/nfs/myx/tmp/heterflow-dis-intel")
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
