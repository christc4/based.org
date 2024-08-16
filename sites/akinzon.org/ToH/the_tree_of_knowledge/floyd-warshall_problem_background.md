# Floyd-Warshall Problem Background

Up: [The Floyd-Warshall Algorithm](the_floyd-warshall_algorithm)
Brother(s):
TARGET DECK

The Floyd-Warshall Algorithm is used to find the [Shortest Path Algorithms|shortest path](shortest_path_algorithms|shortest_path) between all pairs of vertices in a weighted [Graphs|graph](graphs|graph).

It uses the properties of intermediate vertices in graphs.

Given a path $p$ in a graph $\langle v_1,...,v_k \rangle$ an intermediate vertex is a vertex in $p$ that isn't $v_1$ or $v_k$.

$\prod_{ij}^k$ is the set of paths from $v_i$ to $v_j$ whose intermediate paths can be drawn from the set $\{v_1,...,v_k\}$.

For a given node $v_k$, if it isn't in the shortest path from $i$ to $j$, then you try the subproblem that is $\prod_{ij}^{k-1}$.

If it is in the shortest path, it splits the path into two "shortest path" subproblems $\pi_1 \in  \prod_{ik}^{k-1}$ and $\pi_2 \in \prod_{kj}^{k-1}$.
































#### Why:
#### How:









