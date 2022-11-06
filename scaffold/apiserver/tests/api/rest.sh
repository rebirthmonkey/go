#!/usr/bin/env bash

INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"
SECURE_SERVER="127.0.0.1:8443"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Get
DCURL="curl -f -s -XDELETE" # Delete

insecure::healthz()
{
  ${RCURL} http://${INSECURE_SERVER}/healthz
}

insecure::user() # 用于无 auth 验证
{
  # 1. 如果有 test00 用户先清空
  ${DCURL} http://${INSECURE_SERVER}/v1/users/test00; echo

  # 2. 创建 test00
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo

  # 3. 列出所有用户
  ${RCURL} "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo

  # 4. 获取 test00 用户的详细信息
  ${RCURL} http://${INSECURE_SERVER}/v1/users/test00; echo

  # 5. 修改 test00 用户
  ${UCURL} "${Header}" http://${INSECURE_SERVER}/v1/users/test00 \
    -d'{"nickname":"test00_modified","email":"test00_modified@foxmail.com","phone":"1306280xxxx"}'; echo

  # 6. 删除 test00 用户
  ${DCURL} http://${INSECURE_SERVER}/v1/users/test00; echo
}

insecure::secret()  # 运行：./tests/api/rest.sh secret
{
  # 1. 如果有 secret0 密钥则先清空
  ${DCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets/secret0; echo

  # 2. 创建 secret0 密钥
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets \
    -d'{"metadata":{"name":"secret0"},"expires":0,"description":"admin secret"}'; echo

  # 3. 列出所有 secret
  ${RCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets; echo

  # 4. 获取 secret0 密钥的详细信息
  ${RCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets/secret0; echo

  # 5. 修改 secret0 密钥
  ${UCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets/secret0 \
    -d'{"expires":0,"description":"admin secret(modified)"}'; echo

  # 6. 删除 secret0 密钥
  ${DCURL} "${Header}" http://${INSECURE_SERVER}/v1/secrets/secret0; echo
}

insecure::policy()  # 运行：./tests/api/rest.sh policy
{
  # 1. 如果有 policy0 密钥则先清空
  ${DCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies/policy0; echo

 # 2. 创建 policy0 策略
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"policy0"},"policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 3. 列出所有 policies
  ${RCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies; echo

  # 4. 获取 policy0 密钥的详细信息
  ${RCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies/policy0; echo

  # 5. 修改 policy0 策略
  ${UCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies/policy0 \
    -d'{"policy":{"description":"One policy to rule them all(modified).","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 6. 删除 policy0 策略
  ${DCURL} "${Header}" http://${INSECURE_SERVER}/v1/policies/policy0; echo
}


if [[ "$*" =~ insecure:: ]];then
  eval $*
fi
