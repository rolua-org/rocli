package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"rocli/cli"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed tpl/init/main.lua
var tpl_main_lua string

func packcli() {
	cli.Init("rocli", "为 rolua 编写的纯 Go 实现的跨平台项目管理器")

	cli.Add(
		"init", "初始化 rolua 项目",
		func(cmd *cobra.Command) {},
		func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("未指定目标目录")
				return
			}

			proj := args[0]

			if err := os.MkdirAll(proj, 0755); err != nil {
				panic(err)
			}

			t, err := template.New("tpl_main_lua").Parse(tpl_main_lua)
			if err != nil {
				panic(err)
			}

			f, err := os.Create(filepath.Join(proj, "main.lua"))
			if err != nil {
				panic(err)
			}

			defer f.Close()

			execErr := t.Execute(f, map[string]string{
				"ProjectName": proj,
			})

			if execErr != nil {
				panic(execErr)
			}

			f.Sync()
		},
	)

	cli.Parse()
}
