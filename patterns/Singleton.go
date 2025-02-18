package patterns

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	fmt.Printf("Pointer - %v\n", &instance)
	return instance
}

func CallSingleton() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Printf("Is the same - %v\n", s1 == s2)
}
