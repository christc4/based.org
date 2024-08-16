# Substitution in Event-B

Up: [Proof-Based Verification in Event-B](proof-based_verification_in_event-b)
Brother(s):
TARGET DECK

Substitution in Event-B is when *free* occurrences of a variable $x$ are replaced by expression $E$ in the predicate $P$.

Example:

$$ (0 \lt n\ \land\ n \le10 )[n:= 7] \iff (0 \lt 7\ \land\ 7 \le10 ) $$

**Bound variables cannot be substituted**, with bound variables being those that have been quantified in the predicate (using the universal quantifier $\forall$ or existential quantifier $\exists$)

[Multiple Substitution in Event-B](multiple_substitution_in_event-b)































#### Why:
#### How:









