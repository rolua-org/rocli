package conf

import "sync"

type _Conf struct {
	Scripts  map[string]string `json:"scripts"`
	Librarys []string          `json:"librarys"`
}

var Conf = _Conf{
	Scripts:  map[string]string{},
	Librarys: []string{},
}

var mu = sync.Mutex{}
