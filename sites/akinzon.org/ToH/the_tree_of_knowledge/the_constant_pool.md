# The Constant Pool

Up: [Method Area](method_area)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::The JVM

The constant pool is a partition of the method area for each class loaded into the method area by the [Class Loader Subsystem](class_loader_subsystem).

The Constant Pool contains an ordered list of 'constants' used by that class type.

Any entity ([primitives](primitives)/[classes](classes)) that is part of this class' definition can be found in its constant pool.

These are either stored as literals or references to other classes (that this class depends on.)

Constant Pool Resolution is when symbolic references in the constant pool are replaced with direct pointers to those classes, methods, etc.

since its ordered, I think each constant is popped off when needed, so that it is in $\Theta (1)$ time.

What is [the constant pool](the_constant_pool), what does it store (what are these 'constants') and where is it stored? #flashcard 
The constant pool is an ordered list of 'constants', stored for each class. These 'constants' refer to **every entity referred to in the class' definition**, from literals (what you think of when you say constants) to symbolic references to other [classes](classes) and [methods](methods). It is stored in the [method area](method_area).
<!--ID: 1716322907782-->










































#### Why:
#### How:









