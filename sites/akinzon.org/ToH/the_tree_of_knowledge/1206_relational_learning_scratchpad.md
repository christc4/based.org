Thread-local memory/method area - thread-local is private, while the method area is shared between all threads.

Class loader subsystem/method area - the class loader subsystem places the class files it has located + verified to be well-formed bytecode in the method area

class loader subsystem/the constant pool - each class is specifically placed in the constant pool in the method area.

Java stack frame/recursion in Java - any time a recursive method calls itself, a new stack frame is created.

operand stack/java stack frame - a portion of space in a java stack frame is allocated for intermediate calculations that would take place during a method's execution, this is known as the operand stack.

Serialisation/the transient keyword - the transient keyword can be used to prevent a member variable from being serialised.

Serialisation/Default serialisation - the default implementation for serialisation is as follows:
	1) Generate a hashcode for the class
	2) serialise all non-transient and non-static attributes
	3) Use a fixed encoding for primitives and Strings
	4) Recursively do the same for any object attributes.

Object serialisation/serialisation of reference types - objects only get serialised once, so if a reference to the same object is found more than once, it will still only be serialised once.

The heap/method area - the heap and the method area are both example of shared JVM memory, although they have different uses:
	- The heap is used to store objects as collections of their member variables, along with a pointer to the object's type in the method area
	- The method area is used to store class information (such as the class name, whether it's a class or an interface, and field information) as well as methods and statics. It also stores the constant pool, which is an ordered list of constants for each type. Constant here refers to every entity referred to in that class' definition, from literals to symbolic references to methods and other classes.

DO NOT USE JAVA RECURSION LIKE `while(true)` OR YOU ARE FINISHED BOY!

recursion/stack frame data - every time a method recurses, some stack frame data is created: 
	- A pointer to the method type's constant pool
	- the return address, which is where program flow should return to once the called method is done
	- A pointer to the method's exception table, which details all the possible caught exceptions by a method, and how to handle those exceptions.

stack frame/local variables in the stack frame - local variables are stored as an ordered, `[0]`-indexed array of words.
index 0 is typically reserved for the object itself - `this`. It's the same reasoning as in Python where the `self` keyword is required at the start of every non-static method.
For static methods, index 0 is not the object (duh)

Structure of a file/file i/o - all files end with an end-of-file marker.

Heaps and Arrays/method area - for arrays, as well as the array size + array contents, the heap also stores a pointer to the array class in the method area

The exception class/struct of an exception - the exception class is the base class all exceptions inherit from. 
An exception's structure is:
	- the exception type
	- information about the exception (what happened)
	- the state of the program as the exception occurred.

Thread locks/thread deadlocks - thread deadlocks occur when two or more threads are waiting for each other's lock

the externalizable interface/the transient keyword - both of these are ways to specify which attribute gets serialised. However, they go about it in different ways. Implementing the externalizable interface  means implementing methods that do the serialisation. using the transient keyword is compatible with the default serialisation implementation, which makes it simpler to use

Functional interfaces/functors/lambdas - functional interfaces are interfaces with just one abstract method. functors are concrete implementations of these functional interfaces. Lambdas are examples of functors, and are anonymous since they have no identifier.

Events/event listeners/event handlers - events are objects which represent user actions. You can listen for these events with event listeners, which don't really do much except for listen for these events. However, event handlers are a special breed of event listeners which perform an action once a listened-for event occurs.

structs and arrays/C multidimensional arrays - declaring arrays inside of structs and vice versa is similar to multidimensional arrays I guess.

Thread synchronisation/atomic blocks - atomic blocks are what we aim for in when synchronising code, we want small units of synchronised code that cannot be broken down any further, rather than synchronising an entire method even when its unwarranted.

Dependency Injection/`easyMock` - once the mock object has been created, we can then inject it into the class to be tested through
	- the constructor `new testClass(mock)`
	- a setter/mutator `test.setObj(mock)`
	- A factory method

Exceptions/why we use exceptions for error handling
just imagine creating an if branch for every single specific error case.
Or even if you tried to generalise as much as possible it would still be much more cluttered than just using exceptions.

Recording more advanced behaviours with `easyMock`/exceptions - you can tell your mock objects which exceptions to throw

prioritising threads/what we can do with threads - we can give threads a higher priority so that they essentially win more race conditions I guess

enum types vs standard static finals - static finals without enum types aren't type-safe, for example a method which takes static finals of type `int` also has to worry about taking every other integer ever, compared to that same method who only needs to worry about specific enum types

c functions and c pointers - pointers can be passed to functions (passing by reference), functions can return pointers.
Functions can also have pointers, which is used in event-driven programming for callbacks

c pointers and c arrays - c arrays are literally pointers, so no:
	- `pointer = &array` cos this is fookin wrong!!!!
	- `pointer = &array[i]` is right though

thread/monitors/wait/notify - a thread can be made to wait in the wait queue for a mutex lock using `wait()`, and notified that the lock has been freed by `notifyAll()`.

`javaFX` nodes/scenes/stages - stages are the top-level container, and act as windows, scenes define what shows up in a stage (scene graph, only one root, the root must be a `Parent` object, only can be altered in the application thread), and nodes are objects representing what actually shows 

bottom-up integration/drivers/low-level modules - bottom-up integration starts with the lower-level modules, adds on the drivers, and works its way up. These drivers need to be very complex due to the lack of a logical point of entry when doing bottom-up integration.

Big-bang integration/test cases - it's very easy to miss test cases with big-bang integration since everything is tested in isolation, without any regard for the bigger picture. If these are found, it's typically very late.

The need for stubs/mocks - when we want to test a class in isolation, we need to isolate it (duh) from all its dependencies (dependency injection) and use fake classes that pass the class we're testing test data, without having to rely on any real computation to get those values. those fake classes should just be dummies that pass the class we're testing the test data, and nothing else.

finally/further exceptions - when an exception is thrown, GG, **we are done.** When an exception is *caught*, program flow goes to the code after the `try..catch` block, which may be the optional `finally` block.