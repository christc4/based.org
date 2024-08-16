# Primitives vs Objects

Up: [Programming Principles](programming_principles)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1202 Programming::Programming Principles

| [Primitives](primitives) | [Objects](objects) |
| ---- | ---- |
| Defined in Java | Defined in [Classes](classes) |
| Stored directly in variables  | Stored as a reference to a [RAM\|memory](ram\|memory) location |
| Passed directly by value | Passed by reference |

#### Definition of Primitives and Objects

Primitives are just defined by Java, we don't need to [Constructors|instantiate](constructors|instantiate) them, we just use them.

Objects can be defined by you creating a class for them to become.


#### Storage of Primitives and Objects

Primitives can easily be stored in a "variable space", think of it as a cup, a primitive can fit in a cup. Objects cannot, and so instead a **reference** to the object's location in memory is stored in the metaphorical cup instead.


#### Passage of Primitives and Objects

Primitive values can easily be passed by copying the value from variable `a` to `b`. On the other hand, objects are too big to safely do that for all the time, so instead their references are passed instead. In other words, you would have a copy of a primitive in two different variables, but the same object would be referenced by two different variables. Changing one variable changes the object, changing it for the other variable too (side effects?)

What are the main differences between how [Primitives](primitives) and [Objects](objects) are defined, stored, and passed #flashcard 
Primitives are defined in java, objects are defined in classes.
Primitive [variables](variables) the actual value, [Objects|object](objects|object) variables store a reference to the object in [RAM|main memory](ram|main_memory).
Primitive values are passed as a copy of the value, while object references are passed instead of the actual object.
`
<!--ID: 1704916722340-->


































#### Why:
#### How:









