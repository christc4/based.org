# Initialising Struct Member Variables in C

Up: [C Structures](c_structures)
Brother(s):
TARGET DECK

Struct variables can have their member variables initialised and changed, like so:

```C
employee.ID = 8;
```

And more generally:

```C
varName.memberN = val;
```

The member operator `.` is the way to access a `struct`'s member variables.

Structs can also be initialised in a similar fashion to arrays:

```C
struct employee worker1 = {1, 23, 4};
```

This initialises only the first 2 member variables:

```C
struct employee worker1 = {1, 23};
```

And you can even choose which members to initialise:

```C
struct employee worker1 = {.ID  = 3};
```


























#### Why:
#### How:









