# Subset Sum Algorithm

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Subset Sum Problem Background](subset_sum_problem_background)

The subset sum algorithm checks each element in the array.

If it is in the solution, the problem is reduced to the subproblem of checking if from the rest of $A$ a subset can be constructed whose sum is $S - A[n]$.

If it is not, the problem is reduced to the subproblem of checking if from the rest of $A$ a subset can be constructed whose sum is $S$.

**Subset Sum Recurrence:**

$$ opt(n, A, S) =

\begin{cases}

A[n] + opt(n-1, A, S-A[n]) \ (A[n] \in \mathcal{O})\\

opt(n-1, A, S) \ (A[n] \notin \mathcal{O})



\end{cases}

$$
































#### Why:
#### How:









