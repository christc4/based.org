# Using Recursion To Compute Powers

Up: [Recursion](recursion)
Brother(s):
TARGET DECK

Say we were computing $0.95^{25}$, we could iteratively multiply 0.95 by itself 24 times, but that is bad.

With recursion, we can slash the number of operations down, using nice maths:

$$x^{2n} = (x^n)^2,\ x^{2n+1} = x \times x^{2n}$$
Both of these rules allow us to break down the exponentiation until we reach a base case, namely $x^1 = x$.

for $0.95^{25}$:

$$ 0.95^{25} = 0.95 \times 0.95^{24} $$

Which gives us
$$ 0.95^{25} = 0.95 \times (0.95^{12})^2 $$
$$ 0.95^{25} = 0.95 \times ((0.95^{6})^2)^2 $$
$$ 0.95^{25} = 0.95 \times (((0.95^{3})^2)^2)^2 $$
$$  0.95^{25} = 0.95 \times (((0.95 \times 0.95^{2})^2)^2)^2  $$
$$  0.95^{25} = 0.95 \times (((0.95 \times (0.95)^{2})^2)^2)^2  $$
From 24 multiplications to 6!!!

[Implementation of power() using Recursion](implementation_of_power()_using_recursion)



























#### Why:
#### How:









