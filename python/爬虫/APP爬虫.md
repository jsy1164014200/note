# APP爬虫

安装 appium ，一个phone的自动化测试工具

> sudo cnpm install -g appium  // 这里使用cnpm，因为npm貌似会出问题

CA证书的安装

　　要捕获https证书，就得解决证书认证的问题，因此需要在通信发生的客户端安装证书，并且设置为受信任的根证书颁布机构。下面介绍6种客户端的安装方法。

　　当我们初次运行mitmproxy或mitmdump时，

　　会在当前目录下生成 ~/.mitmproxy文件夹，其中该文件下包含4个文件，这就是我们要的证书了。

　　mitmproxy-ca.pem 私钥
　　mitmproxy-ca-cert.pem 非windows平台使用
　　mitmproxy-ca-cert.p12 windows上使用
　　mitmproxy-ca-cert.cer 与mitmproxy-ca-cert.pem相同，android上使用

1. Firefox上安装
   preferences-Advanced-Encryption-View Certificates-Import (mitmproxy-ca-cert.pem)-trust this CA to identify web sites
2. chrome上安装
   设置-高级设置-HTTPS/SSL-管理证书-受信任的根证书颁发机构-导入mitmproxy-ca-cert.pem
3. osx上安装
   双击mitmproxy-ca-cert.pem - always trust
4. windows7上安装
   双击mitmproxy-ca-cert.p12-next-next-将所有的证书放入下列存储-受信任的根证书发布机构
5. iOS上安装
   将mitmproxy-ca-cert.pem发送到iphone邮箱里，通过浏览器访问/邮件附件
6. Android上安装
   将mitmproxy-ca-cert.cer 放到sdcard根目录下选择设置-安全和隐私-从存储设备安装证书