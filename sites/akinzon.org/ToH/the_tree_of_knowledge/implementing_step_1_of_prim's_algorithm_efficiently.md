# Implementing step 1 of Prim's Algorithm Efficiently

Up: [High-Level Implementation of Prim's Algorithm](high-level_implementation_of_prim's_algorithm)
Brother(s):
TARGET DECK

Instead of iterating over all edges that cross the [Cuts in Graph Theory|cut](cuts_in_graph_theory|cut) ($S$, $V - S$), the best way to carry out step 1 of Prim's algorithm is to use previous results.

For any node $u$ we can calculate the distance to its neighbours $v$, and store that in $v.key$.

If we now assume that $w(x, y)$ has the smallest weight among all edges that cross the cut, then as per Prim's algorithm, $y$ is removed from $V - S$ and added to $S$.

All we need to now is consider the weights of edges $(y, v)$ for all $v \in adj[y]$.

This is because it makes sense to travel along the shortest path, it is a minimum spanning tree after all!

If we find a path to a node $u$ with a smaller weight than its current $u.key$, this value is updated to the smaller weight.
































#### Why:
#### How:









