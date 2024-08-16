# Prim's Algorithm for Minimum Spanning Trees

Up: [Minimum Spanning Trees](minimum_spanning_trees)
Brother(s):
TARGET DECK

Prim's algorithm is **vertex-based.**

Prim's algorithm works like this:
	- Maintain a [Sets|set](sets|set) $S$ of vertices, starting with a single random vertex.
	- Bearing the [Cuts in Graph Theory|cut](cuts_in_graph_theory|cut) ($S$, $V-S$) in mind each iteration, we find the edge $(u,v)$ with the smallest weight where $u \in S$ and $v \in V - S$.
	- $v$ is added to $S$ and $(u, v)$ is added to $T$.


[High-Level Implementation of Prim's Algorithm](high-level_implementation_of_prim's_algorithm)

[Time Complexity of Prim's Algorithm](time_complexity_of_prim's_algorithm)

[Correctness of Prim's Algorithm](correctness_of_prim's_algorithm)


























#### Why:
#### How:









