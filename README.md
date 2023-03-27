<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [bluebird](#bluebird)
- [中间件](#%E4%B8%AD%E9%97%B4%E4%BB%B6)
  - [Swagger](#swagger)
  - [Pporf](#pporf)
  - [Requestid](#requestid)
- [初始化](#%E5%88%9D%E5%A7%8B%E5%8C%96)
  - [linux](#linux)
  - [windows](#windows)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# bluebird
蓝色是一种温和、舒适的颜色，而蓝鸟象征着快乐和自由。

> 使用golang 1.18+的环境

# 中间件

## Swagger

```bash
# 安装swagger
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/hertz-contrib/swagger
go get github.com/swaggo/files
```

## Pporf

```bash
go get github.com/hertz-contrib/pprof
```

## Requestid

```bash
go get github.com/hertz-contrib/requestid
```



# 初始化

**在项目目录下**

1、需要安装swagger - 已安装可跳过

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2、检查依赖

```bash
go mod tidy
```

## linux

使用containerd，需要安装nerdctl

方式一：镜像

```bash
bash build.sh
nerdctl push ${image}
```

方式二：直接运行

```bash
swag init
go run github.com/stonebirdjx/bluebird
```

## windows

直接运行

```bash
swag init
go run github.com/stonebirdjx/bluebird
```
