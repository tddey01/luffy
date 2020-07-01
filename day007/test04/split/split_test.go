package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	t.Log("测试字符串中包含分隔符的情形")
// 	got := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
// 	}

// }

// // TestnoeSplit 如果字符串中不包含分隔离符， 测试结果是否正确
// func TestNoneSplit(t *testing.T) {
// 	t.Log("如果字符串中不包含分隔离符， 测试结果是否正确")
// 	got := Split("a:b:c", "*")
// 	want := []string{"a:b:c"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
// 	}
// }

// //  go test  -run None -v  单个测试用例
// //  go test  -run="." -v  所有测试用例

// 分割符是多少个字符的
// func TestMultiSepSplit(t *testing.T) {
// 	got := Split("abcfbcabcd", "bc")
// 	want := []string{"a", "f", "a", "d"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
// 	}
// }

//  将多个测试用例放到一起组成 测试组
func TestMultiSepSplit(t *testing.T) {
	// 定义一个测试数据结构体
	type test struct {
		str  string
		sep  string
		want []string
	}

	// 创建一个存放所有测试用例的map
	var tests = map[string]test{
		"noraml": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   test{"a:b:c", "*", []string{"a:b:c"}},
		"multi":  test{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
		"num":    test{"1231", "1", []string{"", "23", ""}},
	}

	for name, item := range tests {
		// ret := Split(item.str, item.sep)
		// if !reflect.DeepEqual(ret, item.want) {
		// 	t.Fatalf("测试用例:%v失败 期望得到:%#v, 实际得到:%#v\n", name, item.want, ret)
		// }
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			//  测试之前做点什么
			t.Log("要开始测试拉！...")
			defer func() {
				t.Log("沙河出太阳")
			}()
			ret := Split(item.str, item.sep)
			t.Log("官大吗 SB")
			if !reflect.DeepEqual(ret, item.want) {
				t.Errorf("期望得到:%#v, 实际得到:%#v", item.want, ret)
			}
		})
	}
}

//  go test -cover -coverprofile=c.out
//  go tool cover -html=c.out

// 基准测试函数格式
func BenchmarkSplit(b *testing.B) {
	b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//  go test  -bench BenchmarkSplit   使用方法
//  go test  -bench BenchmarkSplit -benchmem  内存 优化性能

//  并行测试
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

// // 测试集的Setup与Teardown
// func setupTestCase(t *testing.T) func(t *testing.T) {
// 	t.Log("如有需要在此执行:测试之前的setup")
// 	return func(t *testing.T) {
// 		t.Log("如有需要在此执行:测试之后的teardown")
// 	}
// }

// // 子测试的Setup与Teardown
// func setupSubTest(t *testing.T) func(t *testing.T) {
// 	t.Log("如有需要在此执行:子测试之前的setup")
// 	return func(t *testing.T) {
// 		t.Log("如有需要在此执行:子测试之后的teardown")
// 	}
// }

// func TestSplit(t *testing.T) {
// 	type test struct { // 定义test结构体
// 		input string
// 		sep   string
// 		want  []string
// 	}
// 	tests := map[string]test{ // 测试用例使用map存储
// 		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
// 		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
// 		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
// 		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
// 	}
// 	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
// 	defer teardownTestCase(t)            // 测试之后执行testdoen操作

// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
// 			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
// 			defer teardownSubTest(t)           // 测试之后执行testdoen操作
// 			got := Split(tc.input, tc.sep)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
// 			}
// 		})
// 	}
// }

// 实例函数
func ExampleAdd() {
	fmt.Println(Add("官大妈", "DSB"))
	// OutPut: 官大码DSB
}
