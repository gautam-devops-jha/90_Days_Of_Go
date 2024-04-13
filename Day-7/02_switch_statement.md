# Switch Statements

A switch statement can be considered as a specialized form of an if statement. 

It looks something like this :-

```go
switch test expression {
    case expression1:
        statements
    case expression2, expression3:
        statements
    default:
        statements
}
```

- We start with switch keyword instead of if keyword, and then we are going to have a test expression. 
- similar to if statement we had the if keyword followed by a boolean expression. well in switch statement we only going to have a single value, then fo is going to compare that value to a series of expressions that we provide in what are called the switch's cases. 
- above we see that we got the case keyword followed by another expression. when go execute this test, it's going to enter the switch, evaluates the test expression to derive a single value, and then it's going to look at the cases to see if any of those cases are equivalent to the test expression or to the value that was derived from the test expression.
- if it does then the block of statement will be executed. 
- we can also combine multiple expression into a single case, where we simply  separate those expressions using commas. and then if any of those expressions evaluate to the same value as the test expression, then once again we are going to esecute the block statements that comes after the case.
- Finally, if we want the switch to always evallluate at least one branch, then, just like the if statement has the else statement, switches have what are called the default cases. The default case will execute if non of the other cases have been selected. 


```go
i := 5

switch i {
    case 1: 
        fmt.Println("First case")
    case 2 + 3, 2*i+3: // the fist statement here match so this case will be executed. 
        fmt.Println("Second Case")
    default:
        fmt.Println("default case")
}


// we can use initilizer in switch statement also

switch i:=999; i {
    case 1:
        fmt.Println("First Case")
    case 2 + 3, 2*i+3:
        fmt.Println("Seconnd Case")
    default: // default will execute as non of the above cases results to true
        fmt.Println("default case")
}
```


# Logical Switch 

Switch statement has a specil form in go called logical switch 

- A logical switch is relly nothing new. Based on what we've already seen with switches, we could make a logical switch by doing something like this 
`switch i := 8  true { }` 
- if for the test expression we just provided a Boolean value of true, then every case is going to evaluate to see if it is true or not. 
- the different is our cases start to look very different. instead of having some sort of a value or airthmetic expression, our cases are going to contain boolean expressions. 

```go
switch i := 8; true {
    case i < 5:
        fmt.Println("i is less than 5")
    case i < 10:
        fmt.Println("i is less than 10")
    default:
        fmt.println("i is greater than 10")

}
```