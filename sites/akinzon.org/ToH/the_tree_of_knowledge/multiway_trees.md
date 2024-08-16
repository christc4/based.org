# Multiway Trees

Up: [Trees](trees)
Brother(s):
TARGET DECK

Multiway trees are trees where each non-leaf has a $M$ children.

While the search time for [Binary Trees](binary_trees) is $log_2(n)$ for a tree of size $n$, for multiway trees it's different:
$$ log_M(n) = \frac{log_2(n)}{log_2{(M)}} $$

Making $M \approx 2^n$ can reduce the search time by a factor of $n$.

[B-Trees](b-trees) are one such implementation of these multiway trees.

[Tries](tries) are another used for storing strings composed from elements of an alphabet.

































#### Why:
#### How:









