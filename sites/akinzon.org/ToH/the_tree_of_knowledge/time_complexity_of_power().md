# Time Complexity of power()

Up: [Time Complexity of Recursive Algorithms](time_complexity_of_recursive_algorithms)
Brother(s):
TARGET DECK

Remembering that: 

$$ x^{2n} = (x^n)^2,\ x^{2n+1} = x \times x^{2n} $$

We can infer that:

$$ T(n) = \begin{cases} 0\ (n=1) \\ T(n/2) + 1\ (n\ is\ even) \\
T(\lfloor n/2 \rfloor) + 2\ (n\ is\ odd)
\end{cases} $$
We can then create the inequality:
$$ T(n) \le T(\lfloor n/2 \rfloor ) + 2 $$
Solving this recursive problem requires considering how many times $n$ must be divided to reach 1.

We can use binary and right shifts which represent division by 2 to help us with this.

if we wrote $n$ as $b_mb_{m-1}...b_2b_1$, where:
	$b_i \in \{0,1\}$
	$b_m = 1$
	$m$ is the number of digits in $n$'s binary representation
	$\lfloor n/2 \rfloor$ is one right shift of $n$, $b_mb_{m-1}...b_2$  
	$m-1$ right shifts on $n$ = 1
The time taken is therefore the same is the time taken for those $m-1$ divides, creating the inequality

$$ T(n) \le 2(m-1) $$

The time now depends on the size of m. With $m$ binary digits we can represent the range

$$ 2^{m-1} \le n < 2^m  $$
taking logs
$$ m - 1 \le log_2(n) < m $$
And so:

$$ T(n) \le 2 log_2(n) $$
Therefore:
$$ T(n)\ is\ O(log(n)) $$


























#### Why:
#### How:









