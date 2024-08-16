# Testing Questions

Up: [Software Testing](software_testing)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::Testing

What is [Black-Box Testing](black-box_testing)? What is its main goal and who does the testing, and why? #flashcard 
Black-Box Testing is a form of testing where the tester has no knowledge of the underlying system.
The main goal is to test the system's input and output behaviours (if input `x` produces desired output `y`).
As such, it can be carried out by anyone since it requires no knowledge of the system.
<!--ID: 1713821313732-->

What is [White-Box Testing](white-box_testing)? What is its main goal and who does the testing, and why? #flashcard 
White-Box Testing is a form of testing where the tester has knowledge of and access to the underlying system.
The main goal is to test the system's [methods](methods) and [data structures](data_structures).
As such, it can only be carried out by people with good programming and system knowledge, such as testers and developers.
<!--ID: 1713821410746-->

What is a big risk of [White-Box Testing](white-box_testing), and why? #flashcard 
Since the tester is required to have access to the underlying system, a bad actor can steal the source code.
<!--ID: 1713821443147-->

What is a [Unit Testing|unit test](unit_testing|unit_test)? When can they be done, and what kind of box test is it? Who does them? #flashcard 
Unit tests are tests carried out on small pieces of code. 
They can be done anytime.
They are a form of [White-Box Testing](white-box_testing).
This is done by developers.
<!--ID: 1713821602466-->


What is an [Integration Testing|integration test](integration_testing|integration_test)? When can they be done, and what kind of box test is it? Who does them? #flashcard 
Integration tests carry out tests on increasing amounts of the system.
These are done after [unit testing](unit_testing), but before [system testing](system_testing).
These are a mixture of [White-Box Testing](white-box_testing) and [Black-Box Testing](black-box_testing).
These are done by testers and developers.
<!--ID: 1713821715192-->

What is a [System Testing|system test](system_testing|system_test)? When can they be done, and what kind of box test is it? Who does them? #flashcard 
System tests are tests done on the entire system.
These are done at the end of development.
These are mainly [Black-Box Testing|black-box tests](black-box_testing|black-box_tests)
Therefore, they can be done by anyone.
<!--ID: 1713821812759-->

What are [Mock and Stub Objects](mock_and_stub_objects)? What are they mainly used for? Who creates which? #flashcard 
[Mock objects](mock_objects) are [objects](objects) used to mimic the behaviour of real objects. Mock objects are mainly used to verify system behaviour. Mock objects are usually created by 3rd-party libraries.
[Stub Objects](stub_objects) are objects used for testing by having the minimal amount of [methods](methods) for the test to run. Stub objects are mainly used to verify system state. Stubs are usually created by developers.
<!--ID: 1713822025358-->


What is [Dependency Injection](dependency_injection)? What are three ways dependency injections can be done? #flashcard 
Dependency injection is a [Refactoring Techniques|refactoring technique](refactoring_techniques|refactoring_technique) done by using [stub objects](stub_objects) at system boundaries and extracting an interface from them to allow for injection. 
Dependency injections can be done through [Constructors|constructors](constructors|constructors), [Accessor Methods|accessor](accessor_methods|accessor)/[Mutator Methods|mutator](mutator_methods|mutator) methods and [factory methods](factory_methods).
<!--ID: 1713822223679-->

What's the difference between [testing to pass](testing_to_pass) and [testing to fail](testing_to_fail)? #flashcard 
Testing to pass is ensuring that the system does what the spec details it should do.
Testing to fail is using specific inputs to find bugs in the system
<!--ID: 1713892034534-->

What are [Drivers (Testing)|drivers](drivers_(testing)|drivers), in the context of [testing](testing)? What kind of module are they typically used on? #flashcard 
Drivers are the "[The Main Method|main](the_main_method|main) programs" that initiate and oversee the execution of an isolated component under test, passing it the necessary data to test it.
They are typically used to test [Low-Level Modules|lower-level modules](low-level_modules|lower-level_modules).
<!--ID: 1714070673872-->

What is [Top-Down Integration](top-down_integration), and [Top-Down Integration Evaluation|evaluate](top-down_integration_evaluation|evaluate) it #flashcard 
Top-Down integration is a method of [Integration Testing](integration_testing) where the [High-Level Modules](high-level_modules) are tested first, then the [Low-Level Modules|lower-level modules](low-level_modules|lower-level_modules) (which are [Stub Objects|stubs](stub_objects|stubs) where necessary) that it calls, so on and so forth until the lowest-level modules have been reached.
**Advantages:** Major flaws are found early, it's easy to locate flaws, and the [Drivers (Testing)|drivers](drivers_(testing)|drivers) are easier to write since top-down follows the logical structure of a program
**Disadvantages:** You need so many [mock and stub objects](mock_and_stub_objects) from the beginning and the lower-level modules aren't tested as 
rigorously as you'd like since they're not being tested in isolation
<!--ID: 1714071044437-->

What is [Bottom-Up Integration](bottom-up_integration), and [Bottom-Up Integration Evaluation|evaluate it](bottom-up_integration_evaluation|evaluate_it) #flashcard 
Bottom-Up integration is a method of [Integration Testing](integration_testing) where the [Low-Level Modules](low-level_modules) are tested first (and are called by [Drivers (Testing)|drivers](drivers_(testing)|drivers)), then the [High-Level Modules|higher-level modules](high-level_modules|higher-level_modules), so on until the highest-level modules have been reached.
**Advantages:** [Stub Objects|stubs](stub_objects|stubs) aren't used as much, low-level modules are well-tested, easy to locate faults (and early too)
**Disadvantages:** Needs drivers since the entry-points to these low-level modules aren't as clear as [Top-Down Integration](top-down_integration), and since we're working our way up not down, the logic of the overall solution is tested closer to the end.
<!--ID: 1714071244164-->

What is [Sandwich Integration](sandwich_integration), and [Sandwich Integration Evaluation|evaluate](sandwich_integration_evaluation|evaluate) it #flashcard 
Sandwich integration is a method of [Integration Testing](integration_testing) that combines [Top-Down Integration](top-down_integration) and [Bottom-Up Integration](bottom-up_integration).
**Advantages:** Same as top-down and bottom-up, + their disadvantages are mitigated/cancelled out
**Disadvantages:** It's rather complicated
<!--ID: 1714071355647-->

What is [Big Bang Integration](big_bang_integration), and [Big Bang Integration Evaluation|evaluate](big_bang_integration_evaluation|evaluate) it #flashcard 
Big bang integration is a method of [Integration Testing](integration_testing) where each component/module is tested in [Testing in Isolation|isolation](testing_in_isolation|isolation), then integrated.
**Advantages:** saves time, good for small systems
**Disadvantages:** Easy to miss faults and test cases, and if found, these faults are found late, can't be started  until all components are complete.
<!--ID: 1714071461568-->

What are the [four types of errors found in testing](four_types_of_errors_found_in_testing)? #flashcard 
System failure: Event where system does something it shouldn't do
System Error: Erroneous system state that leads to system failure
System fault: Characteristic that leads to system error
Human error/mistake: Human action that introduces faults
<!--ID: 1716148052976-->

What makes [equivalence class tests](equivalence_class_tests) strong/weak? what makes them normal/robust? #flashcard 
Strong equivalence class tests test multiple variables, assuming multiple points of failure (multiple fault assumption.)
Weak ECTs test single variables, assuming a single point of failure (single fault assumption.)
Robust ECTs also test [Equivalence Classes (Programming)|ECs](equivalence_classes_(programming)|ecs) that contain invalid inputs, normal ECTs do not.
<!--ID: 1716148647107-->

Why is it important to [Boundary ECT Testing|test boundary values](boundary_ect_testing|test_boundary_values) in [Equivalence Class Tests](equivalence_class_tests)? #flashcard 
Boundary values are important to test since these values are the most likely to cause errors.
<!--ID: 1716148721809-->

What are the 6 benefits of [Unit Testing](unit_testing)? #flashcard 
Takes less time to write them - Faster
Can find bugs early
Can help with improving design decisions
Makes your test documentation more detailed
Encourages better code
Safer in the way that small units are easy to manage
<!--ID: 1716306544389-->


How does `expect()` work in [EasyMock](easymock)? #flashcard 
`expect()` allows us to set the behaviour for our [Mock Objects|mock object](mock_objects|mock_object), we tell it what inputs to expect and what output to return, amongst many other things.
<!--ID: 1716306544452-->


What is the problem with creating [mock objects](mock_objects) by hand? Why is [EasyMock](easymock) so beneficial as a result? #flashcard 
Creating [mock objects](mock_objects) yourself over and over and filling them with the test values would be insanely tedious.
With EasyMock, all you need is the [Interfaces|interface](interfaces|interface) you'll use as a mock, and defining expected behaviour is done with `expect()`! 
<!--ID: 1716306544465-->

Why does [Top-Down Integration](top-down_integration) encourage early prototyping of the application? #flashcard 
Because you start with an abstract, high-level version of the application and refine it over time, which is akin to the process of prototyping.
<!--ID: 1716306743799-->


















#### Why:
#### How:









