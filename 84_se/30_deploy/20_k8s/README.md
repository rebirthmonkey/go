# 云上 k8s 部署

## 部署MySQL

同上例

## 部署 k8s

- 上传镜像到 docker.io

- 上传配置文件到 TKE 集群的各个节点：整个 /configs 目录

```shell
scp -i xxx.pem -r ./configs root@175.178.199.5:/data/apiserver  # k8s node
ssh -i xxx.pem root@175.178.199.5
```

- 本地 kubectl switch 到云上的 context

```shell
cp configs/cls-lquwlzum-config /tmp/
export KUBECONFIG=/tmp/cls-lquwlzum-config
kubectl config --kubeconfig=configs/cls-lquwlzum-config get-contexts
kubectl config --kubeconfig=configs/cls-lquwlzum-config use-context  cls-lquwlzum-100006513780-context-default
kubectl cluster-info
```

- 通过本地 kubectl 启动云上 workload：其中 service 采用 LoadBalancer 可确保外网访问

```shell
kubectl apply -f apiserver.yaml
```

- 为 CLB 添加对外的安全组

- 通过 CLB 的 VIP 接入
```shell
curl -X GET http://VIP:8080/v1/users
./test.sh api::test::user
```

