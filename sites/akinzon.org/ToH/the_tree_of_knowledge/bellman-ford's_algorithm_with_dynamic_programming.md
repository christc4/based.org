# Bellman-Ford's Algorithm with Dynamic Programming

Up: [Dynamic Programming](dynamic_programming)
Brother(s): [Bellman-Ford Shortest Path Algorithm](bellman-ford_shortest_path_algorithm)
TARGET DECK

[Bellman-Ford Problem Background](bellman-ford_problem_background)

Given that this alteration of the algorithm is looking for the shortest path using **at most** $i$ edges,

The shortest path is the same as the smallest out of:
	- The shortest path to a node $u$ in at most $i-1$ edges
	- The shortest path to any neighbour $v$ of $u$ + the cost of the edge $(v, u)$

**Bellman-Ford Altered Recurrence:**

$$
d[u, i] = min 

\begin{cases}
d[u, i-1] \\

min_{v \in V} (d[v, i-1] + c_{v,u})




\end{cases}
$$





























#### Why:
#### How:









