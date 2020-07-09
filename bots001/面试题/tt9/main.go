package main

import (
	"sync"
	_"github.com/axgle/mahonia"
)

//  下面的迭代方法有什么问题
//  可以换成有缓冲的通道
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for _, value := range set.s {
			ch <- value
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {

}
