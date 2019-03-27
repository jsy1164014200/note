main的返回值0表示成功，非0表示其他错误类型。

1. cin
2. cout
3. cerr
4. clog

标准库定义的名字都在std命名空间

| 类型        | 最小尺寸     |      |
| ----------- | ------------ | ---- |
| bool        | undefine     |      |
| char        | 8            |      |
| wchar_t     | 16           |      |
| char16_t    | 16           |      |
| char32_t    | 32           |      |
| short       | 16           |      |
| int         | 16           |      |
| long        | 32           |      |
| long long   | 64           |      |
| float       | 6位有效数字  |      |
| double      | 10位有效数字 |      |
| long Double | 10位有效数字 |      |
|             |              |      |
|             |              |      |
|             |              |      |
|             |              |      |
|             |              |      |
|             |              |      |
|             |              |      |

还有 unsigned int（简写unsigned） ，unsigned long

char（这种具体是后面哪一种由编译器确定） signed char unsigned char



### 类型转换

bool = int  // 0 or 1

int = bool  // 0 or 1

浮点与整形，可能会损失精度

无符号的数 <-- 超过范围，，，会取模余数

有符号的数 <-- 超过范围，，会未定义，可能不报错，可能崩溃，可能产生垃圾



```C++
int i = 100;
if (i) { // 类型转换
    i = 0
}
```

无符号数加上有符号数，后者会出现隐式转换



### 字面值常量

1. 20  十进制
2. 025 八进制
3. 0x14 十六进制
4. 0e1
5. .001

字符和字符串字面值

1. 'a'
2. "sdfjk"  会自动补上 '\0'

转义字符

1. \1234 八进制的 \后面的3位，所以分成了 \123  4 
2. \x1234 十六进制的转义字符，后面的 4位都会被带上，而且超过了会丢弃

通过前后缀改变字面值的类型

1. L'a'   ---> wchar_t
2. u8"ha"  ----> utf8编码
3. 43ULL unsigned long long 
4. 1E-3F float
5. 3.23L long double 



### string

C++ std:string 是可变长度的字符序列类型





### 初始化与赋值

```c++
// 普通初始化
int units_sold = 0;
int units_sold3(0);
// 列表初始化 (如果存在精度损失会报错)
long long int haha = 1;
int ubnits_sold{haha}; // error : compile error
int units_sold1{0};
int units_sold2 = {0};
```

在函数外面未初始化的会被默认初始化为0

对与在函数内部未初始化的 是未定义（即值不确定）

### 区分声明和定义

声明只是被程序所知道

> extern int i;  //声明
>
> extern int i = 10;  // 定义



显式调用外部变量

```c
int c;
int main() {
    int c = 10;
    std::cout<<c;
    std::cout<<::c;
    return 0;
}
```







### 复合类型（指针 引用）

引用必须被初始化

1. 引用绑定后 不能再重新绑定到另一个值
2. 引用只能绑定在 对象上，不能给字面量绑定值

> int ival = 1024;
>
> int &refVal = ival;



指针是一个对象（基本类型）

1. 可以不初始化



空指针

1. int *p1 = nullptr; // === int *p1=0;
2. int *p2 = NULL;

if (pi)   // 如果pi是空指针 那么就是false，其他都是true



void*指针（能做的事比较有限）

1. 跟其他指针比较
2. 输入输出
3. 赋值给另一个 void*
4. 不能操作其指的对象



指针的引用（从右往左读 容易理解）

```C++
    int i = 42;
    int *p = nullptr;
    int *&r = p;
    r = &i;
    *r = 0;
    std::cout<<i;
    return 0;

//0
```





### const

对于const变量，如果要想在多个文件中都能使用那么必须加extern

默认const会被编译器在本文件中替换，这样就不能共享了。



对常量的引用（常量引用 只是说 不能通过引用来修改，const引用可以指向一个非const)

```c++
const int ci = 1;
const int &r = ci;
r = 42; // error 不能更改值
int &r = ci; //error 不能用非常量引用引用 const
```

常量引用的几个特例

```c
   //int &a = 10;  // 不合法
    const int &b = 10; // 正确
```



同样，对于常量指针，他们不能修改原来的值

> double a = 1.0;
>
> const double  *ptr = a;
>
> // 但是指针指向的值 可以修改
>
> double b = 2.0;
>
> ptr = b;

总的来说，指向常量的指针，引用只是他们自以为 不能修改原来的值，所以不能修改，





const 指针

因为 指针也是一个基本类型对象，所以能const

> int errNumb =0;
>
> int *const curErr = &errNumb;
>
> const int *const c = &errNumb;



这样一来

1. *curErr = 10; 正确
2. *c = 20; 不正确



顶层指针是指本身是一个const，底层指针是说 指向的对象是一个常量

但是用于声明 const int &r = ai; 引用都是底层cosnt，因为引用本身就不能修改。





常量表达式

如何确定？在编译一开始就能知道值的就是 常量表达式

1. const int i = 10 // shi
2. int i = 10 //bushi
3. const int sz = get_size() ; //bushi

constexpr 语句

1. constexpr int mf = 20;
2. constexpr int limit = mf +1;
3. constexpr int sz = size();  // 只用当size是一个constexpr表达式时才是正确的

一个constexpr 常量表达式 指针一定是 constexpr int *p = nullptr;



常量表达式会将对象置为顶层的const

```c++
int j = 0;
constexpr int i = 2;

int main() {

    constexpr const int *p = &i;
    constexpr int *p1 = &j;

    return 0;
}
```





### typedef 

同 using

using test = double ; == typedef double test;

### auto类型推断

auto一般会忽略掉顶层const，底层的const则会保留下来

所以如果要使推断出来的类型是一个顶层const

> const auto f =ci;

例如

1. const int ci = i;
2. auto &n = i, *p2 = &ci;  // 错误，因为 i 是 int ，，p2 是const int 



### decltype类型指示符

> decltype(f()) sum = x; 根据f的返回值来确定返回类型，，，，不会实际调用f

处理顶层const 与 底层const的方式与 auto不同

1. const int ci = 0,&cj = ci;
2. decltype(ci) x = 0; // x类型是int
3. decltype(cj) y = x;// y是const int&，y绑定到x
4. decltype(cj) z; //错误，z是一个引用

两个关键点

1. 

decltype(*p) 是表达式，所以得到的是 引用类型

2. 

decltype( xxx ) 得到的是xxx的类型

而decltype( (xxx) ) 得到的是表达式的值，永远都是 引用





### struct 

11标准给定 默认会初始化 成员

```c
struct Sales_data {
    std::string bookNo;
    unsigned units_sold = 0;
    double revenue = 0.0;
} accum, trans, *salesptr;
```

头文件的名字与类名一样