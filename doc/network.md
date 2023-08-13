# RFC文档学习

### Introduction

 

# Go语言学习

## 1. 搭建一个项目

```shell
go mod init "github.com/repo_name"
```

go.mod文件可以作为项目的包管理器，例如java中的pom.xml Python中的requirement.txt，在go.mod中包含了项目所需要的依赖，当项目使用到了其他的包，go mod会将依赖关系添加到文件中

## 2. 如何设定项目入口

