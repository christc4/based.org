# Time Complexity of the Factorial

Up: [Time Complexity of Recursive Algorithms](time_complexity_of_recursive_algorithms)
Brother(s):
TARGET DECK

We can compute this by counting the multiplications:

$$ T(n) = T(n-1) + 1, \forall n > 0 $$
$$ T(0) = 0 $$
Which gives us the [Counting via Generating Function|generating function](counting_via_generating_function|generating_function) of:
$$ t_n = \begin{cases} 
0, n = 0 \\
t_{n-1} + 1, n >0
\end{cases} $$
Which we solve to get $n$.

Or alternatively,

![Pasted image 20240207175449.png](pasted_image_20240207175449.png)
































#### Why:
#### How:









