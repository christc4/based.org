# Top-Down Integration Evaluation

Up: [Top-Down Integration](top-down_integration)
Brother(s):
TARGET DECK


**Advantages:**
	- Because the entry point to the test is likely to be an entry point to the actual app, the [Drivers (Testing)|drivers](drivers_(testing)|drivers) are easy to write.
	- Since the tests are step-by-step, you will know on what step a fault occurs.
	- The major flaws are found out early.

**Disadvantages:**
	- So many [mock and stub objects](mock_and_stub_objects) are needed from the beginning.
	- The [Low-Level Modules|lowest-level modules](low-level_modules|lowest-level_modules) aren't tested in [Testing in Isolation|isolation](testing_in_isolation|isolation), but from within a larger context. This testing may not be rigorous/sufficient enough. Therefore, the low-level module may not be that reusable.


































#### Why:
#### How:









