# How B-Trees Speed up Searching

Up: [B-Trees](b-trees)
Brother(s):
TARGET DECK

B-Trees are like [Binary Trees](binary_trees) but on crack: Instead of one node containing one value and having a `left` subtree full of values smaller than it and a `right` subtree full of larger values, they have a node that contains **multiple values** with more than 2 children!!!

Functionally similar to binary trees in the way that you can navigate to subtrees based on value, it's also different in the way that due to more values per nodes, there are subtrees for cases when our search parameter $x$ is in between two values!!

In a normal binary tree, it would take more than one jump to get to our search parameter, but it can take as little as a single jump with B-Trees!

![Pasted image 20240220192517.png](pasted_image_20240220192517.png)

Getting to 7 only requires us to check for $5 \le 7 \le 9$ in one go, rather than $5 \le 7$ and then $7 \le 9$ afterwards!!! 

































#### Why:
#### How:









