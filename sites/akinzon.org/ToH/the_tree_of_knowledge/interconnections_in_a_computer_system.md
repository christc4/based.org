# Interconnections in a Computer System

Up: [Digital Logic](digital_logic)
Brother:
TARGET DECK

The key takeaway is that different computer modules must be connected to each other, with different types of connections needed for each type of unit.

- Memory receives read/write control signals, as well as data and an address, and outputs data. Memory itself consists of N words from 0 to N-1. ![Pasted image 20231119205749.png](pasted_image_20231119205749.png)
- 
- I/O modules function in a similar way to memory (as far as the OS is concerned), with data passing through any one of its M ports. Furthermore, there are data paths for both internal and external devices so that I/O modules can communicate with both.![Pasted image 20231119205759.png](pasted_image_20231119205759.png)
- The [The Processor|CPU](the_processor|cpu) reads in instructions, data and also interrupt signals, and can output an address (to send to memory), control signals and also data.![Pasted image 20231119210049.png](pasted_image_20231119210049.png) 




































#### Why:
#### How:









