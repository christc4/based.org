# Implementing the Runnable Interface

Up: [Creating Threads in Java](creating_threads_in_java)
Brother(s):
TARGET DECK

Implementing the `Runnable` [Interfaces|interface](interfaces|interface) is one way of creating threads.

This interface has only one abstract method: `run()`.

You create a class that implements this interface, provide an implementation for the `run()` [Methods|method](methods|method).

#### But doesn't `start()` start the new thread?

Yes, and this is where we use the [Constructors|constructor](constructors|constructor) of the `Thread` class.

It takes any [classes|class](classes|class) that implements `Runnable` as a parameter.

Once instantiated, call `start()`!!!

[Implementing the Runnable Interface Syntax](implementing_the_runnable_interface_syntax)

[Extending the Thread Class vs Implementing the Runnable Interface](extending_the_thread_class_vs_implementing_the_runnable_interface)

































#### Why:
#### How:









