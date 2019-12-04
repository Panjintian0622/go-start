package string_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化默认为零值
	s = "hello"
	t.Log(len(s)) //5
	//s[1]='3' //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Log(s)           //严
	t.Log(len(s))      //3
	s = "中"
	t.Log(len(s))  //3
	c := []rune(s) //取出unicode
	t.Log(len(c))  //1
	//t.Log("rune size:",unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0]) //中 unicode 4e2d
	t.Logf("中 UTF8 %x", s)       //中 UTF8 e4b8ad
}

//字符串的遍历
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) //中 4e2d...
	}
}
