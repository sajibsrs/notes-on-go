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

### 1.3 # Array length
We can determine what is the length of an array with:

```go
len(array)
```

## 2 # Code

```go
package main

import "fmt"

func main() {
	arrOne := [...]int{1, 2, 3}
	var arrTwo [3]string

	arrTwo[0] = "a"
	arrTwo[1] = "b"
	arrTwo[2] = "c"

	fmt.Println("arrOne:", arrOne)
	fmt.Println("arrTwo:", arrTwo)
}
```