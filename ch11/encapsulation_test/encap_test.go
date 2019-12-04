package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct{
	Id string
	Name string
	Age int
}
func(e Employee) String() string{
fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}
//func(e *Employee) String() string{
//	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
//}

func TestCreateEmployeeObj(t *testing.T){
	e:=Employee{"0","Bob",20}
	e1:= Employee{Name:"Mike",Age:20}
	e2:=new(Employee)//返回指针
	e2.Id="2"
	e2.Age=22
	e2.Name="Rose"
	t.Log(e)//{0 Bob 20}
	t.Log(e1)//{ Mike 20}
	t.Log(e2)//&{2 Rose 22}
	t.Logf("e is %T",e)//e is encapsulation_test.Employee
	t.Logf("e is %T",e2)//e is *encapsulation_test.Employee
}



func TestStructOperations(t *testing.T){
	e:=Employee{"0","Bob",20}
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}