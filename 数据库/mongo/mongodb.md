# mongodb权威指南笔记



## mongo 应用程序设计范式

• 内嵌数据和引用数据之间的权衡；
• 优化技巧；
• 数据一致性；
• 模式迁移；
• 不适合使用 MongoDB 作为数据存储的场景。



### 范式与反范式

范式化（normalization）是将数据分散到多个不同的集合，不同集合之间可
以相互引用数据。虽然很多文档可以引用某一块数据，但是这块数据只存储在一个
集合中。所以，如果要修改这块数据，只需修改保存这块数据的那一个文档就行了。
但是， MongoDB 没有提供连接（join）工具，所以在不同集合之间执行连接查询需
要进行多次查询。

反范式化（denormalization）与范式化相反：将每个文档所需的数据都嵌入在文档
内部。每个文档都拥有自己的数据副本，而不是所以文档共同引用同一个数据副本。
这意味着，如果信息发生了变化，那么所有相关文档都需要进行更新，但是在执行
查询时，只需要一次查询，就可以得到所有数据

一个典型的反范式化的例子。

{
"_id" : ObjectId("512512a5d86041c7dca81914"),
"name" : "John Doe",
"classes" : [
{
"class" : "Trigonometry",
"credits" : 3,
"room" : "204"
},
{
"class" : "Physics",
"credits" : 3,
"room" : "159"
},
{
"class" : "Women in Literature",
"credits" : 3,
"room" : "14b"
},
{
"class" : "AP European History",
"credits" : 4,
"room" : "321"
}
]
}

一般来说，数据生成越频繁，就越不应该将这些数据内嵌到其他文档中。如果内嵌
字段或者内嵌字段数量是无限增长的，那么应该将这些内容保存在单独的集合中，
使用引用的方式进行访问，而不是内嵌到其他文档中。评论列表或者活动列表等信
息应该保存在单独的集合中，不应该内嵌到其他文档中。

最后，如果某些字段是文档数据的一部分，那么需要将这些字段内嵌到文档中。如
果在查询文档时经常需要将某个字段排除，那么这个字段应该放在另外的集合中，
而不是内嵌在当前的文档中。

![选区_109](/home/jsy/图片/选区_109.png)



### 基数与集合的关系

一个集合中包含的对其他集合的引用数量叫做基数（cardinality）。常见的关系有一
对一、一对多、多对多。假如有一个博客应用程序。每篇博客文章（post）都有一
个标题（title），这是一个一对一的关系。每个作者（author）可以有多篇文章，这
是一个一对多的关系。每篇文章可以有多个标签（tag），每个标签可以在多篇文章
中使用，所以这是一个多对多的关系。
在 MongoDB 中， many（多）可以被分拆为两个子分类： many（多）和 few（少）。
假如，作者和文章之间可能是一对少的关系：每个作者只发表了为数不多的几篇文
章。博客文章和标签可能是多对少的关系：文章数量实际上很可能比标签数量多。
博客文章和评论之间是一对多的关系：每篇文章都可以拥有很多条评论。

只要确定了少与多的关系，就可以比较容易地在内嵌数据和引用数据之间进行权衡。
通常来说，“少”的关系使用内嵌的方式会比较好，“多”的关系使用引用的方式比
较好。

综上所述：

1. 一对一
2. 一对多 ，引用
3. 一对少 ， 内嵌
4. 少对少 （少的内嵌，多的引用）
5. 少对多 （少的内嵌，多的引用）
6. 多对多 （少的内嵌，多的引用）



### 优化数据操作

1. 优化文档增长

如果增长程度是可预知的，可以为文档预留足够的增长空间，这样可以避免文档移动，可以提高写入速度。检查一下填充因子：如果它大约是 1.2 或者更大，可以考虑手动填充。

{
"_id" : ObjectId(),
"restaurant" : "Le Cirque",
"review" : "Hamburgers were overpriced."
"userId" : ObjectId(),
"tags" : [],
"garbage" : "......................................................"+
".............................................................."+
".............................................................."
}如果文档中有一个字段需要增长，应该尽可能将这个字段放在文档最后的位置（ "garbage" 之前）。这样可以稍微提高一点点的性能，因为如果 "tags" 字段发生
了增长， MongoDB 不需要重写 "tags" 后面的字段。

### 删除旧数据

有三种常见的方式用于删除旧数据：使用固定集合，使用 TTL 集合，或者
定期删除集合。

使用 Ruby、 Python 和 Java 驱动程序时尤其要注意这个问题，因为这三种语言的驱
动程序都使用了连接池（connection pool）。为了提高效率，这些驱动程序会建立多
个与服务器之间的连接（也就是一个连接池），将请求通过不同的连接发送到服务
器。但是它们都有各自的机制来保证一系列相关的请求会被同一个连接处理。



个人博客设计

如果数组已经存在， "$push" 会向已有的数组末尾加入一个元素，要是没有就
创建一个新的数组。例如，假设要存储博客文章，要添加一个用于保存数组的
"comments"（评论）键。可以向还不存在的 "comments" 数组添加一条评论，这
个数组会被自动创建，并加入一条评论



注意，通常不需要保留 "createdAt" 这样的字段，因为 ObjectIds 里包含了一个
用于标明文档创建时间的时间戳。但是，在预置或者初始化计数器时，或者是对于
不使用 ObjectIds 的集合来说， "$setOnInsert" 是非常有用的。

常见的分页操作

db.stock.find({"desc" : "mp3"}).limit(50).skip(50).sort({"price" : -1})

## 1. mongodb介绍

MongoDB 是一款强大、灵活，且易于扩展的通用型数据库。它能扩展出非常多
的功能，如二级索引（secondary index）、范围查询（range query）、排序、聚合
（aggregation），以及地理空间索引（geospatial index）。

### 对比sql

1. 与关系型数据库相比，面向文档的数据库不再有“行”（row）的概念，取而代之的
   是更为灵活的“文档”（document）模型。通过在文档中嵌入文档和数组，面向文
   档的方法能够仅使用一条记录来表现复杂的层次关系，这与使用现代面向对象语言
   的开发者对数据的看法一致。
2. 再有预定义模式（predefined schema）：文档的键（key）和值（value）不
   再是固定的类型和大小。由于没有固定的模式，根据需要添加或删除字段变得更容
   易了。
3. MongoDB 的设计采用横向扩展。面向文档的数据模型使它能很容易地在多台服务
   器之间进行数据分割。 MongoDB 能自动处理跨集群的数据和负载，自动重新分配
   文档，以及将用户请求路由到正确的机器上。这样，开发者能够集中精力编写应用
   程序，而不需要考虑如何扩展的问题。如果一个集群需要更大的容量，只需要向集
   群添加新服务器， MongoDB 就会自动将现有数据向新服务器传送。



### mongo功能的特点

1. 索引（ indexing）
   MongoDB 支持通用二级索引，允许多种快速查询，且提供唯一索引、复合索引、地理空间索引，以及全文索引。
2. MongoDB 支持“聚合管道”（aggregation pipeline）。用户能通过简单的片段创建复杂的聚合，并通过数据库自动优化。
3. 特殊的集合类型
   MongoDB 支持存在时间有限的集合，适用于那些将在某个时刻过期的数据
4.  文件存储（ file storage）
      MongoDB 支持一种非常易用的协议，用于存储大文件和文件元数据

### 基本概念

- 文档是 MongoDB 中数据的基本单元，非常类似于关系型数据库管理系统中的行，但更具表现力。
- 类似地， 集合（collection）可以看作是一个拥有动态模式（dynamic schema）的表。
-  MongoDB 的一个实例可以拥有多个相互独立的数据库（database），每一个数据库都拥有自己的集合。
- 每一个文档都有一个特殊的键 "_id"， 这个键在文档所属的集合中是唯一的
- MongoDB 自带了一个简单但功能强大的 JavaScript shell，可用于管理 MongoDB的实例或数据操作。



文档：

1. 键不能含有 \0（空字符）。这个字符用于表示键的结尾。
2. . 和 $ 具有特殊意义，只能在特定环境下使用
3. 文档中的键 / 值对是有序的： ｛"x" : 1, "y"： 2｝ 与｛"y": 2, "x": 1｝ 是不同
   的。通常，字段顺序并不重要，无须让数据库模式依赖特定的字段顺序（MongoDB会对字段重新排序）。在某些特殊情况下，字段顺序变得非常重要

集合

1. 集合就是一组文档。如果将 MongoDB 中的一个文档比喻为关系型数据库中的一行，那么一个集合就相当于一张表。
2. 动态模式（一个集合里面的文档是各式各样的）
3. 要将不同的文档 分开在不同的集合中存储

子集合

组织集合的一种惯例是使用“.”分隔不同命名空间的子集合。例如，一个具有博客
功能的应用可能包含两个集合，分别是 blog.posts 和 blog.authors。这是为了使组织
结构更清晰，这里的 blog 集合（这个集合甚至不需要存在）跟它的子集合没有任何
关系。
虽然子集合没有任何特别的属性，但它们却非常有用，因而很多 MongoDB 工具都
使用了子集合

- GridFS（一种用于存储大文件的协议）使用子集合来存储文件的元数据，这样就
  可以与文件内容块很好地隔离开来。
- 大多数驱动程序都提供了一些语法糖，用于访问指定集合的子集合。例如，在数
  据库 shell 中， db.blog 代表 blog 集合，而 db.blog.posts 代表 blog.posts 集合



数据库

1. 要记住一点，数据库最终会变成文件系统里的文件，而数据库名就是相应的文件名
2. 有一些数据库是保留的
   - admin
     从身份验证的角度来讲，这是“root”数据库。如果将一个用户添加到 admin 数据库，这个用户将自动获得所有数据库的权限。再者，一些特定的服务器端命令也只能从 admin 数据库运行，如列出所有数据库或关闭服务器
   - local
     这个数据库永远都不可以复制，且一台服务器上的所有本地集合都可以存储在这个数据库中
   - config
     MongoDB 用于分片设置时，分片信息会存储在 config 数据库中。



命名空间

把数据库名添加到集合名前，得到集合的完全限定名，即命名空间（namespace）。
例如，如果要使用 cms 数据库中的 blog.posts 集合，这个集合的命名空间就是 cms.
blog.posts。命名空间的长度不得超过 121 字节，且在实际使用中应小于 100 字节

### mongod服务器

1. mongod 在没有参数的情况下会使用默认数据目录 /data/db，启动
   MongoDB 前，先创建数据目录（如 mkdir -p /data/db/），以确保对该目录有写权限，这点非常重要。
2. 监听 27017 端口



### mongodb shell 是一个JavaScript shell

help 命令可以查看 自带的帮助文档

可以通过 db.help() 查看数据库级别的帮助，使用 db.foo.help() 查看集合级别
的帮助。

如果想知道一个函数是做什么用的，可以直接在 shell 输入函数名（函数名后不要
输入小括号），这样就可以看到相应函数的 JavaScript 实现代码

$ mongo --nodb

> 启 动 之 后， 在 需 要 时 运 行 new Mongo(hostname) 命 令 就 可 以 连 接 到 想 要 的
> mongod 了：
> 》conn = new Mongo("some-host:30000")
> connection to some-host:30000
> 》 db = conn.getDB("myDB")
> myDB
> 执行完这些命令之后，就可以像平常一样使用 db 了。任何时候都可以使用这些命
> 令来连接到不同的数据库或者服务器。



> 它可以运行所有的JavaScript命令
>
> 为了方便习惯使用 SQL shell 的用户， shell 还包含一些非 JavaScript 语法的扩展。这些扩展并不提供额外的功能，而是一些非常棒的语法糖。

连接后，db默认指向test上数据库

命令

1. db  显示当前数据库
2. use db 

创建

1. db.blog.insert({post:"xxx"})



读取

1. db.blog.find(查询文档)   自动显示最多 20 个匹配的文档
2. db.blog.findOne(查询文档)



更新

1. db.blog.update({title: "xxx"}, post)



删除

1. db.blog.remove({title: "xxx"})



### 设置一些操作脚本js

使用 load("/home/myUser/
my-scripts/defineConnectTo.js") 命 令 来 加 载 defineConnectTo.js。 注 意，
load 函数无法解析 ~ 符号。

```js
/**
* 连接到指定的数据库，并且将 db 指向这个连接
*/
var connectTo = function(port, dbname) {
if (!port) {
port = 27017;
}
if (!dbname) {
dbname = "test";
}
db = connect("localhost:"+port+"/"+dbname);
return db;
};
```

在用户主目录下创
建一个名为 .mongorc.js 的文件，可以在 进入mongoshell时自己运行

例如可以：

```js
var no = function() {
print("Not on my watch.");
};
// 禁止删除数据库
db.dropDatabase = DB.prototype.dropDatabase = no;
// 禁止删除集合
DBCollection.prototype.drop = no;
// 禁止删除索引
DBCollection.prototype.dropIndex = no;


另一个方便的提示是显示当前使用的数据库：
prompt = function() {
if (typeof db == 'undefined') {
return '(nodb)> ';
}
// 检查最后的数据库操作
try {
db.runCommand({getLastError:1});
}
catch (e) {
print(e);
}
return db+"> ";
};


现在，如果想要编辑一个变量，可以使用 "edit 变量名 " 这个命令，比如：
> var wap = db.books.findOne({title: "War and Peace"})
> edit wap
修改完成之后，保存并退出编辑器。变量就会被重新解析然后加载回 shell。
在 .mongorc.js 文件中添加一行内容， EDITOR=" 编辑器路径 ";，以后就不必单独设
置 EDITOR 变量了。
```







### 基本数据类型

- null

- true  或者 false

- 数值

shell 默认使用 64 位浮点型数值。因此，以下数值在 shell 中是很“正常”的：
{"x" : 3.14}或：{"x" : 3}
对于整型值，可使用 NumberInt 类（表示 4 字节带符号整数）或 NumberLong
类（表示 8 字符带符号整数），分别举例如下：{"x" : NumberInt("3")}

{"x" : NumberLong("3")}

- 字符串

- 日期被存储为自新纪元以来经过的毫秒数，不存储时区,{"x": new Date()}

- 正则表达式

查询时，使用正则表达式作为限定条件，语法也与 JavaScript 的正则表达式语法
相同，例如：{ "x" : /foobar/i }

- 数组   **JavaScript数组，可以存储不同类型**

- 内嵌文档

`{"x" : {"foo" : "bar"}}`

- Object id

Object id 是一个 12 字节的 ID，是文档的唯一标识。

![选区_104](/home/jsy/图片/选区_104.png)

与 _ id 区别开来：

MongoDB 中存储的文档必须有一个 "_ id" 键。这个键的值可以是任何类型的，默
认是个 ObjectId 对象。在一个集合里面，每个文档都有唯一的 "_ id"，确保集合里
面每个文档都能被唯一标识.



- 二进制数据

二进制数据是一个任意字节的字符串。它不能直接在 shell 中使用。如果要将非
UTF-8 字符保存到数据库中，二进制数据是唯一的方式

- 代码

`{"x" : function() { /* ... */ }}`



## 2. 数据操作

### 插入

+ `db.foo.insert({"bar" : "baz"})``

+ ``db.foo.batchInsert([{"_id" : 0}, {"_id" : 1}, {"_id" : 2}])` MongoDB 能接受的最大消息长度是 48 MB，所以在一次批量插入中能
  插入的文档是有限制的。如果试图插入 48 MB 以上的数据，多数驱动程序会将这个批量插入请求拆分为多个 48 MB 的批量插入请求。

  在批量插入中遇到错误时，如果希望 batchInsert 忽略错误并且继续执行后续插入，可以使用 continueOnError 选项。

### 删除

- `db.foo.remove()`
  上述命令会删除 foo 集合中的所有文档。但是不会删除集合本身，也不会删除集合的元信息
- db.foo.drop() 会删除整个集合

### 更新

简单的修改

```js
> joe = db.people.findOne({"name" : "joe", "age" : 20});
{
"_id" : ObjectId("4b2b9f67a1f631733d917a7c"),
"name" : "joe",
"age" : 20
}
> joe.age++;
> db.people.update({"name" : "joe"}, joe);
E11001 duplicate key on update
调用 update 时，数据库会查找一个 "name" 值为 "Joe" 的文档。找
到的第一个是 65 岁的 Joe。然后数据库试着用变量 joe 中的内容替换找到的文档，
但是会发现集合里面已经有一个具有同样 "_id" 的文档。所以，更新就会失败，因
为 "_id" 值必须唯一。为了避免这种情况，最好确保更新时总是指定一个唯一文
档，例如使用 "_id" 这样的键来匹配

db.people.update({"_id"：ObjectId("4b2b9f67a1f631733d917a7c")}, joe)
```

#### $inc 自增

"$inc" 只能用于整型、长整型或双精度浮点型的值。

db.analytics.update({"url" : "www.example.com"}, {"$inc" : {"pageviews" : 100}})

#### $set 添加 修改 

db.users.update({"_id" : ObjectId("4b253b067525f35f94b60a31")}，{"$set" : {"favorite book" : "War and Peace"}})

#### $set 修改子文档

db.blog.posts.update({"author.name" : "joe"},{"$set" : {"author.name" : "joe schmoe"}})

#### $unset 删除键

db.users.update({"name" : "joe"},{"$unset" : {"favorite book" : 1}})

#### $push 增加数组元素

db.blog.posts.update({"title" : "A blog post"},
... {"$push" : {"comments" :
... {"name" : "joe", "email" : "joe@example.com",
... "content" : "nice post."}}})

#### $each  + $push

db.stock.ticker.update({"_id" : "GOOG"},
... {"$push" : {"hourly" : {"$each" : [562.776, 562.790, 559.123]}}})

#### $slice + $each + $push
使用，这样就可以保证数组不会超出设定好的最大长度，这实际上就得到了一个最多包含 N 个元素的数组：

db.movies.find({"genre" : "horror"},
... {"$push" : {"top10" : {
... "$each" : ["Nightmare on Elm Street", "Saw"],
... "$slice" : -10}}})
这个例子会限制数组只包含最后加入的 10 个元素。 "$slice" 的值必须是负整数。

#### $sort +  $slice + $each + $push

用于 排行榜之类的

db.movies.find({"genre" : "horror"},
... {"$push" : {"top10" : {
... "$each" : [{"name" : "Nightmare on Elm Street", "rating" : 6.6},
... {"name" : "Saw", "rating" : 4.3}],
... "$slice" : -10,
... "$sort" : {"rating" : -1}}}})

#### 把数组当做集合来使用

db.papers.update({"authors cited" : {"$ne" : "Richie"}},
... {$push : {"authors cited" : "Richie"}})

更好的使用 $addToSet

db.users.update({"_id" : ObjectId("4b2d75476cc613d5ee930164")},
... {"$addToSet" : {"emails" : "joe@gmail.com"}})

#### 把数组当做堆栈来使用

是把数组看成队列或者栈，可以用 "$pop"，
这个修改器可以从数组任何一端删除元素。 {"$pop":{"key":1}} 从数组末尾删除
一个元素， {"$pop":{"key":-1}} 则从头部删除

有时需要基于特定条件来删除元素，而不仅仅是依据元素位置，这时可以使用
"$pull"。例如，有一个无序的待完成事项列表：

db.lists.insert({"todo" : ["dishes", "laundry", "dry cleaning"]})
如果我们先完成了洗衣服（laundry），就可以用下面的方式删除它：
db.lists.update({}, {"$pull" : {"todo" : "laundry"}})

#### 按顺序访问数组

```js
db.blog.posts.findOne()
{
"_id" : ObjectId("4b329a216cc613d5ee930192"),
"content" : "...",
"comments" : [
	{
		"comment" : "good post",
		"author" : "John",
		"votes" : 0
	},
	{
		"comment" : "i thought it was too short",
		"author" : "Claire",
		"votes" : 3
	},
	{
		"comment" : "free watches",
		"author" : "Alice",
		"votes" : -1
	}
]}
如果想增加第一个评论的投票数量，可以这么做：
> db.blog.update({"post" : post_id},
... {"$inc" : {"comments.0.votes" : 1}})
或者使用 子文档 查询后的 占位符$
db.blog.update({"comments.author" : "John"},
... {"$set" : {"comments.$.author" : "Jim"}})

```

#### upsert

upsert 是一种特殊的更新。要是没有找到符合更新条件的文档，就会以这个条
件和更新文档为基础创建一个新的文档。如果找到了匹配的文档，则正常更新。
upsert 非常方便，不必预置集合，同一套代码既可以用于创建文档又可以用于更
新文档。

db.analytics.update({"url" : "/blog"}, {"$inc" : {"pageviews" : 1}}, true)
这行代码和之前的代码作用完全一样，但它更高效，并且是原子性的

#### 多文档更新

db.users.update({"birthday" : "10/13/1978"},
... {"$set" : {"gift" : "Happy Birthday!"}}, false, true)

### 查询

1. 使用 find 或者 findOne 函数和查询文档对数据库执行查询；
2.  使用 $ 条件查询实现范围查询、数据集包含查询、不等式查询，以及其他一些查询； 
3. 查询将会返回一个数据库游标，游标只会在你需要时才将需要的文档批量返回；
4. 还有很多针对游标执行的元操作，包括忽略一定数量的结果，或者限定返回结果
   的数量，以及对结果排序。

#### 普通find

`db.users.find({"username" : "joe", "age" : 27})`  and查询

db.users.find({}, {"username" : 1, "email" : 1})

{
"_id" : ObjectId("4ba0f0dfd22aa494fd523620"),
"username" : "joe",
"email" : "joe@example.com"
}

默认会带上 _ id 除非在 { "_ id" : 0}

#### 大于小与范围 查询条件

"$lt" , "lte", "$gt", "gte"    <  .<=.  >. >= 

db.users.find({"age" : {"$gte" : 18, "$lte" : 30}})

db.users.find({"registered" : {"$lt" : new Date("2018/01/01")}})

db.raffle.find({"ticket_no" : {"$in" : [725, 542, 390]}})

db.raffle.find({"ticket_no" : {"$nin" : [725, 542, 390]}})

#### or  not  mod查询  

db.raffle.find({"$or" : [{"ticket_no" : 725}, {"winner" : true}]})

db.users.find({"id_num" : {"$mod" : [5, 1]}})

db.users.find({"id_num" : {"$not" : {"$mod" : [5, 1]}}})

#### null

db.c.find({"y" : null})
{ "_ id" : ObjectId("4ba0f0dfd22aa494fd523621"), "y" : null }
但是， null 不仅会匹配某个键的值为 null 的文档，而且还会匹配不包含这个键的
文档。所以，这种匹配还会返回缺少这个键的所有文档：

db.c.find({"z" : null})
{ "_ id" : ObjectId("4ba0f0dfd22aa494fd523621"), "y" : null }
{ "_ id" : ObjectId("4ba0f0dfd22aa494fd523622"), "y" : 1 }
{ "_ id" : ObjectId("4ba0f148d22aa494fd523623"), "y" : 2 }

查询存在且为null 的

db.c.find({"z" : {"$in" : [null], "$exists" : true}})

#### 正则表达式查询

db.users.find({"name" : /joe/i})

> MongoDB 可以为前缀型正则表达式（比如 /^joey/）查询创建索引，所
> 以这种类型的查询会非常高效。

#### 查询数组

询数组元素与查询标量值是一样的。例如，有一个水果列表，如下所示：

db.food.insert({"fruit" : ["apple", "banana", "peach"]})
下面的查询：
db.food.find({"fruit" : "banana"})
会成功匹配该文档。这个查询好比我们对一个这样的（不合法）文档进行查询：
{"fruit" : "apple", "fruit" : "banana", "fruit" : "peach"}。



要找到既有 "apple" 又有 "banana" 的文档，可以使用 "$all" 来查询：

db.food.find({fruit : {$all : ["apple", "banana"]}})



db.food.find({"fruit.2" : "peach"})
数组下标都是从 0 开始的，所以上面的表达式会用数组的第 3 个元素和 "peach" 进
行匹配。



查询数组的长度 $size

db.food.find({"fruit" : {"$size" : 3}})



slice切片操作符

假设现在有一个博客文章的文档，我们希望返回前 10 条评论，可以这样做：

db.blog.posts.findOne(criteria, {"comments" : {"$slice" : 10}})
也可以返回后 10 条评论，只要在查询条件中使用 -10 就可以了：

db.blog.posts.findOne(criteria, {"comments" : {"$slice" : -10}})
"$slice" 也可以指定偏移值以及希望返回的元素数量，来返回元素集合中间位置
的某些结果：

db.blog.posts.findOne(criteria, {"comments" : {"$slice" : [23, 10]}})
这个操作会跳过前 23 个元素，返回第 24～ 33 个元素。如果数组不够 33 个元素，
则返回第 23 个元素后面的所有元素。
除非特别声明，否则使用 "$slice" 时将返回文档中的所有键。别的键说明符都是
默认不返回未提及的键，这点与 "$slice" 不太一样。



返回任意一个匹配的结果

db.blog.posts.find({"comments.name" : "bob"}, {"comments.$" : 1})
{
"_id" : ObjectId("4b2d75476cc613d5ee930164"),
"comments" : [
{
"name" : "bob",
"email" : "bob@example.com",
"content" : "good post."
}
]
}



#### 查询数组可能会出现的bug

假如有如下所示的文档：
{"x" : 5}
{"x" : 15}
{"x" : 25}
{"x" : [5, 25]}
如果希望找到 "x" 键的值位于 10 和 20 之间的所有文档，直接想到的查询方式是使
用 db.test.find({"x" : {"$gt" : 10, "$lt" : 20}})，希望这个查询的返
回文档是 {"x" : 15}。但是，实际返回了两个文档：

db.test.find({"x" : {"$gt" : 10, "$lt" : 20}})
{"x" : 15}
{"x" : [5, 25]}
5 和 25 都不位于 10 和 20 之间，但是这个文档也返回了，因为 25 与查询条件中的
第一个语句（大于 10）相匹配， 5 与查询条件中的第二个语句（小于 20）相匹配。

#### 内嵌文档的查询

db.people.find({"name.first" : "Joe", "name.last" : "Schmoe"})

将 URL 作为键保存时经常会遇到此类问题。一种解决方法就是在插入前或者提取后
执行一个全局替换，将“.”替换成一个 URL 中的非法字符。

#### 内嵌数组的查询

db.blog.find({"comments" : {"$elemMatch" : {"author" : "joe","score" : {"$gte" : 5}}}})

不能使用db.blog.find({"comments.author" :"joe"， "comments.score" : {"$gte" : 5}}

#### where 查询

不是非常必要时，一定要避免使用 "$where" 查询，因为它们在速度上要比常规
查询慢很多。每个文档都要从 BSON 转换成 JavaScript 对象，然后通过 "$where"
表达式来运行。而且 "$where" 语句不能使用索引，所以只在走投无路时才考虑
"$where" 这种用法。先使用常规查询进行过滤，然后再使用 "$where" 语句，这
样组合使用可以降低性能损失。如果可能的话，使用 "$where" 语句前应该先使用
索引进行过滤， "$where" 只用于对结果进行进一步过滤。

#### limit、 skip和sort

常见的分页操作

db.stock.find({"desc" : "mp3"}).limit(50).skip(50).sort({"price" : -1})

但是利用skip  效率不高

例如，要按照 "date" 降序显示文档列表。可以用如下方式获取结果的第一页：

var page1 = db.foo.find().sort({"date" : -1}).limit(100)
然后，可以利用最后一个文档中 "date" 的值作为查询条件，来获取下一页：
var latest = null;
// 显示第一页
while (page1.hasNext()) {
latest = page1.next();
display(latest);
}
// 获取下一页
var page2 = db.foo.find({"date" : {"$gt" : latest.date}});
page2.sort({"date" : -1}).limit(100);
这样查询中就没有 skip 了。





## 3. 索引

### 索引介绍

不使用索引的查询称为全表扫描（这个术语来自关系型数据库），也就是说，服务器
必须查找完一整本书才能找到查询结果。这个处理过程与我们在一本没有索引的书
中查找信息很像

db.users.find({username: "user101"}).explain()
{
"cursor" : "BasicCursor",
"nscanned" : 1000000,
"nscannedObjects" : 1000000,
"n" : 1,
"millis" : 721,
"nYields" : 0,
"nChunkSkips" : 0,
"isMultiKey" : false,
"indexOnly" : false,
"indexBounds" : {
}
}

"nscanned" 是 MongoDB 在完成这个查询的过程中扫描的文档总数。可以看到，
这个集合中的每个文档都被扫描过了。也就是说，为了完成这个查询， MongoDB
查看了每一个文档中的每一个字段



### 索引类型

1. 唯一索引  db.users.ensureIndex({"username" : 1}, {"unique" : true})有一个唯一索引可能你已经比较熟悉了，就是 "_id" 索引，这个索引会在创建集合
   时自动创建

   db.people.ensureIndex({"username" : 1}, {"unique" : true, "dropDups" : true})

   删除已经存在的 重复值

2. 复合唯一索引

3. 稀疏索引，在有些情况下，你可能希望唯一索引只对包含
   相应键的文档生效。如果有一个可能存在也可能不存在的字段，但是当它存在时，它必须是唯一的，这时就可以将 unique 和 sparse 选项组合在一起使用

   b.ensureIndex({"email" : 1}, {"unique" : true, "sparse" : true})。非唯一的稀疏
   索引

4. 普通索引

db.collectionName.getIndex()查看所有索引

db.collectionName.dropIndex() 删除索引



### 创建索引

db.users.ensureIndex({"username" : 1})

然而，使用索引是有代价的：对于添加的每一个索引，每次写操作（插入、更新、删除）都将耗费更多的时间。这是因为，当数据发生变动时， MongoDB 不仅要更新文档，还要更新集合上的所有索引。因此， MongoDB 限制每个集合上最多只能有 64 个索引。通常，在一个特定的集合上，不应该拥有两个以上的索引。于是，挑选合适的字段建立索引非常重要。

#### 复合索引

在下面的排序里，"username" 上的索引没什么作用：

db.users.find().sort({"age" : 1, "username" : 1})
这里先根据 "age" 排序再根据 "username" 排序，所以 "username" 在这里发挥的
作用并不大。为了优化这个排序，可能需要在 "age" 和 "username" 上建立索引：
db.users.ensureIndex({"age" : 1, "username" : 1})
这样就建立了一个复合索引（compound index）。

查询那些不包含 "x" 键的文档可以使用这样的索引 ({"x": {"$exists" : false}}

### 索引对象和数组

MongoDB 允许深入文档内部，对嵌套字段和数组建立索引。嵌套对象和数组字段
可以与复合索引中的顶级字段一起使用，虽然它们比较特殊，但是大多数情况下与
“正常”索引字段的行为是一致的。

{
"username" : "sid",
"loc" : {
"ip" : "1.2.3.4",
"city" : "Springfield",
"state" : "NY"
}
}

db.users.ensureIndex({"loc.city" : 1})

注意，对嵌套文档本身（"loc"）建立索引，与对嵌套文档的某个字段（"loc.
city"）建立索引是不同的。对整个子文档建立索引，只会提高整个子文档的查
询速度。在上面的例子中，只有在进行与子文档字段顺序完全匹配的子文档查询
时（比如 db.users.find({"loc" : {"ip" : "123.456.789.000"， "city"
: "Shelbyville"， "state" : "NY"}}})），查询优化器才会使用 "loc" 上的索
引。无法对形如 db.users.find({"loc.city" : "Shelbyville"}) 的查询使用
索引。



假如有一个博客文章的集合，其中每个文档表示一篇文章。每篇文章都有一个
"comments" 字段，这是一个数组，其中每个元素都是一个评论子文档。如果想要
找出最近被评论次数最多的博客文章，可以在博客文章集合中嵌套的 "comments"
数组的 "date" 键上建立索引：

无法将整个数组作为一个实体建立索引：对数组建
立索引，实际上是对数组中的每个元素建立索引，而不是对数组本身建立索引。

### 使用索引的原则

1. 优先使用排在 复合索引前面的字段进行查询
2. 考虑各个查询 是否能使用索引
3. 一般说来，应该在基数比较高的键上建立索引，或者至少应该把基数较高的键放在复合索引的前面（低基数的键之前）
4. 使用explain()和hint()

hint（） 表示强制使用某个索引

db.c.find({"age" : 14, "username" : /.*/}).hint({"username" : 1, "age" : 1})

db.entries.find({"created_at" : {"$lt" : hourAgo}}).hint({"$natural" : 1}) 强制进行全表扫描

实际返回的文档数量： "n"。它无法反映出 MongoDB 在执行这个查询的过
程中所做的工作：搜索了多少索引条目和文档。索引条目是使用 "nscanned" 描述
的。 "nscannedObjects" 字段的值就是所扫描的文档数量。最后，如果要对结果
集进行排序，而 MongoDB 无法对排序使用索引，那么 "scanAndOrder" 的值就会
是 true。也就是说， MongoDB 不得不在内存中对结果进行排序，这是非常慢的，
而且结果集的数量要比较小

### 并不一定是索引好

| 索引通常适用的情况 | 全表扫描通常适用的情况 |
| ------------------ | ---------------------- |
|                    |                        |

|集合较大 |集合较小|

|文档较大| 文档较小|

|选择性查询| 非选择性查询|



## 4. 特殊的集合和索引

### 1. 固定集合

普通”集合是动态创建的，而且可以自动增长以容纳更多的数据。
MongoDB 中还有另一种不同类型的集合，叫做固定集合，固定集合需要事先创建
好，而且它的大小是固定的。

说到固定大小的集合，有一个很有趣
的问题：向一个已经满了的固定集合中插入数据会怎么样？答案是，固定集合的行
为类似于循环队列。如果已经没有空间了，最老的文档会被删除以释放空间，新插
入的文档会占据这块空间

（一个圆形的存储）

固定集合可以用于记录日志，尽管它们不够灵活。虽然可以在创建时指定集合大小，
但无法控制什么时候数据会被覆盖（固定集合不能被分片）

固定集合必须显示创建

 db.createCollection("my_collection", {"capped" : true, "size" : 100000});

动态集合转换成 固定的集合

db.runCommand({"convertToCapped" : "test", "size" : 10000});

使用自然顺序（固定集合的顺序就是 插入的顺序）

db.my_collection.find().sort({"$natural" : -1})

### 2. 没有_ id 索引的集合

在调用 createCollection
创建集合时指定 autoIndexId 选项为 false，创建集合时就不会自动在 "_id" 上
创建索引。

> 如果创建了一个没有 "_id" 索引的集合，那就永远都不能复制它所在的
> mongod 了。复制操作要求每个集合上都要有 "_id" 索引（对于复制操作，
> 能够唯一标识集合中的每一个文档是非常重要的）。

### 3. TTL索引

如果需要更加灵活的老化移出系统（age-out system），可以使用 TTL 索引
（time-to-live index， 具有生命周期的索引），这种索引允许为每一个文档设置一个超时时间。一个文档到达预设置的老化程度之后就会被删除。这种类型的索引对于缓
存问题（比如会话的保存）非常有用。

在 ensureIndex 中指定 expireAfterSecs 选项就可以创建一个 TTL 索引：

// 超时时间为 24 小时
db.test.ensureIndex({"lastUpdated" : 1}, {"expireAfterSeconds" : 60 × 60 × 24})

### 4. 全文本索引

使用全文本索引可以非常快地进行文本搜索，就如同内置了多种语言
分词机制的支持一样。

创建任何一种索引的开销都比较大，而创建全文本索引的成本更高。在一个操作频
繁的集合上创建全文本索引可能会导致 MongoDB 过载，所以应该是离线状态下创
建全文本索引，或者是在对性能没要求时。创建全文本索引时要特别小心谨慎，内
存可能会不够用（除非你有 SSD）。

db.blobs.ensureIndex({"title" : "text", "desc" : "text", "author" : "text"})

默认的权重是 1，权重的范围可以是 1~1 000 000 000。使用上面的代码设置权重之
后， "title" 字段成为其中最重要的字段， "author" 其次，最后是 "desc"（没有
指定，因此它的权重是默认值 1）。

 db.users.ensureIndex({"profil" : "text", "intérêts" : "text"},
... {"default_language" : "french"})

### 5. 地理位置索引

### 6. GridFS索引



## 5. 聚合操作

db.articles.aggregate({"$project" : {"author" : 1}},
... {"$group" : {"_id" : "$author", "count" : {"$sum" : 1}}},
... {"$sort" : {"count" : -1}},
... {"$limit" : 5})

### $match

通常，在实际使用中应该尽可能将 "$match" 放在管道的前面位置。这样做有两个
好处：一是可以快速将不需要的文档过滤掉，以减少管道的工作量；二是如果在投
射和分组之前执行 "$match"，查询可以使用索引

### $project

1. 过滤

db.articles.aggregate({"$project" : {"author" : 1, "_id" : 0}}) 只会返回一个 author 字段

2. 重命名

db.users.aggregate({"$project" : {"userId" : "$_id", "_id" : 0}})

3. 聚合框架无法在下面的排序操作
   中使用这个索引，尽管人眼一下子就能看出下面代码中的 "newfieldname" 与
   "originalfieldname" 表示同一个字段。

   db.articles.aggregate({"$project" : {"newfieldname" : "$originalfieldname"}},
   ... {"$sort" : {"newfieldname" : 1}})
   所以，应该尽量在修改字段名称之前使用索引。

#### 数学表达式

db.employees.aggregate(
... {
... "$project" : {
... "totalPay" : {
... "$add" : ["$salary", "$bonus"]
... }} })

• "$add" : [expr1[, expr2, ..., exprN]]
这个操作符接受一个或多个表达式作为参数，将这些表达式相加。
• "$subtract" : [expr1, expr2]
接受两个表达式作为参数，用第一个表达式减去第二个表达式作为结果。
• "$multiply" : [expr1[, expr2, ..., exprN]]
接受一个或者多个表达式，并且将它们相乘。
• "$divide" : [expr1, expr2]
接受两个表达式，用第一个表达式除以第二个表达式的商作为结果。
• "$mod" : [expr1, expr2]
接受两个表达式，将第一个表达式除以第二个表达式得到的余数作为结果。

#### 日期表达式

"$year"、"$month"、 "$week"、 "$dayOfMonth"、 "$dayOfWeek"、 "$dayOfYear"、"$hour"、 "$minute" 和 "$second"

db.employees.aggregate(
... {
... "$project" : {
... "hiredIn" : {"$month" : "$hireDate"}
... }
... })

结合使用 

db.employees.aggregate(
... {
... "$project" : {
... "tenure" : {
... "$subtract" : [{"$year" : new Date()}, {"$year" : "$hireDate"}]
... }
... }
... })

#### 字符串表达式

• "$substr" : [expr, startOffset, numToReturn]
其中第一个参数 expr 必须是个字符串，这个操作会截取这个字符串的子串（从
第 startOffset 字节开始的 numToReturn 字节，注意，是字节，不是字符。在
多字节编码中尤其要注意这一点） expr 必须是字符串。
• "$concat" : [expr1[, expr2, ..., exprN]]
将给定的表达式（或者字符串）连接在一起作为返回结果。
• "$toLower" : expr
参数 expr 必须是个字符串值，这个操作返回 expr 的小写形式。
• "$toUpper" : expr
参数 expr 必须是个字符串值，这个操作返回 expr 的大写形式。

#### 逻辑表达式

• "$cmp" : [expr1, expr2]
比较 expr1 和 expr2。如果 expr1 等于 expr2，返回 0；如果 expr1 < expr2，
返回一个负数；如果 expr1 >expr2，返回一个正数。
• "$strcasecmp" : [string1, string2]
比较 string1 和 string2，区分大小写。只对罗马字符组成的字符串有效。
• "$eq"/"$ne"/"$gt"/"$gte"/"$lt"/"$lte" : [expr1, expr2]
对 expr1 和 expr2 执行相应的比较操作，返回比较的结果（true 或 false）。
下面是几个布尔表达式。
• "$and" : [expr1[, expr2, ..., exprN]]
如果所有表达式的值都是 true，那就返回 true，否则返回 false。
• "$or" : [expr1[, expr2, ..., exprN]]
只要有任意表达式的值为 true，就返回 true，否则返回 false。
• "$not" : expr
对 expr 取反。
还有两个控制语句。
• "$cond" : [booleanExpr, trueExpr, falseExpr]
如果 booleanExpr 的值是 true，那就返回 trueExpr，否则返回 falseExpr。
• "$ifNull" : [expr, replacementExpr]
如果 expr 是 null，返回 replacementExpr，否则返回 expr。

db.students.aggregate(
... {
... "$project" : {
... "grade" : {
... "$cond" : [
... "$teachersPet",
... 100, // if
... { // else
... "$add" : [
... {"$multiply" : [.1, "$attendanceAvg"]},
... {"$multiply" : [.3, "$quizzAvg"]},
... {"$multiply" : [.6, "$testAvg"]}
... ]
... }
... ]
... }
... }
... })



### $group

1. "$sum" : value
2. "$avg" : value
3. "$max" : expr
4. "$min" : expr
5. "$first" : expr
6. "$last" : expr
7. "$addToSet" : expr
8. "$push" : expr

### $unwind

如果要得
到特定用户的所有评论（只需要得到评论，不需要返回评论所属的文章），使用普通
的查询是不可能做到的。但是，通过提取、拆分、匹配，就很容易了：

db.blog.aggregate({"$project" : {"comments" : "$comments"}},
... {"$unwind" : "$comments"},
... {"$match" : {"comments.author" : "Mark"}})

### $sort

如果要对大量的文档进行排序，强烈建议在管道的第一阶段进行排序，这时的排序操作可以使用索引。否则，排序过程就会比较慢，而且会占用大量内存

### $limit

### $skip

### mapreduce











## 附录

### 填充因子

填充因子是 MongoDB 为每个新文档预留的增长空间。可以运行 db.coll.stats()
查看填充因子。执行上面的更新之前， "paddingFactor" 字段的值是 1：根据实际
的文档大小，为每个新文档分配精确的空间，不预留任何增长空间，如图 3-1 所示。
让其中一个文档增大之后，再次运行这个命令（如图 3-2 所示），会发现填充因子增
加到了 1.5：为每个新文档预留其一半大小的空间作为增长空间，如图 3-2 所示。如
果随后的更新导致了更多次的文档移动，填充因子会持续变大（虽然不会像第一次
移动时的变化那么大）。如果不再有文档移动，填充因子的值会缓慢降低，如图 3-3
所示

![选区_105](/home/jsy/图片/选区_105.png)

### PyMongo 的聚合操作

```python
from pymongo import MongoClient
from bson.son import SON
from pprint import pprint

client = MongoClient("127.0.0.1",username="jsy",password="jsy1164+",authSource="new",authMechanism='SCRAM-SHA-1')

db = client.new

def aggregate_data():
    db.collection_one.drop()
    result = db.collection_one.insert_many([
        {
            "x": 1,
            "tags": ["daa","cat"]
            },
        {   
            "x": 2,
            "tags": ["cat"]
            },  
        {   
            "x": 2,
            "tags": ["mouse","cat","dog"]
            },  
        {   
            "x":3,
            "tags": []
            }   
        ])  
    #print(result.inserted_ids) # [ObjectId('5c56b7226e1e71090d214468'), ObjectId('5c56b7226e1e71090d214469'), ObjectId('5c56b7226e1e71090d21446a'), ObjectId('5c56b7226e1e71090d21446b')]
        
    result = db.collection_one.aggregate([
        {"$unwind":"$tags"},
        {"$group": {"_id": "$tags","count": {"$sum":1}}},
        {"$sort": SON([("count",-1),("_id",-1)])} # SON or OrderedDict
        ])  
    # 上面的等于 db.command("aggregate","things", pipeline=pipeline, explain=True)
    #pprint(list(result))[{'_id': 'cat', 'count': 3},
# {'_id': 'mouse', 'count': 1},
# {'_id': 'dog', 'count': 1},
# {'_id': 'daa', 'count': 1}]

from pymongo import InsertOne,DeleteMany, ReplaceOne, UpdateOne
def bulk_data():
    db.collection_two.bulk_write([
        DeleteMany({}),
        InsertOne({"_id":1}),
        InsertOne({"_id":2}),
        InsertOne({"_id":3}),
        InsertOne({"_id":5}),
        InsertOne({"_id":6})
        ], ordered=False) # order指定是否 是有序操作
    

if __name__ == "__main__":

    client.close()
```



