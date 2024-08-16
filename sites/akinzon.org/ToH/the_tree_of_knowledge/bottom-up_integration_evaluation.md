# Bottom-Up Integration Evaluation

Up: [Bottom-Up Integration](bottom-up_integration)
Brother(s):
TARGET DECK

**Advantages:**
	- The step on which faults occur is easy to find since the testing happens step-by-step
	- Few [Stub Objects|stubs](stub_objects|stubs) are used
	- The [low-level modules](low-level_modules) are rigorously tested, lending credence to their reusability
	- Testing is done in parallel with development
	- Major design flaws are found out early on

**Disadvantages:**
	- Needs [Drivers (Testing)|drivers](drivers_(testing)|drivers) since the entry-points to the modules aren't as clear as with [Top-Down Integration](top-down_integration)
	- The logic of the overall solution is tested closer to the end since we're testing the lower-level modules first, and they do not convey the overall logic of the system as well as [High-Level Modules|higher-lever modules](high-level_modules|higher-lever_modules), they simply do a task.



































#### Why:
#### How:









