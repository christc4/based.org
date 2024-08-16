# Default Serialisation Process

Up: [Serialisation](serialisation)
Brother(s): [Default Deserialisation Process](default_deserialisation_process)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

Here's how `writeObject()` does it:
1) Write a hashcode of the current Class object for the object being serialised.
2) Serialise each field, ignoring:
	- static fields
	- [The Transient Keyword|transient](the_transient_keyword|transient) fields
3) Use a fixed encoding (through `DataOutputStream`) to write [primitives](primitives) and [Strings](strings) properly
4) Recursively call `writeObject()` on any reference types (member variables that are objects) **that haven't been serialised already.**

What is the actual process of serialising [objects](objects)? #flashcard 
Write a hashcode for the object
Serialise all non-static and non-transient fields
Use a fixed encoding to write [primitives](primitives) and [Strings](strings) properly
Recursively do the same for any member variables that are objects
<!--ID: 1707421590859-->




































#### Why:
#### How:









