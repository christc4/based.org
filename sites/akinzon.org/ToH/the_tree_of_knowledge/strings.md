# Strings

Up: [Objects](objects)
Brother(s):
TARGET DECK

Strings are a data type that is used to store text.

[Comparing Strings to Each Other](comparing_strings_to_each_other)

Strings are a sequence of [characters](characters), but they aren't just represented as an array of [Unicode](unicode) values.

They are full objects in their own right!

Whenever you use a string in Java, even if it's a string literal, the [The Java Virtual Machine|JVM](the_java_virtual_machine|jvm) allocates space on the [Heaps|heap](heaps|heap) for the object that represents that string.

Duplicate string literals (e.g. using "hello" multiple times in the same program) all point to the same object in the heap.

Strings are surrounded by double quotes:

```java
String hi = "hello";
```

[Constructing a String in Java](constructing_a_string_in_java)

[A new String is Born](a_new_string_is_born)

[Strings as Anonymous Objects](strings_as_anonymous_objects)

[Strings - The Immutable Object](strings_-_the_immutable_object)

[Useful String Methods](useful_string_methods)























#### Why:
#### How:









