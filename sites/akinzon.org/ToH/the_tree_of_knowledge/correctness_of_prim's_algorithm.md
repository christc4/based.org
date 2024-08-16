# Correctness of Prim's Algorithm

Up: [Prim's Algorithm for Minimum Spanning Trees](prim's_algorithm_for_minimum_spanning_trees)
Related Note(s): [Why The General Algorithm for MSTs Works](why_the_general_algorithm_for_msts_works)
TARGET DECK

This is the logic behind why Prim's Algorithm is correct:
	 - At the beginning of every iteration (except the first), Prim's algorithm selects a vertex $u$ to be removed from $Q$, where `u.key` is the smallest.
	 - Since $u$ hasn't been added yet, this means that $(u.p,\ u)$ is a [Light Edges in Graph Theory|light edge](light_edges_in_graph_theory|light_edge) for the [Cuts in Graph Theory|cut](cuts_in_graph_theory|cut) $(Q, V-Q)$.
	 - Furthermore, the unfinished MST $T$ contains edges that connect vertices already in the MST $(V - Q)$, and so the cut respects $T$.
	 - Therefore, according to the [Why The General Algorithm for MSTs Works|theorem](why_the_general_algorithm_for_msts_works|theorem), $(u.p,\ u)$ is safe for $T$.



































#### Why:
#### How:









