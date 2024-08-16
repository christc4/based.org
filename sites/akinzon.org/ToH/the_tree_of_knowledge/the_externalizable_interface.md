# The Externalizable Interface

Up: [Serialisation](serialisation)
Related Note(s): [Interfaces](interfaces)
TARGET DECK


While the *Serializable* interface serialises all non-static properties of [objects](objects), the *Externalizable* interface allows developers to choose which properties get serialised.

Unlike the *Serializable* interface, the *Externalizable interface* comes with two methods (that must be implemented):
	- `writeExternal()`
	- `readExternal()`

Furthermore, any [classes](classes) that implement this interface must also declare a constructor with no parameters (akin to [The Default Constructor](the_default_constructor).)

[Externalizable Interface Example](externalizable_interface_example)

































#### Why:
#### How:









