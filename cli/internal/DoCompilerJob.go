package internal

import (
	"runtime"

	"github.com/spf13/cobra"
)

func DoCompilerJobArgs(cmd *cobra.Command) {
	cmd.Flags().StringP("repo", "r", "github.com/rolua-org/rolua", "解释器仓库")
	cmd.Flags().StringP("version", "v", "v1", "解释器版本")
	cmd.Flags().String("os", runtime.GOOS, "目标系统")
	cmd.Flags().String("arch", runtime.GOARCH, "目标架构")
	cmd.Flags().String("url", "", "解释器下载链接")
}

func DoCompilerJob(cmd *cobra.Command, args []string) {}
