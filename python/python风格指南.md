# python编码风格指南

> 结合了Google风格，以及 Pocoo风格（这两种风格都是对pep8的扩展）

## 导入模块

导入总应该放在文件顶部, 位于模块注释和文档字符串之后, 模块全局变量和常量之前. 导入应该按照从最通用到最不通用的顺序分组:

1. 标准库导入
2. 第三方库导入
3. 应用程序指定导入

每种分组中, 应该根据每个模块的完整包路径按字典序排序, 忽略大小写.

```python
import foo
from foo import bar
from foo.bar import baz
from foo.bar import Quux
from Foob import ar
```



## 缩进：

4个空格。不使用制表符，没有例外。

## 最大行长：

软限制为 79 个字符，不超过 84 个字符。尝试合理放置 break 、 continue 和 return 声明来避免代码过度嵌套。

续行:

**不要反斜杠来续行**，续行应对齐最后一个点号或等于号，或者缩进四个空格:

```
this_is_a_very_long(function_call, 'with many parameters') 
    .that_returns_an_object_with_an_attribute
    
MyModel.query.filter(MyModel.scalar > 120) 
             .order_by(MyModel.name.desc()) 
             .limit(10)
```

如果你在括号内的换行，那么续行应对齐括号:

```
this_is_a_very_long(function_call, 'with many parameters',
                    23, 42, 'and even more')
```

对于有许多元素的元组或列表，在起始括号后立即换行:

```
items = [
    'this is the first', 'set of items', 'with more items',
    'to come in this line', 'like this'
]
```

## 空行：

顶层函数和类由两个空行分隔，其它一个空行。不要使用过多空行来分隔代码 逻辑段。例如:

```
def hello(name):
    print 'Hello %s!' % name


def goodbye(name):
    print 'See you %s.' % name


class MyClass(object):
    """This is a simple docstring"""

    def __init__(self, name):
        self.name = name

    def get_annoying_name(self):
        return self.name.upper() + '!!!!111'
```

## 括号

宁缺毋滥的使用括号

除非是用于实现行连接, 否则不要在**返回语句或条件语句**中使用括号. 不过在元组两边使用括号是可以的.

## Shebang

大部分.py文件不必以#!作为文件的开始. 根据 [PEP-394](http://www.python.org/dev/peps/pep-0394/) , 程序的main文件应该以 #!/usr/bin/python2或者 #!/usr/bin/python3开始.

## 函数和方法

一个函数必须要有文档字符串, 除非它满足以下条件:

1. 外部不可见
2. 非常短小
3. 简单明了

文档字符串应该包含函数做什么, 以及输入和输出的详细描述. 通常, 不应该描述”怎么做”, 除非是一些复杂的算法. 文档字符串应该提供足够的信息, 当别人编写代码调用该函数时, 他不需要看一行代码, 只要看文档字符串就可以了. 对于复杂的代码, 在代码旁边加注释会比使用文档字符串更有意义.

关于函数的几个方面应该在特定的小节中进行描述记录， 这几个方面如下文所述. 每节应该以一个标题行开始. 标题行以冒号结尾. 除标题行外, 节的其他内容应被缩进2个空格.

- Args:

  列出每个参数的名字, 并在名字后使用一个冒号和一个空格, 分隔对该参数的描述.如果描述太长超过了单行80字符,使用2或者4个空格的悬挂缩进(与文件其他部分保持一致). 描述应该包括所需的类型和含义. 如果一个函数接受*foo(可变长度参数列表)或者**bar (任意关键字参数), 应该详细列出*foo和**bar.

- Returns: (或者 Yields: 用于生成器)

  描述返回值的类型和语义. 如果函数返回None, 这一部分可以省略.

- Raises:

  列出与接口有关的所有异常.

```python
def fetch_bigtable_rows(big_table, keys, other_silly_variable=None):
    """Fetches rows from a Bigtable.

    Retrieves rows pertaining to the given keys from the Table instance
    represented by big_table.  Silly things may happen if
    other_silly_variable is not None.

    Args:
        big_table: An open Bigtable Table instance.
        keys: A sequence of strings representing the key of each table row
            to fetch.
        other_silly_variable: Another optional variable, that has a much
            longer name than the other args, and which does nothing.

    Returns:
        A dict mapping keys to the corresponding table row data
        fetched. Each row is represented as a tuple of strings. For
        example:

        {'Serak': ('Rigel VII', 'Preparer'),
         'Zim': ('Irk', 'Invader'),
         'Lrrr': ('Omicron Persei 8', 'Emperor')}

        If a key from the keys argument is missing from the dictionary,
        then that row was not found in the table.

    Raises:
        IOError: An error occurred accessing the bigtable.Table object.
    """
    pass


class SampleClass(object):
    """Summary of class here.

    Longer class information....
    Longer class information....

    Attributes:
        likes_spam: A boolean indicating if we like SPAM or not.
        eggs: An integer count of the eggs we have laid.
    """

    def __init__(self, likes_spam=False):
        """Inits SampleClass with blah."""
        self.likes_spam = likes_spam
        self.eggs = 0

    def public_method(self):
        """Performs operation blah."""
```



## 字符串

即使参数都是字符串, 使用%操作符或者格式化方法格式化字符串. 不过也不能一概而论, 你需要在+和%之间好好判定.

```
Yes: x = a + b
     x = '%s, %s!' % (imperative, expletive)
     x = '{}, {}!'.format(imperative, expletive)
     x = 'name: %s; score: %d' % (name, n)
     x = 'name: {}; score: {}'.format(name, n)
No: x = '%s%s' % (a, b)  # use + in this case
    x = '{}{}'.format(a, b)  # use + in this case
    x = imperative + ', ' + expletive + '!'
    x = 'name: ' + name + '; score: ' + str(n)
```

避免在循环中用+和+=操作符来累加字符串. 由于字符串是不可变的, 这样做会创建不必要的临时对象, 并且导致二次方而不是线性的运行时间. 作为替代方案, 你可以将每个子串加入列表, 然后在循环结束后用 `.join` 连接列表. (也可以将每个子串写入一个 `cStringIO.StringIO` 缓存中.)

```
Yes: items = ['<table>']
     for last_name, first_name in employee_list:
         items.append('<tr><td>%s, %s</td></tr>' % (last_name, first_name))
     items.append('</table>')
     employee_table = ''.join(items)
No: employee_table = '<table>'
    for last_name, first_name in employee_list:
        employee_table += '<tr><td>%s, %s</td></tr>' % (last_name, first_name)
    employee_table += '</table>'
```

在同一个文件中, 保持使用字符串引号的一致性. 使用单引号’或者双引号”之一用以引用字符串, 并在同一文件中沿用. 在字符串内可以使用另外一种引号, 以避免在字符串中使用. GPyLint已经加入了这一检查.

(译者注:GPyLint疑为笔误, 应为PyLint.)

```
Yes:
     Python('Why are you hiding your eyes?')
     Gollum("I'm scared of lint errors.")
     Narrator('"Good!" thought a happy Python reviewer.')
No:
     Python("Why are you hiding your eyes?")
     Gollum('The lint. It burns. It burns us.')
     Gollum("Always the great lint. Watching. Watching.")
```

为多行字符串使用三重双引号”“”而非三重单引号’‘’. 当且仅当项目中使用单引号’来引用字符串时, 才可能会使用三重’‘’为非文档字符串的多行字符串来标识引用. 文档字符串必须使用三重双引号”“”. 不过要注意, 通常用隐式行连接更清晰, 因为多行字符串与程序其他部分的缩进方式不一致.



## TODO注释

为临时代码使用TODO注释, 它是一种短期解决方案. 不算完美, 但够好了.

TODO注释应该在所有开头处包含”TODO”字符串, 紧跟着是用括号括起来的你的名字, email地址或其它标识符. 然后是一个可选的冒号. 接着必须有一行注释, 解释要做什么. 主要目的是为了有一个统一的TODO格式, 这样添加注释的人就可以搜索到(并可以按需提供更多细节). 写了TODO注释并不保证写的人会亲自解决问题. 当你写了一个TODO, 请注上你的名字.

```
# TODO(kl@gmail.com): Use a "*" here for string repetition.
# TODO(Zeke) Change this to use relations.
```



## 表达式和语句

- 常规空格规则：

  不是单词的一元运算符不使用空格（例如： `-` 、 `~` 等等），在圆括号 也是这样。用空格包围二元运算符。对:`exp = -1.05 value = (item_value / item_count) * offset / exp value = my_list[index] value = my_dict['key'] `错:`exp = - 1.05 value = ( item_value / item_count ) * offset / exp value = (item_value/item_count)*offset/exp value=( item_value/item_count ) * offset/exp value = my_list[ index ] value = my_dict ['key'] `

- 禁止 Yoda 语句：

  永远不要用变量来比较常量，而是用常量来比较变量：对:`if method == 'md5':     pass `错:`if 'md5' == method:     pass `

- 比较：

  针对任意类型使用 `==` 和 `!=`针对单一类型使用 `is` 和 `is not` （例如： `foo is not None` ）永远不要与 `True` 或 `False` 作比较（例如永远不要写 `foo == False` ，而应当写 `not foo` ）

- 排除检验：

  使用 `foo not in bar` 而不是 `not foo in bar`

- 实例检验：

  使用 `isinstance(a, C)` 而不是 `type(A) is C` ，但是通常应当避免检验 实例，而应当检验特性。

## 命名约定

- 类名： `CamelCase` ，缩写词大写（ `HTTPWriter` 而不是 `HttpWriter` ）
- 变量名： `lowercase_with_underscores`
- 方法和函数名： `lowercase_with_underscores`
- 常量： `UPPERCASE_WITH_UNDERSCORES`
- 预编译正则表达式： `name_re`

被保护的成员以单个下划线作为前缀，混合类则使用双下划线。

如果使用关键字作为类的名称，那么在名称末尾添加下划线。与内置构件冲突是允许 的，请 **一定不要** 用在变量名后添加下划线的方式解决冲突。如果函数需要访问 一个隐蔽的内置构件，请重新绑定内置构件到一个不同的名字。

- 函数和方法参数:

  类方法: `cls` 作为第一个参数

  实例方法:`self` 作为第一个参数

  用于属性的 lambda 表达式应该把第一个参数替换为 `x` ， 像 `display_name = property(lambda x:x.real_name or x.username)` 中一样

**Python之父Guido推荐的规范**

| Type                       | Public             | Internal                                                     |
| -------------------------- | ------------------ | ------------------------------------------------------------ |
| Modules                    | lower_with_under   | _lower_with_under                                            |
| Packages                   | lower_with_under   |                                                              |
| Classes                    | CapWords           | _CapWords                                                    |
| Exceptions                 | CapWords           |                                                              |
| Functions                  | lower_with_under() | _lower_with_under()                                          |
| Global/Class Constants     | CAPS_WITH_UNDER    | _CAPS_WITH_UNDER                                             |
| Global/Class Variables     | lower_with_under   | _lower_with_under                                            |
| Instance Variables         | lower_with_under   | _lower_with_under (protected) or __lower_with_under (private) |
| Method Names               | lower_with_under() | _lower_with_under() (protected) or __lower_with_under() (private) |
| Function/Method Parameters | lower_with_under   |                                                              |
| Local Variables            | lower_with_under   |                                                              |



## 文档字符串

- 文档字符串约定:

  所有的文档字符串为 Sphinx 可理解的 reStructuredText 格式。它们的形态 因行数不同而不同。如果只有一行，三引号闭合在同一行，否则开头的三引号 与文本在同一行，结尾的三引号独立一行:`def foo():     """This is a simple docstring"""   def bar():     """This is a longer docstring with so much information in there     that it spans three lines.  In this case the closing triple quote     is on its own line.     """ `

- 模块头:

  模块头包含一个 utf-8 编码声明（即使没有使用非 ASCII 字符，也始终推 荐这么做）和一个标准的文档字符串:`# -*- coding: utf-8 -*- """     package.module     ~~~~~~~~~~~~~~      A brief description goes here.      :copyright: (c) YEAR by AUTHOR.     :license: LICENSE_NAME, see LICENSE_FILE for more details. """ `谨记使用合适的版权和许可证文件以利于通过 Flask 扩展审核。

## 注释

为了提高可读性, 注释应该至少离开代码2个空格.

注释的规则与文档字符串类似。两者都使用 reStructuredText 格式。如果一个 注释被用于一个说明类属性，在起始的井号（ `#` ）后加一个冒号:

```
class User(object):
    #: the name of the user as unicode string
    name = Column(String)
    #: the sha1 hash of the password + inline salt
    pw_hash = Column(String)
```

## Main用法

即使是一个打算被用作脚本的文件, 也应该是可导入的. 并且简单的导入不应该导致这个脚本的主功能(main functionality)被执行, 这是一种副作用. 主功能应该放在一个main()函数中.

在Python中, pydoc以及单元测试要求模块必须是可导入的. 你的代码应该在执行主程序前总是检查 `if __name__ == '__main__'` , 这样当模块被导入时主程序就不会被执行.

```
def main():
      ...

if __name__ == '__main__':
    main()
```