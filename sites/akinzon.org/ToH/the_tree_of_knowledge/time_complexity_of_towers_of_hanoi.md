# Time Complexity of Towers of Hanoi

Up: [Time Complexity of Recursive Algorithms](time_complexity_of_recursive_algorithms)
Brother(s):
TARGET DECK

By counting operations, we get the [Counting via Generating Function|generating function](counting_via_generating_function|generating_function):
$$ T(n) = \begin{cases} 0,\ n=0 \\ 2T(n-1) + 1 \end{cases} $$
Solving this gives the solution $2^n-1$, giving $O(2^n)$.




$$ \frac{1}{(1-x)^2} = \biggl(\frac{1}{1-x}\biggr)^2 $$






























#### Why:
#### How:









