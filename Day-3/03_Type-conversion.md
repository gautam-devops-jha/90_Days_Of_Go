# Type conversion 

One of the Go's basic principles, is to be clear over being clever. 

So when we come to converting from one data type to another data type, We see this very very strongly 

Let's just say that we've got two variables, 

```go
var i int = 32
var f float32
```

Now in many languages we would be able to do this.   
`f = i`  
> In Go this is actually an error Go does not support implicit conversion because there is potential for data loss and subtle bugs to be introduced.

We can instead do a explicit conversion like this   
`f = float32(i)`
