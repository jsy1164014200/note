

# flask 测试

Flask 提供的测试渠道是使用 Werkzeug 的 [`Client`](http://werkzeug.pocoo.org/docs/test/#werkzeug.test.Client) 类， 并为你处理本地环境。你可以结合这个渠道使用你喜欢的测试工具。

本文使用 [pytest](https://pytest.org/) 包作为测试的基础框架。你可以像这样使用 `pip` 来安装它:

```bash
pip install pytest
```

首先我们在应用的根文件夹中添加一个测试文件夹。然后创建一个 Python 文件来储 存测试内容（ `test_flaskr.py`）。名称类似 `test_*.py` 的文件会被 pytest **自动发现**。



使用flask_testing  https://pythonhosted.org/Flask-Testing/







# 可拔插试图

对于 REST 式的 API 来说，为每种 HTTP 方法提供相对应的不同函数显得尤为有用。使用[`flask.views.MethodView`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.views.MethodView) 可以轻易做到这点。在这个类中，每个 HTTP 方法 都映射到一个同名函数（函数名称为小写字母）:

```
from flask.views import MethodView

class UserAPI(MethodView):

    def get(self):
        users = User.query.all()
        ...

    def post(self):
        user = User.from_form_data(request.form)
        ...

app.add_url_rule('/users/', view_func=UserAPI.as_view('users'))
```

使用这种方式，不必提供 [`methods`](https://dormousehole.readthedocs.io/en/latest/api.html#flask.views.View.methods) 属性，它会自动使用相应 的类方法。



[restful api实践](https://dormousehole.readthedocs.io/en/latest/views.html)





```
>>> from flask_restful import fields, marshal
>>> import json
>>>
>>> resource_fields = {'name': fields.String}
>>> resource_fields['address'] = {}
>>> resource_fields['address']['line 1'] = fields.String(attribute='addr1')
>>> resource_fields['address']['line 2'] = fields.String(attribute='addr2')
>>> resource_fields['address']['city'] = fields.String
>>> resource_fields['address']['state'] = fields.String
>>> resource_fields['address']['zip'] = fields.String
>>> data = {'name': 'bob', 'addr1': '123 fake street', 'addr2': '', 'city': 'New York', 'state': 'NY', 'zip': '10468'}
>>> json.dumps(marshal(data, resource_fields))
'{"name": "bob", "address": {"line 1": "123 fake street", "line 2": "", "state": "NY", "zip": "10468", "city": "New York"}}'
```



