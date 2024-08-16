# Interrupts

Up: [The Processor](the_processor)
Brother:

Interrupts are a mechanism that allows other modules to interrupt the normal function of the processor.

Interrupts provide processing efficiency by allowing other processes to be given processor time while waiting for a slower process to finish (like an I/O operation), reducing wastefulness of the CPU.

Simply put, the processor can be free to execute other instructions while an I/O operation is in progress.
![Pasted image 20231109204629.png](pasted_image_20231109204629.png)

Central to the idea of interrupts is [Context Switching](context_switching)

TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Digital Logic

What are interrupts and how do they provide processing efficiency #flashcard
Interrupts are a mechanism that allows other modules to interrupt the normal function of the processor.
They provide processing efficiency by letting the processor execute other instructions while an I/O operation is in progress.
<!--ID: 1699562664893-->

- Current instruction (stored in PC?) copied to MBR  

- special interrupt memory location loaded into MAR

- contents of MBR are written to memory 

- PC loaded with interrupt handler routine's address