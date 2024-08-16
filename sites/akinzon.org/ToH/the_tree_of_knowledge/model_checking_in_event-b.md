# Model Checking in Event-B

Up: [Proof-Based Verification in Event-B](proof-based_verification_in_event-b)
Brother(s):
TARGET DECK

Event-B comes with **Pro-B** which is used to [Verification in Event-B|verify](verification_in_event-b|verify) models.

It checks all possible [States in Event-B|states](states_in_event-b|states) and stops when:
	- A state that violates an invariant (error state) is reached
	- All **reachable** states are explored, and no invariant violations occur.

Pro-B can also check for **deadlocks** which happen when a state is reached, and you cannot go any further since there is no option to go any further.



































#### Why:
#### How:









