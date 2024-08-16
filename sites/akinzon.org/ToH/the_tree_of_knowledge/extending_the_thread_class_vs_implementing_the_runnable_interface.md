# [Extending the Thread Class](extending_the_thread_class) vs [Implementing the Runnable Interface](implementing_the_runnable_interface)

Up:
Brother(s):
TARGET DECK

Both ways of [Creating Threads in Java](creating_threads_in_java) work, but the `Runnable` interface is preferred due to:
	- Java's restraints on [inheritance](inheritance):  If a class extends `Thread`, it can't extend any other [classes](classes) ([Implementation vs. Extension](implementation_vs._extension))
	- The `Runnable` interface is better for sharing data



































#### Why:
#### How:









