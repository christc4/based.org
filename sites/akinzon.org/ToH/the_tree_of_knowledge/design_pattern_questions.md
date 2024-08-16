# Design Pattern Questions

Up: [Design Patterns](design_patterns)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1216 Software Modelling and Design::Design Patterns

What are the three [design pattern types](design_pattern_types)? (SCB) #flashcard 
**Structural:** Used to assemble [objects](objects) and [classes](classes) into larger structures
**Creational:** Used to create and instantiate objects
**Behavioural:** Governs the behaviour of objects and classes.
<!--ID: 1713898042216-->

What is meant by [Program to Interfaces, not Implementation](program_to_interfaces,_not_implementation)? #flashcard 
This is the idea that [classes](classes) should communicate with each other through their [methods|method](methods|method) [interfaces](interfaces) and not through implementation specifics.
<!--ID: 1713898372183-->


What is meant by [Encapsulate What Varies](encapsulate_what_varies)? #flashcard 
This is the idea that code that is susceptible to change should be wrapped in its own [Methods|method](methods|method)/[Classes|class](classes|class).
<!--ID: 1713898372204-->


What is the [Composite Design Pattern](composite_design_pattern)? #flashcard 
The Composite Design Pattern models a system as a complex object composed of simpler objects. Components in this system communicate through their [interfaces](interfaces).
<!--ID: 1713898498512-->


What is the [Adapter Design Pattern](adapter_design_pattern)? How does it work? #flashcard 
The Adapter Design Pattern acts as the intermediary for your codebase and an incompatible 3rd-party library.
The Adapter has an [interfaces|interface](interfaces|interface) that one [Stub Objects|object](stub_objects|object) can call, and converts the output of those calls into a format that the second object can understand.
<!--ID: 1713898601863-->

What is the [Bridge Design Pattern](bridge_design_pattern)? #flashcard 
The Bridge Design Pattern simplifies [inheritance](inheritance) hierarchies by decomposing the base [Classes|class](classes|class) into simple, more atomic classes each with their own inheritance hierarchy. These classes are then linked to each other via [Associations in UML|association](associations_in_uml|association).
<!--ID: 1713898680069-->


What is the [Decorator Design Pattern](decorator_design_pattern)? #flashcard 
The Decorator Design Pattern allows for the features of an [Objects|object](objects|object) to be customised by wrapping it in decorator objects that contain the [methods](methods) for the new feature rather than creating a new class to represent the original object with the new feature
<!--ID: 1713898845344-->

What is the [Observer Design Pattern](observer_design_pattern)? #flashcard 
At the centre of the Observer Design Pattern is an [Objects|object](objects|object) that broadcasts its changes in state/actions to its observers (listeners), which then perform an action based on what was broadcasted.
<!--ID: 1713898933798-->


How do [Factory Methods](factory_methods) work? What can be done to them? #flashcard 
Factory methods are [methods](methods) that create [objects](objects) for the owning class.
They can be [Overriding|overridden](overriding|overridden) to create subclasses of the base class specified in the base factory method. 
<!--ID: 1713899515899-->

























#### Why:
#### How:









