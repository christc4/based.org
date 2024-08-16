# JavaFX Questions

Up: [JavaFX](javafx)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::JavaFX

What are [JavaFX Stages](javafx_stages)? #flashcard 
Stages are top-level containers in which everything is held.
They act as windows for the GUI.
<!--ID: 1715689329892-->




What are [JavaFX Scenes](javafx_scenes)? What is the limitation for stages and scenes? #flashcard 
Scenes hold the contents for a window, defining what shows on the window.
Each stage can hold one scene at a time (but the scene itself can be changed!)
<!--ID: 1715689329923-->





How does [JavaFX var|var](javafx_var|var) work in [JavaFX](javafx)? #flashcard  
Var infers the type of a variable based on the right hand side of the assignment statement

What are the rules for [JavaFX var|var](javafx_var|var)? #flashcard 
Var variables cannot must be initialised once declared
Var variables can only be used locally
Once a var's type has been determined, it cannot be changed.
<!--ID: 1715689329934-->


What is [JavaFX Binding](javafx_binding)? How does it work in Java? #flashcard 
Binding is a way of synchronising the values of two components/variables.
In Java, [JavaFX Properties|properties](javafx_properties|properties) are used to represent bindable variables to facilitate binding.
<!--ID: 1715689329947-->


Is [JavaFX Binding](javafx_binding) only one-way? #flashcard 
No, binding can be two-way as well as one-way.
<!--ID: 1715689329959-->

What are [Events](events) in Event-Driven Programming? What do they contain? #flashcard 
Events are [objects](objects) which model user actions.
They contain information about the event, such as what happened.
<!--ID: 1715690404004-->


What are [Event Listeners](event_listeners) and [Event Handlers|handlers](event_handlers|handlers)? #flashcard 
Listeners are objects that "listen" for events.
Handlers are special listeners that perform an action once an event occurs.
<!--ID: 1715690404015-->


Can [events](events) only have one [Event Listeners|listener](event_listeners|listener)? What does this mean? #flashcard 
Events can have more than one listener. This means that multiple actions will happen once the listened-for event happens.
<!--ID: 1715690404027-->


What are the 2 [benefits of nested classes](benefits_of_nested_classes)? #flashcard 
[Nested classes](nested_classes) provide a logical grouping for [classes](classes), which makes it easier for developers to understand the relations between classes. This in turn makes the classes more readable and maintainable.
Nested classes also allow for access control and access restriction, which increases information hiding and [encapsulation](encapsulation). 
<!--ID: 1715690404036-->


What are [functional interfaces](functional_interfaces)? What do they allow us to do? #flashcard 
Functional interfaces are [interfaces](interfaces) with only one abstract method.
Functional interfaces allow us to pass [C Functions and C Pointers|function pointers](c_functions_and_c_pointers|function_pointers) to other Java [methods](methods).
<!--ID: 1715690404045-->


What are [Functors](functors)? What are [Lambdas](lambdas)? What is [the method reference operator](the_method_reference_operator) used for? #flashcard 
Functors are instances of [classes](classes) that implement [functional interfaces](functional_interfaces).
[Lambdas](lambdas) are example of functors, they are shorthand method references.
The method reference operator is used to treat [objects](objects) like [methods](methods) and pass an object's methods like a [C Functions and C Pointers|function pointer](c_functions_and_c_pointers|function_pointer).
<!--ID: 1715690404056-->



























#### Why:
#### How:









