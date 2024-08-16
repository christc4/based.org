# Operation(Method) Syntax

Up: [UML Syntax](uml_syntax)
Brother(s):
TARGET DECK

Operations/[Methods](methods) are formatted like this in UML:

`[visibilityModifier] operationName (parameterList): returnType`

`visibilityModifier`: [Visibility in UML](visibility_in_uml)
`returnType, type`: Boolean, Integer, Real, String, (Other defined [classes](classes))

Each parameter in `parameterList` is defined like this:

`kind name : type = defaultValue` (assuming there is a default value)

There are three possibilities for `kind`:
	 - `in`: method doesn't change parameter, and only takes it as input.
	 - `out`: method doesn't take parameter as input, but changes it.
	 - `inout`: method takes parameter as input and changes it.

For `out` and `inout`, the value is either returned, or its reference changed to the updated value.

Don't spam `inout`, since it's only necessary for taking in and modifying values. Spamming it makes the methods confusing.




































#### Why:
#### How:









