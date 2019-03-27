# 给远程服务器配置ssh key 免密码登录

## SSH介绍

SSH全称Secure Shell，也称Socket Socket Shell，是一种网络协议，他给管理员提供一种安全的方式访问远程计算机，是一种公钥加密方式。在ssh protocol version 2中提供5种类型密钥，分别是：

- RSA
- RSA1
- DSA
- ECDSA
- ED25519

使用`ssh-keygen`默认生成rsa密钥对



## ssh连接原理

1. USER1发起SSH请求

2. 服务器生成随机数R1发给USER1，USER1用私钥加密生成R2。

3. USER1把R2回发给服务器，服务器用公钥解密并对比R1，相同则成功连接。

## ssh 登录步骤

1. 在本机生成 密钥对

   > ssh-keygen -t rsa  (默认就是rsa加密)

   - 输入文件 默认 /home/xxx/.ssh/id_rsa
   - 输入使用密码（防止别人用你的电脑登录）
   - 公钥文件在/home/xxx/.ssh/id_rsa.pub中

2. 将公钥 复制到 服务器的认证列表中

   服务器认证列表是一个文件，可以理解为<存储用户SSH公钥的地方>，因为SSH是一个验证过程，所以服务器需要事先保存对方的公钥，这样管理员就可以指定哪些用户(准确说是密钥对)可以登录了。

   在服务器的配置文件`/etc/ssh/sshd_config`中记录的着认证列表的目录。

   文件中的`#AuthorizedKeysFile     %h/.ssh/authorized_keys`表示认证列表文件

   所以只要将 在本机生成的公钥 复制上去就可以了

3. 登录

   ```bash
   >>> ssh root@111.111.111.111 -p 22 -i ~/.ssh/id_rsa
   ```

4. 设置服务器不允许密码登录

   在服务器配置文件`/etc/ssh/sshd_config`里小小的设置一下

   把其中的`PasswordAuthentication`中的yes改成no就不可再用密码登陆了



小bug

```
chmod 600 ~/.ssh/id_rsa
# 更改 rsa文件的权限
ssh-add 
# 添加 秘钥到 ssh-agent

然后才能使用
```



