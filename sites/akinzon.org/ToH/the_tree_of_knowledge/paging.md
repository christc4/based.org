# Paging

Up: [The Operating system](the_operating_system)
Brother(s): [Segmentation](segmentation), [Partitioning](partitioning)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Operating System

**Paging** is when you break a process into chunks known as 'pages' and break [RAM|memory](ram|memory) into frames where the size of pages = the size of frames.

these chunks/pages are usually multiple lines of instructions.

For each process, [The Operating system](the_operating_system) keeps a page table for each process, keeping track of each page.

Page numbers are converted to frame numbers using the page table

A benefit of this is that it minimises the amount of unused memory. 

The best benefit is that we can load a program **larger than main memory itself.**

This concept lends itself well to [Virtual memory](virtual_memory)


Logical to physical address mapping:

![Pasted image 20231122205057.png](pasted_image_20231122205057.png)

The page table base may be stored elsewhere and concatenated with the page number to find the corresponding frame number, which is concatenated with the page offset to find the exact instruction line.

![Pasted image 20231122205225.png](pasted_image_20231122205225.png)


[TLB](tlb) - a useful piece of kit for paging

![Pasted image 20231122205930.png](pasted_image_20231122205930.png)

What is meant by 'Paging'? #flashcard 
Paging is when you break down a process into chunks called 'pages', and breaking memory into its own chunks called 'frames', where the size of frames and pages is equal.
<!--ID: 1700688002115-->


What is the main process of demand paging(virtual memory), and how is this achieved? #flashcard 
Once a process is divided into pages, each page is swapped on and off main memory when needed. The OS stores all a process' pages in a page table, with the page number + offset of instruction being used to translate to the address of the instruction in main memory.
<!--ID: 1700688002122-->


What is a key benefit of Paging? #flashcard 
A key benefit of [Paging](paging) is that it allows for programs larger than main memory to be run.
<!--ID: 1700688002128-->























#### Why:
#### How:









