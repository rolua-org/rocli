package cli

import "github.com/spf13/cobra"

var root *cobra.Command

func Init(name, desc string) {
	root = &cobra.Command{
		Use:   name,
		Short: desc,
		Run: func(cmd *cobra.Command, args []string) {
			root.Help()
		},
	}
}
