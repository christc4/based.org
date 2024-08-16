# Memory Allocation + Initialisation in C

Up: [Memory Allocation in C](memory_allocation_in_c)
Brother(s):
TARGET DECK

Memory that was allocated using `malloc()` stays as-is: This is to say that the contents of those memory locations is not changed by `malloc()`.

`calloc()` is a function similar to `malloc()` but is different in the way that it takes 2 arguments instead of 1:
	- The size of one individual unit
	- The number of units to allocate

Furthermore, `calloc()` initialises the bytes in all memory blocks to 0.


































#### Why:
#### How:









