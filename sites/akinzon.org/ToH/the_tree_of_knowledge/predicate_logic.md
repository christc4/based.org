# Predicate Logic

Up: [Logic](logic)
Brother(s):
TARGET DECK

**Predicate Logic** is a much more richer and expressive language compared to [Propositional Logic](propositional_logic).

In **Propositional logic,** a world is an assignment of truth values to propositional variables.

In **Predicate logic,** a world is a set, together with relations and [Functions|functions](functions|functions).

Given a set A:
	- An n-ary function is a function of type $A^n → A$
	- An n-ary relation is a subset of $A^n$

#### Representing statements in predicate logic

Take the definition of an injective function:
	"A function is injective only when for all x and y in A, if f(x) = f(y) then x = y"

As a predicate logic statement, this would be:
	∀x. ∀y. (f(x) = f(y)) → (x=y)

Predicate logic worlds model/satisfy the formula only when f is an injective function in that world

[Backus-Naur Form](backus-naur_form) for Predicate Logic:
	T (terms) ::= V (variables) | C (constants) | f(T, T,...,T)
	F (formulae) ::= R( T, T,...,T) | T = T | ⊥ | ¬F | F ∧ F | F ∨ F | F → F | ∃x. F | ∀x. F

[Predicate Semantics](predicate_semantics)






































#### Why:
#### How:









