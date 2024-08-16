# Visibility in UML

Up: [UML Syntax](uml_syntax)
Brother(s):
TARGET DECK

This has a close mapping to [Access Levels for Encapsulation|access modifiers](access_levels_for_encapsulation|access_modifiers) in OOP.

**Private (-)** attributes/[methods](methods) cannot be accessed outside of its [classes|class](classes|class).

**Public (+)** attributes can be accessed anywhere.

**Protected (#)** attributes can only be accessed by the class and those that inherit from it.

In this example:

| Car |
| ---- |
| - mileage<br># parts |
| + getPrice() |
The `mileage` attribute is private, `parts` is protected, and `getPrice()` is public.




































#### Why:
#### How:









