# Variables
variables provide us a mechanism to allow our programs to remember specific pieces of information for us to use later.

- There are three different ways that we can declare variables in Go

    1. ``var myName string`` :- we start with var keyword, we are going to provide the name for our variable, in this case the variable name is **myName** and then we're going to follow that with the type. 
         > All variable in Go are strongly typed so we have to provide the type when declaring the variable in one way or another.  

    2. If we already know the value of varible we are initilizing to we can combine that initialization with declaration, `var myName string = "Gautam"`

    2. `var myName = "Gautam"` Go compiler already know what datatype to inferr so we can use inffered type variable. When we declare a variable and assign a string to it, Go knows well that has to be a string
        - In Go the above format is fairly rarely used, becaus ethis situation is actually so common, we have a shorthand notation, and that's called the short declaration syntax, and that looks like this `myName := "Gautam"` 

