# Default Deserialisation Process

Up: [Deserialisation](deserialisation)
Brother(s): [Default Serialisation Process](default_serialisation_process)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

This process has some prerequisites:
	- The superclass is also serialisable
	- Or the superclass has a [Constructors|constructor](constructors|constructor) with no parameters.

What are the prerequisites for [deserialisation](deserialisation)? #flashcard 
That the superclass is serialisable, or that it has a constructor with no parameters.
<!--ID: 1707421752283-->


![Pasted image 20240201160003.png](pasted_image_20240201160003.png)

Once we have established that either one of the two prerequisites is true, we can begin:
1) Read the class type from the stream (Class A)
2) Find the highest serialisable [Superclasses|superclass](superclasses|superclass) (Class B)
3) Call the no-parameter constructor of the superclass' superclass (Class C)
	- This builds an [Objects|object](objects|object) and provides state to this superclass' fields. The state of all fields from the highest serialisable class are then read.
4) The above process is then repeated all the way down to the target class (Class A).


What is the process of [deserialisation](deserialisation)? #flashcard 
Read the class type from the stream
Find the highest serialisable superclass
Call the no-parameter constructor of the superclass' superclass
Update the state of all the fields from that class down to the target class
<!--ID: 1707421874790-->




































#### Why:
#### How:









