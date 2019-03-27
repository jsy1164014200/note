# 1. 创建虚拟环境

```python
python -m venv <name>
激活虚拟环境
切换到已成功创建虚拟环境的目录，执行如下命令，激活虚拟环境
Windows:cmd
Scripts\activate.bat
退出
Scripts\deactivate

如果用powershell要 修改一下.lsp文件的执行权限
然后执行activate.lsp

pip list
pip freeze
```

# 2. Django的使用

1. Django创建项目：django-admin startproject demo

=====================

2. 使用Django创建一个应用：python manage.py startapp booktest
3. 运行一个项目：python manage.py runserver port(8080)

模型：就是数据库的操作

    生成迁移文件：python manage.py makemigrations

    1. 所有创建的应用要在setting的INSTALLED_APPS中注册，第三方包也是一样
    2. python manage.py migrate执行迁移，根据sql语句创建表
    **python manage.py shell**运行一个shell,在里面可以进行数据库操作

    管理后台

    1. python manage.py createsuperuser,按提示输入用户名，邮箱，密码
        python manage.py runserver
    2. 启动服务器，"127.0.0.1/admin"访问，输入用户名，密码进行登录
    3. 进入站点后，可以对groups，users进行管理
    4. 将数据库models注册到管理后台里面

url 

模板视图

# 3. Django模型详解

orm 对象-关系-映射：不需要面对相应的SQL语句

import pymysql
pymysql.install_as_MySQLdb()
1. 使用MySQL数据库 pip install -i https://pypi.tuna.tsinghua.edu.cn/simple pymysql
create databases test2 charset=utf8

python manage.py inspectdb > test_2/models.py 使用数据库生成模型类

    字段类型：
        1. AutoField: 自动增长的主键，一般不用定义，会自动添加到表中
        2. BooleanField: true/false 字段，对应的组件是 checkboxInput
        3. NullBooleanField: null/true/false
        4. CharField(max_length= ) 字符串，对应的组件是textInput
        5. integerField:整数
        6. DecimalField(max_digits=None,decimal_places=None) 十进制浮点数，decimalField.max_digits:位数总数，decimalField.decimal_places:小数点后面的位数
        7. FloatField 浮点数
        8. DateField[auto_now=False,auto_now_add=False]
        8. TimeField
        9. DateTimeField
        10. FileField #一般不用
        11. ImageField #一般不用

    元选项
        1. 在model类中定义Meta类，设置原信息
        2. 原信息db_table 定义数据表的名称，默认是 <app_name>_<model_name>
        3. ordering:对象的默认排序字段，获取对象的列表时使用，字符前面加-表示倒序
            排序会增加数据库的开销

    模型的成员
        1. object:是manager类型的对象，用于与数据库进行交互
        2. 当定义的模型类没有指定管理器，Django会为模型提供一个名为object的管理器
        3. 支持明确指定类型的管理器（当为模型类指定管理器后，Django不再生成object）

    管理器Manager
        1. 它是进行数据库查询的接口，Django应用的每个模型都拥有至少一个管理器
        2. 用两种情况要自定义：
            a. 添加管理方法
            b. 修改管理器返回的原始查询集，重写get_queryset方法
            
    模型查询
        1. all() 会调用get_queryset()
        2. filter(key=value,key=value) 
        3. exclude()
        4. order_by()
        5. values() 一个对象构成一个字典，
        6. get() 返回多个对象
        7. count() 返回查询的总条数
        8. first() 返回第一个对象
        9. last()
        10. exists() 判断是否有数据

    字段查询
        1. 实现where子名，作为filter() ，exclude(),get()的参数
        2. 语法：属性名称__比较运算符=值
        3. 对于外键，使用属性名_id 表示外键的值
        4. 转义：like语句中使用了%与，匹配数据中的%与，在过滤器中直接写
        5. example： filter(title__contains="%") => where title like "%%%" 表示查找标题中包含%的

    比较运算符
        1. exact：表示判断 等于 大小写敏感
        2. contains:是否包含 
        3. startswith,endswith 
        4. isnull,isnotnull
        5. 在filter exclude前面加 i 表示不区分大小写
        6. pk_in=[1,2,3]
        7. year,month,day,week_day,hour,minute,second
            filter(bpub_date__year=1980)
        8. gt gte lt lte 打小于
        9. 跨关系查询：处理join查询
            语法：模型类名<属性名><比较>
            没有__<比较>部分表示等于，同于inner join
            可反向使用，及在关联的两个模型中都可以使用

    聚合函数
        1. aggregate() 聚合
        2. count Avg Count Max Min Sum

    F对象: 
        1. 可以使用模型的字段A 与字段B进行比较，如果A 写在了等号的左边，则B出现在等号的右边，需要通过F对象构造 list.filter(bread__gte=F('bcomment'))
        2. Django支持对F()对象使用算数运算符 lsit.filter(bread__gte=F('bcommet')*2)
        3. F()对象中还可以使用 "模型类__列名"
        4. date/time

    Q对象
        1. Q(pk__gte=2) | Q()
        

# 4. 视图详解

    show(*args,**kw) 关键字参数，位置参数
```python
def show(*args,**kwargs):
    print(args)
    print(kwargs)

show(12,34,55,2443) 
# (12, 34, 55, 2443)
#{}
def show(*args,**kwargs):
    print(args)
    print(kwargs)
# 位置参数要写在 关键字参数前面
show(123423,p1=12,p2=34,p3=55)
# (123423,)
#{'p1': 12, 'p2': 34, 'p3': 55}
```

404.HTML只需要放在 setting里面的 模板目录下即可显示，不需要配置URL
500.html

request解析
1. path ：路径
2. method
3. encoding ： 一个字符串，数据的编码格式
4. GET 一个类似字典的对象，包含get请求的所有参数
5. POST 通get
6. FILES 一个类似字典的对象，包含所有上传的文件
7. COOKIES： 一个标准的Python字典，包含所有的cookie，键值对
8. session： 类似字典的对象

方法
1. is_ajax() 是不是XMLHttpRequest发起的请求

QueryDict对象 GET属性 POST 属性
1. 定义在Django.http.QueryDict
2. request对象的属性GET，POST都是QueryDict对象
3. 与Python对象不同，处理一个键带有多个值的情况
4. get() 根据键获取值  只能获取一个值，如果有多个值，获取最后一个
5. getlist() 获取多个值

response解析

HttpResponse()

HttpResponseRedirect()

JsonResponse() : _init__(data)

1. content 表示返回的内容
2. charset 表示编码集
3. status_code 响应的状态码
4. content-type 输出的MIME类型

方法
1. init 使用页面内容实例化httpResponse对象
2. write(content) 以文件的方式写
3. flush() 以文件的方式输出缓存区
4. set_cookie(key,value="",max_age=None,expires=None)
        key value 是字符串类型
        max_age 是一个整数，指定多少秒数后过期
        expires是一个datetime或者timedelta对象，在这个指定的日期过期
        以上两个二选一
        如果不指定，两个星期后过期
5. delete_cookie(key) 删除cookie 
                            response.delete_cookie('t1')
```python
def handle_index(request):
    response=HttpResponse()
    response.set_cookie('t1','abc')
    return response
```

response 使用模板
```python
from django.http import *
from django.template import loader
context={}
html = loader.get_template('booktest/index.html').render(context)
response=HttpResponse(html)
return response
```

简写
```python
from django.shortcuts import render
context={'request':request}
return render(request,'viewstudy/index.html',context)
```


session 机制

启用session ,在INSTALLED_APPS列表中添加django.contrib.sessions
在MIDDLEWARE_CLASSES列表中添加django.contrib.sessions.middleware.SessionMiddleware

使用session 
- 启用会话后，每个HttpRequest对象将具有一个session属性，它是类字典对象
- get(key,default=None) 根据键获取会话的值
- clear() 清除所有会话
- flush() 删除当前的会话数据并删除cookie
- del request.session['member_id'] 删除会话
- demo 

```python
def session(request):
    
    name = request.session.get('name','未登录')
    context={'name':name}
    return render(request,'viewstudy/session.html',context)


def session1(request):
    return render(request,'viewstudy/session1.html',{})

def session2(request):
    del request.session['name'] # 注意 此处设置session是给request设置，但是cookie是给response设置
    return HttpResponseRedirect('/viewstudy/session/')

def session_handle(request):
    name = request.POST['name']
    request.session['name']=name
    response = HttpResponseRedirect('/viewstudy/session/')
    return response
```

session原理
- set_expiry(value) 设置会话过期时间
- 默认两个星期
- 整数，则为妙
- imedelta对象，则为当前时间加上这个日期
- 0 ，关闭就过期
- None，永不过期

base64编码
配置session的存储地方：session_engine，保存在MySQL数据库的django_engine表中


