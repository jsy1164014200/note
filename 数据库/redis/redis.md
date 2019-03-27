

redis 没有简单类型，只有 string hash list set zset

发布订阅的设计模式 ： 不需要主动请求数据，订阅一个频道，会自动推送

# Redis in Action 笔记

## 1. 基础数据类型

![选区_106](/home/jsy/图片/选区_106.png)

### STRING  字符串，整数，浮点

get 

set

del 

![选区_110](/home/jsy/图片/选区_110.png)

对于空字符串或者是，不存在的键执行加减，会默认当做0来算。

![选区_111](/home/jsy/图片/选区_111.png)



### LIST 列表

lpush,rpush,lpop,rpop,

lindex,lrange

![选区_107](/home/jsy/图片/选区_107.png)

![选区_112](/home/jsy/图片/选区_112.png)

lrange  支持负数做切片

例如要返回数组的所有元素就可以用 lrange list 0 -1

![选区_113](/home/jsy/图片/选区_113.png)



### SET 无序集合

> 可以看做是一个 分数为 1 的zset

sadd  添加

srem  移除

smembers 显示所有包含的元素

sismember  是否包含（如何集合包含的元素比较多，可能查询的比较慢） 

sinter  交

sunion 并集

sdiff 差集

![选区_114](/home/jsy/图片/选区_114.png)

sdiff

![选区_115](/home/jsy/图片/选区_115.png)

### HASH



![选区_108](/home/jsy/图片/选区_108.png)

hset

hget

hgetall

hdel 

![选区_116](/home/jsy/图片/选区_116.png)

![选区_117](/home/jsy/图片/选区_117.png)

同string 一样，如果对散列中不存在的值进行加减 也会默认当做0



### ZSET

有序集合是redis里面唯一一个能够根据成员访问元素，又可以根据分值以及分值排列顺序来访问元素的结构。

zadd  name  score value

zrange  name start stop withscores

zrangebyscore  name  startscore stoopscore  withscores

zrem  name x y withscores

![选区_118](/home/jsy/图片/选区_118.png)

![选区_119](/home/jsy/图片/选区_119.png)

![选区_120](/home/jsy/图片/选区_120.png)

![选区_121](/home/jsy/图片/选区_121.png)

![选区_122](/home/jsy/图片/选区_122.png)



## 2. 其他命令 - sort multi exec 自动过期

#### sort

![选区_125](/home/jsy/图片/选区_125.png)

#### 事务（不同于 sql 的事务）

是指一系列操作都执行完毕后，服务器才会执行其他客户端命令。

> 对于python ，就是在一个pipeline() 方法中，就会创建一个事务

![选区_126](/home/jsy/图片/选区_126.png)

使用事务的好处：

1. 能够移除竞争条件
2. 提高性能，降低与服务器之间的通讯次数

#### 键的过期时间

> redis 没有办法对键里面的元素设置过期时间

![选区_127](/home/jsy/图片/选区_127.png)



## 3. 数据安全与性能保障

### 1. 持久化

> 快照(snapshotting) , 只追加文件(append-only file) 

























# 附录

## redis 发布与订阅



![选区_123](/home/jsy/图片/选区_123.png)

![选区_124](/home/jsy/图片/选区_124.png)













redis 并不支持嵌套结构，可以使用键名来模拟嵌套结构：

user:123 表示 储存用户信息的散列，user:123:posts 表示发表文章的有序集合。

del , type , rename 通用的命令









