# How To Serialise an Object

Up: [Object Serialisation](object_serialisation)
Brother(s):
TARGET DECK

We use a `java.io.ObjectOutputStream`, a type of [Java Output Streams|output stream](java_output_streams|output_stream) specifically designed to work with objects.

`ObjectOutputStream` objects take another output stream as a parameter, typically the thing we are going to write to (like `FileOutputStream`).
`
```java
ObjectOutputStream oos = new ObjectOutputStream(out);
```

Where `out` is an `OutputStream`.

Then, we use the [Methods|method](methods|method) 

```java
void writeObject(Object obj)
```

to actually serialise the `obj` object, and then write it to the underlying output stream.

[Object Serialisation Example](object_serialisation_example)


































#### Why:
#### How:









