# Class
Go is an object-oriented language. 
However, it does not have any class keyword. 
We can associate functions directly by structures and other data types. 


## Structure
structures are a collection of multiple data types as a single entity.
```go
type student struct { 
    rollNo int
    name string 
}

stud := student{1, "Johnny"}
```

## Method
Associate methods to type


## Interface
Interfaces are defined as a set of methods. 

```go
type Rect struct { 
    width float64
    height float64 
}

type Shape interface { 
    Area() float64
    Perimeter() float64 
}
```

