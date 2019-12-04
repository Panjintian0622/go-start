package slice_test
import("testing")

func TestSliceInit(t *testing.T)  {
  var s0 [] int
  t.Log(len(s0),cap(s0))//0 0
  s0 = append(s0,1)
  t.Log(len(s0),cap(s0))//1 1
  s1:= []int{1,2,3,4}
  t.Log(len(s1),cap(s1))//4 4

  s2:=make([]int ,3,5)
  t.Log(len(s2),cap(s2))//3,5
  s2 = append(s2,1)
  t.Log(s2[0],s2[1],s2[2],s2[3])//0 0 0 1
  t.Log(len(s2),cap(s2))//4 5
}

func TestSliceGrowing(t *testing.T)  {
  s:=[]int{}
  for i:=0;i<10;i++{
    s= append(s,i)
    t.Log(len(s),cap(s))
  }
  //0 1
  //2 2
  //3 4
  //4 4
  //5 8
  //6 8
  //7 8
  //8 8
  //9 16
  //10 16
}

func TestSliceShareMemory(t *testing.T)  {
  year:=[]string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
  Q2:=year[3:6]
  t.Log(Q2,len(Q2),cap(Q2))//3 9
  summer:=year[5:8]
  t.Log(summer,len(summer),cap(summer))//5 7
  summer[0]="unknow"
  t.Log(Q2)//Apr May unknow
  t.Log(year)

}

func TestSliceComping(t *testing.T)  {
  // a:=[]int{1,2,3,4}
  // b:=[]int{1,2,3,4}
  // if a==b { 不可进行比较
  //   t.Log()
  // }
}
