# Three Methods To Deal With Pipeline Issues

Up: [Pipelining, Parallelisation and superscalar architecture](pipelining,_parallelisation_and_superscalar_architecture)

Prefetching both branches then using the pipeline that satisfies the branch condition  
  
Loop buffer (for loop branches) - Keep last n instructions in very fast memory, in a similar fashion to cache  
  
Branch prediction - checking the opcode to see if the instruction would be more likely to result in a jump (like the compare operation),  
or basing decision on previous history
