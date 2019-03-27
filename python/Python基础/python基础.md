# python 基础

1. help()   在解释器中输入 help() 进入一个 交互式的 帮助信息界面 显示 你想要的模块信息
2.  // 才是整除
3. 正负数的 除法都是向下圆整的 -10//3 = -4'
4. ** 幂运算优先级 比 负号高  -3**2 = -9
5. 0x   十六  0 八进制 0b 二进制
6. python的变量是没有默认值的，，定义也就必须赋值
7. round() 是向最接近它的圆整，一样近的时候 向偶数圆整
8. (1+2j) * (2+3j)  python 原生支持 负数
9. 注意 解释器 打印的 东西 跟print打印的不一样,这也是 repr()  与str()的区别

> : "hellow\nkdfj"
>
> : 'hellow\nkdf'
>
> : print("hellow\nkdfj")
>
> : hellow
>
>  kdfj

10. ‘’‘bababa’‘’   """dsjkfsdkfjd"""表示原始字符串
11. print(r'') print(r"") 与 三引号 repr一样都表示 原始字符串 ，除了 要对 单 双引号进行转义

重新认识一下print()

> print('hello','world',sep='_',end='')  # sep 默认是空格  end默认是换行

## 列表 元组

切片 索引

1. 拷贝序列 a = b[:]  *所以切片赋值是深拷贝*

2. 切片 指定起点，终点，步长

3. 列表的加，乘法 都是拷贝 ，不会出现 浅拷贝的问题
4. 成员资格检查  用in
5. 使用 del 来删除元素
6. 插入元素用 number[1:1] 这样就能给 Number 第一个元素后面插入元素



常用方法，以下都是 实例的方法

- append()
- clear()
- copy()   === [:]
- extend([])  比加法效率高，，因为加法会 先拷贝 再相加
- inedx() 找到第一次出现的索引  找不到会抛出异常
- pop()
- remove(元素)
- reverse() 
- sort(key=function)  指定一个排序函数

元祖 

1. 一定要区别 (43)  (43,)
2.  values = 1,2,3  是元组



## 字符串

1. > "hello %s,%s enfkdj for ya?" % ('dsjfk','dsjfk')  这个是python2 中常用的 现在基本不用这个 用format

2. `"{1}dkfjs{0}kf".format("kjfk","dkjf")`使用编号

3. "{var}kjdfk{bar}kdsj".format(var="dkf",bar="kdjslj") 使用命名占位符

> 关于格式用法翻书 45面 《python基础教程》

常用方法，实例方法(常见的错误，因为一般字符串是常量，所以 它一般不会修改本身，而是返回修改后的值)

- center(int)

- find("dkfj") 找到返回索引 否则返回-1  现在一般用 in

- join()

- lower()

- replace("","")

- split("")

- strip(" *!")在开头结尾 删除 所有的 列举出的字符

- translate(str.maketrans("cs","kz","kv"),delete)  可以指定要删除的字符

- isbababababa()  判断字符串是否符合一些条件字典

## 字典dict

与列表类似的基本操作

1. 构造字典
   - dict(name="dfjk",age="dj",bar="dfj",key="kdfj")  
   - dict([('name','jsy'),('age',18)])
2. len(d) 
3. del d[k]
4. k in d 

与字符串结合的操作

> "fdkjk{name},sdfsd{age}".format_map(dict([('name','jsy'),('age',18)]))

常用的方法

- clear() 清除
- copy()  它是浅拷贝，只是拷贝了一层 （具体的在python 核心编程中）
- fromkeys(['key1','key2'])  值对应是None
- get('name',default)
- items()
- keys()
- pop(key)
- popitem() 随机弹出一个
- setdefault(key,value)
- update(dict)  使用另一个字典添加进去
-  values()



## 集合 set



## 赋值魔法

1. 序列解包(可迭代对象解包)
   - x,y,z = 1,2,3
   - x,y = y,x
   - key, value = {'name':'jsy'}
   - a,*b,c = [1,2,3,4,5]
2. 链式赋值
   - `x = y = func()`  ===  `y = func()  x = y`  !==  `y = func()  x= func()`
3.  `a = "haha" if 1==1 else "boom"`  三元表达式
4. is  判断的是 两个地址
5. python 可以用链式比较  `0 < a < 2`

6. 迭代字典 `for key,value in d.items() : `
7. 并行迭代 使用zip()
8. 充分使用 元组，序列，字典的 推导操作

解包操作

1. `params = 1,2`  add(*params) === add(1,2)
2. `params = {"a":1,"b":2}` add(**params) === add(a=1,b=2)



> 在闭包中 如果要使用 外部的变量 要指定 nonlocal



简单的递归

> def test(n):
>
> ​	return 1 if n == 1 else n*test(n-1)



reduce

> reduce(lambda x,y:x+y,[1,2,3,4])   # 10

## 类

1. 类中的static 属性 只能用 类名.修改，实例只会覆盖 不会修改

常用的判断方法

1. issubclass(Son,Father)
2. isinstance(x,class)
3. Son.__ bases __
4. setattr(' ',x)  getattr('',default)



### 异常 警告

try except finally

from warnings import warn 

warn("babababa")



### 魔法方法

1. __ del __析构

2. __ init __ 如果是子类一定要记得 在其中 super. __ init __()

3. 基本的序列 协议

   ```python
   def check_index(key):
       if not isinstance(key,int): 
           raise TypeError
       if key < 0: 
           raise IndexError
           
   class ArithmeticSquence:
       def __init__(self,start=0,step=1):
           self.start = start
           self.step = step
           self.changed = {}
       
       def __getitem__(self,key):
           check_index(key)
           return self.changed.get(key,self.start + key*self.step)
           
       def __setitem__(self,key,value):
           check_index(key)
           self.changed[key] = value
       
       def __len__():
           return 0
       
       def __del__():
           print("del success")
           
   ```

4. 可以从 list,dict,str派生出来

   ```python
   class CounterList(list):
       def __init__(self,*args):
           super(CounterList,self).__init__(args)
           self.counter = 0
          
       def __getitem__(self,index):
           self.counter += 1
           return super(CounterList,self).__getitem__(index)
   ```


property 函数 装饰器

@staticmethod

@classmethod

### 迭代器协议

- __ iter __()
- __ next __()

生成器  迭代器 都是可迭代的

### 生成器推导

> g = ((i+2)**2 for i in range(2,27))



## with 上下文管理器

__ enter __  进入时被执行，返回 对象能用 as后的变量接受 

__ exit __

## 文件操作

1. w
2. r
3. a
4. r+ w+
5. x

open后常用的方法(文件对象是可以迭代的)

1. write()
2. read(n)  读几个字符
3. seek(n) 光标移动到第几个位置
4. tell()  打印现在在哪
5. writelines(['','',''])
6. readlines()   返回一个 数组
7. readline() 

```python

```

## cgi socket 相关

区分一下同步异步 阻塞非阻塞

> 详见博文，有很多[例子](https://www.jianshu.com/p/aed6067eeac9)

1. 同步与异步是站在消息通知机制角度来说的（同步可能需要时刻去关心询问线程处理结果，异步注册了回调机制，无需关心）
2. 阻塞和非阻塞是站在线程等待调用结果的线程状态这个角度来说的，阻塞则是线程挂起等待调用结果返回；非阻塞是在等待结果的过程中，线程任然是活动状态，可能处理其他的任务



linux上的几种io模型[博文](https://www.jianshu.com/p/486b0965c296)

![输入图片说明](https://static.oschina.net/uploads/img/201604/20144245_Wtld.png)



同步阻塞 io

同步非阻塞io

多路复用io  [select poll epoll](https://www.jianshu.com/p/dfd940e7fca2)  以同步的方式 处理多个 io操作

异步阻塞io

异步非阻塞io



## 测试驱动编程

unittest

```python
import unittest
class ProductTestCase(unittest.TestCase):
   	def test_intergers(self):
        for babab:
            for babab:
        		self.assertEqual(p,x*y,"intetdksfjkjbababa")
                
if __name__ == "__main__":
    unittest.main()
```



doctest

```python
def square(x):
    '''
    计算平方返回结果
    >>> square(2)
    4
    >>> square(3)
    9
    '''
    return x*x

if __name__ == '__main__':
    import doctest
    doctest.testmod(square_module)
    
>>> python test.py -v   verbose 详尽
```



## pyLint可以用来检查代码 格式

## 性能分析 

> import cProfile
>
> cProfile.run(str)



## java   c 扩展python

见书 295 扩展python

## 可配置文件

```python
[numbers]
pi: 3232.2323
[messages]
greeting: Welcom to blog
    
from configparser import ConfigParser
CONFIGFILE = 'area.ini'
config = ConfigParser()
config.read(CONFIGFILE)
config['messages'].get('greeting')
config['numbers'].getfloat('pi')
```

## 日志

logging

```python
import logging

logging.basicConfig(level=logging.INFO,filename='mylog')

logging.info("babab")

logging.info("end")
```

