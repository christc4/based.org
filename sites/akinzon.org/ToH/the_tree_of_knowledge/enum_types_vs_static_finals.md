# [Enum Types|Enum Types](enum_types|enum_types) vs [Static Final Variables|Static Finals](static_final_variables|static_finals)

Up: [Enum Types](enum_types)
Brother(s):
TARGET DECK

Let's say we used static finals instead of enum types..

![Pasted image 20240514095625.png](pasted_image_20240514095625.png)

Seems cool, right?

**WRONG!**

cool, we can pass these to methods, but since these are really just `int`s under the hood, we can also pass any `int` variable!

![Pasted image 20240514095636.png](pasted_image_20240514095636.png)
^ Never knew that `-10` was a country...

You could probably define a range for the valid `int`s that refer to a country, but that can clutter code.

Furthermore, these are literally just `int`s, which tells you nothing about that they represent.
































#### Why:
#### How:









