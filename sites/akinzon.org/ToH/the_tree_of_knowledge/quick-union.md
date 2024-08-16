# Quick-Union

Up: [Union-Find](union-find)
Brother(s):
TARGET DECK

In the quick-union implementation of union-find, the items are stored as a forest of [trees](trees), each with their own root.

The id of an element is the root of the tree. An element can be its own root (value of id = value of element)

Equivalence classes are `union()`ed by combining the roots; the root of the equivalence class to be changed is made a child of the equivalence class it is to be a part of.

To find the equivalence class of an arbitrary element, we introduce the `root(x)` function to find the root of the tree `x` is a part of, since the root represents the equivalence class.
 
[Time Complexity of Quick-Union](time_complexity_of_quick-union)

[Improvements to Quick-Union](improvements_to_quick-union)

























#### Why:
#### How:









