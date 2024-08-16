# Local Variables in the Stack Frame

Up: [Anatomy of a Stack Frame](anatomy_of_a_stack_frame)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

Any time [methods](methods) are called, their parameters and local variables are stored on the method's stack frame as a zero-indexed (starts at 0) array of words (words being addressable units.)

Bytecode instructions that access local variables use these array indices.

Array entries are either [primitives](primitives) or references to [objects](objects) in [the heap](the_heap).

The first entry in the array for non-static methods `array[0]` always is a reference to the object that called said method. This is used as the value of `this`, similar to calling non-static methods with the first parameter being `self` in python.

This isn't the same for static methods of course, hence why you cannot access instance variables from within a static method.

How are local variables stored in the [The Java Stack Frame|stack frame](the_java_stack_frame|stack_frame) #flashcard 
They are stored in a zero-indexed array of words.
<!--ID: 1707423219953-->

When are [objects](objects) stored in the 0th index of of the [The Java Stack Frame|stack frame](the_java_stack_frame|stack_frame) array #flashcard 
When the calling method is not static
<!--ID: 1707423305257-->





































#### Why:
#### How:









