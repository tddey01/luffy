package split

import (
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
			got := Split(item.str, item.sep)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("excepted:%#v, got:%#v", item.want, got)
			}
		})
	}
}
