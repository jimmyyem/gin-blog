# gin-blog

#### 介绍
a simple blog based on gin framework

#### 软件架构
MySQL/MariaDB as database


#### 安装教程

1. 先导入database里所有tables
2. 安装swagger，并且生成docs目录
```cgo
    go install github.com/swaggo/swag/cmd/swag
    swagger init
```
3. go mod vendor

#### 使用说明

1.  启动MySQL in Docker.
2.  

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
