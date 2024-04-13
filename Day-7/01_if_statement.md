# if Statement

if statement is one of the most common branching constructs that we are going to run into no matter what programming languge we use.

There are multiple forms of if statement that allow us model different scenarios that our programs runs into.

1. `if test { ... }` :- The most basic is using the if keyword followed by some sort of a test. this test is going to resolve in boolean value, so we are going to end up with a true or a false here. 
    - if the test evaluated to true then block of statement in curly braces will execute.
    - if it evaluates to false, then they will not execute.

2. `if test { ... } else { ... }` :- here when test is evaluated,
    - if it evaluatess to true,  then the first block of statements will execute 
    - if it evaluates to false, then the else block will execute

3. `if test { ... } else if test { ... }` here also we have if statement which tests and evaluates the test
    - If test evaluates to true, the statement under curly brace will execute, if test evaluates to false, it jums to next else if statement
    - If the second else if test evaluates to tru the statement under curly brace will execute, otherwise it's not going to exeecute any statement at all.

We can have multiple else if clause, but atlease 1 if clause and 0 or 1 else clause. 


4. `if initializer; test { ... }` in go we have something called initilizer, in counter based for loops, the first statemnet was initilizer. then we had the test and then we had the post clause.  
    - if statements in go also allows an optional initilizer.


```go
i := 5
if i < 5 { 
    fmt.Println("i is less than 5") // if is not less than 5 so this statement will not execute.
} else {
    fmt.Println("i is at least 5") // else statement  will execute because if statement evaluates to false
}
fmt.Println("after the if statement") // this statement will execute because it is out of the if statement



// else if 

i := 5

if i < 5 {
    if i < 5 {
        fmt.Println("i is less that 5")
    } else if i < 10 {
        fmt.Println("i is less than 10")
    } else {
        fmt.Println("i is at least 10")
    }
    fmt.Println("after the if statement")
}


// initilizer

if i := 5; i < 5 {
    if i < 5 {
        fmt.Println("i is less that 5")
    } else if i < 10 {
        fmt.Println("i is less than 10")
    } else {
        fmt.Println("i is at least 10")
    }
    fmt.Println("after the if statement")
}

```