package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	count := 1000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)

	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}

	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintln("user_%d", i))
			fmt.Println(".")
		}(i)
	}
	gw.Wait()
	fmt.Println("Done")
}

/*
func main() {
	count := 1000
	u := UserAges{ages: map[string]int{}}
	for i := 0; i < count; i++ {
		u.Add(fmt.Sprint("user_%d", i), i)
	}

	for i := 0; i < count; i++ {
		u.Get(fmt.Sprintln("user_%d", i))
		fmt.Println(".")
	}
}
*/
