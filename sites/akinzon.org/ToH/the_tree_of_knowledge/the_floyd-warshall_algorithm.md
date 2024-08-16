# The Floyd-Warshall Algorithm

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Floyd-Warshall Problem Background](floyd-warshall_problem_background)

The Floyd-Warshall Algorithm is a method for calculating the shortest path between any 2 nodes $i$ and $j$  in a [Graphs|graph](graphs|graph).

Let $d_{ij}^k$ be the weight of the shortest path from $v_i$ to $v_j$ and the intermediate vertices of $i$ and $j$'s path be the set $\{v_1,...,v_k\}$.

The algorithm considers these cases:
	 - If $k$ is 0, that is, there are no intermediate nodes the weight of the shortest path is simply the cost of the edge $(i, j).$
	 - If $k$ is not 0, we take the smaller of these two cases:
		 - $v_k$ **is not** part of the shortest path so the algorithm instead considers the subproblem $d^{k-1}_{ij},$ checking $v_{k-1}$ instead.
		 - $v_k$ **is** part of the shortest path so we split on $v_k$ into two smaller subproblems of calculating $d^{k-1}_{ik}$ and $d^{k-1}_{kj}$.


**Floyd-Warshall Recurrence:**

$$ d_{ij}^k = 

\begin{cases}

w(i,j)\ \text{     if k = 0 }\\
min(d^{k-1}_{ij},d^{k-1}_{ik} + d^{k-1}_{kj})\ \text{if k > 0}




\end{cases}
$$


































#### Why:
#### How:









