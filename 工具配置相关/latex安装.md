一、选择安装LaTeX发行版

LaTeX有很多发型版，TeX Live就是其中一种。TeX Live 是 TUG (TeX User Group) 维护和发布的 TeX 系统，可说是「官方」的 TeX 系统。TeX Live可以保持在跨操作系统平台、跨用户的一致性。而且TeX Live在Ubuntu18.04上的安装也比较方便。

sudo apt-get install texlive-full

二、安装XeLaTeX编译引擎

sudo apt-get install texlive-xetex

三、安装中文支持包，使用的是xeCjK，中文处理技术也有很多，xeCJK是成熟且稳定的一种。

sudo apt-get install texlive-lang-chinese

四、安装图形化界面

图形化界面有TeXworks，TeXmaker和TeXstudio等，很多在Ubuntu终端都可以直接安装，下载尝试了TeXmaker和TeXstudio，因为是初学者也没有感觉到太大的不同。而他们本身是一个工具，个人觉得选择自己喜欢的就可以，个人跟喜欢TeXstudio的界面风格。至于安装方式终端apt-get install即可

sudo apt-get install texstudio

这里要注意，编译时需要设置编译器为XeLaTeX，TeXstudio中在Options->Configure TeXstudio->Build->Default Compiler中更改默认编译器为XeLaTeX即可。在配置中可以更改软件界面语言，将Options->Configure TeXstudio->General->Language更改为zh-CN即可将界面设置为中文。

五、编译测试

新建文件，在文本编辑框输入

\documentclass{article}

\usepackage{xeCJK}

\begin{document}

hello,你好

\end{document}

按Ｆ５进行编译预览。


