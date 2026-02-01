package conf

import (
	"encoding/json"
	"os"
)

func Load() {
	mu.Lock()
	defer mu.Unlock()

	if !IsExist() {
		return
	}

	projf, err := os.Open("project.json")
	if err != nil {
		panic(err)
	}

	defer projf.Close()

	if err := json.NewDecoder(projf).Decode(&Conf); err != nil {
		panic(err)
	}
}
