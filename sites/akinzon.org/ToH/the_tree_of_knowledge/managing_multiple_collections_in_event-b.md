# Managing Multiple Collections in Event-B

Up: [Representing Collections in Event-B](representing_collections_in_event-b)
Brother(s):
TARGET DECK

With our [Total Functions in Event-B|total](total_functions_in_event-b|total) [Functions (Advanced)|injective](functions_(advanced)|injective) function that we use to map items in collections to their positions, there is a limitation of this approach when it comes to managing multiple collections.

The [Event-B Invariants|invariant](event-b_invariants|invariant) that no two elements can have the same position is **too strong,** since it does not specify that both elements must be in the same collection for the comparison to mean anything.

With a bit of [Predicate Logic|First-Order Logic](predicate_logic|first-order_logic), we can fix this though!

If we specify that:

For all items (in the same [Carrier Sets|carrier set](carrier_sets|carrier_set)) in the same collection as our item before any sort of comparison, we can efficiently manage multiple collections at the same time!

This is how we express that with logic:

$$ 
\begin{align}
\forall x, y\cdot x \in \text{carrierSet}\ \land y \in \text{carrierSet}\\
\land\ x \ne y\\
\land\ \text{collection}(x) = \text{collection}(y)\\
\implies\ \text{position}(x) \ne \text{position}(y)
\end{align}
$$





































#### Why:
#### How:









