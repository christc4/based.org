# Event-B Questions

Up: [Event-B](event-b)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1216 Software Modelling and Design::Event-B




What is the [Benefits of Types in Event-B|benefit of assigning types](benefits_of_types_in_event-b|benefit_of_assigning_types) in [Event-B](event-b)? #flashcard 
Giving everything a type in Event-B makes our model more clear.
It also prevents us from making errors since the computer can type check and ensure that our statements make sense
<!--ID: 1709756834302-->




When we say a [Sets|set](sets|set) $S \subseteq T$, what type is that set? #flashcard 
It is of the type $\mathbb{P}(T)$, since all subsets of a set are members of the powerset.
<!--ID: 1709756834310-->




When we say $s \in X \times Y$, what are we saying? #flashcard 
This means that $s$ is an [Ordered Pairs|ordered pair](ordered_pairs|ordered_pair) who is a member of the Cartesian Product of $X$ and $Y$, and so $X \times Y$ is $s$' type.
<!--ID: 1709756834316-->




What is the [Relational Image in Event-B|relational image](relational_image_in_event-b|relational_image) and [Relational Inverse in Event-B|relational inverse](relational_inverse_in_event-b|relational_inverse) analogous to, respectively? #flashcard 
Relational image is analogous to returning values associated with keys
Relational inverse is analogous to returning keys that map to the specified values.
<!--ID: 1709756834324-->





What is the difference between [Partial Functions in Event-B|partial functions](partial_functions_in_event-b|partial_functions) and [Total Functions in Event-B|total functions](total_functions_in_event-b|total_functions)? #flashcard 
Partial functions do not necessarily map all of the domain to a value in the range, but total functions do.
<!--ID: 1709756834330-->




Is every inverse [Relations in Event-B|relation](relations_in_event-b|relation) a [Functions in Event-B|function](functions_in_event-b|function)? #flashcard 
No, because some range values may have been mapped to by multiple domain values, and so the inverse cannot be a a function because one value cannot map to many values in a function.
<!--ID: 1709756834337-->


What is [Extension Refinement](extension_refinement) in Event-B? #flashcard 
Extension Refinement is when more functionality and detail is added to the model, such as new variables and invariants, and extending/creating events to use new variables.
<!--ID: 1711101648192-->


What is the guarantee of [Refinement in Event-B](refinement_in_event-b)? #flashcard 
In Event-B, refinement preserves safety: we can prove that invariants are kept even after refinement.
<!--ID: 1711101743566-->


If you were to [Modelling Lists in Event-B|model a list in Event-B](modelling_lists_in_event-b|model_a_list_in_event-b), what would be an important component, and why? #flashcard 
An important component would be to add a [Total Functions in Event-B|total](total_functions_in_event-b|total), [Functions (Advanced)|injective](functions_(advanced)|injective) function mapping of each item to its position in the list.
It is important that it is a total function, since each item needs a position.
It is also important that it is an injective function, since no two items are allowed to have the same position.
<!--ID: 1711101992925-->

How can [Modelling Queues in Event-B|queues](modelling_queues_in_event-b|queues) and [Modelling Stacks in Event-B|stacks](modelling_stacks_in_event-b|stacks) be modelled in Event-B by extending the method for [Modelling Lists in Event-B|modelling lists](modelling_lists_in_event-b|modelling_lists)? #flashcard 
For queues: item to be enqueued has largest position, item to be dequeued has smallest position.
For stacks: item that is pushed has largest position, and item that is popped has largest position too.
<!--ID: 1711102156975-->


If you were [managing multiple collections in Event-B](managing_multiple_collections_in_event-b), what would be an important thing to consider? #flashcard 
An important thing to consider is that [Event-B Invariants|invariants](event-b_invariants|invariants) between two elements of the same collection are **explicit in specifying that they are in the same collection**, otherwise the invariants would be too strong/restrictive.
<!--ID: 1711102305489-->

What are two types of [validation in Event-B](validation_in_event-b)? #flashcard 
Requirements validation: How well requirements meet stakeholder needs
Model validation: How well the model aligns with the requirements.
<!--ID: 1711103236382-->

What are the three types of [verification in Event-B](verification_in_event-b)? #flashcard 
Model verification: A model's correctness/consistency (How well it maintains invariants and refines other models)
Formal Verification: Verifying a model's correctness mathematically given some anticipated properties
Code Verification: How well a program implements a model
<!--ID: 1711103348625-->


What are [proof obligations in Event-B](proof_obligations_in_event-b), and how are they used? #flashcard 
Proof obligations are mathematical theorems that are **derived** from a formal model (as in, you're obliged to proved these things.)
They can be used to verify model properties such as if a model preserves invariants.
<!--ID: 1711103459568-->


What is the main idea behind [Inductive Invariants](inductive_invariants)? #flashcard 
This is the idea that if an invariant holds before an event, then given a event with guards and variable assignment, the invariant should still hold.
<!--ID: 1711103517466-->

What are [Sequents in Event-B](sequents_in_event-b)? #flashcard 
Sequents consist of hypotheses and goals, in the form $H ⊢ G$. Sequents are valid if $G$ follows from $H$.
<!--ID: 1711103586873-->

What is [Substitution in Event-B](substitution_in_event-b)? #flashcard 
Substitution in Event-B is when free (not bounded by a quantifier) occurrences of a variable are replaced by a concrete value.
Basically variable assignment.
<!--ID: 1711103684242-->








#### Why:
#### How:









