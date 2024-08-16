# Time Complexity of the Bellman-Ford Shortest Path Algorithm

Up: [Bellman-Ford Shortest Path Algorithm](bellman-ford_shortest_path_algorithm)
Brother(s):
TARGET DECK

The Bellman-Ford algorithm runs the [Edge Relaxation|relax()](edge_relaxation|relax()) function on $|E|$ edges, $|V| - 1$ times.

The time complexity is $O(|V| \cdot |E|)$ as a result.

In the worst case there can be $O(n^2)$ edges, so the algorithm is $O(n^3)$ in the worst case.

In the worst worst case (due to [Negative Cycles in Graph Theory|negative cycles](negative_cycles_in_graph_theory|negative_cycles)) without some way of terminating the algorithm, the algorithm will never terminate.


































#### Why:
#### How:









