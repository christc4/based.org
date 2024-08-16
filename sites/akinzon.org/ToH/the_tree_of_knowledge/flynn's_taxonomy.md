# Flynn's Taxonomy

Up: [Pipelining, Parallelisation and superscalar architecture](pipelining,_parallelisation_and_superscalar_architecture)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Pipelining and Superscalar architecture

#### What is it???
Flynn's Taxonomy is a classification of computer architectures based on the number of instructions and data streams

#### There are four different architectures

**Single Instruction, Single Data** 
![Pasted image 20231205164913.png](pasted_image_20231205164913.png)
^ This is typical of old computers with single cores. No traces of parallelisation/[Superscalar architecture](superscalar_architecture)

****

**Multiple Instructions, Single Data**
![Pasted image 20231205165053.png](pasted_image_20231205165053.png)
^ Example: Multiplying some data by 5 then dividing by 2 (2 different instructions, both on the same data)

****

**Multiple Instruction, Single Data**
![Pasted image 20231205165216.png](pasted_image_20231205165216.png)
^ Found in typical modern, multicore processors.

****

**[Single Input, Multiple Data](single_input,_multiple_data)**

![Pasted image 20231205165539.png](pasted_image_20231205165539.png)
^ Example: [SSE and AVX](sse_and_avx)

































#### Why:
#### How:









