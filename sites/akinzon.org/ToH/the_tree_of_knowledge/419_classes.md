# 419 [Classes](classes)

Up: [Testing in Isolation](testing_in_isolation)
Brother(s):
TARGET DECK

**Benefits:**
	- We don't actually need to make this class
	- We can assume this fake class is correct so we don't need to consider its correctness for our tests on the class we really care about
	- We don't need to consider the classes that this "class B" relies on
	- We can speed up test execution too, since the fake class' method won't be as computationally expensive

**Drawbacks:**
	- We have to put in extra effort just to make fake classes.

**Solution:**
	- Make fake [subclasses](subclasses) of this "class B"


> [!tip] DID YOU KNOW?: Instances of these classes are known as [Mock and Stub Objects|Mock/Stub Objects](mock_and_stub_objects|mock/stub_objects)? 

[Drivers (Testing)](drivers_(testing))



































#### Why:
#### How:









