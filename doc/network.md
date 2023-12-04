# 网络知识学习

## 通过代理访问目标网站

使用 `curl` 命令通过 SOCKS5 代理访问网站的示例如下：

```shell
curl --socks5-hostname <proxy_host>:<proxy_port> http://www.baidu.com
```

在这个命令中，你需要将 `<proxy_host>` 替换为 SOCKS5 代理服务器的主机名或 IP 地址，将 `<proxy_port>` 替换为代理服务器的端口号。然后，`http://www.baidu.com` 是你想访问的目标网站。

例如，如果你的 SOCKS5 代理主机是 `127.0.0.1`，端口是 `1080`，你可以这样执行命令：

```shell
curl --socks5-hostname 127.0.0.1:1080 http://www.baidu.com
```

请注意，`--socks5-hostname` 参数用于指定 SOCKS5 代理服务器。如果你使用的是其他代理协议，如 HTTP 代理，可以使用相应的参数（例如 `--proxy` 或 `--http-proxy`）来指定。

* 本项目即可采用类似的方法验证服务器的正确性

### Introduction about SOCKS Protocal

 防火墙系统就像一个应用层的网关，提供访问控制（TELNET, FTP, SMTP, SSH）

TELNET类似于ssh，可以通过与远程主机构建ssh连接实现远程操控，将本地的指令发送至远端并将结果返回至本地展示。FTP可以在本地与FTP服务器之间建立连接，并通过连接传输文件。**这些协议都通过构建连接实现安全可靠的传输，为高权限的用户提供的内容传递。**SOCKS可为这种类型的协议提供了基础的框架，在基于socks的服务器上实现类似的功能。

C/S架构的出现标志着我们需要一种更强的身份验证以及访问控制，SOCKS为C/S架构提供了更为通用的框架

## Procedure for TCP-based clients

1. **连接阶段：**当一个tcp-based的客户端想建立一条与目标服务器之间的连接，这条连接会经过防火墙系统，那么它必须先连接到SOCKS服务器上的对应的SOCKS端口（1080），端口连接成功进入**协商阶段**
2. **协商阶段：**

# Go语言学习

## 1. 搭建一个项目

```shell
go mod init "github.com/repo_name"
```

go.mod文件可以作为项目的包管理器，例如java中的pom.xml Python中的requirement.txt，在go.mod中包含了项目所需要的依赖，当项目使用到了其他的包，go mod会将依赖关系添加到文件中

```shell
go mod tidy
```

go mod tidy会确保项目中模块间的依赖关系，保持依赖间版本的一致性并清除不需要的依赖

## 2. 如何设定项目入口

Package main就代表了项目的入口，主包中的main函数则是入口函数

相同包的文件可以不用引入直接使用，而不同包的函数不可直接使用

当我们想在主包中引入其他包下的文件时，需要通过import的方式引入

```css
.
├── README.md
├── doc
├── go.mod
├── main.go
└── util
    └── util.go
```

