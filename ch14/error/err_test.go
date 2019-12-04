package error_test

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"testing"
)
var LessThanTwoError = errors.New("n should be not less than 2")
func GetFibonacci(n int) ([]int,error)  {
	if n<2{
		return nil,LessThanTwoError
	}
	fibList :=[]int{1,1}
	for i:=2;i<n;i++{
		fibList =append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}
func TestGetFib(t *testing.T)  {
	if v,err:=GetFibonacci(-10);err!=nil{
		t.Error(err)
	}else{
		t.Log(v)
	}

}