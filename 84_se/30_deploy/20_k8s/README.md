# 云上 k8s 部署

## 部署MySQL

同上例

## 部署 k8s

- 上传镜像到 docker.io

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
./scripts/conf2yaml.sh
kubectl apply -f manifests/config.yaml
kubectl apply -f manifests/cert.yaml
kubectl apply -f apiserver.yaml  # 这里的k8s采用LB作为service，与本地不同
```

- 为 CLB 添加对外的安全组

- 通过 CLB 的 VIP 接入
```shell
curl -X GET http://VIP:8080/v1/users
./test.sh api::test::user
```

