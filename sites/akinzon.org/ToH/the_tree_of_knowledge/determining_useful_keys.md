# Determining Useful Keys

Up: [Keys in Database Systems](keys_in_database_systems)
Brother(s):
TARGET DECK

To find all [Superkeys and Candidate Keys|candidate keys](superkeys_and_candidate_keys|candidate_keys), we'll first need to find all [Superkeys and Candidate Keys|superkeys](superkeys_and_candidate_keys|superkeys).

To find all superkeys, we'll first need to check if a set of attributes chosen at random is a superkey.

This is to say that given a relation schema $R$, a set $F$ of [Functional Dependencies|FDs](functional_dependencies|fds) on $R$, and a FD $X \rightarrow Y$, we will be able to determine if $$F ⊧ X \rightarrow Y$$


This question can be solved with [Closures in Database Systems|closures](closures_in_database_systems|closures)!

Another important part of determining useful keys is considering how the data may change in the future ([functional dependencies](functional_dependencies) that once were no longer are)

































#### Why:
#### How:









