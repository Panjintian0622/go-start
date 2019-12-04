
诞生2007于谷歌

软件开发的新挑战
1多核硬件架构
2超大规模分布式计算机集群
3Web模式导致的前所未有的开发规模和更新速度

简单
高效 指针
生产力 复合VS集成
docker k8s

应用程序入口
1必须是main包：package main
2必须是main方法：func main()
3文件名不一定是main.go

退出返回值
与其他语言主要区别

GO中main函数不支持任何返回值
通过os.Exit来返回状态


获取命令行参数
func main()不接收参数 通过os.Args获取参数 go run hello.go chao

编写测试程序
1源文件以_test结尾：xxx_test.go
2测试方法名以Test开头：func TestXXX(t *testing.T){...}
3必须引用testing包
使用go test 运行

变量
1 var v_name v_type  var a int
2 var v_name = value var d =true
3 v_name := value 注意 := 左侧如果没有声明新的变量，就产生编译错误
var intVal int 
intVal :=1 // 这时候会产生编译错误
intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
多变量声明
var (
    vname1 v_type1
    vname2 v_type2
)
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)
变量赋值
赋值可以自动进行类型推断
在一个赋值语句中可以对多个变量进行同时赋值
值类型和引用类型
a=b b拷贝a的内存地址，a内容发生变化，b也跟着变
常量定义
显式类型定义 const b string = "abc"
隐式类型定义： const b = "abc"
const a, b, c = 1, false, "str" //多重赋值
iota 使用时为0，每次➕1
const (
    a = iota
    b
    c
)//a=0,b=1,c=2
 const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
//0 1 2 ha ha 100 100 7 8
const (
    i=1<<iota
    j=3<<iota
    k
    l
)
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
简单表述:
i=1：左移 0 位,不变仍为 1;
j=3：左移 1 位,变为二进制 110, 即 6;
k=3：左移 2 位,变为二进制 1100, 即 12;
l=3：左移 3 位,变为二进制 11000,即 24。

数据类型

go语言不支持隐式转换
别名和原有类型也不支持隐式转换
类型预定义值
math.MaxInt64
math.MaxFloat64
math.MaxUint32
指针类型
不支持指针运算
String类型默认是空字符串，不是nil

运算符


GO语言没有前置++，--，（++a）


用==比较数组
相同维数且含有相同个数元素的数组才可以比较
每个元素都相同的才相等





&^按位清零（右边为1，结果为0，右边为0，左边是啥结果就是啥）
1 &^ 0 ---1
1 &^ 1 ---0
0 &^ 1 ---0
0 &^ 0 ---0

循环
GO语言仅支持循环关键字for
for j:=7:j<=9;j++ //不需要括号
表示while循环
while 条件循环
while(n<5)
n:=0
for n<5{
n++
fmt.Println(n)
}

无限循环
while(true)
n:=0
for{
...
}

条件语句
if没有括号
condition表达式结果必须为布尔值
支持变量赋值
if var condition; condition{}

switch
1调教表达式不限制为常数或者整数；
2单个case中，可以出现多个结果选项，使用逗号分隔；
3与C语言等规定相反，Go语言不需要用break来明确退出一个case；
4可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个if...else相同

数组的声明
var a [3] int //声明并初始化默认为零值
a[0]=1
b:=[3]int{1,2,3}//声明同时初始化
c:=[2][2]int{{1,2},{3,4}}//多维数组初始化
遍历方式for idx/*索引*/,e/*元素*/:= range arr3 for _,e:= range arr3
数组截取
a[开始索引(包含),结束索引(不包含)]
a :=[...]int{1,2,3,4,5}
a[1,2]//2
a[1,3]//2,3
a[1,lem(a)]//2,3,4,5
a[1:]//2,3,4,5
a[:3]//1,2,3

切片

切片初始化三种方式
var s0 []int
s0 = append(s0,1)
1 s:=[]int{}
2 s1 :=[]int{1,2,3}
3 s2:=make([]int ,2,4)
/*[] type,len,cap 其中len个元素会被初始化为默认零值，未初始化元素不可以访问*/
切片共享存储结构
 


数组VS切片
1容量是否可以伸缩
2是否可以进行比较

