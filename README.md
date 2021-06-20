# go-micro
go-micro是一个微服务项目，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点，让你快速了解go微服务项目

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