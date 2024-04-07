# Conatants

Constants and variable looks same the difference is that the value cannot be changed

```go 
const a = 42 // constant (implicitly typed)
const b string = "hello, world" // explicitly typed constant
```

The difference between these contant will show how it can be used.
- Implicitly typed constant go treats the value as it were a literal so whereever we use constant a it's going to just basically copy in the value 42. so we could use a to assign to an integer, we could assign it to a floating point number, we can assignn it to any numeric type that we want, because go is going to treat that as literal and we could assign the literal value 42 to a floating point number. 
- With explicitly typed constant go iis going to add an additional control So in this second example, b is only going to be used where a string data type would be allowed. 

```go
const c = a // one constant can be assigned to another


// We can declare multiple constants ar once 

const (
    d = true
    e = 3.14
)
```

```go

// unassigned constant receive previous value if in a group
const (
    a = "foo"
    b 
)
```

- We can use expression on constant `const c = 2 * 5` and `const d = "hello, "+"world"`

- We Can't assign any function to constant `const e = someFunction()` WILL NOT WORK

# iota
Iota is a special symbol that's used only in the context of assigning the constants.
- iota is related to position in constant group

```go
const a = iota  // 0


const (
    b = iota // 0  ----> iota starts at zero on first line
    c        // 1 --> constant expression copied, iota increments
    d = 3 * iota // ---> iota increments again ---> 3 * 2  ---> // 6 
)

// if you go inti another constant block iota reset to zero
const (
    e = iota
)

```

