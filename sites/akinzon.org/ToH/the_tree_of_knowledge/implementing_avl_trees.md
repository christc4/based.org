# Implementing AVL Trees

Up: [AVL Trees](avl_trees)
Brother(s):
TARGET DECK

Implementing AVL trees requires additional information for each node: its **balance factor.**

This tells us how balanced each tree is, with
`balanceFactor = (height of left subtree) - (height of right subtree)`

The balance factor for proper AVL trees has three possible values:

$$ balanceFactor = \begin{cases} -1,\ h(s_r) > h(s_l) \\ 0,\ h(s_r) \ =\ h(s_l) \\ +1, h(s_r) < h(s_l) \end{cases} $$


To ensure that AVL tree implementations are self-balancing, the `balanceFactor` for each affected node would be updated after each add/remove operation, and if the `balanceFactor` becomes greater than 1 or less than -1, [Rotating Trees|tree rotations](rotating_trees|tree_rotations) can be used to rebalance the tree.



































#### Why:
#### How:









