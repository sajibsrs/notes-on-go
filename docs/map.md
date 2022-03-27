# # Map
Map in Go is similar to other language call **HashMap**, **Dictionary**, **HashTable**, **Associative array** etc.


## 1 # Declaring map

### 1.1 # With map syntax
Syntax `map[keyType]valueType`:

```go
mp := map[string]string
```
In above example **mp** is `nil`.

### 1.2 # With make syntax
Syntax `make(map[keyType]valueType)`:

```go
mp := make(map[string]string)
```
In above example **mp** is not `nil`.

## 2 # Initialization

### 2.1 # Method one

```go
mp := map[string]string{}

mp["hello"] = "world"
mp["good"]  = "morning"
```

*Note: If you assign more than one element with the same key, only the **last one** will be accepted.*

### 2.2 # Method two

```go
mp := map[string]string{"hello": "world", "good": "morning"}
```

### 2.3 # Method three

```go
mp := make(map[string]int)

mp["juice"] = 1
mp["eggs"] = 12
```

## 3 # Accessing map elements

```go
mp := make(map[string]int)

mp["juice"] = 1
mp["eggs"] = 12

fmt.Println("Eggs count: ", mp["eggs"]);
```

## 4 # Key existence
When receiving a value by it's key from a map, we can check if that key even exists. In the example below the **first** return value holds the value and the **second** one (optional) holds a **boolean** value. If key exists it returns true or false otherwise.

```go
mp := map[string]string{"nice": "weather", "lovely": "environment"}

val, st := mp["nice"]

fmt.Println("val and status: ", val, st)
```

If we are not going to use the second return value, we can ignore it with `_` (underscore).


```go
mp := map[string]string{"nice": "weather", "lovely": "environment"}

val, _ := mp["nice"]

fmt.Println("val: ", val)
```

or simply don't use the second variable at all.