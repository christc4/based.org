# Stack Frame Data

Up: [Anatomy of a Stack Frame](anatomy_of_a_stack_frame)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

Stack frames also contain data that aid in [Methods|method](methods|method) execution.

This portion of the frame contains:
	 - A pointer to the method type's [The Constant Pool|constant pool](the_constant_pool|constant_pool)
	 - The address of the bytecode instruction to return to after completing this method (return address
	 - A pointer to the method's exception table which the JVM uses to handle any exceptions thrown while the method is running.

If an exception is indeed thrown, the JVM uses the exception table to look for any `catch` clauses ([exception handling](exception_handling)) that correspond to the thrown exception. If none are found, the method terminates, its stack frame is popped, and the exception is rethrown using the exception table of the previous frame.

What is stored in the [The Java Stack Frame|stack frame](the_java_stack_frame|stack_frame)'s frame data #flashcard 
A pointer to the method type's [The Constant Pool|constant pool](the_constant_pool|constant_pool)
The Return address
A pointer to the method's exception table
<!--ID: 1707423404892-->



































#### Why:
#### How:









