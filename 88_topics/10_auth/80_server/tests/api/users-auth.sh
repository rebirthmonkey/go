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

api::test::healthz()
{
  ${RCURL} http://${INSECURE_SERVER}/healthz
}

api::test::user()
{
  # 1. 创建 admin 的 basic auth header

  # 2. 用 admin，如果有 test00、test01 用户先清空
  ${DCURL} http://${INSECURE_SERVER}/v1/users/test00; echo

  # 3. 用 admin，创建 test00、test01 用户
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo

  # 4. 创建 test00 的 basic auth header

  # 5. 用 test00，列出所有用户
  ${RCURL} "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo

  # 6. 用 test00，获取 test01 用户的详细信息
  ${RCURL} http://${INSECURE_SERVER}/v1/users/test00; echo

  # 7. 用 test00，修改 test01 用户
  ${UCURL} "${Header}" http://${INSECURE_SERVER}/v1/users/test00 \
    -d'{"nickname":"test00_modified","email":"test00_modified@foxmail.com","phone":"1306280xxxx"}'; echo

  # 8. 用 test00，删除 test00、test01 用户
  ${DCURL} http://${INSECURE_SERVER}/v1/users/test00; echo

  api::log::info "$(echo -e '\033[32mcongratulations, /v1/user test passed!\033[0m')"
}


if [[ "$*" =~ api::test:: ]];then
  eval $*
fi
