#!/usr/bin/env bash


API_ROOT=.
INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"  # Docker port
SECURE_SERVER="127.0.0.1:8443"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Get
DCURL="curl -f -s -XDELETE" # Delete

# Print out some info that isn't a top level status line
api::log::info() {
  local V="${V:-0}"
  if [[ ${IAM_VERBOSE} < ${V} ]]; then
    return
  fi

  for message; do
    echo "${message}"
  done
}

api::test::basic()  # 运行：./tests/api/auth.sh api::test::basic
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=`echo -n 'admin:Admin@2021'|base64`
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  # 2. 用 admin，如果有 test00、test01 用户先清空
  ${DCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test00; echo
  ${DCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 3. 用 admin，创建 test00、test01 用户
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test01"},"password":"User@2022","nickname":"01","email":"test01@gmail.com","phone":"1306280xxxx"}'; echo

  # 4. 创建 test00 的 basic auth header
  basic00=`echo -n 'test00:User@2022'|base64`
  Header00="-HAuthorization: Basic ${basic00}"

  # 5. 用 test00，列出所有用户
  ${RCURL} "${Header}" "${Header00}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo

  # 6. 用 test00，获取 test01 用户的详细信息
  ${RCURL} "${Header}" "${Header00}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 7. 用 test00，修改 test01 用户
  ${UCURL} "${Header}" "${Header00}" http://${INSECURE_SERVER}/v1/users/test01 \
    -d'{"nickname":"test01_modified","email":"test00_modified@foxmail.com","phone":"1306280xxxx"}'; echo

  # 8. 用 test00，删除 test00、test01 用户
  ${DCURL} "${Header}" "${Header00}" http://${INSECURE_SERVER}/v1/users/test01; echo
  ${DCURL} "${Header}" "${Header00}" http://${INSECURE_SERVER}/v1/users/test00; echo

  api::log::info "$(echo -e '\033[32mcongratulations, /v1/user test passed!\033[0m')"
}

api::test::login()
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=`echo -n 'admin:Admin@2021'|base64`
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_AUTHSERVER}/login \
    -d'{"username":"admin","password":"Admin@2021"}' |  jq '.token' | sed 's/"//g'
}

api::test::jwt()  # 运行：./tests/api/auth.sh api::test::jwt
{
  # 1. 用 admin 的用户名、密码，获得 JWT token
  JWT="-HAuthorization: Bearer $(api::test::login)"

  # 2. 用 admin 的 JWT，如果有 secret0 密钥则先清空
  ${DCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets/secret0; echo

  # 3. 用 admin 的 JWT，创建 secret0 密钥
  ${CCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets \
    -d'{"metadata":{"name":"secret0"},"expires":0,"description":"admin secret"}'; echo

  # 4. 用 admin 的 JWT，列出所有密钥
  ${RCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets; echo

  # 6. 用 admin 的 JWT，获取 secret0 密钥的详细信息
  ${RCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets/secret0; echo

  # 7. 用 admin 的 JWT，修改 secret0 密钥
  ${UCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets/secret0 \
    -d'{"expires":0,"description":"admin secret(modified)"}'; echo

  # 8. 用 admin 的 JWT，删除 secret0 密钥
  ${DCURL} "${Header}" "${JWT}" http://${INSECURE_AUTHSERVER}/v1/secrets/secret0; echo

  api::log::info "$(echo -e '\033[32mcongratulations, /v1/user test passed!\033[0m')"
}


if [[ "$*" =~ api::test:: ]];then
  eval $*
fi
