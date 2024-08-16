# Thread-Local Memory

Up: [JVM Private Memory](jvm_private_memory)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

In these segments of private memory, the [Instruction cycle|program counter](instruction_cycle|program_counter) holds the address of the next bytecode instruction (with these instructions being stored in the [method area](method_area)), with one program counter per thread.

The [JVM Stacks|stack](jvm_stacks|stack) holds state information, such as each [The Java Stack Frame|stack frame](the_java_stack_frame|stack_frame) for each currently running method. 



Works similar to [Context Switching](context_switching)

How does the program counter work in the [The Java Virtual Machine (Continued)|JVM](the_java_virtual_machine_(continued)|jvm)? #flashcard 
The Program Counter stores a [C Pointers|pointer](c_pointers|pointer) of the next bytecode instruction to execute.
<!--ID: 1716281492641-->




























#### Why:
#### How:









