# 后台roadmap

> 个人觉得的一个合格后台工程师  (不是架构师)  应该具备的知识

## 前端

1. html
2. css
3. JavaScript(>=es6 or  Typescript)
4. react全家桶 或者 vue技术栈
5. ajax

## linux

- 基本的命令
- 系统相关知识 ，推荐《鸟哥的linux私房菜》
- 了解 linux内核相关

## 数据库

- 学习SQL ，熟悉关系型数据库的设计，三范式 推荐学习mysql 《高性能mysql》
- Google搜一下数据库最佳实践

- 文档型数据库，MongoDB 推荐官网加《mongodb权威指南》着重看反范式设计
- 基于内存的数据库 ，推荐redis，《Redis设计与实现》《redis in action》

## 编程语言

> 推荐动静结合，例如java+python 。举例python

### python

- 语法
- 编码规范（pep8，Google风格，prooce等）
- 高级用法
- 框架（flask，Django，tornado）等
- 框架的第三方工具（restful，数据库，单元测试相关插件）
- 前后端交互基本数据格式：json、form、query、param
- 基本认证方式：Cookie、Session、JWT
- 实践模板引擎
- restful api
- 实践graphql
- 微服务
- rpc，protobuf等

### golang

## 设计模型

《设计模式》

## 文件格式

json, yaml, xml, toml, csv等，序列化和反序列化

## 函数式编程

- 函数式
- 函数柯里化
- 高阶函数
- 可以看一下将其运用好极致的 haskell

## 爬虫

- 模拟登陆，自动化测试工具的使用等

- 熟悉反爬机制
- 熟悉代理，cookie池，请求头等
- 分布式爬虫，可以学习scrapy

## 开发部署工具

- 科学上网（可以自己搭vps，不贵）
- 熟练git、github基本操作
- 使用 macos 或者linux系统 
- nginx/openresty
- 使用 postman 或者其他软件进行后台接口测试
- 使用 docker、docker-compose
- Fiddler/Wireshark/Charles 抓包



## 微信开发

看微信开发文档



## 操作系统

- 进程、线程、协程

- 进程控制、进程通信、锁、PV操作

- 内存管理，缓存机制

- 线程调度
-  epoll poll select 
- 异步同步，阻塞非阻塞

## 网络

1. http协议 https   可阅读《图解http》
2. 计算机网络七层协议  可阅读《计算机网络-自顶向下》《图解tcp/ip》等书
3. http2.0

## 安全

1. hash算法，对称，非对称加密rsa等
2. web安全
   - 防火墙配置
   - DDOS攻击
   - 中间人攻击
   - sql注入
   - ip欺骗
   - xss攻击
   - csrf
   - 远程脚本执行

## 算法

- 看书《算法导论》《算法》第四版，java描述
- 刷题---leetcode

## 机器学习

- 分类算法  推荐 机器学习西瓜书
- 回归算法  
- 神经网络  推荐花书
- 信息熵，损失函数，评估（召回率，准确率等）
- 流程 ： 数据预处理 ---  跑算法  ---  调参数（网格搜索等） ---  评估模型 --- 循环往复 



## 分布式

- 学习区块链

- Hadoop、Spark、Flink、Ray 等大数据平台、工具

核心问题：

- 一致性问题

- 共识算法（Paxos 与 Raft 等等）

- FLP 不可能性原理

- CAP 原理

- ACID 原则

- 拜占庭问题

- 可靠性指标

- MapReduce、GFS、Bigtable