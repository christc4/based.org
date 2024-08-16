# Instruction cycle

Up: [The Processor](the_processor)``
Brother(s):
TARGET DECK

In the instruction (FDE) cycle, the processor fetches an instruction, decodes it, and then executes it.

### Fetch phase

The contents of the Program Counter (PC) are copied to the Memory Address Register (MAR), which are sent to [cache](cache)s via the address bus to see if it can be found there first. Simultaneously, the Control Unit sends a memory read request to allow for the instruction/data to be read. If not (cache miss), it is sent to [RAM](ram). (If not found there, a [Paging|page fault](paging|page_fault) is triggered and the required page is swapped into RAM). Anyway, once it is fetched, it is stored in the Memory Buffer Register (MBR), and the contents of the PC are incremented by 1. The contents of the MBR are copied to the Current Instruction Register (CIR) 


### Decode phase

The instruction is then decoded/interpreted to see what is asking the processor to do. As a result, some of the data needed may still be in memory, and will need to be fetched.


### Execute phase 

![Pasted image 20231124135124.png](pasted_image_20231124135124.png)

Some kind of process is performed on the data. At the hardware level, all possible operations are performed, and only the output of the one we need is selected using a [Multiplexors|multiplexor](multiplexors|multiplexor).



































#### Why:
#### How:









