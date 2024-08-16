# Referencing Structs using their pointers

Up: [Structs and Pointers](structs_and_pointers)
Brother(s):
TARGET DECK

Struct member variables can be accessed through a `struct`'s pointer, like so:

```c
(*pointerName).member;
```

with the reason for brackets being that the member operator binds tighter and takes preference over the dereference operator.

Another way of doing it would be with the structure pointer operator:

```c
pointername->member;
```
































#### Why:
#### How:









