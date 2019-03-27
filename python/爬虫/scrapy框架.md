# scrapy



项目目录

scrapy.cfg	配置文件，部署信息

project/

​	__ init __.p

​	items.py			定义了所有item的数据结构

​	pipelines.py			item pipeline都在这里

​	settings.py			项目的全局配置

​	middlewares.py		Spider Middlewares 和 Downloader Middlewares的实现

​	spiders/

​		__ init __.py

​		spider1.py

​		spider2.py



## demo

> scrapy startproject tutorial

> cd tutorial && scrapy genspider quotes quotes.toscrape.com

构造Request

1. response.urljoin("xxx") 	用来生成绝对url
2. yield scrapy.Request(url=url, callback=self.parse)

然后

> scrapy crawl quotes

如果不使用pipelines保存结果

1. scrapy crawl quotes -o quotes.json
2. scrapy crawl quotes -o quotes.jl  或者 scrapy crawl quotes -o quotes.jsonlines
3. scrapy crawl quotes -o quotes.csv
4.    -o  .xml
5.  -o .pickle
6.  -o .marshal
7.  -o ftp://user:pass@ftp.example.com/path/to/quotes.csv

> pipeline的简单用法 ，具体demo见 /home/jsy/deskpt/python/scrapy_s/tutorial



### scrapy解析数据

1. response.css()
2. response.xpaht()

具体说一下xpath的用法

1. /开头表示从根节点开始
2.   ./开头从上一个节点开始
3. 所有selectList类型都可以调用 xpath()
4. extract() 从selector类型中获取数据
5. extract_first("default value") 不用担心数组越界的问题
6. /text()  @href等。。





### spider的用法

1. name  爬虫名字
2. allowed_domains
3. start_urls
4. custom_settings 是专属spider的配置，会覆盖 全局 设置 ，，，字典类型
5. crawler。
6. settings  用于获取全句配置

常用的方法

1. start_requests() 。用于修改 start_urls的值，一般用来刚开始爬取 post请求页面
2. parse() ，默认的回调函数
3. closed() 



### Downloader Middleware用法

可以再两个地方修改 调度的内容

1. schedule调度出队列的Request 发送给Doanload之前
2. 在response发送给spider之前，也就是解析之前。

对应下面三个方法

1. process_request(request, spider)
2. process_response(request, response, spider)
3. process_exception(request, exception, spider)

> 以上三个方法的 返回类型不同，它的执行逻辑也不同，具体见 python3网络爬虫实战，489

### spider Middlewares 用法

修改内容

1. spider收到response之前
2. spider request发送给Scheduler之前
3. spider Item 发送给ItemPipeline之前

方法

1. process_spider_input(resonse, spider)

   一定会被调用，返回None，或者抛出异常

2. process_spider_output(response, result, spider)

   返回request 或者 Item

3. process_spider_exception(response, exception, spider)

4. process_spider_requests(start_reauests, spider)





### scrapy 中有专门内置的 image pipeline 

用来下载东西

1.  get_media_requests(self, request, response, info)
2. file_path(self, result, item, info)
3. item_completed(self, item, info )