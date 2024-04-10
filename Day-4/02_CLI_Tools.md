# CLI Tools in Standard Library 

There are two locations, well actually we are gonna be using 3 withing Standard Library that are very useful to be aware of when you are building CLI applications. 

1. **OS Package** :- The OS package contains a lot of functionality that allows our programs, to interact with the operation system that the program is running on. So if you are runnig on a MAC the OS package is going to allow you to interact with the MAC operating system.   
    -   Three members of OS package that are going to be particularly useful to use are
        * Stdin object  :- Allows us to read input from the standard input source, which by default is going to be our keyboard. 
        * Stdout object :- Used to write to standard output device of a computer, which is typically a monitor.
        * stderr object :- similar to standard output, but it's a special channel that's set aside to allow us to write error information out, in case we need to redirect that information for a logging system or something like that. 

2. **fmt Package** :- The fmt package contains two sections of functionality, But it's all about string management.
    - Scan :- The scan functions allow us to read information in from various sources, including for example, stdin. 
    - Print :- Prints something on Standard output 

3. **bufio** :- The OS package gives us the raw access to keyboard data. so we are going to be pullling in one keystroke at a time. However that can be a little bit cumbersome when you're writing CLI applications.
    - The bufio package stands for buffered input/output, and it allows us to gather groups of text together into more logical or more useful units. 
    - basically buffio reads one line at a time and fmt reads one character at a time. 


    