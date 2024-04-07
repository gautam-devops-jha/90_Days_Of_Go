# Simple data types in Go
here simple datatype means that I am talking about the data types that contain one and only one value. 

- In Go there are only 4 simple data types
    1. Strings
    2. Numbers
    3. Booleans
    4. Errors
- all of these types have one thing in common and that is they can contain one, and only one value, and that's preety much all that they have in common. 


## Strings

- Strings in Go represent a collection of one or more UTF-8 code points. in other words, they contain one or more letters or numbers or other symbols that we might use depending on the languages that your users are using. 
- Strings comes in two flavour in Go 
    1. Quoted string :- The string that are series of characters, delimited by quotation marks (""), These represents what's called an interpreted string in Go. 
        * ex:- "This is a string"
    2. Raw String:- this type of string don't use quotation marks, it uses what's called a backtick or grave mark (``).
        * ex:- \`This is also a string\` 

- See this string carefully :- "this is a escape character: \n it creates a newline"
    * notice the string `\n` is an escape sequence then We have a text called it creates newline. 
    * `\n` is a special command that's being sent to Go for interpretation. With an interpreted string like we have here, (notice we are using quotation mark.) this is going to create a newline character. 
    * So if we were, for example, to print this out to our standard output, we would get this result. 
    ```go
    this is an escape character:
        it creates newline
    ```
    - If we have the same string with the grave marks (raw string) the escape character is not interpreted. so if we print this out it will going to print exactly like we saww it in code.
    ```go
    this is an escape character: \n it creates a newline
    ```

- One of the neat thing that we can do with raw strings is if you have a whitespace, so notie here in this  raw strings we actuallly got a carriage return.  
   > \`raw strings  
    ignore new lines\`  


    * This is valid go code, Go is not going to have any problem with the fact that we've got a carriage return right in the middle of that statement. and what is go going to do? Well, it's not going to interpret the whitespace. when we print it we will get it reproduced exactly like we see it. 
    * We run into this often when we see datat structure that we are creating. For example, using XML or JSON. 

## Numbers

- Numbers in go comes in four basic categories. 
    1. Integers :- 99, 0, -937
        * The most common type that we see in Go is `int` datatype, There are other integer data types also. we will see later
        
    2. Unsigned integers :- 0, 15, 7329
        * Unsigned integers are kind of like integers in that they're whole numbers, but instead of being able to go negative unsigned integers have a minimum value of 0
        * The most common datatype you will see in unsigned integer is `uint`. 
    3. Floating point numbers :- 6.02e23, 3.1415, 0.25
        * The most common floating point number we see in GO is `float32` which represent 32-bit floating-point numbers, and `float64` which represent 64-bit floating-point number. 
    4. Complex Numbers :- 1+2i, 0.833i, 6.02e23+3.1415i
        *  complex numbers are supported as first class citizen in GO, and they represent a number that has a real component, on that first example you see the real component is number 1, followed by an imaginary component, which if you don't work with complex numbers, isn't super important. 
        * The imaginary number is simply multiplying that number by the square root 0f -1, which is useful in many fields of math and statistics. 
        * Similar to floating point numbers they are two types `complex64` which is made up of two 32-bit floating point numbers, a real 32-bit floating point and an imagionary 32-bit floating point number. And `complex128` is made up of two float64s 

## Booleans

- Booleans are simplest datatyppe in programming and probably the one that you're going to run into the most often.
- A Boolean value simply represent `true` or `false` and in Go those are first class citizens.
- So we don't have true and false as aliases of other numbers like other languages do. **True is true and False is false**


## Error type 
The error type built-in interface type is the convential interface for representing an error condition, with the nil value representing no error. 