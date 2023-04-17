#!/bin/bash

# 写一个 shell 脚本，用于通过c调用 HTTP GET http://127.0.0.1:8080/healthz 接口
# 1. 定义变量
URL="http://127.0.0.1:8080/healthz"

# 2. 调用结果
response=$(curl -X GET $URL)

# 3. 打印结果
echo $response
