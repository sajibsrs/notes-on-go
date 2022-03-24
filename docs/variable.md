# Variable

## 1 # Declaration and assignment
There are couple of ways of declaring the variable in **Go**, depending on their **scope** and **type**. For now we'll see declaring simple **integer** variable. Variables with different data types will be discussed in their own section.

### 1.1 # Just the declaration:

```go
var value int;
```

### 1.2 # Declaration and assignment:
When declaration and assignment happens at once, data type can be **omitted**. In that case Go **infers** the data type from the value.

Method one:

```go
var value = 10
```

Method two:

```go
value := 10
```
*Note: Method two is only available **inside a function** scope.*
