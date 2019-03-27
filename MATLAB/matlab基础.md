# MATLAB

## 常用命令

1. clear
2. who
3. whos
4. clc
5. type 显示文件内容



1. 常用的矩阵操作：

rem 用来取模

logic  === bool nei 类型

length 返回向量的 长度

size 返回规模

numel 返回个数

reshape

fliplr  上下反转

flipud   左右反转

rot90  可以接参数旋转任意度数

repmat  用来重复矩阵多少次



2. 多维矩阵：

mat = 3 x 3 

mat(:,:,1)  == mat

mat(:,:,2) =  3 x 3 

这样就变成了三维矩阵



4. 类型：

int8(200) == 127 matlab不会截断，会饱和运算。

char('abcd' + 1)



5. 输入输出

disp('')

fprintf(' %d \n', 4^3)  %f %c %s



6. 保存文件，加载文件

   save   -append  -ascii  如果有-ascii就不会 带.dat

   load -ascii

带不带 .dat 区别在于时候是存在一个变量 ，或者多个变量



7. 逻辑运算符

||

&&

~

xor()

同样  | & 分别用于矩阵的 逐个元素 && ||



sum(xx,1)

sum(xx,2)

sum(xx,'all')

prod([x x x])  计算所有元素乘积

switch  xxxx

​	case {xx,xx,xx} 

​		grade = 'a'

​	case 8

​		grade = 'B'

​	otherwise

end



for i = 1:3 

end



8. 一个最简单的gui

menu(title, xxx,xxx,xxx)

9. is函数 用来判断 是不是一种类型



cumsum  cumprod  累加和，累加积

diff 函数 返回逐个元素的 差



tic  / toc  可以用来确定程序的执行时间



10. 字符串变量

strcat  删除前空格， 不删除后空格

blanks(4)

delblank()

strtrim()  trim操作

upper

lower

strcmp  比较字符串

findstr  在长字符串中找出 短字符串的 索引

strrep  替换 有三个参数

strtok  分割

eval('plot(x)')  和python中的 eval一样



其他高级操作

1.  vec >  5 返回一个 logic组成的向量
2. 逻辑索引   mat(vec)     % vec 是一个逻辑索引
3. logic 强制转换  logic([1 1 1 0 1 0 ])  但是 true false 函数更加高效
4. any  all  类似 python里面的 简写 &&  || 
5. find(vec > 5)    返回符合条件的索引 
6. vec1 == vec2  返回一个 logic 矩阵 ，然后使用 all就可以比较了 。不过内置的isequal 实现了这个功能



1. 元胞数组矩阵 ，一个可以存储不同类型数据的矩阵

{} 与 （） 的区别  前者 得到的还是 元胞数组，后者得到的 它本身类型





## 官网基础教程

> 一切皆是数组

如果未指定输出变量，MATLAB 将使用变量 `ans`（*answer* 的缩略形式）来存储计算结果。

```
sin(a)
ans =

    0.8415
```

如果语句以分号结束，MATLAB 会执行计算，但不在命令行窗口中显示输出。

```
e = a*b;
```

按向上 (↑) 和向下箭头键 (↓) 可以重新调用以前的命令。在空白命令行中或在键入命令的前几个字符之后按箭头键。例如，要重新调用命令 `b = 2`，请键入 `b`，然后按向上箭头键。

### 1. 矩阵数组

#### 创建数组

要创建每行包含四个元素的数组，请使用逗号 (`,`) 或空格分隔各元素。

```
a = [1 2 3 4]
a = 1×4

     1     2     3     4
```

这种数组为*行向量*。

要创建包含多行的矩阵，请使用分号分隔各行。

```
a = [1 2 3; 4 5 6; 7 8 10]
a = 3×3

     1     2     3
     4     5     6
     7     8    10
```

创建矩阵的另一种方法是使用 `ones`、`zeros` 或 `rand` 等函数。例如，创建一个由零组成的 5×1 列向量。

```
z = zeros(5,1)
z = 5×1

     0
     0
     0
     0
     0
```

#### 简单的运算

MATLAB 允许您使用单一的算术运算符或函数来处理矩阵中的所有值。

```
a + 10
ans = 3×3

    11    12    13
    14    15    16
    17    18    20
sin(a)
ans = 3×3

    0.8415    0.9093    0.1411
   -0.7568   -0.9589   -0.2794
    0.6570    0.9894   -0.5440
```

要转置矩阵，请使用单引号 (`'`)：

```
a'
ans = 3×3

     1     4     7
     2     5     8
     3     6    10
```

您可以使用 `*` 运算符执行标准矩阵乘法，这将计算行与列之间的内积。例如，确认矩阵乘以其逆矩阵可返回单位矩阵：

```
p = a*inv(a)
p = 3×3

    1.0000         0   -0.0000
         0    1.0000         0
         0         0    1.0000
```

请注意，`p` 不是整数值矩阵。MATLAB 将数字存储为浮点值，算术运算可以区分实际值与其浮点表示之间的细微差别。使用 `format` 命令可以显示更多小数位数：

```
format long
p = a*inv(a)
p = 3×3

   1.000000000000000                   0  -0.000000000000000
                   0   1.000000000000000                   0
                   0                   0   0.999999999999998
```

使用以下命令将显示内容重置为更短格式

```
format short
```

`format` 仅影响数字显示，而不影响 MATLAB 对数字的计算或保存方式。



要执行元素级乘法（就是对矩阵中的每个元素做运算），请使用 `.*` 运算符：

```
p = a.*a
p = 3×3

     1     4     9
    16    25    36
    49    64   100
```

乘法、除法和幂的矩阵运算符分别具有执行元素级运算的对应数组运算符。例如，计算 `a` 的各个元素的三次方：

```
a.^3
ans = 3×3

           1           8          27
          64         125         216
         343         512        1000
```

#### 串联

*串联*是连接数组以便形成更大数组的过程。实际上，第一个数组是通过将其各个元素串联起来而构成的。成对的方括号 `[]` 即为串联运算符。

```
A = [a,a]
A = 3×6

     1     2     3     1     2     3
     4     5     6     4     5     6
     7     8    10     7     8    10
```

使用逗号将彼此相邻的数组串联起来称为*水平*串联。每个数组必须具有相同的行数。同样，如果各数组具有相同的列数，则可以使用分号*垂直*串联。

```matlab
A = [a; a]
A = 6×3

     1     2     3
     4     5     6
     7     8    10
     1     2     3
     4     5     6
     7     8    10
```

#### 复数

复数包含实部和虚部，虚数单位是 `-1` 的平方根。

```
sqrt(-1)
ans = 0.0000 + 1.0000i
```

要表示复数的虚部，请使用 `i` 或 `j`。

```
c = [3+4i, 4+3j; -i, 10j]
c = 2×2 complex

   3.0000 + 4.0000i   4.0000 + 3.0000i
   0.0000 - 1.0000i   0.0000 +10.0000i
```

#### 数组索引(不同于python，golang索引)

MATLAB® 中的每个变量都是一个可包含许多数字的数组。如果要访问数组的选定元素，请使用索引。

以 4×4 幻方矩阵 `A` 为例：

```
A = magic(4)
A =

    16     2     3    13
     5    11    10     8
     9     7     6    12
     4    14    15     1
```

引用数组中的特定元素有两种方法。最常见的方法是指定行和列下标，例如

```
A(4,2)
ans =

    14
```

另一种方法不太常用，但有时非常有用，即使用单一下标按顺序向下遍历每一列：

```
A(8)
ans =

    14
```

使用单一下标引用数组中特定元素的方法称为*线性索引*。



您可以在赋值语句左侧指定当前维外部的元素。数组大小会增大以便容纳新元素。

```
A(4,5) = 17
A =

    16     2     3    13     0
     5    11    10     8     0
     9     7     6    12     0
     4    14    15     1    17
```

要引用多个数组元素，请使用冒号运算符，这使您可以指定一个格式为 `start:end` 的范围。例如，列出 `A` 前三行及第二列中的元素：

```
A(1:3,2)
ans =

     2
    11
     7
```

单独的冒号（没有起始值或结束值）指定该维中的所有元素。例如，选择 `A` 第三行中的所有列：

```
A(3,:)
ans =

     9     7     6    12     0
```

此外，冒号运算符还允许您使用较通用的格式 `start:step:end` 创建等距向量值。

```
B = 0:10:100
B =

     0    10    20    30    40    50    60    70    80    90   100
```

如果省略中间的步骤（如 `start:end` 中），MATLAB 会使用默认步长值 `1`

##### 下标

x(L) 一列一列的查找

[`find`](https://ww2.mathworks.cn/help/matlab/ref/find.html) 函数可用于确定与指定逻辑条件相符的数组元素的索引。`find` 以最简单的形式返回索引的列向量。转置该向量以便获取索引的行向量。例如，再次从丢勒的幻方矩阵开始。（请参阅 [magic 函数](https://ww2.mathworks.cn/help/matlab/learn_matlab/matrices-and-magic-squares.html#f2-484)。）

```
k = find(isprime(A))'
```

使用一维索引选取幻方矩阵中的质数的位置：

```
k =
     2     5     9    10    11    13
```

使用以下命令按 `k` 确定的顺序将这些质数显示为行向量

```
A(k)

ans =
     5     3     2    11     7    13
```

将 `k` 用作赋值语句的左侧索引时，会保留矩阵结构：

```
A(k) = NaN

A =
    16   NaN   NaN   NaN
   NaN    10   NaN     8
     9     6   NaN    12
     4    15    14     1
```

##### 冒号运算符

要获取非单位间距，请指定增量。例如，

```
100:-7:50
```

为

```
100    93    86    79    72    65    58    51
```

而

```
0:pi/4:pi
```

为

```
0    0.7854    1.5708    2.3562    3.1416
```

包含冒号的下标表达式引用部分矩阵：

```
A(1:k,j)
```

表示 `A` 第 `j` 列中的前 `k` 个元素。因此，

```
sum(A(1:4,4))
```

计算第四列的总和。但是，执行此计算有一种更好的方法。冒号本身引用矩阵行或列中的*所有*元素，而关键字 [`end`](https://ww2.mathworks.cn/help/matlab/ref/end.html) 引用*最后一个*行或列。因此，

```
sum(A(:,end))
```

计算 `A` 最后一列中的元素的总和：

```
ans =
     34
```

为什么 4×4 幻方矩阵的幻数和等于 34？如果将介于 1 到 16 之间的整数分为四个总和相等的组，该总和必须为

```
sum(1:16)/4
```

当然，也即

```
ans =
     34
```

##### 删除行和列

只需使用一对方括号即可从矩阵中删除行和列。首先

```
X = A;
```

然后，要删除 `X` 的第二列，请使用

```
X(:,2) = []
```

这会将 `X` 更改为

```
X =
    16     2    13
     5    11     8 
     9     7    12
     4    14     1
```

如果您删除矩阵中的单个元素，结果将不再是矩阵。因此，以下类似表达式

```
X(1,2) = []
```

将会导致错误。但是，使用单一下标可以删除一个元素或元素序列，并将其余元素重构为一个行向量。因此

```
X(2:2:10) = []
```

生成

```
X =
    16     9     2     7    13    12     1
```

### 2. 工作区变量的保存与应用

工作区变量

*工作区*包含在 MATLAB® 中创建或从数据文件或其他程序导入的变量。例如，下列语句在工作区中创建变量 `A` 和 `B`。

```
A = magic(4);
B = rand(3,5,2);
```

使用 `whos` 可以查看工作区的内容。

```
whos
  Name      Size             Bytes  Class     Attributes

  A         4x4                128  double              
  B         3x5x2              240  double              
```

此外，桌面上的“工作区”窗格也会显示变量。

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/workspace_zh_CN.png)

退出 MATLAB 后，工作区变量不会保留。使用 `save` 命令保存数据以供将来使用，

```
save myfile.mat
```

通过保存，系统会使用 `.mat` 扩展名将工作区保存在当前工作文件夹中一个名为 MAT 文件的压缩文件中。

要清除工作区中的所有变量，请使用 `clear` 命令。

使用 `load` 将 MAT 文件中的数据还原到工作区。

```
load myfile.mat
```



### 3. 文本字符串

当您处理文本时，将字符序列括在单引号中。可以将文本赋给变量。

```
myText = 'Hello, world';
```

如果文本包含单引号，请在定义中使用两个单引号。

```
otherText = 'You''re right'
otherText = 
'You're right'
```

与所有 MATLAB® 变量一样，`myText` 和 `otherText` 为数组。其*类*或数据类型为 `char`（*character* 的缩略形式）。

```
whos myText
  Name        Size            Bytes  Class    Attributes

  myText      1x12               24  char               
```

您可以使用方括号串联字符数组，就像串联数值数组一样。

```
longText = [myText,' - ',otherText]
longText = 
'Hello, world - You're right'
```

要将数值转换为字符，请使用 `num2str` 或 `int2str` 等函数。

```
f = 71;
c = (f-32)/1.8;
tempText = ['Temperature is ',num2str(c),'C']
tempText = 
'Temperature is 21.6667C'
```



### 矩阵操作

##### 求和

`sum`对列求和 

##### 转置

A'

##### 对角矩阵

>  sum(diag(A))



##### 交换矩阵的列行

> A = B(:,[1 3 2 4])

##### 生成矩阵的函数

1. zeros
2. ones
3. rand  均匀分布
4. randn  正态分布

> “等于”关系运算符 `==` 要求实部和虚部相等。其他二进制关系运算符 `>`、`<`、`>=` 和 `<=` 忽略数字的虚部，而仅考虑实部。

##### 矩阵运算符



表达式使用大家熟悉的算术运算符和优先法则。

| `+`   | 加法         |
| ----- | ------------ |
| -     | 减法         |
| `*`   | 乘法         |
| `/`   | 除法         |
| `\`   | 左除         |
| `^`   | 幂           |
| `'`   | 复共轭转置   |
| `( )` | 指定计算顺序 |

##### 数组运算符



如果矩阵不用于线性代数运算，则成为二维数值数组。数组的算术运算按元素执行。这意味着，加法和减法运算对数组和矩阵都是相同的，但乘法运算不相同。MATLAB 的乘法数组运算表示法中包含点，也就是小数点。



运算符列表包括

| `+`  | 加法           |
| ---- | -------------- |
| `-`  | 减法           |
| `.*` | 逐元素乘法     |
| `./` | 逐元素除法     |
| `.\` | 逐元素左除     |
| `.^` | 逐元素幂       |
| `.'` | 非共轭数组转置 |



如果使用数组乘法将丢勒的幻方矩阵自乘

```
A.*A
```

则会生成一个数组，该数组包含介于 1 至 16 之间的整数的平方，并且以不常见的顺序排列：

```
ans =
   256     9     4   169
    25   100   121    64
    81    36    49   144
    16   225   196     1
```



## 多维数组的玩法

语句

```
p = perms(1:4);
```

生成 4! = 24 置换`1:4`。第 `k` 个置换为行向量 `p(k,:)`。然后，

```
A = magic(4);
M = zeros(4,4,24);

for k = 1:24
   M(:,:,k) = A(:,p(k,:));
end
```

语句

```
sum(M,d)
```

通过改变第 `d` 个下标来计算总和。因此

```
sum(M,1)
```

是一个含有 24 个行向量副本的 1×4×24 数组

```
34    34    34    34
```

而

```
sum(M,2)
```

是一个含有 24 个列向量副本的 4×1×24 数组

```
34    
34    
34    
34
```

最后，

```
S = sum(M,3)
```

在序列中添加 24 个矩阵。结果的大小为 4×4×1，因此它看似是 4×4 数组：

```
S =
   204   204   204   204
   204   204   204   204
   204   204   204   204
   204   204   204   204
```

### 函数

如果存在多个输出参数，请将其括在方括号中：

```
[maxA,location] = max(A)
maxA = 5
location = 3

```

将任何字符输入括在单引号中：

```
disp('hello world')
hello world

```

要调用不需要任何输入且不会返回任何输出的函数，请只键入函数名称：

```
clc
```

`clc` 函数清除命令行窗口。

有关初等数学函数的列表，请键入

```
help elfun
```

有关更多高等数学函数和矩阵函数的列表，请键入

```
help specfun
help elmat
```

一些特殊函数提供了有用的常量值。

| [`pi`](https://ww2.mathworks.cn/help/matlab/ref/pi.html)     | 3.14159265...         |
| ------------------------------------------------------------ | --------------------- |
| [`i`](https://ww2.mathworks.cn/help/matlab/ref/i.html)       | 虚数单位 G−1          |
| [`j`](https://ww2.mathworks.cn/help/matlab/ref/j.html)       | 与 `i` 相同           |
| [`eps`](https://ww2.mathworks.cn/help/matlab/ref/eps.html)   | 浮点相对精度 ε=2−52   |
| [`realmin`](https://ww2.mathworks.cn/help/matlab/ref/realmin.html) | 最小浮点数 2−1022     |
| [`realmax`](https://ww2.mathworks.cn/help/matlab/ref/realmax.html) | 最大浮点数 (2−ε)21023 |
| [`Inf`](https://ww2.mathworks.cn/help/matlab/ref/inf.html)   | 无穷大                |
| [`NaN`](https://ww2.mathworks.cn/help/matlab/ref/nan.html)   | 非数字                |







#### 输出格式

```
x = [4/3 1.2345e-6]

format short

   1.3333    0.0000

format short e

   1.3333e+000  1.2345e-006

format short g

   1.3333  1.2345e-006

format long

   1.33333333333333   0.00000123450000

format long e

   1.333333333333333e+000    1.234500000000000e-006

format long g

   1.33333333333333               1.2345e-006

format bank

   1.33          0.00

format rat

   4/3          1/810045

format hex
 
   3ff5555555555555   3eb4b6231abfd271
```

如果矩阵的最大元素大于 103 或小于 10-3，MATLAB 会对短格式和长格式应用常用缩放因子。

除了上面显示的 `format` 函数，

```
format compact
```

会不显示在输出中出现的多个空行。这样，您可以在屏幕或窗口中查看更多信息。如果要进一步控制输出格式，请使用 [`sprintf`](https://ww2.mathworks.cn/help/matlab/ref/sprintf.html) 和 [`fprintf`](https://ww2.mathworks.cn/help/matlab/ref/fprintf.html) 函数。

##### 取消输出



如果您在仅键入语句后按 **Return** 或 **Enter**，MATLAB 会在屏幕上自动显示结果。但是，如果使用分号结束行，MATLAB 会执行计算，但不会显示任何输出。当生成大型矩阵时，此功能尤其有用。例如，

```
A = magic(100);
```

##### 输入长语句



如果语句无法容纳在一行中，请使用省略号（三个句点）`...`，后跟 **Return** 或 **Enter** 以指示该语句在下一行继续。例如，

```
s = 1 -1/2 + 1/3 -1/4 + 1/5 - 1/6 + 1/7 ...
      - 1/8 + 1/9 - 1/10 + 1/11 - 1/12;
```

`=`、`+` 和 - 符号周围的空白是可选的，但可提高可读性。











## 结构体

### 结构体





结构体是多维 MATLAB 数组，包含可按文本*字段标志符*访问的元素。例如，

```
S.name = 'Ed Plum';
S.score = 83;
S.grade = 'B+'
```

创建一个具有三个字段的标量结构体：

```
S = 
     name: 'Ed Plum'
    score: 83
    grade: 'B+'
```

与 MATLAB 环境中的所有其他内容一样，结构体也为数组，因此可以插入其他元素。在本示例中，数组的每个元素都是一个具有若干字段的结构体。可以一次添加一个字段，

```
S(2).name = 'Toni Miller';
S(2).score = 91;
S(2).grade = 'A-';
```

也可以使用一个语句添加整个元素：

```
S(3) = struct('name','Jerry Garcia',... 
               'score',70,'grade','C')
```

现在，结构体非常大以致仅输出摘要：

```
S = 
1x3 struct array with fields:
    name
    score
    grade
```

将不同字段重新组合为其他 MATLAB 数组的方法有许多种。这些方法大多基于*逗号分隔列表*的表示法。键入

```
S.score
```

与键入

```
S(1).score, S(2).score, S(3).score
```

相同，这是一个逗号分隔列表。

如果将生成此类列表的表达式括在方括号中，MATLAB 会将该列表中的每一项都存储在数组中。在本示例中，MATLAB 创建一个数值行向量，该向量包含结构体数组 `S` 的每个元素的 `score` 字段：

```
scores = [S.score]
scores =
    83    91    70

avg_score = sum(scores)/length(scores)
avg_score =
   81.3333
```

要根据某个文本字段（例如，`name`）创建字符数组，请对 `S.name` 生成的逗号分隔列表调用 [`char`](https://ww2.mathworks.cn/help/matlab/ref/char.html) 函数：

```
names = char(S.name)
names =
   Ed Plum    
   Toni Miller
   Jerry Garcia
```

同样，通过将生成列表的表达式括入花括号中，可以根据 `name` 字段创建元胞数组：

```
names = {S.name}
names = 
    'Ed Plum'    'Toni Miller'    'Jerry Garcia'
```

要将结构体数组的每个元素的字段赋值给结构体外部的单独变量，请指定等号左侧的每个输出，并将其全部括在方括号中：

```
[N1 N2 N3] = S.name
N1 =
   Ed Plum
N2 =
   Toni Miller
N3 =
   Jerry Garcia
```

#### 动态字段名称

访问结构体中的数据的最常用方法是指定要引用的字段的名称。访问结构体数据的另一种方法是使用动态字段名称。这些名称将字段表示为变量表达式，MATLAB 会在运行时计算这些表达式。此处显示的点-括号语法将 `expression` 作为动态字段名称：

```
structName.(expression)
```

使用标准 MATLAB 索引语法创建此字段的索引。例如，要在字段名称中计算 `expression`，并在行 `7` 中的 `1` 至 `25` 列内获取该字段的值，请使用

```
structName.(expression)(7,1:25)
```

**动态字段名称示例-**  下面显示的 `avgscore` 函数可用于计算考试的平均分数，并使用动态字段名称检索 `testscores` 结构体中的信息：

```
function avg = avgscore(testscores, student, first, last)
for k = first:last
   scores(k) = testscores.(student).week(k);
end
avg = sum(scores)/(last - first + 1);
```

您可以运行此函数，并对动态字段 `student` 使用不同值。首先，对包含 25 周内的分数的结构体进行初始化：

```
testscores.Ann_Lane.week(1:25) = ...
  [95 89 76 82 79 92 94 92 89 81 75 93 ...
   85 84 83 86 85 90 82 82 84 79 96 88 98];

testscores.William_King.week(1:25) = ...
  [87 80 91 84 99 87 93 87 97 87 82 89 ...
   86 82 90 98 75 79 92 84 90 93 84 78 81];
```

现在，运行 `avgscore`，并在运行时使用动态字段名称为 `testscores` 结构体提供学生姓名字段：

```
avgscore(testscores, 'Ann_Lane', 7, 22)
ans = 
   85.2500

avgscore(testscores, 'William_King', 7, 22)
ans = 
   87.7500
```









### 5. 画图

#### 线图

要创建二维线图，请使用 `plot` 函数。例如，绘制从 0 到 ![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_eq01_zh_CN.png) 之间的正弦函数值：

```
x = 0:pi/100:2*pi;
y = sin(x);
plot(x,y)

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_01_zh_CN.png)

可以标记轴并添加标题。

```
xlabel('x')
ylabel('sin(x)')
title('Plot of the Sine Function')

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_02_zh_CN.png)

通过向 `plot` 函数添加第三个输入参数，您可以使用红色虚线绘制相同的变量。

```
plot(x,y,'r--')

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_03_zh_CN.png)

`'r--'` 为*线条设定*。每个设定可包含表示线条颜色、样式和标记的字符。标记是在绘制的每个数据点上显示的符号，例如，`+`、`o` 或 `*`。例如，`'g:*'` 请求绘制使用 `*` 标记的绿色点线。

请注意，为第一幅绘图定义的标题和标签不再被用于当前的*图窗*窗口中。默认情况下，每次调用绘图函数、重置坐标区及其他元素以准备新绘图时，MATLAB® 都会清除图窗。

要将绘图添加到现有图窗中，请使用 `hold on`。在使用 `hold off` 或关闭窗口之前，当前图窗窗口中会显示所有绘图。

```
x = 0:pi/100:2*pi;
y = sin(x);
plot(x,y)

hold on

y2 = cos(x);
plot(x,y2,':')
legend('sin','cos')

hold off

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_04_zh_CN.png)

#### 三维绘图

三维图通常显示一个由带两个变量的函数（即 *z = f (x,y*)）定义的曲面图。

要计算 *z*，请首先使用 `meshgrid` 在此函数的域中创建一组 (*x,y*) 点。

```
[X,Y] = meshgrid(-2:.2:2);                                
Z = X .* exp(-X.^2 - Y.^2);

```

然后，创建曲面图。

```
surf(X,Y,Z)

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_05_zh_CN.png)

`surf` 函数及其伴随函数 `mesh` 以三维形式显示曲面图。`surf` 使用颜色显示曲面图的连接线和面。`mesh` 生成仅以颜色标记连接定义点的线条的线框曲面图。

#### 子图

使用 `subplot` 函数可以在同一窗口的不同子区域显示多个绘图。

`subplot` 的前两个输入表示每行和每列中的绘图数。第三个输入指定绘图是否处于活动状态。例如，在图窗窗口的 2×2 网格中创建四个绘图。

```
t = 0:pi/10:2*pi;
[X,Y,Z] = cylinder(4*cos(t));
subplot(2,2,1); mesh(X); title('X');
subplot(2,2,2); mesh(Y); title('Y');
subplot(2,2,3); mesh(Z); title('Z');
subplot(2,2,4); mesh(X,Y,Z); title('X,Y,Z');

```

![img](https://ww2.mathworks.cn/help/matlab/learn_matlab/gs2dand3dplotsexample_06_zh_CN.png)

### 6. matlab脚本编程

> edit  txt.txt

#### 循环及条件语句

在脚本中，可以使用关键字 `for`、`while`、`if` 和 `switch` 循环并有条件地执行代码段。

例如，创建一个名为 `calcmean.m` 的脚本，该脚本使用 `for` 循环来计算 5 个随机样本的均值和总均值。

```matlab
nsamples = 5;
npoints = 50;

for k = 1:nsamples
   iterationString = ['Iteration #',int2str(k)];
   disp(iterationString)
   currentData = rand(npoints,1);
   sampleMean(k) = mean(currentData)
end
overallMean = mean(sampleMean)

if overallMean < .49
   disp('Mean is less than expected')
elseif overallMean > .51
   disp('Mean is greater than expected')
else
   disp('Mean is within the expected range')
end
```

#### 帮助函数文档

- 使用 `doc` 命令在单独的窗口中打开函数文档。

  ```
  doc mean
  
  ```

- 在键入函数输入参数的左括号之后暂停，此时命令行窗口中会显示相应函数的提示（函数文档的语法部分）。

  ```
  mean
  ```

- 使用 `help` 命令可在命令行窗口中查看相应函数的简明文档。

  ```
  help mean
  
  ```































