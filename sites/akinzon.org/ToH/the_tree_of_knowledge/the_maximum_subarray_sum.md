# The Maximum Subarray Sum

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

[Maximum Subarray Sum Background](maximum_subarray_sum_background)

The maximum subarray sum works by calculating $max_i\ S_{ij}$ for all $j$, and then choosing the largest of those to get $max_{ij}\ S_{ij}$.

It [Iteration|iterates](iteration|iterates) through the array, finding the maximum sum for each subarray.

With all maximum sums computed, it then returns the maximum of all the sums

At the most basic level, the maximum sum for a subarray could be either be $A[i]$ or $A[i] + M[i-1]$, depending on which one is bigger.
	- It would be $A[i]$ when $M[i-1]$ is negative or 0
	- It would be $A[i] + M[i-1]$ otherwise, since even 1000000000001 is greater than 1000000000000, even if only by one.

**Maximum Subarray Sum Recurrence:** 
$$ M_{j+1} = max(M_j + A[j+1], A[j+1]) $$


**Akin Style:**

$$ M[j] = max \begin{cases} \text{(M[j-1] + A[j] > A[j])}\ \ \ \ \ M[j-1] + A[j]\ \text{(increment end of maximum subarray)}\\ \text{(else)}\ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \ \  A[j]\ \text{(increment start of maximum subarray)} \end{cases}  $$

We disregard $M_j$ if it will only detract from the maximum subarray sum (by being negative OR negative and smaller than a negative $A[j+1]$.)

If $M_j$ is no longer part of the maximum sum, then $i$ isn't where $M_j$ would have it be, and it is now equal to $j$.

This solution also solves all subproblems!
























#### Why:
#### How:









