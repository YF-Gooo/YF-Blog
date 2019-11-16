## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)

# 添加https自签名
    mkdir cert
    openssl req -newkey rsa:4096 -nodes -keyout ./cert/server.key -out ./cert/server.csr
    openssl x509 -signkey ./cert/server.key -in ./cert/server.csr -req -days 365 -out ./cert/server.crt

