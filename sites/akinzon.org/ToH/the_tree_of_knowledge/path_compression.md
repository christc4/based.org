# Path Compression

Up: [Improvements to Quick-Union](improvements_to_quick-union)
Brother(s):
TARGET DECK

Path compression involves the shortening the distance to the root for nodes in the path to the root from a node `x` when we call `find(x)`.

What this means is that when we call `find(x)`, we not only find the root and therefore [Equivalence Classes|equivalence class](equivalence_classes|equivalence_class) of `x`, but also for all nodes we had to traverse to before finding the root!

This feature of `find(x)` is then used to make all nodes in that path a direct child of the root and flatten the [Trees|tree](trees|tree), and turn it into a [Multiway Trees|multiway tree](multiway_trees|multiway_tree).




































#### Why:
#### How:









