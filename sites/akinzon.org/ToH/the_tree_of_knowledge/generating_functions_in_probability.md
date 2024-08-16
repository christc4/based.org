# Generating Functions in Probability

Up: [Combinatorics in Probability](combinatorics_in_probability)
Brother(s): [Counting via Generating Function](counting_via_generating_function)
TARGET DECK

> [!TIP]- 2d6: toss a six-sided die twice, what is the probability that the sum is 4 ?
> $Ω(6-sided\ die) = \{1, 2, 3, 4, 5, 6\}$ with probabilities $\{p_1, p_2, p_3, p_4, p_5, p_6\}$
> $G_{1d6}(x) = p_1x + p_2x^2 + p_3x^3 + p_4x^4 + p_5x^5 + p_6x^6$
> $G_{2d6}(x) = $G_{1d6}(x) · G_{1d6}(x)$
> $\mathbb{P}(sum\ is\ 4) = [x^4]G_{2d6}=p_1p_3 + p_2p_2 + p_3p_1$ (3 possible ways to get 4 from polynomial)
> If the die is fair $p_i=\frac{1}{6}$, so the probability is $\frac{3}{36}=\frac{1}{12}$  


If, for example, one of the die had 4 sides instead of 6, its generating function would resemble:
$$ (x^1 + x^2 + x^3 + x^4 ) $$


































#### Why:
#### How:









