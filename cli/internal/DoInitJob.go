package internal

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed tpl/init/main.lua
var tpl_main_lua string

//go:embed tpl/init/project.json
var tpl_project_json []byte

//go:embed tpl/init/ropacker-ignore
var tpl_ropacker_ignore string

func DoInitJob(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("未指定目标目录")
		return
	}

	proj := args[0]
	clidir := filepath.Join(proj, "__rocli__")

	if err := os.MkdirAll(filepath.Join(clidir, "compiler"), 0755); err != nil {
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

	err = t.Execute(luaf, map[string]string{
		"ProjectName": proj,
	})

	if err != nil {
		panic(err)
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

	t, err = template.New("tpl_ropacker_ignore").Parse(tpl_ropacker_ignore)
	if err != nil {
		panic(err)
	}

	ignoref, err := os.Create(filepath.Join(proj, "ropacker-ignore"))
	if err != nil {
		panic(err)
	}

	defer ignoref.Close()

	selfp, err := os.Executable()
	if err != nil {
		panic(err)
	}

	err = t.Execute(ignoref, map[string]string{
		"CLIName": filepath.Base(selfp),
	})

	if err != nil {
		panic(err)
	}

	ignoref.Sync()
}
