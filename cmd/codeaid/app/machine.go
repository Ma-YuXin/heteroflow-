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
			fmt.Println(analyzer.Similarity())
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
