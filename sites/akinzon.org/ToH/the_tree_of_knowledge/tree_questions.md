# Tree Questions

Up: [Trees](trees)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Abstract Data Types

What is the red rule for [Red-Black Trees](red-black_trees), and what is the black rule for red-black trees? #flashcard 
Red rule: The children of a red node must be black
Black rule: the number of black nodes in all paths from the root to a leaf (or a node with one child) must be equal.
<!--ID: 1709497278353-->


What separates [B+ Trees](b+_trees) from [B-Trees](b-trees)? #flashcard 
In a B+Tree, all values are stored in roots. In B-Trees, non-leaf nodes can store values.
<!--ID: 1715460005019-->

What are the [B+ Tree Rules](b+_tree_rules)? #flashcard 
Only leaves store data
Non-leaf nodes store up to $M-1$ keys
Non-leaf nodes can have between $\lceil M/2 \rceil$ and $M$ children
Leaf nodes (except the root) are at the same level and have between $\lceil L/2 \rceil$ and $L$ data entries.
<!--ID: 1715460128431-->

How do [Tries](tries) work? #flashcard 
Tries are multiway trees that store words as compactly as possible using prefixes. Each string is stored as a path from root to leaf based on each letter in the string, which links to a subtrie based on that current prefix, where all children start with that current prefix.
<!--ID: 1715460728898-->




































#### Why:
#### How:









