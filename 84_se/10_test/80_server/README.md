# API Server

API Server 是一个基于 Go 语言开发的 demo，用于对资源访问进行 CRUD 操作。

API Server 同时也具有以下能力：

1. 讲解如何用 Go 做企业级应用的开发，配合理论部分，包含了 Go 项目开发各个知识点和构建思路的讲解，也会包含一线研发经验和建议。

2. 作为一个开发脚手架，供开发者克隆后二次开发，快速构建自己的 API 应用。

## 功能特性

本项目用到了Go 应用开发的大部分核心技能点，主要包括 Go 应用参数配置、启动运行、Gin    及 Grpc Web 服务、Mysql 等数据库对接、Log、错误处理、目录及代码规范、测试、自动构建等方面。

## 软件架构

<img src="docs/images/image-20220827180807512.png" alt="image-20220827180807512" style="zoom: 50%;" />

## 快速开始

### 依赖检查

1. 服务器能访问外网

2. 操作系统：MacOS 或 CentOS Linux 8.x (64-bit)

> 本安装脚本基于 MacOS 安装，建议你选择 CentOS 8.x 系统。其它 Linux 发行版也能安装，不过需要手动调试。

### 快速部署

快速部署请参考：[APIServer 部署指南](docs/guide/installation/README.md)

### 构建

如果你需要重新编译本项目，可以执行以下步骤：

```shell
cd $(pwd) # 进入当前目录
make
```

构建后的二进制文件保存在 `_output/platforms/linux/amd64/` 目录下。

## 使用指南

[User Guide](docs/guide/README.md)

## 如何贡献

欢迎贡献代码，贡献流程可以参考 [developer's documentation](docs/devel/development.md)。

## 社区

You are encouraged to communicate most things via pull requests.

## 关于作者

- Wukong SUN <rebirthmonkey@gmail.com>

## 许可证

APIServer is licensed under the MIT. See [LICENSE](LICENSE) for the full license text.
