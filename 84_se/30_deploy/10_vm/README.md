# 云上 VM 部署

## 部署MySQL

- 在控制台申请 MySQL DB
- 导入数据：

```shell
mysql -h gz-cdb-audi4a1v.sql.tencentcdb.com -P 57102 -u root -p < configs/apiserver.sql
```

- 数据验证：

```shell
mysql -h gz-cdb-audi4a1v.sql.tencentcdb.com -P 57102 -u root -p  # 创建 xyzshop DB
```

```sql
use iam;
show tables;
select * from user;
```

## CVM部署应用

- 上传应用及配置文件：其中 apiserver 执行文件为先前 build 的 _output 文件。

```shell
cd 10_vm
scp -i xxx.pem -r ./* root@1.14.254.147:/data/apiserver
ssh -i xxx.pem root@1.14.254.147
cd /etc & mkdir apiserver
mv /data/apiserver/configs/* .
```

- 启动应用

```shell
cd /data/apiserver
./apiserver -c /etc/apiserver/apiserver.yaml
```

- 测试接口

```shell
curl -X GET http://1.14.254.147:8080/v1/users
./test api::test::user
```
