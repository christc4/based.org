# Tips and Tricks for [Sequence Diagrams](sequence_diagrams)

Up: [Making Good UML Models](making_good_uml_models)
Brother(s):
TARGET DECK

1st column should be the main [System Actors|actor](system_actors|actor) for the [Use Cases|use case](use_cases|use_case)

2nd column should be a **[Elements in a Sequence Diagram|boundary object](elements_in_a_sequence_diagram|boundary_object)** that the actor must interact with to gain access to the system (like a login page)

3rd column should be a **control object** that manages the rest of the use case

(Can you visualise the flow? Don't memorise this, **visualise!!!**)

Create control objects at the beginning of an event, which then create the boundary objects.

Control and boundary objects should be able to access entities, but not the other way around.

(like if you want login info to be accessed by the login page, but not the other way around because that's beetlejuice behaviour)




































#### Why:
#### How:









