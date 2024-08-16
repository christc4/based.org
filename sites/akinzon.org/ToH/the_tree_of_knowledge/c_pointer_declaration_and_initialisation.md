# C Pointer Declaration and Initialisation

Up: [C Pointers](c_pointers)
Brother(s):
TARGET DECK


This is how pointers are declared in C:

```c
datatype *pointerVariable;
```

the asterisk tells the compiler that this is a pointer.

Pointers are initialised like this:

```C
pointer = &variable;
```

We do **NOT** do this:

```c
*pointer = &variable;
```

since using the asterisk on a pointer returns the value stored at that memory location, and not the memory location itself.

[C Array Pointer Declaration and Initialisation](c_array_pointer_declaration_and_initialisation)


























#### Why:
#### How:









