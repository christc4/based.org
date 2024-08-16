# [C Structures|Structs](c_structures|structs) and [C Arrays|Arrays](c_arrays|arrays)

Up:
Brother(s):
TARGET DECK


An array can store multiple structs.

This is how such arrays are declared:

```C
struct structName arrayName[n];
```

Or with the use of [typedef](typedef) to shorten declaration:

```c
newStructName arrayName[n];
```

Arrays of structures are initialised in a similar way to [C Multidimensional Arrays|multidimensional arrays](c_multidimensional_arrays|multidimensional_arrays).

Structures can also have arrays as members, like so:

```C
struct student
{
	int ID;
	int marks[5];
}
```































#### Why:
#### How:









