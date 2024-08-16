# Vertex Cover of size $k$ Problem

Up: [Approximation Algorithms and Vertex Covers](approximation_algorithms_and_vertex_covers), [Backtracking](backtracking)
Brother(s):
TARGET DECK

This is an [Complexity Classes|NP-Complete](complexity_classes|np-complete) problem for verifying if there is a vertex cover of size $k$.

This is done by selecting an edge $(u, v)$ where either:
	- $v$ is in the solution: 
		- Construct [Graphs|graph](graphs|graph) $G_v$ from $G$ by removing $v$ and all its incident edges
		- $G_v$'s vertex cover is of size $k-1$.
	- $u$ is in the solution:
		- Construct [Graphs|graph](graphs|graph) $G_u$ from $G$ by removing $u$ and all its incident edges
		- $G_u$'s vertex cover is of size $k-1$.
After this, [Recursion|recurse](recursion|recurse) on $G_u$ and $G_v$





































#### Why:
#### How:









