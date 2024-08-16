# Anatomy of a Stack Frame

Up: [The Java Stack Frame](the_java_stack_frame)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

Stack frames contain sufficient space for the [Methods|method](methods|method) they were created to represent. That is to say, the size of a stack frame is *method-dependent*.

This space is used to store:
	- [Local Variables in the Stack Frame|Local variables](local_variables_in_the_stack_frame|local_variables)
	- Parameters
	- [Stack Frame Data|Frame data](stack_frame_data|frame_data)
	- [The Operand Stack|Intermediate computation data](the_operand_stack|intermediate_computation_data)

What is the [anatomy of a stack frame](anatomy_of_a_stack_frame)? (4) Are all stack frames of equal size? #flashcard 
Stack frames have space for:
local variables
parameters
[Stack Frame Data](stack_frame_data)
space for intermediate computations
No, not all stack frames are of equal size. A stack frame's size depends on how much it needs to store, and is therefore *method-dependent.*
<!--ID: 1716281494714-->




































#### Why:
#### How:









