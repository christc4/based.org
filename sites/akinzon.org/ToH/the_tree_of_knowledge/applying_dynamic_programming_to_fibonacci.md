# Applying Dynamic Programming to Fibonacci

Up: [Dynamic Programming](dynamic_programming)
Brother(s):
TARGET DECK

The classic [Recursion|recursive](recursion|recursive) algorithm for computing Fibonacci numbers is very wasteful, since it calculates the same Fibonacci number multiple times!

Surely there is a better way to go about things...

[Memoisation](memoisation) is the key!

By memoising the results of previous subproblems, we only have to calculate each Fibonacci number exactly once, leading to a much greater time complexity.

A [Bottom-Up Solutions|bottom-up](bottom-up_solutions|bottom-up) approach would use memoisation by starting at 0, and calculating each Fibonacci number on the way to the target number $n$, while using the previous results, $n-1$ and $n-2$ to calculate $n$.































#### Why:
#### How:









