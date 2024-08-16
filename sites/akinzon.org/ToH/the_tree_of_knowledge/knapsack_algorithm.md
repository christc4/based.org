# Knapsack Algorithm

Up: [Dynamic Programming](dynamic_programming)
Related Note(s): [The Fractional Knapsack Algorithm](the_fractional_knapsack_algorithm)
TARGET DECK

[Knapsack Problem Background](knapsack_problem_background)

The 0-1 Knapsack algorithm starts off in a similar way to the fractional knapsack algorithm by considering items with the largest value : weight ratio.

If it can select an item, it will do so, and continue with the subproblem of selecting items for the remaining capacity.

If it cannot, it continues with the subproblem of selecting the remaining items.

The final value is the largest of the two selections.



**Knapsack Recurrence (Consider an array of items $A$, capacity $C$, and index $n$):**

$$ opt(n, A, C) = max 
\begin{cases}

v_{A[n]} + opt(n-1, A, C-w_{A[n]})\ (A[n] \in \mathcal{O}) \\

opt (n-1, A, C)\ (A[n] \notin \mathcal{O})


\end{cases}

$$
































#### Why:
#### How:









