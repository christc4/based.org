# Control Flow Questions

Up: [Control Flow](control_flow)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::Control Flow



What are [Exceptions](exceptions) in Java? #flashcard 
Exceptions are [objects](objects) that contain information about errors.
<!--ID: 1707990412664-->



Why are [Exceptions](exceptions) the [Why we don't use if statements for exception handling|best way](why_we_don't_use_if_statements_for_exception_handling|best_way) to handle errors? #flashcard 
If `if..else` statements were used to catch every single error case, it would lead to cluttered code.
<!--ID: 1707990412673-->



What do [Exceptions](exceptions) [Structure of an Exception|contain](structure_of_an_exception|contain)? #flashcard 
Exceptions contain information about the error, the error type, and the state of the program when the error occured.
<!--ID: 1707990412679-->




What is the difference between [checked exceptions](checked_exceptions) and [unchecked exceptions](unchecked_exceptions)? #flashcard 
Checked exceptions inherit from the `Exception` class, and generally require exception handling since they are out of your control.
Unchecked Exceptions inherit from the `RuntimeException` class and generally require the programmer to improve the code (COS IT'S BAD)
<!--ID: 1707990412685-->



What does anything [Statics|static](statics|static) in Java belong to #flashcard 
Statics belong to its class and not to any specific object.
<!--ID: 1707990796882-->




What are 3 benefits of using [Static Final Variables](static_final_variables) (constants) and why? #flashcard 
Their value only needs to be declared once, and can be used anywhere (in the same class.) 
This makes it easier to change as well.
It also keeps code self-documenting!
<!--ID: 1707990796908-->




What is a key difference between normal [methods](methods) and [static methods](static_methods), and what is a consequence of this? #flashcard 
Normal methods are called on their [objects](objects), whereas static methods are called on the class itself.
A consequence of this is that you can use the class name to call static methods.
<!--ID: 1707990796920-->



What is the difference between processes and [Threads|threads](threads|threads)? #flashcard 
Processes are a set of activities that produce a result when chained together,
but threads are the smallest set of programmed instructions that a scheduler can manage.
<!--ID: 1707998793449-->




What is [Java's Support for Threading](java's_support_for_threading)? #flashcard 
Java supports threading with its `Thread` class
The [The Java Virtual Machine (Continued)|JVM](the_java_virtual_machine_(continued)|jvm) gives each thread [JVM Private Memory|private memory](jvm_private_memory|private_memory) for the [JVM Stacks|java stack](jvm_stacks|java_stack) and program counter.
Threads can also [JVM Shared Memory|share](jvm_shared_memory|share) access to the [method area](method_area) and [the heap](the_heap).
<!--ID: 1707998793467-->




Is it better to [Extending the Thread Class|extend the thread class](extending_the_thread_class|extend_the_thread_class) or [Implementing the Runnable Interface|implement the runnable interface?](implementing_the_runnable_interface|implement_the_runnable_interface?) #flashcard 
It is better to implement the runnable interface because it does not restrict the class to only being able to extend the `Thread` class ([Implementation vs. Extension](implementation_vs._extension))
Furthermore, it is better for sharing data.
<!--ID: 1707998793478-->



What are the differences between [User Threads](user_threads) and [Daemon Threads](daemon_threads)? #flashcard 
User threads are user tasks, high priority, and prevent the [The Java Virtual Machine (Continued)|JVM](the_java_virtual_machine_(continued)|jvm) from closing the program. 
Daemon threads are background tasks, low priority, and do not prevent the JVM from closing the program.
<!--ID: 1708020105336-->



Despite [Objects](objects) being used as their own [Thread Locks|locks](thread_locks|locks) by default in [The Synchronized Keyword|synchronized](the_synchronized_keyword|synchronized) [methods](methods), when may it be useful to use [the synchronized block](the_synchronized_block)? (apart from it being better for smaller pieces of code) #flashcard 
If they wanted to specify a different object as the thread lock, it would be beneficial to use a `synchronized` block.
<!--ID: 1708119019854-->





What are [Thread Locks](thread_locks)? #flashcard 
Thread locks are [objects](objects) used to restrict access to synchronised [methods](methods)/code blocks.
<!--ID: 1708119019876-->




What is [Mutual Exclusion](mutual_exclusion) in threading #flashcard 
Mutual exclusion is when some code can only be run by one thread at any given time.
<!--ID: 1708119019884-->


How do we deal with [Unchecked Exceptions](unchecked_exceptions)? #flashcard 
We fix our code (cos it sucks) and sanitise user inputs.
<!--ID: 1716284107013-->

When an [Exceptions|exception](exceptions|exception) is caught, does program flow return back to where the exception occurred? #flashcard 
No! Rather, it returns to after the `try..catch` block or to the optional `finally` block.
<!--ID: 1716284267835-->

What does [yielding a thread](yielding_a_thread) do? #flashcard 
Making a thread yield causes it to give up its time to allow other threads to have its processor time.
<!--ID: 1716285267630-->

What are the 6 [Java Thread States](java_thread_states)? #flashcard 
New - The thread has just been created
Runnable - The thread has been `start()`ed
Blocked - The thread is waiting for a [Mutual Exclusion|mutex](mutual_exclusion|mutex) lock
Waiting - The thread is waiting for another thread to do something (notify it)
Timed Waiting - The thread is waiting for some time to run out
Terminated - The thread has finished running.
<!--ID: 1716285613288-->

What are race conditions in the context of [Multithreading (Programming)|multithreading](multithreading_(programming)|multithreading)? #flashcard 
Race conditions are when two or more threads compete for resources.
<!--ID: 1716287607652-->


Although [thread synchronisation](thread_synchronisation) is very useful, what are its downsides? When should we actually use it? #flashcard 
If one thread holds a lock for too long, it can cause other threads to block while they are waiting for the same lock. Furthermore, synchronisation incurs a runtime overhead.
We should always lock when updating shared fields and accessing fields that may be updated, and never lock when invoking [methods](methods) on other [objects](objects).
<!--ID: 1716287607677-->

What are the building blocks of [Monitors](monitors), and why? #flashcard 
[Thread Synchronisation](thread_synchronisation), [Making a Thread Wait|wait()](making_a_thread_wait|wait()) and [Notifying Threads|notify/notifyAll()](notifying_threads|notify/notifyall())
Synchronisation is important in preventing race conditions for data (textbook multithreading)
`wait()/notifyAll()` is important in preventing race conditions for the mutex lock (depends on your program logic)
<!--ID: 1716287897968-->


















#### Why:
#### How:









