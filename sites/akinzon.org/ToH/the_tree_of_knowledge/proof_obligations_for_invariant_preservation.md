# Proof Obligations for Invariant Preservation

Up: [Proof Obligations in Event-B](proof_obligations_in_event-b)
Brother(s):
TARGET DECK

The proof obligations for invariant preservation use the idea of [inductive invariants](inductive_invariants).

Given variables $v$, an invariant $Inv$ (where $v$ is not quantified by $Inv$), and an event:

$$ Ev ≙ \text{any } x \text{ where } Grd \text{ then } v := Exp$$

The proof obligation is to prove this [Sequents in Event-B|sequent](sequents_in_event-b|sequent):

$$ Inv,\ Grd ⊢\ Inv[v:=Exp] $$

This is proving that an updated invariant holds for the previous invariant and a guard, for a given event.






























#### Why:
#### How:









