# Gin HTTPS Server

与 insecure server 唯一不同之处在于，curl 命令需要加上 --insecure flag

```shell
curl --insecure -X GET https://127.0.0.1:8443/v1/users
```

一次性测试所有命令

```bash
cat configs/config.yaml | sed "s#{{CERT-FILE}}#$(pwd)/configs/cert/server.pem#g" | sed "s#{{PRIVATE-KEY-FILE}}#$(pwd)/configs/cert/server.key#g"  > configs/config-out.yaml
go run cmd/apiserver.go -c configs/config-out.yaml
sleep 10
#create
curl --insecure -X POST -H "Content-Type: application/json" \
-d '{"metadata":{"name":"user99", "password":"admin"},"description":"admin user"}' \
http://127.0.0.1:8080/v1/users
# list
curl --insecure -X GET http://127.0.0.1:8080/v1/users
# get
curl --insecure -X GET http://127.0.0.1:8080/v1/users/user99
# update
curl --insecure -X PUT -H "Content-Type: application/json" \
-d '{"metadata":{"name":"user99"},"nickname":"xxx"}' \
http://127.0.0.1:8080/v1/users/user99
# delete
curl --insecure -X DELETE http://127.0.0.1:8080/v1/users/user99
```

> 执行完后使用`kill -9 $!`结束后台任务
