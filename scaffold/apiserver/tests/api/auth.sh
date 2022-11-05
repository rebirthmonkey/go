#!/usr/bin/env bash

INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"  # Docker port
SECURE_SERVER="127.0.0.1:8443"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Get
DCURL="curl -f -s -XDELETE" # Delete

insecure::basic()  # 运行：./tests/api/auth.sh insecure::basic
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  # 2. 用 admin，如果有 test00、test01 用户先清空
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test00; echo
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 3. 用 admin，创建 test00、test01 用户
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test01"},"password":"User@2022","nickname":"01","email":"test01@gmail.com","phone":"1306280xxxx"}'; echo

  # 4. 创建 test00 的 basic auth header
  basic00=$(echo -n 'test00:User@2022'|base64)
  Header00="-HAuthorization: Basic ${basic00}"

  # 5. 用 test00，列出所有用户
  ${RCURL} "${Header00}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo

  # 6. 用 test00，获取 test01 用户的详细信息
  ${RCURL} "${Header00}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 7. 用 test00，修改 test01 用户
  ${UCURL} "${Header}" "${Header00}" http://${INSECURE_SERVER}/v1/users/test01 \
    -d'{"nickname":"test01_modified","email":"test00_modified@foxmail.com","phone":"1306280xxxx"}'; echo

  # 8. 用 test00，删除 test00、test01 用户
  ${DCURL} "${Header00}" http://${INSECURE_SERVER}/v1/users/test01; echo
  ${DCURL} "${Header00}" http://${INSECURE_SERVER}/v1/users/test00; echo
}

insecure::login()
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/login \
    -d'{"username":"admin","password":"Admin@2021"}' |  jq '.token' | sed 's/"//g'
}

insecure::jwt()  # 运行：./tests/api/auth.sh insecure::jwt
{
  # 0. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  # 1. 用 admin 的 basic，如果有 test01 用户先清空
  ${DCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 2. 用 admin 的 basic，创建 test01 用户
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test01"},"password":"User@2022","nickname":"01","email":"test01@gmail.com","phone":"1306280xxxx"}'; echo

  # 3. 用 admin 的 basic，获得 admin 的 JWT token
  JWT="-HAuthorization: Bearer $(insecure::login)"

  # 4. 用 admin 的 JWT，如果有 secret1 密钥则先清空
  ${DCURL} "${Header}" "${JWT}" http://${INSECURE_SERVER}/v1/secrets/secret1; echo

  # 5. 用 admin 的 JWT，创建用于 test01 的密钥 secret1
  ${CCURL} "${Header}" "${JWT}" http://${INSECURE_SERVER}/v1/secrets \
    -d'{"metadata":{"name":"secret1"},"username":"test01", "expires":0,"description":"test01 secret"}'; echo

  # 6. 获取 test01 的秘钥 secret1
  basic01=$(echo -n 'test01:User@2022'|base64)
  HeaderBasic01="-HAuthorization: Basic ${basic01}"
  token=$(${CCURL} "${Header}" "${HeaderBasic01}" http://${INSECURE_SERVER}/login \
    -d'{"username":"test01","password":"User@2022"}' |  jq '.token' | sed 's/"//g')
  JWT01="-HAuthorization: Bearer ${token}"

  # 7. 用 test01 的秘钥 secret1，列出所有密钥
  ${RCURL} "${Header}" "${JWT01}" http://${INSECURE_SERVER}/v1/secrets; echo

  # 8. 用 secret1 获取 secret1 密钥的详细信息
  ${RCURL} "${Header}" "${JWT01}" http://${INSECURE_SERVER}/v1/secrets/secret1; echo

  # 9. 用 secret1，修改 secret1 密钥
  ${UCURL} "${Header}" "${JWT01}" http://${INSECURE_SERVER}/v1/secrets/secret1 \
    -d'{"expires":0,"description":"test01 secret(modified)"}'; echo

  # 10. 用 secret1 删除 secret1 密钥
  ${DCURL} "${Header}" "${JWT01}" http://${INSECURE_SERVER}/v1/secrets/secret1; echo
}


if [[ "$*" =~ insecure:: ]];then
  eval $*
fi
