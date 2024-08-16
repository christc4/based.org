# Foundations of Shortest Path Solutions

Up: [Shortest Path Algorithms](shortest_path_algorithms)
Brother(s):
TARGET DECK

For us to calculate the shortest path, we need a couple of things first:

For the nodes in the [Graphs|graph](graphs|graph) to store its predecessor, $v.\pi$
	- At termination this will be used to return the actual shortest path itself

We also maintain a value, $v.\delta$ that represents the total weight of the shortest path from source to $v$.
	- While the program is running, this value is an upper bound on that value.

We also need an algorithm for [Edge Relaxation](edge_relaxation).




































#### Why:
#### How:









