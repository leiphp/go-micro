# go-micro
go-micro是一个微服务项目，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点，让你快速了解go微服务项目

## 项目说明
* 服务说明  `开启服务之前，需要开启etcd注册服务2379`  
1. web.bat web管理服务,端口8082
2. gw.bat 网关服务,端口8080  
3. user.bat 用户服务,端口9090 
4. course.bat 课程服务，端口9091  
5. coursehttp 课程http服务，端口9000 
  
* 系统配置文件  
app.yaml
* 业务配置文件全部从nacos取  
部分配置修改--要重启  
部分配置修改--直接覆盖  

## 安装protobuf相关工具  
go get github.com/golang/protobuf/protoc-gen-go

## go-micro自己生成工具  
go get github.com/micro/protoc-gen-micro/v2

## 生成proto
protoc --proto_path=src/protos  --micro_out=src/Users --go_out=src/Users Users.proto

## 下载microv2工具集  
go get github.com/micro/micro/v2

## 卸载包
go clean -i github.com/micro/micro

## 网关及服务
```
micro.Name("go.micro.api.user"))  //服务名  
set MICRO_API_NAMESPACE=go.micro  //网关名  
```

## 接口预览

### 1-1 使用`网关`调用`user`服务

**请求方式**

`POST` `http://localhost:8080/user/userService/test`

**Header**  
`Content-Type: application/json`

**Body参数说明**  

|参数|类型|说明|  
|-|-|-|  
|id|string| 1-1 用户id，如`239933`|  

**Response返回示例**  

```
{
    "ret": "users239933"
}
```

### 1-2 使用`网关`调用`course`服务

**请求方式**

`POST` `http://localhost:8080/course/courseService/listForTop`

**Header**  
`Content-Type: application/json`

**Body参数说明**  

|参数|类型|说明|  
|-|-|-|  


**Response返回示例**  

```
{
    "result": [
        {
            "course_id": 101,
            "course_name": "java课程"
        },
        {
            "course_id": 102,
            "course_name": "php课程"
        }
    ]
}
```
## nacos配置中心
配置不要有制表符，一定用4个空格  

格式一：  
```yaml
mysql:
    ip: 127.0.0.1
    port: 3306
redis:
    ip: 127.0.0.1
    port: 6379
```
格式二：  
```go
mysql:
    dsn: root:123456@tcp(localhost:3306)/gomicro?charset=utf8mb4
    maxidle: 5
    maxopen: 10
redis:
    ip: 127.0.0.1
    port: 6379
```