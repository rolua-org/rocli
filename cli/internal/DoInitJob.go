package internal

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed tpl/init/main.lua
var tpl_main_lua string

//go:embed tpl/init/project.json
var tpl_project_json []byte

func DoInitJob(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("未指定目标目录")
		return
	}

	proj := args[0]
	ignore := filepath.Join(proj, "__ignore__")

	if err := os.MkdirAll(filepath.Join(ignore, "compiler"), 0755); err != nil {
		panic(err)
	}

	t, err := template.New("tpl_main_lua").Parse(tpl_main_lua)
	if err != nil {
		panic(err)
	}

	luaf, err := os.Create(filepath.Join(proj, "main.lua"))
	if err != nil {
		panic(err)
	}

	defer luaf.Close()

	execErr := t.Execute(luaf, map[string]string{
		"ProjectName": proj,
	})

	if execErr != nil {
		panic(execErr)
	}

	luaf.Sync()

	projf, err := os.Create(filepath.Join(proj, "project.json"))
	if err != nil {
		panic(err)
	}

	defer projf.Close()

	if _, err := projf.Write(tpl_project_json); err != nil {
		panic(err)
	}
}
