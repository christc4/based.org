# Deleting from Trees

Up: [Trees](trees)
Brother(s):
TARGET DECK

When deleting a node from a tree, there are three cases to consider:
	- **The node has no children:** Mark the node's parent as pointing to null
	- **The node has one child:** Make the node's parent point to the node's child
	- **The node has two children:** Replace the node with its successor (smallest node $\ge$ our node), then delete the successor according to its own situation (is [Recursion|recursive](recursion|recursive).)




































#### Why:
#### How:









