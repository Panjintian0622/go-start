package condition_test
import("testing")
func TestMultSec(t *testing.T)  {
  if a:=1==1;a{
    t.Log("1==1")
  }
  // if v,err:=someFun(); err=nil{
  //   t.Log("1==1")
  // }else{
  //   t.Log("1==1")
  // }
}

func TestSwitchMultCase(t *testing.T)  {
  for i:=0;i<5;i++{
    switch i{
    case 0,2:
      t.Log("Even")
    case 1,3:
      t.Log("Odd")
    default:
      t.Log("it is not 0-3")
    }
  }
}

func TestSwitchCaseCondition(t *testing.T)  {
  for i:=0;i<5;i++{
    switch {
    case i%2==0:
      t.Log("Even")
    case i%2 ==1:
      t.Log("Odd")
    default:
      t.Log("unknow")
    }
  }
}
