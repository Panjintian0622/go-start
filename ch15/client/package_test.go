package client

import (
	"ch15/services"
	"testing"
)

func TestPackage(t *testing.T)  {
	t.Log(services.GetFibonacci(5))//GetFib必须大写开头
}
