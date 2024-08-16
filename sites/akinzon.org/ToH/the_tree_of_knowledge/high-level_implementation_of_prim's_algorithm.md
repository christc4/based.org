# High-Level Implementation of Prim's Algorithm

Up: [Prim's Algorithm for Minimum Spanning Trees](prim's_algorithm_for_minimum_spanning_trees)
Brother(s):
TARGET DECK

Initialise with [sets](sets) $S$ and $T$, set them to both be $∅$. Also initialise with a set for $V - S$.

Then [Iteration|iteratively](iteration|iteratively) do the following: 
	- Find $x \in S, y \in V - S$ such that $w(x, y) \le w(u, v),\ \forall v \in S, v \in V - S$  
		- This means to find an $x$ and a $y$ such that their edge weight is the smallest of any edge weight that crosses the [Cuts in Graph Theory|cut](cuts_in_graph_theory|cut) ($S$, $V-S$).
	- Remove $y$ from $V - S$ and add it to $S$, and add the edge $(x, y)$ to $T$.
	- Repeat until $V - S = \mathbb{∅}$.

Step 1 is not easy to implement efficiently because if done wrong (brute force), you end up with $O(n^2)$ for **each** iteration!!!

[Implementing step 1 of Prim's Algorithm Efficiently](implementing_step_1_of_prim's_algorithm_efficiently)

[Prim's Algorithm in Java](prim's_algorithm_in_java)






























#### Why:
#### How:









