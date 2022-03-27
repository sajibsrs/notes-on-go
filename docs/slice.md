# # Slice
Slice is very similar to go, but more flexible. Unlike array, slice has a dynamic length. Meaning it can grow on demand. Behind the scene slice is created from an underlying array.

Slice has two main property:
1. Length (the number of elements in the slice)
2. Capacity (the number of elements the slice can grow or shrink to)

This two property can be measured with:

```go
length := len(slc)
```

```go
capacity := cap(slc)
```

## 1 # Declaring slice
Slice can be declared ahead of initialization. Code snippet below will create a empty slice with length 0 and capacity 0.

```go
slc := []int{}
```

## 2 # Declaration and initialization

```go
slc := []int{1, 2, 3}
```

##  3 # Slice from an array
A slice can be created from an array.

```go
arr := [5]{1, 2, 3, 4, 5}
slc := arr[2:4]
```

## 4 # Slice with make
When using `make()`, we can define it's second argument as length and third argument as capacity. Capacity is optional.

```go
slc := make([]int, 3)
slc1 := make([]string, 3, 6)
```

*Note: If capacity is not defined, slice capacity will be equal to its length.*

## 5 # Adding elements
We can add new items to slice with `append`. If we create slice with `make` we can add items to it as we do with array and with `append` too. As with `make` function it has a length. But if we create it without length we can only add elements to it with `append` function.

```go
slc := make([]int, 3)
slc[0] = 1
slc[1] = 2
slc[2] = 3

slc = append(slc, 4)
```

*Note: If you try to add more elements than it's length with array method, it will result in an error. But you can use `append` as much as you want.*

```go
slc := []int{}
slc = append(slc, 4)
```

*Note: If we exceed slice length by adding new items, slice capacity automatically increases to accommodate new items. If slice capacity is 3 and we add 4 items it's capacity will increase to 6.*

*Note: The formula of increment is, if capacity is n then next capacity will be 2n then 3n then 4n and so on. Are we wasting some memory if we are not using all capacity? Hell ya, but we have lots of memory now. Right? (pun intended).*