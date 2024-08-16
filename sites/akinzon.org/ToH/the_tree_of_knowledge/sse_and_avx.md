# SSE and AVX

Up: [The Processor](the_processor)
Brother(s): [Other CPU Improvements](other_cpu_improvements) 
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::The Processor

SSE stands for streaming SIMD (Single input, multiple data) extensions, where the same instruction is performed on multiple bytes of data at the same time (sounds quite cool, eh?)

Needs special registers and instructions tho

![Pasted image 20231124153503.png](pasted_image_20231124153503.png)

SSE in action ^  and how it deals with overflows

![Pasted image 20231124153909.png](pasted_image_20231124153909.png)


### SSE uses:
 - Image processing

 - Video processing

 - Array/Vector processing

 - Text processing

 - It's also just a good speed-up in general, so long as it's supported by the application.

 - Also used in [RAID](raid), to speed up parity calculations (RAID 5-6)

 - Even used in [The GPU|GPUs](the_gpu|gpus)



### AVX, AVX-512

Extension of SSE, is even better with its even longer register width


What is SSE, and what is key to it working #flashcard 
[SSE and AVX|SSE](sse_and_avx|sse) can be used to perform the same instruction on multiple pieces of data at the same time through the use of large registers.
<!--ID: 1700856164916-->

























#### Why:
#### How:









