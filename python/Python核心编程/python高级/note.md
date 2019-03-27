1. python 中导入模块不需要写路径，路径都在import sys,sys.path.append("");添加路径即可

2. 重新导入
  `from imp import *`
  `reload(test)`

3. 循环导入
  尽量避免

4. == 跟 is的区别
  `id(a);`能返回变量的地址
- is判断是否指向同一个地址
- == 判断值是否相同

5. 深拷贝跟浅拷贝
- 浅拷贝：a b指向同一个地址`= `
- 深拷贝：理解为完全跟原先的断绝关系
`import copy `then
`c = copy.deepcopy(a)` c指向新的地址

6. copy.copy()的用法
- 对于不变量，直接实行浅拷贝
- 对于可变量，只拷贝一层
```python
a = [1,2,3]
b = [4,5,6]
c = [a,b]
e = copy.copy(c)
id(c) #123456789
id(e) #123456456
a.append(4)
c #[[1,2,3,4],[4,5,6]]
e #[[1,2,3,4],[4,5,6]]
```
对于元组
```python
a = [1,2,3]
b = [4,5,6]
c = (a,b)
e = copy.copy(c)
id(c) #123456789
id(e) #123456789
a.append(4)
c #[[1,2,3,4],[4,5,6]]
e #[[1,2,3,4],[4,5,6]]
```
7. 原码，反码，补码

-1原码：10000000001
-1补码：11111111110
用补码相加
```python
bin(18) #0b10010
oct(18) #0o22
hex(18) #0x12

# 其他进制转成十进制
int("#0b10010",2)
int("#0o10010",8)
int("#0x10010",16)
```

8. 位运算

& | ^ ~ <<  >>

9. 私有化
- xx:公共变量或者方法 public
- _x:保护的变量或者方法 protected 子类可以访问
- __xx:私有变量或者方法 private
- __xx__:魔法方法，Python自带
- xx_ :用于避免跟Python关键字冲突

**但是通过_class__xx访问私有属性**

9. property用法

```python
#class Test(object):
#    def __init__(self):
#        self.__num = 100
    
#    def setNum(self,newNum):
#        self.__num = newNum

#    def getNum(self):
#       return self.__num

#    num = property(getNum,setNum)

####简写用装饰器
class Test(object):
    def __init__(self):
        self.__num = 100
    
    @property
    def num(self):
        return self.__num

    @num.setter
    def num(self,newNum):
        self.__num = newNum

t = Test()
t.num = 200 #会调用setNum
print(t.num) #会调用getNum
```
10. 可迭代：元组，列表，字典，生成器，生成器函数等

判断方法
```python
from collections import Iterable #可迭代，能够for
from collections import Iterator #迭代器
isinstance([],Iterable) # 可迭代 true
isinstance([],Iterator) # 迭代器 false

### 1.生成器是迭代器。 2.可以通过iter([1,3,4])将 可迭代 变成迭代器
```

11. 闭包
```python
def test(number):
    def test_in(number_in):
        print(number+number_in)
    
    return test_in

test(100)(1100)#1200
```

12. 装饰器

```python
def w1(func):
    def inner():
        print("正在验证权限-----")
        func()
    return inner

def test():
    print("调用成功")

test = w1(test) # 相当于给test里面加上了一端验证

test()
# 正在验证权限-----
# 调用成功



#现在就可以把w1理解成一个装饰器
def decorate(func):
    def inner(*args,**kwargs):
        print("正在验证权限-----")
        result = func(*args,**kwargs)
        return result
    return inner

@decorate#这里自动执行 test = w1(test)
def test(a,b,c,d):
    print(a+b+c+d)
    print("调用成功")
    return "hahaha"

a = test(4,5,6,7)
print(a)
```
一种带参数的装饰器
```python
def func_arg(string):
    def decorate(func):
        def inner(*args,**kwargs):
            print("正在验证权限-----"+string)
            result = func(*args,**kwargs)
            return result
        return inner
    return decorate

@func_arg("hehe")#这里自动执行 test = w1(test)
def test(a,b,c,d):
    print(a+b+c+d)
    print("调用成功")
    return "hahaha"

test()
```

13. 作用域
`globals()`
`locals()`
```Python
#给class或者实例对象可以直接添加属性
class Person(object):
    def __init__(self,name,age):
        self.name = name
        self.age = age
    
def run(self):
    print(self.name+"run")    

p1 = Person("p1",10)
p1.addr = "haha"

# 添加方法
import types
p1.run = types.MethodType(run,p1)
```
python中静态方法跟类方法
```python
@staticmethod
def test():
    print("--static---")

@classmethod
def test2():
    print("---test2---")
```

14. 限制添加属性

```python
__slots__  =("name","age") # 只能添加name，age这两个属性，其他的都不能添加
```

15. 生成器（节省空间）

将列表产生形式的`b = [x*2 for x in range(10)]`该成
`b = (x*2 for x in range(10))`

```Python
#在没有第三个变量的情况下交换 a,b
a,b = b,a
a = a+b
b = a-b
a = a-b
```
```python
def createNum():
    a,b = 0,1
    for i in range(5):
        yield b
        a,b = b,a+b

# 该函数只会执行到yield，只是生成一个生成器
# 如果 temp = yield b #temp是none
t = createNum()
t.send("haha")
temp = "haha"
# 之后不停的循环执行
```
应用：可以实现挟程
```python
def test1():
    while True:
        print(11111111111)
        yield None

def test2():
    while True:
        print(2222222222)
        yield None

t1 = test1()
t2 = test2()
while True:
    t1.__next__()
    t2.__next__()
```

16. GC垃圾回收

- 基本类型，int，char，bool,string都是直接指代，不是引用，不存在深拷贝，浅拷贝
- intern机制：字符串，不可修改，开启intern机制，公用对象，引用计数为0则销毁

```python
# gc模块的基本使用

#必须要import gc模块，并且 is_enable()=True  才会启动⾃动垃圾回收
gc.set_debug(flags) #设置gc的debug⽇志，⼀般设置为gc.DEBUG_LEAK
gc.collect([generation]) #显式进⾏垃圾回收，可以输⼊参数，0代表只检查第⼀代的对象，1代表检查⼀，⼆代的对象，2代表检查⼀，⼆，三代的对象，如果不传参数，执⾏⼀个full collection，也就是等于传2。 返回不可达（unreachable objects）对象的数⽬
gc.get_threshold() #获取的gc模块中⾃动执⾏垃圾回收的频率
gc.set_threshold(threshold0[, threshold1[, threshold2]) #设置⾃动执⾏垃圾回收的频率
gc.get_count() #获取当前⾃动执⾏垃圾回收的计数器，返回⼀个⻓度为3的列表
```

17. 内建属性

- __new__ 生成实例所需属性
- __init__ 初始化函数，在__new__之后调用
- __class__ 实例化所在类
- __str__ 实例字符串表示,可读性，pring(类实例),如果没有实现，使用repr结果
- __repr__ 实例字符串表示，准确性
- __del__ 析构 删除实例
- __dict__ 实例自定义属性
- __getattribute__ 属性访问拦截器 
- __bases__ 类的所有父类构成元素

18. functools
```python
import functools
def showarg(*args, **kw):
print(args)
print(kw)
p1=functools.partial(showarg, 1,2,3)
p1() # (1,2,3) {}
p1(4,5,6) # (1,2,3,4,5,6) {}
p1(a='python', b='itcast') # (1,2,3) {"a":"python","b":"itcast"}
p2=functools.partial(showarg, a=3,b='linux')
p2()
p2(1,2)
p2(a='python', b='itcast')
```
```python
import functools
def note(func):
"note function"
	@functools.wraps(func)
	def wrapper():
		"wrapper function"
		print('note something')
		return func()
	return wrapper
@note
def test():
"test function"
	print('I am test')

    
test()
print(test.__doc__)

#这样一来 test里面的函数说明文字就不会被覆盖了
note something
I am test
test function
```

直接用Python搭建一个静态服务器
```python
python -m http.server PORT
```

pdp调试工具


总结了一下Python中几个常用的高级用法，copy.copy() copy.deepcopy() 浅拷贝 的区别，is 跟 == 的区别，常用的装饰器@property，@staticmethod ，@classmethod，生成器，内建属性，闭包，实现函数复用神器：装饰器（有参，无参，不确定参的基本用法），等等……