package singleton

import (
	"sync"
)

var (
	once     sync.Once
	instance *Singleton
)

type Singleton struct {
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}
