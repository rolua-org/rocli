package cli

import (
	_ "embed"
	"rocli/cli/internal"

	"github.com/spf13/cobra"
)

func Load() {
	Init("rocli", "为 rolua 编写的纯 Go 实现的跨平台项目管理器")

	Add(
		"init", "初始化 rolua 项目",
		func(cmd *cobra.Command) {},
		internal.DoInitJob,
	)

	Parse()
}
