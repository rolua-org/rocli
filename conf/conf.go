package conf

import (
	"os"
	"sync"
)

type _Conf struct {
	Compilers []string          `json:"compilers"`
	Scripts   map[string]string `json:"scripts"`
	Librarys  []string          `json:"librarys"`
}

var Conf = _Conf{
	Compilers: []string{},
	Scripts:   map[string]string{},
	Librarys:  []string{},
}

var mu = sync.Mutex{}

func IsExist() bool {
	_, err := os.Stat("project.json")
	return err == nil
}
