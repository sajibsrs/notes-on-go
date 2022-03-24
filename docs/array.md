# # Go Array
Array has a fixed length in Go.

## 1 # Declaring array
In order to declare an array we need two things.
* It's length
* It's element type

*Note: Unlike other dynamically typed languages, Go does not allow type mixing in array. Every element must be of same type.*

### 1.1 # Length defined

```go
array := [3]int{1, 2, 3}
```

### 1.2 # Length inferred

```go
array := [...]int{1, 2, 3}
```
*Note: Inferred length syntax is only available if you assign a value along with the declaration.*

## 2 # Code

```go
package main

import "fmt"

func main() {
	arr_one := [...]int{1, 2, 3}
	var arr_two [3]string

	arr_two[0] = "a"
	arr_two[1] = "b"
	arr_two[2] = "c"

	fmt.Println("arr_one:", arr_one)
	fmt.Println("arr_two:", arr_two)
}
```