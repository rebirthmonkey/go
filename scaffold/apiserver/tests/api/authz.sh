#!/usr/bin/env bash

INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"  # Docker port
SECURE_SERVER="127.0.0.1:8443"
INSECURE_AUTHZ_SERVER="127.0.0.1:8090"
#INSECURE_AUTHZ_SERVER="127.0.0.1:30090"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Get
DCURL="curl -f -s -XDELETE" # Delete

insecure::login()
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/login \
    -d'{"username":"admin","password":"Admin@2021"}' |  jq '.token' | sed 's/"//g'
}

insecure::authz-noauth()
{
  # 1. admin login
  token="-HAuthorization: Bearer $(insecure::login)"

  # 2 如果有 wkpolicy 策略先清空
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/policies/wkpolicy; echo

  # 3 创建 wkpolicy 策略
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"wkpolicy"},"policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 4. 调用 /v1/authz 完成资源鉴权
  $CCURL "${Header}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:maria","action":"delete","resource":"resources:articles:ladon-introduction","context":{"remoteIPAddress":"192.168.0.5"}}'
}

if [[ "$*" =~ insecure:: ]];then
  eval $*
fi