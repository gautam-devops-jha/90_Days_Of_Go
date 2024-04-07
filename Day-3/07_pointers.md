# Pointers and value 

Pointers share memory in Go so if we change the value using pointers it will change the actual value presnt in the variable.

```go
a := 42 
b := &a // address operator (&) returns the address of a variable 

*b //42 
a = 27 

*b // 27 

c = new(int) // built-in "new" function creates pointer to anonymous variable
```

Pointers are primarily used to share memory (but use copies whenever possible).
   **why ?**
What we get into is Go is often tapped to be useed in highly concurrent scenarios such as creating web services. Any time we are sharing memory in a highly concurrent program, we run into th possibility fo creating something called a data race.
- A data race is where we have got multiple tasks in our program that are competing for the same memory and behaviour becomes undefined, In another words we create bugs. If you use copy you never share memory and so therefore by avoiding pointers whenever we can we make it less likely to create data races in our applications.

