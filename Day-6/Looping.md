# Looping

- Go has only for loop, but we have different forms of for loop, depending on what we're trying to do. 


1. **Infinite Loop** :- runs again and again without breaking out of the loop. 
    * `for {...}` 
2. **Loop till condition** :- Run the loop again and again till a condition match 
    * `for condition { ... }` 
3. **Counter Based Loop** :- This is the most complicated  form of the loop in that it allows us three different statements after the for keyword and before that opening curly brace.
    * `for initializer; test; post clause { ... }`  
        - We have initilizer statement that allows us to set up the loop for whatever we need to do, 
        - We have the test statement, which is the same as the condition in the conditional loop
        - Then we have post clause 

## Using the loops

```go
// Infinite loop
i := 1 

for { // infinite loop
    fmt.Println(i)
    i += 1 // this is same as i = i + 1
}


// Loop till Condition

i := 1
for i < 3 {
    fmt.Println(i)
    i += 1
}
fmt.Println("Done!")


// counter Based Loop
for i :=1; i < 3; i++ {
    fmt.Println(i)
}
fmt.Println("Done!")




```

## Looping with Collections

When we are looping through collections, we have the signal of termination builtin, because we are going to iterate through the loop once for every member of the collection.s

We have 3 forms for looping tn collection

1. `for key, value := range collection { ... }` We are going to use this for keyword again, because every loop in Go is a for loop, and then we are going to retrieve up to two values, the key and the value back out. So for the map, the key is going to be the key of map, and for a slice or an array, the key is going to be the index. 
 - Range keyyword tells the go that we are looping through collection. 

2. `for key := range collection { ... }` :- To just get the keys

3. `for _, value := range collection { ... }` :- To just get the values 


```go

// Looping with Collections
 arr := [3]int{101, 102, 103}

 for i, v := range arr {
    fmt.Println(i, v)
 }



```