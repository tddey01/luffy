package mylogger

import (
	"testing"
)

// 单元测试
func TestContLevel(t *testing.T) {
	t.Logf("%v %T\n", DEBUGLevel, DEBUGLevel)
	t.Logf("%v %T\n", INFOLevel, INFOLevel)
}
