# Branch-And-Bound Subset Sum

Up: [Branch-And-Bound](branch-and-bound)
Related Note(s): [Subset Sum Problem Background](subset_sum_problem_background)
TARGET DECK

Given the subset sum problem, the method for solving it using branch-and-bound is as follows:
	- If the current sum $S$ + the next element $x$ = target $t$ $\implies$ the solution has been found!
	- If the current sum $S$ + the next element $x$ $\gt$ target $t$ $\implies$ we've gone too far, backtrack. (upper bound)
	- If the current sum $S$ + sum of remaining values $R$ - next element $x$ (current sum + other remaining values when not choosing $x$) $\lt$ target $t$ $\implies$ we'll **never** reach the target, backtrack (lower bound)



































#### Why:
#### How:









