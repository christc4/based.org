# How Nvidia GPU architecture has evolved over time

Up: [The GPU](the_gpu)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

**Old Nvidia GTX 780, 2013:**

> [!tip]- Image:
> ![Pasted image 20240118095411.png](pasted_image_20240118095411.png)
> **Kepler architecture**
> 7 Billion transistors
> ~900 MHz core clock  
  6 GHz effective memory clock
  12 SMs, 2304 CUDA Cores, 192 CUDA cores per SM
  384 bit memory
  250 Watts
  4.15 TFLOPS (FP32)

**GTX 980, 2014:**

> [!tip]- Image: 
> ![Pasted image 20240118095819.png](pasted_image_20240118095819.png)
> Maxwell Architecture
> Fewer transistors than 780, but much more cores
> More cache
> Less power consumption

**Turing RTX 20 Series, 2018:**

> [!tip]- Image:
> ![Pasted image 20240118100029.png](pasted_image_20240118100029.png)
> MUCH more transistors (18.6B)
> 4608 CUDA cores, 576 [Tensor Cores](tensor_cores), 72 Ray tracing cores

**Ampere, 2020 and Lovelace, 2022:**

> [!tip]- Image:
> ![Pasted image 20240118100511.png](pasted_image_20240118100511.png)
> **RTX 4090**
> ishkiddibaillion (76B) transistors
> 16k CUDA cores 

From these, its clear to see why [Supercomputers](supercomputers) would love GPUs.

What are 6 ways in which [the GPU|Nvidia GPUs](the_gpu|nvidia_gpus) have changed over time, in terms of its components? #flashcard 
More transistors
More CUDA cores
Addition of tensor cores for AI
Addition of [Ray tracing](ray_tracing) cores
More cache 
More bandwidth
<!--ID: 1705597926935-->































#### Why:
#### How:









