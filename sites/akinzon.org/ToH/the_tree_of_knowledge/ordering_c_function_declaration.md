# Ordering C Function Declaration

Up: [C Functions](c_functions)
Brother(s):
TARGET DECK

In C, you cannot use a function before it has been declared in the program.

C compilers compile function calls even before the function has been defined.

This can lead to errors.

This problem is solved in two ways.


**Defining functions in order:**

```C
int funA() {

}

int funB() {
	funA();
}
```


**Declaring functions before they're defined:**

```C
int funA();


int funB() {
	funA();
}

int funA() {

}
```

































#### Why:
#### How:









