# Weighted Quick-Union

Up: [Improvements to Quick-Union](improvements_to_quick-union)
Brother(s):
TARGET DECK

The main goal of the weighted quick-union is to keep the tree as balanced as possible when we perform the `union()` method.

This requires the introduction of a `size[]` array to keep track of the number of items in each [Equivalence Classes|equivalence class](equivalence_classes|equivalence_class).

With these size values, when a `union(x, y)` is performed, the size of `x` and the size of `y` are compared, and the smaller of the two is added as a child of the larger one (remembering that we add by roots!!)

With this improvement, `find` and `union` both become $O(log_2(N))$!!! pretty cool, no?

































#### Why:
#### How:









