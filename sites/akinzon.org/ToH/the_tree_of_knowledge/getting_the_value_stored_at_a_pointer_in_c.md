# Getting The Value Stored at a Pointer in C

Up: [C Pointers](c_pointers)
Brother(s):
TARGET DECK

To get the value stored at a pointer in C, the **dereference operator** is used.

for a pointer `ptr`, to get the value you would use `*ptr`.

Why not just use the original value instead of this bs?

This is because it is important for [Passing Values by Reference|passing by reference](passing_values_by_reference|passing_by_reference).

You can change this value too, in a very natural way:

```C
*ptr = val;
```

Meaning that the value stored at the memory location stored in `ptr` is changed to `val`.
































#### Why:
#### How:









