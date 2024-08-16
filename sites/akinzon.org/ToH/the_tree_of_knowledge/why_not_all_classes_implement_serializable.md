# Why Not All [Classes](classes) Implement Serializable

Up: [Serialisation](serialisation)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

While being able to serialise [objects](objects) is a great thing, we don't make all objects implement the *Serializable* [Interfaces|interface](interfaces|interface). 

**Why?:** Because Java as a programming language is still developing, and serialising an object is like creating a time capsule. Perhaps its implementation gets changed, and so any outdated object of that class can no longer be deserialised.

However, by using a pre-defined `serialVersionUID` we essentially encode the method for serialising and deserialising that object *with the object itself*, allowing it to be deserialised the way its version dictates, and not be made undeserialisable due to it being of an outdated version.

In IntelliJ, we can add a default serial version ID to help us with this.

Why don't all [objects](objects) implement *Serializable* #flashcard 
This is because their implementation may change, and may become undeserialisable in future versions of Java.
<!--ID: 1707421124299-->






































#### Why:
#### How:









