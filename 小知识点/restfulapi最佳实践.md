# restful Api 最佳实践

做出一个好的API设计很难。API表达的是你的数据和你的数据使用者之间的契约。打破这个契约将会招致很多愤怒的邮件，和一大堆伤心的用户-因为他们手机上的App不工作了。而文档化只能达到一半的效果，并且也很难找到一个愿意写文档的程序员。

## definitions

- **资源：**一个对象的单独实例，如一只动物
- **集合：**一群同种对象，如动物
- **HTTP：**跨网络的通信协议
- **客户端：**可以创建HTTP请求的客户端应用程序
- **第三方开发者：**这个开发者不属于你的项目但是有想使用你的数据
- **服务器：**一个HTTP服务器或者应用程序，客户端可以跨网络访问它
- **端点：**这个API在服务器上的URL用于表达一个资源或者一个集合
- **幂等：**无边际效应，多次操作得到相同的结果
- **URL段：**在URL里面已斜杠分隔的内容

## Actions

- GET (选择)：从服务器上获取一个具体的资源或者一个资源列表。
- POST （创建）： 在服务器上创建一个新的资源。
- PUT （更新）：以整体的方式更新服务器上的一个资源。
- PATCH （更新）：只更新服务器上一个资源的一个属性。
- DELETE （删除）：删除服务器上的一个资源。

还有两个不常用的HTTP动词：

- HEAD ： 获取一个资源的元数据，如数据的哈希值或最后的更新时间。
- OPTIONS：获取客户端能对资源做什么操作的信息。

一个好的RESTful API只允许第三方调用者使用这四个半HTTP动词进行数据交互，并且在URL段里面不出现任何其他的动词。

**一般来说，GET请求可以被浏览器缓存（通常也是这样的）。例如，缓存请求头用于第二次用户的POST请求。HEAD请求是基于一个无响应体的GET请求，并且也可以被缓存的。**

## Versioning

一个好的RESTful API会在URL中包含版本信息。另一种比较常见的方案是在请求头里面保持版本信息。但是跟很多不同的第三方开发者一起工作后，我可以很明确的告诉你，在请求头里面包含版本信息远没有放在URL里面来的容易。

## Analytics

当然第三方开发者的通知流程可以以某种条件被自动触发，例如每当一个过时的特性上发生10000次请求时就发邮件通知开发者。

## Root url

- https://example.org/api/v1/*
- https://api.example.com/v1/*

同样也请注意HTTPS前缀，一个好的RESTful API总是基于**HTTPS来发布的**。

## Examples

如果你正在构建一个虚构的API来展现几个不同的动物园，每一个动物园又包含很多动物，员工和每个动物的物种，你可能会有如下的端点信息：

- https://api.example.com/v1/**zoos**
- https://api.example.com/v1/**animals**
- https://api.example.com/v1/**animal_types**
- https://api.example.com/v1/**employees**

针对每一个具体的api

zoos

- GET /zoos:        		List all Zoos (ID and Name, not too much detail)
- POST /zoos:                   Create a new Zoo



- GET /zoos/ZID:              Retrieve an entire Zoo object
- PUT /zoos/ZID:              Update a Zoo (entire object)
- DELETE /zoos/ZID:        Delete a Zoo



- GET /zoos/ZID/animals:  Retrieve a listing of Animals (ID and Name).

animals

- GET /animals:                  List all Animals (ID and Name).
- POST /animals:               Create a new Animal



- GET /animals/AID:          Retrieve an Animal object
- PUT /animals/AID:          Update an Animal (entire object)

animals_types

- GET /animal_types:        Retrieve a listing (ID and Name) of all Animal Types
- GET /animal_types/ATID:  Retrieve an entire Animal Type object

employees

- GET /employees:            Retrieve an entire list of Employees
- POST /employees: Create a new Employee

- GET /employees/EID:     Retreive a specific Employee



- GET /zoos/ZID/employees: Retrieve a listing of Employees (ID and Name) who work at this Zoo
- POST /zoos/ZID/employees: Hire an Employee at a specific Zoo
- DELETE /zoos/ZID/employees/EID: Fire an Employee from a specific Zoo

让文档里所有的东西都有一个关键字是一个好主意。

为了简洁起见，我已经省略了所有API共有的URL前缀。作为沟通方式这没什么问题，但是如果你真要写到API文档中，那就必须包含完整的路径（如，GET http://api.example.com/v1/animal_type/ATID）。

**请注意如何展示数据之间的关系，特别是雇员与动物园之间的多对多关系。通过添加一个额外的URL段就可以实现更多的交互能力。**

## Filtering

尽可能减少那些会影响到第三方开发者的无谓限制

这点很重要，但你可以让客户端自己对结果做一些具体的过滤或限制。这么做最重要的一个原因是可以最小化网络传输，并让客户端尽可能快的得到查询结果。其次是客户端可能比较懒，如果这时服务器能对结果做一些过滤或分页，对大家都是好事。另外一个不那么重要的原因是（从客户端角度来说），对服务器来说响应请求的负载越少越好

这点很重要，但你可以让客户端自己对结果做一些具体的过滤或限制。这么做最重要的一个原因是可以最小化网络传输，并让客户端尽可能快的得到查询结果。其次是客户端可能比较懒，如果这时服务器能对结果做一些过滤或分页，对大家都是好事。另外一个不那么重要的原因是（从客户端角度来说），对服务器来说响应请求的负载越少越好。

- ?limit=10: 减少返回给客户端的结果数量（用于分页）
- ?offset=10: 发送一堆信息给客户端（用于分页）
- ?animal_type_id=1: 使用条件匹配来过滤记录
- ?sortby=name&order=asc:  对结果按特定属性进行排序
- ?state=open: 做查询

## Status Codes

- 200

  OK – [GET]

  - 客户端向服务器请求数据，服务器成功找到它们

- 201

  CREATED – [POST/PUT/PATCH]

  - 客户端向服务器提供数据，服务器根据要求创建了一个资源

- 204

  NO CONTENT – [DELETE]

  - 客户端要求服务器删除一个资源，服务器删除成功

- 304

  not modified

  - HTTP缓存有效

- 400

  INVALID REQUEST – [POST/PUT/PATCH]

  - 客户端向服务器提供了不正确的数据，服务器什么也没做

- 401

  unauthorized

  - 非授权

- 403

  forbidden

  - 鉴权成功，但是该用户没有权限

- 404

  NOT FOUND – [*]

  - 客户端引用了一个不存在的资源或集合，服务器什么也没做

- 405

  method not allowed 

  - 该http 方法不被允许

- 500

  INTERNAL SERVER ERROR – [*]

  - 服务器发生内部错误，客户端无法得知结果，即便请求已经处理成功

## Return Data

- GET /collection: 返回一系列资源对象
- POST /collection: 返回新创建的资源对象



- GET /collection/resource: 返回单独的资源对象
- PUT /collection/resource: 返回完整的资源对象
- DELETE /collection/resource: 返回一个空文档

## Authentication

restful API是无状态的也就是说用户请求的鉴权和cookie以及session无关，每一次请求都应该包含鉴权证明。

现在常用 jwt

OAuth2.0提供了一个非常好的方法去做这件事。在每一个请求里，你可以明确知道哪个客户端创建了请求，哪个用户提交了请求，并且提供了一种标准的访问过期机制或允许用户从客户端注销，所有这些都不需要第三方的客户端知道用户的登陆认证信息。

还有OAuth1.0和xAuth同样适用这样的场景。无论你选择哪个方法，请确保它为多种不同语言/平台上的库提供了一些通用的并且设计良好文档，因为你的用户可能会使用这些语言和平台来编写客户端。

## Content-Type

JSON

## Documentation

老实说，即使你不能百分之百的遵循指南中的条款，你的API也不是那么糟糕。但是，如果你不为API准备文档的话，没有人会知道怎么使用它，那它真的会成为一个糟糕的API。

- 让你的文档对那些未经认证的开发者也可用
- 不要使用文档自动化生成器，即便你用了，你也要保证自己审阅过并让它具有更好的版式。
- 不要截断示例中请求与响应的内容，要展示完整的东西。并在文档中使用高亮语法。
- 文档化每一个端点所预期的响应代码和可能的错误消息，和在什么情况下会产生这些的错误消息

如果你有富余的时间，那就创建一个控制台来让开发者可以立即体验一下API的功能。创建一个控制台并没有想象中那么难，并且开发者们（内部或者第三方）也会因此而拥戴你

## 缓存

HTTP提供了自带的缓存框架。你需要做的是在返回的时候加入一些返回头信息，在接受输入的时候加入输入验证。基本两种方法：

**ETag：**当生成请求的时候，在HTTP头里面加入ETag，其中包含请求的校验和和哈希值，这个值和在输入变化的时候也应该变化。如果输入的HTTP请求包含IF-NONE-MATCH头以及一个ETag值，那么API应该返回304 not modified状态码，而不是常规的输出结果。

**Last-Modified：**和etag一样，只是多了一个时间戳。返回头里的Last-Modified：包含了 [RFC 1123](http://www.ietf.org/rfc/rfc1123.txt) 时间戳，它和IF-MODIFIED-SINCE一致。HTTP规范里面有三种date格式，服务器应该都能处理。

## 使用 gzip

## 只在需要的时候使用“envelope”



[实际的案例](http://dev.enchant.com/api/v1)



[github v3采用的restfulapi](https://developer.github.com/v3/)



slate 制作文档

https://swagger.io/tools/swaggerhub/

在线调试工具







