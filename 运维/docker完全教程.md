# 践行devops（docker 完全教程）

> 作为一个ITer ，开发 测试 运维 ，都必备 docker



kubernetes（k8s）  docker

## 历史

1. 部署慢
2. 成本高
3. 难迁移扩展
4. 资源浪费

虚拟机

1. 一个app一个virtual machine
2. 资源化
3. 容易扩展
4. 容易云化（阿里云）

虚拟机局限性

1. 一个虚拟机是一个操作系统，本身消耗的资源很大
2. 开发与运维的矛盾

容器化

1. 对软件及其依赖的标准化打包
2. 应用的隔离
3. 共享一个os kernel
4. 运行在很多os上

docker-engine

docker-cli

docker-compose

docker-machine



## docker-machine

用来连接远程服务器上的docker，可以在阿里云上创建一个 有docker的虚拟机

## docker-engine

1. 后台进程（dockerd）
2. restful api server
3. cli 接口

管理 network volumn image container 

底层技术支持（linux系统原有的）

1. namespace ： 做隔离，pid ，net， ipc，mut，uts
2. control groups ：做资源限制
3. union file  systems ：container和image的分层

## docker 镜像

1. 文件和meta data的集合（root filesystem）
2. 分层的，并且每一层都可以添加改变，删除文件，成为一个新的image
3. 不同的image可以共享 layer
4. image 本身是read-only的

base image  指Ubuntu centos debian等 在linux Kernel （bootfs）之上的 一层rootfs

命令

- docker image ls  列出所有镜像
- docker pull Ubuntu:14.04  从registry上拉取image
- docker build -t tag . (. 不能少)
- docker history id 
- docker image rm id  同 docker rmi id

docker image build || docker build 

生成临时container 用来创建container ，然后 commit成 image

## docker container

- 通过image构建（copy）
- 在imagelayer 之上 建立了一个container layer 可读写
- 就好比 原型与实例
- image 负责app 的存储和分发，container负责运行app

命令

- docker run -it xxx

- docker container ls 当前在运行的所有容器
- docker container ls -a 所有的容器
- docker container rm id
- docker ps -a  同 docker  container ls -a
- docker rm id  默认是删除 container
- docker container ls -aq 列出所有id
- docker rm (docker container ls -aq)  删除所有的container
- docker rm (docker container ls -f "status=exited" -q)

docker container commit  || docker  commit  这种 方式不提倡（不安全）

能够将 container commit 成image

docker commit 【Options】container repository:tag 

docker exec 命令 

docker exec -it  container_id /bin/bash [ip a 打印端口ip]

docker stop name|container_id	

docker start name | container_id

docker inspect id

## Dockerfile

1. FROM base-image

2. LABEL “Metadata” 

3. RUN xxx && xxx（避免多次RUN ，尽量 把多个RUN 合成一个 用\换行）

   执行命令并且创建新的 image layer

4. WORKDIR /root   没有工作目录会自动创建 （用workdir 不要RUN cd，尽量使用绝对目录）

5. ADD 和COPY

   > ADD hello /
   >
   > ADD test.tar.gz /  # 添加到根目录并且解压缩

   > WORKDIR /root   
   >
   > ADD hello test/   # /root/test/hello

   > WORKDIR /root
   >
   > COPY hello test/

   大部分情况，copy 优与ADD ，ADD 除了 COPY还有额外的功能（解压）

   添加远程文件/目录 请使用 curl  或者 wget

6. ENV MYSQL_VERSION 5.7 设置常量

   > 尽量多使用 ，增加可维护性

7. VOLUME 

8. EXPOSE

9. CMD

   1. 设置容器启动后 默认执行的 命令和参数
   2. 如果在docker run 的时候指定了其他命令，那么 CMD 命令会被忽略
   3. 如果定义了 多个CMD，只有最后一个被执行

10. ENTRYPOINT

    1. 设置容器启动时运行的命令

    2. 不会被忽略，一定会执行

    3. 最佳实践：写一个shell脚本作为entrypoint

       ```dockerfile
       COPY docker-entrypoint.sh /usr/local/bin/
       ENTRYPOINT ["docker-entrypiont.sh"]
       ```

       EXPOSE 27017
       CMD ["mongod"]

> 注： 区分shell格式 与 exec 格式
>
> shell  RUN/CMD/ENTRYPOINT  echo "haha"
>
> exec ["echo","haha"]
>
> 但是 shell 是使用的 bash 执行 命令
>
> exec 是单独的执行所以 如果要用$必须指定bash
>
> ["/bin/bash","-c","echo hello $name"]





## 远程仓库（docker hub）

docker login 

docker push 用户名/xxx

docker pull 

docker logout

## 搭建使用私有的 dockerhub

> 在远程服务器执行 docker run -d -p 5000:5000 --restart always --name registry rCistry:2 

要该名字为 ip/xxxx

docker push ip/xxx 

1. 在 push 到 自己的私有仓库时要 改本机/ect/docker/daemon.json

   {"insecure-registries":["ip"]}

2. sudo systemctl restart docker  

## debug Dockerfile

docker run -it xxx /bin/bash 进入中断的地方



## Docker Network

network namespace

veth 相当于网口 一对 （一个在namespace1，一个在namespace2

尝试下面的网络命令

ip a

sudo ip netns list

sudo ip netns delete test

sudo ip netns add test

sudo ip netns exec test ip a

sudo ip netns exec test ip link set dev lo up # ip link 能对 device 的 MTU 及 MAC 等进行设置，也能 up or down 某个网络端口

实验：

添加一对veth 接口

sudo ip netns add test1

sudo ip netns add test2

sudo ip link add veth-test1 type veth peer name veth-test2

sudo ip link set veth-test1 netns test1

sudo ip link set veth-test2 netns test2

添加ip

sudo ip netns exec test1 ip addr add 192.168.1.1/24 dev veth-test1

sudo ip netns exec test2 ip addr add 192.168.1.1/24 dev veth-test2

sudo ip netns exec test1 ip link set dev veth-test1 up

sudo ip netns exec test2 ip link set dev veth-test2 up



docker 网络

docker 新建一个container 时会 默认使用 Bridge Network，当然可以换

docker network create -d my-bridge

docker network connect my-birdge  xxx

自建的网络与 Bridge0不同的在于 自带了link



docker network ls

单机

1. Bridge Network

   ![选区_088](/home/jsy/图片/选区_088.png)

   访问外网

   ![选区_089](/home/jsy/图片/选区_089.png)

2. Host Network

   docker run  -d --name test1 --network host busybox /bin/sh -c "while true; do sleep 3600; done"

3. None Network

   docker run  -d --name test2 --network none busybox /bin/sh -c "while true; do sleep 3600; done"

   防止别人连接，保证安全



### link参数

docker run  -d --name test2 --link test1 busybox /bin/sh -c "while true; do sleep 3600; done"

docker exec -it test2 /bin/sh

ping test1 就能够ping通 test1



docker run --name redis redis

docker run -d --link redis --name flask-redis-dockerfile -e REDIS_HOST=redis flask-redis-dockerfile

```python
from flask import Flask
from redis import Redis
import os
import socket

app = Flask(__name__)
redis = Redis(host=os.environ.get("REDIS_HOST","127.0.0.1"), port=6379)


@app.route("/")
def hello():
    redis.incr("hits")
    return "hello docker,{},{}".format(redis.get('hits'),socket.gethostname())


if __name__ == "__main__":
    app.run(host="0.0.0.0",port=5000,debug=True)
```

docker file

```dockerfile
FROM python
LABEL maintainer="GNG<gng@bingyan.net>"
COPY . /app
WORKDIR /app 
RUN pip install flask redis
EXPOSE 5000
CMD ["python","app.py"]

```

### docker network 之多机通信

多机

Overlay Network

vxlan（包了一层）



## docker 数据持久化 挂载

1. 本地文件存储
2. plugin 第三方的存储方案



使用volume使用 bind mounting

docker log ls  可以用来列出所有执行出错的 log

docker volume ls  可以列出 挂载目录

docker volume ls 

docker volume rm 

如果要使用 数据库的 持久化 请参考 library docker 

docker run -v mysql:/var/lib/mysql

在本地就使用 docker 开发 环境（这样不用担心部署了）

再link 上 数据库 这样就完全不用 担心本地与 服务器的差异了 

> docker run -d -p xxx:xxx -v (pwd):/project --name web image_name 



## docker-compose 批处理

mysql: 8.0 貌似有bug 一直无法连接

手动部署 多容器太麻烦

yml文件来管理

1. services

   一个service就是一个容器 （从 dockerhub来，或者自己 build）

   service启动类似 docker run  而我们可以给他 指定 network  volume，

2. network

3. volume

命令

1. docker-compose up 

2. stop

3. down

4. **scale service=num  负载均衡,可以 用来 增加 缩小 应用**（但是不能 跨机器）

   ```dockerfile
   docker-compose up --scale web=3 -d 三个一样的web service 访问同一个 数据库
   
   另外要更改
   lb: 
   	image: dockercloud/haproxy
   	links:
   		- web
   	ports:
   		- 8080:80
   ```

5. docker-compose build  之后在 up 



## docker Swam

![选区_090](/home/jsy/图片/选区_090.png)

1. service  一个service 代表一个容器
2. Replicas   表示每个 service 扩展的个数  类似docker-compose --scale

https://labs.play-with-docker.com/p/bh9u80spkul000ehkd1g#bh9u80sp_bh9ub6r0mkfg0080s6i0

可以用来做实验

步骤： 先运行一个 manager

docker swarm init --advertise-addr=192.168.205.10

> balabala join token

docker swarm join --token xxxxxxxxxxxxxx 192.168.205.10:2377

> this node join as a worker 

docker node ls 显示出所有node节点



docker service create [options] image [arg...]

docker service ls

docker service ps 

docker service scale



网络

docker network create -d overlay demo

docker service create --name mysql --env MYSQL_ROOT_PASSWORD=root --env MYSQL_DATABASE=wordpress --network demo --mount type=volume,source=mysql-data,destination=/var/lib/mysql mysql

docker service create --name wordpress -p 80:80 --network demo --env WORDPRESS_DB_PASSWORD=root --env WORDPRESS_DB_HOST=mysql wordpress

这样在所有的 cluster 节点都能访问到

![选区_091](/home/jsy/图片/选区_091.png)

![选区_092](/home/jsy/图片/选区_092.png)

![选区_093](/home/jsy/图片/选区_093.png)

![选区_094](/home/jsy/图片/选区_094.png)









# docker 企业版

docker cloud 

1. 提供容器管理，编排，部署的托管服务

![选区_095](/home/jsy/图片/选区_095.png)



tox  单元测试 









# Kubernetes K8S

![选区_096](/home/jsy/图片/选区_096.png)



![选区_097](/home/jsy/图片/选区_097.png)









周三班长会



168计划 ： 人数到期，不允许无辜到客。宿舍。

未毕业率 1%

60%的深造率

优良学风班比率80%



能力提升课程



文

1. 读书坊

体

