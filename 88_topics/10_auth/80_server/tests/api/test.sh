#!/usr/bin/env bash

# The root of the build/dist directory
#IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
#[[ -z ${COMMON_SOURCED} ]] && source ${IAM_ROOT}/scripts/install/common.sh
#INSECURE_SERVER=${IAM_APISERVER_HOST}:${IAM_APISERVER_INSECURE_BIND_PORT}
#INSECURE_AUTHZSERVER=${IAM_AUTHZ_SERVER_HOST}:${IAM_AUTHZ_SERVER_INSECURE_BIND_PORT}

API_ROOT=.
INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"
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

api::test::user() # 用于无 auth 验证
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

  api::log::info "$(echo -e '\033[32mcongratulations, /v1/user test passed!\033[0m')"
}

api::test::secret()  # 运行：./tests/api/test.sh api::test::secret
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

  api::log::info "$(echo -e '\033[32mcongratulations, /v1/user test passed!\033[0m')"
}


if [[ "$*" =~ api::test:: ]];then
  eval $*
fi
