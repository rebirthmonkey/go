# API 介绍

APIServer 接口文档，相关参考文档如下：

- [更新历史](./CHANGELOG.md)
- [API 接口文档规范](./api_specification.md)
- [通用说明](./generic.md)
- [数据结构](./struct.md)
- API 接口：
    - [User 相关接口](./user.md)
 - [业务错误码](./error_code_generated.md)

## API 概览

### User 相关接口

| 接口名称                                                      | 接口功能     |
| ------------------------------------------------------------- | ------------ |
| [POST /v1/users](./user.md#创建用户)                          | 创建用户     |
| [DELETE /v1/users/:name](./user.md#删除用户)                  | 删除用户     |
| [PUT /v1/users/:name](./user.md#修改用户属性)                 | 修改用户属性 |
| [GET /v1/users/:name](./user.md#查询用户信息)                 | 查询用户信息 |
| [GET /v1/users](./user.md#查询用户列表)                       | 查询用户列表 |



