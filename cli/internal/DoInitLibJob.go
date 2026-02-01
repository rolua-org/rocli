package internal

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"rocli/conf"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed tpl/_init/lib.go
var tpl_lib_go string

func DoInitLibJob(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("未指定扩展库名")
		return
	}

	if !conf.IsExist() {
		fmt.Println("当前目录不受 rocli 管理")
		return
	}

	libname := args[0]
	libdir := filepath.Join("lib", "user")
	libpath := filepath.Join(libdir, libname+".go")

	if err := os.MkdirAll(libdir, 0755); err != nil {
		panic(err)
	}

	libf, err := os.Create(libpath)
	if err != nil {
		panic(err)
	}

	defer libf.Close()

	t, err := template.New("tpl_lib_go").Parse(tpl_lib_go)
	if err != nil {
		panic(err)
	}

	t.Execute(libf, map[string]string{
		"LibName": libname,
	})

	libf.Sync()

	fmt.Println("成功创建", libpath)
}
