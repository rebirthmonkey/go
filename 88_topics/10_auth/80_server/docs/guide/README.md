# 使用手册

## 配置

### 配置运行环境

```bash
cd 80_server
make help
make tools # install all the tools on the local host
make tidy # go mod tidy
```

### 配置参数

根据运行方式（本地、docker 或 k8s）需要对

 `configs/apiserver.yaml`进行配置

- secure/tls/cert-file：
- secure/tls/private-key-file：
- mysql/host：

Makefile 进行配置

- ROOT_PACKAGE

## 基本操作

### Run

```bash
make run
```

### Build

Build current platform

```bash
make clean
make build
make clean
```

Build multiple platforms

```bash
make clean
make build.multiarch
make clean
```

### Docker Image & Run

```bash
make image
docker run -d -v $(pwd)/configs/:/etc/apiserver/ -p 8080:8080 wukongsun/apiserver-amd64
```

注意：因为运行环境不同，在 Docker 镜像启动时，需要修改 apiserver.yaml 的 3 个参数：

- gin/secure/cert-file
- gin/secure/ private-key-file
- mysql/host

### K8s Run

```shell
make deploy
```

### Test

```bash
make test
```

## 具体操作

### 简单接口测试

- 启动 apiserver

```shell
make run
```

- 启动 API 接口测试

```shell
make test.api
```



