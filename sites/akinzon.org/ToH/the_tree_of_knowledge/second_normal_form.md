# Second Normal Form

Up: [The Normal Forms](the_normal_forms)
Brother(s): [First Normal Form](first_normal_form)
TARGET DECK

The Second (Hokage) Normal Form removes partial dependencies from a [Relation Schemas|relation schema](relation_schemas|relation_schema).

Partial dependencies arise when an attribute is [Functional Dependencies|functionally dependent](functional_dependencies|functionally_dependent) on the [Superkeys and Candidate Keys|primary key](superkeys_and_candidate_keys|primary_key), but only part of it.

This is a good sign that the database can be broken down further into a part where data that depends on only part of the primary key can be in its own table.

If you can remove any part of a [Superkeys and Candidate Keys|candidate key](superkeys_and_candidate_keys|candidate_key) (forming a proper subset of the candidate key) and it is still a candidate key, you've got a partial dependency.

































#### Why:
#### How:









