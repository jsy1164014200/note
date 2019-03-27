# erlang趣学指南之并发编程

> $ erlc road.erl
> $ erl -noshell -run road main road.txt

## 1. 并发概念

Erlang 的并发是基于消息传递和 actor 模型的

 3 个原语是创建（spawn）进程、发送消息以及接收消息

1. 创建进程

```erlang
1> F = fun() -> 2 + 2 end.
#Fun<erl_eval.20.67289768>
2> spawn(F).
<0.44.0>  %% pid


4> G = fun(X) -> timer:sleep(10), io:format("~p~n", [X]) end.
#Fun<erl_eval.6.13229925>
5> [spawn(fun() -> G(X) end) || X <- lists:seq(1,10)].
[<0.273.0>,<0.274.0>,<0.275.0>,<0.276.0>,<0.277.0>,
<0.278.0>,<0.279.0>,<0.280.0>,<0.281.0>,<0.282.0>]
2 1 4 3 5 8 7 6
10
9
```

为了证明 shell 本身也是一个常规的进程，可以使用 BIF self/0，这个函数会返回当前进程
的 pid：
6> self().
<0.41.0>
7> exit(self()).
** exception exit: <0.41.0>
8> self().
<0.285.0>

可以看到，因为进程被重启了，所以 pid 变了

2. 发送消息

现在我们来介绍消息传递原语—操作符!，也称为 bang 符号。该操作符的左边是一个 pid，右边可以是任意 Erlang 数据项。这个数据项会被发送给左边的 pid 所代表的进程，这个进程就可以访问它了。下面是一个例子：
9> self() ! hello.
hello
消息会被放到接收进程的邮箱中，但是并没有被读取。上面例子中出现的第二个 hello 是这个发送函数的返回值。这意味着，可以用如下方式给多个进程发送同样的消息：
10> self() ! self() ! double.
double
上面的调用等价于 self() ! (self() ! double)。

要想看到当前邮箱中的内容，如果在 shell
中，可以使用 flush()命令：
11> flush().
Shell got hello
Shell got double
Shell got double
ok

3. 接收消息

receive
Pattern1 when Guard1 -> Expr1;
Pattern2 when Guard2 -> Expr2;
Pattern3 -> Expr3
end

类似case of 语句

```erlang
-module(dolphins).
-compile(export_all).
dolphin1() ->
receive
do_a_flip ->
io:format("How about no?~n");
fish ->
io:format("So long and thanks for all the fish!~n");
_ ->
io:format("Heh, we're smarter than you humans.~n")
end.


12> Dolphin = spawn(dolphins, dolphin1, []).
<0.40.0>
13> Dolphin ! "oh, hello dolphin!".
Heh, we're smarter than you humans.
"oh, hello dolphin!"
14> Dolphin ! fish
fish
```



## 2. 深入多重处理

一个简单的 冰箱 存取操作

```erlang
-module(kitchen).
-compile(export_all).

fridgel(FoodList) -> 
    receive
        {From, {store, Food}} ->
            From ! {self(), ok},
            fridgel([Food|FoodList]);
        {From, {take, Food}} ->
            case lists:member(Food, FoodList) of 
                true -> 
                    From ! {self(), {ok, Food}},
                    fridgel(lists:delete(Food, FoodList));
                false -> 
                    From ! {self(), not_found},
                    fridgel(FoodList)
end;
        terminate -> 
            ok
end.

store(Pid, Food) -> 
    Pid ! {self(), {store, Food}},
    receive
        {Pid, Msg} -> Msg
end.

take(Pid, Food) -> 
    Pid ! {self(), {take, Food}},
    receive
        {Pid, Msg} -> Msg
end.

start(FoodList) -> 
    spawn(?MODULE, fridgel, [FoodList]).
```



#### 加入超时时间

receive 和 after 之间的部分和之前见到的完全一样。当过了 Delay（单位是毫秒）时间
后还没有收到和 Match 模式相匹配的消息，就会执行 after 部分。此时，会执行 Expresson2

```erlang
store(Pid, Food) -> 
    Pid ! {self(), {store, Food}},
    receive
        {Pid, Msg} -> Msg
    after 3000 -> 
        timeout
end.

take(Pid, Food) -> 
    Pid ! {self(), {take, Food}},
    receive
        {Pid, Msg} -> Msg
    after 3000 -> 
        timeout
end.
```

> after 只接收毫秒值，但是实际上还可以使用原子 infinity。

#### 选择性接受



```erlang
important() -> 
    receive
        {Priority, Message} when Priority > 10 -> 
            [Message | important()]
    after 0 -> 
        normal()
end.

normal() -> 
    receive
        {_, Message} -> 
            [Message | normal()]
    after 0 -> 
        []
end.
```

```erlang
1> c(multiproc).
{ok,multiproc}
2> self() ! {15, high}, self() ! {7, low}, self() ! {1, low}, self() ! {17, high}.
{17,high}
3> multiproc:important().
[high,high,low,low]
```

> 消息会存在 消息邮箱中



## 3. 错误与进程

### 链接

链接（link）是两个进程之间的一种特殊关系。当在两个进程间建立了这种关系后，如果其中一个进程由于意外的抛出、出错或者退出（参见第 7 章）而死亡时，另外一个进程也会死亡，把这两个进程独立的生存期绑定成一个关联在一起的生存期。

Erlang 中有一个原生函数 link/1，用于在两个进程间建立一条链接，它的参数是进程的 pid。当调用它时，会在当前进程和参数 pid 标识的进程之间建立一条链接。要去除链接，可以使用unlink/1。

链接在一起的进程组

```erlang
chain(0) ->
    receive
        _ -> ok
    after 2000 ->
        exit("chain dies here")
    end;
chain(N) ->
    Pid = spawn(fun() -> chain(N-1) end),
    link(Pid),
    receive
    	_ -> ok
    end.
```



捕获退出进程

```erlang
1> process_flag(trap_exit, true).
true
2> spawn_link(fun() -> linkmon:chain(3) end).
<0.49.0>
3> receive X -> X end.
{'EXIT',<0.49.0>,"chain dies here"}
```

啊！现在有趣多了。还用图示来说明的话，现在看起来是这样：
[shell] == [3] == [2] == [1] == [0][shell] == [3] == [2] == [1] == *dead*
[shell] == [3] == [2] == *dead*
[shell] == [3] == *dead*
[shell] <-- {'EXIT,Pid,"chain dies here"} -- *dead*
[shell] <-- still alive!



### 异常（新概念）

异常源： spawn_link(fun() -> ok end)
未捕获时的结果： 无任何结果
捕获时的结果： {'EXIT', <0.61.0>, normal}
进程正常退出，没有出现问题。注意，这个结果和 catch exit(normal)的结果有点儿像，只不过元组中多了一个 pid，标识哪个进程死亡了。

异常源： spawn_link(fun() -> exit(reason) end)
未捕获时的结果： ** exception exit: reason
捕获时的结果： {'EXIT', <0.55.0>, reason}
进程由于一个自定义的原因终止了。如果没有捕获退出信号，和它链接的进程就会崩溃。如果捕获了退出信号，就会收到一条消息。

异常源： spawn_link(fun() -> exit(normal) end)
未捕获时的结果： 无任何结果
捕获时的结果： {'EXIT', <0.58.0>, normal}这个调用成功地模拟了进程正常退出的情况。有时，你想把杀死进程当作程序正常流程的一部分，并且不想引发任何异常，那么可以使用这种方法

例子：

异常源： spawn_link(fun() -> 1/0 end)
未捕获时的结果： Error in process <0.44.0> with exit value: {badarith,[{erlang, '/', [1,0]}]}
捕获时的结果： {'EXIT', <0.52.0>, {badarith, [{erlang, '/', [1,0]}]}}
错误（{badarith, Reason}）没有用 try ... catch 块进行捕获，从而升级成了一个'EXIT'。此时，它和 exit(reason)完全一样，不过多了一个描述所发生情况细节的栈跟踪信息。

异常源： spawn_link(fun() -> erlang:error(reason) end)
未捕获时的结果： Error in process <0.47.0> with exit value: {reason,[{erlang, apply, 2}]}
捕获时的结果： {'EXIT', <0.74.0>, {reason, [{erlang, apply, 2}]}}这里的错误情况和 1/0 非常相似。这很正常—erlang:error/1 就是用来做这种事的。

异常源： spawn_link(fun() -> throw(rocks) end)
未捕获时的结果： Error in process <0.51.0> with exit value: {{nocatch,rocks},[{erlang, apply, 2}]}
捕获时的结果： {'EXIT', <0.79.0>, {{nocatch, rocks}, [{erlang, apply,2}]}}
因为没有用 try ... catch 去捕获 throw，所以它先升级成一个错误，随后又升级成一个 EXIT。没有捕获退出信号时，进程死亡了；当捕获退出信号时，则可以很好地处理这个错误。



完全不同的exit/2

异常源： exit(self(), normal)
未捕获时的结果： ** exception exit: normal
捕获时的结果： {'EXIT', <0.31.0>, normal}
当没有捕获退出信号时， exit(self(), normal)和 exit(normal)一样。否则，会收到一条消息，消息的格式和有链接关系的进程正常死亡时该进程收到的消息格式一样。

异常源： exit(spawn_link(fun() -> timer:sleep(50000) end), normal)
未捕获时的结果： 无任何结果
捕获时的结果： 无任何结果  这基本上等同于 exit(Pid, normal)调用。这个命令不会有任何效果，因为不能以
normal 作为原因参数远程杀死一个进程。

异常源： exit(spawn_link(fun() -> timer:sleep(50000) end), reason)
未捕获时的结果： ** exception exit: reason
捕获时的结果： {'EXIT', <0.52.0>, reason}
这个调用会导致链接着的另外一方进程因为 reason 原因终止。结果等同于另外一方进
程自己调用 exit(reason)。

异常源： exit(spawn_link(fun() -> timer:sleep(50000) end), kill)
未捕获时的结果： ** exception exit: killed
捕获时的结果： {'EXIT', <0.58.0>, killed}
很奇怪，死亡进程发送给创建者进程的消息有所变化。现在创建者进程收到的消息是
killed 而不是 kill。这是因为 kill 是一个特殊的退出信号，会在下一节进行介绍。

异常源： exit(self(), kill)
未捕获时的结果： ** exception exit: killed
捕获时的结果： ** exception exit: killed
哇，看看。这个调用似乎根本不可能被捕获。下面的异常也很厉害。

异常源： spawn_link(fun() -> exit(kill) end)
未捕获时的结果： ** exception exit: killed
捕获时的结果： {'EXIT', <0.67.0>, kill}
这就有些令人困惑了。当进程用 exit(kill)杀死自己时，如果我们的进程没有捕获退出信号，它死亡的原因是 killed。但是，当捕获了退出信号时，收到的消息中的原因却是 kill。

#### kill

kill 的含义

虽然进程可以捕获退出信号，但是有时还是想强行杀死（kill）这个进程。也许某个捕获了退出信号的进程陷入了一个死循环，不能再接收任何消息了。 kill 原因是一种特殊的信号，它不能被捕获。这可以保证任何被它终止的进程都肯定会死掉。通常，当其他所有手段都无效时，可以使用 kill。由于 kill 原因无法被捕获，因此当其他进程接收到这个消息时需要被更改为 killed。如果不进行更改，那么所有和它链接在一起的其他进程都会因为同样的 kill 原因相继死亡，并会导致和这些进程具有链接关系的其他进程死亡，从而发生死亡的连锁反应。这也解释了为何在没有捕获退出信号时，链接着的进程如果调用 exit(kill)，本进程收到的是 killed（信号被更改，防止连锁反应），但是在本地捕获了退出信号时，看起来仍然是 kill的原因。如果你感到困惑，不要担心。很多程序员都有这个感受。退出信号是有些奇怪复杂。 不过，除了这里介绍的之外，再没有其他特别的情况了。一旦你理解了这里的内容，基本上就可以理解Erlang 的所有并发错误管理了。



### 监控器

也许杀死进程并不是你想要的。也许你并不希望自己离开后整个世界都随你而去。也许你只想当一个跟踪者。如果是这样，那么监控器（monitor）可能就是你想要的，因为它们不会杀死进程。监控器是一种特殊类型的链接，有如下两点不同：

- 监控器是单向的；

- 在两个进程之间可以设置多个监控器（监控器可以叠加，每个监控器有
  自己的标识）。

  如果一个进程想知道另外一个进程的死活，但是这两个进程之间并没有强
  的业务关联时，可以使用监控器。

创建监控器的函数是erlang:monitor/2，
它的第一个参数永远是原子 process，第二个参数是进程的 pid。
1> erlang:monitor(process, spawn(fun() -> timer:sleep(500) end)).
#Ref<0.0.0.77>
2> flush().
Shell got {'DOWN',#Ref<0.0.0.77>,process,<0.63.0>,normal}
ok
每 当 被 监 控 的 进 程 死 亡 时 ， 监 控 进 程 都 会 收 到 一 条 消 息 ， 格 式 是 {'DOWN',
MonitorReference, process, Pid, Reason}。其中的引用可以用来解除对一个进程的
监控。记住，监控器是可叠加的，因此会收到多条 DOWN 消息。引用可以唯一确定一条 DOWN 消息。还要说明一点，和链接一样，监控器也有一个具有原子性质的函数，可以在创建进程的同时
监控它： spawn_monitor/1-3。
3> {Pid, Ref} = spawn_monitor(fun() -> receive _ -> exit(boom) end end).
{<0.73.0>,#Ref<0.0.0.100>}
4> erlang:demonitor(Ref).
true
5> Pid ! die.
die
6> flush().
ok



7> f().
ok
8> {Pid, Ref} = spawn_monitor(fun() -> receive _ -> exit(boom) end end).
{<0.35.0>,#Ref<0.0.0.35>}
9> Pid ! die.
die
10> erlang:demonitor(Ref, [flush, info]).
false
11> flush().
ok
info 选项用来指示某个监控器在被解除时是否还存在。这也是为何在第 10 行的调用中返回
了 false。 flush 选项会把邮箱中存在的 DOWN 消息都清除掉，因此 shell 中的 flush()命令
没有从 shell 进程的邮箱中找到任何消息。

### 命名进程

可以使用函数 erlang:register(Name,Pid)为进程命名。如果进程死亡了，它会自动，失去自己的名字。也可以使用函数 unregister/1 手工解除进程的名字注册。

```erlang
restarter() ->
process_flag(trap_exit, true),
Pid = spawn_link(?MODULE, critic, []),
register(critic, Pid),
receive
{'EXIT', Pid, normal} -> %正常死亡
ok;
{'EXIT', Pid, shutdown} -> %手动终止，不是崩溃
ok;
{'EXIT', Pid, _} ->
restarter()
end
    
    
    judge2(Band, Album) ->
Ref = make_ref(),
critic ! {self(), Ref, {Band, Album}},
receive
{Ref, Criticism} -> Criticism
after 2000 ->
timeout
end.
critic2() ->
receive
{From, Ref, {"Rage Against the Turing Machine", "Unit Testify"}} ->
From ! {Ref, "They are great!"};
{From, Ref, {"System of a Downtime", "Memoize"}} ->
From ! {Ref, "They're not Johnny Crash but they're good."};
{From, Ref, {"Johnny Crash", "The Token Ring of Fire"}} ->
From ! {Ref, "Simply incredible."};
{From, Ref, {_Band, _Album}} ->
From ! {Ref, "They are terrible!"}
end,
critic2().
```

## 4. 一个小的提醒应用

#### 需求

增加一个事件。事件中包含最后期限（给出警告的时间）、事件名称以及事件描述。
当事件的最后期限到达时，显示警告。
根据名字取消事件。
通过命令行和系统交互，也可以扩展为其他方式（如 GUI、 Web 页面、即时消息软件或
者电子邮件）。

#### 分析

事件服务器的任务如下：
 接收来自客户进程的订阅；
 把来自事件进程的消息转发给每个订阅者进程；
 接收增加事件的消息（需要时会启动 x、 y 和 z 进程）；
 接收取消事件的消息，随后杀死事件进程。
事件服务器可以被客户进程终止，事件服务器的代码可以通过 shell 重新加载。
客户进程的任务如下：
 向事件服务器发起订阅，并接收通知消息；
 请求服务器增加一个包含具体内容的事件；
 请求服务器取消一个事件；
 监控服务器（为了知道服务器进程是否死亡）；
 在需要时，关闭事件服务器。

![选区_129](/home/jsy/图片/选区_129.png)



在真实的应用中，把每个待提醒的事件都表示为一个进程的做法可能有些过度了，并且难以
扩展到大量事件的场合。不过，因为这个应用只有你一个用户，所以这种设计没有问题。还可以
使用如 timer:send_after/2-3 之类的函数避免创建过多的进程。



#### 目录结构

1. ebin/
2. include/
3. priv/
4. src/
5. conf/
6. doc/
7. lib/
8. deps/

 ebin/目录存放编译过的源代码；
 include/目录存放可以被其他应用包含的.hrl 文件（私有的.hrl 文件通常放在 src/目录中）；
 priv/目录存放需要和 Erlang 交互的可执行文件，如特定的驱动库等。在本项目中并没有
用到这个目录；
 src/目录存放所有的.erl 文件。





Erlang 中函数有本地和外部两种调用方式，这是一个很重要的概念。本地调用指的是那
些可以针对非导出函数进行的调用。本地调用的格式为 Name(Args)。外部调用只能用于导
出函数，格式为 Module:Function(Args)。外部调用更准确的名字是完全限定调用（ fully
qualified call）

当 VM 中加载了一个模块的两个版本时，所有本地调用都针对进程中当前运行的模块版本。
不过，完全限定调用则是永远针对代码服务器中模块的最新版本。因此，如果在完全限定调用中
进行了本地调用，那么这些本地调用使用的都是最新版本的代码。

![选区_130](/home/jsy/图片/选区_130.png)





> erl -make
>
> erl -pa ebin/





```erlang

```