# 5. 模板

DTL :django.template包中
setting中有关于模板的值
DIRS：总目录 APP_DIRS每个应用的模板目录

dirs = [os.path.join(BASE_DIR,"template")]

模板处理：
1. 变量{{  }}
2. 标签{%  %}
3. 过滤器
4. 注释{#  #}

- . 的查询顺序 1. 字典查询 2. 属性 方法查询 3. 数字索引查询
- 模板中调用方法时不能传递参数 用.   book.handlefunction
- 如果变量不存在，或者出错，将插入一个空字符串

标签
- {% for in %}
{%empty%}
{%endfor%}

- {{ forloop.conter }}可以得到循环的第几次
- {% if %}
{%elif%}
{%endif%}
- {%comment%}
{%endcomment%}
- {%include "foo/bar.html" %} 加载模板并已标签类的参数解析
- URL反向解析 {% url 'name' p1 p2%}
- csrf_token:这个标签用于跨站请求伪造保护{% csrf_token%}
- and ,or,布尔标签
- block extends 继承
- autoescape HTML转义

过滤器 {%变量|过滤器%} {{name|lower}}编程小写

- divisibleby 是否整除
- if list|length >1
- 过滤器能够串联 name|lower|upper
- 过滤器能过传递参数 list|join:","
- default value|default:"There is none"
- date 根据格式 对一个date变量格式化 value|date:'Y-m-d'
- escape

注释 
- {# 单行注释 #}
-  {%comment%}
    {%endcomment%}


反向解析：根据URL生成URL 
- include(('viewstudy.urls','viewstudy'),namespace="viewstudy")
- path('getTest1/',getTest1,name='getTest1'),
- <a href="{% url 'viewstudy:detail' '233' %}">sfsdfsd</a>

模板继承
- block标签：在父模板中留坑，字幕版填充
{% block block_name %}
{%endblock block_name %}
- extends 继承标签
{%extend 'viewstudy/base.html'%}


HTML转义

给的 字符串 会直接转义，原样输出escape  
    safe不转义
    或者{%autoescape off%}
    {%endautoescape%}

会自动转义的字符
- <  &lt ;
- -> &gt ;
- ' &#39 ;
- " &quot ;
- & &amp ;

csrf跨站请求伪造
{% csrf_token %} 放在from表单中 input type=hidden name= value

验证码
- viewsUtil.py 中定义函数 verifycode

# 6. 高级应用

静态文件处理
- 在setting中设置静态内容
- STATIC_URL = '/static/'
- STATICFILES_DIRS=[os.path.join(BASE_DIR,'static')]

中间件

- 激活添加到 MIDDLEWARE_CLASSES元组中
- 每个中间件都是一个Python类 ，可以定义方法
    1. __ init__： 不要参数，服务器响应第一个请求时候调用一次，用于确定时候启用当前中间件，
    2. process_request(request): 执行视图之前被调用，返回None或HttpResponse对象
    3. process_view(request,view_func,view_args,view_kwargs): 调用视图之前被调用，在每个请求上调用，返回None或者HttpResponse对象
    4. process_template_response(request,response) 在视图刚好执行完毕后调用，在每个请求上调用，返回实现了render方法的响应对象
    5. process_response(request,response) 所有的响应返回浏览器之前被调用，在每个请求上调用，返回HttpResponse对象
    6. process_exception(request,response,exception) 在所有的视图抛出异常后调用

- 使用中间件，可以干扰整个处理过程，每次请求中都会执行中间件这个方法

- demo 在setting.py同级目录下创建myexception.py文件，定义类MyException

上传图片

<input type="file" name="" />
注意：FILES只有在请求的方法为POST且提交的form表单带有 enctype="multipart/form-data"的情况下才会包含数据，admin若使用模型上传文件：将属性定义成models.ImageField类型 pic=models.ImageField(upload_to="cars/")

图片储存路径
- 创建media文件夹
- 图片上传后，存在/static/media/cars/图片文件
- 打开setting.py文件，增加media_root  MEDIA_ROOT=os.path.join(BASE_DIR,"static/media")


分页 Paginator对象
- Paginator(列表，int) 返回分页对象，参数为列表数据，每页数据的条数
- count:对象总数
- num_pages 页面总数
- page_range: 页码列表，从1开始
- page(num) 下标从1开始，返回一个page对象

page对象的属性
- object_list
- number
- paginator 得到对应的paginator对象
