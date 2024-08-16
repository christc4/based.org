# How Do Tree Iterators Work?

Up: [Implementation of Tree Iterators](implementation_of_tree_iterators)
Brother(s):
TARGET DECK

To iterate through the elements in a tree iterator, we:
	- Start at the leftmost branch
	- To find the next node we do the following:
		- If there is a right child, move right then go as far left as possible
		- Else, go up until a node with a right child is found.


> [!TIP]- Trace the iteration through this tree. ![Pasted image 20240212092731.png](pasted_image_20240212092731.png)
> ![Pasted image 20240212092938.png](pasted_image_20240212092938.png)







































#### Why:
#### How:









