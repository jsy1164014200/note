# golang圣经笔记（二）

## 共享变量的并发

## sync包

sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。

本包的类型的值不应被拷贝。

type Locker

type Once

```
var once sync.Once
onceBody := func() {
    fmt.Println("Only once")
}
done := make(chan bool)
for i := 0; i < 10; i++ {
    go func() {
        once.Do(onceBody)
        done <- true
    }()
}
for i := 0; i < 10; i++ {
    <-done
}
```

Output:

```
Only once
```

如果once.Do(f)被多次调用，只有第一次调用会执行f，即使f每次调用Do 提供的f值不同。需要给每个要执行仅一次的函数都建立一个Once类型的实例。

Do用于必须刚好运行一次的初始化。因为f是没有参数的，因此可能需要使用闭包来提供给Do方法调用：

```
config.once.Do(func() { config.init(filename) })
```





- func (o *Once) Do(f func())

type Mutex

- [func (m *Mutex) Lock()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#Mutex.Lock)
- func (m *Mutex) Unlock()

type RWMutex

- [func (rw *RWMutex) Lock()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#RWMutex.Lock)
- func (rw *RWMutex) Unlock()
- [func (rw *RWMutex) RLock()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#RWMutex.RLock)
- [func (rw *RWMutex) RUnlock()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#RWMutex.RUnlock)
- func (rw *RWMutex) RLocker() Locker

type Cond

- func NewCond(l Locker) *Cond
- [func (c *Cond) Broadcast()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#Cond.Broadcast)
- func (c *Cond) Signal()
- [func (c *Cond) Wait()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#Cond.Wait)

type WaitGroup

```
var wg sync.WaitGroup
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
wg.Wait()
```

- func (wg *WaitGroup) Add(delta int)
- func (wg *WaitGroup) Done()
- [func (wg *WaitGroup) Wait()](https://studygolang.com/static/pkgdoc/pkg/sync.htm#WaitGroup.Wait)

type Pool

- [func (p *Pool) Get() interface{}](https://studygolang.com/static/pkgdoc/pkg/sync.htm#Pool.Get)
- func (p *Pool) Put(x interface{})

类似一个 并发安全的 队列









## sync讲解

数据竞争的定义，因为实在太重要了：数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生。根据上述定义，有三种方式可以避免数据竞争：

第一种方法是不要去写变量。

第二种避免数据竞争的方法是，避免从多个goroutine访问变量，不要使用共享数据来通信；使用通信来共享数据。一个提供对一个指定的变量通过cahnnel来请求的goroutine叫做这个变量的监控(monitor)goroutine。例如broadcaster goroutine会监控(monitor)clients map的全部访问。

```go
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
    var balance int // balance is confined to teller goroutine
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

func init() {
    go teller() // start the monitor goroutine
}
```

第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问。这种方式被称为“互斥”

### 1. sync.Mutex 互斥锁

添加一个二元信号量

```go
var (
    sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
    balance int
)

func Deposit(amount int) {
    sema <- struct{}{} // acquire token
    balance = balance + amount
    <-sema // release token
}

func Balance() int {
    sema <- struct{}{} // acquire token
    b := balance
    <-sema // release token
    return b
}

//////////////////等同于下面的
import "sync"

var (
    mu      sync.Mutex // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance = balance + amount
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    b := balance
    return b
}
```



### 2. sync.RWMutex读写锁

```go
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}

```

### 3. 对于只初始化一次的函数 用 sync.Once

```go
func loadIcons() {
    icons = make(map[string]image.Image)
    icons["spades.png"] = loadIcon("spades.png")
    icons["hearts.png"] = loadIcon("hearts.png")
    icons["diamonds.png"] = loadIcon("diamonds.png")
    icons["clubs.png"] = loadIcon("clubs.png")
}

var loadIconsOnce sync.Once
var icons map[string]image.Image
// Concurrency-safe.
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}

```



竞争条件检查器：只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，并且会记录下每一个读或者写共享变量的goroutine的身份信息。

### 实战一个 并发的非阻塞缓存

做一个无阻塞的缓存，这种工具可以帮助我们来解决现实世界中并发程序出现但没有现成的库可以解决的问题。这个问题叫作缓存，也就是说我们需要缓存函数的返回结果，这样在对函数进行调用的时候，我们就只需要一次计算，之后只要返回计算的结果就可以了

```go
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetBody)
	for url := range incomingURLS {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Println("%s, %s, %d bytes", url, time.Since(start), len(value.([]byte)))
	}
}

```



> 这部分 有时间再看一遍



## 测试

> go test

### example test

Examples

该包还运行并验证示例代码。示例函数可以包括以 "Output：" 开头的行注释，并在运行测试时与函数的标准输出进行比较。 （比较时会忽略前导和尾随空格。）这些是一个 example 的例子：

```
func ExampleHello() {
        fmt.Println("hello")
        // Output: hello
}

func ExampleSalutations() {
        fmt.Println("hello, and")
        fmt.Println("goodbye")
        // Output:
        // hello, and
        // goodbye
}
```

"Unordered output:" 形式的注释，和 "Output:" 类似，但是能够以任意顺序匹配行：

```
func ExamplePerm() {
    for _, value := range Perm(4) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

没有输出注释的示例函数被编译但不执行。

example 声明的命名约定：包，函数 F，类型 T，类型 T 上的方法 M 依次是：

```
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

可以为 包/类型/函数/方法 提供多个 example 函数，这通过在名称上附加一个不同的后缀来实现。后缀必须是以小写字母开头。

```
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

当一个文件包含一个示例函数，同时至少一个其他函数，类型，变量或常量声明，或没有测试或基准函数时，这个测试文件作为示例存在，通常命名为 example_test.go

-run 和 -bench 命令行标志的参数是与测试名称相匹配的非固定的正则表达式。对于具有多个斜杠分隔元素（例如子测试）的测试，该参数本身是斜杠分隔的，其中表达式依次匹配每个名称元素。因为它是非固定的，一个空的表达式匹配任何字符串。例如，使用 "匹配" 表示 "其名称包含"：

```
go test -run ''      # Run 所有测试。
go test -run Foo     # Run 匹配 "Foo" 的顶层测试，例如 "TestFooBar"。
go test -run Foo/A=  # 匹配顶层测试 "Foo"，运行其匹配 "A=" 的子测试。
go test -run /A=1    # 运行所有匹配 "A=1" 的子测试。
```



#### 并行测试

子测试也可用于控制并行性。所有的子测试完成后，父测试才会完成。在这个例子中，所有的测试是相互并行运行的，当然也只是彼此之间，不包括定义在其他顶层测试的子测试：

```
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

在并行子测试完成之前，Run 方法不会返回，这提供了一种测试后清理的方法：

```
func TestTeardownParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

#### 使用Main 来执行测试

Main

测试程序有时需要在测试之前或之后进行额外的设置（setup）或拆卸（teardown）。有时, 测试还需要控制在主线程上运行的代码。为了支持这些和其他一些情况, 如果测试文件包含函数:

```
func TestMain(m *testing.M)
```

那么生成的测试将调用 TestMain(m)，而不是直接运行测试。TestMain 运行在主 goroutine 中, 可以在调用 m.Run 前后做任何设置和拆卸。应该使用 m.Run 的返回值作为参数调用 os.Exit。在调用 TestMain 时, flag.Parse 并没有被调用。所以，如果 TestMain 依赖于 command-line 标志 (包括 testing 包的标记), 则应该显示的调用 flag.Parse。

一个简单的 TestMain 的实现：

```
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
    // 如果 TestMain 使用了 flags，这里应该加上 flag.Parse()
	os.Exit(m.Run())
}
```

### B T

- [type B](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B)

- - [func (c *B) Error(args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Error)
  - [func (c *B) Errorf(format string, args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Errorf)
  - [func (c *B) Fail()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Fail)
  - [func (c *B) FailNow()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.FailNow)
  - [func (c *B) Failed() bool](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Failed)
  - [func (c *B) Fatal(args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Fatal)
  - [func (c *B) Fatalf(format string, args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Fatalf)
  - [func (c *B) Log(args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Log)
  - [func (c *B) Logf(format string, args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Logf)
  - [func (c *B) Name() string](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Name)
  - [func (b *B) ReportAllocs()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.ReportAllocs)
  - [func (b *B) ResetTimer()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.ResetTimer)
  - [func (b *B) Run(name string, f func(b *B)) bool](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Run)
  - [func (b *B) RunParallel(body func(*PB))](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.RunParallel)
  - [func (b *B) SetBytes(n int64)](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.SetBytes)
  - [func (b *B) SetParallelism(p int)](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.SetParallelism)
  - [func (c *B) Skip(args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Skip)
  - [func (c *B) SkipNow()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.SkipNow)
  - [func (c *B) Skipf(format string, args ...interface{})](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Skipf)
  - [func (c *B) Skipped() bool](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.Skipped)
  - [func (b *B) StartTimer()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.StartTimer)
  - [func (b *B) StopTimer()](https://studygolang.com/static/pkgdoc/pkg/testing.htm#B.StopTimer)

```
func (b *B) RunParallel(body func(*PB))
```

以并行的方式执行给定的基准测试。 RunParallel 会创建出多个 goroutine ，并将 b.N 分配给这些 goroutine 执行， 其中 goroutine 数量的默认值为 GOMAXPROCS 。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在 RunParallel 之前调用 SetParallelism 。RunParallel 通常会与 -cpu 标志一同使用。

body 函数将在每个 goroutine 中执行，这个函数需要设置所有 goroutine 本地的状态， 并迭代直到 pb.Next 返回 false 值为止。因为 StartTimer 、 StopTimer 和 ResetTimer 这三个函数都带有全局作用，所以 body 函数不应该调用这些函数；除此之外，body 函数也不应该调用 Run 函数。





### 测试的要求

文件名以 *_test.go 为例

测试函数用 Test* 开头，测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头

```go
func TestIsPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want  bool
    }{
        {"", true},
        {"a", true},
        {"aa", true},
        {"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
    }
    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalindrome(%q) = %v", test.input, got)
        }
    }
}
```

表格测试

和其他编程语言或测试框架的assert断言不同，t.Errorf调用也没有引起panic异常或停止测试的执行。即使表格中前面的数据导致了测试的失败，表格后面的测试数据依然会运行测试，因此在一个测试中我们可能了解多个失败的信息。我们可以使用t.Fatal或t.Fatalf停止当前测试函数。它们必须在和测试函数同一个goroutine内调用。

```go
func TestIsPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want  bool
    }{
        {"", true},
        {"a", true},
        {"aa", true},
        {"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
    }
    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalindrome(%q) = %v", test.input, got)
        }
    }
}
```



白盒测试和黑盒测试

```Go
package storage

import (
    "fmt"
    "log"
    "net/smtp"
)

func bytesInUse(username string) int64 { return 0 /* ... */ }

// Email sender configuration.
// NOTE: never put passwords in source code!
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`

func CheckQuota(username string) {
    used := bytesInUse(username)
    const quota = 1000000000 // 1GB
    percent := 100 * used / quota
    if percent < 90 {
        return // OK
    }
    msg := fmt.Sprintf(template, used, percent)
    auth := smtp.PlainAuth("", sender, password, hostname)
    err := smtp.SendMail(hostname+":587", auth, sender,
        []string{username}, []byte(msg))
    if err != nil {
        log.Printf("smtp.SendMail(%s) failed: %s", username, err)
    }
}
```

可以通过测试扩展包的方式解决循环依赖的问题，也就是在net/url包所在的目录声明一个独立的url_test测试扩展包。其中测试扩展包名的`_test`后缀告诉go test工具它应该建立一个额外的包来运行测试。我们将这个扩展测试包的导入路径视作是net/url_test会更容易理解，但实际上它并不能被其他任何包导入。

有时候测试扩展包也需要访问被测试包内部的代码，例如在一个为了避免循环导入而被独立到外部测试扩展包的白盒测试。在这种情况下，我们可以通过一些技巧解决：我们在包内的一个_test.go文件中导出一个内部的实现给测试扩展包。因为这些代码只有在测试时才需要，因此一般会放在export_test.go文件中。

例如，fmt包的fmt.Scanf函数需要unicode.IsSpace函数提供的功能。但是为了避免太多的依赖，fmt包并没有导入包含巨大表格数据的unicode包；相反fmt包有一个叫isSpace内部的简易实现。

为了确保fmt.isSpace和unicode.IsSpace函数的行为一致，fmt包谨慎地包含了一个测试。是一个在测试扩展包内的白盒测试，是无法直接访问到isSpace内部函数的，因此fmt通过一个秘密出口导出了isSpace函数。export_test.go文件就是专门用于测试扩展包的秘密出口。

```Go
package fmt

var IsSpace = isSpace
```





较好的测试函数， 不适用assert 使用Errorf  详细的表达信息

```go
func TestSplit(t *testing.T) {
    s, sep := "a:b:c", ":"
    words := strings.Split(s, sep)
    if got, want := len(words), 3; got != want {
        t.Errorf("Split(%q, %q) returned %d words, want %d",
            s, sep, got, want)
    }
    // ...
}
```



### 测试覆盖率

```
$ go tool cover
Usage of 'go tool cover':
Given a coverage profile produced by 'go test':
    go test -coverprofile=c.out

Open a web browser displaying annotated source code:
    go tool cover -html=c.out
...


$ go test -run=Coverage -coverprofile=c.out gopl.io/ch7/eval
ok      gopl.io/ch7/eval         0.032s      coverage: 68.5% of statements
$ go tool cover -html=c.out
```



### 基准测试

基准测试是测量一个程序在固定工作负载下的性能。在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个`*testing.B`类型的参数；`*testing.B`参数除了提供和`*testing.T`类似的方法，还有额外一些和性能测量相关的方法。它还提供了一个整数N，用于指定操作执行的循环次数。

```go
import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}

$ go test -bench=.

goos: linux
goarch: amd64
pkg: test_app/tests
BenchmarkIsPalindrome-8         1000000000               2.52 ns/op
PASS
ok      test_app/tests  2.781s

```



### 测试可视化

```
$ go test -run=NONE -bench=ClientServerParallelTLS64 \
    -cpuprofile=cpu.log net/http
 PASS
 BenchmarkClientServerParallelTLS64-8  1000
    3141325 ns/op  143010 B/op  1747 allocs/op
ok       net/http       3.395s

$ go tool pprof -text -nodecount=10 ./http.test cpu.log
2570ms of 3590ms total (71.59%)
Dropped 129 nodes (cum <= 17.95ms)
Showing top 10 nodes out of 166 (cum >= 60ms)
    flat  flat%   sum%     cum   cum%
  1730ms 48.19% 48.19%  1750ms 48.75%  crypto/elliptic.p256ReduceDegree
   230ms  6.41% 54.60%   250ms  6.96%  crypto/elliptic.p256Diff
   120ms  3.34% 57.94%   120ms  3.34%  math/big.addMulVVW
   110ms  3.06% 61.00%   110ms  3.06%  syscall.Syscall
    90ms  2.51% 63.51%  1130ms 31.48%  crypto/elliptic.p256Square
    70ms  1.95% 65.46%   120ms  3.34%  runtime.scanobject
    60ms  1.67% 67.13%   830ms 23.12%  crypto/elliptic.p256Mul
    60ms  1.67% 68.80%   190ms  5.29%  math/big.nat.montgomery
    50ms  1.39% 70.19%    50ms  1.39%  crypto/elliptic.p256ReduceCarry
    50ms  1.39% 71.59%    60ms  1.67%  crypto/elliptic.p256Sum
```



## 反射

一个例子

```go
func Sprint(x interface{}) string {
   type stringer interface {
      String() string
   }
   switch x := x.(type) {
   case stringer:
      return x.String()
   case string:
      return x
   case int:
      return strconv.Itoa(x)
   case bool:
      if x {
         return "true"
      }
      return "false"
   default:
      return "???"
   }
}
```

### reflect.Type和reflect.Value

所有类型：

Bool, String 和 所有数字类型的基础类型; Array 和 Struct 对应的聚合类型; Chan, Func, Ptr, Slice, 和 Map 对应的引用类似; 接口类型; 还有表示空值的无效类型. (空的 reflect.Value 对应 Invalid 无效类型.)



函数 reflect.TypeOf 接受任意的 interface{} 类型, 并返回对应动态类型的reflect.Type:

```Go
t := reflect.TypeOf(3)  // a reflect.Type
fmt.Println(t.String()) // "int"
fmt.Println(t)          // "int"
```

将一个具体的值转为接口类型会有一个隐式的接口转换操作, 它会创建一个包含两个信息的接口值: 操作数的动态类型(这里是int)和它的动态的值(这里是3).

因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型. 因此, 下面的代码将打印 "*os.File" 而不是 "io.Writer". 稍后, 我们将看到 reflect.Type 是具有识别接口类型的表达方式功能的.

```Go
var w io.Writer = os.Stdout
fmt.Println(reflect.TypeOf(w)) // "*os.File"   ===   fmt.Printf("%T", os.Stdout)
```

要注意的是 reflect.Type 接口是满足 fmt.Stringer 接口的. 因为打印动态类型值对于调试和日志是有帮助的, fmt.Printf 提供了一个简短的 %T 标志参数, 内部使用 reflect.TypeOf 的结果输出:

```Go
fmt.Printf("%T\n", 3) // "int"
```



reflect 包中另一个重要的类型是 Value. 一个 reflect.Value 可以持有一个任意类型的值. 函数 reflect.ValueOf 接受任意的 interface{} 类型, 并返回对应动态类型的reflect.Value. 和 reflect.TypeOf 类似, reflect.ValueOf 返回的结果也是对于具体的类型, 但是 reflect.Value 也可以持有一个接口值.

```Go
v := reflect.ValueOf(3) // a reflect.Value
fmt.Println(v)          // "3"
fmt.Printf("%v\n", v)   // "3"
fmt.Println(v.String()) // NOTE: "<int Value>"
```

调用 Value 的 Type 方法将返回具体类型所对应的 reflect.Type:

```Go
t := v.Type()           // a reflect.Type
fmt.Println(t.String()) // "int"
```

逆操作是调用 reflect.ValueOf 对应的 reflect.Value.Interface 方法. 它返回一个 interface{} 类型表示 reflect.Value 对应类型的具体值:

```Go
v := reflect.ValueOf(3) // a reflect.Value
x := v.Interface()      // an interface{}
i := x.(int)            // an int
fmt.Printf("%d\n", i)   // "3"
```

Value 的 Kind（）

```go
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + "value"
	}
}

```



### 递归打印 类型值

```Go
func display(path string, v reflect.Value) {
    switch v.Kind() {
    case reflect.Invalid:
        fmt.Printf("%s = invalid\n", path)
    case reflect.Slice, reflect.Array:
        for i := 0; i < v.Len(); i++ {
            display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
        }
    case reflect.Struct:
        for i := 0; i < v.NumField(); i++ {
            fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
            display(fieldPath, v.Field(i))
        }
    case reflect.Map:
        for _, key := range v.MapKeys() {
            display(fmt.Sprintf("%s[%s]", path,
                formatAtom(key)), v.MapIndex(key))
        }
    case reflect.Ptr:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            display(fmt.Sprintf("(*%s)", path), v.Elem())
        }
    case reflect.Interface:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
            display(path+".value", v.Elem())
        }
    default: // basic types, channels, funcs
        fmt.Printf("%s = %s\n", path, formatAtom(v))
    }
}
```



### 通过反射 修改值

```Go
x := 2                   // value   type    variable?
a := reflect.ValueOf(2)  // 2       int     no
b := reflect.ValueOf(x)  // 2       int     no
c := reflect.ValueOf(&x) // &x      *int    no
d := c.Elem()            // 2       int     yes (x)
```

其中a对应的变量则不可取地址。因为a中的值仅仅是整数2的拷贝副本。b中的值也同样不可取地址。c中的值还是不可取地址，它只是一个指针`&x`的拷贝。实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。

我们可以通过调用reflect.Value的CanAddr方法来判断其是否可以被取地址：

```Go
fmt.Println(a.CanAddr()) // "false"
fmt.Println(b.CanAddr()) // "false"
fmt.Println(c.CanAddr()) // "false"
fmt.Println(d.CanAddr()) // "true"
```

每当我们通过指针间接地获取的reflect.Value都是可取地址的，即使开始的是一个不可取地址的Value。在反射机制中，所有关于是否支持取地址的规则都是类似的。例如，slice的索引表达式e[i]将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。以此类推，reflect.ValueOf(e).Index(i)对于的值也是可取地址的，即使原始的reflect.ValueOf(e)不支持也没有关系。



```Go
x := 2
d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
px := d.Addr().Interface().(*int) // px := &x
*px = 3                           // x = 3
fmt.Println(x)                    // "3"
```

或者，不使用指针，而是通过调用可取地址的reflect.Value的reflect.Value.Set方法来更新对于的值：

```Go
d.Set(reflect.ValueOf(4))
fmt.Println(x) // "4"
```



### 通过反射 修改 结构体中的tag值

```go
// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
    if err := req.ParseForm(); err != nil {
        return err
    }

    // Build map of fields keyed by effective name.
    fields := make(map[string]reflect.Value)
    v := reflect.ValueOf(ptr).Elem() // the struct variable
    for i := 0; i < v.NumField(); i++ {
        fieldInfo := v.Type().Field(i) // a reflect.StructField
        tag := fieldInfo.Tag           // a reflect.StructTag
        name := tag.Get("http")
        if name == "" {
            name = strings.ToLower(fieldInfo.Name)
        }
        fields[name] = v.Field(i)
    }

    // Update struct field for each parameter in the request.
    for name, values := range req.Form {
        f := fields[name]
        if !f.IsValid() {
            continue // ignore unrecognized HTTP parameters
        }
        for _, value := range values {
            if f.Kind() == reflect.Slice {
                elem := reflect.New(f.Type().Elem()).Elem()
                if err := populate(elem, value); err != nil {
                    return fmt.Errorf("%s: %v", name, err)
                }
                f.Set(reflect.Append(f, elem))
            } else {
                if err := populate(f, value); err != nil {
                    return fmt.Errorf("%s: %v", name, err)
                }
            }
        }
    }
    return nil
}
```

### 访问结构体方法

```Go
// Print prints the method set of the value x.
func Print(x interface{}) {
    v := reflect.ValueOf(x)
    t := v.Type()
    fmt.Printf("type %s\n", t)

    for i := 0; i < v.NumMethod(); i++ {
        methType := v.Method(i).Type()
        fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
            strings.TrimPrefix(methType.String(), "func"))
    }
}
```

