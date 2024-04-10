# Structs

- Maps, arrays and slices are homogenious data type, all the values are of same data type
- A struct is a fixed-size aggregate type, so it's similar to an array in that aspect. But it's values can be different types. 
- Struct is a hetrogeneous data structure, and it's the only heterogeneous aggregate type that we have available in Go. 
- we can contain anything that we want within the struct. the only constraint is we have to define what what aare the the things which struct will contain as we are writing our program, we can't change that at runtime.
- Structs are comparable



```go
var s struct {  // declare an anonymous struct
    name string
    id int 
}

fmt.Println(s) // {"" 0}

s.name = "Gautam" // assign value to field 
fmt.Println(s.name) // Gautam


// The above struct is anonymous struct to use these struct we have to define it again and again


// To solve this issue we can create a custom type


type myStruct struct{
    name string
    id int
}

var s myStruct // declare variable with custom type
fmt.Println(s) // {"" 0}

s = myStruct{
    name: "Gautam"
    id: 42 
}

fmt.Println(s) // {"Gautam" 42}

// if you copy the struct and update the value in one of the variable the values is not going to refelect on another 
s2 := s
s.name = "Tricia" 
fmt.Println(s, s2) // {"Tricia" 42} {"Arthur" 42}

s == s2 // false 


```