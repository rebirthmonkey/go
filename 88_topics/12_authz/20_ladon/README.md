# Landon

## 简介

Ladon 是用 Go 语言编写的用于实现访问控制策略的库，类似于 RBAC（Role Based Access Control）和 ACL（Access Control Lists）。但与 RBAC 和 ACL 相比，Ladon 可以实现更细粒度的访问控制，并且能够在更为复杂的环境中（例如多租户、分布式应用程序和大型组织）工作。Ladon 解决了这个问题：在特定的条件下，谁能够 / 不能够对哪些资源做哪些操作。

### Policy

为了解决这个问题，Ladon 引入了授权策略。授权策略是一个有语法规范的文档，这个文档描述了谁在什么条件下能够对哪些资源做哪些操作。Ladon 可以用请求的上下文，去匹配设置的授权策略，最终判断出当前授权请求是否通过。

策略（Policy）由若干元素构成，用来描述授权的具体信息，你可以把它们看成一组规则。核心元素包括主题（Subject）、操作（Action）、效力（Effect）、资源（Resource）以及生效条件（Condition）。元素保留字仅支持小写，它们在描述上没有顺序要求。对于没有特定约束条件的策略，Condition 元素是可选项。一条策略包含下面 6 个元素：

- 主题（Subject），主题名是唯一的，代表一个授权主题。例如，“ken” or  “printer-service.mydomain.com”。
- 操作（Action），描述允许或拒绝的操作。
- 效力（Effect），描述策略产生的结果是“允许”还是“拒绝”，包括 allow（允许）和  deny（拒绝）。
- 资源（Resource），描述授权的具体数据。
- 生效条件（Condition），描述策略生效的约束条件。
- 描述（Description），策略的描述。

下面是一个 Ladon 的授权策略样例：

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:<peter|ken>", "users:maria", "groups:admins"],
  "actions" : ["delete", "<create|update>"],
  "effect": "allow",
  "resources": [
    "resources:articles:<.*>",
    "resources:printer"
  ],
  "conditions": {
    "remoteIP": {
        "type": "CIDRCondition",
        "options": {
            "cidr": "192.168.0.1/16"
        }
    }
  }
}
```

### Request

有了授权策略，就可以传入 reqeust，由 Ladon 来决定请求是否能通过授权。下面是一个请求示例：

```json
{
  "subject": "users:peter",
  "action" : "delete",
  "resource": "resources:articles:ladon-introduction",
  "context": {
    "remoteIP": "192.168.0.5"
  }
}
```

可以看到，在  remoteIP="192.168.0.5" 生效条件（Condition）下，针对主题（Subject） users:peter  对资源（Resource） resources:articles:ladon-introduction 的 delete  操作（Action），授权策略的效力（Effect）是 allow 的。所以 Ladon 会返回如下结果：

```json
{
    "allowed": true
}
```

### Condition

<img src="figures/image-20220506081304192.png" alt="image-20220506081304192" style="zoom:50%;" />

### AuditLogger

Ladon 还支持授权审计，用来记录授权历史。可以通过在 ladon.Ladon 中附加一个 ladon.AuditLogger 来实现：

```json
import "github.com/ory/ladon"
import manager "github.com/ory/ladon/manager/memory"

func main() {

    warden := ladon.Ladon{
        Manager: manager.NewMemoryManager(),
        AuditLogger: &ladon.AuditLoggerInfo{}
    }

    // ...
}
```

AuditLogger 会在授权时打印调用的策略到标准错误。AuditLogger 是一个 interface：

## 注意事项

在使用 Ladon 的过程中，有两个地方需要注意：所有检查都区分大小写，因为主题值可能是区分大小写的 ID。如果 ladon.Ladon 无法将策略与请求匹配，会默认授权结果为拒绝，并返回错误。

## Lab

```shell
go run example.go
# 新建终端，并在新建的终端继续
curl -s -X POST -H "Content-Type: application/json" \
-d '{"subject": "Tony", "action" : "delete", "resource": "resources:hair"}' \
http://127.0.0.1:8080/check
```

或者在同一个终端窗口中执行

```bash
go run example.go &
sleep 10

curl -s -X POST -H "Content-Type: application/json" \
-d '{"subject": "Tony", "action" : "delete", "resource": "resources:hair"}' \
http://127.0.0.1:8080/check
```