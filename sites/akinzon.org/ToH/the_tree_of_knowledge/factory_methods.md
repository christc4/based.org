# Factory Methods

Up: [Factory Design Pattern](factory_design_pattern)
Brother(s):
TARGET DECK

Factory Methods are [methods](methods) used by [classes](classes) to decide which object is created.

The benefit of having these [Objects|object](objects|object)-creators as methods is that these methods can be [Overriding|overridden](overriding|overridden) by [subclasses](subclasses) to further specify which objects are created, since the subclass may need a specific version of the general class created by the base class' factory method.

The main limit on this is that all created objects must be linked by the same base class (base class B), which is the return type defined in the base class' (base class A's) factory method.



































#### Why:
#### How:









