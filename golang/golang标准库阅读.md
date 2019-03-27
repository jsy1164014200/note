# go标准库

## *builtin  内建函数

#### Constants

```
const (
    true  = 0 == 0 // 无类型布尔值
    false = 0 != 0 // 无类型布尔值
)
```

true 和false是两个无类型布尔值。

```
const iota = 0 // 无类型整数值
```

iota是一个预定义的标识符，代表顺序按行增加的无符号整数，每个const声明单元（被括号括起来）相互独立，分别从0开始。

type bool

```
type bool bool
```

布尔类型。

type byte

```
type byte byte
```

8位无符号整型，是uint8的别名，二者视为同一类型。

type rune

```
type rune rune
```

32位有符号整形，int32的别名，二者视为同一类型。

type int

```
type int int
```

至少32位的有符号整形，但和int32/rune并非同一类型。

type int8

```
type int8 int8
```

8位有符号整形，范围[-128, 127]。

type int16

```
type int16 int16
```

16位有符号整形，范围[-32768, 32767]。

type int32

```
type int32 int32
```

32位有符号整形，范围[-2147483648, 2147483647]。

type int64

```
type int64 int64
```

64位有符号整形，范围[-9223372036854775808, 9223372036854775807]。

type uint

```
type uint uint
```

至少32位的无符号整形，但和uint32不是同一类型。

type uint8

```
type uint8 uint8
```

8位无符号整型，范围[0, 255]。

type uint16

```
type uint16 uint16
```

16位无符号整型，范围[0, 65535]。

type uint32

```
type uint32 uint32
```

32位无符号整型，范围[0, 4294967295]。

type uint64

```
type uint64 uint64
```

64位无符号整型，范围[0, 18446744073709551615]。

type float32

```
type float32 float32
```

所有IEEE-754 32位浮点数的集合，12位有效数字。

type float64

```
type float64 float64
```

所有IEEE-754 64位浮点数的集合，16位有效数字。

type complex64

```
type complex64 complex64
```

具有float32 类型实部和虚部的复数类型。

type complex128

```
type complex128 complex128
```

具有float64 类型实部和虚部的复数类型。

type uintptr

```
type uintptr uintptr
```

可以保存任意指针的位模式的整数类型。

type string

```
type string string
```

8位byte序列构成的字符串，约定但不必须是utf-8编码的文本。字符串可以为空但不能是nil，其值不可变。

type error

```
type error interface {
    Error() string
}
```

内建error接口类型是约定用于表示错误信息，nil值表示无错误。

type Type

```
type Type int
```

在本文档中代表任意一个类型，但同一个声明里只代表同一个类型。

```
var nil Type // Type必须是指针、通道、函数、接口、映射或切片
```

nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值。

type Type1

```
type Type1 int
```

在本文档中代表任意一个类型，但同一个声明里只代表同一个类型，用于代表和Type不同的另一类型。

type IntegerType

```
type IntegerType int
```

在本文档中代表一个有符号或无符号的整数类型。

type FloatType

```
type FloatType float32
```

在本文档中代表一个浮点数类型。

type ComplexType

```
type ComplexType complex64
```

在本文档中代表一个复数类型。

#### functions

func [real](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#198)

```
func real(c ComplexType) FloatType
```

返回复数c的实部。

func [imag](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#203)

```
func imag(c ComplexType) FloatType
```

返回复数c的虚部。

func [complex](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#194)

```
func complex(r, i FloatType) ComplexType
```

使用实部r和虚部i生成一个复数。

func [new](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#187)

```
func new(Type) *Type
```

内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针。

func [make](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#182)

```
func make(Type, size IntegerType) Type
```

内建函数make分配并初始化一个类型为切片、映射、或通道的对象。其第一个实参为类型，而非值。make的返回类型与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：

```
切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；
     它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的切片。
映射：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，这种情况下就会分配一个
     小的起始大小。
通道：通道的缓存根据指定的缓存容量初始化。若 size为零或被省略，该信道即为无缓存的。
```

func [cap](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#164)

```
func cap(v Type) int
```

内建函数cap返回 v 的容量，这取决于具体类型：

```
数组：v中元素的数量，与 len(v) 相同
数组指针：*v中元素的数量，与len(v) 相同
切片：切片的容量（底层数组的长度）；若 v为nil，cap(v) 即为零
信道：按照元素的单元，相应信道缓存的容量；若v为nil，cap(v)即为零
```

func [len](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#155)

```
func len(v Type) int
```

内建函数len返回 v 的长度，这取决于具体类型：

```
数组：v中元素的数量
数组指针：*v中元素的数量（v为nil时panic）
切片、映射：v中元素的数量；若v为nil，len(v)即为零
字符串：v中字节的数量
通道：通道缓存中队列（未读取）元素的数量；若v为 nil，len(v)即为零
```

func [append](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#134)

```
func append(slice []Type, elems ...Type) []Type
```

内建函数append将元素追加到切片的末尾。若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组。append返回更新后的切片，因此必须存储追加后的结果。

```
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

作为特例，可以向一个字节切片append字符串，如下：

```
slice = append([]byte("hello "), "world"...)
```

slice 底层原理

```go
func main() {
	arr := [4]int{1, 2, 3, 4}
	sli := arr[0:2:3]
	c := append(sli, 10)

	fmt.Println(arr, sli, c)
}
// [1 2 10 4] [1 2] [1 2 10]
func main() {
	arr := [4]int{1, 2, 3, 4}
	sli := arr[0:2:3]
	c := append(sli, 10, 20)

	fmt.Println(arr, sli, c)

}
// [1 2 3 4] [1 2] [1 2 10 20]
```

func [copy](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#141)

```
func copy(dst, src []Type) int
```

内建函数copy将元素从来源切片复制到目标切片中，也能将字节从字符串复制到字节切片中。copy返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。来源和目标的底层内存可以重叠。

```go
func main() {
	arr := [4]int{1, 2, 3, 4}
	sli := arr[0:2:3]
	c := make([]int, 3)

	i := copy(c, sli)
	fmt.Println(i)
	fmt.Println(arr, sli, c)

}
// 2
// [1 2 3 4] [1 2] [1 2 0]
```



func [delete](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#146)

```
func delete(m map[Type]Type1, key Type)
```

内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。

func [close](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#213)

```
func close(c chan<- Type)
```

内建函数close关闭信道，该通道必须为双向的或只发送的。它应当只由发送者执行，而不应由接收者执行，其效果是在最后发送的值被接收后停止该通道。在最后的值从已关闭的信道中被接收后，任何对其的接收操作都会无阻塞的成功。对于已关闭的信道，语句：

```
x, ok := <-c
```

还会将ok置为false。

func [panic](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#226)

```
func panic(v interface{})
```

内建函数panic停止当前Go程的正常执行。当函数F调用panic时，F的正常执行就会立刻停止。F中defer的所有函数先入后出执行后，F返回给其调用者G。G如同F一样行动，层层返回，直到该Go程中所有函数都按相反的顺序停止执行。之后，程序被终止，而错误情况会被报告，包括引发该恐慌的实参值，此终止序列称为恐慌过程。

func [recover](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#237)

```
func recover() interface{}
```

内建函数recover允许程序管理恐慌过程中的Go程。在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。在此情况下，或当该Go程不在恐慌过程中时，或提供给panic的实参为nil时，recover就会返回nil。

```go
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling...g")
	panic("Error")

	fmt.Println("Returned normally from g.")
}

Calling...g
Recovered in f Error
Returned normally from f.

```





func [print](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#243)

```
func print(args ...Type)
```

内建函数print以特有的方法格式化参数并将结果写入标准错误，用于自举和调试。

func [println](https://github.com/golang/go/blob/master/src/builtin/builtin.go?name=release#250)

```
func println(args ...Type)
```

println类似print，但会在参数输出之间添加空格，输出结束后换行。





## archive  压缩包读写

### archive/tar

> import "archive/tar"

tar包实现了tar格式压缩文件的存取。本包目标是覆盖大多数tar的变种，包括GNU和BSD生成的tar文件。

```go
func main() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive xxxx"},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Printf("内容: %s\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

}
```

Constants

Variables

type Header

- [func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#FileInfoHeader)
- func (h *Header) FileInfo() os.FileInfo

type Reader

- func NewReader(r io.Reader) *Reader
- [func (tr *Reader) Next() (*Header, error)](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#Reader.Next)
- [func (tr *Reader) Read(b [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#Reader.Read)

type Writer

- [func NewWriter(w io.Writer) *Writer](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#NewWriter)
- func (tw *Writer) WriteHeader(hdr *Header) error
- [func (tw *Writer) Write(b [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#Writer.Write)
- [func (tw *Writer) Flush() error](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm#Writer.Flush)
- func (tw *Writer) Close() error

### archive/zip

> import "archive/zip"

zip包提供了zip档案文件的读写服务。

关于ZIP64：

为了向下兼容，FileHeader同时拥有32位和64位的Size字段。64位字段总是包含正确的值，对普通格式的档案未见它们的值是相同的。对zip64格式的档案文件32位字段将是0xffffffff，必须使用64位字段

Constants

Variables

type Compressor

type Decompressor

func RegisterCompressor(method uint16, comp Compressor)

func RegisterDecompressor(method uint16, d Decompressor)

type FileHeader

- [func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileInfoHeader)
- [func (h *FileHeader) FileInfo() os.FileInfo](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileHeader.FileInfo)
- [func (h *FileHeader) Mode() (mode os.FileMode)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileHeader.Mode)
- [func (h *FileHeader) SetMode(mode os.FileMode)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileHeader.SetMode)
- [func (h *FileHeader) ModTime() time.Time](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileHeader.ModTime)
- [func (h *FileHeader) SetModTime(t time.Time)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#FileHeader.SetModTime)

type File

- [func (f *File) DataOffset() (offset int64, err error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#File.DataOffset)
- [func (f *File) Open() (rc io.ReadCloser, err error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#File.Open)

type Reader

- [func NewReader(r io.ReaderAt, size int64) (*Reader, error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#NewReader)

type ReadCloser

- [func OpenReader(name string) (*ReadCloser, error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#OpenReader)
- [func (rc *ReadCloser) Close() error](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#ReadCloser.Close)

type Writer

- [func NewWriter(w io.Writer) *Writer](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#NewWriter)
- [func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#Writer.CreateHeader)
- [func (w *Writer) Create(name string) (io.Writer, error)](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm#Writer.Create)
- func (w *Writer) Close() error

## *bufiio  有缓存的io

> import "bufio"

bufio实现了有缓存的io，它包装了一个io.Reader 或者 io.Writer 接口对象。且同时还提供了缓冲和一些文本I/O的帮助函数的对象。



## *bytes

> import "bytes"

[func Compare(a, b ]byte) int

[func Equal(a, b ]byte) bool

[func EqualFold(s, t ]byte) bool

func Runes(s []byte) []rune

[func HasPrefix(s, prefix ]byte) bool

[func HasSuffix(s, suffix ]byte) bool

[func Contains(b, subslice ]byte) bool

[func Count(s, sep ]byte) int

[func Index(s, sep ]byte) int

[func IndexByte(s ]byte, c byte) int

[func IndexRune(s ]byte, r rune) int

[func IndexAny(s ]byte, chars string) int

[func IndexFunc(s ]byte, f func(r rune) bool) int

[func LastIndex(s, sep ]byte) int

[func LastIndexAny(s ]byte, chars string) int

[func LastIndexFunc(s ]byte, f func(r rune) bool) int

func Title(s []byte) []byte

func ToLower(s []byte) []byte

func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte

func ToUpper(s []byte) []byte

func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte

func ToTitle(s []byte) []byte

func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte

func Repeat(b []byte, count int) []byte

func Replace(s, old, new []byte, n int) []byte

func Map(mapping func(r rune) rune, s []byte) []byte

func Trim(s []byte, cutset string) []byte

func TrimSpace(s []byte) []byte

func TrimFunc(s []byte, f func(r rune) bool) []byte

func TrimLeft(s []byte, cutset string) []byte

func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

func TrimPrefix(s, prefix []byte) []byte

func TrimRight(s []byte, cutset string) []byte

func TrimRightFunc(s []byte, f func(r rune) bool) []byte

func TrimSuffix(s, suffix []byte) []byte

func Fields(s []byte) [][]byte

func FieldsFunc(s []byte, f func(rune) bool) [][]byte

func Split(s, sep []byte) [][]byte

func SplitN(s, sep []byte, n int) [][]byte

func SplitAfter(s, sep []byte) [][]byte

func SplitAfterN(s, sep []byte, n int) [][]byte

func Join(s [][]byte, sep []byte) []byte

type Reader

- [func NewReader(b ]byte) *Reader
- [func (r *Reader) Len() int](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.Len)
- [func (r *Reader) Read(b [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.Read)
- [func (r *Reader) ReadByte() (b byte, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.ReadByte)
- [func (r *Reader) UnreadByte() error](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.UnreadByte)
- [func (r *Reader) ReadRune() (ch rune, size int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.ReadRune)
- [func (r *Reader) UnreadRune() error](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.UnreadRune)
- [func (r *Reader) Seek(offset int64, whence int) (int64, error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.Seek)
- [func (r *Reader) ReadAt(b [\]byte, off int64) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.ReadAt)
- [func (r *Reader) WriteTo(w io.Writer) (n int64, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Reader.WriteTo)

type Buffer

```
var b bytes.Buffer // A Buffer needs no initialization.
b.Write([]byte("Hello "))
fmt.Fprintf(&b, "world!")
b.WriteTo(os.Stdout)
```

Output:

```
Hello world!
```

- [func NewBuffer(buf [\]byte) *Buffer](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#NewBuffer)
- [func NewBufferString(s string) *Buffer](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#NewBufferString)
- [func (b *Buffer) Reset()](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Reset)
- [func (b *Buffer) Len() int](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Len)
- [func (b *Buffer) Bytes() [\]byte](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Bytes)
- [func (b *Buffer) String() string](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.String)
- [func (b *Buffer) Truncate(n int)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Truncate)
- [func (b *Buffer) Grow(n int)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Grow)
- [func (b *Buffer) Read(p [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Read)
- [func (b *Buffer) Next(n int) [\]byte](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Next)
- [func (b *Buffer) ReadByte() (c byte, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.ReadByte)
- [func (b *Buffer) UnreadByte() error](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.UnreadByte)
- [func (b *Buffer) ReadRune() (r rune, size int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.ReadRune)
- [func (b *Buffer) UnreadRune() error](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.UnreadRune)
- [func (b *Buffer) ReadBytes(delim byte) (line [\]byte, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.ReadBytes)
- [func (b *Buffer) ReadString(delim byte) (line string, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.ReadString)
- [func (b *Buffer) Write(p [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.Write)
- [func (b *Buffer) WriteString(s string) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.WriteString)
- [func (b *Buffer) WriteByte(c byte) error](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.WriteByte)
- [func (b *Buffer) WriteRune(r rune) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.WriteRune)
- [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)](https://studygolang.com/static/pkgdoc/pkg/bytes.htm#Buffer.ReadFrom)
- func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)



## compress

> 一堆压缩文件的读写

## *container 常用的数据结构

### container/heap

> import "container/heap"

intheap example

```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```



PriorityQueue example

```go
type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
```

type [Interface](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#30)

```
type Interface interface {
    sort.Interface
    Push(x interface{}) // 向末尾添加元素
    Pop() interface{}   // 从末尾删除元素
}
```

任何实现了本接口的类型都可以用于构建最小堆。最小堆可以通过heap.Init建立，数据是递增顺序或者空的话也是最小堆。最小堆的约束条件是：

```
!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
```

注意接口的Push和Pop方法是供heap包调用的，请使用heap.Push和heap.Pop来向一个堆添加或者删除元素。

func [Init](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#41)

```
func Init(h Interface)
```

一个堆在使用任何堆操作之前应先初始化。Init函数对于堆的约束性是幂等的（多次执行无意义），并可能在任何时候堆的约束性被破坏时被调用。本函数复杂度为O(n)，其中n等于h.Len()。

func [Push](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#52)

```
func Push(h Interface, x interface{})
```

向堆h中插入元素x，并保持堆的约束性。复杂度O(log(n))，其中n等于h.Len()。

func [Pop](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#61)

```
func Pop(h Interface) interface{}
```

删除并返回堆h中的最小元素（不影响约束性）。复杂度O(log(n))，其中n等于h.Len()。等价于Remove(h, 0)。

func [Remove](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#71)

```
func Remove(h Interface, i int) interface{}
```

删除堆中的第i个元素，并保持堆的约束性。复杂度O(log(n))，其中n等于h.Len()。

func [Fix](https://github.com/golang/go/blob/master/src/container/heap/heap.go?name=release#85)

```
func Fix(h Interface, i int)
```

在修改第i个元素后，调用本函数修复堆，比删除第i个元素后插入新元素更有效率。

复杂度O(log(n))，其中n等于h.Len()。

### container/list

> import "container/list"

list包实现了双向链表。要遍历一个链表。

```go
func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
// 1
// 2
// 3
// 4
```

type Element

- [func (e *Element) Next() *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#Element.Next)
- [func (e *Element) Prev() *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#Element.Prev)

type List

- func New() *List
- [func (l *List) Init() *List](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.Init)
- [func (l *List) Len() int](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.Len)
- [func (l *List) Front() *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.Front)
- [func (l *List) Back() *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.Back)
- [func (l *List) PushFront(v interface{}) *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.PushFront)
- [func (l *List) PushFrontList(other *List)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.PushFrontList)
- [func (l *List) PushBack(v interface{}) *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.PushBack)
- [func (l *List) PushBackList(other *List)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.PushBackList)
- [func (l *List) InsertBefore(v interface{}, mark *Element) *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.InsertBefore)
- [func (l *List) InsertAfter(v interface{}, mark *Element) *Element](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.InsertAfter)
- [func (l *List) MoveToFront(e *Element)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.MoveToFront)
- [func (l *List) MoveToBack(e *Element)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.MoveToBack)
- [func (l *List) MoveBefore(e, mark *Element)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.MoveBefore)
- [func (l *List) MoveAfter(e, mark *Element)](https://studygolang.com/static/pkgdoc/pkg/container_list.htm#List.MoveAfter)
- func (l *List) Remove(e *Element) interface{}

### container/ring

> import "container/ring"

环形链表

type Ring

- [func New(n int) *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#New)
- [func (r *Ring) Len() int](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Len)
- [func (r *Ring) Next() *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Next)
- [func (r *Ring) Prev() *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Prev)
- [func (r *Ring) Move(n int) *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Move)
- [func (r *Ring) Link(s *Ring) *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Link)
- [func (r *Ring) Unlink(n int) *Ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm#Ring.Unlink)
- func (r *Ring) Do(f func(interface{}))

## *context   记得加内容



## crypto  加密解密包

> import "crypto"

crypto包收集了常用的密码算法 常量

type PublicKey

type PrivateKey

type Hash

- [func (h Hash) Available() bool](https://studygolang.com/static/pkgdoc/pkg/crypto.htm#Hash.Available)
- [func (h Hash) Size() int](https://studygolang.com/static/pkgdoc/pkg/crypto.htm#Hash.Size)
- [func (h Hash) New() hash.Hash](https://studygolang.com/static/pkgdoc/pkg/crypto.htm#Hash.New)

func RegisterHash(h Hash, f func() hash.Hash)

### ##crypto/cipher  加密器

```
import "crypto/cipher"
```

cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现。

type Block  Block接口代表一个使用特定密钥的底层块加/解密器。它提供了加密和解密独立数据块的能力。

type BlockMode  BlockMode接口代表一个工作在块模式（如CBC、ECB等）的加/解密器。

- [func NewCBCDecrypter(b Block, iv ]byte) BlockMode

  ```go
  func main() {
  	key := []byte("example key 1234")
  	plaintext := []byte("exampleplaintext")
  	if len(plaintext)%aes.BlockSize != 0 {
  		panic("plaintext is not a multiple of the block size")
  	}
  	// 创建一个 aes 的秘钥（暗号）
  	block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}
  
  	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
  	iv := ciphertext[:aes.BlockSize]
  
  	// 加入一个随机 串
  	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
  		panic(err)
  	}
  	// 一个加密器
  	mode := cipher.NewCBCEncrypter(block, iv)
  	// iv 的用途 ：返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
  	//NewCBCDecrypter 的时候的 iv要跟它一样
  	fmt.Printf("%x\n", iv)
  	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
  	fmt.Printf("%x\n", ciphertext)
  }
  ```

- [func NewCBCEncrypter(b Block, iv ]byte) BlockMode

  ```go
  
  func main() {
  	key := []byte("example key 1234")
  	ciphertext, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")
  	block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}
  
  	if len(ciphertext) < aes.BlockSize {
  		panic("ciphertext too short")
  	}
  	iv := ciphertext[:aes.BlockSize]
  	ciphertext = ciphertext[aes.BlockSize:]
  	if len(ciphertext)%aes.BlockSize != 0 {
  		panic("ciphertext is not a multiple of the block size")
  	}
  	mode := cipher.NewCBCDecrypter(block, iv)
  	mode.CryptBlocks(ciphertext, ciphertext)
  	fmt.Printf("%s\n", ciphertext)
  }
  
  ```


type Stream

```go
key := []byte("example key 1234")
plaintext := []byte("some plaintext")
block, err := aes.NewCipher(key)
if err != nil {
    panic(err)
}
// The IV needs to be unique, but not secure. Therefore it's common to
// include it at the beginning of the ciphertext.
ciphertext := make([]byte, aes.BlockSize+len(plaintext))
iv := ciphertext[:aes.BlockSize]
if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    panic(err)
}
stream := cipher.NewCFBEncrypter(block, iv)
stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
// It's important to remember that ciphertexts must be authenticated
// (i.e. by using crypto/hmac) as well as being encrypted in order to
// be secure.
```



- [func NewCFBDecrypter(block Block, iv ]byte) Stream
- [func NewCFBEncrypter(block Block, iv ]byte) Stream
- [func NewCTR(block Block, iv ]byte) Stream
- [func NewOFB(b Block, iv ]byte) Stream

type StreamReader

- [func (r StreamReader) Read(dst [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/crypto_cipher.htm#StreamReader.Read)

type StreamWriter

- [func (w StreamWriter) Write(src [\]byte) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/crypto_cipher.htm#StreamWriter.Write)
- [func (w StreamWriter) Close() error](https://studygolang.com/static/pkgdoc/pkg/crypto_cipher.htm#StreamWriter.Close)

type AEAD

- [func NewGCM(cipher Block) (AEAD, error)](https://studygolang.com/static/pkgdoc/pkg/crypto_cipher.htm#NewGCM)





### crypto/aes

AES加密算法

Constants

type KeySizeError

- [func (k KeySizeError) Error() string](https://studygolang.com/static/pkgdoc/pkg/crypto_aes.htm#KeySizeError.Error)

[func NewCipher(key ]byte) (cipher.Block, error)



### crypto/des 

TDEA算法  和   DES算法

Constants

type KeySizeError

- [func (k KeySizeError) Error() string](https://studygolang.com/static/pkgdoc/pkg/crypto_des.htm#KeySizeError.Error)

[func NewCipher(key ]byte) (cipher.Block, error)

DES算法

[func NewTripleDESCipher(key ]byte) (cipher.Block, error)

TDEA算法

### crypto/rc4

### crypto/rsa

ras （非对称）

### crypto/dsa

> import "crypto/dsa"

dsa包实现了DSA算法（非对称）

[
Variables](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#pkg-variables)

[type ParameterSizes](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#ParameterSizes)

[type Parameters](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#Parameters)

[type PublicKey](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#PublicKey)

[type PrivateKey](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#PrivateKey)

[func GenerateParameters(params *Parameters, rand io.Reader, sizes ParameterSizes) (err error)](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#GenerateParameters)

[func GenerateKey(priv *PrivateKey, rand io.Reader) error](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#GenerateKey)

[func Sign(rand io.Reader, priv *PrivateKey, hash [\]byte) (r, s *big.Int, err error)](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#Sign)

[func Verify(pub *PublicKey, hash [\]byte, r, s *big.Int) bool](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm#Verify)



func [GenerateKey](https://github.com/golang/go/blob/master/src/crypto/dsa/dsa.go?name=release#151)

func [GenerateParameters](https://github.com/golang/go/blob/master/src/crypto/dsa/dsa.go?name=release#55)

```
func GenerateParameters(params *Parameters, rand io.Reader, sizes ParameterSizes) (err error)
```

GenerateParameters函数随机设置合法的参数到params。即使机器很快，函数也可能会花费很多时间来生成参数。

```
func GenerateKey(priv *PrivateKey, rand io.Reader) error
```

GenerateKey生成一对公钥和私钥；priv.PublicKey.Parameters字段必须已经（被GenerateParameters函数）设置了合法的参数。

func [Sign](https://github.com/golang/go/blob/master/src/crypto/dsa/dsa.go?name=release#194)

```
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
```

使用私钥对任意长度的hash值（必须是较大信息的hash结果）进行签名，返回签名结果（一对大整数）。私钥的安全性取决于密码读取器的熵度（随机程度）。

注意根据FIPS 186-3 section 4.6的规定，hash必须被截断到亚组的长度，本函数是不会自己截断的。

func [Verify](https://github.com/golang/go/blob/master/src/crypto/dsa/dsa.go?name=release#249)

```
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
```

使用公钥认证hash和两个大整数r、s构成的签名，报告签名是否合法。

注意根据FIPS 186-3 section 4.6的规定，hash必须被截断到亚组的长度，本函数是不会自己截断的。



### crypto/ecdsa

ecdsa包实现了椭圆曲线数字签名算法（非对称）

### crypto/elliptic

elliptic包实现了几条覆盖素数有限域的标准椭圆曲线。

### *##crypto/hmac hash加密

hmac包实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）。

HMAC是使用key标记信息的加密hash。接收者使用相同的key逆运算来认证hash。

出于安全目的，接收者应使用Equal函数比较认证码：

```go
// 如果messageMAC是message的合法HMAC标签，函数返回真
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
```

[func Equal(mac1, mac2 [\]byte) bool](https://studygolang.com/static/pkgdoc/pkg/crypto_hmac.htm#Equal)

[func New(h func() hash.Hash, key [\]byte) hash.Hash](https://studygolang.com/static/pkgdoc/pkg/crypto_hmac.htm#New)



### crypto/md5

md5哈希算法

```go
func main() {
	a := md5.Sum([]byte("There pretzels are making me thirsty."))
	b := md5.Sum([]byte("There pretzels are making me thirsty."))
	fmt.Printf("%x", a)
	fmt.Println()
	fmt.Printf("%x", b)
}

```



func [Sum](https://github.com/golang/go/blob/master/src/crypto/md5/md5.go?name=release#129)

```
func Sum(data []byte) [Size]byte
```

返回数据data的MD5校验和。

Example

func [New](https://github.com/golang/go/blob/master/src/crypto/md5/md5.go?name=release#49)

```
func New() hash.Hash
```

返回一个新的使用MD5校验的hash.Hash接口。



### crypto/sha1

### crypto/sha256

sha256包实现了SHA224和SHA256哈希算法

### crypto/sha512

sha512包实现了SHA384和SHA512哈希算法

### crypto/rand





## crypto/tls

安全套接字

## crypto/ x509

证书相关

## database

sql 数据库的驱动

## *encoding  数据格式解析

encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口。encoding/gob、encoding/json、encoding/xml三个包都会检查使用这些接口。因此，只要实现了这些接口一次，就可以在多个包里使用。标准包内建类型time.Time和net.IP都实现了这些接口。接口是成对的，分别产生和还原编码后的数据。

[type BinaryMarshaler](https://studygolang.com/static/pkgdoc/pkg/encoding.htm#BinaryMarshaler)

[type BinaryUnmarshaler](https://studygolang.com/static/pkgdoc/pkg/encoding.htm#BinaryUnmarshaler)

[type TextMarshaler](https://studygolang.com/static/pkgdoc/pkg/encoding.htm#TextMarshaler)

[type TextUnmarshaler](https://studygolang.com/static/pkgdoc/pkg/encoding.htm#TextUnmarshaler)



### encoding/base64

base16(又称hex)  base32



```go
func main() {
	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}

data := []byte("any + old & data")
str := base64.StdEncoding.EncodeToString(data)
fmt.Println(str)
```



Variables

type CorruptInputError

- func (e CorruptInputError) Error() string

func [NewEncoding](https://github.com/golang/go/blob/master/src/encoding/base64/base64.go?name=release#34)

```
func NewEncoding(encoder string) *Encoding
```

使用给出的字符集生成一个*Encoding，字符集必须是64字节的字符串。

type Encoding

- func NewEncoding(encoder string) *Encoding
- func (enc *Encoding) DecodedLen(n int) int
- [func (enc *Encoding) Decode(dst, src ]byte) (n int, err error)
- [func (enc *Encoding) DecodeString(s string) (]byte, error)
- func (enc *Encoding) EncodedLen(n int) int
- [func (enc *Encoding) Encode(dst, src [\]byte)](https://studygolang.com/static/pkgdoc/pkg/encoding_base64.htm#Encoding.Encode)
- [func (enc *Encoding) EncodeToString(src ]byte) string

func NewDecoder(enc *Encoding, r io.Reader) io.Reader

func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

最后两个是流解码器

```go
input := []byte("foo\x00bar")
encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
encoder.Write(input)
// Must close the encoder when finished to flush any partial blocks.
// If you comment out the following line, the last partial block "r"
// won't be encoded.
encoder.Close()
```



### encoding/csv

```
import "encoding/csv"
```

csv读写逗号分隔值（csv）的文件。

一个csv分拣包含零到多条记录，每条记录一到多个字段。每个记录用换行符分隔。最后一条记录后面可以有换行符，也可以没有。

```
field1,field2,field3
```

空白视为字段的一部分。

换行符前面的回车符会悄悄的被删掉。

忽略空行。只有空白的行（除了末尾换行符之外）不视为空行。

以双引号"开始和结束的字段成为受引字段，其开始和结束的引号不属于字段。

资源：

```
normal string,"quoted-field"
```

产生如下字段：

```
{`normal string`, `quoted-field`}
```

受引字段内部，如果有两个连续的双引号，则视为一个单纯的双引号字符：

```
"the ""word"" is true","a ""quoted-field"""
```

生成：

```
{`the "word" is true`, `a "quoted-field"`}
```

受引字段里可以包含换行和逗号：

```
"Multi-line
field","comma is ,"
```

生成：

```
{`Multi-line
field`, `comma is ,`}
```



Variables

type ParseError

- func (e *ParseError) Error() string

type Reader

```
type Reader struct {
    Comma            rune // 字段分隔符（NewReader将之设为','）
    Comment          rune // 一行开始位置的注释标识符
    FieldsPerRecord  int  // 每条记录期望的字段数
    LazyQuotes       bool // 允许懒引号
    TrailingComma    bool // 忽略，出于后端兼容性而保留
    TrimLeadingSpace bool // 去除前导的空白
    // 内含隐藏或非导出字段
}
```

- [func NewReader(r io.Reader) *Reader](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#NewReader)
- [func (r *Reader) Read() (record [\]string, err error)](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#Reader.Read)
- [func (r *Reader) ReadAll() (records [\][]string, err error)](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#Reader.ReadAll)

type Writer

- [func NewWriter(w io.Writer) *Writer](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#NewWriter)
- [func (w *Writer) Write(record [\]string) (err error)](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#Writer.Write)
- [func (w *Writer) WriteAll(records [\][]string) (err error)](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm#Writer.WriteAll)
- func (w *Writer) Flush()
- func (w *Writer) Error() error



### encoding/gob

gob包管理gob流——在编码器（发送器）和解码器（接受器）之间交换的binary值。一般用于传递远端程序调用（RPC）的参数和结果，如net/rpc包就有提供。

本实现给每一个数据类型都编译生成一个编解码程序，当单个编码器用于传递数据流时，会分期偿还编译的消耗，是效率最高的。



### encoding/hex

base16 进制的编解码

Variables

type InvalidByteError

- [func (e InvalidByteError) Error() string](https://studygolang.com/static/pkgdoc/pkg/encoding_hex.htm#InvalidByteError.Error)

func DecodedLen(x int) int

[func Decode(dst, src ]byte) (int, error)

[func DecodeString(s string) (]byte, error)

func EncodedLen(n int) int

[func Encode(dst, src ]byte) int

[func EncodeToString(src ]byte) string

[func Dump(data ]byte) string

func Dumper(w io.Writer) io.WriteCloser



### encoding/json

> 一般用 simplejson方便些

### encoding/xml

> 见包工具笔记



## errors  新建error类型

func [New](https://github.com/golang/go/blob/master/src/errors/errors.go?name=release#9)

```
func New(text string) error
```

```go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}



const name, id = "bimmler", 17
err := fmt.Errorf("user %q (id %d) not found", name, id)
if err != nil {
    fmt.Print(err)
}
```

## flag 命令行参数解析

```go
// 单个参数解析
var ip = flag.Int("flagname", 1234, "help message for flagname")
// 或者将 flag绑定到一个变量
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}

```

你可以自定义一个用于flag的类型（满足Value接口）并将该类型用于flag解析，如下：

```
flag.Var(&flagVal, "name", "help message for flagname")
```

对这种flag，默认值就是该变量的初始值。

在所有flag都注册之后，调用：

```
flag.Parse()
```

来解析命令行参数写入注册的flag里。

解析之后，flag的值可以直接使用。如果你使用的是flag自身，它们是指针；如果你绑定到了某个变量，它们是值。

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

解析后，flag后面的参数可以从flag.Args()里获取或用flag.Arg(i)单独获取。这些参数的索引为从0到flag.NArg()-1。



## fmt 格式输出

#### Printing

verb：

通用：

```
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	值的类型的Go语法表示
%%	百分号
```

布尔值：

```
%t	单词true或false
```

整数：

```
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X"
```

浮点数与复数的两个组分：

```
%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
```

字符串和[]byte：

```
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f）
%X	每个字节用两字符十六进制数表示（使用A-F）    
```

指针：

```
%p	表示为十六进制，并加上前导的0x    
```

没有%u。整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）。

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

```
%f:    默认宽度，默认精度
%9f    宽度9，默认精度
%.2f   默认宽度，精度2
%9.2f  宽度9，精度2
%9.f   宽度9，精度0 
```

```
'+'	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
' '	对数值，正数前加空格而负数前加负号；
'-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
'#'	切换格式：
  	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
 	对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
 	对%U（%#U），输出Unicode格式后，如字符可打印，还会输出空格和单引号括起来的go字面值；
  	对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
'0'	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
```





```
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```

会生成"22 11"

[type Stringer](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Stringer)

```
type Stringer interface {
    String() string
}
```

[type GoStringer](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#GoStringer)

```
type GoStringer interface {
    GoString() string
}
```

实现了GoStringer接口的类型（即有GoString方法），定义了该类型值的go语法表示。当采用verb %#v格式化一个操作数时，会调用GoString方法来生成输出的文本。

[type State](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#State)

[type Formatter](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Formatter)

[type ScanState](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#ScanState)

[type Scanner](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Scanner)

[func Printf(format string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Printf)

[func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fprintf)

[func Sprintf(format string, a ...interface{}) string](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sprintf)

[func Print(a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Print)

[func Fprint(w io.Writer, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fprint)

[func Sprint(a ...interface{}) string](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sprint)

[func Println(a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Println)

[func Fprintln(w io.Writer, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fprintln)

[func Sprintln(a ...interface{}) string](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sprintln)

[func Errorf(format string, a ...interface{}) error](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Errorf)

[func Scanf(format string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Scanf)

[func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fscanf)

[func Sscanf(str string, format string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sscanf)

[func Scan(a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Scan)

[func Fscan(r io.Reader, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fscan)

[func Sscan(str string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sscan)

[func Scanln(a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Scanln)

[func Fscanln(r io.Reader, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Fscanln)

[func Sscanln(str string, a ...interface{}) (n int, err error)](https://studygolang.com/static/pkgdoc/pkg/fmt.htm#Sscanln)





## html 文本的解析

func [EscapeString](https://github.com/golang/go/blob/master/src/html/escape.go?name=release#229)

```
func EscapeString(s string) string
```

EscapeString函数将特定的一些字符转为逸码后的字符实体，如"<"变成"&lt;"。

它只会修改五个字符：<、>、&、'、"。

UnescapeString(EscapeString(s)) == s总是成立，但是两个函数顺序反过来则不一定成立。

func [UnescapeString](https://github.com/golang/go/blob/master/src/html/escape.go?name=release#243)

```
func UnescapeString(s string) string
```

UnescapeString函数将逸码的字符实体如"&lt;"修改为原字符"<"。它会解码一个很大范围内的字符实体，远比函数EscapeString转码范围大得多。例如"&aacute;"解码为"á"，"&#225;"和"&xE1;"也会解码为该字符。

### html/template

本包是对[text/template](http://godoc.org/text/template)包的包装，两个包提供的模板API几无差别，可以安全的随意替换两包。

本包使用的安全模型假设模板的作者是可信任的，但用于执行的数据不可信。

示例：

```
import "text/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

生成：

```
Hello, <script>alert('you have been pwned')</script>!
```

但在html/template包里会根据上下文自动转义：

```
import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

生成安全的转义后HTML输出：

```
Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
```

所以一般使用 html/template 里面的api





## image 图片

image实现了基本的2D图片库。

基本接口叫作Image。图片的色彩定义在image/color包。

Image接口可以通过调用如NewRGBA和NewPaletted函数等获得；也可以通过调用Decode函数解码包含GIF、JPEG或PNG格式图像数据的输入流获得。解码任何具体图像类型之前都必须注册对应类型的解码函数。注册过程一般是作为包初始化的副作用，放在包的init函数里。因此，要解码PNG图像，只需在程序的main包里嵌入如下代码：

```
import _ "image/png"
```



> 文档里面有 一个demo





# *io 库

因为这些接口和原语是对底层实现完全不同的低水平操作的包装，除非得到其它方面的通知，客户端不应假设它们是并发执行安全的。

## Variables

```
var EOF = errors.New("EOF")
```

EOF当无法得到更多输入时，Read方法返回EOF。当函数一切正常的到达输入的结束时，就应返回EOF。如果在一个结构化数据流中EOF在不期望的位置出现了，则应返回错误ErrUnexpectedEOF或者其它给出更多细节的错误。

```
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
```

当从一个已关闭的Pipe读取或者写入时，会返回ErrClosedPipe。

```
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
```

某些使用io.Reader接口的客户端如果多次调用Read都不返回数据也不返回错误时，就会返回本错误，一般来说是io.Reader的实现有问题的标志。

```
var ErrShortBuffer = errors.New("short buffer")
```

ErrShortBuffer表示读取操作需要大缓冲，但提供的缓冲不够大。

```
var ErrShortWrite = errors.New("short write")
```

ErrShortWrite表示写入操作写入的数据比提供的少，却没有显式的返回错误。

```
var ErrUnexpectedEOF = errors.New("unexpected EOF")
```

ErrUnexpectedEOF表示在读取一个固定尺寸的块或者数据结构时，在读取未完全时遇到了EOF。



## Type Reader

interface {Read(p []byte) (n int, err error)}

Reader接口用于包装基本的读取方法。

Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。即使Read方法返回值n < len(p)，本方法在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，Read按惯例会返回可以读取到的数据，而不是等待更多数据。

当Read在读取n > 0个字节后遭遇错误或者到达文件结尾时，会返回读取的字节数。它可能会在该次调用返回一个非nil的错误，或者在下一次调用时返回0和该错误。一个常见的例子，Reader接口会在输入流的结尾返回非0的字节数，返回值err == EOF或err == nil。但不管怎样，下一次Read调用必然返回(0, EOF)。

**调用者应该总是先处理读取的n > 0字节再处理错误值。这么做可以正确的处理发生在读取部分数据后的I/O错误，也能正确处理EOF事件。**

**如果Read的某个实现返回0字节数和nil错误值，表示被阻碍；调用者应该将这种情况视为未进行操作。**



## Type Writer

interface {Write(p []byte) (n int, err error) }

Writer接口用于包装基本的写入方法。

Write方法len(p) 字节数据从p写入底层的数据流。它会返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提取结束的错误。Write必须返回非nil的错误，如果它返回的 n < len(p)。Write不能修改切片p中的数据，即使临时修改也不行。

## Type Closer

## Type Seeker

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64 ,error)
}
```

Seeker接口用于包装基本的移位方法。

Seek方法设定下一次读写的位置：偏移量为offset，校准点由whence确定：0表示相对于文件起始；1表示相对于当前位置；2表示相对于文件结尾。Seek方法返回新的位置以及可能遇到的错误。

移动到一个绝对偏移量为负数的位置会导致错误。移动到任何偏移量为正数的位置都是合法的，但其下一次I/O操作的具体行为则要看底层的实现。

## type ReaderAt

```
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

ReaderAt接口包装了基本的ReadAt方法。

ReadAt从底层输入流的偏移量off位置读取len(p)字节数据写入p， 它返回读取的字节数(0 <= n <= len(p))和遇到的任何错误。当ReadAt方法返回值n < len(p)时，它会返回一个非nil的错误来说明为啥没有读取更多的字节。在这方面，ReadAt是比Read要严格的。即使ReadAt方法返回值 n < len(p)，它在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，ReadAt会阻塞直到获取len(p)个字节数据或者遇到错误。在这方面，ReadAt和Read是不同的。如果ReadAt返回时到达输入流的结尾，而返回值n == len(p)，其返回值err既可以是EOF也可以是nil。

如果ReadAt是从某个有偏移量的底层输入流（的Reader包装）读取，ReadAt方法既不应影响底层的偏移量，也不应被底层的偏移量影响。

ReadAt方法的调用者可以对同一输入流执行并行的ReadAt调用。

## type [WriterAt](https://github.com/golang/go/blob/master/src/io/io.go?name=release#212)

```
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

WriterAt接口包装了基本的WriteAt方法。

WriteAt将p全部len(p)字节数据写入底层数据流的偏移量off位置。它返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提前中止的错误。当其返回值n < len(p)时，WriteAt必须放哪会一个非nil的错误。

如果WriteAt写入的对象是某个有偏移量的底层输出流（的Writer包装），WriteAt方法既不应影响底层的偏移量，也不应被底层的偏移量影响。

ReadAt方法的调用者可以对同一输入流执行并行的WriteAt调用。（前提是写入范围不重叠）

## type [ReaderFrom](https://github.com/golang/go/blob/master/src/io/io.go?name=release#156)

```
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

ReaderFrom接口包装了基本的ReadFrom方法。

ReadFrom方法从r读取数据直到EOF或者遇到错误。返回值n是读取的字节数，执行时遇到的错误（EOF除外）也会被返回。

## type [WriterTo](https://github.com/golang/go/blob/master/src/io/io.go?name=release#167)

```
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

WriterTo接口包装了基本的WriteTo方法。

WriteTo方法将数据写入w直到没有数据可以写入或者遇到错误。返回值n是写入的字节数，执行时遇到的任何错误也会被返回。



## func [Pipe](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#186)

```
func Pipe() (*PipeReader, *PipeWriter)
```

Pipe创建一个同步的内存中的管道。它可以用于连接期望io.Reader的代码和期望io.Writer的代码。一端的读取对应另一端的写入，直接在两端拷贝数据，没有内部缓冲。可以安全的并行调用Read和Write或者Read/Write与Close方法。Close方法会在最后一次阻塞中的I/O操作结束后完成。并行调用Read或并行调用Write也是安全的：每一个独立的调用会依次进行。

## type [PipeReader](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#124)

```
type PipeReader struct {
    // 内含隐藏或非导出字段
}
```

PipeReader是一个管道的读取端。

#### func (*PipeReader) [Read](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#133)

```
func (r *PipeReader) Read(data []byte) (n int, err error)
```

Read实现了标准Reader接口：它从管道中读取数据，会阻塞直到写入端开始写入或写入端被关闭。

#### func (*PipeReader) [Close](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#139)

```
func (r *PipeReader) Close() error
```

Close关闭读取器；关闭后如果对管道的写入端进行写入操作，就会返回(0, ErrClosedPip)。

#### func (*PipeReader) [CloseWithError](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#145)

```
func (r *PipeReader) CloseWithError(err error) error
```

CloseWithError类似Close方法，但将调用Write时返回的错误改为err。

## type [PipeWriter](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#151)

```
type PipeWriter struct {
    // 内含隐藏或非导出字段
}
```

#### func (*PipeWriter) [Write](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#160)

```
func (w *PipeWriter) Write(data []byte) (n int, err error)
```

Write实现了标准Writer接口：它将数据写入到管道中，会阻塞直到读取器读完所有的数据或读取端被关闭。

#### func (*PipeWriter) [Close](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#166)

```
func (w *PipeWriter) Close() error
```

Close关闭写入器；关闭后如果对管道的读取端进行读取操作，就会返回(0, EOF)。

#### func (*PipeWriter) [CloseWithError](https://github.com/golang/go/blob/master/src/io/pipe.go?name=release#172)

```
func (w *PipeWriter) CloseWithError(err error) error
```

CloseWithError类似Close方法，但将调用Read时返回的错误改为err

## func [TeeReader](https://github.com/golang/go/blob/master/src/io/io.go?name=release#476)

```
func TeeReader(r Reader, w Writer) Reader
```

TeeReader返回一个将其从r读取的数据写入w的Reader接口。所有通过该接口对r的读取都会执行对应的对w的写入。没有内部的缓冲：写入必须在读取完成前完成。写入时遇到的任何错误都会作为读取错误返回。

## func [MultiReader](https://github.com/golang/go/blob/master/src/io/multi.go?name=release#31)

```
func MultiReader(readers ...Reader) Reader
```

MultiReader返回一个将提供的Reader在逻辑上串联起来的Reader接口。他们依次被读取。当所有的输入流都读取完毕，Read才会返回EOF。如果readers中任一个返回了非nil非EOF的错误，Read方法会返回该错误。

## func [MultiWriter](https://github.com/golang/go/blob/master/src/io/multi.go?name=release#57)

```
func MultiWriter(writers ...Writer) Writer
```

MultiWriter创建一个Writer接口，会将提供给其的数据写入所有创建时提供的Writer接口。

## func [Copy](https://github.com/golang/go/blob/master/src/io/io.go?name=release#341)

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。

## func [CopyN](https://github.com/golang/go/blob/master/src/io/io.go?name=release#317)

```
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

从src拷贝n个字节数据到dst，直到在src上到达EOF或发生错误。返回复制的字节数和遇到的第一个错误。

只有err为nil时，written才会等于n。如果dst实现了ReaderFrom接口，本函数很调用它实现拷贝。

## f unc [ReadAtLeast](https://github.com/golang/go/blob/master/src/io/io.go?name=release#283)

```
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

ReadAtLeast从r至少读取min字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 如果min比buf的长度还大，函数会返回ErrShortBuffer。只有返回值err为nil时，返回值n才会不小于min。

## func [ReadFull](https://github.com/golang/go/blob/master/src/io/io.go?name=release#306)

```
func ReadFull(r Reader, buf []byte) (n int, err error)
```

ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 只有返回值err为nil时，返回值n才会等于len(buf)。

## func [WriteString](https://github.com/golang/go/blob/master/src/io/io.go?name=release#269)

```
func WriteString(w Writer, s string) (n int, err error)
```

WriteString函数将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法。



# io/ioutil

### func [ReadAll](https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go?name=release#41)

```
func ReadAll(r io.Reader) ([]byte, error)
```

ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回的err为nil而非EOF。因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。

### func [ReadFile](https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go?name=release#49)

```
func ReadFile(filename string) ([]byte, error)
```

ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。

### func [WriteFile](https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go?name=release#76)

```
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。

### func [ReadDir](https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go?name=release#100)

```
func ReadDir(dirname string) ([]os.FileInfo, error)
```

返回dirname指定的目录的目录信息的有序列表。

### func [TempDir](https://github.com/golang/go/blob/master/src/io/ioutil/tempfile.go?name=release#74)

```
func TempDir(dir, prefix string) (name string, err error)
```

在dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径。如果dir是空字符串，TempDir使用默认用于临时文件的目录（参见os.TempDir函数）。 不同程序同时调用该函数会创建不同的临时目录，调用本函数的程序有责任在不需要临时文件夹时摧毁它。

### func [TempFile](https://github.com/golang/go/blob/master/src/io/ioutil/tempfile.go?name=release#47)

```
func TempFile(dir, prefix string) (f *os.File, err error)
```

在dir目录下创建一个新的、使用prefix为前缀的临时文件，以读写模式打开该文件并返回os.File指针。如果dir是空字符串，TempFile使用默认用于临时文件的目录（参见os.TempDir函数）。不同程序同时调用该函数会创建不同的临时文件，调用本函数的程序有责任在不需要临时文件时摧毁它。



# log 日志库

## Type Logger

Logger类型表示一个活动状态的记录日志的对象，它会生成一行行的输出写入一个io.Writer接口。每一条日志操作会调用一次io.Writer接口的Write方法。Logger类型的对象可以被多个线程安全的同时使用，它会保证对io.Writer接口的顺序访问。

```go
var buf bytes.Buffer
logger := log.New(&buf,"Logger: ", log.Lshortfile)
logger.Print("Hello, log file!")
fmt.Print(&buf)
```

#### func [New](https://github.com/golang/go/blob/master/src/log/log.go?name=release#54)

```
func New(out io.Writer, prefix string, flag int) *Logger
```

New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

type Logger

- [func New(out io.Writer, prefix string, flag int) *Logger](https://studygolang.com/static/pkgdoc/pkg/log.htm#New)

- [func (l *Logger) Flags() int](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Flags)

- [func (l *Logger) SetFlags(flag int)](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.SetFlags)

- [func (l *Logger) Prefix() string](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Prefix)

- func (l *Logger) SetPrefix(prefix string)

- [func (l *Logger) Output(calldepth int, s string) error](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Output)

- [func (l *Logger) Printf(format string, v ...interface{})](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Printf)

- [func (l *Logger) Print(v ...interface{})](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Print)

- [func (l *Logger) Println(v ...interface{})](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Println)

- func (l *Logger) Fatalf(format string, v ...interface{})

  相当于 加了一个 os.Exit(1)

- func (l *Logger) Fatal(v ...interface{})

- [func (l *Logger) Fatalln(v ...interface{})](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Fatalln)

- func (l *Logger) Panic(v ...interface{})

  相当于加了个 panic()

- [func (l *Logger) Panicf(format string, v ...interface{})](https://studygolang.com/static/pkgdoc/pkg/log.htm#Logger.Panicf)

- func (l *Logger) Panicln(v ...interface{})



下面的api用于对付标准输出 logger

func Flags() int

func SetFlags(flag int)

func Prefix() string

func SetPrefix(prefix string)

func SetOutput(w io.Writer)

func Printf(format string, v ...interface{})

func Print(v ...interface{})

func Println(v ...interface{})

func Fatalf(format string, v ...interface{})

func Fatal(v ...interface{})

func Fatalln(v ...interface{})

func Panicf(format string, v ...interface{})

func Panic(v ...interface{})

func Panicln(v ...interface{})



## log/syslog

用于 对接操作系统的日志系统



# math

Constants

func NaN() float64

func IsNaN(f float64) (is bool)

func Inf(sign int) float64

func IsInf(f float64, sign int) bool

func Float32bits(f float32) uint32

func Float32frombits(b uint32) float32

func Float64bits(f float64) uint64

func Float64frombits(b uint64) float64

func Signbit(x float64) bool

func Copysign(x, y float64) float64

func Ceil(x float64) float64

func Floor(x float64) float64

func Trunc(x float64) float64

func Modf(f float64) (int float64, frac float64)

func Nextafter(x, y float64) (r float64)

func Abs(x float64) float64

func Max(x, y float64) float64

func Min(x, y float64) float64

func Dim(x, y float64) float64

func Mod(x, y float64) float64

func Remainder(x, y float64) float64

func Sqrt(x float64) float64

func Cbrt(x float64) float64

func Hypot(p, q float64) float64

func Sin(x float64) float64

func Cos(x float64) float64

func Tan(x float64) float64

func Sincos(x float64) (sin, cos float64)

func Asin(x float64) float64

func Acos(x float64) float64

func Atan(x float64) float64

func Atan2(y, x float64) float64

func Sinh(x float64) float64

func Cosh(x float64) float64

func Tanh(x float64) float64

func Asinh(x float64) float64

func Acosh(x float64) float64

func Atanh(x float64) float64

func Log(x float64) float64

func Log1p(x float64) float64

func Log2(x float64) float64

func Log10(x float64) float64

func Logb(x float64) float64

func Ilogb(x float64) int

func Frexp(f float64) (frac float64, exp int)

func Ldexp(frac float64, exp int) float64

func Exp(x float64) float64

func Expm1(x float64) float64

func Exp2(x float64) float64

func Pow(x, y float64) float64

func Pow10(e int) float64

func Gamma(x float64) float64

func Lgamma(x float64) (lgamma float64, sign int)

func Erf(x float64) float64

func Erfc(x float64) float64

func J0(x float64) float64

func J1(x float64) float64

func Jn(n int, x float64) float64

func Y0(x float64) float64

func Y1(x float64) float64

func Yn(n int, x float64) float64

## math/big

用于大数计算

## math/complx

## math/rand

# mime 文件类型

func [AddExtensionType](https://github.com/golang/go/blob/master/src/mime/type.go?name=release#56)

```
func AddExtensionType(ext, typ string) error
```

函数将扩展名和mimetype建立偶联；扩展名应以点号开始，例如".html"。

func [FormatMediaType](https://github.com/golang/go/blob/master/src/mime/mediatype.go?name=release#21)

```
func FormatMediaType(t string, param map[string]string) string
```

函数根据[RFC 2045](http://tools.ietf.org/html/rfc2045)和 [RFC 2616](http://tools.ietf.org/html/rfc2616)的规定将媒体类型t和参数param连接为一个mime媒体类型，类型和参数都采用小写字母。任一个参数不合法都会返回空字符串。

func [ParseMediaType](https://github.com/golang/go/blob/master/src/mime/mediatype.go?name=release#101)

```
func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
```

函数根据[RFC 1521](http://tools.ietf.org/html/rfc1521)解析一个媒体类型值以及可能的参数。媒体类型值一般应为Content-Type和Conten-Disposition头域的值（参见[RFC 2183](http://tools.ietf.org/html/rfc2183)）。成功的调用会返回小写字母、去空格的媒体类型和一个非空的map。返回的map映射小写字母的属性和对应的属性值。

func [TypeByExtension](https://github.com/golang/go/blob/master/src/mime/type.go?name=release#45)

```
func TypeByExtension(ext string) string
```

函数返回与扩展名偶联的MIME类型。扩展名应以点号开始，如".html"。如果扩展名未偶联类型，函数会返回""。

内建的偶联表很小，但在unix系统会从本地系统的一或多个mime.types文件（参加下表）进行增补。

```
/etc/mime.types
/etc/apache2/mime.types
/etc/apache/mime.types
```

Windows系统的mime类型从注册表获取。文本类型的字符集参数默认设置为"utf-8"。

## mime/multipart

### type [File](https://github.com/golang/go/blob/master/src/mime/multipart/formdata.go?name=release#142)

```
type File interface {
    io.Reader
    io.ReaderAt
    io.Seeker
    io.Closer
}
```

File是一个接口，实现了对一个multipart信息中文件记录的访问。它的内容可以保持在内存或者硬盘中，如果保持在硬盘中，底层类型就会是*os.File。

### type [FileHeader](https://github.com/golang/go/blob/master/src/mime/multipart/formdata.go?name=release#122)

```
type FileHeader struct {
    Filename string
    Header   textproto.MIMEHeader
    // 内含隐藏或非导出字段
}
```

FileHeader描述一个multipart请求的（一个）文件记录的信息。

func (*FileHeader) [Open](https://github.com/golang/go/blob/master/src/mime/multipart/formdata.go?name=release#131)

```
func (fh *FileHeader) Open() (File, error)
```

Open方法打开并返回其关联的文件。







# net 网络库

## net/http

http包提供了HTTP客户端和服务端的实现。

Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。

```
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

程序在使用完回复后必须关闭回复的主体。

```
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：

```
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")
// ...
req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：

```
tr := &http.Transport{
	TLSClientConfig:    &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

Client和Transport类型都可以安全的被多个go程同时使用。出于效率考虑，应该一次建立、尽量重用。

ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

```
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```

要管理服务端的行为，可以创建一个自定义的Server：

```
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```







## net/http/httptest

func [NewRequest](https://github.com/golang/go/blob/master/src//net/http/httptest/httptest.go#L41)

```go
func NewRequest(method, target string, body io.Reader) *http.Request
```

NewRequest 返回一个新的服务器访问请求，这个请求可以传递给 http.Handler 以便进行测试。

target 参数的值为 [RFC 7230](http://tools.ietf.org/html/rfc7230) 中提到的“请求目标”（request-target)： 它可以是一个路径或者一个绝对 URL。如果 target 是一个绝对 URL，那么 URL 中的主机名（host name）将被使用；否则主机名将为 example.com。

当 target 的模式为 https 时，TLS 字段的值将被设置为一个非 nil 的随意值（dummy value）。

Request.Proto 总是为 HTTP/1.1。

如果 method 参数的值为空， 那么使用 GET 方法作为默认值。

body 参数的值可以为 nil；另一方面，如果 body 参数的值为 *bytes.Reader 类型、 *strings.Reader 类型或者 *bytes.Buffer 类型，那么 Request.ContentLength 将被设置。

为了使用的方便，NewRequest 将在 panic 可以被接受的情况下，使用 panic 代替错误。

如果你想要生成的不是服务器访问请求，而是一个客户端 HTTP 请求，那么请使用 net/http 包中的 NewRequest 函数。



type [ResponseRecorder](https://github.com/golang/go/blob/master/src/net/http/httptest/recorder.go?name=release#15)  记录响应内容

```
type ResponseRecorder struct {
    Code      int           // HTTP回复的状态码
    HeaderMap http.Header   // HTTP回复的头域
    Body      *bytes.Buffer // 如非nil，会将Write方法写入的数据写入bytes.Buffer
    Flushed   bool
    // 内含隐藏或非导出字段
}
```

ResponseRecorder实现了http.ResponseWriter接口，它记录了其修改，用于之后的检查。

```go
handler := func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "something failed", http.StatusInternalServerError)
}
req, err := http.NewRequest("GET", "http://example.com/foo", nil)
if err != nil {
    log.Fatal(err)
}
w := httptest.NewRecorder()
handler(w, req)
fmt.Printf("%d - %s", w.Code, w.Body.String())
```



type [Server](https://github.com/golang/go/blob/master/src/net/http/httptest/server.go?name=release#21)  一个测试服务器

```
type Server struct {
    URL      string // 格式为http://ipaddr:port，没有末尾斜杠的基地址
    Listener net.Listener
    // TLS是可选的TLS配置，在TLS开始后会填写为新的配置。
    // 如果在未启动的Server调用StartTLS方法前设置，已经存在的字段会拷贝进新配置里。
    TLS *tls.Config
    // Config可能会在调用Start/StartTLS方法之前调用NewUnstartedServer时被修改。
    Config *http.Server
    // 内含隐藏或非导出字段
}
```

Server是一个HTTP服务端，在本地环回接口的某个系统选择的端口监听，用于点对点HTTP测试。

```go
func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", greeting)
}

// Hello, client
```



















### 2. net/url

> request.Url 中可以用来解析一些东西

### 3. net/smtp

demo

```go
func main() {
	sendEmail("<h1> I'm Dad. </h1>", []string{"1141741348@qq.com"})
}

func sendEmail(htmlBody string, to []string) {
	user := "1164014200@qq.com"
	auth := smtp.PlainAuth("", "1164014200@qq.com", "授权码", "smtp.qq.com")
	//to := []string{"1141741348@qq.com"}

	nickname := "邮件AI"
	subject := "邮件轰炸"
	content_type := "Content-Type: text/html; charset=UTF-8"
	body := htmlBody

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
```

下面是低级 api

流程

```go
// Connect to the remote SMTP server.
c, err := smtp.Dial("mail.example.com:25")
if err != nil {
    log.Fatal(err)
}
// Set the sender and recipient first
if err := c.Mail("sender@example.org"); err != nil {
    log.Fatal(err)
}
if err := c.Rcpt("recipient@example.net"); err != nil {
    log.Fatal(err)
}
// Send the email body.
wc, err := c.Data()
if err != nil {
    log.Fatal(err)
}
_, err = fmt.Fprintf(wc, "This is the email body")
if err != nil {
    log.Fatal(err)
}
err = wc.Close()
if err != nil {
    log.Fatal(err)
}
// Send the QUIT command and close the connection.
err = c.Quit()
if err != nil {
    log.Fatal(err)
}
```



[type ServerInfo](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#ServerInfo)

[type Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Auth)

- [func CRAMMD5Auth(username, secret string) Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#CRAMMD5Auth)
- [func PlainAuth(identity, username, password, host string) Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#PlainAuth)

```go
// Set up authentication information.
auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
// Connect to the server, authenticate, set the sender and recipient,
// and send the email all in one step.
to := []string{"recipient@example.net"}
msg := []byte("This is the email body.")
err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
if err != nil {
    log.Fatal(err)
}
```

[type Client](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client)

- [func Dial(addr string) (*Client, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Dial)
- [func NewClient(conn net.Conn, host string) (*Client, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#NewClient)
- [func (c *Client) Extension(ext string) (bool, string)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Extension)
- [func (c *Client) Hello(localName string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Hello)
- [func (c *Client) Auth(a Auth) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Auth)
- [func (c *Client) Verify(addr string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Verify)
- [func (c *Client) StartTLS(config *tls.Config) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.StartTLS)
- [func (c *Client) Mail(from string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Mail)
- [func (c *Client) Rcpt(to string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Rcpt)
- [func (c *Client) Data() (io.WriteCloser, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Data)
- [func (c *Client) Reset() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Reset)
- [func (c *Client) Quit() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Quit)
- [func (c *Client) Close() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Close)

[func SendMail(addr string, a Auth, from string, to [\]string, msg []byte) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#SendMail)







## net/smtp 邮箱发

```go
// Connect to the remote SMTP server.
c, err := smtp.Dial("mail.example.com:25")
if err != nil {
    log.Fatal(err)
}
// Set the sender and recipient first
if err := c.Mail("sender@example.org"); err != nil {
    log.Fatal(err)
}
if err := c.Rcpt("recipient@example.net"); err != nil {
    log.Fatal(err)
}
// Send the email body.
wc, err := c.Data()
if err != nil {
    log.Fatal(err)
}
_, err = fmt.Fprintf(wc, "This is the email body")
if err != nil {
    log.Fatal(err)
}
err = wc.Close()
if err != nil {
    log.Fatal(err)
}
// Send the QUIT command and close the connection.
err = c.Quit()
if err != nil {
    log.Fatal(err)
}
```

[type ServerInfo](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#ServerInfo)

[type Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Auth)

- [func CRAMMD5Auth(username, secret string) Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#CRAMMD5Auth)
- [func PlainAuth(identity, username, password, host string) Auth](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#PlainAuth)

[type Client](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client)

- [func Dial(addr string) (*Client, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Dial)
- [func NewClient(conn net.Conn, host string) (*Client, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#NewClient)
- [func (c *Client) Extension(ext string) (bool, string)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Extension)
- [func (c *Client) Hello(localName string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Hello)
- [func (c *Client) Auth(a Auth) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Auth)
- [func (c *Client) Verify(addr string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Verify)
- [func (c *Client) StartTLS(config *tls.Config) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.StartTLS)
- [func (c *Client) Mail(from string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Mail)
- [func (c *Client) Rcpt(to string) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Rcpt)
- [func (c *Client) Data() (io.WriteCloser, error)](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Data)
- [func (c *Client) Reset() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Reset)
- [func (c *Client) Quit() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Quit)
- [func (c *Client) Close() error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#Client.Close)

[func SendMail(addr string, a Auth, from string, to [\]string, msg []byte) error](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm#SendMail)



### 6. net/http/cgi

cgi包实现了CGI（Common Gateway Interface，公共网关协议），参见[RFC 3875](http://tools.ietf.org/html/rfc3875)。

注意使用CGI意味着对每一个请求开始一个新的进程，这显然要比使用长期运行的服务程序要低效。本包主要是为了兼容现有的系统。

### 7. net/http/cookiejar

### 8. net/http/httptrace 追寻请求

### 9. net/http/pprof

pprof包通过它的HTTP服务端提供pprof可视化工具期望格式的运行时剖面文件数据服务。关于pprof的更多信息，参见http://code.google.com/p/google-perftools/。

本包一般只需导入获取其注册HTTP处理器的副作用。处理器的路径以/debug/pprof/开始。

要使用pprof，在你的程序里导入本包：

```
import _ "net/http/pprof"
```

如果你的应用还没有运行http服务器，你需要开始一个http服务器。添加"net/http"包和"log"包到你的导入列表，然后在main函数开始处添加如下代码：

```
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

然后使用pprof工具查看堆剖面：

```
go tool pprof http://localhost:6060/debug/pprof/heap
```

或查看周期30秒的CPU剖面：

```
go tool pprof http://localhost:6060/debug/pprof/profile
```

或查看go程阻塞剖面：

```
go tool pprof http://localhost:6060/debug/pprof/block
```

要查看所有可用的剖面，在你的浏览器阅读<http://localhost:6060/debug/pprof/>。要学习这些运转的设施，访问：

```
http://blog.golang.org/2011/06/profiling-go-programs.html
```









# sort 排序

sort

```go

type Person struct {
	Name string
	Age  int
}

//func (p Person) String() string {
//	return fmt.Sprintf("%s: %d", p.Name, p.Age)
//}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := ByAge{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
}
```

sortKeys

```go
// A couple of type definitions to make the units clear.
type earthMass float64
type au float64
// A Planet defines the properties of a solar system object.
type Planet struct {
    name     string
    mass     earthMass
    distance au
}
// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Planet) bool
// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(planets []Planet) {
    ps := &planetSorter{
        planets: planets,
        by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
    }
    sort.Sort(ps)
}
// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
    planets []Planet
    by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}
// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
    return len(s.planets)
}
// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
    s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}
// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
    return s.by(&s.planets[i], &s.planets[j])
}
var planets = []Planet{
    {"Mercury", 0.055, 0.4},
    {"Venus", 0.815, 0.7},
    {"Earth", 1.0, 1.0},
    {"Mars", 0.107, 1.5},
}
// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func Example_sortKeys() {
    // Closures that order the Planet structure.
    name := func(p1, p2 *Planet) bool {
        return p1.name < p2.name
    }
    mass := func(p1, p2 *Planet) bool {
        return p1.mass < p2.mass
    }
    distance := func(p1, p2 *Planet) bool {
        return p1.distance < p2.distance
    }
    decreasingDistance := func(p1, p2 *Planet) bool {
        return !distance(p1, p2)
    }
    // Sort the planets by the various criteria.
    By(name).Sort(planets)
    fmt.Println("By name:", planets)
    By(mass).Sort(planets)
    fmt.Println("By mass:", planets)
    By(distance).Sort(planets)
    fmt.Println("By distance:", planets)
    By(decreasingDistance).Sort(planets)
    fmt.Println("By decreasing distance:", planets)
    // Output: By name: [{Earth 1 1} {Mars 0.107 1.5} {Mercury 0.055 0.4} {Venus 0.815 0.7}]
    // By mass: [{Mercury 0.055 0.4} {Mars 0.107 1.5} {Venus 0.815 0.7} {Earth 1 1}]
    // By distance: [{Mercury 0.055 0.4} {Venus 0.815 0.7} {Earth 1 1} {Mars 0.107 1.5}]
    // By decreasing distance: [{Mars 0.107 1.5} {Earth 1 1} {Venus 0.815 0.7} {Mercury 0.055 0.4}]
}
```

sortMultikeys

```go
// A Change is a record of source code changes, recording user, language, and delta size.
type Change struct {
    user     string
    language string
    lines    int
}
type lessFunc func(p1, p2 *Change) bool
// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
    changes []Change
    less    []lessFunc
}
// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes []Change) {
    ms.changes = changes
    sort.Sort(ms)
}
// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
    return &multiSorter{
        less: less,
    }
}
// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
    return len(ms.changes)
}
// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
    ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}
// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
    p, q := &ms.changes[i], &ms.changes[j]
    // Try all but the last comparison.
    var k int
    for k = 0; k < len(ms.less)-1; k++ {
        less := ms.less[k]
        switch {
        case less(p, q):
            // p < q, so we have a decision.
            return true
        case less(q, p):
            // p > q, so we have a decision.
            return false
        }
        // p == q; try the next comparison.
    }
    // All comparisons to here said "equal", so just return whatever
    // the final comparison reports.
    return ms.less[k](p, q)
}
var changes = []Change{
    {"gri", "Go", 100},
    {"ken", "C", 150},
    {"glenda", "Go", 200},
    {"rsc", "Go", 200},
    {"r", "Go", 100},
    {"ken", "Go", 200},
    {"dmr", "C", 100},
    {"r", "C", 150},
    {"gri", "Smalltalk", 80},
}
// ExampleMultiKeys demonstrates a technique for sorting a struct type using different
// sets of multiple fields in the comparison. We chain together "Less" functions, each of
// which compares a single field.
func Example_sortMultiKeys() {
    // Closures that order the Change structure.
    user := func(c1, c2 *Change) bool {
        return c1.user < c2.user
    }
    language := func(c1, c2 *Change) bool {
        return c1.language < c2.language
    }
    increasingLines := func(c1, c2 *Change) bool {
        return c1.lines < c2.lines
    }
    decreasingLines := func(c1, c2 *Change) bool {
        return c1.lines > c2.lines // Note: > orders downwards.
    }
    // Simple use: Sort by user.
    OrderedBy(user).Sort(changes)
    fmt.Println("By user:", changes)
    // More examples.
    OrderedBy(user, increasingLines).Sort(changes)
    fmt.Println("By user,<lines:", changes)
    OrderedBy(user, decreasingLines).Sort(changes)
    fmt.Println("By user,>lines:", changes)
    OrderedBy(language, increasingLines).Sort(changes)
    fmt.Println("By language,<lines:", changes)
    OrderedBy(language, increasingLines, user).Sort(changes)
    fmt.Println("By language,<lines,user:", changes)
    // Output:
    // By user: [{dmr C 100} {glenda Go 200} {gri Smalltalk 80} {gri Go 100} {ken Go 200} {ken C 150} {r Go 100} {r C 150} {rsc Go 200}]
    // By user,<lines: [{dmr C 100} {glenda Go 200} {gri Smalltalk 80} {gri Go 100} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
    // By user,>lines: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken Go 200} {ken C 150} {r C 150} {r Go 100} {rsc Go 200}]
    // By language,<lines: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {ken Go 200} {glenda Go 200} {rsc Go 200} {gri Smalltalk 80}]
    // By language,<lines,user: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
}
```





# path 解析路径

[Variables](https://studygolang.com/static/pkgdoc/pkg/path.htm#pkg-variables)

[func IsAbs(path string) bool](https://studygolang.com/static/pkgdoc/pkg/path.htm#IsAbs)

[func Split(path string) (dir, file string)](https://studygolang.com/static/pkgdoc/pkg/path.htm#Split)

[func Join(elem ...string) string](https://studygolang.com/static/pkgdoc/pkg/path.htm#Join)

[func Dir(path string) string](https://studygolang.com/static/pkgdoc/pkg/path.htm#Dir)

[func Base(path string) string](https://studygolang.com/static/pkgdoc/pkg/path.htm#Base)

[func Ext(path string) string](https://studygolang.com/static/pkgdoc/pkg/path.htm#Ext)

[func Clean(path string) string](https://studygolang.com/static/pkgdoc/pkg/path.htm#Clean)

[func Match(pattern, name string) (matched bool, err error)](https://studygolang.com/static/pkgdoc/pkg/path.htm#Match)



# os 不依赖平台的操作

> 主要是 环境变量的操作，文件的 操作

还有

1. exec 执行命令
2. signal 信号的操作
3. user 用户id



下面是一个简单的例子，打开一个文件并从中读取一些数据：

```
file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
```

如果打开失败，错误字符串是自解释的，例如：

```
open file.go: no such file or directory
```

文件的信息可以读取进一个[]byte切片。Read和Write方法从切片参数获取其内的字节数。

```
data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
```

[func Hostname() (name string, err error)](https://studygolang.com/static/pkgdoc/pkg/os.htm#Hostname)

[func Getpagesize() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getpagesize)

[func Environ() [\]string](https://studygolang.com/static/pkgdoc/pkg/os.htm#Environ)

[func Getenv(key string) string](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getenv)

[func Setenv(key, value string) error](https://studygolang.com/static/pkgdoc/pkg/os.htm#Setenv)

[func Clearenv()](https://studygolang.com/static/pkgdoc/pkg/os.htm#Clearenv)

[func Exit(code int)](https://studygolang.com/static/pkgdoc/pkg/os.htm#Exit)

[func Expand(s string, mapping func(string) string) string](https://studygolang.com/static/pkgdoc/pkg/os.htm#Expand)

[func ExpandEnv(s string) string](https://studygolang.com/static/pkgdoc/pkg/os.htm#ExpandEnv)

[func Getuid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getuid)

[func Geteuid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Geteuid)

[func Getgid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getgid)

[func Getegid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getegid)

[func Getgroups() ([\]int, error)](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getgroups)

[func Getpid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getpid)

[func Getppid() int](https://studygolang.com/static/pkgdoc/pkg/os.htm#Getppid)









# go web 实战笔记

## 1. Http包

一个简单的web服务器

```go
package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie")
}

func main() {
    // 直接使用了包级别的。net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主handler，我们不需要将它传给ListenAndServe函数；nil值就可以工作。
    
    // 并且 每个 处理函数都是在一个新的 gorunti里执行的
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

Go的http有两个核心功能：Conn、ServeMux

```go
func main() {
    db := database{"shoes": 50, "socks": 5}
    mux := http.NewServeMux()  // 是http.hander接口
    mux.Handle("/list", http.HandlerFunc(db.list)) // 使用强制类型转换
    mux.Handle("/price", http.HandlerFunc(db.price))
    // 或者使用
    mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
    log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
```











## 2. xml解析

```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	//Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("test.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	fmt.Println(data)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)

}


// {{ servers} 1 [{{ server} Shanghai_VPN 127.0.0.1} {{ server} Beijing_VPN 127.0.0.2}]}

```

我们看到函数定义了两个参数，第一个是XML数据流，第二个是存储的对应类型，目前支持struct、slice和string，XML包内部采用了反射来进行数据的映射，所以v里面的字段必须是导出的。`Unmarshal`解析的时候XML元素和字段怎么对应起来的呢？这是有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会对应字段名。必须注意一点的是解析的时候tag、字段名、XML元素都是大小写敏感的，所以必须一一对应字段。

- 如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段,如上面例子中的servers。
- 如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段，如上servername和serverip定义。
- 如果某个struct字段的tag定义了中含有`",attr"`，那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义。
- 如果某个struct字段的tag定义 型如`"a>b>c"`,则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
- 如果某个struct字段的tag定义了`"-"`,那么不会为该字段解析匹配任何xml数据。
- 如果struct字段后面的tag定义了`",any"`，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
- 如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。

上面详细讲述了如何定义struct的tag。 只要设置对了tag，那么XML解析就如上面示例般简单，tag和XML的element是一一对应的关系，如上所示，我们还可以通过slice来表示多个同级元素。

## 3. xml 输出

假若我们不是要解析如上所示的XML文件，而是生成它，那么在go语言中又该如何实现呢？ xml包中提供了`Marshal`和`MarshalIndent`两个函数，来满足我们的需求。这两个函数主要的区别是第二个函数会增加前缀和缩进，函数的定义如下所示：

```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"shanghai", "113"})
	v.Svs = append(v.Svs, server{"beijing", "23"})
	output, err := xml.MarshalIndent(v, "  ", "	")
	if err != nil {
		fmt.Printf("error %v: ", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

======================
<?xml version="1.0" encoding="UTF-8"?>
  <servers version="1">
        <server>
                <serverName>shanghai</serverName>
                <serverIP>113</serverIP>
        </server>
        <server>
                <serverName>beijing</serverName>
                <serverIP>23</serverIP>
        </server>
  </servers>⏎ 
```

- XMLName不会被输出
- tag中含有`"-"`的字段不会输出
- tag中含有`"name,attr"`，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
- tag中含有`",attr"`，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
- tag中含有`",chardata"`，输出为xml的 character data而非element。
- tag中含有`",innerxml"`，将会被原样输出，而不会进行常规的编码过程
- tag中含有`",comment"`，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
- tag中含有`"omitempty"`,如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string
- tag中含有`"a>b>c"`，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出



## 4. Json 处理

### 1. 解析到 struct

```gp
package main

import (
	"encoding/json"
	"fmt"
)

const str = `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`



type Server struct {
    ServerName string `json:"serverName"`
    ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
    Servers []Server `json:"servers"`
}


type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

```

在解析的时候，如何将json数据与struct字段相匹配呢？例如JSON的key是`Foo`，那么怎么找对应的字段呢？

- 首先查找tag含有`Foo`的可导出的struct字段(首字母大写)
- 其次查找字段名是`Foo`的导出字段
- 最后查找类似`FOO`或者`FoO`这样的除了首字母之外其他大小写不敏感的导出字段

聪明的你一定注意到了这一点：能够被赋值的字段必须是可导出字段(即首字母大写）。同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题。

### 2. 解析到interface

JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：

- bool 代表 JSON booleans,
- float64 代表 JSON numbers,
- string 代表 JSON strings,
- nil 代表 JSON null.

```go
package main

import (
	"encoding/json"
	"fmt"
)

var a = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

func main() {
	var f map[string]interface{}
	json.Unmarshal(a, &f)
	fmt.Println(f["Name"].(string))
}

```



### 3. 生成Json

```go
type Server struct {
    // ID 不会导出到JSON中
    ID int `json:"-"`

    // ServerName2 的值会进行二次JSON编码
    ServerName  string `json:"serverName"`
    ServerName2 string `json:"serverName2,string"`

    // 如果 ServerIP 为空，则不输出到JSON串中
    ServerIP   string `json:"serverIP,omitempty"`
}

s := Server {
    ID:         3,
    ServerName:  `Go "1.0" `,
    ServerName2: `Go "1.0" `,
    ServerIP:   ``,
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)





package main

import (
	"encoding/json"
	"fmt"
)

var a = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

func main() {
	//var f map[string]interface{}
	//json.Unmarshal(a, &f)
	//fmt.Println(f["Name"].(string))
	//b, _ := json.Marshal(f)
	//fmt.Println(string(b))

	var test map[string]interface{} // = nil 默认是nil 所以要初始化
	test = make(map[string]interface{})

	test["a"] = struct {
		a string
		b string
	}{a: "a", b: "b"}
	test["b"] = 2
	b, _ := json.Marshal(test)
	fmt.Println(string(b))

}
```



## 5. go正则化

> 官网有语法

### 1. 是否匹配

```go
func Match(pattern string, b []byte) (matched bool, error error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, error error)
func MatchString(pattern string, s string) (matched bool, error error)


func (re *Regexp) Match(b []byte) bool
func (re *Regexp) MatchReader(r io.RuneReader) bool
func (re *Regexp) MatchString(s string) bool



func main() {
	fmt.Println(IsIP("127.0.0.1"))
}

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, ip); !m {
		return false
	}
	return true
}

```

### 2. 使用正则获取数据

```Go
func (re *Regexp) ReplaceAll(src, repl []byte) []byte
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
func (re *Regexp) ReplaceAllString(src, repl string) string
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error.")
		return
	}

	src := string(body)
	// 标签全部换成小写
	re, _ := regexp.Compile(`\<[\s\S]+?\>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除style
	re, _ = regexp.Compile(`\<style[\s\S]+?\</style\>`)
	src = re.ReplaceAllString(src, "")

	// 去除 script
	re, _ = regexp.Compile(`\<script[\s\S]+?\</script\>`)
	src = re.ReplaceAllString(src, "")

	fmt.Println(strings.TrimSpace(src))
}

```

常用的几个方法：

```Go
func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindSubmatchIndex(b []byte) []int
```



使用

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	re, _ := regexp.Compile(`[a-z]{2,4}`)
	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// 查找符合正则 的所有slice ， n < 0 表示返回全部符合的字符串，不然就是指定的长度
	all := re.FindAll([]byte(a), -1)
	for _, v := range all {
		fmt.Println(string(v))
	}

	// 查找 符合条件的 index位置，开始的位置和 结束的位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")
	// 类似 group 函数 , 第一个元素是匹配的全部元素，第二个是第一个（） 里面的，依次
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)

	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

}

```



## 6. go 模板处理

```Go
func handler(w http.ResponseWriter, r *http.Request) {
    t := template.New("some template") //创建一个模板
    t, _ = t.ParseFiles("tmpl/welcome.html", nil)  //解析模板文件
    user := GetUser() //获取当前用户信息
    t.Execute(w, user)  //执行模板的merger操作
}
```

Go语言的模板通过`{{}}`来包含需要在渲染时被替换的字段，`{{.}}`表示当前的对象，这和Java或者C++中的this类似，如果要访问当前对象的字段通过`{{.FieldName}}`,但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的),否则在渲染的时候就会报错，请看下面的这个例子：

```go
package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "jsy"}
	t.Execute(os.Stdout, p)
}
```

何来循环的输出这些内容呢？我们可以使用`{{with …}}…{{end}}`和`{{range …}}{{end}}`来进行数据的输出。

```go
package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := &Friend{Fname: "minux.ma"}
	f2 := &Friend{Fname: "xushiwei"}

	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
{{range .Emails}}
an email {{.}}
{{end}}
{{with .Friends}}
{{range .}}
	myfriend name is {{.Fname}}
{{end}}
{{end}}
`)

	p := Person{UserName: "jsy",
		Emails:  []string{"s@s", "2@s"},
		Friends: []*Friend{f1, f2}}

	t.Execute(os.Stdout, p)
}

```



条件处理

```Go
package main

import (
    "os"
    "text/template"
)

func main() {
    tEmpty := template.New("template test")
    tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
    tEmpty.Execute(os.Stdout, nil)

    tWithValue := template.New("template test")
    tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
    tWithValue.Execute(os.Stdout, nil)

    tIfElse := template.New("template test")
    tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
    tIfElse.Execute(os.Stdout, nil)
}
```



管道

```Go
{{. | html}}
```



模板变量

```Go
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
```

模板中使用函数

```Go
t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
```

其实，在模板包内部已经有内置的实现函数，下面代码截取自模板包里面

```Go
var builtins = FuncMap{
    "and":      and,
    "call":     call,
    "html":     HTMLEscaper,
    "index":    index,
    "js":       JSEscaper,
    "len":      length,
    "not":      not,
    "or":       or,
    "print":    fmt.Sprint,
    "printf":   fmt.Sprintf,
    "println":  fmt.Sprintln,
    "urlquery": URLQueryEscaper,
}
```



嵌套模板

我们平常开发Web应用的时候，经常会遇到一些模板有些部分是固定不变的，然后可以抽取出来作为一个独立的部分，例如一个博客的头部和尾部是不变的，而唯一改变的是中间的内容部分。所以我们可以定义成`header`、`content`、`footer`三个部分。Go语言中通过如下的语法来申明

```Go
{{define "子模板名称"}}内容{{end}}
```

通过如下方式来调用：

```Go
{{template "子模板名称"}}
```

接下来我们演示如何使用嵌套模板，我们定义三个文件，`header.tmpl`、`content.tmpl`、`footer.tmpl`文件，里面的内容如下

```html
//header.tmpl
{{define "header"}}
<html>
<head>
    <title>演示信息</title>
</head>
<body>
{{end}}

//content.tmpl
{{define "content"}}
{{template "header"}}
<h1>演示嵌套</h1>
<ul>
    <li>嵌套使用define定义子模板</li>
    <li>调用使用template</li>
</ul>
{{template "footer"}}
{{end}}

//footer.tmpl
{{define "footer"}}
</body>
</html>
{{end}}
```



## 7. 文件操作

### 目录操作

文件操作的大多数函数都是在os包里面，下面列举了几个目录操作的：

- func Mkdir(name string, perm FileMode) error

  创建名称为name的目录，权限设置是perm，例如0777

- func MkdirAll(path string, perm FileMode) error

  根据path创建多级子目录，例如astaxie/test1/test2。

- func Remove(name string) error

  删除名称为name的目录，当目录下有文件或者其他目录时会出错

- func RemoveAll(path string) error

  根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test", 0777)
	os.MkdirAll("test/test1/test2", 0777)
	err := os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("test")
}


```



### 文件操作

新建文件可以通过如下两个方法

- func Create(name string) (file *File, err Error)

  根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。

- func NewFile(fd uintptr, name string) *File

  根据文件描述符创建相应的文件，返回一个文件对象

通过如下两个方法来打开文件：

- func Open(name string) (file *File, err Error)

  该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。

- func OpenFile(name string, flag int, perm uint32) (file *File, err Error)

  打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

写文件函数：

- func (file *File) Write(b []byte) (n int, err Error)

  写入byte类型的信息到文件

- func (file *File) WriteAt(b []byte, off int64) (n int, err Error)

  在指定位置开始写入byte类型的信息

- func (file *File) WriteString(s string) (ret int, err Error)

  写入string信息到文件

读文件函数：

- func (file *File) Read(b []byte) (n int, err Error)

  读取数据到b中

- func (file *File) ReadAt(b []byte, off int64) (n int, err Error)

  从off开始读取数据到b中

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}


func main() {
    userFile := "asatxie.txt"
    fl, err := os.Open(userFile)        
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
}

```

Go语言里面删除文件和删除文件夹是同一个函数

- func Remove(name string) Error



## 8. 字符串处理

- func Contains(s, substr string) bool

  字符串s中是否包含substr，返回bool值

```Go
fmt.Println(strings.Contains("seafood", "foo"))
fmt.Println(strings.Contains("seafood", "bar"))
fmt.Println(strings.Contains("seafood", ""))
fmt.Println(strings.Contains("", ""))
```

- func Join(a []string, sep string) string

  字符串链接，把slice a通过sep链接起来

```Go
s := []string{"foo", "bar", "baz"}
fmt.Println(strings.Join(s, ", "))
//Output:foo, bar, baz
```

- func Index(s, sep string) int

  在字符串s中查找sep所在的位置，返回位置值，找不到返回-1

```Go
fmt.Println(strings.Index("chicken", "ken"))
fmt.Println(strings.Index("chicken", "dmr"))
//Output:4
//-1
```

- func Repeat(s string, count int) string

  重复s字符串count次，最后返回重复的字符串

```Go
fmt.Println("ba" + strings.Repeat("na", 2))
//Output:banana
```

- func Replace(s, old, new string, n int) string

  在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

```Go
fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
//Output:oinky oinky oink
//moo moo moo
```

- func Split(s, sep string) []string

  把s字符串按照sep分割，返回slice

```Go
fmt.Printf("%q\n", strings.Split("a,b,c", ","))
fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
fmt.Printf("%q\n", strings.Split(" xyz ", ""))
fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
//Output:["a" "b" "c"]
//["" "man " "plan " "canal panama"]
//[" " "x" "y" "z" " "]
//[""]
```

- func Trim(s string, cutset string) string

  在s字符串的头部和尾部去除cutset指定的字符串

```Go
fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
//Output:["Achtung"]
```

- func Fields(s string) []string

  去除s字符串的空格符，并且按照空格分割返回slice

```Go
fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
//Output:Fields are: ["foo" "bar" "baz"]
```



### 字符串的 append

字符串转化的函数在strconv中，如下也只是列出一些常用的：

- Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。

```Go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := make([]byte, 0, 100)
    str = strconv.AppendInt(str, 4567, 10)
    str = strconv.AppendBool(str, false)
    str = strconv.AppendQuote(str, "abcdefg")
    str = strconv.AppendQuoteRune(str, '单')
    fmt.Println(string(str))
}
```



### 把其他类型转化成 字符串

- Format 系列函数把其他类型的转换为字符串

```go
package main

import ( "fmt" "strconv" )

func main() { 
    a := strconv.FormatBool(false) 
    b := strconv.FormatFloat(123.23, 'g', 12, 64) 
    c := strconv.FormatInt(1234, 10) 
    d := strconv.FormatUint(12345, 10) 
    e := strconv.Itoa(1023) 
    fmt.Println(a, b, c, d, e) 
}
```



### 把字符串转化成其他类型

- Parse 系列函数把字符串转换为其他类型

```go
package main

import (
    "fmt"
    "strconv"
)
func checkError(e error){
    if e != nil{
        fmt.Println(e)
    }
}
func main() {
    a, err := strconv.ParseBool("false")
    checkError(err)
    b, err := strconv.ParseFloat("123.23", 64)
    checkError(err)
    c, err := strconv.ParseInt("1234", 10, 64)
    checkError(err)
    d, err := strconv.ParseUint("12345", 10, 64)
    checkError(err)
    e, err := strconv.Atoi("1023")
    checkError(err)
    fmt.Println(a, b, c, d, e)
}
```

## 9. web编程相关库

### 1. socket编程

常用的Socket类型有两种：流式Socket（SOCK_STREAM）和数据报式Socket（SOCK_DGRAM）。流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用。



#### parseIP 使用

```go
func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s ip-addr \n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	os.Exit(0)
}

```

#### TCPConn

```Go
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error)
```

`TCPConn`可以用在客户端和服务器端来读写数据。

还有我们需要知道一个`TCPAddr`类型，他表示一个TCP的地址信息，他的定义如下：

```Go
type TCPAddr struct {
    IP IP
    Port int
}
```

在Go语言中通过`ResolveTCPAddr`获取一个`TCPAddr`

```Go
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
```

- net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only),TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个).
- addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22".

#### TCP Client

```Go
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
```

- net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
- laddr表示本机地址，一般设置为nil
- raddr表示远程的服务地址

#### TCP Server

```Go
func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
func (l *TCPListener) Accept() (c Conn, err os.Error)
```

参数说明同DialTCP的参数一样。下面我们实现一个简单的时间同步服务，监听7777端口



实例

```go
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
```



### 2. websocket

### 3. rpc

1. god  Http rpc
2. tcp rpc
3. json rpc

> 具体的请看 go web编程

```
import "net/rpc"
```

rpc包提供了通过网络或其他I/O连接对一个对象的导出方法的访问。服务端注册一个对象，使它作为一个服务被暴露，服务的名字是该对象的类型名。注册之后，对象的导出方法就可以被远程访问。服务端可以注册多个不同类型的对象（服务），但注册具有相同类型的多个对象是错误的。

```go
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
```


事实上，方法必须看起来像这样：

```
func (t *T) MethodName(argType T1, replyType *T2) error
```

其中T、T1和T2都能被encoding/gob包序列化。这些限制即使使用不同的编解码器也适用。（未来，对定制的编解码器可能会使用较宽松一点的限制）

方法的第一个参数代表调用者提供的参数；第二个参数代表返回给调用者的参数。方法的返回值，如果非nil，将被作为字符串回传，在客户端看来就和errors.New创建的一样。如果返回了错误，回复的参数将不会被发送给客户端。

服务端可能会单个连接上调用ServeConn管理请求。更典型地，它会创建一个网络监听器然后调用Accept；或者，对于HTTP监听器，调用HandleHTTP和http.Serve。

想要使用服务的客户端会创建一个连接，然后用该连接调用NewClient。

更方便的函数Dial（DialHTTP）会在一个原始的连接（或HTTP连接）上依次执行这两个步骤。

生成的Client类型值有两个方法，Call和Go，它们的参数为要调用的服务和方法、一个包含参数的指针、一个用于接收接个的指针。

Call方法会等待远端调用完成，而Go方法异步的发送调用请求并使用返回的Call结构体类型的Done通道字段传递完成信号。

httprpc demo

server.go

```go
package main

import (
	"errors"
	"github.com/gpmgo/gopm/modules/log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Que, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Que = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

```

client.go

```go
package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Que, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Reply)
}

```



































































