package constant_test

import(
  "testing"
)
const(
  Monday = iota+1
  Tuesday
  Wednesday
)
const(
  Readable = 1<<iota
  Writeable
  Executable
)
func TestConstantTry(t *testing.T)  {
  t.Log(Monday,Tuesday)
}

func TestConstantTry1(t *testing.T)  {
  a:=7//0111
  a = a &^ Readable
  t.Log(a&Readable==Readable,a&Writeable ==Writeable ,a&Executable==Executable)//true true true
}
