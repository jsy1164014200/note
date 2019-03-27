# erlang 趣学指南之函数式编程

1. 没有变量
2. 没有赋值 只有模式匹配





> 记得句尾加.

1. help(). 
2. erlang中的shell 是一组 shell 实例
3. ctrl + g > h 进入shell 帮助

## 基础知识

#### 1. 数值类型（类似python，不关心float int）

```erlang
2 + 15.
17
5 / 2.
2.5
5 div 2.
2
5 rem 2.
1

进制表示法
2#101010.
42
8#0677.
447
16#AE
174
```

#### 2. 不变的变量

```shell
One = 1.
1
Un = Uno = One = 1.
1
Two = One + One 
2
Two = 2.
2
Two = Two + 1
** exception error : no match of right hand side value 3
= 号（不是变量）负责比较 操作 并且在不相等时 报错
如果等号两边都是变量，并且走边是个未绑定的变量（没有任何值与它关联），erlang会自动把右边的值绑定到左边的变量上

变量不能小写
two = 2. 报错

_ 表示 不关心其值（类golang）

```

在shell 中可以用 f(Variable). 来删除某个变量，f(). 删除所有。

#### 3. 原子atom

原子是字面量， 唯一的值就是它自己（你看到的就是你能得到的，没有更多了）

```shell
> atom.
atom
> atoms_rule.
atoms_rule
> 'Atoms can be cheated'.
'Atoms can be cheated'
> atom = 'atom'.
atom
```

原子如果不以小写字符开头，或者其中包含有除字母，数字，下划线以及@符号之外的其他字符，那么必须放在' ' 之间。![选区_128](/home/jsy/图片/选区_128.png)

保留字：`after` `and` `andalso` `band` `begin` `bnot` `bor` `bsl` `bsr`  `bxor` `case` `catch` `cond` `div` `end` `fun` `if `  `let` `not` `of` `or` `orelse` `query` `receive` `rem` `try` `when` `xor` 

#### 4. 布尔代数和比较操作符

**要理解，true，false 只是特殊的atom**

```shell
> true and false
true
> false or true
true
> true xor false 
true
> not false
true
> not (true and true)
false 

6> 5 =:= 5.
true
7> 1 =:= 0.
false
8> 1 =/= 0.
true
9> 5 =:= 5.0.
false
10> 5 == 5.0.
true
11> 5 /= 5.0.
false
如果你不想区分，可以使用==和/=操作符。因此，关键在于要想清楚是否需要精确的相等性。根据经验，你应该始终都用=:=和=/=，只有在明确知道确实不需要精确相等性时才换成==和/=。

12> 1 < 2.
true
13> 1 < 1.
false
14> 1 >= 1.
true
15> 1 =< 1.
true
特别注意最后一个，是倒写的。
```

> 注意：布尔操作符不同于 js，python golang 里面的，不会短路运算，如果要短路运算，可以用 andalso , orelse

```shell
14> 0 == false.
false
15> 1 < false.
true
Erlang 中并没有布尔值 true 和 false 之类的东西。数据项 true 和 false 都是原子，只要你始终认为 false 和 true 只能表示字面意思
```

> 注意 比较操作中，数据类型乊间的大小顺序是： number < atom < reference < fun< port < pid < tuple < list < bit string。

#### 5. 元组

{Element1, Element2,....,ElementN}

```shell
1> X = 10, Y = 4.
4
2> Point = {X,Y}.
> f(Point).
ok

3> Point = {4,5}.
{4,5}
4> {X,Y} = Point.
{4,5}
5> X.
4
6> {X,_} = Point.
{4,5}

模式匹配（过程式语言中的赋值）

```

元组可以嵌套，里面可以方任意值。

#### 6. 列表

列表中可以包含任何东西—数值、原子、元组以及其他列表—所有能想象得到的东西都可以放到这个数据结构中。

```shell
1> [1, 2, 3, {numbers,[4,5,6]}, 5.34, atom].
[1,2,3,{numbers,[4,5,6]},5.34,atom]

2> [97, 98, 99].
"abc"
erlang中，字符串就是列表。
3> [97,98,99,4,5,6].
[97,98,99,4,5,6]
4> [233].
"é"
在打印数值列表时，只有当其中至少有一个数字代表的不是字符时， Erlang 才会把它们作为数值打印出来。在 Erlang 中，根本就没有真正的字符串！
```

++  -- 操作符

```shell
5> [1,2,3] ++ [4,5].
[1,2,3,4,5]
6> [1,2,3,4,5] -- [1,2,3].
[4,5]
7> [2,4,2] -- [2,4].
[2]
8> [2,4,2] -- [2,4,2].
[]

右结合
9> [1,2,3] -- [1,2] -- [3].
[3]
10> [1,2,3] -- [1,2] -- [2].
[2,3]
```

列表中第一个元素称为 头（head），剩余部分称为尾（tail）

```shell
使用bif (biuld in function) 获取它
> hd([1,2,3,4]).
1
> tl([1,2,3,4]).
[2,3,4]

Erlang 通过模式匹配提供了一种简单的方式来分离列表的头和尾： [Head|Tail]。
13> list = [2,3,4].
[2,3,4]
14> NewList = [1|List].
[1,2,3,4]
15> [Head|Tail] = NewList.
[1,2,3,4]
16> Head.
1
17> Tail.
[2,3,4]
18> [NewHead|NewTail] = Tail.
[2,3,4]
19> NewHead.
2
```

任何列表都可以用下面这个公式构建： [Term1 | [Term2| [... | [TermN]]]]。
所以，可以把列表递归地定义成一个头元素，后面跟着一个尾，而尾本身又是一个头元素后面跟
着更多的头。

```shell
20> [1 | []].
[1]
21> [2 | [1 | []]].
[2,1]
22> [3 | [2 | [1 | []]]].
[3,2,1]
```

##### 区分良构列表与非良构列表：

以[1|2]这种形式构建出来的列表称为非良构列表（ improper list）。非良构列表可以用于[Head|Tail]这种形式的模式匹配，但是在 Erlang 的标准函数（即使是 length()）中使用时会失败。这是因为 Erlang 期望的是良构列表（ proper list）。良构列表的最后一个元素是空列表。当定义一个像[2]这样的数据项时，这个列表会被自动地转换成良构形式。同样， [1|[2]]也会被转换成良构形式。尽管非良构列表是合乎语法的，但是除了某些用户自定义数据结构，它们的用途非常有限

> 总结 [] 会自动转为 良构，[1|2] 就规定了2 是尾，所以是非良构

##### 列表推导

```shell
1> [2*N || N <- [1,2,3,4]].
[2,4,6,8]

增加约束条件，可以有多个条件，用，号隔开
2> [X || X <- [1,2,3,4,5,6,7,8,9,10], X rem 2 =:= 0].
[2,4,6,8,10]

Pattern <- List（生成器表达式） , Condtion1,Condtion2,....
3> RestaurantMenu = [{steak, 5.99}, {beer, 3.99}, {poutine, 3.50}, {kitten, 20.99},
{water, 0.00}].
[{steak,5.99},
{beer,3.99},
{poutine,3.5},
{kitten,20.99},
{water,0.0}]
4> [{Item, Price*1.07} || {Item,Price} <- RestaurantMenu, Price >= 3, Price =< 10].
[{steak,6.409300000000001},{beer,4.2693},{poutine,3.745}]

多个生成器表达式
[X+Y || X <- [1,2], Y <- [3,4]].
[4,5,5,6]
```

NewList = [Expression || GeneratorExp1, GeneratorExp2, ..., GeneratorExpN,
Condition1, Condition2, ... ConditionM]

#### 7. 二进制数据

在 Erlang 的位（bit）语法中，会把二进制数据用<<和>>括起来，并把数据分隔成多个易理解的区段，区段之间用逗号分隔。一个区段就是一个二进制的位序列（不一定要在字节边界上，尽管默认情况下会这样）。

```shell
1> Color = 16#F09A29.
15768105
2> Pixel = <<Color:24>>.
<<240,154,41>>
```

八位 显示成一个字节。

```erlang
3> Pixels = <<213,45,132,64,76,32,76,0,0,234,32,15>>.
<<213,45,132,64,76,32,76,0,0,234,32,15>>
4> <<Pix1,Pix2,Pix3,Pix4>> = Pixels.
** exception error: no match of right hand side value <<213,45,132,64,76,32,76,0,0,
234,32,15>>
5> <<Pix1:24, Pix2:24, Pix3:24, Pix4:24>> = Pixels.
<<213,45,132,64,76,32,76,0,0,234,32,15>>
6> <<R:8, G:8, B:8>> = <<Pix1:24>>.
<<213,45,132>>
7> R.
213
```

其他的语法糖

```erlang
8> <<R:8, Rest/binary>> = Pixels.
<<213,45,132,64,76,32,76,0,0,234,32,15>>
9> R.
213
```

Value
Value:Size
Value/TypeSpecifierList
Value:Size/TypeSpecifierList

类型
可能的取值为 integer、 float、 binary、 bytes、 bitstring、 bits、 utf8、
utf16 和 utf32。如果不指定类型， Erlang 就会默认使用 integer 类型。
这个字段表示所使用的二进制数据的类型。注意， bytes 是 binary 的缩写， bits 是
bitstring 的缩写。
符号类型
可能的取值为 signed 和 unsigned。默认是 unsigned。只有当类型为 integer 时，
这个字段才有意义。
字节序
可能的取值为 big、 little 和 native。字节序的默认值是 big，因为它是网络协议
编码中使用的标准字节序。
只有当类型为 integer、 utf16、 utf32 或者 float 时，字节序才有意义。这个选
项和系统如何读取二进制数据有关。例如， BMP 图片文件的头格式中用 4 个字节的整数来
表示文件大小。如果一个文件的大小是 72 个字节，那么一个小字节序（little-endian）系统会
把 它 表 示 成 <<72,0,0,0>> ， 而 一 个 大 字 节 序 （ big-endian ） 系 统 会 把 它 表 示 成
<<0,0,0,72>>。在 Erlang 中，前者会解释成 72，后者则会被解释成 1 207 959 552

单位
写成 unit:Integer 这样的形式。
单位指的是每个区段的大小。取值范围为1～256。对于integer、float和bitstring
类型来说，单位默认设置为 1；对于 binary 类型，单位默认设置为 8。类型 utf8、 utf16
和 utf32 无需定义 unit。 Size 和单位的乘积等于要提取的区段中的位数，并且必须能被 8
整除。单位的大小通常用来确保字节对齐。
数据类型的默认长度可以通过组合不同的二进制区段描述加以改变。例如，
<<25:4/unit:8>>会把数值 25 编码成一个 4 字节的整数 ，其形象的表示就是
<<0,0,0,25>>。 <<25:2/unit:16>>会得到同样的结果， <<25:1/unit:32>>也一样。
一般来讲， Erlang 会接受 <<25:Size/unit:Unit>>这样的形式，并用 Size 乘以 Unit
算出表示这个值需要的空间。同样，空间的大小要能被 8 整除。

```shell
10> <<X1/unsigned>> = <<-44>>.
<<"Ô">>
11> X1.
212
12> <<X2/signed>> = <<-44>>.
<<"Ô">>
13> X2.
-44
14> <<X2/integer-signed-little>> = <<-44>>.
<<"Ô">>
15> X2.
-44
16> <<N:8/unit:1>> = <<72>>.
<<"H">>
17> N.
72
18> <<N/integer>> = <<72>>.
<<"H">>
19> <<Y:4/little-unit:8>> = <<72,0,0,0>>.
<<72,0,0,0>>
20> Y.
72
```

Erlang 中也提供了标准的二进制操作（向左移位、向右移位以及二进制的 and、 or、 xor和 not）。相应的操作符是 bsl（按位左移）、 bsr（按位右移）、 band、 bor、 bxor 以及 bnot。



二进制字符串

二进制字符串的语法是<<"this is a binary string!">>。与列表相比，二进制字符
串的缺点是，它的模式匹配和操作方法没有列表那么简单。因此，通常当所存储的文本无需频繁操作或者空间效率是个实际问题时，才会使用二进制字符串。

> 注意 尽管二迚制字符串非常轻量，还是不要用它们去标记数据。例如，使用字符串常量去表达{<<"temperature">>,50}是很有吸引力的，不过，对于这种情况，应该永进使用原子。如果使用原子，那么在对不同的值迚行比较时几乎没有开销，并且不管原子多长，比较总可以在常量时间完成，而二迚制字符串的比较则是线性时间的。另一方面，也不要因为原子更轻量，就用原子取代字符串。字符串是可以运算的（分隔、正则表达式等），而原子除了可以比较，什么都不能做

二进制列表推导式

2> Pixels = <<213,45,132,64,76,32,76,0,0,234,32,15>>.
<<213,45,132,64,76,32,76,0,0,234,32,15>>
3> RGB = [ {R,G,B} || <<R:8,G:8,B:8>> <= Pixels ].
[{213,45,132},{64,76,32},{76,0,0},{234,32,15}]



## 模块

```shell
1> erlang:element(2, {a,b,c}).
b
2> element(2, {a,b,c}).
b
3> lists:seq(1,4).
[1,2,3,4]
4> seq(1,4).
** exception error: undefined shell command seq/2
```

erlang 模块中的会被 自动引入，叫bif函数

所有模块属性都采用-Name(Attribute).的形式。要让你的模块能够编译，下面的模块属性必须定义：

-module(name).

其中 Name 是一个原子。你调用其他模块中的函数时，使用的就是这个名字。函数调用采用的形式是 M:F(A)，其中 M 是模块名， F是函数名， A 是参数。

要想导出函数，需要使用另一个属性：
-export([Function1/Arity, Function2/Arity,..., FunctionN/Arity]).(Arity 元数，表示函数能接受的参数个数)

函数定义的语法遵循 Name(Args)-> Body.这样的形式， Name 必须是一个原子， Body可以是一个或者多个用逗号分隔的 Erlang 表达式。函数以一个句点结束。注意，和许多命令式语言不同， Erlang 中没有 return 关键字。 return 毫无用处！无需显式说明，函数中最后一个表达式的执行结果会被自动作为返回值传递给调用者。

add(A,B) ->

​	A + B.

#### 编译代码

Erlang 代码会被编译成字节码，这样 VM 就能执行它了。调用编译器有多种方法。最常用的一种是在命令行中调用它，例如：
$ erlc flags file.erl
如果是在 shell 或者模块中，可以像这样编译代码：
compile:file(Filename)
还有一种方法，在开发代码时经常使用，就是在 shell 中编译：
c()



举例在shell 中：

```shell
1> cd("/path/to/where/you/saved/the-module/").
"Path Name to the directory you are in"
ok
2> c(useless).
{ok,useless}
3> useless:add(7,2).
9
4> useless:hello().
Hello, world!
ok
5> useless:greet_and_add_two(-3).
Hello, world!
-1
```

这是因为 Erlang 中的函数和表达式必须要有返回值，虽然在其他语言中可能并不
需要。正因为如此， io:format/1 函数返回 ok 表示情况正常：没有错误发生。

#### 编译选项

Erlang 提供了许多编译选项，用来对一个模块的编译方式进行控制。你可以在 Erlang 文档中
看到所有的编译选项。下面列出的是一些最常用的选项。
-debug_info
像调试器、代码覆盖率统计以及静态分析之类的 Erlang 工具都是使用模块中的调试信息
来完成工作的。一般来说，建议这个编译选项一直开启。比起不开启这个选项所节省的那一
点点字节码空间来说，你可能更需要这个选项所带来的好处。
-{outdir,Dir}
默认情况下， Erlang 编译器会将生成的.beam 文件放置到当前目录下。可以用这个选项
指定编译文件的存放路径。
-export_all
这个选项会让编译器忽略文件中已定义的-export 模块属性，把文件中的所有函数都
导出。这个选项在测试和开发新代码时特别有用，但在产品代码中严禁使用。
-{d,Macro}和{d,Macro,Value}
这个选项定义了一个可以在模块中使用的宏，其中 Macro 是个原子。这个选项在单元
测试中用得最多，因为它能确保模块中的测试函数只在明确需要时才会被创建和导出。如果
元组中没有定义第三个元素， Value 会被默认设置为 true。
在编译 useless 模块时，如果想使用编译选项，可以通过如下两种方式：
7> compile:file(useless, [debug_info, export_all]).
{ok,useless}
8> c(useless, [debug_info, export_all]).
{ok,useless}

还可以在模块内部通过模块属性来定义编译选项。要达到与前面第 7 行和第 8 行一样的编译效果，可以在模块中增加如下内容：
-compile([debug_info, export_all]).



#### 使用宏定义

-define(MACRO, some_value).

变量宏 -define(HOUR, 3600). % 单位是秒

函数宏 -define(sub(X,Y), X-Y).

Erlang 中有一些预定义的宏，下面列出了其中的一部分：
 ?MODULE，会被替换成当前模块的名字，是一个原子；
 ?FILE，会被替换成当前文件的名字，是一个字符串；
 ?LINE，会被替换成该宏所在的代码行的行号。

```shell
-ifdef(DEBUGMODE).
-define(DEBUG(S), io:format("dbg: "++S)).
-else.
-define(DEBUG(S), ok).
-endif.

-ifdef(TEST).
my_test_function() ->
run_some_tests().
-endif.
然后，使用前面介绍的编译选项，我们可以选择是否定义 DEBUGMODE 或者 TEST，就像这
样： c(Module, [{d,'TEST'},{d,'DEBUGMODE'}]).
```



#### 元数据

```erlang
9> useless:module_info().
[{exports,[{add,2},
{hello,0},
{greet_and_add_two,1},
{module_info,0},
{module_info,1}]},
{imports,[]},
{attributes,[{vsn,[174839656007867314473085021121413256129]}]},
{compile,[{options,[]},
{version,"4.8"},
{time,{2013,2,13,2,56,32}},
{source,"/home/ferd/learn-you-some-erlang/useless.erl"}]}]
10> useless:module_info(attributes).
[{vsn,[174839656007867314473085021121413256129]}]
```



##  *函数

```erlang
greet(male,Name) -> 
    io:format("Hello, Mr. ~s!", [Name]);
greet(female, Name) -> 
    io:format("Hello, Mrs. ~s!", [Name]);
greet(_, Name) -> 
    io:format("Hello, ~s!", [Name]).
                  
```

> io:format 函数的格式化输出是通过替换字符串中的标记符来完成的。波浪号（ ~）字符用来指示一个标记符。有些标记符是内置的，如~n，会被替换成一个换行符。其他标记符几乎都是用来指示数据的格式化方法的。例如，函数调用 io:format("~s!~n",["Hello"]).中包含了标记符~s，它接受字符串和二迚制字符串作为参数。最终输出的消息是"Hello!\n"。另外一个广泛使用的标记符是~p，它会打印一个 Erlang 数据项，打印的方式和 Erlang shell 打印数据项的方式完全一样（增加必要的缩迚和其他格式）

在 shell 中出现模式不匹配的情况时， Erlang 会很生气地抛出来一条出错消息。而当函数中的一个模式匹配失败时（如 greet(male,Name)）， Erlang 会继续查找函数的下一个模式（这个例子中，就是 greet(female, Name)）进行匹配，如果匹配成功，就运行该匹配对应的语句。

其中的每一条函数声明都被称作一个函数子句（function clause）。函数子句之间必须用分号（;）分隔，所有函数子句一起形成一个完整的函数定义。可以把整个函数定义看作是一个更大的语句，这也是为什么最后一个函数子句以句点结束的原因。



#### 模式匹配

```erlang
same(X,X) ->
true;
same(_,_) ->
false.

当调用 same(a,a)时，第一个 X 未绑定，所以会自动获取值 a。然后，当 Erlang 处理第二
个参数时，发现 X 已经绑定了。于是， Erlang 就把 X 已经绑定的值和作为第二个参数传递给函数
的 a 进行比较，检查它们是否匹配。模式匹配成功，所以函数返回 true。如果两个值不同，模
式匹配失败，
```



```erlang
valid_time({Date = {Y,M,D}, Time = {H,Min,S}}) ->
io:format("The Date tuple (~p) says today is: ~p/~p/~p,~n",[Date,Y,M,D]),
io:format("The time tuple (~p) indicates: ~p:~p:~p.~n", [Time,H,Min,S]);
valid_time(_) ->
io:format("Stop feeding me wrong data!~n").
```

#### 卫语句

```erlang
old_enough(X) when X >= 16 -> true;
old_enough(_) -> false.
```

卫表达式有一条基本规则，要想成功，它必须返回 true。如果返回了 false 或者抛出了异常，就表明卫语句失败。

andalso orelse（短路运算符） 相当于卫语句中的 ,  ;   （或者说 ;更像or这个非短路运算符）

但是也存在区别：如果卫语句 X >= N;N >= 0 的前半部分在执行时抛出了异常，后半部分仍然会被求值。如果 X >= N orelse N >= 0 的前半部分抛出了异常，那么后半部分就会被跳过，整个卫语句就会失败。

在卫语句中，只有 andalso 和 orelse可以嵌套使用。也就是说， (A orelse B) andalso C 是一个有效的卫语句，而(A; B),C 不是。

```erlang
right_age(X) when X >= 16, X =< 104 ->
	true;
right_age(_) ->
	false.

或者
wrong_age(X) when X < 16; X > 104 ->
    true;
wrong_age(_) ->
	false.
```



#### if 语句（卫模式）

```erlang
    1 -module(what_the_if).
    2 -export([heh_fine/0]).
    3 
    4 heh_fine() ->
    5     if 1 =:= 1 ->
    6            works
    7     end, 
    8     if 1 =:= 2; 1 =:= 1 ->
    9            works
   10     end, 
   11     if 1 =:= 2, 1 =:= 1 ->  %% 报错，永远都匹配不到
   12            fails
   13     end.


oh_god(N) ->
if N =:= 2 -> might_succeed;
true -> always_does %% 这是 Erlang if 的'else!'
end.
```

if elseif  else

```erlang
>> 16 help_me(Animal) ->
   17     Talk = if Animal == cat -> "meow";
   18               Animal == beef -> "bark";
   19               true -> "fgdadfgna"
   20            end,
   21     {Animal, "says" ++ Talk ++ "!"}.
```

无论是 else 还是 true 分支都应该极力避免。写 if 表达式时，如果能够覆盖
所有的逻辑分支，而不是依赖于一个包罗一切的分支，通常会更容易理解。



#### case of语句

```erlang

  1 -module(cases).
  2 -export([insert/2]).
  3 
  4 insert(X,[]) -> 
  5     [X];
  6 insert(X,Set) ->
  7     case lists:member(X,Set) of
  8         true -> Set;
  9         false -> [X|Set]
 10     end.


beach(Temperature) ->
    case Temperature of
        {celsius, N} when N >= 20, N =< 45 ->
            'favorable';
        {kelvin, N} when N >= 293, N =< 318 ->
            'scientifically favorable';
        {fahrenheit, N} when N >= 68, N =< 113 ->
            'favorable in the US';
        _ ->
            'avoid beach'
    end.

%% 也可以写成
beachf({celsius, N}) when N >= 20, N =< 45 ->
	'favorable';
...
beachf(_) ->
	'avoid beach'.
```

#### 如何选择（if，case，函数）

## erlang类型系统

#### 动态强类型

一般会认为静态类型语言比动态类型语言更加安全。虽然对于多数其他动
态类型语言来说这种观点也许是对的，但是对于 Erlang，这个观点并不成立，并且有真实的证据可以证明这一点。

尽管绝大多数语言和类型系统都旨在写出没有错误的程序，但是 Erlang 却认为错
误肯定会发生，所以在语言中提供了一些特性，基于这些特性可以很容易对错误进行平滑处理，并且不会造成不必要的停机时间。所以， Erlang 的动态类型系统不是程序可靠性和安全性的障碍。

Erlang 同时还是一个强类型语言。弱类型语言会在不同的数据项之间做隐式的类型转换。例如，如果 Erlang 是弱类型的，就能支持这个操作： 6 = 5 + "1"。但是，因为 Erlang 是强类型的，所以执行这样的操作会引发不正确参数（bad argument）异常

#### 类型转换

```erlang
> erlang:list_to_integer("54").
54
2> erlang:integer_to_list(54).
"54"
3> erlang:list_to_integer("54.32").
** exception error: bad argument
4> erlang:list_to_float("54.32").
54.32
5> erlang:atom_to_list(true).
"true"
6> erlang:list_to_binary("hi there").
<<"hi there">>
7> erlang:binary_to_list(<<"hi there">>).
"hi there
```

完整列表

atom_to_binary/2 integer_to_list/1 list_to_integer/2
atom_to_list/1 integer_to_list/2 list_to_pid/1
binary_to_atom/2 iolist_to_atom/1 list_to_tuple/1
binary_to_existing_atom/2 iolist_to_binary/1 pid_to_list/1
binary_to_list/1 list_to_atom/1 port_to_list/1
binary_to_term/1 list_to_binary/1 ref_to_list/1
binary_to_term/2 list_to_bitstring/1 term_to_binary/1
bitstring_to_list/1 list_to_existing_atom/1 term_to_binary/2
float_to_list/1 list_to_float/1 tuple_to_list/1
fun_to_list/1

#### 类型检查

有一些专门负责检测数据类型的函数。它们接收一个参数，如果参数的数据类型正确，就返回 true，否则，就返回 false。它们是为数不多的、可以在卫表达式中使用的函数中的一部分，也称为类型检测 BIF。

is_atom/1 is_function/1 is_port/1
is_binary/1 is_function/2 is_record/2
is_bitstring/1 is_integer/1 is_record/3
is_boolean/1 is_list/1 is_reference/1
is_builtin/3 is_number/1 is_tuple/1
is_float/1 is_pid/1

> 在所有可以在卫表达式中使用的函数中，类型检测 BIF 几乎占了大半。其余的也都是
> BIF，只是不用于类型测试，包括 abs(Number)、 bit_size(Binary)、 byte_size
> (Binary)、 element(N, Tuple)、 float(Term)、 hd(List)、 length(List)、
> node() 、 node(Pid|Ref|Port) 、 round(Number) 、 self() 、 tl(List) 、
> trunc(Number)和tuple_size(Tuple)。其中，函数node/1 和self/0 与分布式Erlang
> 以及进程/actor 有关

为什么没有提供这样一个函数，它可以直接返回传给它的数据项的类型（就
像这样： type_of(X) -> Type）。原因很简单： Erlang 只针对正确的情况编程。你的程序只需要处理那些你知道肯定会发生以及所期望的情况，对于除此以外的其他情况，都应该尽快抛出异常。因此，如果提供一个 type_of(X)函数，就会怂恿人们在代码中写出条件分支，有点像这样：

```erlang
my_function(Exp) ->
case type_of(Exp) of
binary -> Expression1;
list -> Expression2
end.

符合 Erlang 语言内在的声明性
my_function(Exp) when is_binary(Exp) -> Expression1;
my_function(Exp) when is_list(Exp) -> Expression2.
```



## *尾递归（函数式编程之没有循环）

```erlang
fac(0) -> 1;
fac(N) when N > 0 -> N*fac(N-1).

%% 求数组长度
len([]) -> 0;
len([_|T]) -> 1 + len(T).
```



```erlang
tail_fac(N) -> tail_fac(N,1).

tail_fac(0,Acc) -> Acc;
tail_fac(N,Acc) when N > 0 -> tail_fac(N-1,N*Acc).
```

tail_fac(4) = tail_fac(4,1)
tail_fac(4,1) = tail_fac(4-1, 4*1)
tail_fac(3,4) = tail_fac(3-1, 3*4)
tail_fac(2,12) = tail_fac(2-1, 2*12)
tail_fac(1,24) = tail_fac(1-1, 1*24)
tail_fac(0,24) = 24
看到差别了吗？这个版本中，程序占用的内存不会超过两个数据项，所以对空间的使用是常量的。计算 4 的阶乘和计算 1 000 000 的阶乘所耗费的空间是一样的（如果不考虑 4!在实际表示上比 1 000 000!小很多的话）。有了将阶乘函数改成尾递归的经验，你也许能够看出如何把这个模式应用在 len/1 函数上。我们要把递归调用变成独立的。如果你喜欢形象一点的说法，可以这样想象：给函数增加一个参数，把那个+1 的部分放到函数调用中去。

```erlang
tail_len(L) -> tail_len(L,0).
tail_len([], Acc) -> Acc;
tail_len([_|T], Acc) -> tail_len(T,Acc+1).
```



多几个练习

```erlang
>> 12 duplicate(0,_) ->
   13     [];
   14 duplicate(N,Term) when N > 0 ->
   15     [Term|duplicate(N-1,Term)].
   16 
   17 
>> 18 tail_duplicate(N,Term) ->
   19     tail_duplicate(N,Term,[]).
   20 
>> 21 tail_duplicate(0,_,List) ->
   22     List;
   23 tail_duplicate(N,Term,List) when N > 0 -> 
   24     tail_duplicate(N-1, Term, [Term|List]).
```

#### reverse

```erlang
reverse([]) -> [];
reverse([H|T]) -> reverse(T)++[H].

tail_reverse(L) -> tail_reverse(L,[]).
tail_reverse([],Acc) -> Acc;
tail_reverse([H|T],Acc) -> tail_reverse(T, [H|Acc]).
```



#### sublist

```erlang
sublist(_,0) -> [];
sublist([],_) -> [];
sublist([H|T],N) when N > 0 -> [H|sublist(T,N-1)].


tail_sublist(L, N) -> tail_sublist(L, N, []).
tail_sublist(_, 0, SubList) -> SubList;
tail_sublist([], _, SubList) -> SubList;
tail_sublist([H|T], N, SubList) when N > 0 ->
	reverse(tail_sublist(L, N, [])).
```

#### zip 

```erlang
zip([],[]) -> [];
zip([X|Xs],[Y|Ys]) -> [{X,Y}|zip(Xs,Ys)].

lenient_zip([],_) -> [];
lenient_zip(_,[]) -> [];
lenient_zip([X|Xs],[Y|Ys]) -> [{X,Y}|lenient_zip(Xs,Ys)].


zip 当碰到空时就停止
>> 26 zip([],_) -> [];
   27 zip(_,[]) -> [];
   28 zip([X|Xs],[Y|Ys]) -> [{X,Y}|zip(Xs,Ys)].
   29 
>> 30 tail_zip(Lone,Ltwo) -> tail_zip(Lone,Ltwo,[]).
   31 
>>k32 tail_zip([],_,Acc) -> Acc;
   33 tail_zip(_,[],Acc) -> Acc;
   34 tail_zip([X|Xs],[Y|Ys],Acc) -> 
   35     tail_zip(Xs,Ys,[{X,Y}|Acc]).

```

> 尾递归优化
> 这里编写的尾递归不会引起内存消耗的增长，因为当 VM 看到一个函数在尾部（函数中最
> 后一个被求值的表达式）调用了自身，它会清除当前的栈帧。这称为尾递归优化（ tail recursion
> optimization， TRO），它是一个更为一般的优化方法的特例，这种更为一般的优化方法称为尾
> 调用优化（ last call optimization， LCO）。
> 当函数体中的最后一个被求值的表达式是另外一个函数调用时，就会迚行 LCO。在迚行
> LCO 时，和 TRO 一样， Erlang VM 会避免存储栈帧。因此，尾递归也适用于多个函数的情况。
> 例如，函数调用链 a() -> b(). b() -> c(). c() -> a().实际上导致了一个无限循环，
> 但不会耗尽内存，因为 LCO 避免了栈溢出。这条原则，再加上累加器的使用，成就了尾递归
> 的可用性

#### 快速排序

```erlang
quicksort([]) -> [];
quicksort([Pivot|Rest]) -> 
    {Smaller, Larger} = partition(Pivot,Rest,[],[]),
    quickosrt(Smaller) ++ [Pivot] ++ quicksort(Larger).

partition(_,[],Smaller,Larger) -> {Smaller,Larger};
partition(Pivot, [H|T],Smaller, Larger) ->
    if H =< Pivot -> partition(Pivot,T,[H|Smaller],Larger);
        H > Pivot -> partition(Pivot, T, Smaller, [H|Larger])
    end.


lc_quicksort([]) -> [];
lc_quicksort([Pivot|Rest]) ->
lc_quicksort([Smaller || Smaller <- Rest, Smaller =< Pivot])
++ [Pivot] ++
lc_quicksort([Larger || Larger <- Rest, Larger > Pivot])
```



#### 实现一颗树

```erlang

```



## 高阶函数

> 一切都是函数

```erlang
map(_, []) -> [];
map(F, [H|T]) -> [F(H)|map(F,T)].
incr(X) -> X + 1.
decr(X) -> X - 1.

5> hhfuns:map(fun hhfuns:incr/1, L).
[2,3,4,5,6]
6> hhfuns:map(fun hhfuns:decr/1, L).
[0,1,2,3,4]

```



#### 匿名函数

```erlang
fun(Args1) ->
	Expression1, Exp2,..., ExpN;
(Args2) ->
	Expression1, Exp2,..., ExpN;
(Args3) ->
	Expression1, Exp2,..., ExpN
end
    
    
    
7> Fn = fun() -> a end.
#Fun<erl_eval.20.67289768>
8> Fn().
a
9> hhfuns:map(fun(X) -> X + 1 end, L).
[2,3,4,5,6]
10> hhfuns:map(fun(X) -> X - 1 end, L).
[0,1,2,3,4]
```

#### 闭包

> 不管匿名函数在哪里，这个被继承的作用域会一直跟随着它，即使把这个
> 匿名函数传递给另外一个函数。
>
> 匿名函数会一直带着继承的作用域

```erlang
base(A) ->
B = A + 1,
F = fun() -> A * B end,
F().
```

#### 过滤器

```erlang
%% 只保留偶数
even(L) -> lists:reverse(even(L,[])).
even([], Acc) -> Acc;
even([H|T], Acc) when H rem 2 == 0 ->
even(T, [H|Acc]);
even([_|T], Acc) ->
even(T, Acc).
%% 只保留年龄大于 60 的男性
old_men(L) -> lists:reverse(old_men(L,[])).
old_men([], Acc) -> Acc;
old_men([Person = {male, Age}|People], Acc) when Age > 60 ->
old_men(People, [Person|Acc]);
old_men([_|People], Acc) ->
old_men(People, Acc).


%% 抽象出一个过滤器
filter(Pred, L) -> lists:reverse(filter(Pred, L,[])).
filter(_, [], Acc) -> Acc;
filter(Pred, [H|T], Acc) ->
case Pred(H) of
true -> filter(Pred, T, [H|Acc]);
false -> filter(Pred, T, Acc)
end.
```



#### 折叠一切

```erlang
%% 找出列表中的最大值
max([H|T]) -> max2(T, H).
max2([], Max) -> Max;
max2([H|T], Max) when H > Max -> max2(T, H);
max2([_|T], Max) -> max2(T, Max).
%% 找出列表中的最小值
min([H|T]) -> min2(T,H).
min2([], Min) -> Min;
min2([H|T], Min) when H < Min -> min2(T,H);
min2([_|T], Min) -> min2(T, Min).
%% 计算列表所有元素的和
sum(L) -> sum(L,0).
sum([], Sum) -> Sum;
sum([H|T], Sum) -> sum(T, H+Sum).


%%  抽象出一个 折叠器
fold(_, Start, []) -> Start;
fold(F, Start, [H|T]) -> fold(F, F(H,Start), T).
```



#### 标准库的实现

Erlang 标准库中提供了许多基于列表的抽象，映射、过滤器和折叠只是其中几个（参见lists:map/2、 lists:filter/2、 lists:foldl/3 和 lists:foldr/3）。此处还有 all/2
和 any/2，这两个函数都以一个谓词为参数，分别用来测试列表中是否所有元素都为 true 以及列表中是否至少有一个元素为 true。

其中还有一个 dropwhile/2 函数，这个函数会略去列表前面的所有元素直到某个不满足谓词的元素。与之相反的是 takewhile/2 函数，这个函数会保留列表前面的所有元素直到某个不满足谓词的元素。作为这两个函数的补充，还有一个 partition/2 函数，它接收一个列表，返回两个列表，其中一个列表中的元素都满足给定的谓词，另外一个则包含剩余不满足的元素。

> 请阅读 lists 标准库实现

#### file 模块

file 模块中也包含有一些常见的文件读写函数，如 file:open/2 和 file:close/1
—它们的功能如其名字所示（打开和关闭文件！）： file:read/2—读取文件的内容（可以
作为字符串，也可以作为二进制）， file:read_line/1—读取一行以及 file: position/3
—把打开文件的位置指针移动到一个指定位置。
这个模块中还包含一组快捷函数，如 file:read_file/1（打开并以二进制方式读取整个
文件内容）、 file:consult/1（打开并把文件解析成 Erlang 数据项）、 file:pread/2（改变
位置指针，然后读取内容）以及 file:pwrite/2（改变位置指针并写入内容）。



## 错误 let it crash

1. 编译错误
2. 逻辑错误（请写 测试）
3. 运行时异常
   - 出错error
   - 退出 exit
   - 抛出 throw



#### 1. 出错异常

> erlang:error(Reason)

#### 2. 退出异常

内部退出 exit/1 

外部退出 exit/2

 erlang:error/1 会返回调用栈而 exit/1 不会。如果当前函数的调用栈很深，参数很多，把这个退出消息复制给每一个监听进程意味着要复制很多数据。在有些情况下，这并不现实

#### 3. 抛出异常

> throw(permission_denied).
>
> ** exception throw: permission_denied



处理异常

```erlang
try Expression of
	SuccessfulPattern1 [Guards] ->
		Expression1;
	SuccessfulPattern2 [Guards] ->
		Expression2
catch
	TypeOfError:ExceptionPattern1 ->
		Expression3;
	TypeOfError:ExceptionPattern2 ->
		Expression4
after
	Expr
end.
```

of 与 After 部分可以不要



例子

```erlang
-module(exceptions).
-compile(export_all).
throws(F) ->
	try F() of
		_ -> ok
	catch
		Throw -> {throw, caught, Throw}
	end.

errors(F) ->
	try F() of
		_ -> ok
	catch
		error:Error -> {error, caught, Error}
	end.
exits(F) ->
	try F() of
		_ -> ok
	catch
		exit:Exit -> {exit, caught, Exit}
	end.
```

> 不要在 尾递归中 使用异常
>
> 异常中被保护的部分是无法做到尾递归的。在执行时， VM 必须要一直保持一个相关的引用，以防异常出现。因为不带有 of 部分的 try ... catch 语句中只有被保护的部分，所以对于需要长时间运行的程序（ Erlang 的擅长领域）来说，在其中进行递归调用非常危险。经过充分数量的迭代乊后，会导致内存耗尽，或者让程序的运行速度变慢。如果把递归调用放在 of和 catch 乊间，就不会受到保护，此时可以对其进行尾调用优化（在第 5 章中介绍过）。不过，如果在 try 表达式中使用了 after 语句，那么就不能进行这种优化了，因为 after 语句中的表达式会被最后执行，所以需要在函数调用列表中对其进行记录。



#### 更多选择

1> catch throw(whoa).
whoa
2> catch exit(die).
{'EXIT',die}
3> catch 1/0.
{{'EXIT',{badarith,[{erlang,'/',[1,0],[]},
{erl_eval,do_apply,6,[{file,"erl_eval.erl"},{line,576}]},
{erl_eval,expr,5,[{file,"erl_eval.erl"},{line,360}]},
{shell,exprs,7,[{file,"shell.erl"},{line,668}]},
{shell,eval_exprs,7,[{file,"shell.erl"},{line,623}]},
{shell,eval_loop,3,[{file,"shell.erl"},{line,608}]}]}}
4> catch 2+2.
4
可以看出，抛出异常和之前一样，但是退出异常和出错异常都被表示为{'EXIT', Reason}。
这是由于出错是在退出之后增加到语言中的（Erlang 的实现者为了向后兼容使用了类似的表示方法）。
我们来尝试另外一个例子：
5> catch doesnt:exist(a,4).
{'EXIT',{undef,[{doesnt,exist,[a,4],[]},
{erl_eval,do_apply,6,[{file,"erl_eval.erl"},{line,576}]},
{erl_eval,expr,5,[{file,"erl_eval.erl"},{line,360}]},
{shell,exprs,7,[{file,"shell.erl"},{line,668}]},
{shell,eval_exprs,7,[{file,"shell.erl"},{line,623}]},
{shell,eval_loop,3,[{file,"shell.erl"},{line,608}]}]}}
错误的类型是 undef，意思是调用的函数没有定义。
紧跟着错误类型后面的列表是调用栈跟踪。下面是对调用栈跟踪的解释。
 调用栈跟踪的顶层元素表示最后一个调用函数（{Module, Function, Arguments}）。
也就是那个没有定义的函数。
 之后的元组是在错误出现之前所调用的函数。此时，它们的格式是{Module, Function,
Arity, Details}。
 Details 字段是一个元组列表，其中包含着文件名和文件行号。在本例中，文件名是
erl_eval.erl 和 shell.erl，因为这两个文件负责解释在 Erlang shell 中输入的代码。
关于调用栈跟踪的解释就这些了，真的！



## 数据结构

### 1. 记录record

>  Erlang 记录只是元组之上的语法糖。

首先，记录（record）是一种拼凑物。它们是在语言实现完毕后临时添加
上去的，因此使用起来有些不方便。但是，如果数据结构比较小，并且想直接
通过名字去访问属性字段，那么使用记录还是很合适的。记录的这种使用方式
和 C 语言中的结构很像。

定义一个record 

```erlang
-module(records).
-compile(export_all).
-record(robot, {name,
	type=industrial,
	hobbies,
	details=[]}).

first_robot() ->
    #robot{name="Mechatron",
        type=handmade,
        details=["Moved by a small man inside"]}.

执行一下
1> c(records).
{ok,records}
2> records:first_robot().
{robot,"Mechatron",handmade,undefined,
["Moved by a small man inside"]}
```

Erlang shell 提供了一条命令 rr(Module)，可以加载 Module 中定义的记录。

除了 rr()外， Erlang 还提供了其他一些函数，可以在 shell 中处理记录。

- rd(Name, Definition)， 定义一个记录，方式和模块中使用的-record(Name,
  Definition)类似。
- rf()，“卸载”所有记录。
- rf( Name)或者 rf([Names])，删除指定的记录定义。
- rl()，把 shell 中当前定义的所有记录打印出来，打印采用的记录格式非常易于复制、粘 贴 到 模 块 中 。 如 果 只 想 打 印 某 些 特 定 的 记 录 ， 可 以 rl(Name) 或 者rl([Names])。



从中读取值

```erlang
5> Crusher = #robot{name="Crusher", hobbies=["Crushing people","petting cats"]}.
#robot{name = "Crusher",type = industrial,
hobbies = ["Crushing people","petting cats"],
details = []}
6> Crusher#robot.hobbies.
["Crushing people","petting cats"]

7> NestedBot = #robot{details=#robot{name="erNest"}}.
#robot{name = undefined,type = industrial,
	hobbies = undefined,
	details = #robot{name = "erNest",type = industrial,
	hobbies = undefined,details = []}}
8> (NestedBot#robot.details)#robot.name.
"erNest"



9> #robot.type.
3
这条语句输出了 type 在底层元组中的元素位置编号。
```



可以在函数中使用模式匹配

```erlang
-record(user, {id, name, group, age}).
	%% 使用模式匹配进行过滤
admin_panel(#user{name=Name, group=admin}) ->
	Name ++ " is allowed!";
admin_panel(#user{name=Name}) ->
	Name ++ " is not allowed".
	%% 可以随意扩展 user 记录，函数无需修改
adult_section(U = #user{}) when U#user.age >= 18 ->
	%% 显示不适合写出来的内容
	allowed;
adult_section(_) ->
	%% 重定向到 Sesame Street 教育网站
	forbidden.

10> c(records).
{ok,records}
11> rr(records).
[robot,user]
12> records:admin_panel(#user{id=1, name="ferd", group=admin, age=96}).
"ferd is allowed!"
13> records:admin_panel(#user{id=2, name="you", group=users, age=66}).
"you is not allowed"
14> records:adult_section(#user{id=21, name="Bill", group=users, age=72}).
allowed
15> records:adult_section(#user{id=22, name="Noah", group=users, age=13}).
forbidden
```

更新记录

```erlang
repairman(Rob) ->
	Details = Rob#robot.details,
	NewRob = Rob#robot{details=["Repaired by 	repairman"|Details]},
	{repaired, NewRob}.

16> c(records).
{ok,records}
17> records:repairman(#robot{name="Ulbert", hobbies=["trying to have feelings"]}).
{repaired,#robot{name = "Ulbert",type = industrial,
	hobbies = ["trying to have feelings"],
	details = ["Repaired by repairman"]}}
```

#### 头文件的用法

records.hrl 

```erlang
%% 这是一个.hrl（头）文件
-record(included, {some_field,
	some_default = "yeah!",
	unimaginative_name}).
```

使用头文件

```erlang
-include("records.hrl").
```

### 2. 键值对存储

1. 属性列表

属性列表就是形如[{Key,Value}]的元组列表。

可以通过 proplists 模块来处理属性列表。这个模块中包含 proplists:delete/2、proplists: get_value/2、 proplists:get_all_values/2、 proplists:lookup/2和 proplists:lookup_all/2 之类的函数。

该模块中没有向列表中增加元素或者更新列表元素的函数。

2. 有序字典

如果在存储小量数据时想使用更完善些的键/值存储，可以考虑使用 orddict 模块。 有序字典是具有一些规范要求的属性列表。每个键都只能出现一次。这个列表是排过序的，因此平均来说，查找速度要快一些。所存储的数据项必须是严格的{Key, Value}形式。不能像属性列表那样，把有序字典当成列表来操作，必须使用 orddict 模块的函数接口来进行所有需要的操作。

常用的 CRUD（创建、读取、更新和删除）函数包括 orddict:store/3，orddict:find/2（当不知道字典中是否包含给定主键时使用）、 orddict:fetch/2（当知道主键存在或者必须存在 时 使 用 ） 以 及 orddict:erase/2 。 可 以 使 用 orddict:new/0 或 者 orddict:from_list/1 创建一个有序字典。

> 一般来讲，对于小于 75 个元素的数据量来说，使用这个够了，但是如果元素更多，请使用后面的。

### 3. 大数据量存储：字典和通用平衡树

Erlang 提供了两种用于大数据量存储的键/值数据结构： 字典（dict）和通用平衡（GB） 树。 字典的接口和有序字典完全一样： dict:store/3、 dict:find/2 以及 dict:fetch/ 2、dict:erase/2。 字典中也具有有序字典中的所有其他函数，如 dict:map/2 和 dict:fold/2（在处理整个数据结构时很有用！）。因此，当有序字典需要向上伸展时， 字典是非常好的选择。

GB 树是由 gb_trees 模块实现的，其中包含的函数要比 dict 中多，在数据结构的使用上提供了更多的直接控制手段。 gb_trees 有两种主要的工作模式：一种针对彻底了解自己数据的情况（称之为智能模式），还有一种针对不能对数据做假设的情况（称之为简单模式）。在简单模式中，数据操作函数为 gb_trees:enter/2、 gb_trees:lookup/2 和 gb_trees:delete_any/2 。 在 智 能 模 式 中 ， 相 应 的 函 数 为 gb_trees:insert/3 、 gb_trees:get/2 、gb_trees:update/3 和 gb_trees:delete/2。还有一个 gb_trees:map/2 函数，它和lists:map/2 等价，不过操作的对象是树（在需要时，它非常好用）。

简单模式函数相比智能模式函数存在劣势的原因在于 GB 树是平衡树，每当插入一个新元素时（或者删除一批元素时），树都需要平衡自己。这需要时间和内存（即使是些无用的检查，最后发现无需变化，也要花时间去确定树已经处于平衡状态）。智能模式函数默认已知晓键值在树中的存在情况。有了这种假设，就可以略过所有的安全检查，从而达到更快的执行速度。

### 4. 集合

1. ordsets
   ordsets 模块集合被实现为一种有序列表。它们主要适用于小集合，是最慢的一种集
   合，不过它的实现是所有集合中最简单、最容易理解的一种。该模块中的一些标准函数有
   ordsets:new/0、 ordsets:is_element/2、 ordsets:add_element/2、 ordsets:
   del_element/2、 ordsets:union/1 和 ordsets:intersection/1。
2. sets
   sets（模块名） 实现使用的底层数据结构和 dict 使用的相似。 sets 模块的接口和
   ordsets 完全一样，不过支持的数据规模要大一些。和 dict 一样， sets 更擅长读密集型
   的处理，如检查某个元素是否在集合中。
3. gb_sets
   gb_sets 模块的底层实现结构是一棵 GB 树，和 gb_trees 模块使用的类似。gb_sets
   和 sets 的关系与 gb_trees 和 dict 的关系一样：在非读取操作方面， gb_sets 要更快
   一些，提供的控制手段也更多一些。 gb_sets 在实现了和 sets、 ordsets 同样接口的同
   时，还增加了其他一些函数。和 gb_tress 一样， gb_sets 中也分智能模式函数和简单模
   式函数，也有迭代器以及对最小值和最大值的快速访问函数。
4. sofs
   可以使用 sofs 模块创建集合的集合。这个模块使用有序列表实现，这个列表被放置在
   一个包含元数据信息的元组中。如果想完全掌控集合和集合族之间的关系，强制集合类型或
   者有其他类似要求时，可以使用这个模块。当想使用数学意义上的集合概念，而又不仅仅是
   值唯一的元素组时，这个模块非常有用

> 建议在大多数情况下尽量使用 gb_sets，当需要一种清晰的表示，想在自己
> 的代码中操纵这种表示时，使用 ordsets，当需要=:=操作符时，使用 sets

> gb_sets、 ordsets 和 sofs 都使用==操作符来迚行值的比较，如果有两个数 2
> 和 2.0，它们会被认为是相同的数。但是，在 sets 模块中使用了=:=操作符

### 5. 有向图

有向图在 Erlang 中被实现为两个模块： digraph 和 digraph_utils。 digraph 模块主要实现了有向图的构造和修改功能—操作边和顶点、寻找路径和环等。 digraph_utils 模块则实现了图的遍历功能（后序和前序），环、 树形图①以及树性质检测，寻找邻居顶点等功能。

### 6. 队列

queue模块

原始 API
原始 API 中包含了队列实现的基础函数，包括创建空队列的 new/0，插入新元素的in/2，以及移除元素的 out/1。其中也包含有像把队列转换成列表、反转队列、检查某个特定值是否在队列中之类的函数。

扩展 API
扩展 API 主要增加了一些内省能力和灵活性。可以用它来进行一些诸如在不移除第一个元素的情况下查看队列的头元素（ get/1 或者 peek/1）、直接移除元素而不关心它的值（drop/1）之类的操作。虽然这些函数并不是队列必要的操作，但是总体来讲它们还是很有用的。



## 函数式的例子

### 逆波兰式计算器

### 从希斯罗到伦敦



执行

$ erlc road.erl
$ erl -noshell -run road main road.txt

使用 escript
Erlang 的 escript 命令提供了一种在不直接启动 erl 的情况下运行 Erlang 程序的简单方
法。简单来讲， escript 命令以一个模块为参数，然后无需编译即可解释执行这个模块。
模块的结构和前面介绍的类似，不过需要更改一下模块的头部说明。不能再使用
-module(Name)属性，而是要使用如下格式：
#!/usr/bin/env escript
%% -*- erlang -*-
%%! -pa 'ebin/' [Other erl Arguments]
main([StringArguments]) ->
...
在脚本启动时，会自动调用 main/1 函数，脚本的启动命令可以是./script-name.erl
或者 escript script-name.erl（后者更适用于 Windows 平台）。模块会像一个普通脚本
一样运行。
如果你想得到 escript 的便利性，但是不想解释执行代码（会慢一些），更希望能够执行
编译后的结果，那么只需在文件中的某个地方增加-mode(compile).模块属性即可。
关于 escript 的更多信息，请参阅 Erlang 自带的相关文档，也可以在线阅读：
http://erlang.org/doc/man/escript.html。













