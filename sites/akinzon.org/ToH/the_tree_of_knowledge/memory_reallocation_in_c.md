# Memory Reallocation in C

Up: [Memory Allocation in C](memory_allocation_in_c)
Brother(s):
TARGET DECK

Sometimes, it may be necessary to change the amount of memory that is being used to store data.

`realloc()` is used to achieve this.

`realloc()` takes two arguments: the pointer to the original block of memory, and the size for the new block of memory.

The data stored in the original block is migrated to the new block

`realloc()` can either extend the existing block of memory, or copy the original block of memory to a new block and free the original one.


































#### Why:
#### How:









