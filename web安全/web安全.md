# web 安全攻击

## 1. XSS

> cross site scripting 跨站脚本攻击

原理

> <d iv >#{content}< /d iv> 
>
> 将 content 改成 一些< script src="www.baidu.com/xx.js"> </ script> 

1. 获取页面数据
2. 获取Cookie
3. 劫持前端逻辑
4. 发送请求
5. 偷取网站的资料
6. 获取用户密码登录状态
7. 欺骗用户



xss攻击的分类

1. 放射型：url参数直接注入
2. 存储型：存储到DB后读取时注入



# web安全深度解析

1. 作为web开发人员要牢记，JavaScript验证只是为了防止用户误操作，而服务器端的验证才是能够防止恶意攻击。（可以拦截 更改post数据）

