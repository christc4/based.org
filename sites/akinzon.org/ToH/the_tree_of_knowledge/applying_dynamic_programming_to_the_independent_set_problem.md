# Applying Dynamic Programming to The Independent Set Problem

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Independent Set Problem Background](independent_set_problem_background)

If we considered how to solve this problem for a chain-like [Graphs|graph](graphs|graph) (basically a [Linked Lists|linked list](linked_lists|linked_list)), the steps to the solution are something like this:

Consider the last node in the graph $x_n$ , and a set $\mathcal{O}$ that contained all the nodes to form the optimal solution.

If the last node was in $\mathcal{O}$, then the next node to consider would be $x_{n-2}$. 

If not, then we consider $x_{n-1}$ and the solution is essentially the same as if $x_n$ wasn't in the graph at all.

Formally describing it as a formula, $opt(n)$: 
$$ opt(n) = \text{max} 
\begin{cases} 

w_n + opt(n-2)\ (x_n\ \in \mathcal{O}) \\
opt(n-1)\ (x_n\ \notin \mathcal{O})


\end{cases} 

$$

From these two cases, we choose the case that results in the larger value at each step, to give us the optimal value.

given base cases of $opt(0) = 0$ and $opt(1) = 1$, the solution will reach the base case and use [memoisation](memoisation) to remember the results to each subproblem to use in order to calculate the final solution.






























#### Why:
#### How:









