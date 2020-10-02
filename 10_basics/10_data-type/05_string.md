# String

- A String is a sequence of Unicode character. 
- String is an immutable type variable. 
- Double quotes are used to declare strings.

```go
var s string = "Hello, World!" // standard declaration
```

```go
s := "Hello, World!" // short declaration
```

## Operation

- len(mystring) --> 12: Used to find the number of characters in mystring
- “hello”+“world” --> “helloworld”: 2 strings are concatenated into a single string
- “world” == “hello” --> False: Equality can be tested using “==” sign
- “a” < “b” --> True: Unicode value can also be used to “b” < “a” False compare
- mystring[0] --> “h” Indexing: String are indexed same as array.
- mystring[1:4] --> "ell": Slicing