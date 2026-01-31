package main

import (
	"rocli/cli"
	"rocli/conf"
)

func main() {
	conf.Load()
	cli.Load()
}
