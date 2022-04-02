# # Struct (Structure)
Structs is used to create a collection of members of different data types.
Arrays are used to store data of same type under a single variable, structs are used to store multiple values of different types under a single variable. A struct can be useful for grouping data together to create records.

## 1 # Declaration

```go
type Person struct {
    name string
    age int
    location string
}
```

## 2 # Access struct members
To access struct member, use `(.)` dot operator.

```go
type Person struct {
    name string
    age int
    location string
}

func main() {
    var p Person

    p.name = "Xavier"
    p.age = 34
    p.location = "NY"

    fmt.Println("Name: ", p.name)
}
```
