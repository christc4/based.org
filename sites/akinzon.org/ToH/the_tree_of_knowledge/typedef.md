# Typedef

Up: [C Principles](c_principles)
Brother(s):
TARGET DECK

`typedef` is a way of assigning new names for data types in C.

For example:

```C
typedef int age;
```

create a new datatype `age` that is equivalent to `int`.

You can do the same for [c pointers|pointers](c_pointers|pointers):

```C
typedef int* intPointer;
```

where `intPointer x;` is equivalent to `int *x;`


You can do the same for [C Arrays|arrays](c_arrays|arrays):

```C
typedef int array3 [3];
```

where `array3 arr;` is equivalent to `int arr[3];`

[Structs and Typedef](structs_and_typedef)























#### Why:
#### How:









