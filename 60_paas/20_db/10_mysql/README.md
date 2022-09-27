# MySQL

## GO SQL Driver

```shell
go run example.go
```

## Gorm

Gorm 是 Go 语言的 ORM 包，功能强大，调用方便。Gorm 有很多特性，开发中常用的核心特性如下：

- 功能全：使用 ORM 操作数据库的接口，Gorm 都有，可满足开发中对数据库调用的各类需求。
- 支持 Hook 钩子方法：这些 Hook 可以应用在 Create、Save、Update、Delete、Find 方法中。
- 开发者友好、调用方便：支持 Auto Migration，关联查询。
- 支持多种关系数据库：如 MySQL、Postgres、SQLite、SQLServer 等。

### 方法

- Open()：连接数据库

- Create()：insert 记录

- Delete()/DeleteAt()：

- Save()：更新记录

- Where().Find()：查询记录

- Order().Find()：

- Limit(2).Offset(5).Find()

- Distinct('name', 'age').Find()

### 配置

#### Tag

- column

#### Scoped/Unscoped

Gorm 支持两种删除方法：软删除和永久删除：

- 软删除（scoped 默认）：软删除是指执行 Delete 时，记录不会被从数据库中真正删除。Gorm 会将 DeletedAt 设置为当前时间，并且不能通过正常的方式查询到该记录。如果模型包含了一个 gorm.DeletedAt 字段，Gorm 在执行删除操作时，会更新该记录为 NULL。Gorm 查询条件中会默认新增 AND deleted_at IS NULL条件，所以这些被设置过 deleted_at 字段的记录（已被软删除）不会被查询到。对于一些比较重要的数据，可以通过软删除的方式删除记录，软删除可以使这些重要的数据后期能够被恢复，并且便于以后的排障。
- 永久删除（unscoped）：如果想永久删除一条记录，可以使用 Unscoped，或 gorm.DeletedAt。

#### TableName()

给 Models 添加 TableName 方法，来告诉 Gorm 该 Models 映射到数据库中的哪张表。

如果没有指定表名，则 Gorm 使用结构体名的蛇形复数作为表名。例如：结构体名为 DockerInstance ，则表名为 dockerInstances 。

### Hook

- BeforeCreate
- AfterCreate
- BeforeUpdate
- AfterUpdate
- BeforeSave
- AfterSave
- BeforeDelete
- AfterDelete
- AfterFind

### Lab

- Mysql DB 初始化

```shell
mysql -h 127.0.0.1 -u root -p < test.sql
```

- [Gorm](20_gorm/10_gorm.go)

```shell
go run 10_gorm.go -H 127.0.0.1:3306 -u root -p P@ssw0rd -d test
```

- [Gorm Model](20_gorm/12_gorm-model.go)

```shell
go run 12_gorm-model.go -H 127.0.0.1:3306 -u root -p P@ssw0rd -d test
```

## apiserver 示例

apiserver 示例后端需要持久化存储，所以采用 MySQL Gorm 作为其持久化存储引擎，具体介绍[在此](80_server/README.md)。

## Ref

1. [Go MySQL](https://zetcode.com/golang/mysql/)
