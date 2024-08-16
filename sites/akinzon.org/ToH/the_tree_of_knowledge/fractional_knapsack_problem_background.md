# Fractional Knapsack Problem Background

Up: [The Fractional Knapsack Algorithm](the_fractional_knapsack_algorithm)
Brother(s):
TARGET DECK

This problem is a bit like filling your bag with as many items as possible in IQ dungeon!

More formally, given $n$ items having value $v_i$ and (physical) weight $w_i$, and your knapsack of capacity $C$:

We want to maximise the value in our bag,
$$\sum_{i=1}^n x_i \cdot v_i$$

And we want to minimise the physical weight of our bag (we're not that bolo) whilst also making sure that our bag can carry everything we want to bring:
$$\sum_{i=1}^n x_i \cdot w_i \le C$$

($x_i$ is the fraction of the item we're going to carry, so of course it's between 0 and 1 inclusive.)


































#### Why:
#### How:









