# latex入门文档

## 入门

TeX 的源代码是后缀为 `.tex` 的纯文本文件。使用任意纯文本编辑器，都可以修改 `.tex` 文件：包括 Windows 自带的记事本程序，也包括专为 TeX 设计的编辑器（TeXworks, TeXmaker, TeXstudio, WinEdt 等），还包括一些通用的文本编辑器（Sublime Text, Atom, Visual Studio Code 等）。你可以在这些能够编辑纯文本文件的编辑器中任选其一作为你的 TeX 编辑器，也可以使用 TeX 发行自带的编辑器。最流行的两个 TeX 发行（TeX Live 和 MiKTeX）都带有 TeXworks 编辑器。



```latex
\documentclass[]{article}

%opening
\title{}
\author{}

\begin{document}

\maketitle

\begin{abstract}

\end{abstract}

\section{}

\end{document}
```

此处的第一行 `\documentclass{article}` 中包含了一个控制序列（或称命令/标记）。所谓控制序列，是以反斜杠 `\` 开头，以第一个**空格或非字母** 的字符结束的一串文字。它们不被输出，但是他们会影响输出文档的效果。这里的控制序列是 `documentclass`，它后面紧跟着的 `{article}` 代表这个控制序列有一个必要的参数，该参数的值为 `article`。这个控制序列的作用，是调用名为 `article` 的文档类。



在注释行之后出现了控制序列 `begin`。这个控制序列总是与 `end` 成对出现。这两个控制序列以及他们中间的内容被称为「环境」；它们之后的第一个必要参数总是**一致的**，被称为环境名。只有在 `document` 环境中的内容，才会被正常输出到文档中去或是作为控制序列对文档产生影响。也就是说，在 `\end{document}` 之后插入任何内容都是无效的。



从 `\documentclass{article}` 开始到 `\begin{document}` 之前的部分被称为导言区。你可以将导言区理解为是对整篇文档进行设置的区域——在导言区出现的控制序列，往往会影响整篇文档的格式。

> 比如，我们通常在导言区设置页面大小、页眉页脚样式、章节标题样式等等。



## 使用中英混排

所谓宏包，就是一系列控制序列的合集。这些控制序列太常用，以至于人们会觉得每次将他们写在导言区太过繁琐，于是将他们打包放在同一个文件中，成为所谓的宏包（台湾方面称之为「巨集套件」）。`\usepackage{}` 可以用来调用宏包。

所以要使用中文应该使用ctex 宏包

> ```
> \documentclass[UTF8]{ctexart}
> ```



## 基本骨架

在文档类 `article`/`ctexart` 中，定义了五个控制序列来调整行文组织结构。他们分别是

- `\section{·}`
- `\subsection{·}`
- `\subsubsection{·}`
- `\paragraph{·}`
- `\subparagraph{·}`

> 在`report`/`ctexrep`中，还有`\chapter{·}`；在文档类`book`/`ctexbook`中，还定义了`\part{·}`。

```latex
\documentclass[UTF8]{ctexart}
\usepackage{amsmath}

\title{数学建模第一次作业}
\author{江诗毅}
\date{\today}

\begin{document}
\maketitle
\tableofcontents


\section{你好中国}
中国位置。
\subsection{北京}
北京在北京.
\subsubsection{天安门}
天安门也在被禁锢。
\paragraph{位置}
在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.
\subparagraph{门}
门.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.在北京中心.

\subsection{Hello 山东}
\paragraph{山东大学} is one of the best university in 山东。


\end{document}
```





注意：

 LaTeX 将一个换行当做是一个简单的空格来处理，如果需要换行另起一段，则需要用两个换行（一个空行）来实现。



## 数学模式

LaTeX 的数学模式有两种：行内模式 (inline) 和行间模式 (display)。前者在正文的行文中，插入数学公式；后者独立排列单独成行，并自动居中。



在行文中，使用 `$ ... $` 可以插入行内公式，行内公式也可以使用 `\(...\)` ，使用 `\[ ... \]` 可以插入行间公式，如果需要对行间公式进行编号，则可以使用 `equation` 环境：

```latex
\begin{equation}
...
\end{equation}
```



### 公式之间的空格

> \quad

### 上下标

在数学模式中，需要表示上标，可以使用 `^` 来实现（下标则是 `_`）。**它默认只作用于之后的一个字符**，如果想对连续的几个字符起作用，请将这些字符用花括号 `{}` 括起来，例如：

```
\[ z = r\cdot e^{2\pi i}. \]
```

示例：

```latex
\documentclass{article}
\usepackage{amsmath}
\begin{document}
Einstein 's $E=mc^2$.

\[ E=mc^2. \]

\begin{equation}
E=mc^2.
\end{equation}
\end{document}
```

> 在这里提一下关于公式标点使用的规范。行内公式和行间公式对标点的要求是不同的：行内公式的标点，应该放在数学模式的限定符之外，而行间公式则应该放在数学模式限定符之内。



### 根式与分式

1. 行内一般用 \sfrac 或者 \tfrac
2. 行间用 \frac

根式用 `\sqrt{·}` 来表示，分式用 `\frac{·}{·}` 来表示（第一个参数为分子，第二个为分母）。

示例代码（请保存后，使用 XeLaTeX 编译，查看效果）：

```
\documentclass{article}
\usepackage{amsmath}
\usepackage{xfrac}
\begin{document}
$\sqrt{x}$, $\frac{1}{2}$.

\( \sfrac{1}{2} \)

\[ \sqrt{x}, \]

\[ \frac{1}{2}. \]
\end{document}
```

可以发现，在行间公式和行内公式中，分式的输出效果是有差异的。如果要强制行内模式的分式显示为行间模式的大小，可以使用 `\dfrac`, 反之可以使用 `\tfrac`。

> 在行内写分式，你可能会喜欢 `xfrac` 宏包提供的 `\sfrac` 命令的效果。

> 排版繁分式，你应该使用 `\cfrac` 命令。



### 运算符

一些小的运算符，可以在数学模式下直接输入；另一些需要用控制序列生成，如

```
\[ \pm\; \times \; \div\; \cdot\; \cap\; \cup\;
\geq\; \leq\; \neq\; \approx \; \equiv \]
```

连加、连乘、极限、积分等大型运算符分别用 `\sum`, `\prod`, `\lim`, `\int` 生成。他们的上下标在行内公式中被压缩，以适应行高。我们可以用 `\limits` 和 `\nolimits` 来强制显式地指定是否压缩这些上下标。例如：

```
$ \sum_{i=1}^n i\quad \prod_{i=1}^n $

$ \sum\limits _{i=1}^n i\quad \prod\limits _{i=1}^n $

\[ \lim_{x\to0}x^2 \quad \int_a^b x^2 dx \]

\[ \lim\nolimits _{x\to0}x^2\quad \int\nolimits_a^b x^2 dx \]
```

多重积分可以使用 `\iint`, `\iiint`, `\iiiint`, `\idotsint` 等命令输入。

```
\[ \iint\quad \iiint\quad \iiiint\quad \idotsint \]
```



### 定界符（括号等）

各种括号用 `()`, `[]`, `\{\}`, `\langle\rangle` 等命令表示；注意花括号通常用来输入命令和环境的参数，所以在数学公式中它们前面要加 `\`。因为 LaTeX 中 `|` 和 `\|` 的应用过于随意，amsmath 宏包推荐用 `\lvert\rvert` 和 `\lVert\rVert` 取而代之。

为了调整这些定界符的大小，amsmath 宏包推荐使用 `\big`, `\Big`, `\bigg`, `\Bigg` 等一系列命令放在上述括号前面调整大小。

> 有时你可能会觉得 amsmath 宏包提供的定界符放大命令不太够用。通常这意味着你的公式太过复杂。此时你应当首先考虑将公式中的部分提出去，以字母符号代替以简化公式。如果你真的想要排版如此复杂的公式，你可以参考我[这篇博文](https://liam.page/2018/11/09/the-bigger-than-bigger-delimiter-in-LaTeX/)。

```
\[ \Biggl(\biggl(\Bigl(\bigl((x)\bigr)\Bigr)\biggr)\Biggr) \]
\[ \Biggl[\biggl[\Bigl[\bigl[[x]\bigr]\Bigr]\biggr]\Biggr] \]
\[ \Biggl \{\biggl \{\Bigl \{\bigl \{\{x\}\bigr \}\Bigr \}\biggr \}\Biggr\} \]
\[ \Biggl\langle\biggl\langle\Bigl\langle\bigl\langle\langle x
\rangle\bigr\rangle\Bigr\rangle\biggr\rangle\Biggr\rangle \]
\[ \Biggl\lvert\biggl\lvert\Bigl\lvert\bigl\lvert\lvert x
\rvert\bigr\rvert\Bigr\rvert\biggr\rvert\Biggr\rvert \]
\[ \Biggl\lVert\biggl\lVert\Bigl\lVert\bigl\lVert\lVert x
\rVert\bigr\rVert\Bigr\rVert\biggr\rVert\Biggr\rVert \]
```



### 省略号

省略号用 `\dots`, `\cdots`, `\vdots`, `\ddots` 等命令表示。`\dots` 和 `\cdots` 的纵向位置不同，前者一般用于有下标的序列。

```latex
\[ x_1,x_2,\dots ,x_n\quad 1,2,\cdots ,n\quad
\vdots\quad \ddots \]
```



### 矩阵

`amsmath` 的 `pmatrix`, `bmatrix`, `Bmatrix`, `vmatrix`, `Vmatrix` 等环境可以在矩阵两边加上各种分隔符。

```
\[ \begin{pmatrix} a&b\\c&d \end{pmatrix} \quad
\begin{bmatrix} a&b\\c&d \end{bmatrix} \quad
\begin{Bmatrix} a&b\\c&d \end{Bmatrix} \quad
\begin{vmatrix} a&b\\c&d \end{vmatrix} \quad
\begin{Vmatrix} a&b\\c&d \end{Vmatrix} \]
```

使用 `smallmatrix` 环境，可以生成行内公式的小矩阵。

```
Marry has a little matrix $ ( \begin{smallmatrix} a&b\\c&d \end{smallmatrix} ) $.
```



### 多行公式

有的公式特别长，我们需要手动为他们换行；有几个公式是一组，我们需要将他们放在一起；还有些类似分段函数，我们需要给它加上一个左边的花括号。

#### 长公式

##### 不对齐

无须对齐的长公式可以使用 `multline` 环境。

```
\begin{multline}
x = a+b+c+{} \\
d+e+f+g
\end{multline}
```

效果：

![img](https://liam.page/uploads/teaching/LaTeX/figures/818901c1jw1e44jzfychej20dv02sjr6.jpg)

如果不需要编号，可以使用 `multline*` 环境代替。

##### 对齐

需要对齐的公式，可以使用 `aligned` *次环境*来实现，它必须包含在数学环境之内。

```
\[\begin{aligned}
x ={}& a+b+c+{} \\
&d+e+f+g
\end{aligned}\]
```



#### 公式组

无需对齐的公式组可以使用 `gather` 环境，需要对齐的公式组可以使用 `align` 环境。他们都带有编号，如果不需要编号可以使用带星花的版本。

```
\begin{gather}
a = b+c+d \\
x = y+z
\end{gather}
\begin{align}
a &= b+c+d \\
x &= y+z
\end{align}
```

#### 分段函数

分段函数可以用`cases`次环境来实现，它必须包含在数学环境之内。

```
\[ y= \begin{cases}
-x,\quad x\leq 0 \\
x,\quad x>0
\end{cases} \]
```

### 图片

关于 LaTeX 插图，首先要说的是：「LaTeX 只支持 `.eps` 格式的图档」这个说法是错误的。

在 LaTeX 中插入图片，有很多种方式。最好用的应当属利用 `graphicx` 宏包提供的 `\includegraphics` 命令。比如你在你的 TeX 源文件同目录下，有名为 `a.jpg` 的图片，你可以用这样的方式将它插入到输出文档中：

```
\documentclass{article}
\usepackage{graphicx}
\begin{document}
\includegraphics{a.jpg}
\end{document}
```

图片可能很大，超过了输出文件的纸张大小，或者干脆就是你自己觉得输出的效果不爽。这时候你可以用 `\includegraphics` 控制序列的可选参数来控制。比如

```
\includegraphics[width = .8\textwidth]{a.jpg}
```

这样图片的宽度会被缩放至页面宽度的百分之八十，图片的总高度会按比例缩放。

> `\includegraphics` 控制序列还有若干其他的可选参数可供使用，一般并用不到。感兴趣的话，可以去查看[该宏包的文档](http://texdoc.net/texmf-dist/doc/latex/graphics/graphicx.pdf)。

### 表格

`tabular` 环境提供了最简单的表格功能。它用 `\hline` 命令表示横线，在列格式中用 `|` 表示竖线；用 `&` 来分列，用 `\\` 来换行；每列可以采用居左、居中、居右等横向对齐方式，分别用 `l`、`c`、`r` 来表示。

```
\begin{tabular}{|l|c|r|}
 \hline
操作系统& 发行版& 编辑器\\
 \hline
Windows & MikTeX & TexMakerX \\
 \hline
Unix/Linux & teTeX & Kile \\
 \hline
Mac OS & MacTeX & TeXShop \\
 \hline
通用& TeX Live & TeXworks \\
 \hline
\end{tabular}
```



### 浮动体

插图和表格通常需要占据大块空间，所以在文字处理软件中我们经常需要调整他们的位置。`figure` 和 `table` 环境可以自动完成这样的任务；这种自动调整位置的环境称作浮动体(float)。我们以 `figure` 为例。

```
\begin{figure}[htbp]
\centering
\includegraphics{a.jpg}
\caption{有图有真相}
\label{fig:myphoto}
\end{figure}
```

`htbp` 选项用来指定插图的理想位置，这几个字母分别代表 here, top, bottom, float page，也就是就这里、页顶、页尾、浮动页（专门放浮动体的单独页面或分栏）。`\centering` 用来使插图居中；`\caption` 命令设置插图标题，LaTeX 会自动给浮动体的标题加上编号。注意 `\label` 应该放在标题命令之后。

> 图片和表格的各种特殊效果，限于篇幅此处无法详叙。请查看最后一章推荐的文档。
>
> 如果你想了解 LaTeX 的浮动体策略算法细节，你可以参考我博客中[关于浮动体的系列文章](https://liam.page/series/#LaTeX-%E4%B8%AD%E7%9A%84%E6%B5%AE%E5%8A%A8%E4%BD%93)
>
> 如果你困惑于「为什么图表会乱跑」或者「怎样让图表不乱跑」，请看[我的回答](https://www.zhihu.com/question/25082703/answer/30038248)。



## 版面设置

### 页边距

设置页边距，推荐使用 `geometry` 宏包。可以在[这里](http://texdoc.net/texmf-dist/doc/latex/geometry/geometry.pdf)查看它的说明文档。

比如我希望，将纸张的长度设置为 20cm、宽度设置为 15cm、左边距 1cm、右边距 2cm、上边距 3cm、下边距 4cm，可以在导言区加上这样几行：

```
\usepackage{geometry}
\geometry{papersize={20cm,15cm}}
\geometry{left=1cm,right=2cm,top=3cm,bottom=4cm}
```

### 页眉页脚

设置页眉页脚，推荐使用 `fancyhdr` 宏包。可以在[这里](http://texdoc.net/texmf-dist/doc/latex/fancyhdr/fancyhdr.pdf)查看它的说明文档。

比如我希望，在页眉左边写上我的名字，中间写上今天的日期，右边写上我的电话；页脚的正中写上页码；页眉和正文之间有一道宽为 0.4pt 的横线分割，可以在导言区加上如下几行：

```
\usepackage{fancyhdr}
\pagestyle{fancy}
\lhead{\author}
\chead{\date}
\rhead{152xxxxxxxx}
\lfoot{}
\cfoot{\thepage}
\rfoot{}
\renewcommand{\headrulewidth}{0.4pt}
\renewcommand{\headwidth}{\textwidth}
\renewcommand{\footrulewidth}{0pt}
```

### 首行缩进

CTeX 宏集已经处理好了首行缩进的问题（自然段前空两格汉字宽度）。因此，使用 CTeX 宏集进行中西文混合排版时，我们不需要关注首行缩进的问题。

> 如果你因为某些原因选择不适用 CTeX 宏集（不推荐）进行中文支持和版式设置，则你需要做额外的一些工作。
>
> - 调用 `indentfirst` 宏包。具体来说，中文习惯于每个自然段的段首都空出两个中文汉字的长度作为首行缩进，但西文行文习惯于不在逻辑节（`\section` 等）之后缩进。使用改宏包可使 LaTeX 在每个自然段都首行缩进。
> - 设置首行缩进长度 `\setlength{\parindent}{2\ccwd}`。其中 `\ccwd` 是 `xeCJK` 定义的宏，它表示当前字号中一个中文汉字的宽度。

### 行间距

我们可以通过 `setspace` 宏包提供的命令来调整行间距。比如在导言区添加如下内容，可以将行距设置为字号的 1.5 倍：

```
\usepackage{setspace}
\onehalfspacing
```

具体可以查看该宏包的[文档](http://texdoc.net/texmf-dist/doc/latex/setspace/README)。

> 请注意用词的差别：
>
> - 行距是字号的 1.5 倍；
> - 1.5 倍行距。
>
> 事实上，这不是设置 1.5 倍行距的正确方法，请参考[这篇博文](https://liam.page/2013/10/17/LaTeX-Linespace/)。另外，[RuixiZhang](https://github.com/RuixiZhang42) 开发了 [zhlineskip](https://github.com/CTeX-org/ctex-kit/tree/master/zhlineskip) 宏包，提供了对中西文混排更细致的行距控制能力。

### 段间距

我们可以通过修改长度 `\parskip` 的值来调整段间距。例如在导言区添加以下内容

```
\addtolength{\parskip}{.4em}
```

则可以在原有的基础上，增加段间距 0.4em。如果需要减小段间距，只需将该数值改为负值即可。







