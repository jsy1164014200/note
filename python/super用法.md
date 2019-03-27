[知乎链接](https://www.zhihu.com/question/20040039)

```python
class Child(Parent):  
    def __init__(self):  
         Parent.__init__(self) 
```

这种方式与super(Child, self).__init__() 没有区别



```python
def super(cls, inst):
    mro = inst.__class__.mro()
    return mro[mro.index(cls) + 1]
```

