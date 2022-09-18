# 云上 VM 部署

- 上传应用及配置文件：其中 apiserver 执行文件为先前 build 的 _output 文件。

```shell
scp -i tam-account/ruan.pem -r ./* root@1.14.254.147:/root
ssh -i tam-account/ruan.pem root@1.14.254.147
cd /etc & mkdir apiserver
mv /root/configs/* .
```

- 启动应用

```shell
cd /root
./apiserver -c /etc/apiserver/apiserver.yaml
```

- 测试接口

```shell
curl -X GET http://1.14.254.147:8080/v1/users
./test api::test::user
```
