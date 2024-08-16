# The Weighted Interval Scheduling Algorithm

Up: [Dynamic Programming](dynamic_programming)
Brother(s): [The Interval Scheduling Algorithm](the_interval_scheduling_algorithm)
TARGET DECK

The goal here is to find the most amount of non-overlapping time intervals, but also **where the weight is a maximum.**

The algorithm is in similar spirit to the [Applying Dynamic Programming to The Independent Set Problem|independent set problem](applying_dynamic_programming_to_the_independent_set_problem|independent_set_problem), since here we consider the last time interval again, and it has two cases as well.

Either it is in $\mathcal{O}$ and therefore we then consider the next time interval that does **not** overlap with this one, or it is not in $\mathcal{O}$ and consider the most recent previous time interval.


$$ opt(n) = \text{max} 
\begin{cases} 

w_n + opt(I_n)\ (x_n\ \in \mathcal{O}) \\
opt(n-1)\ (x_n\ \notin \mathcal{O})


\end{cases} 

$$


Where $I_n$ is the set of intervals that do not overlap with this interval and also come before it.

[Use of the p Array for The WIS Algorithm](use_of_the_p_array_for_the_wis_algorithm)

This array helps to simplify the [Recurrence Relations|recurrence relation](recurrence_relations|recurrence_relation) to:

$$ opt(n) = \text{max} 
\begin{cases} 

w_n + opt(p(n))\ (x_n\ \in \mathcal{O}) \\
opt(n-1)\ (x_n\ \notin \mathcal{O})


\end{cases} 

$$


























#### Why:
#### How:









