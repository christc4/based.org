# Adding an Element to a Set

Up: [Implementation of Sets using Binary Search Trees](implementation_of_sets_using_binary_search_trees)
Brother(s):
TARGET DECK

Now that we can use the [Implementing the Comparable Interface|compareTo()](implementing_the_comparable_interface|compareto()) [methods|method](methods|method) to compare nodes, we can easily add nodes to a set:

```
add (element)
	IF there is no root THEN
		root = element
		size = size + 1
	ELSE
		tempNode = root
		While nodeNotAdded DO
			Compare element to tempNode
			IF element == tempNode THEN
				nodeNotAdded = false
			ELSE IF element < tempNode THEN
				IF tempNode.left is not null THEN
					tempNode = tempNode.left
				ELSE
					tempNode.left = element
					nodeNotAdded = false
					size = size + 1
			ELSE
				IF tempNode.right is not null THEN
					tempNode = tempNode.right
				ELSE
					tempNode.right = element
					nodeNotAdded = false
					size = size + 1
```





































#### Why:
#### How:









