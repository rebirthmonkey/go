# CI-CD

需将 env.DOCKERHUB_USERNAME 环境变量替换为 Docker Hub 用户名。

简单介绍一下这个工作流。这里的 name 字段是工作流的名称，它会展示在 GitHub 网页上。

- on.push.branches 字段的值为 main，代表当 main 分支有新的提交之后，会触发工作流。
- env.DOCKERHUB_USERNAME 是为 Job 配置的全局环境变量，用作镜像 Tag 的前缀。
- jobs.docker 字段定义了一个任务，它的运行环境是 ubuntu-latest，并且由 7 个 Step 组成。jobs.docker.steps 字段定义了 7 个具体的执行阶段。要特别注意的是，uses 字段代表使用 GitHub Action 的某个插件，例如 actions/checkout@v3 插件会帮助 checkout 代码。在这个工作流中，这 7 个阶段会具体执行下面几件事。
- “Checkout”阶段负责将代码 checkout 到运行环境。
- “Set outputs”阶段会输出 sha_short 环境变量，值为 short commit id，这可以方便在后续阶段引用。
- “Set up QEMU”和“Set up Docker  Buildx”阶段负责初始化 Docker 构建工具链。
- “Login to Docker Hub”阶段通过 docker login 来登录到 Docker Hub，以便获得推送镜像的权限。要注意的是，with 字段是向插件传递参数的，这里传递了 username 和 password，值的来源分别是定义的环境变量 DOCKERHUB_USERNAME 和 GitHub Action  Secret，后者我们还会在稍后进行配置。
- “Build backend and push”和“Build frontend and  push”阶段负责构建前后端镜像，并且将镜像推送到 Docker Hub，在这个阶段中，传递了 context、push 和 tags 参数，context 和 tags 实际上就是 docker build 的参数。在 tags 参数中，通过表达式 `${{  env.DOCKERHUB_USERNAME }}` 和 `${{ steps.vars.outputs.sha_short }}` 分别读取了在 YAML 中预定义的 Docker Hub 的用户名，以及在“Set outputs”阶段输出的 short commit id。

.github/workflows/build.yaml 文件

```yaml

name: build

on:
  push:
    branches:
      - 'main'

env:
  DOCKERHUB_USERNAME: lyzhang1999

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      - name: Set outputs
        id: vars
        run: echo "::set-output name=sha_short::$(git rev-parse --short HEAD)"
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v2
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
      - name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ env.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}
      - name: Build backend and push
        uses: docker/build-push-action@v3
        with:
          context: backend
          push: true
          tags: ${{ env.DOCKERHUB_USERNAME }}/backend:${{ steps.vars.outputs.sha_short }}
      - name: Build frontend and push
        uses: docker/build-push-action@v3
        with:
          context: frontend
          push: true
          tags: ${{ env.DOCKERHUB_USERNAME }}/frontend:${{ steps.vars.outputs.sha_short }}
```

