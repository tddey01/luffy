package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 下面代码有什么问题
//  互斥锁修改读写锁
// 更改成读写锁

type userage struct {
	ages map[string]int
	sync.RWMutex
}

func (u *userage) Add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}

func (u *userage) Get(name string) int {
	u.RLock()
	defer u.RUnlock()
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	u := userage{ages: map[string]int{}}
	wg := sync.WaitGroup{}
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go func(i int) {
			u.Add("u"+strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	// time.Sleep(time.Millisecond)
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(u.Get("u" + strconv.Itoa(i)))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}
