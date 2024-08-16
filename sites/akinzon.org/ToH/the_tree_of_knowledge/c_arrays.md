# C Arrays

Up: [C Principles](c_principles)
Brother(s): [Arrays|Arrays (Java)](arrays|arrays_(java))
TARGET DECK

Arrays in C are declared in this format:

```C
<variableType> arrayName[<arrayLength>];
```

They can also be preloaded with some/all items in them:

```C
int array[5] = {1, 2, 3, 4, 5};
```

```C
int array[5] = {[3] = 3, [4] = 71};
```

Though this means that values for the uninitialised elements are not defined, and **will take the value of whatever is in that memory address currently.**

Arrays can also be defined without specifying the number of elements, though **this means that the length of the array will be based on how many elements are in the array at initialisation.**

```C
char word[ ] = {"A", "k", "i", "n"};
```

[C Array Bounds](c_array_bounds)

[C Multidimensional Arrays](c_multidimensional_arrays)

Arrays can be declared with a variable size `n` where `n` is a number taken from [Inputting and Outputting Data in C|user input](inputting_and_outputting_data_in_c|user_input).

[Incrementing Arrays with C Pointers](incrementing_arrays_with_c_pointers)





















#### Why:
#### How:









