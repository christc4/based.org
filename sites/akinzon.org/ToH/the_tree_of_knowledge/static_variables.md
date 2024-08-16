# Static [Variables](variables)

Up: [Statics](statics)
Brother(s):
TARGET DECK

Due to the definition of statics in java, **static variables** therefore act like variables that are global to all instances of a [Classes|class](classes|class), all its [objects](objects) share it.

This is because they all reference the same variable in memory - it is stored in the class' [Method Area](method_area).

The `this` keyword can **NOT** be used to reference static variables in cases where a variable of the same name has been passed to the method, e.g.:

```java
this.var = var;
```

This **ONLY** works for **INSTANCE VARIABLES!!!**

[Static Final Variables](static_final_variables)


































#### Why:
#### How:









