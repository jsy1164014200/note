# 1. pdb调试工具：基于命令行的调试工具
```
l->list显示当前代码
n->next向下执行一行
c->continue继续执行完代码
b num -> 断点
clear num->清除断点 

a ->打印所有形参数据
p params->打印一个参数值

s -> step 进入一个函数
r -> 快速执行完一个函数
```
往程序中加入pdb.set_trace() 在程序中开启pdb模式

# 2. pep8编码规范

# 3. fork创建进程（类Linux才能用）
```python
import os

ret = os.fork() # 主进程的ret>0 子进程的ret=0
if ret == 0:
    while True:
        print("子进程",os.getpid(),os.getppid())
else :
    while True:
        print("父进程",os.getpid())

```
```python
#修改全局变量
import os 
import time

g_num = 100
ret = os.fork()
if ret ==0:
    print("子进程")
    g_num += 1
else:
    time.sleep(3)
    print("父进程",g_num) # 100

# 每个进程的数据不会共享
```
多次fork问题
```python
import os 
import time

ret = os.fork()
if ret == 0:
    print(1)
else:
    print(2)

ret = os.fork()
if ret == 0:
    print(11)
else:
    print(22)

# 共有四个进程
```
# 4. 跨平台版fork

multiprocessing 中的Process类,但是跟fork不同的是，它的父进程必须等待所有子进程都结束后再推出
```python
from multiprocessing import Process
import os


def run_proc(name):
    print("子进程在运行中,name=%s,pid=%d" % (name, os.getpid()))


if __name__ == "__main__":
    print("父进程%d" % os.getpid())
    p = Process(target=run_proc, args=("test",))
    p.start() #开始执行子进程
    p.join([等待时间]) #阻塞，等子进程执行完再往下执行

```
继承Process类
```python
from multiprocessing import Process
class ProcessClass(Process):
    def __init__(self):
        Process.__init__(self)

    def run(self):
        print("子进程运行了")

p = ProcessClass()
p.start()

```
进程池（常用）
```python
from multiprocessing import Pool

def worker(msg): 
    pass

po = Pool(3) # 最大容量
po.apply_async(worker,(i,))# 添加进程 异步添加 po.apply(worker,(i,)) 同步阻塞方式
po.close()# 关闭进程池
po.join()
```

# 5. 进程间通信-queue

Process  and  Queue
```python
from mylitprocessing import Queue,Process

q  = Queue([最大容量]) # 先进先出
q.empty()# 是否空
q.full()#是否满

q.qsize()
q.put() q.put_nowait() # 非阻塞
q.get() q.get_nowait() #非阻塞

p1 = Process(target=w1,(q,))
p2 = Process(target=w2,(q,))

p1.start()
p1.join()
p2.start()
p2.join()
```
Pool  and  Manager
```python
from multiprocess import Manager,Pool

q = Manager.Queue()
po.Pool()
po.apply(writer,(q,)) #同步阻塞
po.apply(reader,(q,)) #
po.close()
po.join()
```

多任务拷贝文件案例
```python
# 创建一个文件夹
# os.mkdir
# 获取 old文件夹中所有文件的名字
# os.listdir(oldFolderName)
# 使用多任务copy
from multiprocessing import Pool, Manager
import os


def copyFileTask(name, oldFolderName, newFolderName, queue):
    queue.put(name)
    fread = open("./" + oldFolderName + "/" + name)
    fwrite = open("./" + newFolderName + "/" + name, "w")

    content = fread.read()
    fwrite.write(content)
    fread.close()
    fwrite.close()


def main():
    oldFolderName = input("请输入文件夹名字;")

    newFolderName = oldFolderName + "-附件"

    os.mkdir(newFolderName)  # 创建文件夹

    fileNames = os.listdir(oldFolderName)  # 得到每一个文件夹

    queue = Manager().Queue()  # 创建队列
    pool = Pool(5)
    for name in fileNames:
        pool.apply_async(copyFileTask, args=(name, oldFolderName, newFolderName, queue))

    num = 0
    allNum = len(fileNames)
    while num < allNum:
        queue.get()
        num += 1
        copyRate = num / allNum * 100
        print("copy的进度是:%.2f%%" % copyRate, end="")

    pool.close()
    pool.join()


if __name__ == "__main__":
    main()

```

# 6. 多线程-threading

一个进程下开多个线程，一个进程运行起来就会有一个箭头（主线程）


类似Process类 threading.Thread(target=func,args)
```python
import threading
import time 
def say():
    print("吃饭")
    time.sleep(1)

if __name__ == "__main__":
    for i in range(5):
        t = threading.Thread(target=say)
        t.start()
```
同样类似class方法
```python
import threading
import time


class myThread(threading.Thread):
    def run(self):
        for i in range(3):
            time.sleep(1)
            print(self.name)
#self.name是子线程

if __name__ == "__main__":
    t = myThread()
    t.start()
```
主线程会等待子进程执行完，不然会产生僵尸进程

多线程共享全局变量，多进程不会共享数据

但要保证数据不会被同时使用

1. 轮询 flag=1 另一个flag=0
2. 互斥锁
```python
#创建锁
mutex = threading.Lock()
#上锁
mutex.acquire()#同时上锁，另一个上不了锁会等待
#释放锁
mutex.release()#通知的方式
```

threading.current_thread().name #Thread-1

# 7. 多线程跟多进程的区别

多线程一定会等待子线程执行完

多线程：共享全局变量，不共享局部变量
多进程：不存在共享，除非用queue

死锁及其解决方法：

1. 设置超时时间mutex.acquire(timeout=2)  **看门狗**
2. 银行家算法

# 8. 缓冲的地方

1. from queue import Queue **队列缓冲**
qsize() put() get()
2. 文件存储

Threadlocal使用全局字典
```python
import threading
global_dict = {}

class myThread(threading.Thread):
    global global_dict
    global_dict[threading.Thread.current_thread()] = name
```

使用threading.local()

避免了传参
```python
import threading

local_school = threading.local()

def process_student():
    std = local_school.student
    print("hello,%s %s"%(std,threading.current_thread().name))

def process_thread(name):
    local_school.student = name
    process_student()

t1 = threading.Thread(target=process_thread,args=("jiangshiyi",),name="Thread-A")
t2 = threading.Thread(target=process_thread,args=("laowang",),name="Thread-B")
t1.start()
t2.start()
t1.join()
t2.join()
```

GIL（上锁）导致多线程 是假的 ，但是关键的地方用C语言来写可以解决

所以多进程效率比多线程高