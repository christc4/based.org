# Daemon Threads

Up: [The Two Types of Threads](the_two_types_of_threads)
Brother(s): [User Threads](user_threads)
TARGET DECK

Daemon Threads sound cool, but they're not treated with the respect they deserve!

Daemon threads consist consist of tasks that happen in the background that exist only to serve the user threads.

If the [The Java Virtual Machine (Continued)|JVM](the_java_virtual_machine_(continued)|jvm) wants to exit, and a daemon thread is running, the JVM will exist (doesn't care at all wow)

Because of this, it's advised not to make all threads daemon threads.

Threads created inside a daemon threads are also daemon threads.

































#### Why:
#### How:









