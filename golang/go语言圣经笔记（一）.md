# golang圣经 阅读笔记（一） --- 一到五章的笔记


golang是21世纪的C语言，执行速度很快，编译型语言

# 1. 入门

> go run test.go 编译，链接，执行一次
> go build test.go 构建成一个EXE
> go fmt == gofmt 自动格式化 该目录下的所有.go文件

安装goimoprts 插件，能够根据代码需要 ，自动地添加或者删除`import`声明
> go get golang.org/x/tools/cmd/goimports //来安装

**go采用类似Java的包管理 措施 但是 main包比较特殊，他定义了一个独立可执行的程序，而不是一个库，在main里面的mian函数也很特殊，它是整个程序的入口。**

实际上，编译器会主动把特定符号后的换行符转换为分号, 因此换行符添加的位置会影响Go代码的正确解析。举个例子, 函数的左括号{必须和func函数声明在同一行上, 且位于末尾，不能独占一行，而在表达式x + y中，可在+后换行，不能在+前换行。

**golang 中 会给每一个变量附上默认值 ,下面是我们常用的两种定义变量的方法**

> var a int == var a int = 0 
> a := 0

os.Args
```go
// 获取命令行参数
package main

import (
	"os"
	"fmt"we
)

func main() {
	params := os.Args[1:len(os.Args)]
	fmt.Println(params)
}



package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	counts := make(map[string]int) // 储存 每个文件 line-n 的字典
	files := os.Args[1:]  // 文件名组成的数组

	if len(files) == 0 {
		countLines(os.Stdin,counts)
	} else{
		for _,file := range files {
			fp,err := os.Open(file)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err)
			} else{
				countLines(fp,counts)
			}
			fp.Close()
		}
	}

	for line,n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}


func countLines(fp *os.File,counts map[string]int){
	input := bufio.NewScanner(fp)
	for input.Scan(){
		counts[input.Text()]++
	}
}
```
基本类型 int float 数组等是 值，不是指针

go语言中 map 变量是指针

一个简单的 小爬虫
```go
	for _,url := range os.Args[1:] {
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
		}
		
		response,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdin,"fetch:%v\n",err)
			os.Exit(1)
		}
		// b,err := ioutil.ReadAll(response.Body)
		io.Copy(os.Stdout,response.Body)
		// response.Body.Close()
		// if err != nil{
		// 	fmt.Fprintf(os.Stderr,"error")
		// 	os.Exit(1)
		// }

		// fmt.Printf("%s",b)
	}
```

一个简单的web应用

```go
import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:3300",nil))

}

func handler(writer http.ResponseWriter,request *http.Request) {
	fmt.Fprintf(writer,"URL.Path = %q\n",request.URL.Path)
}
```

flag 包用来解析 命令行参数
```go
var n = flag.Bool("n",false,"omit trailing newline")
var seq = flag.String("s"," ","separator")

func main() {
	flag.Parse()
	fmt.Print( strings.Join(flag.Args(),*seq) )
	if !*n{
		fmt.Println()
	}
}

```

# 2. 基本结构

**内建函数 new**  
new将创建一个匿名变量，初始化类型的值为默认值，然后返回类型的变量地址，返回的为指针类型*T 
> p := new(int) == var p *int

变量的生命周期

type 类型名 类型

> type wendu float32
> type wendu2 float32
尽管两个 类型是一样的float32 但是不能 用来比较
```go
var a wendu
var b wendu2 
a == b // 会报错 类型不一样 
float32(a) == float32(b)  //正确
```


# 3. go语言的数据结构

- 基础数据结构   （值） 1. 数字 2. 字符串 3. 布尔型
- 复合类型		（值的组合） 1. 数组 2. 结构体
- 引用类型		引用	1. 指针 2. 切片 3. 字典map 4. 函数 5. 通道
- 接口类型		接口结构体

# 4. 基础数据类型

### 1. 整型

int8 int16 int32 int64 uint8 uint16 uint32 uint64

> byte == uint8
> rune == int32

### 2. 浮点 

float32 float64

### 3. 复数

complex64 complex128


### 4. 字符串

不可变的  `` 字面量

utf-8 变长编码 utf-32 固定长度的编码

### 5. 无类型的变量

这是GOLANG中比较特殊的


# 5. 复合类型

先说一下数组与切片的区别 一个是 复合类型，一个是引用类型
```go
func main() {
	// 数组类型 与 切片类型的区别 
	// []声明的数切片类型    [int]声明的是数组类型
	a := [4]int{1,2,3,4}
	copyA := a
	copyA[0] = 100
	fmt.Println(a)
	fmt.Println(copyA)

	b := a[:3]
	c := b
	c[0] = 100
	fmt.Println(b)
	fmt.Println(c)
	[1 2 3 4]
	[100 2 3 4]
	[100 2 3]
	[100 2 3]



	a := []int{1,2,3,4}
	copyA := a
	copyA[0] = 100
	fmt.Println(a)
	fmt.Println(copyA)

	b := a[:3]
	c := b
	c[0] = 100
	fmt.Println(b)
	fmt.Println(c)
	[100 2 3 4]
	[100 2 3 4]
	[100 2 3]
	[100 2 3]
}
```


# 6. 切片

先看一个例子

```go
	// 数组类型 与 切片类型的区别 
	// []声明的数切片类型    [int]声明的是数组类型
	a := [4]int{1,2,3,4}
	copyA := a
	copyA[0] = 100
	fmt.Println(a)
	fmt.Println(copyA)

	b := a[:3]
	c := b
	c[0] = 100
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)


	[1 2 3 4]
[100 2 3 4]
[100 2 3]
[100 2 3]
[100 2 3 4]
```

最后一个 可能跟我想象的不同，仔细看了下 ，发现时 切片从数组中切得时，切片的引用指向的是 原来的数组，所以原来的数组会改变

但是要注意，字符串是不可变的，所有切片引用会重新生成一个字符串

```go
func reverse(s []int){
	for i,j := 0,len(s)-1 ; i<j ; i,j = i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}


func main() {
	a := [...]int{1,1,4,5,6,7}
	reverse(a[:])
	fmt.Println(a)
}
```

内置的append() 函数用于向切片中添加元素
内置的copy(x,y) 函数 用于复制所有的元素 到其中


# 7. Map

map 是一个无序的键值对 字典类型，是引用类型

```go
ages := map[string]int{
		"alice":31,
		"charlie":34,
	}

fmt.Println(ages["alice"])
fmt.Println(ages["a"])  // 如果没有值，返回value的默认值 int 0 string ""      

for key,value := range ages{

} // 遍历的顺序 不确定 因为 map是无序的
```

go语言中的内建函数对于 nil 空值一般是可以安全的操作的

为了判断 map中的一个键值对是否存在，可以这样使用
> age,ok := ages["bob"] if !ok {}
常用的
> if age,ok := ages["bob"] !ok {}

go 语言中没有set集合，一般用map来代替集合
```go
seen := make(map[string]bool)
input := bufio.NewScanner(os.Stdin)
for input.Scan(){
	line := input.Text()
	if !seen[line]{
		seen[line] = true
		fmt.Println(line)
	}
}
```

# 8. 结构体


可以不用链式访问的一种用法
```go
type Point struct{
	x int
	y int
}

type Cricle struct{
	Point
	radius int
}


func main() {
	a := Cricle{Point{1,2},3}
	fmt.Println(a.x)
}
```

利用go语言中的结构体 解析 json格式的数据

> marshal n. 元帅；司仪 vt. 整理；引领；编列 vi. 排列
```go
type Movie struct{
	Title string
	Year int `json:"released"`  // 结构体成员的 元信息，关联到改成员的一些信息 
	Color bool `json:"color,omitempty"` // omitempty 表示当该成员变量的值为空或者零值的时候，不显示改变量
	Actors []string
}


func main() {
	var movies = []Movie{
		{Title:"Casablanca",Year:1888,Color:false,Actors:[]string{"paul newman"}},
		{Title:"Casablanca",Year:1888,Color:false,Actors:[]string{"paul newman"}},
		{Title:"Casablanca",Year:1888,Color:true,Actors:[]string{"paul newman"}},
		{Title:"Casablanca",Year:1888,Color:true,Actors:[]string{"paul newman"}},
	}
	data,err := json.MarshalIndent(movies,"","    ")
	if err != nil {
		log.Fatalf("err",err)
	}
	// 解析json解码字符串
	var years []struct{released int}
	if err := json.Unmarshal(data,&years);err != nil{
		log.Fatalf("error",err)
	}
	fmt.Println(years)
}
```

模板字符串
go 语言模板字符串，类似 node django servlet里面的模板 一样，有自己的语法解析

text/template
```go
//  . 代表 调用他的那个结构体
// range end 之间的内容表示循环
//  | 表示前面的内容作为后面函数的输入 ，类似unix里面的 管道符
const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------------
Number :{{.Number}}
User:{{.User.Login}}
Title:{{.Title | printf "%.64s"}} 
Age :{{.CreatedAt | daysAgo}} days
{{end}}` 

// 解析模板的套路
	report := template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo":daysAgo}).
	Parse(templ))
```
html/template与上面的类似，只不过是 增加了一些让字符串自动转义的特性

# 9. 函数

1. go语言函数可以 返回多个值
2. go语言 匿名函数 func() type {} 匿名函数
3. 闭包  类似 python里面的装饰器，里面的函数能够访问外面的 参数
```go
func squares() func() int {
	var x int
	return func() int{
		x++
		return x*x
	}
}


func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25	
}
```

可变个数 参数的语法
> func sum(vals...int) int{}
> values := []int{1,2,4}    sum(values...)

defer语句 用于在该函数执行完后调用，一般用于释放资源，关闭连接等。

```go
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s",msg)
	return func() {
		log.Printf("exit %s (%s)",msg,time.Since(start))
	}
}

func main() {
	defer trace("haha")()
}
```


# 10. golang方法 -- oop编程

golang中的面向对象编程是 通过添加方法实现的

看一个有关 例子 就大致能够明白了

原因是 golang编译器会自动的为变量 解指针，取地址等

```go
type Point struct{
	X int
	Y int
}

func (p *Point) SetX(x int) {
	p.X = x
}

func main() {
	p := Point{1,2}
	p.SetX(100)
	fmt.Println(p)
}
// 输出的结果为 {100,2}


// 但是更改一下 SetX方法
func (p Point) SetX(x int) {
	p.X = x
}

// 输出结果就变成了 {1,2}
```


内嵌类型的覆盖情况
```go
type Point struct{
	X int
	Y int
}

func (p *Point) SetX(x int) {
	p.X = x
}

type TowPoint struct{
	Point
	X int
	Y int
}
func (t *TowPoint) SetX(x int){
	t.X = x
}


func main() {
	towP := TowPoint{Point{1,2},3,4}
	towP.SetX(100)
	fmt.Println(towP.Point.X) // 1
	fmt.Println(towP.X) // 100
}
```

1. 首字母大小 --- 导出包
2. 首字母小写 --- 不导出包


# 11. 接口

golang面向对象编程的另一半就是 接口

```go
// 只要实现了 接口中定义的方法就是叫实现了这个接口
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)

var _ fmt.Stringer = &s	// ok，指针实现了接口
var _ fmt.Stringer = s // error 指针实现了接口，变量没有实现接口
```
