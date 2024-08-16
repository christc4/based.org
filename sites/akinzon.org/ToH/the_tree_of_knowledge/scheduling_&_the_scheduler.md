# Scheduling & The Scheduler

Up: [The Operating system](the_operating_system)
Brother: [Interrupts](interrupts)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Operating System


The main goal of scheduling is to make sure [The Processor|processing](the_processor|processing) time is used as efficiently as possible, which is carried out by the scheduler.

Scheduling itself is when the OS controls which programs can send their data to the processor to be processed. Operations such as accessing memory and I/O operations are far slower than processing speed, but all these tasks must share the processor, which can only do one thing at a time. Therefore, scheduling is needed to make sure that while these slow operations (Such as I/O waiting on a device/user input) are being waited on to complete, the processor can be free to get on with other things.

![Pasted image 20231118184333.png](pasted_image_20231118184333.png)

Only one program is running, and the processor is waiting for the next instruction. What a waste of time and resources!

![Pasted image 20231118184418.png](pasted_image_20231118184418.png)

Aight, it's getting a little better now, in a portion of the waiting period, another program's instructions are being executed. Noice. But still a bit of a waste, no?

![Pasted image 20231118184511.png](pasted_image_20231118184511.png)

### Are we cooking chat??? 
As you can see, one more program is being executed in the waiting period, and the CPU is being put to WORK. But it likes to work like this :D This is MUCH more efficient than just one program running at a time, because it's never just a monolithic program running like that, it's just its different instructions which can have a large gap between them. Thus, having all of these little instructions running in a nice, neat *schedule* allows for maximal use of the CPU, which is what we wanted!!




























#### Why:
#### How:









