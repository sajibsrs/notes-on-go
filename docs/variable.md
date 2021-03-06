# # Variable

## 1 # Declaration and assignment
There are couple of ways of declaring the variable in **Go**, depending on their **scope** and **type**. For now we'll see declaring simple **integer** variable. Variables with different data types will be discussed in their own section.

### 1.1 # Just the declaration:

```go
var value int
```
*Note: In Go all variable is initialized even if you do not assign a value to it. All variable gets initialized to its **zero** value. Eg. Integer zero value is zero, String zero value is empty string and so forth.*

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
*Note: Method two is only available **inside a function**  scope.*

*Note: It is not possible to use `:=` operator without assigning a value to it.*
