# The types of parallelism

Up: [Pipelining, Parallelisation and superscalar architecture](pipelining,_parallelisation_and_superscalar_architecture)
Brother:

**Instruction-level parallelism** is the number of independent instructions there are that can be parallelised  
  
**Machine-level parallelism** is the number of parallel pipelines there are, which take advantage of instruction-level parallelism.

**Data parallelism** is when data is split up to make independent, parallel tasks

**Task parallelism** is when code itself is split up so it can work on separate [The Processor|CPUs](the_processor|cpus).
  
Machine-level parallelism relies on the ability to parallelise instructions so that independent instructions can be placed in parallel pipelines.