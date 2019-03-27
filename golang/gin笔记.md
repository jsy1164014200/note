# govendor 包管理工具

golang编码解码 简直是麻烦的不行，，，难受
```go
package tools

import (
	"github.com/axgle/mahonia"
)

// 编码
func Encode(str string,formater string) []byte{
	coder := mahonia.NewDecoder(formater)
	_,result,_ := coder.Translate([]byte(str),true)   // 因为byte类型转换默认是utf8的字节类型
	return result
}

func Decode(bytes []byte,formater string) string{ // 因为默认go只支持 utf8
	str := string(bytes)
	coder := mahonia.NewDecoder(formater)
	result := coder.ConvertString(str)
	return result

}
```

hex十六进制与byte数组的互换

govendor 类似node中的 npm 包管理工具

主要用来管理go用到的包
> go get github.com/kardianos/govendor
```sh
init    # 创建 vendor 文件夹和 vendor.json 文件
list     #列出已经存在的依赖包
add     # 从 $GOPATH 中添加依赖包，会加到 vendor.json
update   #从 $GOPATH 升级依赖包
remove   #从 vendor 文件夹删除依赖
status   #列出本地丢失的、过期的和修改的package
fetch  # 从远端库增加新的，或者更新 vendor 文件中的依赖包
sync     #Pull packages into vendor folder from remote repository with revisions
migrate # Move packages from a legacy tool to the vendor folder with metadata.
get     #类似 go get，但是会把依赖包拷贝到 vendor 目录

go tool commands that are wrapped:
      `+<status>` package selection may be used with them
    fmt, build, install, clean, test, vet, generate, tool
```

下载gopm 用来从一些镜像网站上下载包
>go get -u github.com/gpmgo/gopm

基本使用
> gopm get -g golang.ory/x/x  

-g 用来指定是 下载在GOPATH还是当前vendor目录 类似 npm中的 --save -dev

# gin简述

gin极度类似express，路由绑定方式，中间件，监听等……

1. 解析get
> c.DefaultQuery() c.Query()

2. 解析POST
> c.PostForm() c.DefaultPostForm()

3. 解析 POST + query
> c.DefaultQuery c.PostForm

4. Map 解析
> c.QueryMap("ids") c.PostFormMap("ids")

5. 上传file
```go
c.FormFile()
c.SaveUploadFile(file,dst)

router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
```

# gin常用功能

## 1. Grouping routes
也就是路由的嵌套
```go
v1 := router.Group("/v1")
{
  v1.POST("/login", loginEndpoint)
  v1.POST("/submit", submitEndpoint)
  v1.POST("/read", readEndpoint)
}

```
使用不带中间件的router
r := gin.New()

或者使用默认的，自带了弄个了 and recovery
// Default With the Logger and Recovery middleware already attached
r := gin.Default()

## 2. 使用中间件
> r.Use(gin.Logger())
> r.Use(gin.Recovery())


```go
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
  // 只有authorized这个group路由才会使用下面的这个中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```

## 3. 书写日志文件
```go
func main() {
    // Disable Console Color, you don't need console color when writing the logs to file.
    gin.DisableConsoleColor()

    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // 如果想同时在文件，以及console写日志
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
```

## 4. 绑定模型

+ type  - MUST BIND 
  - methods -Bind ,BindJson,BindXml,BindQuery
  - behavior - 这些方法必须使用MustBindWith

+ type - should bind
  - methods - ShouldBind ShouldBindJson,ShouldBindXml,ShouldBIndQuery
  - behavior - ShouldBindWith

```go
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 
		
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</user>
	//	</root>)
	// router.POST("/loginXML", func(c *gin.Context) {
	// 	var xml Login
	// 	if err := c.ShouldBindXML(&xml); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
		
	// 	if xml.User != "manu" || xml.Password != "123" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		return
	// 	} 
		
	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// })

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 
		
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8000")
}
```

## 5. static文件路由

```go
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
```

## 6. 渲染HTML模板

```go
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
```

## 7. custom Middleware 中间件

```go
func Logger() gin.HandlerFunc{
  return func(c *gin.Context){
    t := time.Now()
    c.set("example","123")
    c.Next()
    latency := time.Since(t)
    log.Print(latency)
    status := c.Writer.Status()
    log.Println(status)
  }
}

r.Use(Logger())
```


配置http 
```go
func main() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router, // router 就是一个http.handle
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
```

set and get cookie
```go
func main() {

    router := gin.Default()

    router.GET("/cookie", func(c *gin.Context) {

        cookie, err := c.Cookie("gin_cookie")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

        fmt.Printf("Cookie value: %s \n", cookie)
    })

    router.Run()
}
```