Map
Map声明
m:=map[string]int{"one":1,"two":2}
m1:=map[string]int{}
m1["one"] = 1
m2:=make(map[string]int,10/*容量*/)
//不初始化长度，不能初始化零值
map key不存在时会赋值为0
在访问的Key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在
if v,ok:=m1[3];ok{//返回值，值是否存在
    t.Logf("key 3's valur is %d",v)
  }else{
    t.Log("key 3 is not existing.")
  }
map的遍历
m1:=map[int]int{1:1,2:4,3:9}
  for k,v:=range m1{
    t.Log(k,v)
  }
Map与工厂模式
map的value可以是一个方法
与GO的Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式
  m:=map[int]func(op int)int{}
  m[1] =func (op int)int{return op}
  m[2] =func (op int)int{return op*op}
  m[3] =func (op int)int{return op*op*op}
  t.Log(m[1](2),m[2](2),m[3](2))//2 4 8
实现set map[type]bool
1元素的唯一性
2基本操作
 添加元素
 判断元素是否存在
 删除元素
 元素个数
mySet:=map[int]bool{}
  mySet[1] = true
  n:=1
  if mySet[n]{
    t.Logf("%d is existing",n)
  }else{
      t.Logf("%d is not existing",n)
  }
  mySet[3]=true
  t.Log(len(mySet))
  delete(mySet,1)

字符串
1string是数据类型，不是饮用或指针类型
2string是只读的byte slice，len函数可以他所包含的byte数
3string的byte数组可以存放任何数据

Unicode UTF8
Unicode是字符集，UTF8是unicode的存储实现（转换为字节序列的规则）


字符串的转换strconv.Atoi("10") //将字符串转int
strconv.Itoa(10) //将int转为字符串
函数
1可以有多个返回值
2所有参数都是值传递：slice，map，channel会有传引用的错觉
3函数可以作为变量的值
4函数可以作为参数和返回值
可变参数：
func Sum(ops...int) int{
   ret:=0
   for _,op:=range ops{
      ret+=op
   }
   return ret
}
defer函数 类似于finally
func TestDefer(t *testing.T){
   defer clear()
   fmt.Println("Start")
   panic("err")
}
面向对象
封装数据和行为
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
行为（方法）定义
//第一种方式在实例对应方法被调用时，实例的成员会进行复制
func(e Employee) String() string{
fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
   return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}
//通常情况下为了避免内存拷贝我们使用第二种方式
func(e *Employee) String() string{
   fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
   return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}
接口
type Programmer interface{
   WriteHelloWorld() string
}
type GoProgrammer struct {

}
func (g *GoProgrammer) WriteHelloWorld() string{
   return "fmt.Println(\"Hello World\")"
}
与其他编程语言的区别
1接口是非入侵性，实现不依赖于接口定义
2所以接口的定义可以包含在接口使用者包内

接口变量


自定义类型
type IntVonv  func(op int) int
扩展与复用
type Pet struct {
}

func (p *Pet) Speak() {
   fmt.Print("...")
}
func (p *Pet) SpeakTo(host string){
   p.Speak()
   fmt.Println(" ",host)
}
type Dog struct {
   Pet
}
func (d *Dog) Speak(){
   fmt.Println("wang")
}
//func (d *Dog) SpeakTo(host string){
// d.Speak()
// fmt.Println(" ",host)
//}
func TestDog(t *testing.T){
   dog:= new(Dog)
   dog.SpeakTo("chao")
}
空接口
口接口可以表示任何类型
通过断言来将空接口转换为制定类型
v,ok :=p.(int)//ok=true 时为转换成功

GO错误机制
1没有错误机制
2error类型实现了error接口 type error interface{Error() string}
3可以通过errors.New 来快速创建错误机制 errors.New("n must be in the range [0,1]")
panic 不可恢复的错误，退出前会执行defer指定的内容
os.Exit 退出时不会掉用defer指定的函数，退出时不输出掉用栈信息

包
package
1基本复用模块单元 以首字母大写来表明可被包外代码访问
2代码的package可以和所在的目录不一致
3同一目录里的Go代码的package要保持一致
通过go get 来获取 远程依赖
go get -u 强制从网络更新远程依赖
注意代码在Github上的组织形式，以适应go get
直接以代码路径开始不要有src

init 方法
在main被执行前，所有以来的package的init方法都会被执行
不同包的init函数按照包倒入的依赖关系决定执行顺序
每个包可以有多个init函数
包的每个源文件也可以有多个init函数，这点比较特殊


