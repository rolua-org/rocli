package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

const sep = "/"

var commands = make(map[string]*cobra.Command)

func Add(path, desc string, pre func(cmd *cobra.Command), run func(cmd *cobra.Command, args []string)) {
	parts := strings.Split(path, sep)
	if len(parts) == 0 {
		return
	}

	var currentParent = root
	var currentPath []string

	for i, name := range parts {
		currentPath = append(currentPath, name)
		fullPath := strings.Join(currentPath, sep)

		cmd, exists := commands[fullPath]
		if !exists {
			cmd = &cobra.Command{
				Use:   name,
				Short: desc,
			}

			if i == len(parts)-1 {
				pre(cmd)
				cmd.Run = run

			} else {
				cmd.Run = func(cmd *cobra.Command, args []string) {
					cmd.Help()
				}

			}

			currentParent.AddCommand(cmd)
			commands[fullPath] = cmd
		}

		currentParent = cmd
	}
}
