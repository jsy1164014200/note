# docker bug 记录



1. 在容器内部 ports 暴露端口时 ，一定要 是监听 0.0.0.0：xxx

2. 个人觉得 expose 没有什么用，只要在一个 network namespace

