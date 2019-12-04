package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10) //将int转为字符串
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil { //将字符串转int
		t.Log(10 + i)
	}

}
