# The Closure Algorithm

Up: [Closures in Database Systems](closures_in_database_systems)
Brother(s):
TARGET DECK

**Input:** [Relation Schemas|Relation schema](relation_schemas|relation_schema) $R$ with [Sets|set](sets|set) of attributes $U$, set $F$ of [Functional Dependencies|FDs](functional_dependencies|fds) and a set of attributes $X \subseteq U$

**Output:** $X^+ = \{A_i: F ⊧ X \rightarrow A_i \}$

**Initialisation:** $X_0 = X$

**Repeat:**
	$$X_{n+1} := X_n \cup \{A_i|\exists Y \rightarrow Z \in F | Y \subseteq X_n \land A_i \in Z\}$$
**Until:**
$$ X_{n+1} == X_n$$
(no more changes made)




























#### Why:
#### How:









