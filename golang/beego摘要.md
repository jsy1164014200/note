# beego摘要





## 1. 配置信息

#### 自带的配置

通过`beego.BConfig.AppName="beepkg"`这样来修改 自带的配置

##### App 配置

- AppName

  应用名称，默认是 beego。通过 `bee new` 创建的是创建的项目名。

  `beego.BConfig.AppName = "beego"`

- RunMode

  应用的运行模式，可选值为 `prod`, `dev` 或者 `test`. 默认是 `dev`, 为开发模式，在开发模式下出错会提示友好的出错页面，如前面错误描述中所述。

  `beego.BConfig.RunMode = "dev"`

- RouterCaseSensitive

  是否路由忽略大小写匹配，默认是 true，区分大小写

  `beego.BConfig.RouterCaseSensitive = true`

- ServerName

  beego 服务器默认在请求的时候输出 server 为 beego。

  `beego.BConfig.ServerName = "beego"`

- RecoverPanic

  是否异常恢复，默认值为 true，即当应用出现异常的情况，通过 recover 恢复回来，而不会导致应用异常退出。

  `beego.BConfig.RecoverPanic = true`

- CopyRequestBody

  是否允许在 HTTP 请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件请求除外）。

  `beego.BConfig.CopyRequestBody = false`

- EnableGzip

  是否开启 gzip 支持，默认为 false 不支持 gzip，一旦开启了 gzip，那么在模板输出的内容会进行 gzip 或者 zlib 压缩，根据用户的 Accept-Encoding 来判断。

  `beego.BConfig.EnableGzip = false`

  Gzip允许用户自定义压缩级别、压缩长度阈值和针对请求类型压缩:

  1. 压缩级别, `gzipCompressLevel = 9`,取值为 1~9,如果不设置为 1(最快压缩)
  2. 压缩长度阈值, `gzipMinLength = 256`,当原始内容长度大于此阈值时才开启压缩,默认为 20B(ngnix默认长度)
  3. 请求类型, `includedMethods = get;post`,针对哪些请求类型进行压缩,默认只针对 GET 请求压缩

- MaxMemory

  文件上传默认内存缓存大小，默认值是 `1 << 26`(64M)。

  `beego.BConfig.MaxMemory = 1 << 26`

- EnableErrorsShow

  是否显示系统错误信息，默认为 true。

  `beego.BConfig.EnableErrorsShow = true`

- EnableErrorsRender

  是否将错误信息进行渲染，默认值为 true，即出错会提示友好的出错页面，对于 API 类型的应用可能需要将该选项设置为 false 以阻止在 `dev`模式下不必要的模板渲染信息返回。

##### Web配置

- AutoRender

  是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。

  `beego.BConfig.WebConfig.AutoRender = true`

- EnableDocs

  是否开启文档内置功能，默认是 false

  `beego.BConfig.WebConfig.EnableDocs = true`

- FlashName

  Flash 数据设置时 Cookie 的名称，默认是 BEEGO_FLASH

  `beego.BConfig.WebConfig.FlashName = "BEEGO_FLASH"`

- FlashSeperator

  Flash 数据的分隔符，默认是 BEEGOFLASH

  `beego.BConfig.WebConfig.FlashSeperator = "BEEGOFLASH"`

- DirectoryIndex

  是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。

  `beego.BConfig.WebConfig.DirectoryIndex = false`

- StaticDir

  静态文件目录设置，默认是static

  可配置单个或多个目录:

  1. 单个目录, `StaticDir = download`. 相当于 `beego.SetStaticPath("/download","download")`
  2. 多个目录, `StaticDir = download:down download2:down2`. 相当于 `beego.SetStaticPath("/download","down")` 和 `beego.SetStaticPath("/download2","down2")`

  `beego.BConfig.WebConfig.StaticDir`

- StaticExtensionsToGzip

  允许哪些后缀名的静态文件进行 gzip 压缩，默认支持 .css 和 .js

  `beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}`

  等价 config 文件中

  ```
  StaticExtensionsToGzip = .css, .js
  ```

- TemplateLeft

  模板左标签，默认值是`{{`。

  `beego.BConfig.WebConfig.TemplateLeft="{{"`

- TemplateRight

  模板右标签，默认值是`}}`。

  `beego.BConfig.WebConfig.TemplateRight="}}"`

- ViewsPath

  模板路径，默认值是 views。

  `beego.BConfig.WebConfig.ViewsPath="views"`

- EnableXSRF

  是否开启 XSRF，默认为 false，不开启。

  `beego.BConfig.WebConfig.EnableXSRF = false`

- XSRFKEY

  XSRF 的 key 信息，默认值是 beegoxsrf。 EnableXSRF＝true 才有效

  `beego.BConfig.WebConfig.XSRFKEY = "beegoxsrf"`

- XSRFExpire

  XSRF 过期时间，默认值是 0，不过期。

  `beego.BConfig.WebConfig.XSRFExpire = 0`

##### 监听配置

- Graceful

  是否开启热升级，默认是 false，关闭热升级。

  `beego.BConfig.Listen.Graceful=false`

- ServerTimeOut

  设置 HTTP 的超时时间，默认是 0，不超时。

  `beego.BConfig.Listen.ServerTimeOut=0`

- ListenTCP4

  监听本地网络地址类型，默认是TCP6，可以通过设置为true设置为TCP4。

  `beego.BConfig.Listen.ListenTCP4 = true`

- EnableHTTP

  是否启用 HTTP 监听，默认是 true。

  `beego.BConfig.Listen.EnableHTTP = true`

- HTTPAddr

  应用监听地址，默认为空，监听所有的网卡 IP。

  `beego.BConfig.Listen.HTTPAddr = ""`

- HTTPPort

  应用监听端口，默认为 8080。

  `beego.BConfig.Listen.HTTPPort = 8080`

- EnableHTTPS

  是否启用 HTTPS，默认是 false 关闭。当需要启用时，先设置 EnableHTTPS = true，并设置 `HTTPSCertFile` 和 `HTTPSKeyFile`

  `beego.BConfig.Listen.EnableHTTPS = false`

- HTTPSAddr

  应用监听地址，默认为空，监听所有的网卡 IP。

  `beego.BConfig.Listen.HTTPSAddr = ""`

- HTTPSPort

  应用监听端口，默认为 10443

  `beego.BConfig.Listen.HTTPSPort = 10443`

- HTTPSCertFile

  开启 HTTPS 后，ssl 证书路径，默认为空。

  `beego.BConfig.Listen.HTTPSCertFile = "conf/ssl.crt"`

- HTTPSKeyFile

  开启 HTTPS 之后，SSL 证书 keyfile 的路径。

  `beego.BConfig.Listen.HTTPSKeyFile = "conf/ssl.key"`

- EnableAdmin

  是否开启进程内监控模块，默认 false 关闭。

  `beego.BConfig.Listen.EnableAdmin = false`

- AdminAddr

  监控程序监听的地址，默认值是 localhost 。

  `beego.BConfig.Listen.AdminAddr = "localhost"`

- AdminPort

  监控程序监听的地址，默认值是 8088 。

  `beego.BConfig.Listen.AdminPort = 8088`

- EnableFcgi

  是否启用 fastcgi ， 默认是 false。

  `beego.BConfig.Listen.EnableFcgi = false`

- EnableStdIo

  通过fastcgi 标准I/O，启用 fastcgi 后才生效，默认 false。

  `beego.BConfig.Listen.EnableStdIo = false`

##### Session配置

- SessionOn

  session 是否开启，默认是 false。

  `beego.BConfig.WebConfig.Session.SessionOn = false`

- SessionProvider

  session 的引擎，默认是 memory，详细参见 `session 模块`。

  `beego.BConfig.WebConfig.Session.SessionProvider = ""`

- SessionName

  存在客户端的 cookie 名称，默认值是 beegosessionID。

  `beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"`

- SessionGCMaxLifetime

  session 过期时间，默认值是 3600 秒。

  `beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600`

- SessionProviderConfig

  配置信息，根据不同的引擎设置不同的配置信息，详细的配置请看下面的引擎设置，详细参见 [session 模块](https://beego.me/docs/mvc/controller/zh-CN/module/session.md)

- SessionCookieLifeTime

  session 默认存在客户端的 cookie 的时间，默认值是 3600 秒。

  `beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600`

- SessionAutoSetCookie

  是否开启SetCookie, 默认值 true 开启。

  `beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true`

- SessionDomain

  session cookie 存储域名, 默认空。

  `beego.BConfig.WebConfig.Session.SessionDomain = ""`

##### Log配置

```
log详细配置，请参见 `logs 模块`。
```

- AccessLogs

  是否输出日志到 Log，默认在 prod 模式下不会输出日志，默认为 false 不输出日志。此参数不支持配置文件配置。

  `beego.BConfig.Log.AccessLogs = false`

- FileLineNum

  是否在日志里面显示文件名和输出日志行号，默认 true。此参数不支持配置文件配置。

  `beego.BConfig.Log.FileLineNum = true`

- Outputs

  日志输出配置，参考 logs 模块，console file 等配置，此参数不支持配置文件配置。

  `beego.BConfig.Log.Outputs = map[string]string{"console": ""}`

  or

  `beego.BConfig.Log.Outputs["console"] = ""`

#### 自定义的配置

你也可以在配置文件中配置应用需要用的一些配置信息，例如下面所示的数据库信息：

```
mysqluser = "root"
mysqlpass = "rootpass"
mysqlurls = "127.0.0.1"
mysqldb   = "beego"
```

那么你就可以通过如下的方式获取设置的配置信息:

```
beego.AppConfig.String("mysqluser")
beego.AppConfig.String("mysqlpass")
beego.AppConfig.String("mysqlurls")
beego.AppConfig.String("mysqldb")
```

AppConfig 的方法如下：

- Set(key, val string) error
- String(key string) string
- Strings(key string) []string
- Int(key string) (int, error)
- Int64(key string) (int64, error)
- Bool(key string) (bool, error)
- Float(key string) (float64, error)
- DefaultString(key string, defaultVal string) string
- DefaultStrings(key string, defaultVal []string)
- DefaultInt(key string, defaultVal int) int
- DefaultInt64(key string, defaultVal int64) int64
- DefaultBool(key string, defaultVal bool) bool
- DefaultFloat(key string, defaultVal float64) float64
- DIY(key string) (interface{}, error)
- GetSection(section string) (map[string]string, error)
- SaveConfigFile(filename string) error





#### 不同级别的配置

```
appname = beepkg
httpaddr = "127.0.0.1"
httpport = 9090
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"

[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888
```

读取不同模式下配置参数的方法是“模式::配置参数名”，比如：beego.AppConfig.String(“dev::mysqluser”)。



#### include 方式

iNI 格式配置支持 `include` 方式，引用多个配置文件，例如下面的两个配置文件效果同上：

app.conf

```
appname = beepkg
httpaddr = "127.0.0.1"
httpport = 9090

include "app2.conf"
```

app2.conf

```
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"

[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888
```





#### 通过环境变量配置

配置文件解析支持从环境变量中获取配置项，配置项格式：`${环境变量}`。例如下面的配置中优先使用环境变量中配置的 runmode 和 httpport,如果有配置环境变量 ProRunMode 则优先使用该环境变量值。如果不存在或者为空，则使用 “dev” 作为 runmode。

app.conf

```
runmode  = "${ProRunMode||dev}"
httpport = "${ProPort||9090}"
```



## 2. 路由的使用

> 固定路由，正则路由，自动路由

- beego.Get(router, beego.FilterFunc)
- beego.Post(router, beego.FilterFunc)
- beego.Put(router, beego.FilterFunc)
- beego.Patch(router, beego.FilterFunc)
- beego.Head(router, beego.FilterFunc)
- beego.Options(router, beego.FilterFunc)
- beego.Delete(router, beego.FilterFunc)
- beego.Any(router, beego.FilterFunc)

基本 GET 路由

```
beego.Get("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world"))
})
```

基本 POST 路由

```
beego.Post("/alice",func(ctx *context.Context){
     ctx.Output.Body([]byte("bob"))
})
```

注册一个可以响应任何 HTTP 的路由

```
beego.Any("/foo",func(ctx *context.Context){
     ctx.Output.Body([]byte("bar"))
})
```

#### 注册rpc应用

```go
s := rpc.NewServer()
s.RegisterCodec(json.NewCodec(), "application/json")
s.RegisterService(new(HelloService), "")
beego.Handler("/rpc", s)
```

这个函数其实还有第三个参数就是是否是前缀匹配,默认是 false, 如果设置了 true, 那么就会在路由匹配的时候前缀匹配,即 `/rpc/user` 这样的也会匹配去运行

#### 固定路由

```
beego.Router("/", &controllers.MainController{})
beego.Router("/admin", &admin.UserController{})
beego.Router("/admin/index", &admin.ArticleController{})
beego.Router("/admin/addpkg", &admin.AddController{})
```

#### 正则路由

- beego.Router(“/api/?:id”, &controllers.RController{})

  默认匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”

- beego.Router(“/api/:id”, &controllers.RController{})

  默认匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”，但URL”/api/“匹配失败

- beego.Router(“/api/:id([0-9]+)“, &controllers.RController{})

  自定义正则匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”

- beego.Router(“/user/:username([\\w]+)“, &controllers.RController{})

  正则字符串匹配 //例如对于URL”/user/astaxie”可以匹配成功，此时变量”:username”值为”astaxie”

- beego.Router(“/download/*.*”, &controllers.RController{})

  *匹配方式 //例如对于URL”/download/file/api.xml”可以匹配成功，此时变量”:path”值为”file/api”， “:ext”值为”xml”

- beego.Router(“/download/ceshi/*“, &controllers.RController{})

  *全匹配方式 //例如对于URL”/download/ceshi/file/api.json”可以匹配成功，此时变量”:splat”值为”file/api.json”

- beego.Router(“/:id:int”, &controllers.RController{})

  int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)

- beego.Router(“/:hi:string”, &controllers.RController{})

  string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)

- beego.Router(“/cms_:id([0-9]+).html”, &controllers.CmsController{})

  带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123

可以在 Controller 中通过如下方式获取上面的变量：

```
this.Ctx.Input.Param(":id")
this.Ctx.Input.Param(":username")
this.Ctx.Input.Param(":splat")
this.Ctx.Input.Param(":path")
this.Ctx.Input.Param(":ext")
```

#### 自定义路由方法

```
beego.Router("/",&IndexController{},"*:Index")
```

使用第三个参数，第三个参数就是用来设置对应 method 到函数名，定义如下

- `*`表示任意的 method 都执行该函数
- 使用 httpmethod:funcname 格式来展示
- 多个不同的格式使用 `;` 分割
- 多个 method 对应同一个 funcname，method 之间通过 `,` 来分割

以下是一个 RESTful 的设计示例：

```
beego.Router("/api/list",&RestController{},"*:ListFood")
beego.Router("/api/create",&RestController{},"post:CreateFood")
beego.Router("/api/update",&RestController{},"put:UpdateFood")
beego.Router("/api/delete",&RestController{},"delete:DeleteFood")
```

以下是多个 HTTP Method 指向同一个函数的示例：

```
beego.Router("/api",&RestController{},"get,post:ApiFunc")
```

以下是不同的 method 对应不同的函数，通过 ; 进行分割的示例：

```
beego.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")
```

可用的 HTTP Method：

- *: 包含以下所有的函数
- get: GET 请求
- post: POST 请求
- put: PUT 请求
- delete: DELETE 请求
- patch: PATCH 请求
- options: OPTIONS 请求
- head: HEAD 请求

#### 注释路由

```
// CMS API
type CMSController struct {
    beego.Controller
}

func (c *CMSController) URLMapping() {
    c.Mapping("StaticBlock", c.StaticBlock)
    c.Mapping("AllBlock", c.AllBlock)
}


// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {

}
```

可以在 `router.go` 中通过如下方式注册路由：

```
beego.Include(&CMSController{})
```

beego 自动会进行源码分析，注意只会在 dev 模式下进行生成，生成的路由放在 “/routers/commentsRouter.go” 文件中。



#### namespace的使用

```go
//初始化 namespace
ns :=
beego.NewNamespace("/v1",
    beego.NSCond(func(ctx *context.Context) bool {
        if ctx.Input.Domain() == "api.beego.me" {
            return true
        }
        return false
    }),
    beego.NSBefore(auth),
    beego.NSGet("/notallowed", func(ctx *context.Context) {
        ctx.Output.Body([]byte("notAllowed"))
    }),
    beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
    beego.NSRouter("/changepassword", &UserController{}),
    beego.NSNamespace("/shop",
        beego.NSBefore(sentry),
        beego.NSGet("/:id", func(ctx *context.Context) {
            ctx.Output.Body([]byte("notAllowed"))
        }),
    ),
    beego.NSNamespace("/cms",
        beego.NSInclude(
            &controllers.MainController{},
            &controllers.CMSController{},
            &controllers.BlockController{},
        ),
    ),
)
//注册 namespace
beego.AddNamespace(ns)
```

- 

- NewNamespace(prefix string, funcs …interface{})

  初始化 namespace 对象,下面这些函数都是 namespace 对象的方法,但是强烈推荐使用 NS 开头的相应函数注册，因为这样更容易通过 gofmt 工具看的更清楚路由的级别关系

- NSCond(cond namespaceCond)

  支持满足条件的就执行该 namespace, 不满足就不执行

- NSBefore(filiterList …FilterFunc)

- NSAfter(filiterList …FilterFunc)

  上面分别对应 beforeRouter 和 FinishRouter 两个过滤器，可以同时注册多个过滤器

- NSInclude(cList …ControllerInterface)

- NSRouter(rootpath string, c ControllerInterface, mappingMethods …string)

- NSGet(rootpath string, f FilterFunc)

- NSPost(rootpath string, f FilterFunc)

- NSDelete(rootpath string, f FilterFunc)

- NSPut(rootpath string, f FilterFunc)

- NSHead(rootpath string, f FilterFunc)

- NSOptions(rootpath string, f FilterFunc)

- NSPatch(rootpath string, f FilterFunc)

- NSAny(rootpath string, f FilterFunc)

- NSHandler(rootpath string, h http.Handler)

- NSAutoRouter(c ControllerInterface)

- NSAutoPrefix(prefix string, c ControllerInterface)

  上面这些都是设置路由的函数,详细的使用和上面 beego 的对应函数是一样的

- NSNamespace(prefix string, params …innnerNamespace)

更多的例子代码：

```go
//APIS
ns :=
    beego.NewNamespace("/api",
        //此处正式版时改为验证加密请求
        beego.NSCond(func(ctx *context.Context) bool {
            if ua := ctx.Input.Request.UserAgent(); ua != "" {
                return true
            }
            return false
        }),
        beego.NSNamespace("/ios",
            //CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
            beego.NSNamespace("/create",
                // /api/ios/create/node/
                beego.NSRouter("/node", &apis.CreateNodeHandler{}),
                // /api/ios/create/topic/
                beego.NSRouter("/topic", &apis.CreateTopicHandler{}),
            ),
            beego.NSNamespace("/read",
                beego.NSRouter("/node", &apis.ReadNodeHandler{}),
                beego.NSRouter("/topic", &apis.ReadTopicHandler{}),
            ),
            beego.NSNamespace("/update",
                beego.NSRouter("/node", &apis.UpdateNodeHandler{}),
                beego.NSRouter("/topic", &apis.UpdateTopicHandler{}),
            ),
            beego.NSNamespace("/delete",
                beego.NSRouter("/node", &apis.DeleteNodeHandler{}),
                beego.NSRouter("/topic", &apis.DeleteTopicHandler{}),
            )
        ),
    )

beego.AddNamespace(ns)
```



## 3. Controller 设计

> 匿名组合 beego.Controller 就可以了

```
type xxxController struct {
    beego.Controller
}
```

- Init(ct *context.Context, childName string, app interface{})

  这个函数主要初始化了 Context、相应的 Controller 名称，模板名，初始化模板参数的容器 Data，app 即为当前执行的 Controller 的 reflecttype，这个 app 可以用来执行子类的方法。

- Prepare()

  这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。

- Get()

  如果用户请求的 HTTP Method 是 GET，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Get 请求。

- Post()

  如果用户请求的 HTTP Method 是 POST，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Post 请求。

- Delete()

  如果用户请求的 HTTP Method 是 DELETE，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Delete 请求。

- Put()

  如果用户请求的 HTTP Method 是 PUT，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Put 请求.

- Head()

  如果用户请求的 HTTP Method 是 HEAD，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Head 请求。

- Patch()

  如果用户请求的 HTTP Method 是 PATCH，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Patch 请求.

- Options()

  如果用户请求的HTTP Method是OPTIONS，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Options 请求。

- Finish()

  这个函数是在执行完相应的 HTTP Method 方法之后执行的，默认是空，用户可以在子 struct 中重写这个函数，执行例如数据库关闭，清理数据之类的工作。

#### 处理get post 等传递的参数

- GetString(key string) string
- GetStrings(key string) []string
- GetInt(key string) (int64, error)
- GetBool(key string) (bool, error)
- GetFloat(key string) (float64, error)

```go
func (this *MainController) Post() {
    jsoninfo := this.GetString("jsoninfo")
    if jsoninfo == "" {
        this.Ctx.WriteString("jsoninfo is empty")
        return
    }
}
```

如果你需要的数据可能是其他类型的，例如是 int 类型而不是 int64，那么你需要这样处理：

```
func (this *MainController) Post() {
    id := this.Input().Get("id")
    intid, err := strconv.Atoi(id)
}
```

更多其他的 request 的信息，用户可以通过 `this.Ctx.Request` 获取信息，关于该对象的属性和方法参考手册 [Request](http://gowalker.org/net/http#Request)。

- Method
- URL
- Proto
- ProtoMajor  int 
- ProtoMinor  int 
- Header
- Body  io.ReadCloser
- GetBody()    return  io.ReadCloser, error
- ContentLength
- TransferEncoding  []string
- Host  string
- Form   url.Values
- PostForm  url.Values
- RemoteAddr  string
- RequestURI  string
- TLS  *tls.ConnectionState



#### form直接解析到struct

定义 struct：

```
type user struct {
    Id    int         `form:"-"`
    Name  interface{} `form:"username"`
    Age   int         `form:"age"`
    Email string
}
```

表单：

```
<form id="user">
    名字：<input name="username" type="text" />
    年龄：<input name="age" type="text" />
    邮箱：<input name="Email" type="text" />
    <input type="submit" value="提交" />
</form>
```

Controller 里解析：

```go
func (this *MainController) Post() {
    u := user{}
    if err := this.ParseForm(&u); err != nil {
        //handle error
    }
}
```

- 如果要忽略一个字段，有两种办法，一是：字段名小写开头，二是：`form` 标签的值设置为 `-`



#### 处理json

```go
func (this *ObjectController) Post() {
    var ob models.Object
    var err error
    if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
        objectid := models.AddOne(ob)
        this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
    } else {
        this.Data["json"] = err.Error()
    }
    this.ServeJSON()
}
```



#### 输出数据 json  xml  等

一定要记得在配置文件加入copyrequestbody = true

```go
func (this *AddController) Get() {
    mystruct := { ... }
    this.Data["json"] = &mystruct
    this.ServeJSON()
}
```

调用 ServeJSON 之后，会设置 `content-type` 为 `application/json`，然后同时把数据进行 JSON 序列化输出。



XML 数据直接输出：

```
func (this *AddController) Get() {
    mystruct := { ... }
    this.Data["xml"]=&mystruct
    this.ServeXML()
}
```

调用 ServeXML 之后，会设置 `content-type` 为 `application/xml`，同时数据会进行 XML 序列化输出

#### this.Abort(401)



#### 提前终止执行

```go
type RController struct {
    beego.Controller
}

func (this *RController) Prepare() {
    this.Data["json"] = map[string]interface{}{"name": "astaxie"}
    this.ServeJSON()
    this.StopRun()
}
```

stopRun()



## 4. 处理日志

```go
beego.Emergency("this is emergency")
beego.Alert("this is alert")
beego.Critical("this is critical")
beego.Error("this is error")
beego.Warning("this is warning")
beego.Notice("this is notice")
beego.Informational("this is informational")
beego.Debug("this is debug")
```



我们的程序往往期望把信息输出到 log 中，现在设置输出到文件很方便，如下所示：

```
beego.SetLogger("file", `{"filename":"logs/test.log"}`)
```

更多详细的日志配置请查看 [日志配置](https://beego.me/docs/module/logs.md)

这个默认情况就会同时输出到两个地方，一个 console，一个 file，如果只想输出到文件，就需要调用删除操作：

```
beego.BeeLogger.DelLogger("console")
```







## 总结

1. this.GetString 系列 能得到 queryparam, x-www-form-urlencoded, form-data的数据 
2. id := this.Input().Get("id")
3. 不要使用 this.Abbort("xxx") 返回的是html页面
4. this.StopRun()
5. this.Ctx.Input.RequestBody  []byte 拿来做 json化
6. this.Ctx.Input 上面有一些请求的信息
7. this.Ctx.Request  原生的
8. this.Ctx.ResponseWriter  就是原生的





# mongo driver的使用

```go
package main

import (
   "github.com/gpmgo/gopm/modules/log"
   "gopkg.in/mgo.v2"
   "time"
)

type Person struct {
   Name  string
   Phone string
}

var mongoDialInfo = &mgo.DialInfo{
   Addrs: []string{
      "localhost",
   },
   Database: "admin",
   Username: "root",
   Password: "password",
   Timeout:  60 * time.Second,
}

//[[constraint]]
//name = "github.com/mongodb/mongo-go-driver"
func main() {
   session, err := mgo.DialWithInfo(mongoDialInfo)
   if err != nil {
      log.Fatal(err.Error())
   }
   test := session.DB("new").C("test")
   err = test.Insert(&Person{"Ale", "+55 53 8116 9639"},
      &Person{"Cla", "+55 53 8402 8510"})
   if err != nil {
      log.Fatal(err.Error())
   }

}
```



```go
type User struct 
{ 
Id_ bson.ObjectId `bson:"_id"` 
Name string `bson:"name"` 
Age int `bson:"age"` 
JonedAt time.Time `bson:"joned_at"` 
Interests []string `bson:"interests"` 
}

插入数据
err = c.Insert(&User{ Id_: bson.NewObjectId(), 
				Name: "Jimmy Kuu", 
				Age: 33, 
				JoinedAt: time.Now(),
				Interests: []string{"Develop", "Movie"}, }
				) 

查询
var users []User 
c.Find(nil).All(&users) 
fmt.Println(users)
Query.One()可以获得一个结果，注意如果没有数据或者数量超过一个，One()会报错。

id := "5204af979955496907000001" 
objectId := bson.ObjectIdHex(id) 
user := new(User) 
c.Find(bson.M{"_id": objectId}).One(&user) 
fmt.Println(user)
$and
c.Find(bson.M{"name": bson.M{"$ne": "Jimmy Kuu"}}).All(&users)
$or
c.Find(bson.M{"$or": []bson.M{bson.M{"name": "Jimmy Kuu"}, bson.M{"age": 31}}}).All(&users)

修改
c.Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")}, bson.M{"$set": bson.M{ "name": "Jimmy Gu", "age": 34, }})
其他操作跟原生的 差不多 $inc $push $pull等等

c.Remove(bson.M{"name": "Jimmy Kuu"})
```

所有collection操作

- [func (c *Collection) Bulk() *Bulk](https://godoc.org/gopkg.in/mgo.v2#Collection.Bulk)
- [func (c *Collection) Count() (n int, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.Count)
- [func (c *Collection) Create(info *CollectionInfo) error](https://godoc.org/gopkg.in/mgo.v2#Collection.Create)
- [func (c *Collection) DropCollection() error](https://godoc.org/gopkg.in/mgo.v2#Collection.DropCollection)
- [func (c *Collection) DropIndex(key ...string) error](https://godoc.org/gopkg.in/mgo.v2#Collection.DropIndex)
- [func (c *Collection) DropIndexName(name string) error](https://godoc.org/gopkg.in/mgo.v2#Collection.DropIndexName)
- [func (c *Collection) EnsureIndex(index Index) error](https://godoc.org/gopkg.in/mgo.v2#Collection.EnsureIndex)
- [func (c *Collection) EnsureIndexKey(key ...string) error](https://godoc.org/gopkg.in/mgo.v2#Collection.EnsureIndexKey)
- [func (c *Collection) Find(query interface{}) *Query](https://godoc.org/gopkg.in/mgo.v2#Collection.Find)
- [func (c *Collection) FindId(id interface{}) *Query](https://godoc.org/gopkg.in/mgo.v2#Collection.FindId)
- [func (c *Collection) Indexes() (indexes [\]Index, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.Indexes)
- [func (c *Collection) Insert(docs ...interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Collection.Insert)
- [func (c *Collection) NewIter(session *Session, firstBatch [\]bson.Raw, cursorId int64, err error) *Iter](https://godoc.org/gopkg.in/mgo.v2#Collection.NewIter)
- [func (c *Collection) Pipe(pipeline interface{}) *Pipe](https://godoc.org/gopkg.in/mgo.v2#Collection.Pipe)

```go
pipe := collection.Pipe([]bson.M{{"$match": bson.M{"name": "Otavio"}}})
iter := pipe.Iter()

func (p *Pipe) All(result interface{}) error
func (p *Pipe) AllowDiskUse() *Pipe
func (p *Pipe) Batch(n int) *Pipe
func (p *Pipe) Explain(result interface{}) error
func (p *Pipe) Iter() *Iter
func (p *Pipe) One(result interface{}) error

iter==========
func (iter *Iter) All(result interface{}) error
func (iter *Iter) Close() error
func (iter *Iter) Done() bool
func (iter *Iter) Err() error
func (iter *Iter) For(result interface{}, f func() error) (err error)
func (iter *Iter) Next(result interface{}) bool
func (iter *Iter) Timeout() bool
```



- [func (c *Collection) Remove(selector interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Collection.Remove)
- [func (c *Collection) RemoveAll(selector interface{}) (info *ChangeInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.RemoveAll)
- [func (c *Collection) RemoveId(id interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Collection.RemoveId)
- [func (c *Collection) Repair() *Iter](https://godoc.org/gopkg.in/mgo.v2#Collection.Repair)
- [func (c *Collection) Update(selector interface{}, update interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Collection.Update)
- [func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.UpdateAll)
- [func (c *Collection) UpdateId(id interface{}, update interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Collection.UpdateId)
- [func (c *Collection) Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.Upsert)
- [func (c *Collection) UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Collection.UpsertId)
- [func (c *Collection) With(s *Session) *Collection](https://godoc.org/gopkg.in/mgo.v2#Collection.With)

查询

- [type Query](https://godoc.org/gopkg.in/mgo.v2#Query)

- - [func (q *Query) All(result interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Query.All)
  - [func (q *Query) Apply(change Change, result interface{}) (info *ChangeInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Query.Apply)
  - [func (q *Query) Batch(n int) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Batch)
  - [func (q *Query) Comment(comment string) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Comment)
  - [func (q *Query) Count() (n int, err error)](https://godoc.org/gopkg.in/mgo.v2#Query.Count)
  - [func (q *Query) Distinct(key string, result interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Query.Distinct)
  - [func (q *Query) Explain(result interface{}) error](https://godoc.org/gopkg.in/mgo.v2#Query.Explain)
  - [func (q *Query) For(result interface{}, f func() error) error](https://godoc.org/gopkg.in/mgo.v2#Query.For)
  - [func (q *Query) Hint(indexKey ...string) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Hint)
  - [func (q *Query) Iter() *Iter](https://godoc.org/gopkg.in/mgo.v2#Query.Iter)
  - [func (q *Query) Limit(n int) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Limit)
  - [func (q *Query) LogReplay() *Query](https://godoc.org/gopkg.in/mgo.v2#Query.LogReplay)
  - [func (q *Query) MapReduce(job *MapReduce, result interface{}) (info *MapReduceInfo, err error)](https://godoc.org/gopkg.in/mgo.v2#Query.MapReduce)
  - [func (q *Query) One(result interface{}) (err error)](https://godoc.org/gopkg.in/mgo.v2#Query.One)
  - [func (q *Query) Prefetch(p float64) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Prefetch)
  - [func (q *Query) Select(selector interface{}) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Select)
  - [func (q *Query) SetMaxScan(n int) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.SetMaxScan)
  - [func (q *Query) SetMaxTime(d time.Duration) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.SetMaxTime)
  - [func (q *Query) Skip(n int) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Skip)
  - [func (q *Query) Snapshot() *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Snapshot)
  - [func (q *Query) Sort(fields ...string) *Query](https://godoc.org/gopkg.in/mgo.v2#Query.Sort)
  - [func (q *Query) Tail(timeout time.Duration) *Iter](https://godoc.org/gopkg.in/mgo.v2#Query.Tail)





# redisgo使用

```go
Go Type                 Conversion
[]byte                  Sent as is
string                  Sent as is
int, int64              strconv.FormatInt(v)
float64                 strconv.FormatFloat(v, 'g', -1, 64)
bool                    true -> "1", false -> "0"
nil                     ""
all other types         fmt.Fprint(w, v)
```

```go
Redis type              Go type
error                   redis.Error
integer                 int64
simple string           string
bulk string             []byte or nil if value not present.
array                   []interface{} or nil if value not present.
```

Bool，Int，Bytes，String，Strings和Values函数将回复转换为特定类型的值。为了方便地包含对连接Do和Receive方法的调用，这些函数采用了类型为error的第二个参数。如果错误是非nil，则辅助函数返回错误。如果错误为nil，则该函数将回复转换为指定的类型：

```go
exists，err：= redis.Bool（c.Do（“EXISTS”，“foo”））
if err！= nil { 
    //句柄错误从c.Do返回或输入转换错误。
}
```

Scan函数将数组回复的元素转换为Go类型：

```go
var value1 int 
var value2 string 
reply，err：= redis.Values（c.Do（“MGET”，“key1”，“key2”））
if err！= nil { 
    // handle error 
} 
 if _，err：= redis .Scan（回复，＆value1，＆value2）; 错误！= nil { 
    //句柄错误
}
```