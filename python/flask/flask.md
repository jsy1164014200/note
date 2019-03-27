

# flask 笔记

# 初识flask

> python3 中使用 -m venv  生成虚拟环境 ，requirement.txt文件
>
> (venv) $ pip freeze >requirements.txt
>
> (venv) $ pip install -r requirements.txt

创建虚拟环境

1. virtualenv --no-site-packages venv
2. source ./venv/bin/activate.fish   **如果用的是bash就不用 .fish**
3. deactivate



## 运行

```bash
$ export FLASK_ENV=development
激活调试器。
激活自动重载。
打开 Flask 应用的调试模式。
$ export FLASK_APP=hello.py
$ python -m flask run
或者
$ flask run
```



## 路由的转换器类型

转换器类型：

| `string` | （缺省值） 接受任何不包含斜杠的文本 |
| -------- | ----------------------------------- |
| `int`    | 接受正整数                          |
| `float`  | 接受正浮点数                        |
| `path`   | 类似 `string` ，但可以包含斜杠      |
| `uuid`   | 接受 UUID 字符串                    |

URL 是中规中举的，尾部有一个斜杠，看起来就如同一个文件夹。 访问一个没有斜杠结尾的 URL 时 Flask 会自动进行重定向，帮你在尾部加上一个斜杠。



```
from flask import Flask,url_for 
app = Flask(__name__)

@app.route('/')
def index():
    return 'index'

@app.route('/login')
def login():
    return 'login'

@app.route('/user/<username>')
def profile(username):
    return '{}\'s profile'.format(username)
with app.test_request_context():
    print(url_for('index'))
    print(url_for('login'))
    print(url_for('login', next='/'))
    print(url_for('profile', username='John Doe'))

/
/login
/login?next=/
/user/John%20Doe
```



## http方法

```
from flask import request

@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        return do_the_login()
    else:
        return show_the_login_form()
```



## 测试覆盖

用 coverage run -m pytest

## 生成随机秘钥

```
python -c 'import os; print(os.urandom(16))'
```

# 详细的应用

```python
from flask_restful import Resource, reqparse, abort, Api
from flask import request


TODOS = {
        'todo1': {'task': 'build as API'},
        'todo2': {'task': '????'},
        'todo3': {'task': 'profit'}
        }


def abort_if_todo_doesnt_exist(todo_id):
    if todo_id not in TODOS:
        abort(404,message='Todo {} dosent exist'.format(todo_id))
# 下面这一行配置可以让 所有错误都以json形式返回
#app.config['BUNDLE_ERRORS'] = True

parser = reqparse.RequestParser()
# 如果不在参数里面，那么它将变成null(如果指定了required那么这也将返回help信息), 如果是类型错误，它将返回json数据，显示help里面的错误
parser.add_argument('task', type=int, required=True, help='Rate cannot be converted{error_msg}')
#parser.add_argument('old',dest='new')
# old将被绑定到 args的new下面'
# 默认的 是 flask.Request.values flask.Request.json 里面的值
# Look only in the POST body
#parser.add_argument('name', type=int, location='form')

# Look only in the querystring
#parser.add_argument('PageSize', type=int, location='args')

# From the request headers
#parser.add_argument('User-Agent', location='headers')

# From http cookies
#parser.add_argument('session_id', location='cookies')

# From file uploads
#parser.add_argument('picture', type=werkzeug.datastructures.FileStorage, location='files')






class Todo(Resource):
    @marshal_with(resource_fields, envelope='resource')
    def get(self, **kwargs):
        return db_get_todo()  # Some function that queries the db



# marshal_with 是一个 方便的装饰器 ,相当于
# 但是有与 marshal_with 只能返回一200 状态码，所以 如果要更改，使用下面的
#class Todo(Resource):
#    def get(self, **kwargs):
#        return marshal(db_get_todo(), resource_fields), 200


class UrgentItem(fields.Raw):
    def format(self, value):
        return "Urgent" if value & 0x01 else "Normal"
class UnreadItem(fields.Raw):
    def format(self, value):
        return "Unread" if value & 0x02 else "Read"

# 继承 fields.Raw 可以自定义 自己的属性
resource_fields = {
        # 实现属性重命名
        'name': fields.String(attribute='private_name'),
        'address': fields.String,
        'date_updated': fields.DateTime(dt_format='rfc822')
        }

class Todo(Resource):
    def get(self,todo_id):
        abort_if_todo_doesnt_exist(todo_id)
        return TODOS[todo_id]

    def delete(self,todo_id):
        abort_if_todo_doesnt_exist(todo_id)
        del TODOS[todo_id]
        return '',204

    def put(self,todo_id):
        args = parser.parse_args()
        task = {'task': args['task']}
        TODOS[todo_id] = task
        return task, 201


class TodoList(Resource):
    def get(self):
        return TODOS

    def post(self):
        args = parser.parse_args()
        todo_id = int(max(TODOS.keys()).lstrip('todo')) + 1
        todo_id = 'todo%i' % todo_id
        TODOS[todo_id] = {'task': args['task']}
        return TODOS[todo_id], 201


#  自定义资源装饰器，用来 做一些验证等
def authenticate(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        if not getattr(func, 'authenticated', True):
            return func(*args, **kwargs)

        acct = basic_authentication()  # custom account lookup function

        if acct:
            return func(*args, **kwargs)

        flask_restful.abort(401)
    return wrapper

class Resource(flask_restful.Resource):
    method_decorators = [authenticate]   # applies to all inherited resources



def cache(f):
    @wraps(f)
    def cacher(*args, **kwargs):
        # caching stuff
        return
    return cacher

class MyResource(restful.Resource):
    method_decorators = {'get': [cache]}

    def get(self, *args, **kwargs):
        return something_interesting(*args, **kwargs)

    def post(self, *args, **kwargs):
        return create_something(*args, **kwargs)


```























## cgi, wsgi ,usgi,uSGI

> 它是外部应用程序与Web服务器之间的接口标准

要了解目前比较常见的服务端结构：

假设我们使用 python 的 Django 框架写了一个网站，现在要将它挂在网上运行，我们一般需要：

1. nginx 做为代理服务器：负责静态资源发送（js、css、图片等）、动态请求转发以及结果的回复；
2. uWSGI 做为后端服务器：负责接收 nginx 请求转发并处理后发给 Django 应用以及接收 Django 应用返回信息转发给 nginx；
3. Django 应用收到请求后处理数据并渲染相应的返回页面给 uWSGI 服务器![img](https://upload-images.jianshu.io/upload_images/3027050-9cc25c602d43a977.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1000/format/webp)

**WSGI（Python Web Server GateWay Interface）:**它是用在 python web 框架编写的应用程序与后端服务器之间的规范（本例就是 Django 和 uWSGI 之间），让你写的应用程序可以与后端服务器顺利通信。在 WSGI 出现之前你不得不专门为某个后端服务器而写特定的 API，并且无法更换后端服务器，而 WSGI 就是一种统一规范， 所有使用 WSGI 的服务器都可以运行使用 WSGI 规范的 web 框架，反之亦然。

**uWSGI**: 是一个Web服务器，它实现了WSGI协议、uwsgi、http等协议。用于接收前端服务器转发的动态请求并处理后发给 web 应用程序。

**uwsgi**: 是uWSGI服务器实现的独有的**协议**， 网上没有明确的说明这个协议是用在哪里的，我个人认为它是用于前端服务器与 uwsgi 的通信规范，相当于 FastCGI的作用。当然这只是个人见解，我在[知乎](https://link.jianshu.com?t=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F46945479)进行了相关提问，欢迎共同讨论



