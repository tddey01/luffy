package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 通道
// 生成者消费这模型
// 使用goroutine和channel实现一个简易的生产消费者模型

// 生成者  产生随机 math.rand
//
// 消费者 计算每个随机数的每个位的数字的和  13144134123 =

// 1 生产者  20个随机数
var itemChan chan *item
var resultChan chan *result

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

// 生产者
func producer(ch chan *item) {
	// 生产随机数
	var id int64
	for {
		id++
		number := rand.Int63() // int64整数
		tmp := &item{
			id:  id,
			num: number,
		}
		// 把随机书发送到通道
		ch <- tmp
	}
}

// 计算一个数字每个位的和
func calc(num int64) int64 {
	// 123%10=12...3 sum = 0 + 3
	// 12%10=1...2
	// 1%10=0...3
	var sum int64
	for num > 0 {
		sum = sum + num%10 // sum = 0 + 3
		num = num / 10     // num = 12
	}
	return sum
}

// 消费者
func consumer(ch chan *item, resultChan chan *result) {

	for tmp := range ch { //
		// tmp := <-ch // 结构体指针 *item
		// (*item).num // item.num
		sum := calc(tmp.num)
		retobj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retobj
	}
}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

// 打印结果
func printResuft(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id%v, num:%v sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}
}

func main() {
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)

	go producer(itemChan)
	startWorker(20, itemChan, resultChan)

	printResuft(resultChan)

	// // 打印结果
	// rand.Seed(time.Now().Unix())
	// ret := rand.Int63()    // int64正整数
	// ret1 := rand.Intn(101) // [1 101]
	// fmt.Println(ret)
	// fmt.Println(ret1)
}
