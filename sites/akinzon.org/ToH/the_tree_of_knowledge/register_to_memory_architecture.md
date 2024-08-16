# Register to memory architecture

Up: [Instruction Set Architecture types](instruction_set_architecture_types)
Brother(s): [Accumulator based architecture](accumulator_based_architecture) [Stack based architecture](stack_based_architecture) [Register to register architecture](register_to_register_architecture)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

In this architecture, a value A is loaded from memory into a General purpose [Registers|register](registers|register) RX, with B being ADDed to RX and stored in RY, which is then stored in X.

Benefits - simple code generation, data can be accessed directly from memory

Disadvantages - Instructions require a variable amount of cycles (sounds like [CISC](cisc) eh)
You also get instructions that seem like they can be done in any order but really cannot, possibly resulting in an operand being destroyed.

How does register-memory architecture work, give advantages and disadvantages #flashcard 
In [Register to memory architecture](register_to_memory_architecture), one operand is loaded into a register and the other is used directly from [RAM](ram)
It also results in a simple code generation but its instructions require a variable amount of cycles.
<!--ID: 1700856164900-->

































#### Why:
#### How:









