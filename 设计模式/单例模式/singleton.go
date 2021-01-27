package singleton

import "sync"

// 饿汉式
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = new(Singleton)
}

func GetInstance() *Singleton {
	return singleton
}

// 懒汉式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazySingleInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = new(Singleton)
		})
	}
	return lazySingleton
}
