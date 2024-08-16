# Finding a Lower Bound for the Optimal Solution to Load Balancing

Up: [Approximation Algorithms and Load Balancing](approximation_algorithms_and_load_balancing)
Brother(s):
TARGET DECK

A simple algorithm for load balancing would be a [Greedy Algorithms|greedy](greedy_algorithms|greedy) one that gives the next task to the machine with the lowest load.

This may not give the optimal solution, but it does work as a lower bound on the optimal solution.


Again, we denote $L(M)$ to be the load of machine $M$.

The average load would be:
$$ \frac{1}{m} \sum_{i=1}^{m} L(M_i) = \frac{1}{m} \sum_{j=1}^n t_j $$
Where $t_j$ is the time taken to complete job $j$.

Given that this is the average, then there must be a machine $M_i$ such that its load is greater than the average:
$$ L(M_i) \ge \frac{1}{m}\sum_{k=1}^n  t_k$$
Also, the largest task must be assigned to some machine $M_j$:
$$ L(M_j) \ge max_{1 \le k \le n}\ t_k $$
These two inequalities help to form a lower bound for the optimal solution, which in any case the lower bound would be the bigger of the two:
$$ OPT(I) \ge max\bigg \{  \frac{1}{m} \sum_{j=1}^n t_j,  max_{1 \le i \le n}\ t_i \bigg \}  $$
































#### Why:
#### How:









