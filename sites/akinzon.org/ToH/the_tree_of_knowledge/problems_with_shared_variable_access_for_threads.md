# Problems with Shared Variable Access For Threads

Up: [Thread Synchronisation](thread_synchronisation)
Brother(s):
TARGET DECK

Since all threads in the [The Java Virtual Machine (Continued)|JVM](the_java_virtual_machine_(continued)|jvm) share access to [the heap](the_heap), any [variables](variables) stored within [objects](objects) in the heap can be accessed by multiple threads at the same time.

This can lead to two threads updating the same value at the same time ([Antidependencies](antidependencies)) 

Or one thread updating a value and another trying to read from it ([Data dependencies (and branch issues) that can prevent the parallelisation of instructions and pipelines|true data dependencies](data_dependencies_(and_branch_issues)_that_can_prevent_the_parallelisation_of_instructions_and_pipelines|true_data_dependencies))

In programming, these are known as **race conditions**, where two or more threads are competing to access/modify a shared resource.

































#### Why:
#### How:









