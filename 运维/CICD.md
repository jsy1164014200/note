# devop 之CICD



## gitlab-ci实践

实践

1. 在本地 使用 docker-compose  +  test 做 单元测试，效果测试
2. push 到 dev分支，做 cicd  中的 pipeline（在gitlab 上设置 repository 不允许任何人 push 到master，在general中设置，只有 pipeline通过后 才能 merge request）
3. 自己搭建 runner 服务器 做  生产的 代码风格检查，单元测试，部署
4. 发布 release 也就是部署到 生产服务器

git push origin <local_branch_name>:<remote_branch_name>

jenkins

介绍 gitlab ci 的一篇文章

https://scarletsky.github.io/2016/07/29/use-gitlab-ci-for-continuous-integration/



在gitlab ci服务器上做 测试 风格检查 测试部署

docker registry 私有服务器



pep8

test

docker test deploy 

docker image release  需要一个私有的 docker registry

![选区_100](/home/jsy/图片/选区_100.png)

![选区_098](/home/jsy/图片/选区_098.png)



## Travis CI

Travis 要求项目的根目录下面，必须有一个`.travis.yml`文件。这是配置文件，指定了 Travis 的行为。该文件必须保存在 Github 仓库里面，一旦代码仓库有新的 Commit，Travis 就会去找这个文件，执行里面的命令。

Travis 的运行流程很简单，任何项目都会经过两个阶段。

> - install 阶段：安装依赖
> - script 阶段：运行脚本

### install 字段

`install`字段用来指定安装脚本。

> ```javascript
> install: ./install-dependencies.sh
> ```

如果有多个脚本，可以写成下面的形式。

> ```javascript
> install:
>   - command1
>   - command2
> ```

上面代码中，如果`command1`失败了，整个构建就会停下来，不再往下进行。

如果不需要安装，即跳过安装阶段，就直接设为`true`。

> ```javascript
> install: true
> ```

### script 字段

`script`字段用来指定构建或测试脚本。

> ```javascript
> script: bundle exec thor build
> ```

如果有多个脚本，可以写成下面的形式。

> ```javascript
> script:
>   - command1
>   - command2
> ```

注意，`script`与`install`不一样，如果`command1`失败，`command2`会继续执行。但是，整个构建阶段的状态是失败。

如果`command2`只有在`command1`成功后才能执行，就要写成下面这样。

> ```javascript
> script: command1 && command2
> ```



### 部署

`script`阶段结束以后，还可以设置[通知步骤](https://docs.travis-ci.com/user/notifications/)（notification）和[部署步骤](https://docs.travis-ci.com/user/deployment/)（deployment），它们不是必须的。

部署的脚本可以在`script`阶段执行，也可以使用 Travis 为几十种常见服务提供的快捷部署功能。比如，要部署到 [Github Pages](https://docs.travis-ci.com/user/deployment/pages/)，可以写成下面这样。

> ```javascript
> deploy:
>   provider: pages
>   skip_cleanup: true
>   github_token: $GITHUB_TOKEN # Set in travis-ci.org dashboard
>   on:
>     branch: master
> ```

其他部署方式，请看[官方文档](https://docs.travis-ci.com/user/deployment/)。

### 钩子方法

Travis 为上面这些阶段提供了7个钩子。

> - before_install：install 阶段之前执行
> - before_script：script 阶段之前执行
> - after_failure：script 阶段失败时执行
> - after_success：script 阶段成功时执行
> - before_deploy：deploy 步骤之前执行
> - after_deploy：deploy 步骤之后执行
> - after_script：script 阶段之后执行

完整的生命周期，从开始到结束是下面的流程。

> 1. before_install
> 2. install
> 3. before_script
> 4. script
> 5. after*success or after*failure
> 6. [OPTIONAL] before_deploy
> 7. [OPTIONAL] deploy
> 8. [OPTIONAL] after_deploy
> 9. after_script

下面是一个`before_install`钩子的例子。

> ```javascript
> before_install:
>   - sudo apt-get -qq update
>   - sudo apt-get install -y libxml2-dev
> ```

### 运行状态

最后，Travis 每次运行，可能会返回四种状态。

> - passed：运行成功，所有步骤的退出码都是`0`
> - canceled：用户取消执行
> - errored：`before_install`、`install`、`before_script`有非零退出码，运行会立即停止
> - failed ：`script`有非零状态码 ，会继续运行





可以限制构建的并行数目

可以设置安全列表

```yaml
branches:
	except:
		- dev
	only:
		- master
```

```bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
```



一篇讲node 自动部署的

https://juejin.im/post/5a9e1a5751882555712bd8e1



# 记录一个travis 的bug

before_install:

- openssl aes-256-cbc -K $encrypted_8cc0db86b412_key -iv $encrypted_8cc0db86b412_iv -in id_rsa.enc -out ~/    .ssh/id_rsa -d
- openssl aes-256-cbc -K $encrypted_8cc0db86b412_key -iv $encrypted_8cc0db86b412_iv -in id_rsa.enc -out ~\/    .ssh/id_rsa -d  错误的写法

