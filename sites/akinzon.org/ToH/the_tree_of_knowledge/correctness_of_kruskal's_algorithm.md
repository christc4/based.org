# Correctness of Kruskal's Algorithm

Up: [Kruskal's Algorithm for Minimum Spanning Trees](kruskal's_algorithm_for_minimum_spanning_trees)
Related Note(s): [Why The General Algorithm for MSTs Works](why_the_general_algorithm_for_msts_works)
TARGET DECK

This is the logic of why Kruskal's Algorithm works:
	- Before every iteration, we have a forest of [trees](trees), $G_A = (V,\ A)$.
	- We select an edge $(u, v) \in E - A$ where:
		- $w(u, v)$ is the minimum
		- $u \in T_u$ and $v \notin T_u$, where $T_u$ is a tree in $G_A$ that contains $u$.
	- We can see that:
		- $(T_u, V - T_u)$ would be a [Cuts in Graph Theory|cut](cuts_in_graph_theory|cut)
		- $(u, v)$ is a [Light Edges in Graph Theory|light edge](light_edges_in_graph_theory|light_edge) that crosses the cut
		- Since $A$ is the set of edges in our forest of trees, there are no edges that cross the cut and therefore our cut respects A.
	- Therefore, from the [Why The General Algorithm for MSTs Works|theorem](why_the_general_algorithm_for_msts_works|theorem), we know that $(u, v)$ is a safe edge for $A$.


































#### Why:
#### How:









