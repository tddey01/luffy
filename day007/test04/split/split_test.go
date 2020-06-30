package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	t.Log("测试字符串中包含分隔符的情形")
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
	}

}

// TestnoeSplit 如果字符串中不包含分隔离符， 测试结果是否正确
func TestNoneSplit(t *testing.T) {
	t.Log("如果字符串中不包含分隔离符， 测试结果是否正确")
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
	}
}

//  go test  -run None -v  单个测试用例
//  go test  -run="." -v  所有测试用例

func TestMultiSepSplit(t *testing.T) {
	got := Split("abcfbcabcd", "bc")
	want := []string{"a", "f", "a", "d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v, 实际得到:%v\n", want, got)
	}
}
