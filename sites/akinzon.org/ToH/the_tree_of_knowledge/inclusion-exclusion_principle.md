# Inclusion-Exclusion Principle

Up: [Combinatorics](combinatorics)
Brother(s): [Inclusion-Exclusion Principle (Probability Version)](inclusion-exclusion_principle_(probability_version))
TARGET DECK

![Pasted image 20240102202145.png](pasted_image_20240102202145.png)

The cardinality of the union of two [sets](sets) is the sum of the cardinalities of both sets minus the cardinality of their intersection (so that the same element isn't counted twice.)

Therefore, if the two sets are **disjoint,** the cardinality of the sum of their union is only the sum of their cardinalities.

AKA addition rule

> [!TIP]- How many integers n for $1 ≤ n ≤ 2001$ are divisible by 3 or 4 but not 5?? 
> Inclusion-exclusion principle: 
> $A$ = numbers divisible by 3
> $B$ = numbers divisible by 4
> $C$ = numbers not divisible by 5
> $|(A \cap C) \cup (B \cap C)| = |(A \cap C)| + |(B \cap C)| - |(A \cap C) \cap (B \cap C)|$
> $|(A \cap C) \cup (B \cap C)| = |(A \cap C)| + |(B \cap C)| - |(A \cap B \cap C)|$
> $|(A \cap C) \cup (B \cap C)| = (\lfloor 2001/3 \rfloor - \lfloor 2001/15 \rfloor) + (\lfloor 2001/4 \rfloor - \lfloor 2001/20 \rfloor) - (\lfloor 2001/12 \rfloor - \lfloor 2001/60 \rfloor)$
> $|(A \cap C) \cup (B \cap C)| = 801$

































#### Why:
#### How:









