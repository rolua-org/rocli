package cli

import "fmt"

func Parse() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
