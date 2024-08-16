# Importance of Declaring Pointers with the Right Type

Up: [C Pointers](c_pointers)
Brother(s):
TARGET DECK

It is important to declare pointers with the right value.

Why?

When it comes to [Pointer Incrementing in C|incrementing pointers](pointer_incrementing_in_c|incrementing_pointers), they are incremented by the size of the data type they were declared with.

This works because different data types take up different amounts of memory (which can be checked using [Checking The Size of Variables in C|the size of operator](checking_the_size_of_variables_in_c|the_size_of_operator)), so if you increment by the size of long long integers on an integer array, you could end up with unwanted side effects.

As a matter of fact, the compiler will warn you:

```c
test.c:11:13: warning: incompatible pointer types assigning to ’double *’ from ’int [10]’ [-Wincompatible-pointer-types] pointer = array; ˆ ˜˜˜˜˜ 1 warning generated.
```

































#### Why:
#### How:









