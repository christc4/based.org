# Key Limitation of Dijkstra's Shortest Path

Up: [Dijkstra's Shortest Path Algorithm](dijkstra's_shortest_path_algorithm)
Brother(s):
TARGET DECK

## ALERT!!! ALERT!!! THIS ALGORITHM DOES NOT WORK FOR [GRAPHS](graphs) WITH [Negative Weights in Graph Theory|NEGATIVE EDGES](negative_weights_in_graph_theory|negative_edges)!!!

The negative path may not get updated accordingly, since any nodes that may have a shorter route due to that path may already be in the visited list.

![Pasted image 20240313212433.png](pasted_image_20240313212433.png)

The path a ] c ] b ] d has a cost of 2, but this wasn't updated!!!































#### Why:
#### How:









