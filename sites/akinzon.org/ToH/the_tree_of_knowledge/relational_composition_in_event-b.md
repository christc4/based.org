# Relational Composition in Event-B

Up: [Relations in Event-B](relations_in_event-b)
Related Note(s): [Relation Composition](relation_composition)
TARGET DECK

You can chain 2+ [relations](relations) together using **Relational Composition.**

Given $Q \in S \leftrightarrow T$ and $R \in T \leftrightarrow U$, The composition of $Q$ and $R$ is written as: 

$$ Q;R $$
It is defined as:

$$ \exists y · x \mapsto y \in Q \land y \mapsto z \in R $$
This means that there is a $y$ that is mapped to by $x$ in the relation $Q$, AND there is a mapping of $y \mapsto z$ in $R$, such that $x \mapsto z$ is in $Q;R$.



































#### Why:
#### How:
























































