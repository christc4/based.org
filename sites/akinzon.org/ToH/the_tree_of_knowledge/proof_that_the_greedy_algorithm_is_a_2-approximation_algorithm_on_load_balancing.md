# Proof That The Greedy Algorithm is a 2-approximation Algorithm on Load Balancing

Up: [Approximation Algorithms and Load Balancing](approximation_algorithms_and_load_balancing)
Brother(s):
TARGET DECK


For this proof, we will let $M_x$ be the machine with the largest load when the algorithm terminates (its load is the makespan)

Let $J_y$ be the last task assigned to $M_x$ and have processing time $t_y$.

Let $L^*$ be the load of $M_x$ before $J_y$ was assigned to it.

Since the load of $M_x$ becomes the makespan when $J_y$ is assigned to it, we can say that the makespan here $= L^* + t_y$.

Because $M_x$ was chosen, it must mean that it had the smallest load before $J_y$ was assigned, since that's how the greedy algorithm was defined.

This also implies that $L^*$ was smaller than (or equal to) the running average just before $J_y$ was assigned, $\frac{1}{m} \sum_{j=1}^{y-1} t_j$.

Since all processing times are nonnegative the average of all $n$  jobs, $\bigg ( \frac{1}{m} \sum_{j=1}^{n} t_j \bigg )$, is $\ge$ the running average before $J_y$ was assigned.

Therefore we have:

$$ L^* \le \frac{1}{m} \sum_{j=1}^{y-1} t_j \le \frac{1}{m} \sum_{j=1}^{n} t_j  $$

Given that the [Greedy Algorithms|greedy algorithm](greedy_algorithms|greedy_algorithm)'s makespan is given by $G(I) = L^* + t_y$, we can then argue the following:

$G(I) \le \frac{1}{m} \sum_{j=1}^{n} t_j + max_{1 \le i \le n}\ t_i$ since $L^* \le \frac{1}{m} \sum_{j=1}^{n} t_j$ (look above for why) and $t_y \le  max_{1 \le i \le n} t_i$ since $t_y$ is not necessarily the largest job, but $max\ t_i$ is by necessity.

$\frac{1}{m} \sum_{j=1}^{n} t_j + max_{1 \le i \le n}\ t_i$ is $\le$ $max\bigg \{  \frac{1}{m} \sum_{j=1}^n t_j,  max_{1 \le i \le n}\ t_i \bigg \}$ since if one is larger than the other (lets call this larger one $z$ and the smaller one $w$) then $2z$ must be larger than $z + w$ since $w$ is 'dragging the sum down'. If $z = w$ then $z + w = 2z$.

Now, from the [Finding a Lower Bound for the Optimal Solution to Load Balancing|lower bound](finding_a_lower_bound_for_the_optimal_solution_to_load_balancing|lower_bound) we got, we know that $max\bigg \{  \frac{1}{m} \sum_{j=1}^n t_j,  max_{1 \le i \le n}\ t_i \bigg \} \le OPT(I)$ so therefore: 

$$ 2 \cdot max\bigg \{  \frac{1}{m} \sum_{j=1}^n t_j,  max_{1 \le i \le n}\ t_i \bigg \} \le 2 \cdot OPT(I) $$
And finally:

$$ G(I) \le 2 \cdot OPT(I) $$

And we are done.
























#### Why:
#### How:









