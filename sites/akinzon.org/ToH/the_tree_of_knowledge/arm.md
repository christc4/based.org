# ARM

Up: [Instruction Sets](instruction_sets)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

ARM is a family of [Instruction Set Architecture VS Microarchitecture|instruction set architectures](instruction_set_architecture_vs_microarchitecture|instruction_set_architectures), and stands for Advanced [RISC](risc) Machine.

Typically, it gives you a better performance per watt than x86.

3 processor designs:
- A for Application processors
- R for Real-time processors
- M for [Microcontrollers](microcontrollers)

3 Instruction sets:
 - A32 - 32-bit instruction set (has 16 32-bit registers), registers 13,14,15 are the stack pointer (SP), link register (LR), and program counter respectively (PC). Also has a Current Program Status Register (CPSR)
 - A64 - 64-bit version of A32, and is backwards-compatible with it too (31 64-bit registers), dedicated SP - even has 32 128-bit registers on some [SSE and AVX|SIMD](sse_and_avx|simd) extensions
 - T32 - stores instruction set in 32/16 bits, with a focus on getting more code in a smaller space. Mainly used in [Microcontrollers](microcontrollers)



































#### Why:
#### How:









