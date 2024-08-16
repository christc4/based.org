# C Questions

Up: [C Principles](c_principles)
Related Note(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::C Stuff


How can [C Arrays](c_arrays) be defined and initialised? Does C check array bounds? If not, what could this lead to? #flashcard 
C arrays are defined in the form `<variableType> arrayName[length]`, and can be initialised with all entries filled, some, or even specific indices. C does not check array bounds, so if you use an illegal array bound, you will essentially be using whatever was stored at that memory location.
<!--ID: 1715360715062-->

Can you loop through [C Arrays](c_arrays) using [C Pointers|pointers](c_pointers|pointers)? #flashcard 
Yes!
<!--ID: 1715360816991-->

Does the C compiler automatically [Initialising Variables in C|initialise variables](initialising_variables_in_c|initialise_variables)? If not, what does this mean? #flashcard 
NO! As a result, the variable will retain whatever value it had.
<!--ID: 1715361141857-->



What are [C Pointers](c_pointers)? How are they linked to the [Getting The Value Stored at a Pointer in C|dereference operator](getting_the_value_stored_at_a_pointer_in_c|dereference_operator)? #flashcard 
Pointers in C are variables that store the memory address of a variable.
The dereference operator (`*`) is used to get the value stored in the memory location represented by a pointer.
<!--ID: 1715361141865-->



Can [C Pointers|pointers in C](c_pointers|pointers_in_c) be incremented for iterative purposes? If so, is there anything important to get right about this? #flashcard 
Yes! Pointers in C can be incremented. However, an important note about this is that pointers increment according to the size of the data type of the variable they are pointing to. So if you give a pointer the wrong type, it will increment by the wrong amount.
<!--ID: 1715361141871-->

What are two differences between [Memory Allocation in C|malloc](memory_allocation_in_c|malloc) and [Memory Allocation + Initialisation in C|calloc](memory_allocation_+_initialisation_in_c|calloc)? #flashcard 
Malloc simply allocates memory based on the number given to it. (takes 1 argument)
Calloc can allocate memory based on a the size of the data type used to store a single unit, and the number of units. (takes 2 arguments)
Calloc also initialises all allocated block of memory to 0 whereas malloc does no initialisation.
<!--ID: 1715361281212-->


How does [Memory Reallocation in C|realloc](memory_reallocation_in_c|realloc) work? #flashcard 
`realloc();` works by taking a pointer to the original block of memory to be resized and the size to resize it to. It may either increase the original block of memory or copy it to a new block and free the original memory.
<!--ID: 1715361467348-->

What does [Memory Allocation in C|malloc](memory_allocation_in_c|malloc), [Memory Allocation + Initialisation in C|calloc](memory_allocation_+_initialisation_in_c|calloc) and [Memory Reallocation in C|realloc](memory_reallocation_in_c|realloc) return? What do they return if they fail? #flashcard 
Malloc, calloc and realloc typically return a pointer to the block of memory that has been allocated.
If they fail however, they return a `NULL` pointer, which is a pointer that points to nothing.
<!--ID: 1715361562877-->


What gets passed by value in C, and what gets passed by reference? What does this mean?
How can we make things that get passed by value get passed by reference? #flashcard 
Everything except [C Arrays|arrays](c_arrays|arrays) get passed by value in C, meaning that a copy of the original value is passed to a function rather than the original value itself.
To pass anything else by reference, a [C Pointers|pointer](c_pointers|pointer) to the original value should be passed instead.
<!--ID: 1715361726799-->


Can we create pointers to functions in C? If so, what is one thing we can do with them? #flashcard 
Yes, we can create function pointers! They can be passed to other functions and called from there.
<!--ID: 1715363612821-->


Why is it important to [Ordering C Function Declaration|order function declarations](ordering_c_function_declaration|order_function_declarations) in C? How can the problems be solved? #flashcard 
The C compiler compiles function calls even before the function has been defined, which can lead to errors.
Either you can define functions in order, or declare functions before they're defined.
<!--ID: 1715364142292-->


What is [typedef](typedef) in C? #flashcard 
Typedef can be used to assign new names for datatypes in C.
<!--ID: 1715364248900-->


What are [C Structures](c_structures), and how do they differ from [C Arrays](c_arrays)? #flashcard 
Structures are like composite types which allow you to group together multiple types.
They differ from arrays since arrays can only store one particular type.
<!--ID: 1715364472990-->


Why do we get a conflicting types error when calling a [C Functions|function](c_functions|function) before it's been declared? #flashcard 
This is because the return type is assumed to be an `int`, which will conflict if the calling code is expecting something else.
<!--ID: 1716307998105-->





















#### Why:
#### How:









