# Array & Slice
## Array
An array is a collection of variables of the same data type


## Slice
Slice is an abstraction over Array, it actually uses arrays as an underlying structure. 
To define a slice, you can declare it as an array without specifying its size.

The various operations over slice are:
- append(): add the elements to a slice. If the size of underlying array is not enough then automatically a new array is created and content of the old array is copied to it.
- len(): returns the number of elements presents in the slice.
- cap(): returns the capacity of the underlying array of the slice. 
- copy(): the contents of a source slice are copied to a destination slice.
- <SliceName>[start:end]: returns a slice object containing the elements of base slice from index start to end- 1.

```go
var s []int
s := []int{1,2,3,4,5,6,7,8,9,10}
a := make([]int, 10)
b := make([]int, 0, 10)
c := s[0:4]
```
