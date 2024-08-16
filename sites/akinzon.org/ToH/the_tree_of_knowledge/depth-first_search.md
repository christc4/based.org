# Depth-First Search

Up: [Graph Theory](graph_theory)
Brother(s): [Breadth-First Search](breadth-first_search)
TARGET DECK

Depth-first search involves going as **deep** as possible:

Start at one node, and keep on going until you reach a node with no other neighbours, then backtrack and keep going.

Depth-first search computes 2 values for each node
	- A **discovery time**: Time taken to reach node
	- A **finishing time**: Time taken to come back to this node after traversing down all its neighbours, depth-first.

We add each node to a pending [Stacks|stack](stacks|stack), and nodes are popped off when all their neighbours have been visited.

Depth-First Search (in our implementation) can reach all nodes, even separate [Connected Components in a Graph|connected components](connected_components_in_a_graph|connected_components) since we iterate through all nodes in the graph.

[Time Complexity of Depth-First Search](time_complexity_of_depth-first_search)

[Application of Depth-First Search - Topological Sorting](application_of_depth-first_search_-_topological_sorting)





























#### Why:
#### How:









