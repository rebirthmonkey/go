# Loop

## Condition
### IF
```go
if <Boolean expression> { 
  <Statements>
}
```

### IF-ELSE
```go
if <Boolean expression> { 
  <Statements>
} else { 
  <Statements>
}
```

## Switch
```go
switch {
  case <condition>: 
    <statements>
  case <condition>:
    <statements>
  default:
    <statements>
}
```


## Loop
The Go for loop has four forms:
1. for<initialization>;<condition>;<increment/decrement>{} 
2. for<condition>{}-likeawhileloop
3. for{}-aninfinitewhileloop.
4. forwithrange.

### Range
The range keyword is used in for loop to iterate data in data structures (arrays, slices, string, maps etc.). 

```go
func main() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}
```



