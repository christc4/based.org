# Proof that LP Relaxation Solution for the Weighted Verted Cover is a 2-Approximation

Up: [Weighted Vertex Covers](weighted_vertex_covers)
Brother(s):
TARGET DECK

Since $C = \{v \in V: x_v \ge 0.5\}$ we can say that:
$$ \sum_{v \in C} w(v) \le \sum_{v \in C} 2 \cdot x_v \cdot w(v) $$
(If they were all 0.5 it would be equal, but so long as one of $C$ is greater than 0.5 it would be $\lt$ )

Then, since $x_v \ge 0$ for all $v \in V$, we can say:
$$  \sum_{v \in C} w(v) \le \sum_{v \in V} 2 \cdot x_v \cdot w(v) $$
(if for all the $v \notin C$ it was 0, it would be equal, but so long as one of $V \cap C'$ is greater than 0 it would be $\lt$)
Since $\sum_{v \in V} w(v) \cdot x_v$ is the [The Objective Function|objective function](the_objective_function|objective_function) for the [Linear Programming Problems|LP Problem](linear_programming_problems|lp_problem), and $\sum 2 \cdot x = 2 \cdot \sum x$, we can say:

$$ \sum_{v \in C} w(v) \le 2 \cdot z_r^* $$
where $z_r^*$ is the optimal solution to the LP problem.

And since $z_r^* \le z^*$ owing to our [Finding a Lower Bound for the Weighted Vertex Cover|lower bound](finding_a_lower_bound_for_the_weighted_vertex_cover|lower_bound), we can say:

$$ \sum_{v \in C} w(v) \le 2 \cdot z^* $$
Where $z^*$ is the optimal value for the original problem, and we are finished.





























#### Why:
#### How:









