# Event-B Events

Up: [The Event-B Machine](the_event-b_machine)
Brother(s):
TARGET DECK


Events in Event-B represent how the [The Event-B Machine|machine](the_event-b_machine|machine) can behave, with two things:
	**Guards:** Conditions which must be met for the behaviour to happen
	**Actions:** The actual way in which machine state is changed

> [!tip] These guards and actions are written using [Set Operations|set notation!](set_operations|set_notation!)

There is a special kind of event called *Initialisation*, which has no guards (unconditional) and happens once in a machine before anything else.

During initialisation, it initialises the machine's [variables](variables) to values that are consistent with the machine's [Event-B Invariants|invariants](event-b_invariants|invariants).



































#### Why:
#### How:









