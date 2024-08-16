# Other CPU Improvements

Up: [The Processor](the_processor)
Brother(s): [Multi-core CPUs](multi-core_cpus)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Processor

![Pasted image 20231124144833.png](pasted_image_20231124144833.png)
### Caching hierarchy

### Branch prediction

### Deeper buffers 

### Faster virtualisation

### Instructions per clock ^
More IPC means faster processing

### Decoding width 
Wider decoding is a prefetch benefit

### Micro-op cache ^

### TURBO BOOST.
This is increasing the [The Clock|clock](the_clock|clock) speed for a short burst if only one core is in use

### Integrated memory controllers (IMCs)

^1c51c9

Giving each processor chip its own RAM bank (a chip can access another's RAM bank, but it is slower than accessing its own)

What are integrated memory controllers #flashcard 
[#^1c51c9|Integrated memory controllers](#^1c51c9|integrated_memory_controllers) give each processor chip its own RAM bank, and they can also access others with [#^cf91f7|NUMA](#^cf91f7|numa).
<!--ID: 1700856164924-->


### Non-uniform Memory Access (NUMA)

^cf91f7

The method used for one chip to interact with another chip's RAM

### VECTOR instructions [SSE and AVX](sse_and_avx)

### Neural Processing unit and more specialist process units
Encryption/Decryption
Data streaming
Matrix acceleration
^ These all free the main CPU to get on with other tasks

See [Intel vs AMD vs Apple performance improvements](intel_vs_amd_vs_apple_performance_improvements)





























#### Why:
#### How:









