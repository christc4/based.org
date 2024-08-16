# High-Level Implementation of Kruskal's Algorithm

Up: [Kruskal's Algorithm for Minimum Spanning Trees](kruskal's_algorithm_for_minimum_spanning_trees)
Brother(s):
TARGET DECK

Given a graph $G$ and a weight [Functions|function](functions|function) $w : E \rightarrow \mathbb{R}$.

We build the minimum spanning tree, $T$ by adding edges to $T$ one at a time.

Since we are [Iteration|iterating](iteration|iterating) through the edges from smallest weight to largest, we could sort [VVEEMAD|E](vveemad|e).

The next thing needed is to check that adding edges to the current spanning tree creates a cycle.

The [union-find](union-find) [Data Structures|data structure](data_structures|data_structure) can be useful for this!

[Kruskal's Algorithm in Java](kruskal's_algorithm_in_java)




































#### Why:
#### How:









