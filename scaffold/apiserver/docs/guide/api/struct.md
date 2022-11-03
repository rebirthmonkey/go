# 数据结构

APIServer 系统数据结构。

## ObjectMeta

资源对象元数据，所有资源对象都具有此属性。注意：只有 `name` 是输入参数，其它全是输出参数。

| 参数名称  | 类型   | 必选 | 描述                     |
| --------- | ------ | ---- | ------------------------ |
| id        | uint64 | 否   | 资源 ID，唯一标识一个资源 |
| name      | String | 是   | 资源名称（输入参数）     |
| CreatedAt | String | 否   | 资源创建时间             |
| UpdatedAt | String     |   否   | 资源更新时间             |

## User2

查询用户列表接口中，返回的用户字段信息。

| 参数名称    | 类型                      | 描述               |
| ----------- | ------------------------- | ------------------ |
| metadata    | [ObjectMeta](struct.md#ObjectMeta) | REST 资源的功能属性 |
| nickname    | String                    | 昵称               |
| password    | String                    | 密码               |
| email       | String                    | 邮箱地址           |
| phone       | String                    | 电话号码           |
| totalPolicy | Uint64                    | 用户授权策略个数   |



