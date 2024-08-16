1)Processor/clock - Clock is a component in the processor, is used to time/synchronise processor operations

2)Processor/interrupts - Interrupts are used to temporarily halt the current instruction being executed by the processor so that an instruction with **priority** can be done. They help to keep the processor busy and stop its time from being wasted on slower operations (e.g. I/O)

3)Processor/supercomputers - supercomputers have the best hardware available at the time, including the best processors.

4)Processors/SSE+AVX - SSE is used in processors to perform the same instruction on multiple pieces of data at the same time, and AVX is an extension for SSE that allows for more pieces of data to be processed simultaneously (larger register width.)

5)Processor/Cache - Processors have **cache** memory on them, a small chunk of SRAM that is searched first in the instruction cycle, in the hope of a cache hit (we find what we're looking for in cache.) If we cache miss, we go to RAM and also copy the instruction/data into cache, in case we need it later. 

6)Processor/UEFI - UEFI is the processor-independent version of BIOS

7)Processor/Server Motherboards - these motherboards usually have tons of space for multiple CPUs (due to all the processing they'll have to do)

8)Processor/Chipsets - chipsets connect all the components of a computer together, (including the processor?)

9)Processor/microprocessors - Processors of today really are microprocessors - they fit all components on a single chip.

10)Processors/Harvard Architecture - The Harvard architecture has separate stores for instructions and data, so the processor can fetch them in one cycle/fewer cycles.

11)Processors/Simultaneous multithreading - AKA hyperthreading, this is when a processor core seems to appear like multiple cores by making use of redundant/unused execution units in the processor

12) Processors/Symmetric multiprocessing - A MIMD system where multiple cores all share RAM and 
13) Processors/Von Neumann - This architecture only has one data store for both instructions and data.
14) Data movement - Processor/memory - PC -> MAR -> Cache miss -> RAM -> cache -> MBR - **uses direct memory access, does not need for the data to pass through the ALU.**
15) Processors and Parallelism - Instruction level parallelism is the number of instructions that can be parallelised with each other, machine level parallelism is all about the number of parallel pipelines that can take in parallel instructions, Data parallelism is about making data able to be used in a superscalar architecture, Then there's the one about making sure code is parallelisable as possible, reducing the number of sequential instructions as possible while retaining all functionality. (**Task Parallelism**)
16) ram and SRAM/DRAM - SRAM is a type of ram that stores data using transistors and flash memory, while dynamic RAM stores data as a charge in capacitors, which  leaks over time and must be recharged.
17) RAM and virtual memory - virtual memory is a location on the drive that is used to store pages that have been swapped out of RAM, and also stores pages currently not in use. It helps in allowing us to store programs larger than RAM.
18) RAM and accumulator-based architecture - load value into accumulator -> add with value taken directly from RAM (or cache), store value in accumulator - short instructions, but system has only one intermediate value storage, which is even temporary.
19) RAM and cache - cache acts as a buffer for RAM, the processor searches the cache first, and only goes to RAM if it cannot find what it is looking for in the cache (cache miss) - cache has 3 levels, with the third one protecting the RAM from having to interact with multiple cores.
20) Ram and Segmentation - Segmentation is similar to paging, but is visible to the programmer, who can view RAM as a set of independent address spaces, each with their own sets of permissions and rules. As a result, program parts in different segments can be recompiled independently of each other, and different programs can share the same segment.
21) OS and Task Management - The operating system is able to create **processes**, which are currently running programs and the state of said programs, and compile a list of these processes. Task management also involves setting priority levels for tasks, and managing how interrupts work.
22) OS and Paging - In the case of a page fault, the OS is required to swap in the required page to RAM. Furthermore, the OS stores a **page table**, which is a table of all pages and their locations in RAM. (Page number + offset = location in RAM)
23) OS and logical/physical addresses - The OS allows programs to think they all start at address 0 (logical address), whereas they may have a physical address of 111.
24) Networks and Computer clusters - cluster computing is a network of computers whose network is in place to increase computational power.
25) Networks and TCP/IP model - The TCP/IP model is in place as a means of standardising communication and connection between networks, through the use of 4 layers each with their own substandards.
26) Networks/ end-to-end concept - e-e concept is the idea that networks should provide best-effort connections, and that each host is responsible for reliability and security.
27) Networks and NAT - As a way of trying to solve the IPv4 exhaustion, NAT was used to essentially turn a set of internet-connected devices in the same region/area into one subnetwork with unroutable addresses, but are all represented by a single globally unique address.
28) TLB - translation lookaside buffer - the cache where page tables are stored?
29) RISC and CISC - Reduced instruction set - smaller instruction set, but more instructions per program, instructions take a more predictable number of cycles to complete. Complex instruction set - larger instruction set, fewer instructions, instructions take a more variable number of cycles to complete
30) USB vs SAS/SCSI - USB is faster than SAS, they both serially connect peripherals to a small computer
31) Pipelining and out of order issue/completion - OOI/C is good for pipelining since it decouples decoding and execution, allowing them to be run in parallel where possible. Furthermore, decoded instructions can be used to "see" ahead into the program, possibly benefitting prefetching.
32) OS and drivers - drivers are used by the OS to handle requests where I/O and peripherals are involved, and can convert machine code into device-specific control instructions.
33) PCIe and Chipsets - PCIe is an important part of chipsets since its high bandwidth, use of multiple lanes and high transfer rate (v4 - ~4-5Gbps per lane) means that data can travel around a computer at high speeds. It is also used in disk interfaces with NVMe.
