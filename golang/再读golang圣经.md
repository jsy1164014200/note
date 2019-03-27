```go
// 在Go语言中，返回函数中局部变量的地址也是安全的。
//go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，当发现变量的作用域没有跑出函数范围
// ,就可以在栈上，反之则必须分配在堆。所以不用担心会不会导致memory leak，因为GO语言有强大的垃圾回收机制。
// go语言声称这样可以释放程序员关于内存的使用限制，更多的让程序员关注于程序功能逻辑本身。
// 同样：对于动态new出来的局部变量，go语言编译器也会根据是否有逃逸行为来决定是分配在堆还是栈，而不是直接分配在堆中
func returnp() *int {
   a := 1
   return &a
}

func golandzh() {
   /*在golang中不存在未被定义的值*/
   // 数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，
   // 接口或引用类型（包括slice、map、chan和函数）变量对应的零值是nil。
   // 数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。

   //TODO:var i, j, k int                 // int, int, int
   //TODO:var b, f, s = true, 2.3, "four" // bool, float64, string

   //简短变量声明左边的变量可能并不是全部都是刚刚声明的。如果有一些已经在相同的词法域声明过了
   //，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。
   // 比如err 就是长这样用的

   // 可以在内作用域声明 一个同 外作用域名字相同的变量，，，假如外面有个a，如果是 a = 1 就是给外部变量赋值，如果是 a := 1 就是声明

   // new出来的 指针 分配了地址 ，所以不是nil
   p := new(int)   // p, *int 类型, 指向匿名的 int 变量
   fmt.Println(*p) // "0"

   // 既可以用单值接受，又可以用多值接受的
   v, ok = m[key] // map lookup
   v, ok = x.(T)  // type assertion
   v, ok = <-ch   // channel receive
   a := '过'
   b := 'a'
   var c rune = 10
   fmt.Println(reflect.TypeOf(a)) // int32
   fmt.Println(reflect.TypeOf(b)) // int32
   fmt.Println(b + c)
}

//TODO:golang特殊的变量声明周期。（不是离开局部作用域就被回收，而是当访问不到就被回收）。
//Go语言的自动圾收集器是如何知道一个变量是何时可以被回收的呢？这里我们可以避开完整的技术细节，基本的实现思路是，
//从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。
//如果不存在这样的访问路径，那么说明该变量是不可达的，也就是说它是否存在并不会影响程序后续的计算结果。

// Type 声明的具有相同底层数据结构的类型，不能相互赋值，除非进行显式转换
//对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型
// （译注：如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)）。
// 只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的
// 指针类型，这些转换只改变类型而不会影响值本身。
// 类型错误只会发生在编译阶段,不会发生在运行时。

// 许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印，我们将在7.1节讲述。
type MyString string

func (c MyString) String() string { return fmt.Sprintf("%g MyString", c) }

// 包名前面的注释，是包注释，文档工具会使用它。
// 使用init()函数来 初始化包中的数据

n := 0
for range s {
    n++
}


```



## 常量

> 不能是在运行时确定的值

常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。

所有常量的运算都可以在编译期完成，这样可以减少运行时的工作，也方便其他编译优化。当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越界、任何导致无效浮点数的操作等。

常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex和unsafe.Sizeof。因为它们的值是在编译期就确定的

const中iato的一种用法

```go
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
```











```Go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]

// 多参数
appendInt(x []int, y ...int)
// 解包
append(x, x...)

```



map

禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

Map的迭代顺序是不确定的

```Go
if age, ok := ages["bob"]; !ok { /* ... */ }
```



golang中使用map来模拟set



结构体可以嵌套（设计模式）



JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。在类似的协议中，JSON并不是唯一的一个标准协议。 XML（§7.14）、ASN.1和Google的Protocol Buffers都是类似的协议，并且有各自的特色，但是由于简洁性、可读性和流行程度等原因，JSON是应用最广泛的一个。



## 函数

```go
func add(a, b int) (z int) {
   z = a + b
   return // 并且这个return不能省去。
}
但是这种方式会使得代码难以理解，少使用
```





golang中发生错误时，处理的三种策略：

1. 输出错误信息并结束程序。需要注意的是，这种策略只应在main中执行。

2. 对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。(os.exit(1))

3. 输出错误信息，但是不终止

   - ```Go
     log.Fatalf("Site is down: %v\n", err)
     ```

   - ```Go
     fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
     ```

4. 忽略错误，但是一定要记录忽略的理由

> log包中的所有函数会为没有换行符的字符串增加换行符。



golang提供了一个特殊的error类型，io.EOF表示文件读取到了末尾。



golang中的闭包概念

```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```





### 函数中捕获迭代变量（陷阱）

```go
for _, dir := range tempDirs() {
    dir := dir // declares inner dir, initialized to outer dir
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```







golang中的异常

当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。

在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理，最好的处理方式，就是使用Go的错误机制。



recover捕获panic异常，使得整个程序能继续执行

```go
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }()
    // ...parser...
}
```







方法：OOP编程才有方法

在能够给任意类型定义方法这一点上，Go和很多其它的面向对象的语言不太一样。因此在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。

指针类型不能有方法：

```go
type P *int
func (P) f() { /* ... */ } // compile error: invalid receiver type
```



我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：（因为它是分配在栈上面的）

```django
Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
```







方法表达式子：

两种

1. ```go
   p := Point{1, 2}
   q := Point{4, 6}
   
   distanceFromP := p.Distance        // method value
   fmt.Println(distanceFromP(q))      // "5"
   ```

2. ```go
   distance := Point.Distance   // method expression
   fmt.Println(distance(p, q))  // "5"
   // 第一个参数是接收器
   ```





### 接口

#### 接口值

type ，value

接口的零值就是 type为nil，value也为nil

> var w io.Write 这时type为nil
>
> w  = os.Stdout 这时type为*os.File

接口值可以使用＝＝和！＝来进行比较。两个接口值相等仅当它们都是nil值或者它们的动态类型相同并且动态值也根据这个动态类型的＝＝操作相等。因为接口值是可比较的，所以它们可以用在map的键或者作为switch语句的操作数。







## error

使用 errors.New() 或者

```go
func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```



1. 不要使用w.Write([]byte("Content-Type: "))，会在内存中产生拷贝，应该使用writeString





golang并发模型

CSP Communicating sequential processes

io以及简单的go 关键字用法

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	testSpinner()
}

func testSpinner() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFib(%d)=%d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `_\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
```





服务端

```go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close() // 忽略关闭错误
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
```



客户端

```go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```





通道

```go
//接收者收到数据发生在唤醒发送者goroutine之前（：happens before，这是Go语言并发内存模型的一个关键术语！）。
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```



range可以迭代带缓存的 channel（当channel关闭并且里面没有值了就会停止）

关闭了channel后，可以接受值，但不能发送值（会panic）

cap(chan) 获取容量 len(chan)获取现有元素



并发循环的一个注意点

```go
func makeThumbnails3(filenames []string) {
ch := make(chan struct{})
for _, f := range filenames {
	go func(f string) {  // 这个地方要传f进去就是因为f会随循环改变（闭包）
		thumbnail.ImageFile(f) // NOTE: ignoring errors
		ch <- struct{}{}
	}(f)
} /
/ Wait for goroutines to complete.
for range filenames {
	<-ch
}
}
```





一个会导致 goroutine泄露的例子

```go
func makeThumbnails4(filenames []string) error {
errors := make(chan error)
for _, f := range filenames {
go func(f string) {
_, err := thumbnail.ImageFile(f)
errors <- err
}(f)}
} f
or range filenames {
if err := <-errors; err != nil {
return err // NOTE: incorrect: goroutine leak!
}
} r
eturn nil
}
```

解决方案是采用 带缓存的channel，长度等于循环次数



地道的循环次数的做法

```
// 地道的使用循环迭代
func makeThumbnails(filenames <-chan string) int64 {
   sizes := make(chan int64)
   var wg sync.WaitGroup
   for f := range filenames {
      wg.Add(1)
      go func(f string) {
         defer wg.Done()
         thumb, err := thumbnail.ImageFile(f)
         if err != nil {
            log.Println(err)
            return
         }
         info, _ := os.Stat(thumb)
         sizes <- info.Size()
      }(f)
   }
   go func() {
      wg.Wait()
      close(sizes)
   }()
   var total int64
   for size := range sizes {
      total += size

   }
   return total
}
```





## 一个爬虫的例子

```go
// 加入n （循环次数）让主程序能正常停止运行。
func main() {
	worklist := make(chan []string)
	var n int
	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++ // 注意n不能在 goroutine中执行，并发不安全
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// 保证在同一时间 对Extract的调用不超过20 次
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
```









select用法

```go
select {
	case <- time.After(10 * time.Second):
    case <- time.NewTicker(1* time.Second):
		//
	case <- abort:
		return 
	}
	launch()
```

du的例子

```go
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 限制并发的次数
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dul: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
```

