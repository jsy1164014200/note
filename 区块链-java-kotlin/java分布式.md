# java分布式 开发

1. web : structs, springMVC
2. service: spring
3. dao: Hibernate, MyBatis, spring Database JPA

ssh  ssm

Spring Boot 微服务架构 





java8 配合 gradle4.1 

## gradle

> gradle.org 官网

build.gradle.kts



```kotlin
 Project 和 task 是 Gradle领域的两个主要对象.
 gradle构建的时候, 会根据 配置文件 创建一个 project 实例,执行project
 配置文件中所有代码都会通过 task任务 方式插入到 project中

 project实例 可以再 这个文件里面 显示调用

 task
 application 包含很多我们常用的任务task

 task 生命周期: 1. 扫描 2. 执行
 gradle 会找到每一个任务, 先执行 每一个 任务中闭包的逻辑, 再执行任务,除非先指定执行的doFirst


project.task("task1:build") {
    doFirst {
        println("开始编译任务1")
    }
}

project.task("task2:build") {
    doFirst {
        println("开始编译任务2")
    }
}.dependsOn("task1:build")

project.task("task3:build") {
    doFirst {
        println("开始编译任务3")
    }
}.dependsOn("task2:build")


tasks{
    "opendoor" {
        group = "大象放冰箱"
        doFirst {
            println("xxxx")
        }
    }
    "putelephant" {
        group = "大象放冰箱"
        doFirst {
            println("xxxx")
        }
    }.dependsOn("opendoor")
    "closedoor" {
        group = "大象放冰箱"
        doFirst {
            println("xxxx")
        }
    }
}

 project中的 properties有一些工程相关的属性

task("打印默认属性") {
    doFirst {
        project.properties.forEach {t,any ->
            println("$t----$any")
        }
    }
}
```



```
// Gradle 增量更新
//inputs.dir()
//inputs.file()
//outputs.dir()
//outputs.file()

task("拷贝工作量") {
    // 指定输入源,与输出源
    inputs.dir("src")
    outputs.file("info.txt")

    doFirst {
        var dir = fileTree("src")
        var infoTxt = File("info.txt")
        dir.forEach {
            if (it.isFile) {
                infoTxt.appendText(it.name)
                infoTxt.appendText("\n")
            }
        }
        println("xx")
    }
}
```



```
// 常见的插件
plugins {  // 插件 ,定义了很多可执行的任务
    application  // gradle 自带的 插件 (类似的 还有 java 插件 ,但是java插件没有 发布的 作用) ( war插件 ,打包)
    kotlin("jvm")  // 支持 kotlin的插件
    // 还有很多 第三方的 插件 ,在官网可以查到
}

application {
    mainClassName = "MainKt"
}

repositories {
    mavenCentral()
    jcenter()
}

//关闭 gradle 的默认处理方案
//configuration.all {
//    resolutionStragegy {
//        failOnVersionConflict()
//    }
//}
// 常见的 jar仓库
// mavenCentral()
// Jcenter
dependencies {
    compile(kotlin("stdlib"))
    // https://mvnrepository.com/artifact/org.apache.httpcomponents/httpclient
//    compile group: 'org.apache.httpcomponents', name: 'httpclient', version: '4.5.7'
    compile("org.apache.httpcomponents","httpclient","4.5.7")

    // https://mvnrepository.com/artifact/junit/junit
    testCompile("junit", "junit", "4.13-beta-2")

}
// 依赖分成 编译时依赖, 测试时依赖(TDD测试驱动开发)
```



# Aop编程

> 在不改变源码的情况下对原有的功能进行扩展.

> 有装饰器模式,代理模式.

> aop底层用的是动态代理

> dl注入



# 微服务

> 围绕业务进行组织

1. 去中心化数据管理
2. 基础设施的自动化(自动化测试)
3. 演进式设计

https://www.martinfowler.com/articles/microservices.html

微服务来历



# spring cloud

> 快速构建分布式系统的工具集,快速构建微服务

## 服务注册中心

1. 服务注册与发现
   - Eureka Server 服务注册中心
   - Eureka Client 服务注册者
2. 





















































