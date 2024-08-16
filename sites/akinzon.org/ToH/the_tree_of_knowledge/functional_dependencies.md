# Functional Dependencies

Up: [Keys in Database Systems](keys_in_database_systems)
Brother(s):
TARGET DECK

Given a relation with [Relation Schemas|schema](relation_schemas|schema) $R(A_1,...,A_n,B_1,...,B_m)$, and an instance of $R$ called $r$,

$r$ satisfies the **functional dependency** $$ A_1,...,A_n \rightarrow B_1,...,B_m $$
if wherever two values agree on the values for $A_1,...,A_n$, they also agree on the values for $B_1,...,B_m$.

This means that if two tuples have the same values for $A_1,...,A_n$, they **must be the same tuple**, and have the same values for $B_1,...,B_m$.

We describe this by saying $A_1,...,A_n$ **functionally determines** $B_1,...,B_m$.

$A_1,...,B_n$ is the **determinant** and $B_1,...,B_m$ is the **dependent set.**

If we generalise both as being sets $A$ and $B$ respectively, we could write it as: $$ A \rightarrow B $$

[Functional Dependency Rules](functional_dependency_rules)

[Finding out if a Functional Dependency Holds for a Schema](finding_out_if_a_functional_dependency_holds_for_a_schema)

[Minimal Set of Functional Dependencies](minimal_set_of_functional_dependencies)






















#### Why:
#### How:









