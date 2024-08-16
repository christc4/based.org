# Branch-And-Bound 0-1 Knapsack

Up: [Branch-And-Bound](branch-and-bound)
Related Note(s): [Knapsack Problem Background](knapsack_problem_background)
TARGET DECK

For this problem, we obtain the upper bound using the [Greedy Algorithms|greedy](greedy_algorithms|greedy) solution to the [The Fractional Knapsack Algorithm|fractional knapsack](the_fractional_knapsack_algorithm|fractional_knapsack) problem.

The upper bound at each branch is calculated using this formula:
$$ S \le C \cdot \frac{v_1}{w_1} $$
Once we branch enough to find a high enough solution, any other branch where the calculated upper bound is less than our current maximum is discarded.


































#### Why:
#### How:









