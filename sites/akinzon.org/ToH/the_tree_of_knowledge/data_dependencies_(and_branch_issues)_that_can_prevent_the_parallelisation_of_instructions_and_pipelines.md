# Data dependencies (and branch issues) that can prevent the [The types of parallelism|parallelisation](the_types_of_parallelism|parallelisation) of instructions and pipelines

Up: [Pipelining, Parallelisation and superscalar architecture](pipelining,_parallelisation_and_superscalar_architecture)


Pipelining has to overcome the fact that programs may branch off, which can be a major issue since it could  
result in the wrong instruction being prefetched (since we may not branch to this instruction.)
#### True Data dependency
Instruction B needs the results of instruction A (Contingency)  
  
#### Resource conflict dependency 
Instructions A and B both need resource C but cannot access it simultaneously  
  
#### Procedural dependency 
When an instruction in a branch alters a value in use by another instruction outside of the branch

