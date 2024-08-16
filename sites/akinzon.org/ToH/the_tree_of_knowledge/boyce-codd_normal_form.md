# Boyce-Codd Normal Form

Up: [The Normal Forms](the_normal_forms)
Brother(s):
TARGET DECK

Boyce-Codd is a stricter version of [Third Normal Form](third_normal_form) that states that every non-key attribute in a table must be a "fact" about the key, such that no trivial [functional dependencies](functional_dependencies) remain in the table

Another way of saying this is that every functional dependency in a table **must** be a [Superkeys and Candidate Keys|candidate key](superkeys_and_candidate_keys|candidate_key).

If one attribute determines another yet it is not a candidate key, the [Relation Schemas|relation schema](relation_schemas|relation_schema) is **NOT** in BCNF.

[Going From 3NF to BCNF](going_from_3nf_to_bcnf)


































#### Why:
#### How:









