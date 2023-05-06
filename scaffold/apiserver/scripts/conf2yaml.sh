#!/bin/bash

# 定义变量
CONFIG_MAP_NAME=config  # ConfigMap 名称
CONFIG_FILE_PATH=./configs/apiserver.yaml  # 配置文件路径
NAMESPACE=default  # 命名空间

# 读取配置文件内容
CONFIG_CONTENT=$(cat $CONFIG_FILE_PATH)

# 定义变量
AUTHZ_CONFIG_MAP_NAME=authz_config  # ConfigMap 名称
AUTHZ_CONFIG_FILE_PATH=./configs/authz.yaml  # 配置文件路径
NAMESPACE=default  # 命名空间

# 读取配置文件内容
AUTHZ_CONFIG_CONTENT=$(cat $AUTHZ_CONFIG_FILE_PATH)

# 生成 ConfigMap YAML 文件
cat <<EOF > manifests/config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: $CONFIG_MAP_NAME
  namespace: $NAMESPACE
data:
  apiserver.yaml: |
$(echo "$CONFIG_CONTENT" | sed 's/^/    /')
  authz.yaml: |
$(echo "$AUTHZ_CONFIG_CONTENT" | sed 's/^/    /')
EOF

# 设置变量
CERT_FILE=./configs/cert/server.pem
CERT_NAME=cert
SECRET_FILE=./configs/cert/server.key
SECRET_NAME=secret

# 从证书文件中读取内容
CERT_CONTENT=$(cat $CERT_FILE | base64)
SECRET_CONTENT=$(cat $SECRET_FILE | base64)

# 生成Secret的YAML文件
cat <<EOF > manifests/cert.yaml
apiVersion: v1
kind: Secret
metadata:
  name: $CERT_NAME
  namespace: $NAMESPACE
type: Opaque
data:
  server.pem: $CERT_CONTENT
  server.key: $SECRET_CONTENT
EOF
