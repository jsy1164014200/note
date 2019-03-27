# apidoc的使用

> [官网](http://apidocjs.com/)

```python
"""
@api {get} /user/:id Request User information
@apiVersion 1.0.0
@apiName GetUser
@apiGroup User
@apiPermission admin

@apiDescription xxx

@apiHeader {String} Authorization User authorization
@apiParam {Number,Boolean,String,Object,String[]} id Users unique ID.
	{string="small"} a string that can only contain the word "small" (a constant).
	{string="small","huge"} a string that can contain the words "small" or "huge".
	{number=1,2,3,99} a number with allowed values of 1, 2, 3 and 99.

@apiSuccess {String} firstname Firstname of the User.
@apiSuccess {String} lastname  Lastname of the User.
@apiSuccess {String}   id            The Users-ID.
@apiSuccess {Date}     registered    Registration Date.
@apiSuccess {Date}     name          Fullname of the User.
@apiSuccess {String[]} nicknames     List of Users nicknames (Array of Strings).
@apiSuccess {Object}   profile       Profile data (example for an Object)
@apiSuccess {Number}   profile.age   Users age.
@apiSuccess {String}   profile.image Avatar-Image.
@apiSuccess {Object[]} options       List of Users options (Array of Objects).
@apiSuccess {String}   options.name  Option Name.
@apiSuccess {String}   options.value Option Value.
@apiSuccessExample {json} success-response:
	HTTP/1.1 200 OK
    {
    "firstname": "Jons"
    }

@apiError  xxxx
@apiErrorExample Response (example):
	HTTP/1.1 401 Not Authenticated
"""
```



apidoc 

1. 可以使用继承 @apiDefine UserNotFoundError   @apiUse UserNotFoundError
2. 使用version  版本化 apiVersion
3. 

apidoc -i src/ apidoc/