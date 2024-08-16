# Dynamic Connectivity

Up: [Union-Find](union-find)
Brother(s):
TARGET DECK

Dynamic connectivity is the state of having a set of $N$ objects with a couple of operations:

`union`: connect two objects

`isConnected`: check if two objects are connected via a *path* (does not need to be a direct connection).

This connectivity can change over time, hence it being called dynamic.

Connectivity here is an example of an [Equivalence Relations|equivalence relation](equivalence_relations|equivalence_relation)!
	**Reflexivity:** objects are connected to themselves
	**Symmetry:** if object $A$ is connected to $B$, then $B$ is connected to $A$.
	**Transitivity:** if object $A$ is connected to $B$, and $B$ is connected to $C$, then object $A$ is connected to $C$ as $A,B,C$ all form a path.

All objects that form a path are an [Equivalence Classes|equivalence class](equivalence_classes|equivalence_class).






























#### Why:
#### How:









