## 特别说明

本项目基于 [go-gin-example](https://github.com/EDDYCJY/go-gin-example) 项目，并在此基础上进行修改和优化。

- 更新项目依赖，更新使用最新稳定版的依赖；
- 优化了部分旧的 api， 优化了部分接口调用和 swagger 不匹配等问题；
- 增加了前端项目 `/frontend`, 进入 `/frontend` 项目后，使用 `pnpm install` 安装依赖，使用 `pnpm run dev` 启动项目；
- 同时，给出了相关测试 http 例子，在目录 `/docs/test.http`，在 `vscode` 安装插件 `rest client` 即可使用；


--- 

# Go Gin Example [![rcard](https://goreportcard.com/badge/github.com/EDDYCJY/go-gin-example)](https://goreportcard.com/report/github.com/EDDYCJY/go-gin-example) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/EDDYCJY/go-gin-example) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/EDDYCJY/go-gin-example/master/LICENSE)

An example of gin contains many useful features




[简体中文](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)

## Installation
```
$ go get github.com/EDDYCJY/go-gin-example
```

## How to run

### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://github.com/EDDYCJY/go-gin-example/blob/master/docs/sql/blog.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

### Run front project

```
$ cd frontend
$ pnpm install
$ pnpm run dev
```


### Test Apis

- use VS Code Editor, install plugin `rest client`
- `cd docs` and open `test.http` file
- click `Send Request` button to test apis


Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/EDDYCJY/go-gin-example/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/EDDYCJY/go-gin-example/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/EDDYCJY/go-gin-example/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/EDDYCJY/go-gin-example/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/EDDYCJY/go-gin-example/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/EDDYCJY/go-gin-example/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/EDDYCJY/go-gin-example/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/EDDYCJY/go-gin-example/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/EDDYCJY/go-gin-example/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/EDDYCJY/go-gin-example/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/EDDYCJY/go-gin-example/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```
Swagger doc

![image](https://i.imgur.com/bVRLTP4.jpg)

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis