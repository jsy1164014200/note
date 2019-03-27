

# NGINX

#### 安装

> sudo apt install nginx

#### ufw防火墙工具（基于iptables）

>  [简单使用](https://www.jianshu.com/p/eccb913ac58d)

```sh
>>> sudo ufw app list
Output
Available applications:
Nginx Full
Nginx HTTP
Nginx HTTPS
OpenSSH
>>> sudo ufw allow 'Nginx Full'
>>> sudo ufw status
```

#### 启动nginx

```sh
>>> systemctl status nginx

基本操作
sudo systemctl stop nginx    
sudo systemctl start nginx
sudo systemctl restart nginx
sudo systemctl reload nginx
sudo systemctl disable nginx
sudo systemctl enable nginx
```

### 熟悉几个基本的文件

#### 内容

/var/www/html：实际的Web内容（默认情况下只包含您之前看到的默认Nginx页面）已被提供给/var/www/html目录。这可以通过更改Nginx配置文件来更改。

#### 服务器配置

/etc/nginx：nginx配置目录。所有的Nginx配置文件驻留在这里。
 /etc/nginx/nginx.conf主要的Nginx配置文件。这可以修改为对Nginx全局配置进行更改。
 /etc/nginx/sites-available：可以存储每个站点“服务器块”的目录。Nginx不会使用此目录中找到的配置文件，除非它们链接到sites-enabled目录（见下文）。通常，所有服务器块配置都在此目录中完成，然后通过链接到其他目录来启用。
 /etc/nginx/sites-enabled/：存储启用每个站点“服务器块”的目录。通常，这些是通过链接到目录中找到的配置文件创建的sites-available。
 /etc/nginx/snippets：此目录包含Nginx配置中可以包含的配置片段。潜在的可重复配置段是重构为代码片段的好候选者。

#### 服务器日志

/var/log/nginx/access.log：对于您的Web服务器的每个请求都将记录在此日志文件中，除非Nginx配置为其他方式。
 /var/log/nginx/error.log：任何Nginx错误将记录在此日志中。

## 基本配置

nginx.conf配置文件,基本就分为以下几块

```
main（全局配置）
events（nginx工作模式）   {
  ....
}
http（http设置）        {
  ....
  upstream myproject（负债均衡服务器设置） {
    .....
  }
  server（主机设置）  {
    ....
    location（url配置） {
        ....
    }
  }
  server  {
    ....
    location {
        ....
    }
  }
  ....
}
```

1. main模块

   ```sh
   user nobody nobody; 
   # user 来指定Nginx Worker进程运行用户以及用户组，默认由nobody账号运行。
   worker_processes auto;
   # worker_processes来指定了Nginx要开启的子进程数。每个Nginx进程平均耗费10M~12M内存。根据经验，一般指定1个进程就足够了，如果是多核CPU，建议指定和CPU的数量一样的进程数即可。
   error_log  /usr/local/var/log/nginx/error.log  notice;
   # error_log用来定义全局错误日志文件。日志输出级别有debug、info、notice、warn、error、crit可供选择，其中，debug输出日志最为最详细，而crit输出日志最少。
   pid        /run/nginx.pid;
   # pid用来指定进程id的存储文件位置
   worker_rlimit_nofile 1024;
   # worker_rlimit_nofile用于指定一个nginx进程可以打开的最多文件描述符数目，这里是65535，需要使用命令“ulimit -n 65535”来设置。
   ```

2. events模块

   events模块来用指定nginx的工作模式和工作模式及连接数上限，一般是这样：

   ```sh
   events {
       use epoll; #Linux平台
       # use用来指定Nginx的工作模式。Nginx支持的工作模式有select、poll、kqueue、epoll、rtsig和/dev/poll。其中select和poll都是标准的工作模式，kqueue和epoll是高效的工作模式，不同的是epoll用在Linux平台上，而kqueue用在BSD系统中,对于Linux系统，epoll工作模式是首选。
       worker_connections  1024;
       # worker_connections用于定义Nginx每个进程的最大连接数，即接收前端的最大请求数，默认是1024。最大客户端连接数由worker_processes和worker_connections决定，即Max_clients=worker_processes*worker_connections，在作为反向代理时，Max_clients变为：Max_clients = worker_processes * worker_connections/4。 进程的最大连接数受Linux系统进程的最大打开文件数限制，在执行操作系统命令“ulimit -n 65536”后worker_connections的设置才能生效。
   }
   ```

3. http模块

   http模块可以说是最核心的模块了，它负责HTTP服务器相关属性的配置，它里面的server和upstream子模块，至关重要。

   ```
   http{
       include       mime.types; # include 来用设定文件的mime类型,类型在配置文件目录下的mime.type文件定义，来告诉nginx来识别文件类型
       default_type  application/octet-stream; # default_type设定了默认的类型为二进制流，也就是当文件类型未定义时使用这种方式，例如在没有配置asp 的locate 环境时，Nginx是不予解析的，此时，用浏览器访问asp文件就会出现下载了。
       log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                         '$status $body_bytes_sent "$http_referer" '
                         '"$http_user_agent" "$http_x_forwarded_for"';
                         # log_format用于设置日志的格式，和记录哪些参数，这里设置为main，刚好用于access_log来纪录这种类型。
       access_log  /usr/local/var/log/nginx/access.log  main; 
       sendfile        on;
       tcp_nopush      on;
       tcp_nodelay     on;
       keepalive_timeout  10;
       #gzip  on; # access_log 用来纪录每次的访问日志的文件地址，后面的main是日志的格式样式，对应于log_format的main。sendfile参数用于开启高效文件传输模式。将tcp_nopush和tcp_nodelay两个指令设置为on用于防止网络阻塞。keepalive_timeout设置客户端连接保持活动的超时时间。在超过这个时间之后，服务器会关闭该连接。
       upstream myproject {
           .....
       }
       server {
           ....
       }
   }
   ```

4. server模块

   ```python
   server {
       # server标志定义虚拟主机开始
   	# listen用于指定虚拟主机的服务端口。
   	#server_name用来指定IP地址或者域名，多个域名之间用空格分开。
           listen       8080;
           server_name  localhost 192.168.12.10 www.yangyi.com;
           # 全局定义，如果都是这一个目录，这样定义最简单。
       	# root 表示在这整个server虚拟主机内，全部的root web根目录。注意要和locate {}下面定义的区分开来。
           root   /Users/yangyi/www;
       	# index 全局定义访问的默认首页地址。注意要和locate {}下面定义的区分开来。
           index  index.php index.html index.htm; 
       	# charset用于设置网页的默认编码格式
           charset utf-8;
       	# access_log用来指定此虚拟主机的访问日志存放路径，最后的main用于指定访问日志的输出格式。
           access_log  usr/local/var/log/host.access.log  main;
           aerror_log  usr/local/var/log/host.error.log  error;
           ....
   }
   ```

5. location 模块

   location模块是nginx中用的最多的，也是最重要的模块了，什么负载均衡啊、反向代理啊、虚拟域名啊都与它相关。location 根据它字面意思就知道是来定位的，定位URL，解析URL，所以，它也提供了强大的正则匹配功能，也支持条件判断匹配，用户可以通过location指令实现Nginx对动、静态网页进行过滤处理。

   ```python
   location / {
               root   /Users/yangyi/www;
               index  index.php index.html index.htm;
           }
   #反向代理配置
     location /xxx/ {
                proxy_pass http://127.0.0.1:12345;
                proxy_set_header X-real-ip $remote_addr;
                proxy_set_header Host $http_host;
            }
     #采用uwsgi方式
     location /python/ {
                include uwsgi_params;
                uwsgi_pass 127.0.0.1:33333;
            }
   
   
   
       #访问nginx本机目录的文件
       location / {
               root   /home/itcast/xwp/itcast/;
               index  index.html index.htm;
           }
   
       location  /static/ {
                alias /var/static/;
   ```

6. 反向代理

   1. proxy_pass

      ```
          proxy_pass URL;
          配置块 location if
          此配置将当前请求代理到URL参数指定的服务器上,URL可以是主机名或者IP地址加PORT的形式
          proxy_pass http://localhost:8000;
          也可以结合负载均衡使用<负载均衡会说明这种情况>
          也可以把HTTP转换成HTTPS
          proxy_pass http://192.168.0.1;
          默认情况反向代理不转发请求中的Host头部,如果需要，则 proxy_set_header Host $host;
      ```

   2. proxy_method

      ```
          proxy_method method_name;
          配置块 http server location
          此配置项表示转发时的协议方法名:
              proxy_method POST;
          那么客户端发来的GET请求在转发时方法改为POST;
      ```

   3. proxy_hide_header

      ```
          proxy_hide_header header1;
          配置块 http server location;
          Nginx会将上游服务器的响应转发给客户端,但默认不转发HTTP头部字段(Date Server X-Pad X-Accel-* )
          使用proxy_hide_header可以指定任意头部不能被转发
          proxy_hide_header Cache-Control;
          proxy_hide_header MicrosoftOfficeWebServer;
      ```

   4. proxy_pass_header

      ```
          proxy_pass_header header1;
          配置块 http server location
          功能与 proxy_hide_header相反,是设置哪些头部允许转发.
          proxy_pass_header X-Accel-Redirect;
      ```

   5. proxy_pass_request_body

      ```
          proxy_pass_request_body on|off;
          默认 on
          配置块 http server location;
          确定上游服务器是否向上游服务器转发HTTP包体
      ```

   6. proxy_pass_request_header

      ```
          proxy_pass_request_header on | off;
          默认on
          配置块 http server location
          确定是否转发HTTP头部
      ```

   7. proxy_redirect

      ```
          proxy_redirect [default | off |redirect |replacement]
          默认default
          配置块 http server location
          当上游服务响应时重定向或刷新(HTTP 301 302),proxy_redirect可以重设HTTP头部的location或refresh字段
      
          proxy_redirect http://locahost:8000/two/ http://frontend/one/;
          上游响应302,location是URI是http://locahost:8000/two/some/uri/
          那是实际转发给客户端的是 http://frontend/one/some/uri/;
          可以使用前面提到的ngx_http_core_module模块提供的变量 
          proxy_redirect http://locahost:8000/two/ http://$host:server_port/;
          可以省略replacement参数的主机名部分,这时候用虚拟主机名填充
          proxy_redirect http://locahost:8000/two/ /one/;
      
          使用off参数的时候,将使location和refresh的字段维持不变
          proxy_redirect off;
      
          如果使用的 proxy_redirect default;
          下面两种配置是等效的
              location /{
                  proxy_pass http://upstream:port/two/;
                  proxy_redirect default;
              }
              location /{
                  proxy_pass http://upstream:port/two/;
                  proxy_redirect http://upstream:port/two/ /one/;
              }
      ```

   8. proxy_next_upstream

      ```
          proxy_next_upstream [error |timeout |invalid_header |http_500 |http_502~504 |http_404 | off]
          默认 proxy_next_upstream error timeout;
          配置块 http server location
      
          此配置表示上游一台服务器转发请求出现错误时,继续换一套服务器处理这个请求
          其参数用来说明在那些情况下继续选择下一台上游服务器转发请求.
          error 向上游发起连接 发送请求 读取响应时出错
          timeout 发送请求或读取响应时出错
          invalid_header 上游服务器发送的响应时不合法
          http_500 上游响应500
          http_502 上游响应502
          http_503 上游响应503
          http_504 上游响应504
          http_404 上游响应404
          off      关闭proxy_next_upstream功能 只要一出错就选择另外一台上游再次出发
      Nginx反向代理模块中还提供很多配置,如设置连接的超时时间,临时文件如何存储,如何缓存上游服务器响应等功能.
      ```

   可以通过阅读 ngx_http_proxy_module了解更多详细情况

   ```
               #sudo vim /usr/local/nginx/conf/nginx.conf 
   
               server {
                   listen       80;
                   server_name  localhost;
                   location / {
                   #保证代理机器能访问到 下面的机器并装有nginx  在主机号为100的机器上有响应网页
                   proxy_pass http://192.168.1.100;
                   root   html;
                   index  index.html index.htm;
                   }
               }
               sudo /usr/local/nginx/sbin/nginx -s reload
   ```

   加一些判断条件 获取到 对方请求的主机 防止别人代理到自己的主机上

## 负载均衡

upstram模块

```python
# upstream 模块负债负载均衡模块，通过一个简单的调度算法来实现客户端IP到后端服务器的负载均衡。
# 通过upstream指令指定了一个负载均衡器的名称test.com。这个名称可以任意指定，在后面需要的地方直接调用即可。
upstream test.com{
    ip_hash;
    server 192.168.123.1:80;
    server 192.168.123.2:80 down;
    server 192.168.123.3:8080  max_fails=3  fail_timeout=20s;
    server 192.168.123.4:8080;
}
```

Nginx的负载均衡模块目前支持4种调度算法:

- weight 轮询（默认）。每个请求按时间顺序逐一分配到不同的后端服务器，如果后端某台服务器宕机，故障系统被自动剔除，使用户访问不受影响。weight。指定轮询权值，weight值越大，分配到的访问机率越高，主要用于后端每个服务器性能不均的情况下。
- ip_hash。每个请求按访问IP的hash结果分配，这样来自同一个IP的访客固定访问一个后端服务器，有效解决了动态网页存在的session共享问题。
- fair。比上面两个更加智能的负载均衡算法。此种算法可以依据页面大小和加载时间长短智能地进行负载均衡，也就是根据后端服务器的响应时间来分配请求，响应时间短的优先分配。Nginx本身是不支持fair的，如果需要使用这种调度算法，必须下载Nginx的upstream_fair模块。
- url_hash。按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，可以进一步提高后端缓存服务器的效率。Nginx本身是不支持url_hash的，如果需要使用这种调度算法，必须安装Nginx 的hash软件包。

在HTTP Upstream模块中，可以通过server指令指定后端服务器的IP地址和端口，同时还可以设定每个后端服务器在负载均衡调度中的状态。常用的状态有：

down，表示当前的server暂时不参与负载均衡。

backup，预留的备份机器。当其他所有的非backup机器出现故障或者忙的时候，才会请求backup机器，因此这台机器的压力最轻。

max_fails，允许请求失败的次数，默认为1。当超过最大次数时，返回proxy_next_upstream 模块定义的错误。

fail_timeout，在经历了max_fails次失败后，暂停服务的时间。max_fails可以和fail_timeout一起使用。

注意 当负载调度算法为ip_hash时，后端服务器在负载均衡调度中的状态不能是weight和backup。

备注： nginx的worker_rlimit_nofile达到上限时，再有客户端链接报502错误. 用了log_format指令设置了日志格式之后，需要用access_log指令指定日志文件的存放路径.



负载均衡是由多台服务器以对称的方式组成一个服务器集合，每台服务器都具有等价的地位，都可以单独对外提供服务而无须其他服务器的辅助。通过某种负载分担技术，将外部发送来的请求按照事先设定分配算法分配到对称结构中的某一台服务器上，而接收到请求的服务器独立地回应客户的请求。

均衡负载能够平均分配客户请求到服务器列阵，籍此提供快速获取重要数据，解决大量并发访问服务问题。

1. upstream块

   ```
   upstream name {...}
   配置块 http
   upstream块定义一个上游服务器的集群,便于反向代理中的proxy_pass使用
   
   upstream mynet{
       server www.wopai1.com;
       server www.wopai2.com;
       server www.wopai3.com;
   }
   server {
       location /{
           proxy_pass http://mynet;
       }
   }
   ```

2. server

   ```
   server name [paramenters]
   配置块upstream
   server配置项指定了一台上游服务器的名字,可以是域名 IP地址端口 UNIX句柄
   weight= number;设置向这台服务器转发的权重,默认为1
   max_fails=number;该选项域fail_timeout配合使用
           指在fail_timeout时间段内如果转发上游失败超过number次就认为当前的fail_timeout时间内
           这台服务器不可用,max_fails默认为1 如果设置为0 表示不检查失败次数
   fail_timeout=time; fail_timeout表示该时间内转发多少次失败后就认为上游不可用.默认10s
   down    表示上游服务器永久下线,只能在ip_hash配置时才有效
   backup  在ip_hash配置时无效.只有所有非备份机都失败,才向上游备份服务器转发请求.
   upstream mynet{
       server www.wopai1.com weight=5;
       server www.wopai2.com:8081 max_fails=3 fail_timeout=300s;
       server www.wopai2.com down;
   }
   ```

3. ip_hash

   ```
   配置块 upstream
   希望来自某一个用户的请求始终落在固定的一台服务器上进行处理.
   根据客户端的IP散列计算出一个key,将key按照upstream集群中的上游服务器进行取模,求得的值对应的主机接收转发请求.
   ip_hash不可以与weight同时使用
   如果upstream配置中有一台服务器暂时不可用,不能直接删除该配置,而应该使用down标识.
   upstream mynet{
       ip_hash;
       server www.wowpai1.top;
       server www.wowpai2.top;  
       server www.wowpai3.top down;
   }
   ```

例子,服务器负载均衡基本配置,nginx中可以进行负载均衡的相关设置:

```
        upstream my.net{    #my.net是自定义的命名 在server结构中引用即可

        #代理服务器为 两台机器192.168.22.136 192.168.22.147做负载均衡操作 
        #两台机器上 可以跑apache负载功能更为强大的网页相关任务

        #max_fails 表示尝试出错最大次数 即可认为该服务器 在fail_timeout时间内不可用
        # server servername:port   servername可以写主机名 或者点分式IP
        server 192.168.22.136:80 max_fails=1 fail_timeout=300s;
        server 192.168.22.147:80 max_fails=1 fail_timeout=300s;  
        }


        server {
        listen       80;
        server_name  localhost; 
        location / {
            #upstream 块名
            proxy_pass http://my.net;
            root   html;
            index  index.html index.htm;
        }
```













<!-- TOC -->

- [理清几个术语：](#理清几个术语)
- [nginx（深入理解nginx）](#nginx深入理解nginx)
- [nginx需要的几个包](#nginx需要的几个包)
- [nginx命令入门](#nginx命令入门)
- [nginx 工作概述](#nginx-工作概述)
- [nginx 配置](#nginx-配置)
    - [1. 块配置项](#1-块配置项)
    - [2. 配置项](#2-配置项)
- [用http核心模块搭建一个静态web服务器](#用http核心模块搭建一个静态web服务器)
    - [1. 虚拟主机与请求分发](#1-虚拟主机与请求分发)
    - [2. 文件路径的定义](#2-文件路径的定义)
    - [3. 内存及磁盘资源的分配](#3-内存及磁盘资源的分配)
    - [4. 网络连接的设置](#4-网络连接的设置)
    - [5. MIME类型的设置](#5-mime类型的设置)
    - [6. 对客户端请求的限制](#6-对客户端请求的限制)
- [nginx 反向代理](#nginx-反向代理)
- [负载均衡的一般配置](#负载均衡的一般配置)

<!-- /TOC -->

# 理清几个术语：

1. IDC 互联网数据中心，主要服务包括整机租用、服务器托管、机柜租用、机房租用、专线接入和网络管理服务等。广义上的IDC业务，实际上就是数据中心所提供的一切服务。客户租用数据中心的服务器和带宽，并利用数据中心的技术力量，来实现自己对软、硬件的要求，搭建自己的互联网平台，享用数据中心所提供的一系列服务。
2. ISP--(Internet Service Provider)互联网服务提供商，即向广大用户综合提供互联网接入业务、信息业务、和增值业务的电信运营商。
3. ICP--(Internet Content Provider)互联网内容提供商，向广大用户综合提供互联网信息业务和增值业务的电信运营商。 根据中华人民共和国国务院令第292号《互联网信息服务管理办法》规定，国家对提供互联网信息服务的ICP实行许可证制度。从而，ICP证成为网站经营的许可证，经营性网站必须办理ICP证，否则就属于非法经营。因此，办理ICP证是企业网站合法经营的需要.
4. CDN--(Content Delivery Network)内容分发网络,依靠部署在各地的边缘服务器，通过中心平台的负载均衡、内容分发、调度等功能模块，使用户就近获取所需内容，降低网络拥塞，提高用户访问响应速度和命中率。CDN的关键技术主要有内容存储和分发技术。 CDN的基本原理是广泛采用各种缓存服务器，将这些缓存服务器分布到用户访问相对集中的地区或网络中，在用户访问网站时，利用全局负载技术将用户的访问指向距离最近的工作正常的缓存服务器上，由缓存服务器直接响应用户请求。
5. LVS--(Linux Virtual Server)的简写，意即Linux虚拟服务器，是一个虚拟的服务器集群系统。LVS集群采用IP负载均衡技术和基于内容请求分发技术。调度器具有很好的吞吐率，将请求均衡地转移到不同的服务器上执行，且调度器自动屏蔽掉服务器的故障，从而将一组服务器构成一个高性能的、高可用的虚拟服务器。整个服务器集群的结构对客户是透明的，而且无需修改客户端和服务器端的程序。为此，在设计时需要考虑系统的透明性、可伸缩性、高可用性和易管理性。
6. CGI--(Common Gateway Interface)通用网关接口。CGI规范允许Web服务器执行外部程序，并将它们的输出发送给Web浏览器，CGI将Web的一组简单的静态超媒体文档变成一个完整的新的交互式媒体

# nginx（深入理解nginx）

1. nginx 相比 Apache等服务器 具有更高的性能
2. nginx 高扩展性，模块化明显
3. 高可靠性
4. 低内存消耗 --- 支持高并发的基础
5. 高并发
6. 热部署

# nginx需要的几个包

1. GCC编译器  c编译器
> yum install -y gcc 
2. g++编译器 c++编译器
> yum install -y gcc-c++ 
3. pcre库 nginx用它来解析正则表达式
> yum install pcre pcre-devel
4. zlib库 用来做gzip格式的压缩
> yum install zlib zlib-devel
5. openSSl开发库 支持ssl协议传输http
> yum install openssl openssl-devel


在nginx官网下载源码包，然后解压。
> wget http://nginx.org/download/nginx-1.14.0.tar.gz
解压
> tar -zxvf nginx-1....

然后编译安装 nginx
```sh
./configure
make
make install
```
linux编译安装中configure、make和make install各自的作用

./configure是用来检测你的安装平台的目标特征的。比如它会检测你是不是有CC或GCC，并不是需要CC或GCC，它是个shell脚本。
make是用来编译的，它从Makefile中读取指令，然后编译。
make install是用来安装的，它也从Makefile中读取指令，安装到指定的位置。
AUTOMAKE和AUTOCONF是非常有用的用来发布C程序的东西。

1. configure
这一步一般用来生成 Makefile，为下一步的编译做准备，你可以通过在 configure 后加上参数来对安装进行控制，比如代码:
./configure --prefix=/usr

上面的意思是将该软件安装在 /usr 下面，执行文件就会安装在 /usr/bin （而不是默认的 /usr/local/bin)，资源文件就会安装在 /usr/share（而不是默认的/usr/local/share）。

同时一些软件的配置文件你可以通过指定 --sys-config= 参数进行设定。有一些软件还可以加上 --with、--enable、--without、--disable 等等参数对编译加以控制，你可以通过允许 ./configure --help 察看详细的说明帮助。
2. make
这一步就是编译，大多数的源代码包都经过这一步进行编译（当然有些perl或Python编写的软件需要调用perl或python来进行编译）。

如果 在 make 过程中出现 error ，你就要记下错误代码（注意不仅仅是最后一行），然后你可以向开发者提交 bugreport（一般在 INSTALL 里有提交地址），或者你的系统少了一些依赖库等，这些需要自己仔细研究错误代码。

make 的作用是开始进行源代码编译，以及一些功能的提供，这些功能由他的 Makefile 设置文件提供相关的功能，比如 make install 一般表示进行安装，make uninstall 是卸载，不加参数就是默认的进行源代码编译。

make 是 Linux 开发套件里面自动化编译的一个控制程序，他通过借助 Makefile 里面编写的编译规范进行自动化的调用 gcc 、ld 以及运行某些需要的程序进行编译的程序。一般情况下，他所使用的 Makefile 控制代码，由 configure 这个设置脚本根据给定的参数和系统环境生成。
3. make install
这条命令来进行安装（当然有些软件需要先运行 make check 或 make test来进行一些测试），这一步一般需要你有 root 权限（因为要向系统写入文件）

这样做完之后nginx会默认安装在/usr/local/nginx下面

# nginx命令入门
1. /nginx/sbin/nginx  默认读取 /nginx/conf/nginx.conf文件  -c可以指定文件
2. -t 测试文件是否有误
3. -s stop 强制停止
4. -s quit 优雅的停止，结束所有的任务连接后退出
5. -s reload 重新加载
6. -s reopen 重新打开日志文件 ，可以用来更换日志文件




# nginx 工作概述

部署nginx时都是使用一个master进程来管理多个worker进程。一般worker进程数 是等于CPU核数

在Apache上每个进程在一个时刻只处理
一个请求，因此，如果希望Web服务器拥有并发处理的请求数更多，就要把Apache的进程或
线程数设置得更多，通常会达到一台服务器拥有几百个工作进程，这样大量的进程间切换将
带来无谓的系统资源消耗。而Nginx则不然，一个worker进程可以同时处理的请求数只受限于
内存大小，而且在架构设计上，不同的worker进程之间处理并发请求时几乎没有同步锁的限
制，worker进程通常不会进入睡眠状态，因此，当Nginx上的进程数与CPU核心数相等时（最
好每一个worker进程都绑定特定的CPU核心），进程间切换的代价是最小的。

# nginx 配置

## 1. 块配置项

块配置项由一个配置项名和一对大括号组成,块配置项可以嵌套，内层块直接继承外层块。
```sh
events {

}
http {
    upstream backend{
        server 127.0.0.1:8080;
    }
    gzip on;
    server{
        location /webstatic {
            gzip off
        }
    }
}
```

## 2. 配置项
默认配置项：
1. daemon on|off 
> daemon on 守护进程
2. master_process on|off
> master_process on 以master方式管理进程
3. error_log pathfile level
> error_log logs/error.log error   
level取值 debug info notice warn error crit alert emerg
关闭日志的唯一手段 error_log /dev/null
4. debug_points [stop|abort] 设置断点调试
5. debug_connection [ip|cidr] 只能放在事件类模块中，它是给指定的ip或cidr地址 输出debug级别的日志

正常运行的配置项
1. 定义环境变量
> env VAR|VAR=VALUE   
2. 嵌入其他配置文件
> include pathfile
3. pid 文件路径
> pid path/file  

默认
> pid logs/nginx.pid
4. Nginx worker进程运行的用户及用户组,指定worker进程运行在那个用户及用户组下面
> user username[groupname]  

默认
> user nobody nobody
5. 指定Nginx worker进程能够打开的最大句柄描述符个数
> worker_rlimit_nofile limit
6. 限制信号队列  ,设置每个用户发往Nginx的信号队列大小，如果摸个用户的信号队列满了，再发就会被丢弃
> worker_rlimit_sigpending limit

优化性能的一些配置
1. Nginx worker 进程个数
> worker_processes number
2. ssl 硬件加速
> ssl_engine device
3. 系统调用gettimeofday的执行频率 ,, 一般不用配置
> timer_resolution t
4. worker进程的优先级设置
> worker_priority 0;

5. accept锁延迟时间
> accept_mutex_delay 500ms

6. 批量建立新连接
> multi_accept [on|off]

7. 选择时间模型
> use [kqueue|rtsig|epoll|/dev/poll|select|poll|eventport]

Linux会自动选择epoll，select，poll中最好的
8. 每个worker最大的连接数
> worker_connections number;



事件类配置项
1. 是否打开accept锁
> accept_mutext on;

它是Nginx的负债均衡锁，安排所有worker进程与客户端建立tcp连接，尽量让每一个worker进程建立的连接数量一样
2. lock文件路径
> lock_file path/file

如果Nginx不支持原子锁，所以会让accept锁去实现accept锁
3. 


一些规则：
1. 每个配置项结束要有；
2. 如果配置项中包括语法符号，比如空格符，要使用单引号或者双引号阔住配置项
3. 注释#
4. 大小单位 k M
5. 时间单位ms s m h d w M y

# 用http核心模块搭建一个静态web服务器

所有的http配置项都必须*直属于* http块，server块，location块，upstream块等。。

配置项分成8类进行阐述
## 1. 虚拟主机与请求分发

一般一个ip对应多个主机域名，nginx按照server_name（server块）来定义虚拟主机，每个server块就是一个虚拟主机，它只处理与之相对应的主机域名请求。这样，一台服务器上的Nginx就能以不同的方式处理访问不同主机域名的HTTP请求了。
1. 监听端口
> listen address:port [default(deprecated in 0.8.21)|default_server]  

> 后面能够接的参数[backlog=num|rcvbuf=size|sndbuf=size|accept_filter=filter|deferred|bind|ipv6only=[on|off]|ssl]]; 

> listen 80; 默认的
2. 主机名称
> server_name name[...];

默认
> server_name "";

例如
> server_name www.testweb.com

3. location
> location [=|~|~*|^~|@]/uri/{...}

+ location会尝试根据用户请求中的URI来匹配上面的/uri表达式，如果可以匹配，就选择location{}块中的配置来处理用户请求。当然，匹配方式是多样的，下面介绍location的匹配规则。
    - =表示把URI作为字符串，以便与参数中的uri做完全匹配。
    - ~表示匹配URI时是字母大小写敏感的。
    - ~*表示匹配URI时忽略字母大小写问题
    - ^~表示匹配URI时只需要其前半部分与uri参数匹配即可。 *默认是这种方式*
    - @表示仅用于Nginx服务内部请求之间的重定向，带有@的location不直接处理用户请求

## 2. 文件路径的定义

下面简绍一下文件路径的配置项
1. 以root方式设置资源路径
> root path;

默认 root html

例如
```sh
location /download/ {
    root document;
}
如果有一个请求的url是/download/index/test.html,那么web服务器将会返回服务器上 document/download/index/test.html的内容
```
2. 以alias方式设置资源路径
> alias path;
```sh
location conf {
    alias usr/local/nginx/conf/;
}
如果一个url请求是/conf/nginx.conf,那么实际返回的文件是/usr/local/nginx/conf/nginx.conf
```
请仔细对比它与root的区别

同时，alias后面还可以跟正则表达式
```sh
location ~^/test/(\w+)\.(\w+)$ {
    alias usr/local/nginx/$2/$1.$2;
}
/test/nginx.conf 会返回usr/local/nginx/conf/nginx.conf
```
3. 访问首页
> index file ...;

默认 index index.html
```sh
location {
    root path;
    index index.html html/index.php /index.php;
}
依次尝试 ，指导访问成功其中之一
```
4. 根据http 返回码重定向页面
> error_page code[code....][=|=answer-code]uri|@named_location
```sh
error_page 404 404.html;
error_page 502 503 504 50x.html;
error_page 403 http://example.com/forbidden.html;
error_page 404 @fetch;
```
更改错误码
> error_page 404 =200 empty.gif'
> error_page 404 =403 forbidden.gif;

## 3. 内存及磁盘资源的分配

1. http包体值存储到磁盘文件中
> client_body_in_file_only on|clean|off
> 默认 off

2. http包体尽量写入一个内存buffer中
> client_body_in_single_buffer on|off
> 默认 off
3. 存储http头部的内存大小
> client_header_buffer_size size;
> client_header_buffer_size 1k 默认
4. 存储超大http头部的内存buffer大小
> large_client_header_buffers number size
> 默认 48k
5. client_body_buffer_size size 
6. client_body_temp_path dir-path[level1[level2[level3]]]
7. connection_pool_size
8. request_pool_size


## 4. 网络连接的设置
1. 读取http头部超时时间
> client_header_timeout time  

超时就返回408
2. 读取包体的超时时间
> client_body_timeout 60
3. send_timeout time
4. reset_timeout_connection on|off
5. lingering_close on表示关闭的时候一定要 处理完数据
6. lingering_time time 当上传时间超过时，自动断开连接
7. lingering_timeout time 
8. keepalive_disable[msie6|safari|none] 对某些浏览器禁用该功能
9. keepalive time
10. keepalive n 一个长连接上允许承载的请求最大数


## 5. MIME类型的设置
> type {...};

定义MIME type 到文件扩展名的映射。
```sh
types {
    text/html html;
    text/html conf;
    image/gif gif;
    image/jpeg jpg;
}
```

默认MIME type 当找不到相应的类型 匹配时，用这个
> default_type MIME-type;
> 默认 default_type text/plain;

设置每个散列桶的内存大小
> types_hash_bucket_size 32|64|128

## 6. 对客户端请求的限制

1. 按http方法名限制用户的请求
```sh
limit_except method ...{...}

limit_except GET {
    allow 192.111.111.111/32;
    deny all;
}
```
2. http请求的包体最大值
> client_max_body_size size;
3. 队请求的限速
> limit_rate speed;


设置DNS名字解析服务器的地址
> resolver 127.0.0.1 192.0.2.1;


# nginx 反向代理

反向代理（reverse proxy）方式是指用代理服务器来接受Internet上的连接请求，然后将
请求转发给内部网络中的上游服务器，并将从上游服务器上得到的结果返回给Internet上请求
连接的客户端，此时代理服务器对外的表现就是一个Web服务器。

由于Nginx具有“强悍”的高并发高负载能力，因此一般会作为前端的服务器直接向客户
端提供静态文件服务。但也有一些复杂、多变的业务不适合放到Nginx服务器上，这时会用
Apache、Tomcat等服务器来处理。于是，Nginx通常会被配置为既是静态Web服务器也是反向
代理服务器


1. proxy_pass
> proxy_pass URL;

+ 几种用法
    - 此配置项将当前的请求反向代理到url参数指定的服务器上，url可以是主机名或者ip地址加端口的形式
    > proxy_pass http://localhost:8000/url/
    - 也可以是Unix句柄
    > proxy_pass http://unix:/path/to/backend.socket:uri/
    - 也可以使用负载均衡中
    ```sh
    upstream backend{

    }
    server {
        location / {
            proxy_pass http://backend
        }
    }
    当然用户可以把http转换成更加安全的https
    proxy_pass https://111.111.111.111
    ```
2. proxy_method 
设定转发的方法 例如POST

3. proxy_hide_header
指定不转发的头部 Date Server X-pad等

4. proxy_pass_header

5. proxy_pass_request_body  是否转发体

6. proxy_pass_request_headers 是否转发头

7. proxy_redirect
> proxy_redirect http://localhost:8000/two/  http://frontend/one

例如把 http://localhost:8000/two/some/url/  http://frontend/one/some/url/

8. 

# 负载均衡的一般配置
作为代理服务器，一般都需要向上游服务器的集群转发请求。这里的负载均衡是指选择一种策略，尽量把请求平均地分布到每一台上游服务器上。
1. upstream块
```sh
upstream backend {
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com;
}
server {
    location / {
        proxy_pass http://backend;
    }
}
```
2. server 
> server name[parameters];

+ 配置在upstream里面表示 指定了一台上游服务器的名字，这个名字可以是域名，ip地址，Unix句柄，在它的后面还可以接参数
    - weight=number：设置向这台上游服务器转发的权重，默认为1。
    - max_fails=number：该选项与fail_timeout配合使用，指在fail_timeout时间段内，如果向当前的上游服务器转发失败次数超过number，则认为在当前的fail_timeout时间段内这台上游服务器不可用。max_fails默认为1，如果设置为0，则表示不检查失败次数。
    - fail_timeout=time：fail_timeout表示该时间段内转发失败多少次后就认为上游服务器暂时不可用，用于优化反向代理功能。它与向上游服务器建立连接的超时时间、读取上游服务器的响应超时时间等完全无关。fail_timeout默认为10秒。
    - down：表示所在的上游服务器永久下线，只在使用ip_hash配置项时才有用。
    - backup：在使用ip_hash配置项时它是无效的。它表示所在的上游服务器只是备份服务器，只有在所有的非备份上游服务器都失效后，才会向所在的上游服务器转发请求。

3. ip_hash

> ip_hash

它首先根据客户
端的IP地址计算出一个key，将key按照upstream集群里的上游服务器数量进行取模，然后以取
模后的结果把请求转发到相应的上游服务器中。这样就确保了同一个客户端的请求只会转发
到指定的上游服务器中。

注意：ip_hash 与weight（权重）不可同时使用
