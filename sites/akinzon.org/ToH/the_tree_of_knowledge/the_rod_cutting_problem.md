# The Rod Cutting Problem

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Rod Cutting Problem Background](rod_cutting_problem_background)

The Rod Cutting Algorithm works by enumerating all possible rod cuts, for a rod of length $n$, and chooses the most money-making cut.

**Rod Cutting Recurrence:**

$$

opt(n, p) = max_{i = 1}^n(p[i] + opt(n-i, p))










$$


[Memoisation](memoisation) is crucial to this algorithm's running time, since it prevents it from having to recompute previously computed values.






























#### Why:
#### How:









