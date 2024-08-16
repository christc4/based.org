# Scope

Up: [Programming Principles](programming_principles)
Brother(s):
TARGET DECK

**Scope** defines which part of a program can interact with a [Variables|variable](variables|variable) and also dictates when and where that variable "exists", it only exists inside its scope.

Member variables exist so long as the containing [Classes|class](classes|class) exists, while local variables only exist in their containing [Methods|method](methods|method).

The rule of thumb for scope is that typically they exist within the $\{\}$ that they were declared in.

>[!TIP]- Member variable example
>![Pasted image 20240110192848.png](pasted_image_20240110192848.png)

>[!TIP]- Local variable example
>![Pasted image 20240110192914.png](pasted_image_20240110192914.png)

Member variables can conflict with local variables if they have the same name, so to ensure you're talking about a member variable, reference it with 
```Java
this.variableName
```




























#### Why:
#### How:









