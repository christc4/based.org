# Direct Mapping

Up: [Cache Mapping](cache_mapping)
Brother(s): [Associative Mapping](associative_mapping)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

#### Direct mapping

Each block of main memory maps to one cache line, and can only be found **in that specific one.**

Memory addresses are made up of two parts: The least significant w bits that identify the particular word we want, and the most significant s bits that specify one block of memory. the s bits are interpreted as a tag of size s-r bits and a cache line of size r bits.

![Pasted image 20231123203604.png](pasted_image_20231123203604.png)

![Pasted image 20231123203620.png](pasted_image_20231123203620.png)

^AS YOU CAN SEE, the cache line stored in the memory address is used to return its tag. If the tags do not match, then it is a cache miss. 



































#### Why:
#### How:









