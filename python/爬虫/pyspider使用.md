> linux 安装 出现错误 https://blog.csdn.net/FK103/article/details/79875453



# pyspider 使用

> pyspider all 开启web服务器

架构

1. Scheduler
2. Fetcher  抓取
3. Processor 处理响应，并且提取出要继续爬取的页面。如果生成了新的提取结果，就交给Result Worker



pyspider 默认发送的是http请求，请求的html文档不包含img节点。所以我们应该在index_page()中加入`fetch_type='js'`


### 状态

1. TODO：刚刚创建还未被实现
2. STOP： 停止某个项目
3. CHECKING：修改后
4. DEBUG/RUNING：运行
5. PAUSE：如果多次错误，则会停一段时间



> 如果要删除项目，那么要将状态变成 STOP，并且修改分组为 delete，过24小时，它就会被自动删除。



### 命令行

> pyspider [OPTIONS] COMMAND [ARGS]

Options:

1. -c --config  Filename             指定配置文件的名字
2. --logging-config TEXT             日志配置文件名称
3. --debug              开启调试模式
4. --queue-maxsize INTERER          队列的最大长度
5. --projectdb TEXT                 project数据库连接字符串 默认为：sqlite
6. --taskdb TEXT                      taskdb数据库连接字符串 默认为：sqlite
7. --resultdb TEXT                      resultdb的数据库连接字符串 默认：multiprocessing.Queue 
8. --phantomjs-proxy TEXT             PhantomJS使用的代理，ip:port 形式
9. --data-path TEXT               数据库存放的路径
10. --version            版本
11. --help

配置文件格式

```json
{
    "taskdb":"mysql+taskdb://username:password@host:port/taskdb",
    "projectdb": "mysql+projectdb://username:password@host:port/projectdb",
    "resultdb": "mysql+resultdb://username:password@host:port/resultdb",
    "message_queue": "amqp://username:password@host:port/%2F",
    "webui": {
        "username": "name",
        "password": "password",
        "need-auth": true
    }
}
```



pysider 还可以单独运行某个模块

1. pyspider scheduler [OPTIONS]
2. pyspider fetcher [OPTIONS]
3. pyspider processer [OPTIONS]
4. pyspider webui --port xxx



### crawl()方法详解

1. url 可以是列表，字符串

2. callback，回调函数

3. age，是任务的有效时间，如果某个任务在有效时间 内且已经执行，就不会重复执行。或者，默认有效时间10天。

   ```python
   @config(age=10 * 24 * 60 * 60)
   def callback(self):
       pass
   ```

4. priority 设置爬去的优先级

5. exetime = time.time() + 30* 60 ，设置定时任务

6. retries  定义重试次数，默认是3

7. itag，判断网页时候发生变化，通过节点的值判断。

   ```python
   def index_page(self, response):
       for item in response.doc(".item").items():
           self.crawl(item.find("a").attr.url,callback=detail,itag=item.find(".update-time").text())
   ```

8. auto_recrawl，会按照age指定的时间，重复执行

9. method，默认是get

10. params

    ```python
    def on_start(self):
        self.crawl("http://http", params={"a":123,"b":"c"})
        self.crawl("http://sdfjk/get?a=123&b=c")
    ```

11. data  post时的数据

12. files，上传的文件

13. user_agent

14. headers

15. cookies

16. connect_timeout

17. timeout  最长等待时间

18. allow_redirects ，是否自动重定向 True

19. validata_cert，是否验证证书，HTTPS 默认是True

20. proxy 格式username@password@hostname:port

21. fetch_type  如果指定为 js 则会使用phantomjs 爬取

22. js_script  是 页面加载完成后 会执行的js语句

    ```python
    def on_start(self):
        self.crawl("xxxxx",callback=self.callback,fetch_type="js", js_script="""
        function() {
        	window.scrollTo(0,document.body.scrollHeight)
        	return 123;
        }
        """)
    ```

23. js_run_at

24. js_viewport_width

25. load_images  时候加载图片，默认是否

26. save， 可以在不同的函数间传递参数。

    ```python
    def on_start(self):
        self.crawl("xx",callback=self.callback, save={"page":1})
        
    def callback(self,response):
        return response.save["page"]
    ```



### pyspider 区分不同的任务

pyspider使用任务URL的MD5值区分任务，如果对同一个url post数据，那么可以重写task_id() 来区分不同的任务。这样就不会被设别为同一个任务了。

```python
import json
from pyspider.libs.utils import md5string
def get_taskid(self, task):
    return md5string(task["url"] + json.dumps(task["fetch"].get("data","")))
```

### 通过every装饰器配置定时爬取

@every(minutes=24*60)







































