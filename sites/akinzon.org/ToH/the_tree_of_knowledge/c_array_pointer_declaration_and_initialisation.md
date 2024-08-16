# C Array Pointer Declaration and Initialisation

Up: [C Pointer Declaration and Initialisation](c_pointer_declaration_and_initialisation)
Brother(s):
TARGET DECK

Array pointers are declared in the same way as regular pointers.

However, their initialisation is different:

```C
a = array;
```

and **NOT**

```c
a = &array;
```

since the compiler knows that the array is already a memory address.

But array elements are still regular variables, so they would be initialised using `&`.






























#### Why:
#### How:









