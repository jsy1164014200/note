# acme.sh let's Encrypt 配置https证书

ACME 全称是 Automated Certificate Management Environment，直译过来是自动化证书管理环境的意思，Let's Encrypt 的证书签发过程使用的就是 ACME 协议。有关 ACME 协议的更多资料可以在[这个仓库](https://github.com/ietf-wg-acme/acme/)找到。



Let's Encrypt作为一个公共且免费SSL的项目逐渐被广大用户传播和使用，是由Mozilla、Cisco、Akamai、IdenTrust、EFF等组织人员发起，主要的目的也是为了推进网站从HTTP向HTTPS过度的进程，目前已经有越来越多的商家加入和赞助支持。

Let's Encrypt免费SSL证书的出现，也会对传统提供付费SSL证书服务的商家有不小的打击。到目前为止，Let's Encrypt获得IdenTrust交叉签名，这就是说可以应用且支持包括FireFox、Chrome。

Let's Encrypt 的最大贡献是它的 ACME 协议，第一份全自动服务器身份验证协议，以及配套的基础设施和客户端。这是为了解决一直以来 HTTPS TLS X.509 PKI 信任模型，即证书权威（Certificate Authority, CA）模型缺陷的一个起步。

在客户端-服务器数据传输中，公私钥加密使得公钥可以明文传输而依然保密数据，但公钥本身是否属于服务器，或公钥与服务器是否同属一个身份，是无法简单验证的。证书权威模型通过引入事先信任的第三方，由第三方去验证这一点，并通过在服务器公钥上签名的方式来认证服务器。第三方的公钥则在事先就约定并离线准备好，以备访问时验证签名之用。这个第三方就称为证书权威，简称CA。相应的，CA验证过的公钥被称为证书。



## 配置步骤

### 1. 下载acme.sh

> curl https://get.acme.sh | sh

把 acme.sh 安装到你的 home 目录下:~/.acme.sh/并创建 一个 bash 的 alias, 方便你的使用:

>  acme.sh=~/.acme.sh/acme.sh  # 写到 ./bashrc 里面 source一下 就可以长期使用，或者alias 只在本次登录使用

安装过程不会污染已有的系统任何功能和文件, 所有的修改都限制在安装目录中: ~/.acme.sh/

acme.sh 实现了 acme 协议支持的所有验证协议. 一般有两种方式验证: http 和 dns 验证.

### 2. 添加解析，生成证书

#### http 方式

需要在你的网站根目录下放置一个文件, 来验证你的域名所有权,完成验证. 然后就可以生成证书了.

>  acme.sh --issue -d [http://mydomain.com](http://link.zhihu.com/?target=http%3A//mydomain.com) -d [http://www.mydomain.com](http://link.zhihu.com/?target=http%3A//www.mydomain.com) --webroot /home/wwwroot/[http://mydomain.com/](http://link.zhihu.com/?target=http%3A//mydomain.com/)

只需要指定域名, 并指定域名所在的网站根目录. acme.sh 会全自动的生成验证文件, 并放到网站的根目录, 然后自动完成验证. 最后会聪明的删除验证文件. 整个过程没有任何副作用.

如果你用的 apache服务器, acme.sh 还可以智能的从 apache的配置中自动完成验证, 你不需要指定网站根目录:

>  acme.sh --issue -d [http://mydomain.com](http://link.zhihu.com/?target=http%3A//mydomain.com) --apache

如果你用的 nginx服务器, 或者反代, acme.sh 还可以智能的从 nginx的配置中自动完成验证, 你不需要指定网站根目录:

> acme.sh --issue -d [http://mydomain.com](http://link.zhihu.com/?target=http%3A//mydomain.com) --nginx

注意, 无论是 apache 还是 nginx 模式, acme.sh在完成验证之后, 会恢复到之前的状态, 都不会私自更改你本身的配置. 好处是你不用担心配置被搞坏, 也有一个缺点, 你需要自己配置 ssl 的配置, 否则只能成功生成证书, 你的网站还是无法访问https. 但是为了安全, 你还是自己手动改配置吧.

如果你还没有运行任何 web 服务, 80 端口是空闲的, 那么 acme.sh 还能假装自己是一个webserver, 临时听在80 端口, 完成验证:

> acme.sh --issue -d [http://mydomain.com](http://link.zhihu.com/?target=http%3A//mydomain.com) --standalone

#### dns 方式

1. 方式一：使用 token

`acme.sh` 支持直接使用主流 DNS 提供商的 API 接口来完成域名验证以及一些相关操作

具体 [dnsapi 链接](https://github.com/Neilpang/acme.sh/tree/master/dnsapi)

这里以 阿里云 为例：

首先获取你的阿里云API Key: <https://ak-console.aliyun.com/#/accesskey>

之后在你的终端配置文件中设置：

```
export Ali_Key="sdfsdfsdfljlbjkljlkjsdfoiwje"
export Ali_Secret="jlsdflanljkljlfdsaklkjflsa"
```

之后直接使用如下命令发起申请：

```
acme.sh --issue --dns dns_ali -d example.com -d *.example.com 
```

`Ali_Key` 和 `Ali_Secret` 将被保存在 `~/.acme.sh/account.conf` , 命令中 **dns_ali** 指明使用 阿里的dns

来生成证书，注意这里第一个域名为顶级域名，后面个为泛域名。

> 这种方式将自动为你的域名添加一条 `txt` 解析，验证成功后，这条解析记录会被删除，所以对你来说是无感的，就是要等 `120秒`。

证书生成成功后，默认保存在 `.acme.sh/hostname` 中。

> 若想自定义证书目录，可加上 -w 参数
>
> ```
> acme.sh --issue --dns dns_ali -d *.example.com -w /etc/letsencrypt/*.example.com
> ```

2. 方式二：添加一条 txt 解析记录

命令：

```
acme.sh  --issue  --dns -d example.com -d *.example.com
```

> 需要同时添加裸域名及泛域名。注意要将非泛域名的域名放在前面，否则可能会遇到一些问题。

然后, [acme.sh](http://acme.sh/) 会生成相应的解析记录显示出来, 示例如下:

```
Multi domain='DNS:bitcat.cc,DNS:*.example.com'
Getting domain auth token for each domain
Getting webroot for domain='example.com'
Getting webroot for domain='*.example.com'
Add the following TXT record:
Domain: '_acme-challenge.example.com'
TXT value: '<ACME_CHALLENGE_STRING>'
Please be aware that you prepend _acme-challenge. before your domain
so the resulting subdomain will be: _acme-challenge.example.com
Please add the TXT records to the domains, and retry again.
Please add '--debug' or '--log' to check more details.
See: https://github.com/Neilpang/acme.sh/wiki/How-to-debug-acme.sh
```

记录下其中的 `<ACME_CHALLENGE_STRING>` 并前往你的 DNS 服务提供商，为主机名 `_acme-challenge` 添加一条 TXT 记录，内容即为上述的 `<ACME_CHALLENGE_STRING>`。提交后可以等待一小段时间以便让 DNS 生效。

重新申请签发证书:

```
acme.sh --renew --dns -d example.com -d *.example.comacme.sh
```

注意第二次这里用的是 `--renew`



### nginx使用证书

这里仅用 nginx 服务器配置做示例：

nginx 配置文件重点介绍：

- Nginx 的配置 `ssl_certificate`  和 `ssl_trusted_certificate` 使用 `fullchain.cer` ，而非 `<domain>.cer` ，否则 [SSL Labs](https://www.ssllabs.com/ssltest/) 的测试会报 `Chain issues Incomplete` 错误

- `ssl_dhparam` 通过下面命令生成：

  ```
  sudo mkdir /etc/nginx/ssl
  sudo openssl dhparam -out /etc/nginx/ssl/dhparam.pem 2048
  ```

nginx.conf 配置示例：

```sh
server {
    listen 80;
    server_name mark.example.com;

    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    server_name mark.example.com;

    ssl_certificate /etc/letsencrypt/live/*.example.com/fullchain.cer;
    ssl_certificate_key /etc/letsencrypt/live/*.example.com/*.example.com.key;

    # disable SSLv2
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;

    # ciphers' order matters
    ssl_ciphers "ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA384:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!aNULL";

    # the Elliptic curve key used for the ECDHE cipher.
    ssl_ecdh_curve secp384r1;

    # use command line
    # openssl dhparam -out dhparam.pem 2048
    # to generate Diffie Hellman Ephemeral Parameters
    ssl_dhparam /etc/nginx/ssl/dhparam.pem;

    # let the server choose the cipher
    ssl_prefer_server_ciphers on;

    # turn on the OCSP Stapling and verify
    ssl_stapling on;
    ssl_stapling_verify on;
    ssl_trusted_certificate /etc/letsencrypt/live/*.example.com/fullchain.cer;

    # http compression method is not secure in https
    # opens you up to vulnerabilities like BREACH, CRIME
    gzip off;

    location / {
        root /mnt/var/www/tofar/mark.example.com;
        index index.html;
    }

    error_log  /mnt/log/nginx/mark.example.com/error.log;
    access_log /mnt/log/nginx/mark.example.com/access.log;
}
```



之后重启 nginx 即可：



```
sudo nginx -s reload
or sudo openresty -s reload # 若安装的是openresty
```