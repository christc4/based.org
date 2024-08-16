# Partitioning

Up: [The Operating system](the_operating_system)
Brother: [Swapping](swapping) [Paging](paging)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Operating System

Partitioning is divvying up the available amount of [RAM|main memory](ram|main_memory) so that it can be shared among currently running processes.

There are two types of partitioning: Fixed-size and Variable-size

Fixed-size partitioning gives each incoming process afixed amount of memory

- It's easier to implement since the size is fixed
- But it can be very wasteful to store a small program in more space than it needs
- While the partitions are fixed, their sizes do not need to be - incoming processes are squeezed into the smallest partition that can take them!
- ![Pasted image 20231119201317.png](pasted_image_20231119201317.png)

Variable-size partitioning gives each process the memory that it needs, based on its size

- It's much more efficient with memory usage
- With continued use, 'gaps' in the memory appear, which may make it a nightmare to add new processes to memory...
- ![Pasted image 20231119200805.png](pasted_image_20231119200805.png)
- It's all cleared up using **Compaction**, where the OS shifts all processes such that it creates one big free block of memory (Woohoo!).



































#### Why:
#### How:









