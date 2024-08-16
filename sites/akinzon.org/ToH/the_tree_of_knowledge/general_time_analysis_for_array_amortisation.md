# General Time Analysis for Array Amortisation

Up: [Amortisation of Arrays](amortisation_of_arrays)
Brother(s):
TARGET DECK

If we perform $N$ adds on an array with initial capacity $n$, we perform $m$ copies where:

$$ n \times 2^{m-1} \lt N \le n \times 2^m $$

You can see this in the example, 16 is indeed greater than 8 and $\le$ 16.

m is therefore:

$$ m = 
\left\lceil log_2\left(\frac{N}{n}\right) \right\rceil $$
And the number of elements copied is:
$$ n + 2n + 4n + ... + 2^{m-1}n = n(1+2+...+2^{m-1})= n(2^m-1) $$
Using $\lceil x \rceil \lt x+1$, the total number of operations is:
$$ N + n(2^m - 1) + m = N + n2^{\left\lceil log_2\left(\frac{N}{n}\right) \right\rceil} - n + \left\lceil log_2\left(\frac{N}{n}\right) \right\rceil$$
Which gives us
$$ < N + 2N - n\ log_2\left(\frac{N}{n}\right) + 1 < 4N $$
and an amortised cost of adding elements is $\Theta (1)$ !































#### Why:
#### How:









