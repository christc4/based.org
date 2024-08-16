# Symmetric Multiprocessing

Up: [Pipelining, Parallelisation and superscalar architecture](pipelining,_parallelisation_and_superscalar_architecture)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Pipelining and Superscalar architecture

**Symmetric Multiprocessing (SMP)** is a [Flynn's Taxonomy|MIMD](flynn's_taxonomy|mimd) system where multiple [The Processor|CPUs](the_processor|cpus) and cores share main memory and I/O

Hardware manages *contention*, which is deciding which processor gets the best share of resources at the current time.

It gives a reasonable performance increase to a point, but those gains diminish the more saturated [the bus](the_bus) becomes

In **typical SMP systems**, processors have their own L1/L2 cache and are connected by [Interconnections in a Computer System|interconnects](interconnections_in_a_computer_system|interconnects).

![Pasted image 20231205171048.png](pasted_image_20231205171048.png)

It is important to note that how processes map to processors significantly affects performance, if a processor must access an instruction/data from a processor far from it, it could result in slow access times for that data/instruction. Therefore, it is much easier for processor to communicate with others on the same 'die' (block of processors)

What is symmetric multiprocessing? #flashcard 
[Symmetric multiprocessing](symmetric_multiprocessing) is a [Flynn's Taxonomy|MIMD](flynn's_taxonomy|mimd) system where multiple [The Processor|CPUs](the_processor|cpus) and cores share main memory and I/O, where hardware manages which processors get access to the best share of current resources.
<!--ID: 1701799349753-->






























#### Why:
#### How:









