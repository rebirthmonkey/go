# Map
哈希表是除了数组之外，最常见的数据结构，几乎所有的语言都会有数组和哈希表这两种集合元素，有的语言将数组实现成列表，有的语言将哈希表称作结构体或者字典，但是它们是两种设计集合元素的思路，数组用于表示元素的序列，而哈希表示的是键值对之间映射关系，只是不同语言的叫法和实现稍微有些不同。

A map is a collection of Key-Value pairs. Hash-Table is used to store elements in a Map so it is unordered.

## 声明/创建

```go
hash := make(map[string]int, 3)
```



## 初始化

```go
hash := map[string]int{
	"1": 2,
	"3": 4,
	"5": 6,
}
```

## 访问/赋值

```go
hash["1"] = 2
hash["3"] = 4
hash["5"] = 6
```

