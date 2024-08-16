# Swapping

Up: [The Operating system](the_operating_system)
Brother: [Partitioning](partitioning) [Paging](paging)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Operating System


To begin with, [The Operating system](the_operating_system) stores a queue of processes waiting to execute in the kernel, with the actual processes themselves being stored on the disk.

Processes are swapped **onto** the disk if space is available.

Processes are swapped **off** the disk if they are finished, or if they are currently waiting for an I/O operation to complete.

![Pasted image 20231119195622.png](pasted_image_20231119195622.png)

What is meant by swapping, and when might processes be swapped onto and off the disk? #flashcard 
Swapping is a way of managing what processes need to be stored in RAM and which do not, due to there being a limited amount of main memory. 
Processes are swapped onto the disk if there is space, and are swapped off if they are finished or are waiting for an I/O operation to complete.
<!--ID: 1700423672009-->































#### Why:
#### How:









