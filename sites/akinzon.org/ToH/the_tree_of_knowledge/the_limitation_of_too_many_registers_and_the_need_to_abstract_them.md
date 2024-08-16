# The limitation of too many [Registers|registers](registers|registers) and the need to abstract them

Up: [Instruction Sets](instruction_sets)
Brother(s): 
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

Too many registers means more bits need to be used to represent each register, a higher cost for all the extra registers, and also more data needs to be saved during [Context Switching](context_switching) since there are more registers to have stored data in.


This is why physical registers are abstracted behind logical ones, the system can pretend to have only 16 but in truth it has much more - For example, the AMD 7000 zen 3 has 160 floating point registers and 192 integer registers!!!

A benefit of this abstraction is that it allows for backwards compatibility - no need to reference register 33 or something like that.

Also, it allows for [Register Renaming](register_renaming)

Is it always better to have more registers? why not? #flashcard 
Too many registers means a higher cost and more data to be saved during a [Context Switching|context switch](context_switching|context_switch).
<!--ID: 1700856164873-->


What are the benefits of abstracting physical registers behind logical registers? #flashcard 
It allows for backwards compatibility and [register renaming](register_renaming).
<!--ID: 1700856164881-->




























#### Why:
#### How:









