# MySQL

在用 Go 应用开发时，免不了要和 DB 打交道。每种语言都有优秀的 ORM 可供选择，Go 也不例外，如 gorm、xorm、gorose 等。

## GO SQL Driver

```bash
go run 10_go-sql-driver/example.go
```

## Gorm

Gorm 是 Go 语言的 ORM 包，功能强大、调用方便。Gorm 有很多特性，开发中常用的核心特性如下：

- 功能全：使用 ORM 操作数据库的所有接口 Gorm 都有，可满足开发中对数据库调用的各类需求。
- 开发者友好、调用方便：支持 Auto Migration，关联查询。
- 支持多种关系数据库：如 MySQL、Postgres、SQLite、SQLServer 等。
- 支持 Hook 钩子方法：这些 Hook 可以应用在 Create、Save、Update、Delete、Find 方法中。

### 基本操作

#### 模型定义

Gorm 使用模型（Model）来映射一个数据库表。默认情况下，使用 ID 作为主键，使用结构体名的 snake_cases 作为表名，使用字段名的 snake_case 作为列名，并使用 CreatedAt、UpdatedAt、DeletedAt 字段追踪创建、更新和删除时间。

使用 Gorm 的默认规则，可以减少代码量，但更好的方式是直接指明字段名和表名。例如，有以下模型：

```go
type Animal struct {
  AnimalID int64    // 列名 `animal_id`
  Birthday time.Time  // 列名 `birthday`
  Age   int64     // 列名 `age`
}
```

上述模型对应的表名为 animals，列名分别为 animal_id、birthday 和 age。可以通过以下方式来重命名表名和列名，并将 AnimalID 设置为表的主键：

```go
type Animal struct {
  AnimalID int64   `gorm:"column:animalID;primarykey"` // 将列名设为 `animalID`
  Birthday time.Time `gorm:"column:birthday"`       // 将列名设为 `birthday`
  Age   int64   `gorm:"column:age"`           // 将列名设为 `age`
}

func (a *Animal) TableName() string {
  return "animal"
}
```

##### Tag

上面的代码中，通过 primaryKey 标签指定主键，使用 column 标签指定列名，通过给 Models 添加 TableName 方法指定表名。如果没有指定表名，则 Gorm 使用结构体名的蛇形复数作为表名。

数据库表通常会包含 4 个字段：

- ID：自增字段，也作为主键。
- CreatedAt：记录创建时间。
- UpdatedAt：记录更新时间。
- DeletedAt：记录删除时间（软删除时有用）。

Gorm 也预定义了包含这 4 个字段的 Models，在定义自己的 Models 时，可以直接内嵌到结构体内。

#### DB 连接

在进行数据库的CURD操作之前，首先需要连接数据库。连接数据库时，可能需要设置一些参数，比如最大连接数、最大空闲连接数、最大连接时长等。

```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

如果需要 Gorm 正确地处理 time.Time 类型，在连接数据库时需要带上 parseTime 参数。如果要支持完整的 UTF-8 编码，可将 charset=utf8 更改为 charset=utf8mb4。

Gorm 支持连接池，底层是用 database/sql 包来维护连接池的，连接池设置如下：

```go
sqlDB, err := db.DB()
sqlDB.SetMaxIdleConns(100)              // 设置MySQL的最大空闲连接数（推荐100）
sqlDB.SetMaxOpenConns(100)             // 设置MySQL的最大连接数（推荐100）
sqlDB.SetConnMaxLifetime(time.Hour)  // 设置MySQL的空闲连接最大存活时间（推荐10s）
```

#### 创建/插入记录

可以通过 db.Create 方法来创建一条记录：

```go
type User struct {
  gorm.Model
  Name     string
  Age     uint8
  Birthday   *time.Time
}

user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
result := db.Create(&user) // 通过数据的指针来创建
```

db.Create 函数会返回如下 3 个值：

- user.ID：返回插入数据的主键，这个是直接赋值给 user 变量。
- result.Error：返回error。
- result.RowsAffected：返回插入记录的条数。

当需要插入的数据量比较大时，可以批量插入，以提高插入性能：

```go
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}

DB.Create(&users)

for _, user := range users {
  user.ID // 1,2,3
} 
```

####  删除记录

可以通过 Delete 方法删除记录：

```go
db.Where("name = ?", "jinzhu").Delete(&user)
```

Gorm 也支持根据主键进行删除，例如：

```go
db.Delete(&User{}, 10)
```

不过，还是推荐使用 db.Where 的方式进行删除，这种方式有两个优点：

- 删除方式更通用：使用 db.Where 不仅可以根据主键删除，还能够随意组合条件进行删除。
- 删除方式更显式：db.Where 显式的告诉你删除时的匹配条件，如果使用 db.Delete(&User{}, 10)，还需要确认 User 的主键，如果记错了主键，还可能会引入 Bug。

此外，Gorm 也支持批量删除：

```go
db.Where("name in (?)", []string{"jinzhu", "colin"}).Delete(&User{})
```

##### Scoped 软删除

Gorm 支持两种删除方法：软删除和永久删除

软删除是指执行 Delete 时，记录不会被从数据库中真正删除。Gorm 会将 DeletedAt 设置为当前时间，并且不能通过正常的方式查询到该记录。如果模型包含了一个 gorm.DeletedAt 字段，Gorm 在执行删除操作时，会软删除该记录。如：

```go
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
db.Where("age = ?", 20).Delete(&User{})

// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
db.Where("age = 20").Find(&user)
```

可以看到，Gorm 并没有真正把记录从数据库删除掉，而是只更新了 deleted_at 字段。在查询时，Gorm 查询条件中新增了 AND deleted_at IS NULL 条件，所以这些被设置过 deleted_at 字段的记录不会被查询到。对于一些比较重要的数据，可以通过软删除的方式删除记录，软删除可以使这些重要的数据后期能够被恢复，并且便于以后的排障。

可以通过下面的方式查找被软删除的记录：

```go
// SELECT * FROM users WHERE age = 20;
db.Unscoped().Where("age = 20").Find(&users)
```

##### Unscoped 永久删除

如果想永久删除一条记录，可以使用 Unscoped：

```go
// DELETE FROM orders WHERE id=10;
db.Unscoped().Delete(&order)
```

或也可以在模型中去掉 gorm.DeletedAt。

#### 更新记录

通过 Save 方法可以把 product 变量中所有跟数据库不一致的字段更新到数据库中。具体操作是：先获取某个资源的详细信息，再通过 product.Price = 200 这类赋值语句，对其中的一些字段重新赋值。最后，调用 Save 方法更新这些字段。Gorm 中，最常用的更新方法如下：

```go
db.First(&user)
user.Name = "jinzhu 2"
user.Age = 100
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
db.Save(&user)
```

上述方法会保留所有字段，所以执行 Save 时，需要先执行 First，获取某个记录的所有列的值，然后再对需要更新的字段设置值。还可以指定更新单个列：

```go
// UPDATE users SET age=200, updated_at='2013-11-17 21:34:10' WHERE name='colin';
db.Model(&User{}).Where("name = ?", "colin").Update("age", 200)
```

也可以指定更新多个列：

```go
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE name = 'colin';
db.Model(&user).Where("name", "colin").Updates(User{Name: "hello", Age: 18, Active: false})
```

这里要注意，这个方法只会更新非零值的字段。

### 查询数据

Gorm 支持不同的查询方法，包括检索单个记录、查询所有符合条件的记录和智能选择字段。

#### 检索单个记录

下面是检索单个记录的示例代码：

```go
// 获取第一条记录（主键升序）
// SELECT * FROM users ORDER BY id LIMIT 1;
db.First(&user)

// 获取最后一条记录（主键降序）
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
db.Last(&user)

result := db.First(&user)
result.RowsAffected // 返回找到的记录数
result.Error    // returns error

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

如果 model 类型没有定义主键，则按第一个字段排序。

#### 查询所有符合条件的记录

示例代码如下：

```go
users := make([]*User, 0)

// SELECT * FROM users WHERE name <> 'jinzhu';
db.Where("name <> ?", "jinzhu").Find(&users)
```

#### 智能选择字段

还可以通过 Select 方法，选择特定的字段。可以定义一个较小的结构体来接受选定的字段：

```go
type APIUser struct {
  ID  uint
  Name string
}

// SELECT `id`, `name` FROM `users` LIMIT 10;
db.Model(&User{}).Limit(10).Find(&APIUser{})
```

#### 高级查询

Gorm 支持很多高级查询功能：

- 指定查询记录时的排序方式。
- 查询时指定 Limit & Offset。
- 查询时指定Distinct。
- 查询时指定Count。

##### 指定查询记录时的排序方式

示例代码如下：

```go
// SELECT * FROM users ORDER BY age desc, name;
db.Order("age desc, name").Find(&users)
```

##### 查询时指定 Limit & Offset

Offset 指定从第几条记录开始查询，Limit 指定返回的最大记录数。Offset 和 Limit 值为 -1 时，消除 Offset 和 Limit 条件，示例代码如下：

```go
// SELECT * FROM users OFFSET 5 LIMIT 10;
db.Limit(10).Offset(5).Find(&users)
```

##### 查询时指定 Distinct

Distinct 可以从数据库记录中选择不同的值，示例代码如下：

```go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

##### 查询时指定 Count

 Count 可以获取匹配的条数，示例代码如下：

```go
var count int64

// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)
db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
```

### 原生SQL

Gorm 支持原生查询 SQL 和执行 SQL。原生查询SQL用法如下：

```go
type Result struct {
  ID  int
  Name string
  Age int
}

var result Result
db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

db.Exec("DROP TABLE users")

db.Exec("UPDATE orders SET shipped_at=? WHERE id IN ?", time.Now(), []int64{1,2,3})
```

### 表结构自动迁移

在一些小型项目中，还会用到 Gorm 的表结构自动迁移功能。

```go
db.AutoMigrate(&Product{})
```

不建议在正式的代码中自动迁移表结构。因为变更现网数据库是一个高危操作，现网数据库字段的添加、类型变更等，都需要经过严格的评估才能实施。这里将变更隐藏在代码中，在组件发布时很难被研发人员感知到，如果组件启动，就可能会自动修改现网表结构，也可能会因此引起重大的现网事故。

Gorm 的 AutoMigrate 方法只对新增的字段或索引进行变更，理论上是没有风险的。在实际的 Go 应用中，也有很多人使用 AutoMigrate 方法自动同步表结构。但更倾向于规范化、可感知的操作方式，所以在实际开发中，都是手动变更表结构的。当然，具体使用哪种方法，可以根据需要自行选择。

### Hook

Gorm 支持 hook 功能，例如下面这段代码在插入记录前执行 BeforeCreate 钩子：

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()
  if u.Name == "admin" {
    return errors.New("invalid name")
  }
  return
}
```

Gorm 支持的 Hook 如下表：

| **钩子**     | **触发时机**   |
| ------------ | -------------- |
| BeforeSave   | Save前执行     |
| AfterSave    | Save后执行     |
| BeforeCreate | 插入记录前执行 |
| AfterCreate  | 插入记录后执行 |
| BeforeDelete | 删除记录前执行 |
| AfterDelete  | 删除记录后执行 |
| BeforeUpdate | 更新记录前执行 |
| AfterUpdate  | 更新记录后执行 |
| AfterFind    | 查询记录后执行 |

### Lab

- Mysql DB 初始化

```shell
mysql -h 127.0.0.1 -u root -p < test.sql
```

- [Gorm](20_gorm/10_gorm.go)

```bash
go run 20_gorm/10_gorm.go -H 127.0.0.1:3306 -u root -p P@ssw0rd -d test
```

- [Gorm Model](20_gorm/12_gorm-model.go)

```bash
go run 20_gorm/12_gorm-model.go -H 127.0.0.1:3306 -u root -p P@ssw0rd -d test
```

## apiserver 示例

apiserver 示例后端需要持久化存储，所以采用 MySQL Gorm 作为其持久化存储引擎，具体介绍[在此](80_server/README.md)。

### Question

针对 [apiserver 示例](80_server/README.md)，Gorm 默认在 MySQL 中创建的 DB iam 中，创建的 Table 的名字都是复数，如 users。但如果查询 MySQL，会发现所创建的 Table 名称为单数，如 user。请解释原因，并且附上相关代码予以说明。

## Ref

1. [Go MySQL](https://zetcode.com/golang/mysql/)
