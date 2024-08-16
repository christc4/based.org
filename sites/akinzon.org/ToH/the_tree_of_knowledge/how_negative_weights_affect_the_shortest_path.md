# How Negative Weights Affect The Shortest Path

Up: [Shortest Path Definitions](shortest_path_definitions)
Brother(s):
TARGET DECK

**Indeed,** negative weights are real. They can hurt you.

However, even if this is the case, a shortest path for a graph can be defined, if two things are true:
	- The algorithm used is the [Bellman-Ford Shortest Path Algorithm](bellman-ford_shortest_path_algorithm)
	- There exist no negative cycles in the graph

This is because a graph with negative cycles will allow for the algorithm to keep going on and on, calculating new "shorter" paths since the cost decreases every time due to the properties of [Negative Cycles in Graph Theory|negative cycles](negative_cycles_in_graph_theory|negative_cycles).





































#### Why:
#### How:









