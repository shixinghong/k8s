## admission-webhook

- 本地测试
```shell
git clone https://github.com/shixinghong/k8s.git
cd k8s && go mod tidy 
go run admission-webhook/main.go  # 启动服务端
go run admission-webhook/client/client.go # 启动客户端
```
