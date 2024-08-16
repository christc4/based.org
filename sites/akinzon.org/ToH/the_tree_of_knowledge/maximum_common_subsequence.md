# Maximum Common Subsequence

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Maximum Common Subsequence Problem Background](maximum_common_subsequence_problem_background)

The maximum common subsequence problem is solved by dynamic programming by considering the last character in both strings:
	- If they are the same, the problem reduces to the subproblem that occurs when the character isn't even in the strings ($X_{n-1}$ and $Y_{m-1}$)
	- Else, the answer is the larger out of the subproblem $MCS(X_{n-1}, Y_n)$ and $MCS(X_n, Y_{m-1})$.


**Maximum Common Subsequence Recurrence:**

$$ opt(X_n, Y_m) = \begin{cases} 

1 + opt(X_{n-1}, Y_{m-1}) \ (A[n] \in \mathcal{O}) \\
max(opt(X_n, Y_{m-1}), opt(X_{n-1}, Y_m)) \ (A[n] \notin \mathcal{O})


\end{cases} $$































#### Why:
#### How:









