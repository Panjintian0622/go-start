package type_test
import(
  "testing"
)
type MyInt int64
func TestImplicit(t *testing.T)  {
  var a int32 =1
  var b int64
  // b= a
 b = int64 (a)
 t.Log(b)
 var c MyInt
 c =MyInt(b)
 t.Log(a,b,c)
}

func TestPoint(t *testing.T){
  a:=1
  aPtr := &a
  // aPtr = aPtr+1 不支持运算
  t.Logf("%T %T",a,aPtr)//int *int
}

func TestString(t *testing.T){
  var s string
  t.Log("*"+s+"*")//**
  t.Log(len(s))//0
  // if s == ""{
  //
  // }
}
