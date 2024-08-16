# Virtual memory

Up: [Paging](paging)
Brother: [TLB](tlb)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Operating System



Virtual [RAM|memory](ram|memory) (or demand paging) is an application of [Paging](paging) where the pages swap in and out of memory using [Swapping](swapping). Therefore, memory needs to only have a few pages of a process at a time. If a page to be branched to is not found in main memory, a page fault is triggered and the OS is told to swap the desired page into memory.

What is a page fault? #flashcard 
This is when a page that is to be branched to cannot be found in main memory and the OS is told to retrieve it from the disk and swap it into memory.
<!--ID: 1700689141101-->







































#### Why:
#### How:









