# Strting with our First code

open VS code inside a folder where you want you project should be in. 

**Now in order to start building an application with Go, we need to create a gro project, or in Go terms *module***

## Module 
- Module is the basic unit of a Go project
- Module is unit of sharing also, As we import dependencies, those libraries that we are going to be importing are also modules. 

### How to create a module in go?
- Step 1 :- Open terminal using CTRL+SHIFT+`  
- Step 2 :- Most probably you already in the directory where you want your codes should be at or else you can always change the directory using commmand `cd <your directory path>` for example, I want to mve to the directory called code so `cd code`

- Step 3 :- Initilize module in directory using command `go mod init <module name>` for example, I want to create a module with name day2 as this is my project name, so `go mod init day2`. 
    * It create a .mod file in the directory which contains go version you are using and module name. 
    * Now we are ready to create a go program

## Create a Program 
For creating a program in go we need a source file with extension `.go` 

### Creating Our First program
- Step 1 :- Create a file `main.go` in the same directory where you initilized module 
- Step 2:- Write a Program
```go
// In Go we have to start our program with a Package identifier
package main // here main indicates that we are creating a source file for executable command, or something taht we are going to be able to compile and run later.

//  Within the package main, in order for go to inderstand where to start the application it has to be provided what's called an entry point.
// function called main in package main is the entrypoint for go program

func main() {
	//  the code that needs to be run should be placed within these curly braces
	//  Let's pring Hello Gophers!
	println("Hello Gophers!") // println is the builtin function of go which prints someting out to terminal with a carriage return. that's what the ln means after print ln means new line.

	// Go programmers are gophers

}

```

- Step 3 :- Run the program We have two choice to run the program.
    1. `go build` : Build and Run the program (Build program compile the package and create a executable binary in your currect directory)
    2. `go run` :- compiles the program in a temprory directory and run it and show the output on terminal  
    In development mode we use go run so use command `go run <directory where code is present>` for example my code is in code directory So i run `go run .`

> Remeber You have to run the go run command where the `go.mod` file is present in my case it is in code directory.


## Importing a package 
to import a package in prigram we use import statement with package name. For ex :- `import "fmt"`

Plese see main.go file in code dirctory for more clarification.

