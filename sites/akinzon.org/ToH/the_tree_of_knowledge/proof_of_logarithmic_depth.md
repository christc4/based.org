# Proof of Logarithmic Depth

Up: [Calculating the Minimum Number of Nodes in an AVL Tree](calculating_the_minimum_number_of_nodes_in_an_avl_tree)
Brother(s):
TARGET DECK

we know that $m(h) = m(h − 1) + m(h − 2) + 1$, with $m(1) = 1, m(2) = 2$.

We can now prove inductively that $m(h) ≥ (3/2)^{h−1}$.

The number of elements $n$ that can be stored in an AVL Tree of height $h$ is $n \ge m(h)$ which implies $n \ge (3/2)^{h−1}$.

Taking logs gives us:

$$ log(n) ≥ (h − 1) log(3/2)  $$
Leading to:
$$ h ≤ \frac{log(n)}{log(3/2)} + 1 $$
And owing to the definition of [Big O Notation](big_o_notation), h is $O(log(n))$.



































#### Why:
#### How:









