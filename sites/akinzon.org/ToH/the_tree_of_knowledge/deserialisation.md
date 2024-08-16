# Deserialisation

Up: [Serialisation and Deserialisation](serialisation_and_deserialisation)
Brother(s): [Serialisation](serialisation)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

**Deserialisation** is the reverse process of serialisation, where objects are read from a file. Those objects must have been serialisable with by implementing the *Serializable* [Interfaces|interface](interfaces|interface).

We use an `ObjectInputStream` which takes an `InputStream` [Objects|object](objects|object) as a parameter.

The method that does the object reading is:

```java
Object readObject()
```

which tries to read objects from our underlying [Java Input Streams|input stream](java_input_streams|input_stream).

This method can throw two kinds of [Exceptions|exception](exceptions|exception):
	- `IOException`: Perhaps the file was not found
	- `ClassNotFoundException`: The definition for the [Classes|class](classes|class) that the object was an instance of does not exist as any known type.

[Deserialisation Example](deserialisation_example)

Deserialisation can be customised by implementing your own `writeObject` and `readObject`, but this is rather tricky.

https://howtodoinjava.com/java/serialization/custom-serialization-readobject-writeobject/





























#### Why:
#### How:









