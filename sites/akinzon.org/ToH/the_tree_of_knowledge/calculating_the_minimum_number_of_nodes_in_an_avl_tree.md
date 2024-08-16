# Calculating the Minimum Number of Nodes in an AVL Tree

Up: [AVL Trees](avl_trees)
Brother(s):
TARGET DECK

Recalling the definition of an AVL tree, the minimum number of nodes in an AVL tree of height $h$ is $m(h)$.

At its most basic, this is a root node with two subtrees.

In the best case, the subtrees have equal height $h-1$. In the worst case (with it still being an AVL tree), one subtree is of height $h-1$ and the other $h-2$.

The minimum number of nodes in an AVL tree of height $h$ is therefore:
$$ m(h) = m(h-1) + m(h-2) + 1 $$

Compared with the Fibonacci sequence, $m(h) = f(h+2) - 1$

Since the Fibonacci sequence is exponential, we can say that $m(h)$ is exponential as well.

[Proof of Logarithmic Depth](proof_of_logarithmic_depth)































#### Why:
#### How:









