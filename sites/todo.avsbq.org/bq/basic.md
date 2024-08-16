# Basic explanation of golang stuff

```
if len(os.Args) < 2 {
   fmt.Println("Please provide a filename")
   os.Exit(1)
}
```

This checks if the number of command line arguments passed is less than two, when the program is executed.

os.Args is a slice of strings representing the arguments provided to the program,
os.Args[0] being the program# name itself, obviously the program name itself is considered an argument, and is stored at index 0 in os.Args. If there are fewer than two arguments then the editor will exit

> Of course, you don't want your text editor to do this. You want to provide a name for a new file, upon launching the editor. This is for basic text editors.

```
filename := os.Args[1]

> What is := ?
> This is used to declare and initialize the variable filename with the value of os.Args[1], remember the program name itself is stored at index 0, dont get confused.
It allows you to decale and initialize a variable without explicitly specifyiinng its type, and Go infers the type based on the value ofo the right-hand side of the expression. This syntax is called a short variable declaration.

If you were to uses = alone, it would indicate assignment without declaration, this only works if you do something like this

var filename string
filename =  os.Args[1]

os.O_APPEND means that writes to file should append to end of file rather than overwrite

os.O_WRONLY means that file should be opened for write-only access

0644 this is file mode or permissions ofr file being opened, in octal notation means read/write for owner, read only for group and others


