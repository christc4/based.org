# Cache Mapping

Up: [Cache](cache)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

Why is it needed?
Well, cache is much smaller than RAM, and so there are fewer cache lines than memory blocks.

With cache mapping, blocks of cache are allocated to represent certain memory addresses.

There are two methods - [Direct Mapping](direct_mapping) (simpler) and [Associative Mapping](associative_mapping) (most common)

The benefit of associative and set-associative mapping over direct mapping is that it avoids the issue of clashes - multiple blocks may map to the same cache line and so it would be better to have a set free lines so that clashes would all be stored in that set's cache lines.

What is cache mapping #flashcard 
Cache mapping is allocating blocks of cache to represent memory addresses
<!--ID: 1700775011460-->


What are the two types of cache mapping, and what are their benefits #flashcard 
Direct mapping - memory blocks can only go to a single cache line
Associative mapping - memory blocks can go to a any/a set of cache lines
Direct mapping is easier to implement but Associative mapping deals with the issue of memory blocks going to the same cache line (clashes)
<!--ID: 1700775011478-->


























#### Why:
#### How:









