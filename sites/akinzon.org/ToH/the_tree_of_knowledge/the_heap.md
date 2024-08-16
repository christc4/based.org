# The Heap

Up: [JVM Shared Memory](jvm_shared_memory)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

The heap is a partition of shared memory that stores all the [objects|object](objects|object) instances.

Each object instance is represented by its instance variables, including inherited ones. These variables can be [Primitives](primitives) or other objects in the heap.

What does [the heap](the_heap) store #flashcard 
The heap stores all [Objects|object](objects|object) instances as a collection of their instance variables.
It also stores a [C Pointers|pointer](c_pointers|pointer) to the object's type that is stored in the [method area](method_area).
<!--ID: 1707422508915-->

For each [Objects|object](objects|object) stored in the Heap, does it also store the instance variables for its [Superclasses|superclass](superclasses|superclass)? #flashcard 
Yes.
<!--ID: 1716280374977-->



[Heaps and Arrays](heaps_and_arrays)






































#### Why:
#### How:









