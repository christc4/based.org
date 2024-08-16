# C Multidimensional Arrays

Up: [C Arrays](c_arrays)
Brother(s):
TARGET DECK

Multidimensional arrays are declared similarly to normal arrays:

```C
int ma[3][4];
```

There are many ways of initialising multidimensional arrays:

```C
int ma[3][4] = {
	{1,2,3,4},
	{5,6,7,8},
	{9,10,11,12}
};
```

or:

```C
int ma[3][4] {1,2,3,4,5,6,7,8,9,10,11,12};
```

And C will just chop up the array accordingly.

Specific elements can also be initialised:

```C
int matrix[4][4] = {[3][3] = 3};
```






























#### Why:
#### How:









